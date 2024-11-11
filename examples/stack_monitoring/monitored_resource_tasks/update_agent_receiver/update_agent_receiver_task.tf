// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0
variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "task_name" {}
variable "stack_mon_management_agent_id_resource1" {}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_stack_monitoring_monitored_resource_task" "test_update_agent_receiver_task" {
  #Required
  compartment_id = var.compartment_ocid
  name = var.task_name
  task_details  {
    type = "UPDATE_AGENT_RECEIVER"
    handler_type = "TELEGRAF"
    agent_id =var.stack_mon_management_agent_id_resource1
    is_enable = false
    #receiver_properties {
    #
    #}
  }

  #Optional
  freeform_tags = { "bar-key" = "test_update_agent_receiver_task.value" }
  lifecycle  {
    ignore_changes = [
      freeform_tags, defined_tags, system_tags
    ]

  }
}

data "oci_stack_monitoring_monitored_resource_tasks" "test_update_agent_receiver_tasks" {
  #Required
  compartment_id = oci_stack_monitoring_monitored_resource_task.test_update_agent_receiver_task.compartment_id
  #Optional
  status= "SUCCEEDED"
}
