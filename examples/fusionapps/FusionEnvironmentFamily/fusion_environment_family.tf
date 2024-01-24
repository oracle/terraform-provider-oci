// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "fusion_environment_family_defined_tags_value" {
  default = "value"
}

variable "fusion_environment_family_display_name" {
  default = "displayName"
}

variable "fusion_environment_family_family_maintenance_policy_concurrent_maintenance" {
  default = "PROD"
}

variable "fusion_environment_family_family_maintenance_policy_is_monthly_patching_enabled" {
  default = false
}

variable "fusion_environment_family_family_maintenance_policy_quarterly_upgrade_begin_times" {
  default = "RRULE:FREQ=YEARLY;BYMONTH=2,5,8,11"
}

variable "fusion_environment_family_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "fusion_environment_family_state" {
  default = "ACTIVE"
}

variable "fusion_environment_family_subscription_ids" {
  default = []
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_fusion_apps_fusion_environment_family" "test_fusion_environment_family" {
  #Required
  compartment_id   = var.compartment_id
  display_name     = var.fusion_environment_family_display_name
  subscription_ids = var.fusion_environment_family_subscription_ids

  #Optional
  defined_tags = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.fusion_environment_family_defined_tags_value)
  family_maintenance_policy {

    #Optional
    concurrent_maintenance        = var.fusion_environment_family_family_maintenance_policy_concurrent_maintenance
    is_monthly_patching_enabled   = var.fusion_environment_family_family_maintenance_policy_is_monthly_patching_enabled
    quarterly_upgrade_begin_times = var.fusion_environment_family_family_maintenance_policy_quarterly_upgrade_begin_times
  }
  freeform_tags = var.fusion_environment_family_freeform_tags
}

data "oci_fusion_apps_fusion_environment_families" "test_fusion_environment_families" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name                 = var.fusion_environment_family_display_name
  fusion_environment_family_id = oci_fusion_apps_fusion_environment_family.test_fusion_environment_family.id
  state                        = var.fusion_environment_family_state
}
