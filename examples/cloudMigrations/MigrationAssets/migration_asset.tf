// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "region" {
  default = "us-ashburn-1"
}

variable "compartment_id" {
  default = "compartment_id"
}

variable "migration_asset_availability_domain" {
  default = "oQNt:US-ASHBURN-AD-1"
}

variable "migration_asset_depends_on" {
  default = []
}

variable "migration_asset_display_name" {
  default = "displayName"
}

variable "migration_asset_state" {
  default = "ACTIVE"
}

variable "inventory_asset_id" {
  default = "inventory_asset_id"
}

variable "migration_id" {
  default = "migration_id"
}

variable "snap_shot_bucket_name" {
  default =  "test"
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

resource "oci_cloud_migrations_migration_asset" "test_migration_asset" {
  #Required
  availability_domain        = var.migration_asset_availability_domain
  inventory_asset_id         = var.inventory_asset_id
  migration_id               = var.migration_id
  replication_compartment_id = var.compartment_id
  snap_shot_bucket_name      = var.snap_shot_bucket_name

  #Optional
  migration_asset_depends_on = var.migration_asset_depends_on
  display_name               = var.migration_asset_display_name
  replication_schedule_id    = var.replication_schedule_id
}

data "oci_cloud_migrations_migration_assets" "test_migration_assets" {

  #Optional
  display_name       = var.migration_asset_display_name
  migration_asset_id = oci_cloud_migrations_migration_asset.test_migration_asset.id
  migration_id       = var.migration_id
  state              = var.migration_asset_state
}

