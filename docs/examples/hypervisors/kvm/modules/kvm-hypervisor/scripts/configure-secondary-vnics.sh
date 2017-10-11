#!/bin/bash -eu

# Configure a secondary VNIC on a bare metal instance
# Prerequisite: Add a secondary VNIC to the instance using Console/SDK/API.

# An error exit function
function error_exit {
	echo "$1" 1>&2
	exit 1
}

function usage {
cat <<EOF
Usage: $0 [-n namespace] [-h]
    This script configures a secondary VNIC on a bare metal instance.
    It shows the list of available secondary VNICs and configures the one you select.
    [-n namespace] - Optional parameter to specify the namespace in which the VNIC needs to be configured.
                     It is configured in the default namespace if this option is not specified.
    [-h] - Shows the usage information
EOF

}

#Initialize NS to empty string. Secondary VNIC will be configured in the default namespace by default.
NS=

# Parse command line arguments
while getopts ":hn:" opt; do
    case $opt in
        h)
            usage
            exit 0
            ;;
        n)
            NS=$OPTARG
            ;;
        \?)
            echo "Invalid option: -$OPTARG" >&2
            usage
            exit 1
            ;;
        :)
            echo "Option -$OPTARG requires an argument." >&2
            exit 1
            ;;
    esac
done

# Make sure the script is run as root
[[ $EUID -eq 0 ]] || error_exit "This script must be run as root"

# Load 8021q module for vlan tagging support: https://en.wikipedia.org/wiki/IEEE_802.1Q
modprobe 8021q || error_exit "Failed to load 8021q module"

IMDSURL="http://169.254.169.254/opc/v1/vnics/"

# Get list of private IPs for all VNICs configured on the instance, except the first one in the list.
# Note: the first VNIC is the primary VNIC, which is already configured on the instance
# Note: jq is not available by default on Oracle Linux images. Using Python to avoid dependency on Internet connectivity.
options=($(curl -s $IMDSURL | python -c 'import sys, json; print "\n".join([vnic["privateIp"] for vnic in json.load(sys.stdin)])' | tail -n +2))

# Sanity check that $options isn't empty
! [ -z $options ] || error_exit "Unable to list the private IPs from $IMDSURL"

# Show the list of private IPs and let the user select the index of the IP to be configured on the instance.

COUNTER=1
while [ $COUNTER -lt $((${#options[@]} + 1 )) ]; do
		echo "The counter is $COUNTER"
		index=$COUNTER

		let COUNTER=COUNTER+1

		CURL=$(which curl)
		ROUTE=$(which route)
		IP=$(which ip)
		DHCLIENT=$(which dhclient)
		SSHD=$(which sshd)

		httpCode=$(curl -sL -w "%{http_code}\\n" $IMDSURL/$index/vlanTag -o /dev/null)
		[[ $httpCode == 200 ]] || error_exit "Failed to get vlan tag for selected VNIC from IMDS"

		vlanTag=$($CURL -s $IMDSURL/$index/vlanTag)
		[[ $vlanTag =~ ^[0-9]+$ ]] || error_exit "Invalid format of vlan Tag"

		httpCode=$(curl -sL -w "%{http_code}\\n" $IMDSURL/$index/macAddr -o /dev/null)
		[[ $httpCode == 200 ]] || error_exit "Failed to get macAddr for selected VNIC from IMDS"
		macAddress=$($CURL -s $IMDSURL/$index/macAddr)

		httpCode=$(curl -sL -w "%{http_code}\\n" $IMDSURL/$index/privateIp -o /dev/null)
		[[ $httpCode == 200 ]] || error_exit "Failed to get privateIp for selected VNIC from IMDS"
		privateIp=$($CURL -s $IMDSURL/$index/privateIp)

		httpCode=$(curl -sL -w "%{http_code}\\n" $IMDSURL/$index/virtualRouterIp -o /dev/null)
		[[ $httpCode == 200 ]] || error_exit "Failed to get virtualRouterIp for selected VNIC from IMDS"
		virtualRouterIp=$($CURL -s $IMDSURL/$index/virtualRouterIp)

		httpCode=$(curl -sL -w "%{http_code}\\n" $IMDSURL/$index/subnetCidrBlock -o /dev/null)
		[[ $httpCode == 200 ]] || error_exit "Failed to get subnetCidrBlock for selected VNIC from IMDS"
		subnetCidrBlock=$($CURL -s $IMDSURL/$index/subnetCidrBlock)
		subnetCidrPrefix=$(echo $subnetCidrBlock | awk -F/ '{print $2}')

		echo "Configuring the interface with the IP address: $privateIp"

		# Name of the primary network interface of the instance
		primaryIntfName=$($ROUTE -n | grep '^0.0.0.0' | grep -o '[^ ]*$')
		[ $? -eq 0 ] || error_exit "Unable to get primary network interface name"

		# Name of the new interface.
		# The naming format below is only a convention and can be changed as required
		macVlanIntf="$primaryIntfName.macv.$vlanTag"
		vlanIntf="$primaryIntfName.vlan.$vlanTag"

		# Create a MAC VLAN with provided MAC or random MAC
		$IP link add link $primaryIntfName $macVlanIntf address $macAddress type macvlan \
		  || error_exit "Failed to create a MAC VLAN with provided MAC or random MAC"

		# Create a VLAN on top of the MAC VLAN
		$IP link add link $macVlanIntf name $vlanIntf type vlan id $vlanTag \
		  || error_exit "Failed to create a VLAN on top of the MAC VLAN"

		NSPREFIX=
		if ! [ -z $NS ]; then
		   # Configure the secondary VNIC in namespace $NS
		    NSPREFIX="$IP netns exec $NS"

		    # Create a new namespace
		    $IP netns add $NS \
			  || error_exit "Failed to create a new namespace"

		    # Move the MACVLAN interface and VLAN interface to the namespace
		    $IP link set dev $macVlanIntf netns $NS \
		      || error_exit "Failed to move the MACVLAN interface to the namespace"

		    $IP link set dev $vlanIntf netns $NS \
		      || error_exit "Failed to move the VLAN interface to the namespace"

		    # Start SSHD daemon in the namespace
		    $NSPREFIX $SSHD \
		      || error_exit "Failed to start SSH daemon in the namespace"
		fi

		# Bring the MACVLAN interface and VLAN interface up
		$NSPREFIX $IP link set dev $macVlanIntf up \
		  || error_exit "Failed to bring the MACVLAN interface up"

		$NSPREFIX $IP link set mtu 9000 dev $vlanIntf up \
		  || error_exit "Failed to bring the VLAN interface up"

		# Configure the IP address of the VLAN interface
		$NSPREFIX $IP addr add $privateIp/$subnetCidrPrefix dev $vlanIntf \
		  || error_exit "Failed to configure the IP address of the VLAN interface"

		if ! [ -z $NS ]; then
		    # Set the default gateway for the secondary VNIC in the network namespace
		    $NSPREFIX $ROUTE add default gw $virtualRouterIp \
			  || error_exit "Failed to set the default gateway for the secondary VNIC in the network namespace"
		else
		    # Create custom route rules for this interface
		    rtId=`expr 100 + $vlanTag`
		    rtName=$vlanIntf
		    RTFILE=/etc/iproute2/rt_tables

		    # Check if the route table exists
		    if grep -qs "^$rtId " $RTFILE; then
		        echo "Route table entry exists for this vlanTag. No route rules configured."
		    else
		        echo "$rtId    $rtName" >> $RTFILE
		        # Add a rule to lookup $rtName route table for any packets sourced from $privateIp
		        $IP rule add from $privateIp lookup $rtName \
		          || error_exit "Failed to add a rule to lookup $rtName route table for any packets sourced from $privateIp"

		        # Add a routing rule in $rtName route table to use $virtualRouterIp as the default gateway
		        $IP route add default via $virtualRouterIp dev $vlanIntf table $rtName \
		          || error_exit "Failed to add a routing rule in $rtName route table to use $virtualRouterIp as the default gateway"
		    fi
		fi

done
