#!/bin/bash

cat > generate_virsh_attach.sh <<SCRIPT
#!/bin/bash

# An error exit function
function error_exit {
	echo "\$1" 1>&2
	exit 1
}

ROUTE=\$(which route)
CURL=\$(which curl)

IMDSURL="http://169.254.169.254/opc/v1/vnics/"


httpCode=\$(\$CURL -sL -w "%{http_code}\\n" \$IMDSURL/ -o /dev/null)
[[ \$httpCode == 200 ]] || error_exit "Failed to get list of VNICs from IMDS"


jsonPayload=\$(\$CURL -s \$IMDSURL)

echo \$jsonPayload > myfile.json

vlanTag=\$(echo \$jsonPayload | jq '.[]? | select(.vnicId | contains("${vnic_id}") ) .vlanTag')
[[ \$vlanTag =~ ^[0-9]+\$ ]] || error_exit "Invalid format of vlan Tag"


# Name of the primary network interface of the instance
primaryIntfName=\$(\$ROUTE -n | grep '^0.0.0.0' | grep -o '[^ ]*\$')
[ \$? -eq 0 ] || error_exit "Unable to get primary network interface name"

macVlanIntf="\$primaryIntfName.vlan.\$vlanTag"

cat > attach.xml << EOF
<interface type='direct'>
<mac address='${vnic_mac_address}'/>
<source dev='\$macVlanIntf' mode='passthrough'/>
<target dev='macvtap1'/>
<alias name='net1'/>
<model type='${emulation_model}'/>
</interface>
EOF
SCRIPT
