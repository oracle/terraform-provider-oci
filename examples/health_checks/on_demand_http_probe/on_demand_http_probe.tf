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

variable "region" {
}

variable "compartment_ocid" {
}

variable "http_monitor_defined_tags_value" {
  default = "value"
}

variable "http_monitor_display_name" {
  default = "displayName"
}

variable "http_monitor_freeform_tags" {
  default = {
    "Department" = "Finance"
  }
}

variable "http_monitor_headers" {
  default = "headers"
}

variable "http_monitor_interval_in_seconds" {
  default = 10
}

variable "http_monitor_is_enabled" {
  default = true
}

variable "http_monitor_method" {
  default = "GET"
}

variable "http_monitor_path" {
  default = "/"
}

variable "http_monitor_port" {
  default = "443"
}

variable "http_monitor_protocol" {
  default = "HTTPS"
}

variable "http_monitor_targets" {
  default = ["www.oracle.com"]
}

variable "http_monitor_timeout_in_seconds" {
  default = 10
}

variable "http_monitor_vantage_point_names" {
  default = ["goo-chs"]
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_health_checks_http_probe" "test_http_probe" {
  #Required
  compartment_id      = var.compartment_ocid
  protocol            = var.http_monitor_protocol
  targets             = var.http_monitor_targets
  method              = var.http_monitor_method
  path                = var.http_monitor_path
  port                = var.http_monitor_port
  timeout_in_seconds  = var.http_monitor_timeout_in_seconds
  vantage_point_names = var.http_monitor_vantage_point_names
}

data "oci_health_checks_http_probe_results" "test_http_probe_results" {
  #Required
  probe_configuration_id = oci_health_checks_http_probe.test_http_probe.id
  #Optional
  #start_time_greater_than_or_equal_to = var.http_probe_result_start_time_greater_than_or_equal_to
  #start_time_less_than_or_equal_to    = var.http_probe_result_start_time_less_than_or_equal_to
  #target                              = var.http_probe_result_target
}

output "results" {
  value = data.oci_health_checks_http_probe_results.test_http_probe_results.http_probe_results
}

