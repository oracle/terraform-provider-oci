// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}
variable "region" {}
variable "ssh_public_key" {}

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

variable "mgmt_subnet_cidr" {
  default = "10.0.0.0/24"
}

variable "private_subnet_cidr" {
  default = "10.0.1.0/24"
}

provider "oci" {
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region           = "${var.region}"
}

data "oci_identity_availability_domain" "ad" {
  compartment_id = "${var.tenancy_ocid}"
  ad_number      = 1
}

resource "oci_core_virtual_network" "CoreVCN" {
  cidr_block     = "${var.vcn_cidr}"
  compartment_id = "${var.compartment_ocid}"
  display_name   = "mgmt-vcn"
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
      "max" = 80
      "min" = 80
    }

    protocol = "6"
    source   = "0.0.0.0/0"
  },
    {
      tcp_options {
        "max" = 443
        "min" = 443
      }

      protocol = "6"
      source   = "0.0.0.0/0"
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

resource "oci_core_subnet" "MgmtSubnet" {
  availability_domain = "${data.oci_identity_availability_domain.ad.name}"
  cidr_block          = "${var.mgmt_subnet_cidr}"
  display_name        = "MgmtSubnet"
  compartment_id      = "${var.compartment_ocid}"
  vcn_id              = "${oci_core_virtual_network.CoreVCN.id}"
  route_table_id      = "${oci_core_route_table.MgmtRouteTable.id}"
  security_list_ids   = ["${oci_core_security_list.MgmtSecurityList.id}"]
  dhcp_options_id     = "${oci_core_virtual_network.CoreVCN.default_dhcp_options_id}"
}

resource "oci_core_instance" "NatInstance" {
  availability_domain = "${data.oci_identity_availability_domain.ad.name}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "NatInstance"
  shape               = "${var.instance_shape}"

  create_vnic_details {
    subnet_id              = "${oci_core_subnet.MgmtSubnet.id}"
    skip_source_dest_check = true
  }

  metadata {
    ssh_authorized_keys = "${var.ssh_public_key}"
    user_data           = "${base64encode(file("user_data.tpl"))}"
  }

  source_details {
    source_type = "image"
    source_id   = "${var.instance_image_ocid[var.region]}"
  }

  timeouts {
    create = "10m"
  }
}

# Gets a list of VNIC attachments on the instance
data "oci_core_vnic_attachments" "NatInstanceVnics" {
  compartment_id      = "${var.compartment_ocid}"
  availability_domain = "${data.oci_identity_availability_domain.ad.name}"
  instance_id         = "${oci_core_instance.NatInstance.id}"
}

# Create PrivateIP
resource "oci_core_private_ip" "NatInstancePrivateIP" {
  vnic_id      = "${lookup(data.oci_core_vnic_attachments.NatInstanceVnics.vnic_attachments[0],"vnic_id")}"
  display_name = "NatInstancePrivateIP"
}

resource "oci_core_security_list" "PrivateSecurityList" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "PrivateSecurityList"
  vcn_id         = "${oci_core_virtual_network.CoreVCN.id}"

  egress_security_rules = [{
    protocol    = "all"
    destination = "0.0.0.0/0"
  }]

  ingress_security_rules = [{
    protocol = "6"

    tcp_options {
      "max" = 22
      "min" = 22
    }

    source = "${var.vcn_cidr}"
  }]
}

resource "oci_core_route_table" "PrivateRouteTable" {
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_virtual_network.CoreVCN.id}"
  display_name   = "PrivateRouteTable"

  route_rules {
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = "${oci_core_private_ip.NatInstancePrivateIP.id}"
  }
}

resource "oci_core_subnet" "PrivateSubnet" {
  cidr_block                 = "${var.private_subnet_cidr}"
  display_name               = "PrivateSubnet"
  compartment_id             = "${var.compartment_ocid}"
  vcn_id                     = "${oci_core_virtual_network.CoreVCN.id}"
  route_table_id             = "${oci_core_route_table.PrivateRouteTable.id}"
  security_list_ids          = ["${oci_core_security_list.PrivateSecurityList.id}"]
  dhcp_options_id            = "${oci_core_virtual_network.CoreVCN.default_dhcp_options_id}"
  prohibit_public_ip_on_vnic = "true"
}

resource "oci_core_instance" "PrivateInstance" {
  availability_domain = "${data.oci_identity_availability_domain.ad.name}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "PrivateInstance"
  shape               = "${var.instance_shape}"

  create_vnic_details {
    subnet_id        = "${oci_core_subnet.PrivateSubnet.id}"
    assign_public_ip = false
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
