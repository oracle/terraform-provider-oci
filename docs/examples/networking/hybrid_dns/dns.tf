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
    default = "VM.Standard1.1"
}

variable "InstanceOS" {
    default = "Oracle Linux"
}

variable "InstanceOSVersion" {
    default = "7.4"
}

variable "vcn_cidr" {
    default = "10.0.0.0/16"
}

variable "mgmt_subnet_cidr" {
    default = "10.0.10.0/24"
}

variable "onprem_cidr" {
    default = "172.16.0.0/16"
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
resource "oci_core_virtual_network" "CoreVCN" {
    cidr_block = "${var.vcn_cidr}"
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
            "max" = 53
            "min" = 53
        }
        protocol = "6"
        source = "${var.vcn_cidr}"
    },
    {
        udp_options {
            "max" = 53
            "min" = 53
        }
        protocol = "17"
        source = "${var.vcn_cidr}"
    },
    {
        tcp_options {
            "max" = 53
            "min" = 53
        }
        protocol = "6"
        source = "${var.onprem_cidr}"
    },
    {
        udp_options {
            "max" = 53
            "min" = 53
        }
        protocol = "17"
        source = "${var.onprem_cidr}"
    },
	{
        protocol = "all"
        source = "${var.vcn_cidr}"
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

resource "oci_core_dhcp_options" "MgmtDhcpOptions" {
    compartment_id = "${var.compartment_ocid}"
    vcn_id = "${oci_core_virtual_network.CoreVCN.id}"
    display_name = "MgmtDhcpOptions"
 
    options {
      type = "DomainNameServer"
      server_type = "VcnLocalPlusInternet"
    }
}

resource "oci_core_subnet" "MgmtSubnet" {
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
    cidr_block = "${var.mgmt_subnet_cidr}"
    display_name = "MgmtSubnet"
    compartment_id = "${var.compartment_ocid}"
    vcn_id = "${oci_core_virtual_network.CoreVCN.id}"
    route_table_id = "${oci_core_route_table.MgmtRouteTable.id}"
    security_list_ids = ["${oci_core_security_list.MgmtSecurityList.id}"]
    dhcp_options_id = "${oci_core_dhcp_options.MgmtDhcpOptions.id}"
}

# Gets the OCID of the OS image to use
data "oci_core_images" "OLImageOCID" {
    compartment_id = "${var.compartment_ocid}"
    operating_system = "${var.InstanceOS}"
    operating_system_version = "${var.InstanceOSVersion}"
}

resource "oci_core_instance" "DnsVM" {
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
    compartment_id = "${var.compartment_ocid}"
    display_name = "DnsVM"
    image = "${lookup(data.oci_core_images.OLImageOCID.images[0], "id")}"
    shape = "${var.InstanceShape}"
    create_vnic_details {
        subnet_id = "${oci_core_subnet.MgmtSubnet.id}"
    }
    metadata {
        ssh_authorized_keys = "${var.ssh_public_key}"
    }
    timeouts {
        create = "10m"
    }
}

# Gets a list of VNIC attachments on the DNS instance
data "oci_core_vnic_attachments" "DnsVMVnics" {
    compartment_id = "${var.compartment_ocid}"
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
    instance_id = "${oci_core_instance.DnsVM.id}"
}

# Gets the OCID of the first (default) vNIC
data "oci_core_vnic" "InstanceVnic" {
    vnic_id = "${lookup(data.oci_core_vnic_attachments.DnsVMVnics.vnic_attachments[0],"vnic_id")}"
}

output "DnsVMPrivateIP" {
    value = ["${data.oci_core_vnic.InstanceVnic.private_ip_address}"]
}

output "DnsVMPublicIP" {
    value = ["${data.oci_core_vnic.InstanceVnic.public_ip_address}"]
}

data "template_file" "generate_named_conf" {
    template = "${file("named.conf.tpl")}"
 
    vars {
      vcn_cidr           = "${var.vcn_cidr}"
      onprem_cidr        = "${var.onprem_cidr}"
      onprem_dns_zone    = "${var.onprem_dns_zone}"
      onprem_dns_server1 = "${var.onprem_dns_server1}"
      onprem_dns_server2 = "${var.onprem_dns_server2}"
    }
}

resource "null_resource" "configure-bind" {
  connection {
    type        = "ssh"
    user        = "opc"
    private_key = "${var.ssh_private_key}"
    host        = "${data.oci_core_vnic.InstanceVnic.public_ip_address}"
    timeout     = "30m"
  }

  provisioner "file" {
    content     = "${data.template_file.generate_named_conf.rendered}"
    destination = "~/named.conf"
  }

  provisioner "remote-exec" {
    inline = [
      "sudo yum update -y",
      "sudo yum install bind -y",
      "sudo firewall-offline-cmd --add-port=53/udp",
      "sudo firewall-offline-cmd --add-port=53/tcp",
      "sudo /bin/systemctl restart firewalld",
      "sudo cp ~/named.conf /etc/named.conf",
      "sudo service named restart"
    ]
  }
}

