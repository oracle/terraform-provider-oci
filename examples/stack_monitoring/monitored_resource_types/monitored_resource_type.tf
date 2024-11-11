// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_stack_monitoring_monitored_resource_type" "test_monitored_resource_type" {
  #Required
  compartment_id = var.compartment_ocid
  name = "terraformResourceTypeExample"

  #Optional
  description = "This is a resource type created for terraform test"
  display_name = "Terraform Resource Type Example"
  metric_namespace = "oci_terraform_test"
  source_type = "SM_MGMT_AGENT_MONITORED"
  resource_category = "APPLICATION"
  freeform_tags = { "bar-key" = "test_monitored_resource_type.value" }
  metadata  {
    format = "SYSTEM_FORMAT"
    agent_properties = [ "agent_prop1", "agent_prop2"]
    required_properties = [ "required_prop1", "required_prop2"]
    valid_properties_for_create = [ "valid_prop1", "valid_prop2"]
    valid_properties_for_update = ["valid_prop2"]

  }
  lifecycle  {
    ignore_changes = [
      freeform_tags, defined_tags, system_tags, id, name, state, time_created, time_updated, metadata
    ]

  }
}

data "oci_stack_monitoring_monitored_resource_types" "test_monitored_resource_types" {
  #Required
  compartment_id = oci_stack_monitoring_monitored_resource_type.test_monitored_resource_type.compartment_id
  #Optional
  name = oci_stack_monitoring_monitored_resource_type.test_monitored_resource_type.name
  status= "ACTIVE"
}
