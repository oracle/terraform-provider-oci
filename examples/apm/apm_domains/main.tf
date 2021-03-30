// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

variable "apm_domain_defined_tags_value" {
  default = "definedTags"
}

variable "apm_domain_description" {
  default = "description"
}

variable "apm_domain_display_name" {
  default = "displayName"
}

variable "apm_domain_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "apm_domain_is_free_tier" {
  default = false
}

variable "apm_domain_state" {
  default = "ACTIVE"
}

variable "data_key_data_key_type" {
  default = "PRIVATE"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_apm_apm_domain" "test_apm_domain" {
  #Required
  compartment_id = var.compartment_ocid
  display_name   = var.apm_domain_display_name

  #Optional
  #defined_tags  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.apm_domain_defined_tags_value)
  description   = var.apm_domain_description
  freeform_tags = var.apm_domain_freeform_tags
  is_free_tier  = var.apm_domain_is_free_tier
}

data "oci_apm_apm_domains" "test_apm_domains" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name = var.apm_domain_display_name
  state        = var.apm_domain_state
}

data "oci_apm_data_keys" "test_data_keys" {
  #Required
  apm_domain_id = oci_apm_apm_domain.test_apm_domain.id

  #Optional
  data_key_type = var.data_key_data_key_type
}
