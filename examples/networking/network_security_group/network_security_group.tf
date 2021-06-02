// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
}

variable "compartment_ocid" {
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

variable "instance_image_ocid" {
  type = map(string)

  default = {
    # See https://docs.us-phoenix-1.oraclecloud.com/images/
    # Oracle-provided image "Oracle-Linux-7.5-2018.10.16-0"
    us-phoenix-1   = "ocid1.image.oc1.phx.aaaaaaaaoqj42sokaoh42l76wsyhn3k2beuntrh5maj3gmgmzeyr55zzrwwa"
    us-ashburn-1   = "ocid1.image.oc1.iad.aaaaaaaageeenzyuxgia726xur4ztaoxbxyjlxogdhreu3ngfj2gji3bayda"
    eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaaitzn6tdyjer7jl34h2ujz74jwy5nkbukbh55ekp6oyzwrtfa4zma"
    uk-london-1    = "ocid1.image.oc1.uk-london-1.aaaaaaaa32voyikkkzfxyo4xbdmadc2dmvorfxxgdhpnk6dw64fa3l4jh7wa"
  }
}

resource "oci_core_vcn" "test_vcn" {
  #Required
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_ocid

  #Optional
  display_name = "testVcn"
  dns_label    = "dnslabel"
}

resource "oci_core_network_security_group" "test_network_security_group" {
  #Required
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.test_vcn.id

  #Optional
  display_name = "testNetworkSecurityGroup"
}

resource "oci_core_network_security_group_security_rule" "test_network_security_group_security_rule_1" {
  network_security_group_id = oci_core_network_security_group.test_network_security_group.id
  direction                 = "EGRESS"
  destination               = "10.0.0.0/16"
  protocol                  = "7"
  count                     = 5
}

resource "oci_core_network_security_group_security_rule" "test_network_security_group_security_rule_2" {
  network_security_group_id = oci_core_network_security_group.test_network_security_group.id

  direction   = "EGRESS"
  protocol    = "all"
  destination = "0.0.0.0/0"
}

resource "oci_core_network_security_group_security_rule" "test_network_security_group_security_rule_3" {
  network_security_group_id = oci_core_network_security_group.test_network_security_group.id
  protocol                  = "6"
  direction                 = "INGRESS"
  source                    = "0.0.0.0/0"
  stateless                 = true

  tcp_options {
    destination_port_range {
      min = 22
      max = 22
    }

    source_port_range {
      min = 100
      max = 100
    }
  }
}

resource "oci_core_network_security_group_security_rule" "test_network_security_group_security_rule_4" {
  network_security_group_id = oci_core_network_security_group.test_network_security_group.id
  protocol                  = 1
  direction                 = "INGRESS"
  source                    = "10.0.0.0/16"
  stateless                 = true

  icmp_options {
    type = 3
    code = 0
  }
}

resource "oci_core_network_security_group_security_rule" "test_network_security_group_security_rule_5" {
  network_security_group_id = oci_core_network_security_group.test_network_security_group.id
  destination               = "0.0.0.0/0"
  direction                 = "EGRESS"
  protocol                  = "17"
  stateless                 = false

  udp_options {
    destination_port_range {
      min = 319
      max = 320
    }

    source_port_range {
      min = 100
      max = 100
    }
  }
}

resource "oci_core_instance" "instance1" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "be-instance1"
  shape               = "VM.Standard2.1"

  create_vnic_details {
    subnet_id        = oci_core_subnet.subnet1.id
    display_name     = "primaryvnic"
    assign_public_ip = true
    nsg_ids          = [oci_core_network_security_group.test_network_security_group.id]
  }

  source_details {
    source_type = "image"
    source_id   = var.instance_image_ocid[var.region]
  }
}

resource "oci_core_subnet" "subnet1" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  cidr_block          = "10.0.0.0/24"
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.test_vcn.id

  provisioner "local-exec" {
    command = "sleep 5"
  }
}

resource "oci_core_vlan" "test_vlan" {
  #Required
  cidr_block     = "10.0.2.0/24"
  compartment_id = var.compartment_ocid
  vcn_id = oci_core_vcn.test_vcn.id

  #Optional
  display_name = "testVlan"
  availability_domain = data.oci_identity_availability_domain.ad.name
  nsg_ids = [oci_core_network_security_group.test_network_security_group.id]
}

