#!/bin/bash
## Give ISCSI time to intiate
iscsi="1"
while [ $iscsi = "1" ]; do 
	if [ -f /tmp/iscsi.lock ]; then
		iscsi="1"
		sleep 1 
	else
		iscsi="0"
	fi
done

## Primary Disk Mounting Function
data_mount () {
	echo -e "Mounting /dev/$disk to /data$dcount"
	sudo mkdir -p /data$dcount
	sudo mount -o noatime,barrier=1 -t ext4 /dev/$disk /data$dcount
	UUID=`sudo lsblk -no UUID /dev/$disk`
	echo "UUID=$UUID   /data$dcount    ext4   defaults,noatime,discard,barrier=0 0 1" | sudo tee -a /etc/fstab
	dcount=$((dcount+1))
}

## Check for x>0 devices
echo -n "Checking for disks..."
## Execute - will format all devices except sda for use as data disks in HDFS 
n=0
dcount=0
for disk in `sudo cat /proc/partitions | grep -iv sda | sed 1,2d | gawk '{print $4}'`; do
	echo -e "\nProcessing /dev/$disk"
	sudo mke2fs -F -t ext4 -b 4096 -E lazy_itable_init=1 -O sparse_super,dir_index,extent,has_journal,uninit_bg -m1 /dev/$disk
	data_mount
	sudo /sbin/tune2fs -i0 -c0 /dev/$disk
done;
