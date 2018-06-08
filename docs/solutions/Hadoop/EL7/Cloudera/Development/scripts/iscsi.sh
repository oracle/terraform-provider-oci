#!/bin/bash

#Look for all ISCSI devices in sequence, finish on first failure
touch /tmp/iscsi.lock
v="0"
done="0"
echo -e "Mapping Block Volumes...."
for i in `seq 2 32`; do
        if [ $done = "0" ]; then
                sudo iscsiadm -m discoverydb -D -t sendtargets -p 169.254.2.$i:3260 2>&1 2>/dev/null
                iscsi_chk=`echo -e $?`
                if [ $iscsi_chk = "0" ]; then
                        echo -e "Success for volume $((i-1))."
                        v=$((v+1))
                        continue
                else
                        echo -e "Completed - $((i-2)) volumes found."
                        done="1"
                fi
        fi
done;
if [ $v -gt 0 ]; then
	echo -e "Setting auto-startup for volumes."
        sudo iscsiadm -m node -l
        sudo iscsiadm -m node -n node.startup -v automatic
fi
rm -f /tmp/iscsi.lock
