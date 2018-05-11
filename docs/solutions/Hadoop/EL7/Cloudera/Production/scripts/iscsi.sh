#!/bin/bash
## ISCSI initiator script
## by Zachary Smith (Zachary.Smith@oracle.com)
## Last Update - March 2018

#Look for all ISCSI devices in parallel
touch /tmp/iscsi.lock
for i in `seq 2 32`; do 
        wait="0"
                while [ $wait = "0" ]; do
                        iscsi_ct=`ps -ef | grep iscsiadm | grep -iv grep | wc -l`
                        if [ $iscsi_ct -lt 5 ]; then
                                sudo iscsiadm -m discoverydb -D -t sendtargets -p 169.254.2.$i:3260 &
                                wait="1"
                        else
                                sleep 1
                        fi
                done
done;
sleep 10 
sudo iscsiadm -m node -l
sudo iscsiadm -m node -n node.startup -v automatic
rm -f /tmp/iscsi.lock
