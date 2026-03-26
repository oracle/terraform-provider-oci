// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {
}

variable "occ_metric_alarm_defined_tags_value" {
  default = "value"
}

variable "occ_metric_alarm_description" {
  default = "description"
}

variable "occ_metric_alarm_display_name" {
  default = "displayName"
}

variable "occ_metric_alarm_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "occ_metric_alarm_frequency" {
  default = "WEEKLY"
}

variable "occ_metric_alarm_is_active" {
  default = false
}

variable "occ_metric_alarm_resource_configuration_compute_hw_generation" {
  default = "computeHwGeneration"
}

variable "occ_metric_alarm_resource_configuration_hw_generation" {
  default = "hwGeneration"
}

variable "occ_metric_alarm_resource_configuration_node_type" {
  default = "nodeType"
}

variable "occ_metric_alarm_resource_configuration_occ_metric_alarm_provider" {
  default = "occMetricAlarmProvider"
}

variable "occ_metric_alarm_resource_configuration_resource" {
  default = "COMPUTE"
}

variable "occ_metric_alarm_resource_configuration_shape" {
  default = "shape"
}

variable "occ_metric_alarm_resource_configuration_storage_type" {
  default = "storageType"
}

variable "occ_metric_alarm_resource_configuration_usage_type" {
  default = "usageType"
}

variable "occ_metric_alarm_state" {
  default = "ACTIVE"
}

variable "occ_metric_alarm_subscribers" {
  default = ["topic.dsadsads.sdadasd.dsad"]
}

variable "occ_metric_alarm_threshold" {
  default = 10
}

variable "occ_metric_alarm_threshold_type" {
  default = "PERCENTAGE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_demand_signal_occ_metric_alarm" "test_occ_metric_alarm" {
  #Required
  compartment_id = var.compartment_id
  display_name   = var.occ_metric_alarm_display_name
  frequency      = var.occ_metric_alarm_frequency
  is_active      = var.occ_metric_alarm_is_active
  resource_configuration {
    #Required
    resource   = var.occ_metric_alarm_resource_configuration_resource
    usage_type = var.occ_metric_alarm_resource_configuration_usage_type

    #Optional
    compute_hw_generation     = var.occ_metric_alarm_resource_configuration_compute_hw_generation
    hw_generation             = var.occ_metric_alarm_resource_configuration_hw_generation
    node_type                 = var.occ_metric_alarm_resource_configuration_node_type
    shape                     = var.occ_metric_alarm_resource_configuration_shape
    storage_type              = var.occ_metric_alarm_resource_configuration_storage_type
  }
  threshold = var.occ_metric_alarm_threshold

  #Optional
  description    = var.occ_metric_alarm_description
  freeform_tags  = var.occ_metric_alarm_freeform_tags
  state          = var.occ_metric_alarm_state
  subscribers    = var.occ_metric_alarm_subscribers
  threshold_type = var.occ_metric_alarm_threshold_type
}

data "oci_demand_signal_occ_metric_alarms" "test_occ_metric_alarms" {
  #Required
  compartment_id = var.compartment_id
}