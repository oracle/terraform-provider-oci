// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "migration_plan_defined_tags_value" {
  default = "value"
}

variable "migration_plan_display_name" {
  default = "displayName"
}

variable "migration_plan_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "migration_plan_state" {
  default = "AVAILABLE"
}

variable "migration_plan_strategies_adjustment_multiplier" {
  default = 1.0
}

variable "migration_plan_strategies_metric_time_window" {
  default = "1d"
}

variable "migration_plan_strategies_metric_type" {
  default = "AUTO"
}

variable "migration_plan_strategies_percentile" {
  default = "P50"
}

variable "migration_plan_strategies_resource_type" {
  default = "CPU"
}

variable "migration_plan_strategies_strategy_type" {
  default = "AS_IS"
}

variable "migration_plan_target_environments_availability_domain" {
  default = "availabilityDomain"
}

variable "migration_plan_target_environments_dedicated_vm_host" {
  default = "dedicatedVmHost"
}

variable "migration_plan_target_environments_fault_domain" {
  default = "faultDomain"
}

variable "migration_plan_target_environments_ms_license" {
  default = "msLicense"
}

variable "migration_plan_target_environments_preferred_shape_type" {
  default = "VM"
}

variable "migration_plan_target_environments_subnet" {
  default = "subnet"
}

variable "migration_plan_target_environments_target_environment_type" {
  default = "VM_TARGET_ENV"
}

variable "migration_plan_target_environments_vcn" {
  default = "vcn"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_cloud_migrations_migration_plan" "test_migration_plan" {
  #Required
  compartment_id = var.compartment_id
  display_name   = var.migration_plan_display_name
  migration_id   = oci_cloud_migrations_migration.test_migration.id

  #Optional
  defined_tags             = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.migration_plan_defined_tags_value)
  freeform_tags            = var.migration_plan_freeform_tags
  source_migration_plan_id = oci_cloud_migrations_migration_plan.test_migration_plan.id
  strategies {
    #Required
    resource_type = var.migration_plan_strategies_resource_type
    strategy_type = var.migration_plan_strategies_strategy_type

    #Optional
    adjustment_multiplier = var.migration_plan_strategies_adjustment_multiplier
    metric_time_window    = var.migration_plan_strategies_metric_time_window
    metric_type           = var.migration_plan_strategies_metric_type
    percentile            = var.migration_plan_strategies_percentile
  }
  target_environments {
    #Required
    subnet                  = var.migration_plan_target_environments_subnet
    target_environment_type = var.migration_plan_target_environments_target_environment_type
    vcn                     = var.migration_plan_target_environments_vcn

    #Optional
    availability_domain   = var.migration_plan_target_environments_availability_domain
    dedicated_vm_host     = var.migration_plan_target_environments_dedicated_vm_host
    fault_domain          = var.migration_plan_target_environments_fault_domain
    ms_license            = var.migration_plan_target_environments_ms_license
    preferred_shape_type  = var.migration_plan_target_environments_preferred_shape_type
    target_compartment_id = oci_identity_compartment.test_compartment.id
  }
}

data "oci_cloud_migrations_migration_plans" "test_migration_plans" {

  #Optional
  compartment_id    = var.compartment_id
  display_name      = var.migration_plan_display_name
  migration_id      = oci_cloud_migrations_migration.test_migration.id
  migration_plan_id = oci_cloud_migrations_migration_plan.test_migration_plan.id
  state             = var.migration_plan_state
}

