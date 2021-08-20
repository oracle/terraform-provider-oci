// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

variable "config_config_type" {
  default = "METRIC_GROUP"
}

variable "config_defined_tags_value" {
  default = "value"
}

variable "config_description" {
  default = "description"
}

variable "config_dimensions_name" {
  default = "name"
}

variable "config_dimensions_value_source" {
  default = "valueSource"
}

variable "config_display_name" {
  default = "displayName"
}

variable "config_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "config_metrics_description" {
  default = "description"
}

variable "config_metrics_name" {
  default = "ThreadCpuTime"
}

variable "config_metrics_unit" {
  default = "ms"
}

variable "config_metrics_value_source" {
  default = "valueSource"
}

variable "config_namespace" {
  default = "oracle_apm_monitoring"
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


// We need to create a Span Filter first, since a Metric Group requires
// an existing filter
//
resource "oci_apm_config_config" "test_filter" {
  #Required
  apm_domain_id = oci_apm_apm_domain.test_apm_domain.id
  config_type   = "SPAN_FILTER"
  display_name  = "filterName"

  filter_text   = "kind='SERVER'"
}

data "oci_apm_config_config" "test_filter" {
  apm_domain_id = oci_apm_apm_domain.test_apm_domain.id
  config_id = oci_apm_config_config.test_filter.id
}

resource "oci_apm_config_config" "test_config" {
  #Required
  apm_domain_id = oci_apm_apm_domain.test_apm_domain.id
  config_type   = var.config_config_type
  display_name  = var.config_display_name

  # Optional
  description   = var.config_description
  filter_id     = data.oci_apm_config_config.test_filter.id
  freeform_tags = var.config_freeform_tags
  metrics {

    # Optional
    description  = var.config_metrics_description
    name         = var.config_metrics_name
    #unit        = var.config_metrics_unit
  }

  namespace   = var.config_namespace
}

data "oci_apm_config_configs" "test_configs" {
  #Required
  apm_domain_id = oci_apm_apm_domain.test_apm_domain.id

  #Optional
  config_type  = var.config_config_type
  display_name = var.config_display_name
}
