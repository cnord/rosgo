package ros

import (
	"bytes"
	"container/list"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"reflect"
	"sync"
	"time"
)

type serviceResult struct {
	srv Service
	err error
}

type remoteClientSessionError struct {
	session *remoteClientSession
	err     error
}

func (e *remoteClientSessionError) Error() string {
	return fmt.Sprintf("remoteClientSession %v error: %v", e.session, e.err)
}

type defaultServiceServer struct {
	node             *defaultNode
	service          string
	srvType          ServiceType
	handler          interface{}
	listener         *net.TCPListener
	sessions         *list.List
	sessLock         sync.Mutex
	shutdownChan     chan struct{}
	sessionErrorChan chan error
}

var (
	normalExitErr = fmt.Errorf("Normal exit")
)

func newDefaultServiceServer(node *defaultNode, service string, srvType ServiceType, handler interface{}) *defaultServiceServer {
	logger := node.logger
	server := new(defaultServiceServer)
	if listener, err := listenRandomPort("127.0.0.1", 10); err != nil {
		panic(err)
	} else {
		if tcpListener, ok := listener.(*net.TCPListener); ok {
			server.listener = tcpListener
		} else {
			panic(fmt.Errorf("Server listener is not TCPListener"))
		}
	}
	server.node = node
	server.service = service
	server.srvType = srvType
	server.handler = handler
	server.sessions = list.New()
	server.shutdownChan = make(chan struct{}, 10)
	server.sessionErrorChan = make(chan error, 10)
	address := fmt.Sprintf("rosrpc://%s", server.listener.Addr().String())
	logger.Debugf("ServiceServer listen %s", address)
	_, err := callRosApi(node.masterUri, "registerService",
		node.qualifiedName,
		service,
		address,
		node.xmlrpcUri)
	if err != nil {
		logger.Errorf("Failed to register service %s", service)
		server.listener.Close()
		return nil
	}
	go server.start()
	return server
}

func (s *defaultServiceServer) Shutdown() {
	s.shutdownChan <- struct{}{}
}

// event loop
func (s *defaultServiceServer) start() {
	logger := s.node.logger
	logger.Debugf("service server '%s' start listen %s.", s.service, s.listener.Addr().String())
	s.node.waitGroup.Add(1)
	defer func() {
		logger.Debug("defaultServiceServer.start exit")
		s.node.waitGroup.Done()
	}()

	go s.chanLoop()

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			logger.Debugf("s.listener.Accept() failed: %v", err)
			return
		}
		logger.Debugf("Connected from %s", conn.RemoteAddr().String())
		session := newRemoteClientSession(s, conn)
		s.sessLock.Lock()
		s.sessions.PushBack(session)
		s.sessLock.Unlock()
		go session.start()
	}
}

func (s *defaultServiceServer) chanLoop() {
	logger := s.node.logger
	for {
		select {
		case err := <-s.sessionErrorChan:
			if sessionError, ok := err.(*remoteClientSessionError); ok {
				if sessionError.err != normalExitErr {
					logger.Errorf("session error: %v", err)
				}
				s.sessLock.Lock()
				for e := s.sessions.Front(); e != nil; e = e.Next() {
					if e.Value == sessionError.session {
						logger.Debugf("service session %v removed", e.Value)
						s.sessions.Remove(e)
					}
				}
				s.sessLock.Unlock()
			} else {
				logger.Errorf("session error: %v", err)
			}
		case <-s.shutdownChan:
			logger.Debug("defaultServiceServer.start Receive shutdownChan")
			s.listener.Close()
			logger.Debug("defaultServiceServer.start closed listener")
			_, err := callRosApi(s.node.masterUri, "unregisterService",
				s.node.qualifiedName, s.service, s.node.xmlrpcUri)
			if err != nil {
				logger.Warn("Failed unregisterService(%s): %v", s.service, err)
			}
			logger.Debug("Called unregisterService(%s)", s.service)
			s.sessLock.Lock()
			for e := s.sessions.Front(); e != nil; e = e.Next() {
				session := e.Value.(*remoteClientSession)
				session.quitChan <- struct{}{}
			}
			s.sessions.Init() // Clear all sessions
			s.sessLock.Unlock()
			logger.Debug("defaultServiceServer.start session cleared")
			return
		}
	}
}

type remoteClientSession struct {
	server       *defaultServiceServer
	conn         net.Conn
	quitChan     chan struct{}
	responseChan chan []byte
	errorChan    chan error
}

func newRemoteClientSession(s *defaultServiceServer, conn net.Conn) *remoteClientSession {
	session := new(remoteClientSession)
	session.server = s
	session.conn = conn
	session.responseChan = make(chan []byte, 1)
	session.errorChan = make(chan error, 1)
	return session
}

func (s *remoteClientSession) start() {
	logger := s.server.node.logger
	conn := s.conn
	nodeId := s.server.node.qualifiedName
	service := s.server.service
	md5sum := s.server.srvType.MD5Sum()
	srvType := s.server.srvType.Name()
	var err error
	logger.Debugf("remoteClientSession.start '%s'", s.server.service)
	defer func() {
		logger.Debug("remoteClientSession.start exit")
	}()
	defer func() {
		if err := recover(); err != nil {
			if e, ok := err.(error); ok {
				s.server.sessionErrorChan <- &remoteClientSessionError{s, e}
			} else {
				e = fmt.Errorf("Unkonwn error value")
				s.server.sessionErrorChan <- &remoteClientSessionError{s, e}
			}
		} else {
			e := normalExitErr
			s.server.sessionErrorChan <- &remoteClientSessionError{s, e}
		}
	}()

	// 1. Read request header
	isProbeReq := false
	conn.SetDeadline(time.Now().Add(10 * time.Millisecond))
	if resHeaders, readErr := readConnectionHeader(conn); readErr != nil {
		panic(readErr)
	} else {
		logger.Debug("TCPROS Connection Header:")
		resHeaderMap := make(map[string]string)
		for _, h := range resHeaders {
			resHeaderMap[h.key] = h.value
			logger.Debugf("  `%s` = `%s`", h.key, h.value)
		}
		if probe, ok := resHeaderMap["probe"]; ok && probe == "1" {
			logger.Debug("TCPROS header 'probe' detected.")
			isProbeReq = true
		} else {
			if resHeaderMap["service"] != service ||
				resHeaderMap["md5sum"] != md5sum {
				logger.Fatalf("Incompatible message type: %s(%s) != %s(%s)",
					resHeaderMap["service"], resHeaderMap["md5sum"],
					service, md5sum)
			}
		}
	}

	// 2. Write response header
	var headers []header
	headers = append(headers, header{"service", service})
	headers = append(headers, header{"md5sum", md5sum})
	headers = append(headers, header{"type", srvType})
	headers = append(headers, header{"callerid", nodeId})
	logger.Debug("TCPROS Response Header")
	for _, h := range headers {
		logger.Debugf("  `%s` = `%s`", h.key, h.value)
	}
	conn.SetDeadline(time.Now().Add(10 * time.Millisecond))
	if err = writeConnectionHeader(headers, conn); err != nil {
		panic(err)
	}
	if isProbeReq {
		return
	}

	// 3. Read request
	logger.Debug("Reading message size...")
	var msgSize uint32
	conn.SetDeadline(time.Now().Add(10 * time.Millisecond))
	if err = binary.Read(conn, binary.LittleEndian, &msgSize); err != nil {
		panic(err)
	}
	logger.Debugf("  %d", msgSize)
	resBuffer := make([]byte, int(msgSize))
	logger.Debug("Reading message body...")
	conn.SetDeadline(time.Now().Add(10 * time.Millisecond))
	if _, err = io.ReadFull(conn, resBuffer); err != nil {
		panic(err)
	}

	s.server.node.jobChan <- func() {
		srv := s.server.srvType.NewService()
		reader := bytes.NewReader(resBuffer)
		err = srv.ReqMessage().Deserialize(reader)
		if err != nil {
			s.errorChan <- err
		}
		args := []reflect.Value{reflect.ValueOf(srv)}
		fun := reflect.ValueOf(s.server.handler)
		results := fun.Call(args)

		if len(results) != 1 {
			logger.Debug("Service callback return type must be 'error'")
			s.errorChan <- err
			return
		}
		result := results[0]
		if result.IsNil() {
			logger.Debug("Service callback success")
			var buf bytes.Buffer
			_ = srv.ResMessage().Serialize(&buf)
			s.responseChan <- buf.Bytes()
		} else {
			logger.Debug("Service callback failure")
			var ok bool
			if err, ok = result.Interface().(error); ok {
				s.errorChan <- err
			} else {
				s.errorChan <- fmt.Errorf("Service handler has invalid signature")
			}
		}
	}

	timeoutChan := time.After(1000 * time.Millisecond)
	select {
	case resMsg := <-s.responseChan:
		// 4. Write OK byte
		var ok byte = 1
		conn.SetDeadline(time.Now().Add(10 * time.Millisecond))
		if err = binary.Write(conn, binary.LittleEndian, &ok); err != nil {
			panic(err)
		}
		// 5. Write response
		logger.Debug(len(resMsg))
		size := uint32(len(resMsg))
		conn.SetDeadline(time.Now().Add(10 * time.Millisecond))
		if err = binary.Write(conn, binary.LittleEndian, size); err != nil {
			panic(err)
		}
		conn.SetDeadline(time.Now().Add(10 * time.Millisecond))
		if _, err = conn.Write(resMsg); err != nil {
			panic(err)
		}
	case err = <-s.errorChan:
		logger.Error(err)
		// 4. Write OK byte
		var ok byte = 0
		conn.SetDeadline(time.Now().Add(10 * time.Millisecond))
		if err = binary.Write(conn, binary.LittleEndian, &ok); err != nil {
			panic(err)
		}
		errMsg := err.Error()
		size := uint32(len(errMsg))
		conn.SetDeadline(time.Now().Add(10 * time.Millisecond))
		if err = binary.Write(conn, binary.LittleEndian, size); err != nil {
			panic(err)
		}
		conn.SetDeadline(time.Now().Add(10 * time.Millisecond))
		if _, err = conn.Write([]byte(errMsg)); err != nil {
			panic(err)
		}
	case <-timeoutChan:
		panic(fmt.Errorf("service callback timeout"))
	}
}
