#!/bin/bash
# iscsiattach.sh - Scan and automatically attach new iSCSI targets
#
# Author: Steven B. Nelson, Sr. Solutions Architect
#       Oracle Bare Metal Cloud Services
#
# 20 April 2017
# Copyright Oracle, Inc.  All rights reserved.

# Make FIFO pipes for the two loops below
mkfifo discpipe
mkfifo sesspipe

# Set the address ranges based on the Block Storage version
V1ADDR="169.254.0.2"
V2ADDR="169.254.2.0"

# Set the block storage version
BSV="v1"

# If the BSV is v2, we need to scan all 254 addresses, otherwise,
# we scan 1. :-(

if [ ${BSV} = "v2" ]
then
	numAddrs=254
	BASEADDR=${V2ADDR}
else
	numAddrs=3
	BASEADDR=${V1ADDR}
fi

# Set a base address incrementor so we can loop through all the
# addresses.
addrCount=0

echo "Scanning "${numAddrs}" for new targets.  Stand by."
while [ ${addrCount} -le ${numAddrs} ]
do
	# Set the current address to attempt to attach.
	if [ ${BSV} = "v2" ]
	then
		CURRADDR=`echo ${BASEADDR} | awk -F\. '{
last=$4+'${addrCount}'
print $1"."$2"."$3"."last
}'`
	else
		CURRADDR=`echo ${BASEADDR} | awk -F\. '{
last=$3+'${addrCount}'
print $1"."$2"."last"."$4
}'`
	fi

	# We use ping to see if the target is even there.
	# Skip to the next address if we cant ping it.
	ping -q -c 1 -W 1 ${CURRADDR} > /dev/null 2>&1
	result=$?
	if [ ${result} -ne 0 ]
	then
		(( addrCount = addrCount + 1 ))
		continue
	fi

	echo "Connecting to "${CURRADDR}
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
done
echo "Scan Complete."

# Remove the FIFOs
find . -maxdepth 1 -type p -exec rm {} \;
