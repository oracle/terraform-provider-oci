// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" { }
variable "user_ocid" { }
variable "fingerprint" { }
variable "private_key_path" { }
variable "region" { }
variable "compartment_ocid" { }

variable "config_config_type" {
  default = "SPAN_FILTER"
}

variable "config_defined_tags_value" {
  default = "value"
}

variable "config_description" {
  default = "description"
}

variable "config_display_name" {
  default = "displayName"
}

variable "config_filter_text" {
  # Must be a valid filter
  default = "kind='SERVER'"
}

variable "config_freeform_tags" {
  default = { "bar-key" = "value" }
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
  description   = var.apm_domain_description
  freeform_tags = var.apm_domain_freeform_tags
  is_free_tier  = var.apm_domain_is_free_tier
}


resource "oci_apm_config_config" "test_config" {
  #Required
  apm_domain_id = oci_apm_apm_domain.test_apm_domain.id
  config_type   = var.config_config_type
  display_name  = var.config_display_name

  #Optional
  #defined_tags = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.config_defined_tags_value)
  description  = var.config_description
  filter_text   = var.config_filter_text
  freeform_tags = var.config_freeform_tags
}

data "oci_apm_config_configs" "test_configs" {
  #Required
  apm_domain_id = oci_apm_apm_domain.test_apm_domain.id

  #Optional
  config_type  = var.config_config_type
  display_name = var.config_display_name
}
