// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "dr_protection_group_association_peer_region" {
  default = "us-ashburn-1"
}

variable "dr_protection_group_association_role" {
  default = "STANDBY"
}

variable "dr_protection_group_defined_tags_value" {
  default = "value"
}

variable "dr_protection_group_display_name" {
  default = "example-standby-drpg"
}

variable "dr_protection_group_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "dr_protection_group_log_location_bucket" {
  default = "bucket"
}

variable "dr_protection_group_log_location_namespace" {
  default = "namespace"
}

variable "dr_protection_group_members_is_movable" {
  default = false
}

variable "dr_protection_group_members_member_type" {
  default = "VOLUME_GROUP"
}

variable "dr_protection_group_state" {
  default = "ACTIVE"
}

data "oci_disaster_recovery_dr_protection_groups" "test_dr_protection_groups" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name           = var.dr_protection_group_display_name
  state                  = var.dr_protection_group_state
}