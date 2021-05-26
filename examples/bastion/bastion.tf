// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

variable "bastion_bastion_lifecycle_state" {
  default = "ACTIVE"
}

variable "bastion_client_cidr_block_allow_list" {
  default = ["0.0.0.0/0"]
}

variable "bastion_defined_tags_value" {
  default = "value"
}

variable "bastion_name" {
  default = "bastionExample"
}

variable "bastion_freeform_tags" {
  default = {
    "bar-key" = "bastion_test"
  }
}

variable "bastion_max_session_ttl_in_seconds" {
  default = 1800
}

variable "tag_namespace_description" {
  default = "Just a test"
}

variable "tag_namespace_name" {
  default = "testexamples-tag-namespace"
}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_bastion_bastion" "test_bastion" {
  #Required
  bastion_type                   = "STANDARD"
  compartment_id                 = var.compartment_ocid
  target_subnet_id               = oci_core_subnet.test_subnet.id

  #Optional
  client_cidr_block_allow_list = var.bastion_client_cidr_block_allow_list
  defined_tags                 = {
    "${oci_identity_tag_namespace.bastion_tag_namespace1.name}.${oci_identity_tag.bastion_tag1.name}" = var.bastion_defined_tags_value
  }
  name                         = var.bastion_name
  freeform_tags                = var.bastion_freeform_tags
  max_session_ttl_in_seconds   = var.bastion_max_session_ttl_in_seconds
}

data "oci_bastion_bastions" "test_bastions" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  bastion_id              = oci_bastion_bastion.test_bastion.id
  bastion_lifecycle_state = var.bastion_bastion_lifecycle_state
  name                    = var.bastion_name
}

data "oci_core_services" "test_bastion_services" {
}

data "oci_identity_availability_domain" "bastion_ad" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}