// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "batch_context_defined_tags_value" {
  default = "value"
}

variable "batch_context_description" {
  default = "description"
}

variable "batch_context_display_name" {
  default = "displayName"
}

variable "batch_context_entitlements" {
  default = "entitlements"
}

variable "batch_context_fleets_max_concurrent_tasks" {
  default = 10
}

variable "batch_context_fleets_name" {
  default = "name"
}

variable "batch_context_fleets_shape_disk_size_in_gbs" {
  default = 10
}

variable "batch_context_fleets_shape_memory_in_gbs" {
  default = 10
}

variable "batch_context_fleets_shape_ocpus" {
  default = 10
}

variable "batch_context_fleets_shape_type" {
  default = "FIXED_GPU_FLEET_SHAPE"
}

variable "batch_context_fleets_type" {
  default = "SERVICE_MANAGED_FLEET"
}

variable "batch_context_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "batch_context_id" {
  default = "id"
}

variable "batch_context_job_priority_configurations_tag_key" {
  default = "tagKey"
}

variable "batch_context_job_priority_configurations_tag_namespace" {
  default = "tagNamespace"
}

variable "batch_context_job_priority_configurations_values" {
  default = "values"
}

variable "batch_context_job_priority_configurations_weight" {
  default = 10
}

variable "batch_context_logging_configuration_is_job_task_events_propagation_enabled" {
  default = false
}

variable "batch_context_logging_configuration_type" {
  default = "OCI_LOGGING"
}

variable "batch_context_network_nsg_ids" {
  default = []
}

variable "batch_context_state" {
  default = "AVAILABLE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_batch_batch_context" "test_batch_context" {
  #Required
  compartment_id = var.compartment_id
  fleets {
    #Required
    max_concurrent_tasks = var.batch_context_fleets_max_concurrent_tasks
    name                 = var.batch_context_fleets_name
    shape {
      #Required
      memory_in_gbs = var.batch_context_fleets_shape_memory_in_gbs
      ocpus         = var.batch_context_fleets_shape_ocpus
      type          = var.batch_context_fleets_shape_type

      #Optional
      disk_size_in_gbs = var.batch_context_fleets_shape_disk_size_in_gbs
      shape_name       = oci_core_shape.test_shape.name
    }
    type = var.batch_context_fleets_type
  }
  network {
    #Required
    subnet_id = oci_core_subnet.test_subnet.id

    #Optional
    nsg_ids = var.batch_context_network_nsg_ids
  }

  #Optional
  defined_tags  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.batch_context_defined_tags_value)
  description   = var.batch_context_description
  display_name  = var.batch_context_display_name
  entitlements  = var.batch_context_entitlements
  freeform_tags = var.batch_context_freeform_tags
  job_priority_configurations {
    #Required
    tag_key       = var.batch_context_job_priority_configurations_tag_key
    tag_namespace = var.batch_context_job_priority_configurations_tag_namespace
    values        = var.batch_context_job_priority_configurations_values
    weight        = var.batch_context_job_priority_configurations_weight
  }
  logging_configuration {
    #Required
    log_group_id = oci_logging_log_group.test_log_group.id
    log_id       = oci_apm_traces_log.test_log.id
    type         = var.batch_context_logging_configuration_type

    #Optional
    is_job_task_events_propagation_enabled = var.batch_context_logging_configuration_is_job_task_events_propagation_enabled
  }
}

data "oci_batch_batch_contexts" "test_batch_contexts" {

  #Optional
  compartment_id = var.compartment_id
  display_name   = var.batch_context_display_name
  id             = var.batch_context_id
  state          = var.batch_context_state
}

