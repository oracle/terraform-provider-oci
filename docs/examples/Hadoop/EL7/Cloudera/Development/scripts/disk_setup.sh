#!/bin/bash
## Node Disk Setup Script 
## by Zachary Smith (Zachary.Smith@oracle.com)
## Last Update - March 2018
## Give ISCSI time to intiate
sleep 20
## Check for x>0 devices
echo -n "Checking for disks..."
x=0
while [ $x = "0" ]; do 	
	disk_count=`sudo cat /proc/partitions | grep -iv sda | sed 1,2d | gawk '{print $4}' | wc -l`
	if [ $disk_count = "0" ]; then
		sleep 1
		echo -n "."
		continue
	else
	## Execute - will format all devices except sda for prep - if additional storage is put on Master Nodes then this may need modification
	nvme_check=`cat /proc/partitions | grep nvme`
	nvme_chk=`echo -e $?`
	n=0
	for disk in `sudo cat /proc/partitions | grep -iv sda | sed 1,2d | gawk '{print $4}'`; do
		echo -e "\nProcessing /dev/$disk"
		sudo mke2fs -F -t ext4 -b 4096 -E lazy_itable_init=1 -O sparse_super,dir_index,extent,has_journal,uninit_bg -m1 /dev/$disk
		if [ "$nvme_chk" = "1" ]; then
			worker_check=`hostname | grep Worker`
			worker_check=`echo -e $?`
			if [ $worker_check = "0" ]; then  
				sudo mkdir -p /data$n
				sudo mount -o noatime,barrier=1 -t ext4 /dev/$disk /data$n
				UUID=`sudo lsblk -no UUID /dev/$disk`
			 	echo "UUID=$UUID   /data$n    ext4   defaults,noatime,discard,barrier=0 0 1" | sudo tee -a /etc/fstab
			else
				if [ -d "/opt/cloudera" ]; then
	                                sudo mkdir -p /data$n
	                                sudo mount -o noatime,barrier=1 -t ext4 /dev/$disk /data$n
        	                        UUID=`sudo lsblk -no UUID /dev/$disk`
                	                echo "UUID=$UUID   /data$n    ext4   defaults,noatime,discard,barrier=0 0 1" | sudo tee -a /etc/fstab
				else
					sudo mkdir -p /opt/cloudera
					sudo mount -o noatime,barrier=1 -t ext4 /dev/$disk /opt/cloudera
                                        UUID=`sudo lsblk -no UUID /dev/$disk`
                                        echo "UUID=$UUID   /opt/cloudera    ext4   defaults,noatime,discard,barrier=0 0 1" | sudo tee -a /etc/fstab
				fi
			fi
		else
			nvme_disk=`echo -e $disk | grep nvme`
			nvme_disk_chk=`echo -e $?`
			if [ $nvme_disk_chk = "0" ]; then 
	                        sudo mkdir -p /data$n
        	                sudo mount -o noatime,barrier=1 -t ext4 /dev/$disk /data$n
                	        UUID=`sudo lsblk -no UUID /dev/$disk`
                        	echo "UUID=$UUID   /data$n    ext4   defaults,noatime,discard,barrier=0 0 1" | sudo tee -a /etc/fstab
			else
				if [ -d /opt/cloudera ]; then 
					## If cloudera mount point exists, then assume additional disks are for data
	 	                        sudo mkdir -p /data$n
        		                sudo mount -o noatime,barrier=1 -t ext4 /dev/$disk /data$n
                        		UUID=`sudo lsblk -no UUID /dev/$disk`
		                        echo "UUID=$UUID   /data$n    ext4   defaults,noatime,discard,barrier=0 0 1" | sudo tee -a /etc/fstab
				else
					## Mount first non-NVME device as /opt/cloudera
					sudo mkdir -p /opt/cloudera
					sudo mount -o noatime,barrier=1 -t ext4 /dev/$disk /opt/cloudera
					UUID=`sudo lsblk -no UUID /dev/$disk`
					echo "UUID=$UUID   /opt/cloudera    ext4   defaults,noatime,discard,barrier=0 0 1" | sudo tee -a /etc/fstab
				fi
			fi
		fi
		sudo /sbin/tune2fs -i0 -c0 /dev/$disk
		n=$((n+1))
	done;
	x=$disk_count
	fi
done
