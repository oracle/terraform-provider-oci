// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "long_term_backup_defined_tags_value" {
  default = "value"
}

variable "long_term_backup_display_name" {
  default = "displayName"
}

variable "long_term_backup_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "long_term_backup_id" {
  default = "id"
}

variable "protected_database_id" {}

variable "long_term_backup_retention_period_retention_count" {
  default = 1
}

variable "long_term_backup_retention_period_retention_period_type" {
  default = "YEAR"
}

variable "long_term_backup_retention_point_in_time" {
  default = null
}

variable "long_term_backup_retention_scn" {
  default = null 
}

variable "long_term_backup_state" {
  default = "ACTIVE"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_recovery_long_term_backup" "test_long_term_backup" {
  #Required
  protected_database_id = var.protected_database_id
  retention_period {
    #Required
    retention_count       = var.long_term_backup_retention_period_retention_count
    retention_period_type = var.long_term_backup_retention_period_retention_period_type
  }

  #Optional
  display_name            = var.long_term_backup_display_name
  freeform_tags           = var.long_term_backup_freeform_tags
  retention_point_in_time = var.long_term_backup_retention_point_in_time
  retention_scn           = var.long_term_backup_retention_scn
}

data "oci_recovery_long_term_backups" "test_long_term_backups" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name          = var.long_term_backup_display_name
  id                    = oci_recovery_long_term_backup.test_long_term_backup.id
  protected_database_id = var.protected_database_id
  state                 = var.long_term_backup_state
}

