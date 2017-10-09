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

variable "InstanceOS" {
    default = "Oracle Linux"
}
variable "InstanceOSVersion" {
    default = "7.4"
}
provider "oci" {
    tenancy_ocid = "${var.tenancy_ocid}"
    user_ocid = "${var.user_ocid}"
    fingerprint = "${var.fingerprint}"
    private_key_path = "${var.private_key_path}"
    private_key_password = "${var.private_key_password}"
    region = "${var.region}"
}
data "oci_identity_availability_domains" "ADs" {
    compartment_id = "${var.tenancy_ocid}"
}
variable "VCN-CIDR" {
    default = "10.0.0.0/16"
}

resource "oci_core_virtual_network" "CoreVCN" {
    cidr_block = "${var.VCN-CIDR}"
    compartment_id = "${var.compartment_ocid}"
    display_name = "mgmt-vcn"
}

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
        source = "${var.VCN-CIDR}"
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
    cidr_block = "10.0.0.0/24"
    display_name = "MgmtSubnet"
    compartment_id = "${var.compartment_ocid}"
    vcn_id = "${oci_core_virtual_network.CoreVCN.id}"
    route_table_id = "${oci_core_route_table.MgmtRouteTable.id}"
    security_list_ids = ["${oci_core_security_list.MgmtSecurityList.id}"]
    dhcp_options_id = "${oci_core_virtual_network.CoreVCN.default_dhcp_options_id}"
}

# Gets the OCID of the OS image to use
data "oci_core_images" "OLImageOCID" {
    compartment_id = "${var.compartment_ocid}"
    operating_system = "${var.InstanceOS}"
    operating_system_version = "${var.InstanceOSVersion}"
}

resource "oci_core_instance" "NatInstance" {
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
    compartment_id = "${var.compartment_ocid}"
    display_name = "NatInstance"
    image = "${lookup(data.oci_core_images.OLImageOCID.images[0], "id")}"
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

# Gets a list of VNIC attachments on the instance
data "oci_core_vnic_attachments" "NatInstanceVnics" {
    compartment_id = "${var.compartment_ocid}"
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
    instance_id = "${oci_core_instance.NatInstance.id}"
}

# Create PrivateIP
resource "oci_core_private_ip" "NatInstancePrivateIP" {
    vnic_id = "${lookup(data.oci_core_vnic_attachments.NatInstanceVnics.vnic_attachments[0],"vnic_id")}"
    display_name = "NatInstancePrivateIP"
}

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
        source = "${var.VCN-CIDR}"
    }]
}

resource "oci_core_route_table" "PrivateRouteTable" {
    compartment_id = "${var.compartment_ocid}"
    vcn_id = "${oci_core_virtual_network.CoreVCN.id}"
    display_name = "PrivateRouteTable"
    route_rules {
        cidr_block = "0.0.0.0/0"
        network_entity_id = "${oci_core_private_ip.NatInstancePrivateIP.id}"
    }
}

resource "oci_core_subnet" "PrivateSubnet" {
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
    cidr_block = "10.0.1.0/24"
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
    image = "${lookup(data.oci_core_images.OLImageOCID.images[0], "id")}"
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
