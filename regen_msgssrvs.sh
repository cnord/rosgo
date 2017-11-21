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

normalize_msgdir() {
	basename "$1" | awk -F: 'BEGIN{deftype="msg"} $1~/_srvs/{deftype="srv"} $2~/^$/{$2=deftype} {print $1 ":" $2}'
}

# Regenerate only choosen message's packs
if [[ "$@" != "" ]]; then
	declare "_msgs_dirs=$(for p in $@; do
		[ -z "$(normalize_msgdir "$p" | awk -F: '{print $1}')" ] && continue
		echo "$p"
	done)"
	if [ -n "$_msgs_dirs" ]; then
		msgs_dirs="$_msgs_dirs"
	else
		echo "The directory list is empty or you set incorrect directory's names" >&2
		exit 1
	fi
fi

for dir in $msgs_dirs; do
  dir="$(normalize_msgdir "$dir")"
  declare "$(echo "$dir"    | awk -F: '{printf "MSGDIR=%s",     $1}')"
  declare "$(echo "$MSGDIR" | awk -F_ '{printf "MSGDIRBASE=%s", $NF}')"
  declare "$(echo "$dir"    | awk -F: '{printf "MSGTYPE=%s",    $2}')"
  echo "package $MSGDIR ..."
  mkdir -p $base_dir/${MSGTYPE}s/${MSGDIR}
  [ -d $share_dir/${MSGDIR}/$MSGTYPE/ ] || continue
  for file in $(find $share_dir/${MSGDIR}/$MSGTYPE/ -mindepth 1 -maxdepth 1 -name "*.${MSGTYPE}"); do
    target=$base_dir/${MSGDIRBASE}/${MSGDIR}/${file##*/}
    ros-gen-go $MSGTYPE --make-copy --package=${MSGDIR} --in=$file --out=$target.go
  done
done
