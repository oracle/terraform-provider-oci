/*
 * Sets up a VCN with:
 * - NAT gateway
 * - bastion subnet and bastion instance
 * - private subnet that routes all traffic to the NAT
 * - a test instance in the private subnet
 *
 * After applying, you should be to to ssh into the private instance
 * via the bastion and verify internet access via the NAT.
 */

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}

variable "region" {}

variable "compartment_id" {}

variable "ssh_public_key_path" {}

// zero-based ad number
variable "ad_number" {
  default = 1
}

variable "vcn_cidr" {
  default = "10.0.0.0/16"
}

variable subnet_cidr_offset {
  default = 5
}

variable "instance_image_id" {
  type = "map"

  default = {
    // Oracle-provided image "Oracle-Linux-7.5-2018.10.16-0"
    // See https://docs.us-phoenix-1.oraclecloud.com/images/
    us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaaoqj42sokaoh42l76wsyhn3k2beuntrh5maj3gmgmzeyr55zzrwwa"

    us-ashburn-1   = "ocid1.image.oc1.iad.aaaaaaaageeenzyuxgia726xur4ztaoxbxyjlxogdhreu3ngfj2gji3bayda"
    eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaaitzn6tdyjer7jl34h2ujz74jwy5nkbukbh55ekp6oyzwrtfa4zma"
    uk-london-1    = "ocid1.image.oc1.uk-london-1.aaaaaaaa32voyikkkzfxyo4xbdmadc2dmvorfxxgdhpnk6dw64fa3l4jh7wa"
  }
}

variable "instance_shape" {
  default = "VM.Standard2.1"
}

locals {
  bastion_subnet_prefix = "${cidrsubnet(var.vcn_cidr, var.subnet_cidr_offset, 0)}"
  private_subnet_prefix = "${cidrsubnet(var.vcn_cidr, var.subnet_cidr_offset, 1)}"

  ad = "${lookup(data.oci_identity_availability_domains.ads.availability_domains[var.ad_number],"name")}"

  tcp_protocol  = "6"
  all_protocols = "all"
  anywhere      = "0.0.0.0/0"
}

provider "oci" {
  region           = "${var.region}"
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
}

data "oci_identity_availability_domains" "ads" {
  compartment_id = "${var.tenancy_ocid}"
}

resource "oci_core_virtual_network" "this" {
  cidr_block     = "${var.vcn_cidr}"
  dns_label      = "pp"
  compartment_id = "${var.compartment_id}"
  display_name   = "proxy_prototype"
}

resource "oci_core_nat_gateway" "nat_gateway" {
  compartment_id = "${var.compartment_id}"
  vcn_id         = "${oci_core_virtual_network.this.id}"
  display_name   = "nat_gateway"
}

resource "oci_core_internet_gateway" "ig" {
  compartment_id = "${var.compartment_id}"
  display_name   = "proxy_prototype"
  vcn_id         = "${oci_core_virtual_network.this.id}"
}

resource "oci_core_subnet" "bastion" {
  availability_domain = "${local.ad}"
  cidr_block          = "${local.bastion_subnet_prefix}"
  display_name        = "bastion"
  compartment_id      = "${var.compartment_id}"
  vcn_id              = "${oci_core_virtual_network.this.id}"
  route_table_id      = "${oci_core_route_table.bastion.id}"

  security_list_ids = [
    "${oci_core_security_list.bastion.id}",
  ]

  dns_label                  = "bastion"
  prohibit_public_ip_on_vnic = false
}

resource "oci_core_route_table" "bastion" {
  compartment_id = "${var.compartment_id}"
  vcn_id         = "${oci_core_virtual_network.this.id}"
  display_name   = "bastion"

  route_rules {
    destination       = "${local.anywhere}"
    network_entity_id = "${oci_core_internet_gateway.ig.id}"
  }
}

resource "oci_core_security_list" "bastion" {
  compartment_id = "${var.compartment_id}"
  display_name   = "bastion"
  vcn_id         = "${oci_core_virtual_network.this.id}"

  ingress_security_rules {
    source   = "${local.anywhere}"
    protocol = "${local.tcp_protocol}"

    tcp_options {
      "min" = 22
      "max" = 22
    }
  }

  egress_security_rules {
    destination = "${var.vcn_cidr}"
    protocol    = "${local.tcp_protocol}"

    tcp_options {
      "min" = 22
      "max" = 22
    }
  }
}

resource "oci_core_instance" "bastion" {
  availability_domain = "${local.ad}"
  compartment_id      = "${var.compartment_id}"
  display_name        = "bastion"
  shape               = "${var.instance_shape}"

  source_details {
    source_id   = "${var.instance_image_id[var.region]}"
    source_type = "image"
  }

  create_vnic_details {
    subnet_id = "${oci_core_subnet.bastion.id}"
  }

  metadata {
    ssh_authorized_keys = "${file(var.ssh_public_key_path)}"
  }

  timeouts {
    create = "10m"
  }
}

output "bastion_public_ip" {
  value = "${oci_core_instance.bastion.public_ip}"
}

resource "oci_core_subnet" "private" {
  availability_domain = "${local.ad}"
  cidr_block          = "${local.private_subnet_prefix}"
  display_name        = "private"
  compartment_id      = "${var.compartment_id}"
  vcn_id              = "${oci_core_virtual_network.this.id}"
  route_table_id      = "${oci_core_route_table.private.id}"

  security_list_ids = [
    "${oci_core_security_list.private.id}",
  ]

  dns_label                  = "private"
  prohibit_public_ip_on_vnic = true
}

resource "oci_core_route_table" "private" {
  compartment_id = "${var.compartment_id}"
  vcn_id         = "${oci_core_virtual_network.this.id}"
  display_name   = "private"

  route_rules = [
    {
      destination       = "${local.anywhere}"
      destination_type  = "CIDR_BLOCK"
      network_entity_id = "${oci_core_nat_gateway.nat_gateway.id}"
    },
  ]
}

resource "oci_core_security_list" "private" {
  compartment_id = "${var.compartment_id}"
  display_name   = "private"
  vcn_id         = "${oci_core_virtual_network.this.id}"

  ingress_security_rules {
    source   = "${local.bastion_subnet_prefix}"
    protocol = "${local.tcp_protocol}"

    tcp_options {
      "min" = 22
      "max" = 22
    }
  }

  egress_security_rules {
    destination = "${local.anywhere}"
    protocol    = "${local.all_protocols}"
  }
}

resource "oci_core_instance" "private" {
  availability_domain = "${local.ad}"
  compartment_id      = "${var.compartment_id}"
  display_name        = "private_test_instance"
  shape               = "${var.instance_shape}"

  source_details {
    source_id   = "${var.instance_image_id[var.region]}"
    source_type = "image"
  }

  create_vnic_details {
    subnet_id        = "${oci_core_subnet.private.id}"
    assign_public_ip = false
  }

  metadata {
    ssh_authorized_keys = "${file(var.ssh_public_key_path)}"
  }

  timeouts {
    create = "10m"
  }
}

output "private_instance_ip" {
  value = "${oci_core_instance.private.private_ip}"
}

output "example_ssh_command" {
  value = "ssh -i $PRIVATE_KEY_PATH -o ProxyCommand=\"ssh -i $PRIVATE_KEY_PATH opc@${oci_core_instance.bastion.public_ip} -W %h:%p %r\" opc@${oci_core_instance.private.private_ip}"
}
