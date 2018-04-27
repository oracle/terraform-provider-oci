variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "private_key_password" {}
variable "compartment_ocid" {}
variable "region" {}
variable "ssh_public_key" {}
variable "ssh_private_key" {}

# Choose an Availability Domain
variable "AD1" {
    default = "1"
}

variable "AD2" {
    default = "2"
}

variable "InstanceShape" {
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

variable "vcn_cidr" {
    default = "10.0.0.0/16"
}

variable "mgmt_subnet_cidr1" {
    default = "10.0.0.0/24"
}

variable "mgmt_subnet_cidr2" {
    default = "10.0.1.0/24"
}

variable "onprem_cidr" {
    default = "172.16.0.0/16"
}

variable "onprem_dns_zone" {
    default = "customer.net"
}

variable "onprem_dns_server1" {
    default = "172.16.0.5"
}

variable "onprem_dns_server2" {
    default = "172.16.31.5"
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
resource "oci_core_virtual_network" "CoreVCN" {
    cidr_block = "${var.vcn_cidr}"
    compartment_id = "${var.compartment_ocid}"
    display_name = "mgmt-vcn"
    dns_label = "mgmtvcn"
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
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD1 - 1],"name")}"
    cidr_block = "${var.mgmt_subnet_cidr1}"
    display_name = "MgmtSubnet"
    dns_label = "mgmtsubnet"
    compartment_id = "${var.compartment_ocid}"
    vcn_id = "${oci_core_virtual_network.CoreVCN.id}"
    route_table_id = "${oci_core_route_table.MgmtRouteTable.id}"
    security_list_ids = ["${oci_core_security_list.MgmtSecurityList.id}"]
    dhcp_options_id = "${oci_core_dhcp_options.MgmtDhcpOptions.id}"
}

resource "oci_core_subnet" "MgmtSubnet2" {
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD2 - 1],"name")}"
    cidr_block = "${var.mgmt_subnet_cidr2}"
    display_name = "MgmtSubnet2"
    dns_label = "mgmtsubnet2"
    compartment_id = "${var.compartment_ocid}"
    vcn_id = "${oci_core_virtual_network.CoreVCN.id}"
    route_table_id = "${oci_core_route_table.MgmtRouteTable.id}"
    security_list_ids = ["${oci_core_security_list.MgmtSecurityList.id}"]
    dhcp_options_id = "${oci_core_dhcp_options.MgmtDhcpOptions.id}"
}

resource "oci_core_instance" "DnsVM" {
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD1 - 1],"name")}"
    compartment_id = "${var.compartment_ocid}"
    display_name = "DnsVM"
    image = "${var.InstanceImageOCID[var.region]}"
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

resource "oci_core_instance" "DnsVM2" {
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD2 - 1],"name")}"
    compartment_id = "${var.compartment_ocid}"
    display_name = "DnsVM2"
    image = "${var.InstanceImageOCID[var.region]}"
    shape = "${var.InstanceShape}"
    create_vnic_details {
        subnet_id = "${oci_core_subnet.MgmtSubnet2.id}"
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
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD1 - 1],"name")}"
    instance_id = "${oci_core_instance.DnsVM.id}"
}

data "oci_core_vnic_attachments" "DnsVMVnics2" {
    compartment_id = "${var.compartment_ocid}"
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD2 - 1],"name")}"
    instance_id = "${oci_core_instance.DnsVM2.id}"
}

# Gets the OCID of the first (default) vNIC
data "oci_core_vnic" "DnsVMVnic" {
    vnic_id = "${lookup(data.oci_core_vnic_attachments.DnsVMVnics.vnic_attachments[0],"vnic_id")}"
}

data "oci_core_vnic" "DnsVMVnic2" {
    vnic_id = "${lookup(data.oci_core_vnic_attachments.DnsVMVnics2.vnic_attachments[0],"vnic_id")}"
}

# Update the default DHCP options to use custom DNS servers
resource "oci_core_default_dhcp_options" "default-dhcp-options" {
    manage_default_resource_id = "${oci_core_virtual_network.CoreVCN.default_dhcp_options_id}"

    // required
    options {
        type = "DomainNameServer"
        server_type = "CustomDnsServer"
        custom_dns_servers = [  "${data.oci_core_vnic.DnsVMVnic.private_ip_address}",
                                "${data.oci_core_vnic.DnsVMVnic2.private_ip_address}" ]
    }

  // optional
  options {
    type = "SearchDomain"
    search_domain_names = [ "${oci_core_virtual_network.CoreVCN.dns_label}.oraclevcn.com" ]
  }
}

output "DnsServer1" {
    value = ["${data.oci_core_vnic.DnsVMVnic.private_ip_address}"]
}

output "DnsServer2" {
    value = ["${data.oci_core_vnic.DnsVMVnic2.private_ip_address}"]
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

resource "null_resource" "configure-bind-vm1" {
  connection {
    type        = "ssh"
    user        = "opc"
    private_key = "${var.ssh_private_key}"
    host        = "${data.oci_core_vnic.DnsVMVnic.public_ip_address}"
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

resource "null_resource" "configure-bind-vm2" {
  connection {
    type        = "ssh"
    user        = "opc"
    private_key = "${var.ssh_private_key}"
    host        = "${data.oci_core_vnic.DnsVMVnic2.public_ip_address}"
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

