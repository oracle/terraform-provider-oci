variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}
variable "region" {}
variable "ssh_public_key" {}
variable "ssh_private_key" {}

# Choose an Availability Domain
variable "AD" {
    default = "1"
}

variable "InstanceShape" {
    default = "VM.Standard1.2"
}

variable "InstanceShape2" {
    default = "VM.Standard1.1"
}

variable "InstanceImageOCID" {
    type = "map"
    default = {
        // See https://docs.us-phoenix-1.oraclecloud.com/images/
        // Oracle-provided image "Oracle-Linux-7.4-2018.02.21-1"
        us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaaupbfz5f5hdvejulmalhyb6goieolullgkpumorbvxlwkaowglslq"
        us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaajlw3xfie2t5t52uegyhiq2npx7bqyu4uvi2zyu3w3mqayc2bxmaa"
        eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaa7d3fsb6272srnftyi4dphdgfjf6gurxqhmv6ileds7ba3m2gltxq"
        uk-london-1 = "ocid1.image.oc1.uk-london-1.aaaaaaaaa6h6gj6v4n56mqrbgnosskq63blyv2752g36zerymy63cfkojiiq"
    }
}

# The First VCN
variable "vcn_cidr" {
    default = "10.0.0.0/16"
}

variable "mgmt_subnet_cidr" {
    default = "10.0.0.0/24"
}

variable "private_subnet_cidr" {
    default = "10.0.1.0/24"
}

# The Second VCN
variable "vcn_cidr2" {
    default = "10.1.0.0/16"
}

variable "mgmt_subnet_cidr2" {
    default = "10.1.0.0/24"
}

variable "private_subnet_cidr2" {
    default = "10.1.1.0/24"
}

provider "oci" {
    tenancy_ocid = "${var.tenancy_ocid}"
    user_ocid = "${var.user_ocid}"
    fingerprint = "${var.fingerprint}"
    private_key_path = "${var.private_key_path}"
    region = "${var.region}"
}

data "oci_identity_availability_domains" "ADs" {
    compartment_id = "${var.tenancy_ocid}"
}

# First VCN
resource "oci_core_virtual_network" "CoreVCN" {
    cidr_block = "${var.vcn_cidr}"
    compartment_id = "${var.compartment_ocid}"
    display_name = "VCN-1"
}

# Second VCN
resource "oci_core_virtual_network" "CoreVCN2" {
    cidr_block = "${var.vcn_cidr2}"
    compartment_id = "${var.compartment_ocid}"
    display_name = "VCN-2"
}

# First VCN configuration
resource "oci_core_internet_gateway" "MgmtIG" {
    compartment_id = "${var.compartment_ocid}"
    display_name = "MgmtIG"
    vcn_id = "${oci_core_virtual_network.CoreVCN.id}"
}

resource "oci_core_route_table" "MgmtRouteTable" {
    compartment_id = "${var.compartment_ocid}"
    vcn_id = "${oci_core_virtual_network.CoreVCN.id}"
    display_name = "MgmtRouteTable"
    route_rules {
        cidr_block = "0.0.0.0/0"
        network_entity_id = "${oci_core_internet_gateway.MgmtIG.id}"
    }
}

resource "oci_core_security_list" "MgmtSecurityList" {
    compartment_id = "${var.compartment_ocid}"
    display_name = "MgmtSecurityList"
    vcn_id = "${oci_core_virtual_network.CoreVCN.id}"

    egress_security_rules = [{
        protocol = "all"
        destination = "0.0.0.0/0"
    }]

    ingress_security_rules = [{
        tcp_options {
            "max" = 80
            "min" = 80
        }
        protocol = "6"
        source = "0.0.0.0/0"
    },
    {
        tcp_options {
            "max" = 443
            "min" = 443
        }
        protocol = "6"
        source = "0.0.0.0/0"
    },
	{
        protocol = "all"
        source = "${var.vcn_cidr}"
    },
        {
        protocol = "all"
        source = "${var.vcn_cidr2}"
    },
    {
        protocol = "6"
        source = "0.0.0.0/0"
        tcp_options {
            "min" = 22
            "max" = 22
        }
    },
    {
        protocol = "1"
        source = "0.0.0.0/0"
        icmp_options {
            "type" = 3
            "code" = 4
        }
    }]
}

resource "oci_core_subnet" "MgmtSubnet" {
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
    cidr_block = "${var.mgmt_subnet_cidr}"
    display_name = "MgmtSubnet"
    compartment_id = "${var.compartment_ocid}"
    vcn_id = "${oci_core_virtual_network.CoreVCN.id}"
    route_table_id = "${oci_core_route_table.MgmtRouteTable.id}"
    security_list_ids = ["${oci_core_security_list.MgmtSecurityList.id}"]
    dhcp_options_id = "${oci_core_virtual_network.CoreVCN.default_dhcp_options_id}"
}

# Creating the Bridge Instance
resource "oci_core_instance" "BridgeInstance" {
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
    compartment_id = "${var.compartment_ocid}"
    display_name = "BridgeInstance"
    image = "${var.InstanceImageOCID[var.region]}"
    shape = "${var.InstanceShape}"
    create_vnic_details {
        subnet_id = "${oci_core_subnet.MgmtSubnet.id}"
        skip_source_dest_check = true
    }
    metadata {
        ssh_authorized_keys = "${var.ssh_public_key}"
        user_data = "${base64encode(file("user_data.tpl"))}"
    }
    timeouts {
        create = "10m"
    }
}

resource "oci_core_security_list" "MgmtSecurityList2" {
    compartment_id = "${var.compartment_ocid}"
    display_name = "MgmtSecurityList2"
    vcn_id = "${oci_core_virtual_network.CoreVCN2.id}"

    egress_security_rules = [{
        protocol = "all"
        destination = "0.0.0.0/0"
    }]

    ingress_security_rules = [{
        tcp_options {
            "max" = 80
            "min" = 80
        }
        protocol = "6"
        source = "0.0.0.0/0"
    },
    {
        tcp_options {
            "max" = 443
            "min" = 443
        }
        protocol = "6"
        source = "0.0.0.0/0"
    },
        {
        protocol = "all"
        source = "${var.vcn_cidr2}"
    },
    {
        protocol = "6"
        source = "0.0.0.0/0"
        tcp_options {
            "min" = 22
            "max" = 22
        }
    },
    {
        protocol = "1"
        source = "0.0.0.0/0"
        icmp_options {
            "type" = 3
            "code" = 4
        }
    }]
}

resource "oci_core_subnet" "MgmtSubnet2" {
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
    cidr_block = "${var.mgmt_subnet_cidr2}"
    display_name = "MgmtSubnet2"
    compartment_id = "${var.compartment_ocid}"
    vcn_id = "${oci_core_virtual_network.CoreVCN2.id}"
    security_list_ids = ["${oci_core_security_list.MgmtSecurityList2.id}"]
    dhcp_options_id = "${oci_core_virtual_network.CoreVCN2.default_dhcp_options_id}"
}

# Gets a list of VNIC attachments on the instance
data "oci_core_vnic_attachments" "BridgeInstanceVnics" {
    compartment_id = "${var.compartment_ocid}"
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
    instance_id = "${oci_core_instance.BridgeInstance.id}"
}

# Create PrivateIP
resource "oci_core_private_ip" "BridgeInstancePrivateIP" {
    vnic_id = "${lookup(data.oci_core_vnic_attachments.BridgeInstanceVnics.vnic_attachments[0],"vnic_id")}"
    display_name = "BridgeInstancePrivateIP"
}

# Get the OCID of the first (default) VNIC
data "oci_core_vnic" "BridgeInstanceVnic1" {
    vnic_id = "${lookup(data.oci_core_vnic_attachments.BridgeInstanceVnics.vnic_attachments[0],"vnic_id")}"
}

# Frist VCN Private instance details
resource "oci_core_security_list" "PrivateSecurityList" {
    compartment_id = "${var.compartment_ocid}"
    display_name = "PrivateSecurityList"
    vcn_id = "${oci_core_virtual_network.CoreVCN.id}"
    egress_security_rules = [{
        protocol = "all"
        destination = "0.0.0.0/0"
    }]
    ingress_security_rules = [{
        protocol = "6"
        tcp_options {
            "max" = 22
            "min" = 22
        }
        source = "${var.vcn_cidr}"
    },
        {
        protocol = "all"
        source = "${var.vcn_cidr}"
    },
        {
        protocol = "all"
        source = "${var.vcn_cidr2}"
    }]
}

resource "oci_core_route_table" "PrivateRouteTable" {
    compartment_id = "${var.compartment_ocid}"
    vcn_id = "${oci_core_virtual_network.CoreVCN.id}"
    display_name = "PrivateRouteTable"
    route_rules {
        cidr_block = "0.0.0.0/0"
        network_entity_id = "${oci_core_private_ip.BridgeInstancePrivateIP.id}"
    }
}

resource "oci_core_subnet" "PrivateSubnet" {
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
    cidr_block = "${var.private_subnet_cidr}"
    display_name = "PrivateSubnet"
    compartment_id = "${var.compartment_ocid}"
    vcn_id = "${oci_core_virtual_network.CoreVCN.id}"
    route_table_id = "${oci_core_route_table.PrivateRouteTable.id}"
    security_list_ids = ["${oci_core_security_list.PrivateSecurityList.id}"]
    dhcp_options_id = "${oci_core_virtual_network.CoreVCN.default_dhcp_options_id}"
    prohibit_public_ip_on_vnic = "true"
}

resource "oci_core_instance" "PrivateInstance" {
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
    compartment_id = "${var.compartment_ocid}"
    display_name = "PrivateInstance"
    image = "${var.InstanceImageOCID[var.region]}"
    shape = "${var.InstanceShape}"
    create_vnic_details {
      subnet_id = "${oci_core_subnet.PrivateSubnet.id}"
      assign_public_ip = false
    }
    metadata {
      ssh_authorized_keys = "${var.ssh_public_key}"
    }
    timeouts {
      create = "10m"
    }
}

# Second VCN private instance details
resource "oci_core_security_list" "PrivateSecurityList2" {
    compartment_id = "${var.compartment_ocid}"
    display_name = "PrivateSecurityList2"
    vcn_id = "${oci_core_virtual_network.CoreVCN2.id}"
    egress_security_rules = [{
        protocol = "all"
        destination = "0.0.0.0/0"
    }]
    ingress_security_rules = [{
        protocol = "6"
        tcp_options {
            "max" = 22
            "min" = 22
        }
        source = "${var.vcn_cidr2}"
    },
        {
        protocol = "all"
        source = "${var.vcn_cidr}"
    },
        {
        protocol = "all"
        source = "${var.vcn_cidr2}"
    }]
}

resource "oci_core_subnet" "PrivateSubnet2" {
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
    cidr_block = "${var.private_subnet_cidr2}"
    display_name = "PrivateSubnet2"
    compartment_id = "${var.compartment_ocid}"
    vcn_id = "${oci_core_virtual_network.CoreVCN2.id}"
    route_table_id = "${oci_core_route_table.PrivateRouteTable2.id}"
    security_list_ids = ["${oci_core_security_list.PrivateSecurityList2.id}"]
    dhcp_options_id = "${oci_core_virtual_network.CoreVCN2.default_dhcp_options_id}"
    prohibit_public_ip_on_vnic = "true"
}

resource "oci_core_instance" "PrivateInstance2" {
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
    compartment_id = "${var.compartment_ocid}"
    display_name = "PrivateInstance2"
    image = "${var.InstanceImageOCID[var.region]}"
    shape = "${var.InstanceShape2}"
    create_vnic_details {
      subnet_id = "${oci_core_subnet.PrivateSubnet2.id}"
      assign_public_ip = false
    }
    metadata {
      ssh_authorized_keys = "${var.ssh_public_key}"
    }
    timeouts {
      create = "10m"
    }
}

# Creating secondary VNIC on BridgeInstance and attaching it to Second VCN Mgmt subnet
resource "oci_core_vnic_attachment" "SecondaryVnicAttachment" {
    create_vnic_details {
        subnet_id = "${oci_core_subnet.MgmtSubnet2.id}"
        display_name = "SecondaryVNIC"
        skip_source_dest_check = true
    }
    instance_id = "${oci_core_instance.BridgeInstance.id}"
}

# Gets a list of VNIC attachments on the instance
data "oci_core_vnic_attachments" "BridgeInstanceVnics2" {
    compartment_id = "${var.compartment_ocid}"
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
    instance_id = "${oci_core_instance.BridgeInstance.id}"
}

# Gets the OCID of the second VNIC
data "oci_core_vnic" "BridgeInstanceVnic2" {
    vnic_id = "${oci_core_vnic_attachment.SecondaryVnicAttachment.vnic_id}"
}

# Gets a list of private IPs on the second VNIC
data "oci_core_private_ips" "BridgeInstancePrivateIP2" {
    vnic_id = "${data.oci_core_vnic.BridgeInstanceVnic2.id}"
}

resource "oci_core_route_table" "PrivateRouteTable2" {
    compartment_id = "${var.compartment_ocid}"
    vcn_id = "${oci_core_virtual_network.CoreVCN2.id}"
    display_name = "PrivateRouteTable2"
    route_rules {
        cidr_block = "0.0.0.0/0"
        network_entity_id = "${lookup(data.oci_core_private_ips.BridgeInstancePrivateIP2.private_ips[0],"id")}"
    }
}

# Configurations for setting up the secondary VNIC
resource "null_resource" "configure-secondary-vnic" {
  connection {
    type        = "ssh"
    user        = "opc"
    private_key = "${var.ssh_private_key}"
    host        = "${data.oci_core_vnic.BridgeInstanceVnic1.public_ip_address}"
    timeout     = "30m"
}
  provisioner "remote-exec" {
    inline = [
      "sudo wget https://docs.cloud.oracle.com/iaas/Content/Resources/Assets/secondary_vnic_all_configure.sh",
      "sudo chmod 777 secondary_vnic_all_configure.sh",

      "sudo ./secondary_vnic_all_configure.sh -c ${lookup(data.oci_core_private_ips.BridgeInstancePrivateIP2.private_ips[0],"id")}",
      "sudo ip route add ${var.vcn_cidr2} dev ens4 via ${oci_core_subnet.MgmtSubnet2.virtual_router_ip}"
    ]
  }
}

# Outputing required info for users
output "Bridge Instance Public IP" {
    value = "${data.oci_core_vnic.BridgeInstanceVnic1.public_ip_address}"
}

output "PrivateInstance1 Private IP" {
    value = "${oci_core_instance.PrivateInstance.private_ip}"
}

output "PrivateInstance2 Private IP" {
    value = "${oci_core_instance.PrivateInstance2.private_ip}"
}

output "SSH login to the Bridge Instance" {
    value = "ssh -A opc@${data.oci_core_vnic.BridgeInstanceVnic1.public_ip_address}"
}
