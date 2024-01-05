// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "migration_defined_tags_value" {
  default = "value"
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
  default = "AVAILABLE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_cloud_migrations_migration" "test_migration" {
  #Required
  compartment_id = var.compartment_id
  display_name   = var.migration_display_name

  #Optional
  defined_tags            = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.migration_defined_tags_value)
  freeform_tags           = var.migration_freeform_tags
  is_completed            = var.migration_is_completed
  replication_schedule_id = oci_cloud_migrations_replication_schedule.test_replication_schedule.id
}

data "oci_cloud_migrations_migrations" "test_migrations" {

  #Optional
  compartment_id = var.compartment_id
  display_name   = var.migration_display_name
  migration_id   = oci_cloud_migrations_migration.test_migration.id
  state          = var.migration_state
}

