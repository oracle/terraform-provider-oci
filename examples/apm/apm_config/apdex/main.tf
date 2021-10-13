// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" { }
variable "user_ocid" { }
variable "fingerprint" { }
variable "private_key_path" { }
variable "region" { }
variable "compartment_ocid" { }

variable "config_config_type" {
  default = "APDEX"
}

variable "config_defined_tags_value" {
  default = "value"
}

variable "config_display_name" {
  default = "displayName"
}

variable "config_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "rule_display_name" {
  default = "rule name"
}

variable "rule_is_apply_to_error_spans" {
  default = true
}

variable "rule_is_enabled" {
  default = true
}

variable "rule_priority" {
  default = 1
}

variable "rule_satisfied_response_time" {
  default = 2000
}

variable "rule_tolerating_response_time" {
  default = 5000
}

variable "rule_filter_text" {
  default = "kind='SERVER'"
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

resource "oci_apm_config_config" "test_apdex" {
  #Required
  apm_domain_id = oci_apm_apm_domain.test_apm_domain.id
  config_type   = var.config_config_type
  display_name  = var.config_display_name

  #Optional
  #defined_tags = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.config_defined_tags_value)
  freeform_tags = var.config_freeform_tags

  rules {
      display_name = var.rule_display_name
      filter_text = var.rule_filter_text
      is_apply_to_error_spans = var.rule_is_apply_to_error_spans
      is_enabled = var.rule_is_enabled
      priority = var.rule_priority
      satisfied_response_time = var.rule_satisfied_response_time 
      tolerating_response_time = var.rule_tolerating_response_time 
  }
}

data "oci_apm_config_config" "test_apdex" {
  apm_domain_id = oci_apm_apm_domain.test_apm_domain.id
  config_id = oci_apm_config_config.test_apdex.id
}


data "oci_apm_config_configs" "test_configs" {
  #Required
  apm_domain_id = oci_apm_apm_domain.test_apm_domain.id

  #Optional
  config_type  = data.oci_apm_config_config.test_apdex.config_type
  display_name = data.oci_apm_config_config.test_apdex.display_name
}
