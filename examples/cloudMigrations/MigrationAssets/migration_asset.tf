// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "migration_asset_availability_domain" {
  default = "availabilityDomain"
}

variable "migration_asset_depends_on" {
  default = []
}

variable "migration_asset_display_name" {
  default = "displayName"
}

variable "migration_asset_state" {
  default = "AVAILABLE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_cloud_migrations_migration_asset" "test_migration_asset" {
  #Required
  availability_domain        = var.migration_asset_availability_domain
  inventory_asset_id         = oci_cloud_migrations_inventory_asset.test_inventory_asset.id
  migration_id               = oci_cloud_migrations_migration.test_migration.id
  replication_compartment_id = oci_identity_compartment.test_compartment.id
  snap_shot_bucket_name      = oci_objectstorage_bucket.test_bucket.name

  #Optional
  depends_on              = var.migration_asset_depends_on
  display_name            = var.migration_asset_display_name
  replication_schedule_id = oci_cloud_migrations_replication_schedule.test_replication_schedule.id
}

data "oci_cloud_migrations_migration_assets" "test_migration_assets" {

  #Optional
  display_name       = var.migration_asset_display_name
  migration_asset_id = oci_cloud_migrations_migration_asset.test_migration_asset.id
  migration_id       = oci_cloud_migrations_migration.test_migration.id
  state              = var.migration_asset_state
}

