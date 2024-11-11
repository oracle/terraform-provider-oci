// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0
variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "task_name" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_stack_monitoring_monitored_resource_task" "test_update_resource_types_config_task" {
  #Required
  compartment_id = var.compartment_ocid
  name = var.task_name
  task_details {
    type         = "UPDATE_RESOURCE_TYPE_CONFIGS"
    handler_type = "TELEGRAF"
    resource_types_configuration {
      resource_type = "telegraf_abc"
      handler_config {
        collector_types = [ "abc" ]
        telemetry_resource_group = "telegraf_abc"
      }
    }
  }

  #Optional
  freeform_tags = { "bar-key" = "test_update_resource_types_config_task.value" }
  lifecycle  {
    ignore_changes = [
      freeform_tags, defined_tags, system_tags
    ]

  }
}

data "oci_stack_monitoring_monitored_resource_tasks" "test_update_resource_types_config_tasks" {
  #Required
  compartment_id = oci_stack_monitoring_monitored_resource_task.test_update_resource_types_config_task.compartment_id
  #Optional
  status= "SUCCEEDED"
}
