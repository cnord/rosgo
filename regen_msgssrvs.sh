#!/bin/bash -e

share_dir="$(dirname "${ROS_ROOT}")"
base_dir="$(dirname "$BASH_SOURCE")"
# uuid_msgs is a dependecy for geographic_msgs
msgs_dirs='
	actionlib_msgs
	common_msgs
	control_msgs
	diagnostic_msgs
	geometry_msgs
	uuid_msgs
	geographic_msgs
	map_msgs
	mavros_msgs
	mavros_msgs:srv
	nav_msgs
	rosgraph_msgs
	sensor_msgs
	shape_msgs
	smach_msgs
	std_msgs
	std_srvs:srv
	stereo_msgs
	tf2_msgs
	trajectory_msgs
	visualization_msgs
'

# Regenerate only choosen messages
if [[ "$@" != "" ]]; then
	declare "_msgs_dirs=$(for p in $@; do
		dir="$(dirname $(realpath $p))"
		ptype="$(echo "$(basename $p)" | awk -F: 'BEGIN{deftype="msg"} $1~/_srvs/{deftype="srv"} {if ($2=="") {print deftype} else {print $2}}')"
		name="$(echo "$(basename $p)"  | awk -F: '{print $1}')"
		[ -z "$name" ] && continue
		# there is no checking of exists folder due to it could be a first time call
		echo "$name:$ptype"
	done)"
	if [ -n "$_msgs_dirs" ]; then
		msgs_dirs="$_msgs_dirs"
	fi
fi

for dir in $msgs_dirs; do
  declare "$(echo "$dir" | awk -F: '{printf "MSGDIR=%s", $1; }')"
  declare "$(echo "$dir" | sed "s/^.*_//" | awk -F: '{printf "MSGDIRBASE=%s", $1; }')"
  declare "$(echo "$dir" | awk -F: '{msgtype=$2; if ($2=="") { msgtype="msg" }; printf "MSGTYPE=%s", msgtype; }')"
  echo "package $MSGDIR ..."
  mkdir -p $base_dir/${MSGTYPE}s/${MSGDIR}
  [ -d $share_dir/${MSGDIR}/$MSGTYPE/ ] || continue
  for file in $(find $share_dir/${MSGDIR}/$MSGTYPE/ -mindepth 1 -maxdepth 1 -name "*.${MSGTYPE}"); do
    target=$base_dir/${MSGDIRBASE}/${MSGDIR}/${file##*/}
    cp $file $target
    ros-gen-go $MSGTYPE --package=${MSGDIR} --in=$file --out=$target.go
	done
done
