// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}
variable "region" {}
variable "ssh_public_key" {}
variable "ssh_private_key" {}

variable "instance_shape" {
  default = "VM.Standard2.1"
}

variable "instance_image_ocid" {
  type = "map"

  default = {
    // See https://docs.us-phoenix-1.oraclecloud.com/images/
    // Oracle-provided image "Oracle-Linux-7.5-2018.10.16-0"
    us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaaoqj42sokaoh42l76wsyhn3k2beuntrh5maj3gmgmzeyr55zzrwwa"

    us-ashburn-1   = "ocid1.image.oc1.iad.aaaaaaaageeenzyuxgia726xur4ztaoxbxyjlxogdhreu3ngfj2gji3bayda"
    eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaaitzn6tdyjer7jl34h2ujz74jwy5nkbukbh55ekp6oyzwrtfa4zma"
    uk-london-1    = "ocid1.image.oc1.uk-london-1.aaaaaaaa32voyikkkzfxyo4xbdmadc2dmvorfxxgdhpnk6dw64fa3l4jh7wa"
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
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region           = "${var.region}"
}

data "oci_identity_availability_domain" "ad1" {
  compartment_id = "${var.tenancy_ocid}"
  ad_number      = 1
}

data "oci_identity_availability_domain" "ad2" {
  compartment_id = "${var.tenancy_ocid}"
  ad_number      = 2
}

resource "oci_core_virtual_network" "CoreVCN" {
  cidr_block     = "${var.vcn_cidr}"
  compartment_id = "${var.compartment_ocid}"
  display_name   = "mgmt-vcn"
  dns_label      = "mgmtvcn"
}

resource "oci_core_internet_gateway" "MgmtIG" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "MgmtIG"
  vcn_id         = "${oci_core_virtual_network.CoreVCN.id}"
}

resource "oci_core_route_table" "MgmtRouteTable" {
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_virtual_network.CoreVCN.id}"
  display_name   = "MgmtRouteTable"

  route_rules {
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = "${oci_core_internet_gateway.MgmtIG.id}"
  }
}

resource "oci_core_security_list" "MgmtSecurityList" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "MgmtSecurityList"
  vcn_id         = "${oci_core_virtual_network.CoreVCN.id}"

  egress_security_rules = [{
    protocol    = "all"
    destination = "0.0.0.0/0"
  }]

  ingress_security_rules = [{
    tcp_options {
      "max" = 53
      "min" = 53
    }

    protocol = "6"
    source   = "${var.vcn_cidr}"
  },
    {
      udp_options {
        "max" = 53
        "min" = 53
      }

      protocol = "17"
      source   = "${var.vcn_cidr}"
    },
    {
      tcp_options {
        "max" = 53
        "min" = 53
      }

      protocol = "6"
      source   = "${var.onprem_cidr}"
    },
    {
      udp_options {
        "max" = 53
        "min" = 53
      }

      protocol = "17"
      source   = "${var.onprem_cidr}"
    },
    {
      protocol = "all"
      source   = "${var.vcn_cidr}"
    },
    {
      protocol = "6"
      source   = "0.0.0.0/0"

      tcp_options {
        "min" = 22
        "max" = 22
      }
    },
    {
      protocol = "1"
      source   = "0.0.0.0/0"

      icmp_options {
        "type" = 3
        "code" = 4
      }
    },
  ]
}

resource "oci_core_dhcp_options" "MgmtDhcpOptions" {
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_virtual_network.CoreVCN.id}"
  display_name   = "MgmtDhcpOptions"

  options {
    type        = "DomainNameServer"
    server_type = "VcnLocalPlusInternet"
  }
}

resource "oci_core_subnet" "MgmtSubnet" {
  availability_domain = "${data.oci_identity_availability_domain.ad1.name}"
  cidr_block          = "${var.mgmt_subnet_cidr1}"
  display_name        = "MgmtSubnet"
  dns_label           = "mgmtsubnet"
  compartment_id      = "${var.compartment_ocid}"
  vcn_id              = "${oci_core_virtual_network.CoreVCN.id}"
  route_table_id      = "${oci_core_route_table.MgmtRouteTable.id}"
  security_list_ids   = ["${oci_core_security_list.MgmtSecurityList.id}"]
  dhcp_options_id     = "${oci_core_dhcp_options.MgmtDhcpOptions.id}"
}

resource "oci_core_subnet" "MgmtSubnet2" {
  availability_domain = "${data.oci_identity_availability_domain.ad2.name}"
  cidr_block          = "${var.mgmt_subnet_cidr2}"
  display_name        = "MgmtSubnet2"
  dns_label           = "mgmtsubnet2"
  compartment_id      = "${var.compartment_ocid}"
  vcn_id              = "${oci_core_virtual_network.CoreVCN.id}"
  route_table_id      = "${oci_core_route_table.MgmtRouteTable.id}"
  security_list_ids   = ["${oci_core_security_list.MgmtSecurityList.id}"]
  dhcp_options_id     = "${oci_core_dhcp_options.MgmtDhcpOptions.id}"
}

resource "oci_core_instance" "DnsVM" {
  availability_domain = "${data.oci_identity_availability_domain.ad1.name}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "DnsVM"
  shape               = "${var.instance_shape}"

  create_vnic_details {
    subnet_id = "${oci_core_subnet.MgmtSubnet.id}"
  }

  metadata {
    ssh_authorized_keys = "${var.ssh_public_key}"
  }

  source_details {
    source_type = "image"
    source_id   = "${var.instance_image_ocid[var.region]}"
  }

  timeouts {
    create = "10m"
  }
}

resource "oci_core_instance" "DnsVM2" {
  availability_domain = "${data.oci_identity_availability_domain.ad2.name}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "DnsVM2"
  shape               = "${var.instance_shape}"

  create_vnic_details {
    subnet_id = "${oci_core_subnet.MgmtSubnet2.id}"
  }

  metadata {
    ssh_authorized_keys = "${var.ssh_public_key}"
  }

  source_details {
    source_type = "image"
    source_id   = "${var.instance_image_ocid[var.region]}"
  }

  timeouts {
    create = "10m"
  }
}

# Gets a list of VNIC attachments on the DNS instance
data "oci_core_vnic_attachments" "DnsVMVnics" {
  compartment_id      = "${var.compartment_ocid}"
  availability_domain = "${data.oci_identity_availability_domain.ad1.name}"
  instance_id         = "${oci_core_instance.DnsVM.id}"
}

data "oci_core_vnic_attachments" "DnsVMVnics2" {
  compartment_id      = "${var.compartment_ocid}"
  availability_domain = "${data.oci_identity_availability_domain.ad2.name}"
  instance_id         = "${oci_core_instance.DnsVM2.id}"
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
    type        = "DomainNameServer"
    server_type = "CustomDnsServer"

    custom_dns_servers = ["${data.oci_core_vnic.DnsVMVnic.private_ip_address}",
      "${data.oci_core_vnic.DnsVMVnic2.private_ip_address}",
    ]
  }

  // optional
  options {
    type                = "SearchDomain"
    search_domain_names = ["${oci_core_virtual_network.CoreVCN.dns_label}.oraclevcn.com"]
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
      "sudo service named restart",
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
      "sudo service named restart",
    ]
  }
}
