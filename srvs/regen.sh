#!/bin/bash -e

share_dir="${ROS_ROOT}/.."
srvs_dirs='
	std_srvs
	mavros_msgs
'

for dir in $srvs_dirs; do
  echo "package $dir ..."
  mkdir -p $dir
  for file in $(find $share_dir/$dir/srv/ -maxdepth 1 -type f -name '*.srv'); do
		target=$dir/${file##*/}
    cp $file $target
    ros-gen-go srv --package=$dir --in=$file --out=$target.go
	done
done
