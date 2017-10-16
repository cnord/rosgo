#!/bin/bash

BASEDIR="$(realpath "$(dirname $BASH_SOURCE)")"
PACKAGE="$(awk '/^package /{print $2}' "$BASEDIR/main.go")"

find "$BASEDIR" -mindepth 1 -maxdepth 1 -type f -iname "*.tmpl" | while read TMPLNAME; do
	TMPLNAME="$(basename "$TMPLNAME")"
	TMPLCAMEL="$(echo "$TMPLNAME" | sed 's/\.\([a-z]\)/\U\1/g')"
	cat > "$BASEDIR/$TMPLNAME.go" << EOFEOF
package $PACKAGE

import (
	"os"
	"time"
)

func ${TMPLCAMEL}Bytes() ([]byte, error) {
	return bindataRead(
		_${TMPLCAMEL},
		"${TMPLNAME}",
	)
}

func ${TMPLCAMEL}() (*asset, error) {
	bytes, err := ${TMPLCAMEL}Bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "", size: int64(len(bytes)), mode: os.FileMode(436), modTime: time.Unix(1507566040, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _${TMPLCAMEL} = []byte{
$(gzip -c ${TMPLNAME} | xxd -i) }

EOFEOF

done
