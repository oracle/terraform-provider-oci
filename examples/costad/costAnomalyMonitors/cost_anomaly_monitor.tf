// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "cost_anomaly_monitor_cost_alert_subscription_map_operator" {
  default = "AND"
}

variable "cost_anomaly_monitor_cost_alert_subscription_map_threshold_absolute_value" {
  default = 10
}

variable "cost_anomaly_monitor_cost_alert_subscription_map_threshold_relative_percent" {
  default = 10
}

variable "cost_anomaly_monitor_defined_tags_value" {
  default = "value"
}

variable "cost_anomaly_monitor_description" {
  default = "description"
}

variable "cost_anomaly_monitor_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "cost_anomaly_monitor_name" {
  default = "name"
}

variable "cost_anomaly_monitor_region" {
  default = []
}

variable "cost_anomaly_monitor_state" {
  default = "AVAILABLE"
}

variable "cost_anomaly_monitor_target_resource_filter" {
  default = "targetResourceFilter"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_costad_cost_anomaly_monitor" "test_cost_anomaly_monitor" {
  #Required
  compartment_id         = var.compartment_id
  name                   = var.cost_anomaly_monitor_name
  target_resource_filter = var.cost_anomaly_monitor_target_resource_filter

  #Optional
  cost_alert_subscription_map {

    #Optional
    cost_alert_subscription_id = oci_costad_cost_alert_subscription.test_cost_alert_subscription.id
    operator                   = var.cost_anomaly_monitor_cost_alert_subscription_map_operator
    threshold_absolute_value   = var.cost_anomaly_monitor_cost_alert_subscription_map_threshold_absolute_value
    threshold_relative_percent = var.cost_anomaly_monitor_cost_alert_subscription_map_threshold_relative_percent
  }
  defined_tags  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.cost_anomaly_monitor_defined_tags_value)
  description   = var.cost_anomaly_monitor_description
  freeform_tags = var.cost_anomaly_monitor_freeform_tags
}

data "oci_costad_cost_anomaly_monitors" "test_cost_anomaly_monitors" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  name             = var.cost_anomaly_monitor_name
  region           = var.cost_anomaly_monitor_region
  state            = var.cost_anomaly_monitor_state
  target_tenant_id = oci_costad_target_tenant.test_target_tenant.id
}

