#!/bin/bash
## Node Disk Setup Script 
## by Zachary Smith (Zachary.Smith@oracle.com)
## Last Update - March 2018
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

data_mount () {
	echo -e "Mounting /dev/$disk to /data$dcount"
	sudo mkdir -p /data$dcount
	sudo mount -o noatime,barrier=1 -t ext4 /dev/$disk /data$dcount
	UUID=`sudo lsblk -no UUID /dev/$disk`
	echo "UUID=$UUID   /data$dcount    ext4   defaults,noatime,discard,barrier=0 0 1" | sudo tee -a /etc/fstab
	dcount=$((dcount+1))
}

parcel_mount () { 
	echo -e "Mounting /dev/$disk to /opt/cloudera"
	sudo mkdir -p /opt/cloudera
	sudo mount -o noatime,barrier=1 -t ext4 /dev/$disk /opt/cloudera
	UUID=`sudo lsblk -no UUID /dev/$disk`
	echo "UUID=$UUID   /opt/cloudera    ext4   defaults,noatime,discard,barrier=0 0 1" | sudo tee -a /etc/fstab
}

## This goes away with Custom Boot Volumes in TF via API
tmp_mount () {
	echo -e "Mounting /dev/$disk to /tmp"
	sudo mount -o noatime,barrier=1 -t ext4 /dev/$disk /tmp
	UUID=`sudo lsblk -no UUID /dev/$disk`
	echo "UUID=$UUID   /tmp    ext4   defaults,noatime,discard,barrier=0 0 1" | sudo tee -a /etc/fstab
	chmod 0777 /tmp/
}

## Check for x>0 devices
echo -n "Checking for disks..."
## Execute - will format all devices except sda for prep - if additional storage is put on Master Nodes then this may need modification
nvme_check=`cat /proc/partitions | grep nvme`
nvme_chk=`echo -e $?`
n=0
tmpdisk=0
dcount=0
for disk in `sudo cat /proc/partitions | grep -iv sda | sed 1,2d | gawk '{print $4}'`; do
	disk_size=`sudo cat /proc/partitions | grep $disk | gawk '{print $3}'`
	echo -e "\nProcessing /dev/$disk"
	sudo mke2fs -F -t ext4 -b 4096 -E lazy_itable_init=1 -O sparse_super,dir_index,extent,has_journal,uninit_bg -m1 /dev/$disk
	if [ "$nvme_chk" = "1" ]; then
		if [ -d /opt/cloudera ]; then 	
			if [ $disk_size -lt "268435456" ]; then
				if [ $tmpdisk = "0" ]; then
					## Assume non-root disk under 256GB is used for TMP
					tmp_mount
					tmpdisk="1"
				else
					echo -e "---------------------------------------------------------------"
					echo -e "--- ERROR /dev/$disk under 256GB should not be HDFS target ----"
					echo -e "---------------------------------------------------------------"
				fi
			else
				## Assume any additional disks are for HDFS
               		        data_mount
			fi
		else
			if [ $disk_size -eq "268435456" ]; then
				## Assume first 256GB disk is used for CDH parcels
				parcel_mount
			elif [ $disk_size -lt "268435456" ]; then 
				tmp_mount
				tmpdisk="1"
			else
				data_mount
			fi
		fi
	else
		nvme_disk=`echo -e $disk | grep nvme`
		nvme_disk_chk=`echo -e $?`
		if [ $nvme_disk_chk = "0" ]; then 
			data_mount
		else
			if [ -d /opt/cloudera ]; then
				# If cloudera mount point exists, then assume additional disks are for data
				data_mount
			else
				## Mount first non-NVME device as /opt/cloudera
				parcel_mount
			fi
		fi
	fi
	sudo /sbin/tune2fs -i0 -c0 /dev/$disk
done;
