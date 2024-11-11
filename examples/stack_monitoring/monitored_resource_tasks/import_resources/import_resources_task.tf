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

resource "oci_stack_monitoring_monitored_resource_task" "test_monitored_resource_task" {
  #Required
  compartment_id = var.compartment_ocid
  name = var.task_name
  task_details  {
    namespace = "oci_terraform_namespace"
    type = "IMPORT_OCI_TELEMETRY_RESOURCES"
    source = "OCI_TELEMETRY_NATIVE"
    console_path_prefix = "consolePathPrefix"
    external_id_mapping = "id"
    resource_group = "tf_group"
    service_base_url = "http://test.com"
}

  #Optional
  freeform_tags = { "bar-key" = "test_monitored_resource_task.value" }
  lifecycle  {
    ignore_changes = [
      freeform_tags, defined_tags, system_tags
    ]

  }
}

data "oci_stack_monitoring_monitored_resource_tasks" "test_monitored_resource_tasks" {
  #Required
  compartment_id = oci_stack_monitoring_monitored_resource_task.test_monitored_resource_task.compartment_id
  #Optional
  status= "SUCCEEDED"
}
