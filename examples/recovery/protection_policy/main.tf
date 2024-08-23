// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

variable "protection_policy_backup_retention_period_in_days" {
  default = 14
}

variable "protection_policy_display_name" {
  default = "displayName"
}

variable "protection_policy_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "protection_policy_owner" {
  default = "customer"
}

variable "protection_policy_state" {
  default = "ACTIVE"
}

variable "protection_policy_policy_locked_date_time" {
  default = "2025-01-01T00:00:00.000Z"
}

variable "must_enforce_cloud_locality" {
  default = false
}

resource "oci_recovery_protection_policy" "test_protection_policy" {
  #Required
  backup_retention_period_in_days = var.protection_policy_backup_retention_period_in_days
  compartment_id                  = var.compartment_id
  display_name                    = var.protection_policy_display_name

  #Optional
  policy_locked_date_time         = var.protection_policy_policy_locked_date_time
  must_enforce_cloud_locality     = var.must_enforce_cloud_locality
  freeform_tags = var.protection_policy_freeform_tags
}

data "oci_recovery_protection_policies" "test_protection_policies" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name         = var.protection_policy_display_name
  owner                = var.protection_policy_owner
  protection_policy_id = oci_recovery_protection_policy.test_protection_policy.id
  state                = var.protection_policy_state
}