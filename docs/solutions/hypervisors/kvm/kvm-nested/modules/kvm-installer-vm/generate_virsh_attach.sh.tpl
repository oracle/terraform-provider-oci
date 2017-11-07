#!/bin/bash

cat > generate_virsh_attach.sh <<SCRIPT
#!/bin/bash

# An error exit function
function error_exit {
	echo "\$1" 1>&2
	exit 1
}

IP=\$(which ip)

# Name of the network interface of the instance for the following mac address
intfName=\$(\$IP -o link  | awk '/${lower(vnic_mac_address)}/{print substr(\$2, 1, length(\$2)-1)}')
[ \$? -eq 0 ] || error_exit "Unable to get network interface name"



cat > attach.xml << EOF
<interface type='direct'>
<mac address='${vnic_mac_address}'/>
<source dev='\$intfName' mode='passthrough'/>
<target dev='macvtap1'/>
<alias name='net1'/>
<model type='${emulation_model}'/>
</interface>
EOF
SCRIPT
