#!/bin/bash
## ISCSI initiator script
## by Zachary Smith (Zachary.Smith@oracle.com)
## Last Update - March 2018

#Look for all ISCSI devices in parallel
for i in `seq 1 254`; do 
	sudo iscsiadm -m discoverydb -D -t sendtargets -p 169.254.2.$i:3260 &
done;
sleep 10 
sudo iscsiadm -m node -l
sudo iscsiadm -m node -n node.startup -v automatic
