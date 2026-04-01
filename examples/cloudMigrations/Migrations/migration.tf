// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "region" {
  default = "us-ashburn-1"
}

variable "compartment_id" {
  default = "compartment_id"
}

variable "migration_display_name" {
  default = "displayName"
}

variable "migration_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "migration_is_completed" {
  default = false
}

variable "migration_state" {
  default = "ACTIVE"
}

variable "replication_schedule_id" {
  default = "replication_schedule_id"
}



provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = "terraform-federation-test"
  region              = var.region
  # version             = "8.3.0"
}

resource "oci_cloud_migrations_migration" "test_migration" {
  #Required
  compartment_id = var.compartment_id
  display_name   = var.migration_display_name

  #Optional
  freeform_tags           = var.migration_freeform_tags
  is_completed            = var.migration_is_completed
  replication_schedule_id = var.replication_schedule_id
}

data "oci_cloud_migrations_migrations" "test_migrations" {

  #Optional
  compartment_id = var.compartment_id
  display_name   = var.migration_display_name
  migration_id   = oci_cloud_migrations_migration.test_migration.id
  state          = var.migration_state
}

