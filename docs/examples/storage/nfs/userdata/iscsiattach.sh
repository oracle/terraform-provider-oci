#!/bin/bash
# iscsiattach.sh - Scan and automatically attach new iSCSI targets
#
# Author: Steven B. Nelson, Sr. Solutions Architect
#       Oracle Cloud Infrastructure
#
# 20 April 2017
# Copyright Oracle, Inc.  All rights reserved.

BASEADDR="169.254.2.2"

# Set a base address incrementor so we can loop through all the
# addresses.
addrCount=0

while [ ${addrCount} -le 32 ]
do

	CURRADDR=`echo ${BASEADDR} | awk -F\. '{last=$4+'${addrCount}';print $1"."$2"."$3"."last}'`

	clear
	echo "Attempting connection to ${CURRADDR}"

	mkfifo discpipe
	# Find all the iSCSI Block Storage volumes attached to the instance but
	# not configured for use on the instance.  Basically, get a list of the
	# volumes that the instance can see, the loop through the ones it has,
	# and add volumes not already configured on the instance.
	#
	# First get the list of volumes visible (attached) to the instance

	iscsiadm -m discovery -t st -p ${CURRADDR}:3260 | grep -v uefi | awk '{print $2}' > discpipe 2> /dev/null &

	# If the result is non-zero, that generally means that there are no targets available or
	# that the portal is reachable but not active.  We make no distinction between the two
	# and simply skip ahead.
	result=$?
	if [ ${result} -ne 0 ]
	then
		(( addrCount = addrCount + 1 ))
		continue
	fi
	 
	# Loop through the list (via the named FIFO pipe below)
	while read target
	do
	    mkfifo sesspipe
	    # Get the list of the currently attached Block Storage volumes
	    iscsiadm -m session -P 0 | grep -v uefi | awk '{print $4}' > sesspipe 2> /dev/null &
	     
	    # Set a flag, and loop through the sessions (attached, but not configured)
	    # and see if the volumes match.  If so, skip to the next until we get
	    # through the list.  Session list is via the pipe.
	    found="false"
	    while read session
	    do
		if [ ${target} = ${session} ]
		then
		    found="true"
		    break
		fi
	    done < sesspipe
	     
	    # If the volume is not found, configure it.  Get the resulting device file.
	    if [ ${found} = "false" ]
	    then
		iscsiadm -m node -o new -T ${target} -p ${CURRADDR}:3260
		iscsiadm -m node -o update -T ${target} -n node.startup -v automatic
		iscsiadm -m node -T ${target} -p ${CURRADDR}:3260 -l
		sleep 10
	    fi
	done < discpipe
	
	(( addrCount = addrCount + 1 ))
	find . -maxdepth 1 -type p -exec rm {} \;
done
echo "Scan Complete."