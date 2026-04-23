// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "cost_anomaly_event_cost_impact" {
  default = 1.0
}

variable "cost_anomaly_event_cost_impact_percentage" {
  default = 1.0
}

variable "cost_anomaly_event_defined_tags_value" {
  default = "value"
}

variable "cost_anomaly_event_feedback_response" {
  default = "ACCURATE_ANOMALY"
}

variable "cost_anomaly_event_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "cost_anomaly_event_name" {
  default = "name"
}

variable "cost_anomaly_event_region" {
  default = []
}

variable "cost_anomaly_event_time_anomaly_event_end_date" {
  default = "timeAnomalyEventEndDate"
}

variable "cost_anomaly_event_time_anomaly_event_start_date" {
  default = "timeAnomalyEventStartDate"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_costad_cost_anomaly_event" "test_cost_anomaly_event" {
  #Required
  cost_anomaly_event_id = oci_costad_cost_anomaly_event.test_cost_anomaly_event.id

  #Optional
  defined_tags      = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.cost_anomaly_event_defined_tags_value)
  feedback_response = var.cost_anomaly_event_feedback_response
  freeform_tags     = var.cost_anomaly_event_freeform_tags
}

data "oci_costad_cost_anomaly_events" "test_cost_anomaly_events" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  cost_anomaly_monitor_id       = oci_costad_cost_anomaly_monitor.test_cost_anomaly_monitor.id
  cost_impact                   = var.cost_anomaly_event_cost_impact
  cost_impact_percentage        = var.cost_anomaly_event_cost_impact_percentage
  name                          = var.cost_anomaly_event_name
  region                        = var.cost_anomaly_event_region
  target_tenant_id              = oci_costad_target_tenant.test_target_tenant.id
  time_anomaly_event_end_date   = var.cost_anomaly_event_time_anomaly_event_end_date
  time_anomaly_event_start_date = var.cost_anomaly_event_time_anomaly_event_start_date
}

