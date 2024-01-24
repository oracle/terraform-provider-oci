// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

#Use Route Table Attachment to avoid cyclic dependency between Subnet and Route Table

# This example uses the `oci_core_route_table_attachment` resource to resolve this dependency cycle problem:
#   oci_core_vnic_attachment.example_vnic_attachment
#   oci_core_private_ip.private_ip
#   oci_core_route_table.example_route_table
#   oci_core_subnet.example_subnet
#   oci_core_instance.example_instance

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

data "oci_identity_availability_domain" "ad" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}

variable "instance_shape" {
  default = "VM.Standard2.1"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_core_vcn" "example_vcn" {
  cidr_block     = "10.1.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "exampleVCN"
  dns_label      = "tfexamplevcn"
}

resource "oci_core_subnet" "example_subnet" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  cidr_block          = "10.1.20.0/24"
  display_name        = "exampleSubnet"
  dns_label           = "tfexamplesubnet"

  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.example_vcn.id
}

resource "oci_core_route_table_attachment" "example_route_table_attachment" {
  subnet_id      = oci_core_subnet.example_subnet.id
  route_table_id = oci_core_route_table.example_route_table.id
}

resource "oci_core_private_ip" "private_ip" {
  vnic_id        = oci_core_vnic_attachment.example_vnic_attachment.vnic_id
  display_name   = "someDisplayName"
  hostname_label = "somehostnamelabel"
}

resource "oci_core_route_table" "example_route_table" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.example_vcn.id
  display_name   = "exampleRouteTable"

  route_rules {
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = oci_core_private_ip.private_ip.id
  }
}

resource "oci_core_vnic_attachment" "example_vnic_attachment" {
  create_vnic_details {
    assign_public_ip       = true
    subnet_id              = oci_core_subnet.example_subnet.id
    skip_source_dest_check = true
  }

  instance_id = oci_core_instance.example_instance.id
}

# Create Instance
resource "oci_core_instance" "example_instance" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid
  display_name        = "testInstance"
  shape               = var.instance_shape

  create_vnic_details {
    hostname_label         = "instance"
    subnet_id              = oci_core_subnet.example_subnet.id
    skip_source_dest_check = true
    assign_public_ip       = true
  }

  source_details {
    source_type = "image"
    source_id   = var.instance_image_ocid[var.region]
  }
}

