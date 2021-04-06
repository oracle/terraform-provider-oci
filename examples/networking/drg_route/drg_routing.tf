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
  tenancy_ocid = var.tenancy_ocid
  user_ocid = var.user_ocid
  fingerprint = var.fingerprint
  private_key_path = var.private_key_path
  region = var.region
}

variable "instance_image_ocid" {
  type = map(string)

  default = {
    # See https://docs.us-phoenix-1.oraclecloud.com/images/
    # Oracle-provided image "Oracle-Linux-7.5-2018.10.16-0"
    us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaaoqj42sokaoh42l76wsyhn3k2beuntrh5maj3gmgmzeyr55zzrwwa"
    us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaageeenzyuxgia726xur4ztaoxbxyjlxogdhreu3ngfj2gji3bayda"
    eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaaitzn6tdyjer7jl34h2ujz74jwy5nkbukbh55ekp6oyzwrtfa4zma"
    uk-london-1 = "ocid1.image.oc1.uk-london-1.aaaaaaaa32voyikkkzfxyo4xbdmadc2dmvorfxxgdhpnk6dw64fa3l4jh7wa"
  }
}

resource "oci_core_vcn" "test_vcn_a" {
  // required
  cidr_block = "10.0.0.0/16"
  compartment_id = var.compartment_ocid

  // optional
  display_name = "testVcnA"
  dns_label = "dnslabelA"
}


resource "oci_core_vcn" "test_vcn_b" {
  // required
  cidr_block = "20.0.0.0/16"
  compartment_id = var.compartment_ocid

  // optional
  display_name = "testVcnB"
  dns_label = "dnslabelB"
}


resource "oci_core_drg" "test_drg" {
  // required
  compartment_id = var.compartment_ocid

  // optional
  display_name = "testDrg"
}

resource "oci_core_drg_route_distribution" "test_drg_route_distribution" {
  // required
  drg_id = oci_core_drg.test_drg.id
  distribution_type = "IMPORT"

  // optional
  display_name = "testDrgRouteDistribution"

}

resource "oci_core_drg_attachment" "test_drg_attachment_a" {
  // required
  drg_id = oci_core_drg.test_drg.id
  vcn_id = oci_core_vcn.test_vcn_a.id

}

resource "oci_core_drg_attachment" "test_drg_attachment_b" {
  // required
  drg_id = oci_core_drg.test_drg.id
  vcn_id = oci_core_vcn.test_vcn_b.id

}

resource "oci_core_drg_route_distribution_statement" "test_drg_route_distribution_statements" {
  // required
  drg_route_distribution_id = oci_core_drg_route_distribution.test_drg_route_distribution.id
  action = "ACCEPT"

  match_criteria {
    match_type= "DRG_ATTACHMENT_TYPE"
    attachment_type = "VCN"
  }

  priority = 10


}

data "oci_core_drg_route_distribution" "test_drg_route_distribution_data" {
  // required
  drg_route_distribution_id = oci_core_drg_route_distribution.test_drg_route_distribution.id
}

resource "oci_core_drg_route_table" "test_drg_route_table" {
  drg_id = oci_core_drg.test_drg.id
}

resource "oci_core_drg_route_table_route_rule" "test_drg_route_table_route_rule" {
  #Required
  drg_route_table_id = oci_core_drg_route_table.test_drg_route_table.id
  destination                = "10.0.0.0/8"
  destination_type           = "CIDR_BLOCK"
  next_hop_drg_attachment_id = oci_core_drg_attachment.test_drg_attachment_a.id
}
