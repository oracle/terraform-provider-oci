// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "protection_policy_backup_retention_period_in_days" {
  default = 10
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


resource "oci_recovery_protection_policy" "test_protection_policy" {
  #Required
  backup_retention_period_in_days = var.protection_policy_backup_retention_period_in_days
  compartment_id                  = var.compartment_id
  display_name                    = var.protection_policy_display_name

  #Optional
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
