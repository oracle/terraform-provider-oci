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

variable "ping_monitor_defined_tags_value" {
  default = "value"
}

variable "ping_monitor_display_name" {
  default = "displayName"
}

variable "ping_monitor_freeform_tags" {
  default = {
    "Department" = "Finance"
  }
}

variable "ping_monitor_interval_in_seconds" {
  default = 10
}

variable "ping_monitor_is_enabled" {
  default = false
}

variable "ping_monitor_port" {
  default = 80
}

variable "ping_monitor_protocol" {
  default = "TCP"
}

variable "ping_monitor_targets" {
  default = ["www.oracle.com"]
}

variable "ping_monitor_timeout_in_seconds" {
  default = 10
}

variable "ping_monitor_vantage_point_names" {
  default = ["goo-chs"]
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_health_checks_ping_monitor" "test_ping_monitor" {
  #Required
  compartment_id      = var.compartment_ocid
  display_name        = var.ping_monitor_display_name
  interval_in_seconds = var.ping_monitor_interval_in_seconds
  protocol            = var.ping_monitor_protocol
  targets             = var.ping_monitor_targets

  #Optional
  #defined_tags  = {"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"}
  freeform_tags = var.ping_monitor_freeform_tags

  is_enabled          = var.ping_monitor_is_enabled
  port                = var.ping_monitor_port
  timeout_in_seconds  = var.ping_monitor_timeout_in_seconds
  vantage_point_names = var.ping_monitor_vantage_point_names
}

data "oci_health_checks_ping_monitors" "test_ping_monitors" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name = oci_health_checks_ping_monitor.test_ping_monitor.display_name
}

output "monitors" {
  value = data.oci_health_checks_ping_monitors.test_ping_monitors.ping_monitors
}

