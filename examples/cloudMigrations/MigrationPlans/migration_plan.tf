// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "region" {
  default = "us-ashburn-1"
}

variable "compartment_id" {
  default = "compartment_id"
}

variable "migration_plan_display_name" {
  default = "displayName"
}

variable "migration_plan_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "migration_plan_state" {
  default = "ACTIVE"
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
  default = "oQNt:US-ASHBURN-AD-1"
}

variable "migration_plan_target_environments_ms_license" {
  default = "msLicense"
}

variable "migration_plan_target_environments_preferred_shape_type" {
  default = "VM"
}

variable "migration_plan_target_environments_subnet" {
  default = "migration_plan_target_environments_subnet"
}

variable "migration_plan_target_environments_target_environment_type" {
  default = "VM_TARGET_ENV"
}

variable "migration_plan_target_environments_vcn" {
  default = "migration_plan_target_environments_vcn"
}

variable "migration_id" {
  default = "migration_id"
}



provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = "terraform-federation-test"
  region              = var.region
  # version             = "8.3.0"
}

resource "oci_cloud_migrations_migration_plan" "test_migration_plan" {
  #Required
  compartment_id = var.compartment_id
  display_name   = var.migration_plan_display_name
  migration_id   = var.migration_id

  #Optional
  freeform_tags            = var.migration_plan_freeform_tags
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
    ms_license            = var.migration_plan_target_environments_ms_license
    preferred_shape_type  = var.migration_plan_target_environments_preferred_shape_type
    target_compartment_id = var.compartment_id
  }
}

data "oci_cloud_migrations_migration_plans" "test_migration_plans" {

  #Optional
  compartment_id    = var.compartment_id
  display_name      = var.migration_plan_display_name
  migration_id      = var.migration_id
  migration_plan_id = oci_cloud_migrations_migration_plan.test_migration_plan.id
  state             = var.migration_plan_state
}

