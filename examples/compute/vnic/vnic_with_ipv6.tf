// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_core_vcn" "test_vcn_with_ipv6" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "CompleteVCN"
  dns_label      = "examplevcn"
  is_ipv6enabled = true
}

resource "oci_core_subnet" "test_subnet_with_ipv6" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  cidr_block          = "10.0.1.0/24"
  display_name        = "TestSubnet"
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.test_vcn_with_ipv6.id
  route_table_id      = oci_core_vcn.test_vcn_with_ipv6.default_route_table_id
  security_list_ids   = [oci_core_vcn.test_vcn_with_ipv6.default_security_list_id]
  dhcp_options_id     = oci_core_vcn.test_vcn_with_ipv6.default_dhcp_options_id
  dns_label           = "examplesubnet"
  ipv6cidr_blocks     = ["${substr(oci_core_vcn.test_vcn_with_ipv6.ipv6cidr_blocks[0], 0, length(oci_core_vcn.test_vcn_with_ipv6.ipv6cidr_blocks[0]) - 2)}${64}"]
}

resource "oci_core_instance" "test_instance_with_ipv6" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "TestInstance"
  shape               = "VM.Standard2.1"

  source_details {
    source_type = "image"
    source_id   = var.instance_image_ocid[var.region]
  }

  create_vnic_details {
    subnet_id      = oci_core_subnet.test_subnet_with_ipv6.id
    hostname_label = "testinstance"
  }

  metadata = {
    ssh_authorized_keys = var.ssh_public_key
  }

  timeouts {
    create = "60m"
  }
}

resource "oci_core_vnic_attachment" "secondary_vnic_attachment_with_ipv6" {
  instance_id  = oci_core_instance.test_instance_with_ipv6.id
  display_name = "SecondaryVnicAttachment_${count.index}"

  create_vnic_details {
    subnet_id                 = oci_core_subnet.test_subnet_with_ipv6.id
    display_name              = "SecondaryVnic_${count.index}"
    assign_public_ip          = true
    skip_source_dest_check    = true
    assign_private_dns_record = true
    assign_ipv6ip             = true

    ipv6address_ipv6subnet_cidr_pair_details {
      ipv6_subnet_cidr = oci_core_subnet.test_subnet_with_ipv6.ipv6cidr_blocks[0]
      ipv6_address = "${substr(oci_core_subnet.test_subnet_with_ipv6.ipv6cidr_blocks[0], 0, length(oci_core_subnet.test_subnet_with_ipv6.ipv6cidr_blocks[0]) - 7)}${1000}"
    }
  }

  count = var.secondary_vnic_count
}

data "oci_core_vnic" "secondary_vnic_with_ipv6" {
  count = var.secondary_vnic_count
  vnic_id = element(
  oci_core_vnic_attachment.secondary_vnic_attachment_with_ipv6.*.vnic_id,
  count.index,
  )
}

output "secondary_ipv6_addresses" {
  value = [data.oci_core_vnic.secondary_vnic_with_ipv6.*.ipv6addresses]
}
