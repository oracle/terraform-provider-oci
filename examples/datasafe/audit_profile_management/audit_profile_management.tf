// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "compartment_id" {
  type    = string
  default = "<ocid>"
}

variable "target_id" {
  type    = string
  default = "<ocid>"
}
variable "target_type" {
  type    = string
  default = "TARGET_DATABASE_GROUP"
  #default = "TARGET_DATABASE"
}

variable "display_name" {
  type    = string
  default = "Audit_1"
}

variable "description" {
  type    = string
  default = "Description"
}

variable "freeform_tags" {
  type = map(string)
  default = {
    Department = "Finance"
  }
}

variable "is_override_global_paid_usage" {
  type    = bool
  default = true
}

variable "is_paid_usage_enabled" {
  type    = bool
  default = true
}

variable "offline_months" {
  type    = number
  default = 10
}

variable "online_months" {
  type    = number
  default = 10
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

/*
* In case of target group 
* Create Resource - Create audit profile
* Update resource - Updates the audit profile 
* Delete resource - Deletes audit profile
* ------------------------------------------
* In case of a target
* Create resource - Will fetch the profile details via target and compartment ocid either by a GET call or tracking Work request. No creation 
* Update resource - Updates audit profile 
* Destroy resource - Nothing is created so nothing will be destroyed
*/
resource "oci_data_safe_audit_profile_management" "test_audit_profile_management" {
  // Required
  compartment_id = var.compartment_id
  target_id                      = var.target_id
  target_type                    = var.target_type
  // Optional
  display_name                   = var.displayName
  description                    = var.description
  freeform_tags                  = var.freeform_tags
  is_override_global_paid_usage  = var.is_override_global_paid_usage
  is_paid_usage_enabled          = var.is_paid_usage_enabled
  offline_months                 = var.offline_months
  online_months                  = var.online_months
  change_retention_trigger       = 1
  is_override_global_retention_setting = true
}