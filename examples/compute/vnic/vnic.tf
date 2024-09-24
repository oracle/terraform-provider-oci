// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "compartment_ocid" {
}

variable "region" {
}

variable "ssh_public_key" {
}

variable "secondary_vnic_count" {
  default = 1
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

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_identity_availability_domain" "ad" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "CompleteVCN"
  dns_label      = "examplevcn"
}

resource "oci_core_subnet" "test_subnet" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  cidr_block          = "10.0.1.0/24"
  display_name        = "TestSubnet"
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.test_vcn.id
  route_table_id      = oci_core_vcn.test_vcn.default_route_table_id
  security_list_ids   = [oci_core_vcn.test_vcn.default_security_list_id]
  dhcp_options_id     = oci_core_vcn.test_vcn.default_dhcp_options_id
  dns_label           = "examplesubnet"
}

resource "oci_core_network_security_group" "test_network_security_group" {
  #Required
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.test_vcn.id

  #Optional
  display_name = "TestNetworkSecurityGroup"
}

resource "oci_core_instance" "test_instance" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "TestInstance"
  shape               = "VM.Standard2.1"

  source_details {
    source_type = "image"
    source_id   = var.instance_image_ocid[var.region]
  }

  create_vnic_details {
    subnet_id      = oci_core_subnet.test_subnet.id
    hostname_label = "testinstance"

    security_attributes = {
      "oracle-zpr.sensitivity.value" = "low"
      "oracle-zpr.sensitivity.mode" = "enforce"
    }
  }

  metadata = {
    ssh_authorized_keys = var.ssh_public_key
  }

  timeouts {
    create = "60m"
  }
}

resource "oci_core_vnic_attachment" "secondary_vnic_attachment" {
  instance_id  = oci_core_instance.test_instance.id
  display_name = "SecondaryVnicAttachment_${count.index}"

  create_vnic_details {
    subnet_id                 = oci_core_subnet.test_subnet.id
    display_name              = "SecondaryVnic_${count.index}"
    assign_public_ip          = true
    skip_source_dest_check    = true
    assign_private_dns_record = true
    nsg_ids                   = [oci_core_network_security_group.test_network_security_group.id]
  }

  count = var.secondary_vnic_count
}

data "oci_core_vnic" "secondary_vnic" {
  count = var.secondary_vnic_count
  vnic_id = element(
    oci_core_vnic_attachment.secondary_vnic_attachment.*.vnic_id,
    count.index,
  )
}

output "primary_ip_addresses" {
  value = [
    oci_core_instance.test_instance.public_ip,
    oci_core_instance.test_instance.private_ip,
  ]
}

output "secondary_public_ip_addresses" {
  value = [data.oci_core_vnic.secondary_vnic.*.public_ip_address]
}

output "secondary_private_ip_addresses" {
  value = [data.oci_core_vnic.secondary_vnic.*.private_ip_address]
}

