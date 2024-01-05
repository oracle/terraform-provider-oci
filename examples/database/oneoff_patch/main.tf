// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
}

variable "compartment_id" {
}

variable "oneoff_patch_db_version" {
  default = "19.0.0.0"
}

variable "oneoff_patch_defined_tags_value" {
  default = {}
}

variable "oneoff_patch_display_name" {
  default = "19.15_RU"
}

variable "oneoff_patch_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "oneoff_patch_one_off_patches" {
  default = ["31908573"]
}

variable "oneoff_patch_release_update" {
  default = "19.15.0.0"
}

variable "oneoff_patch_state" {
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_database_oneoff_patch" "test_oneoff_patch" {
  #Required
  compartment_id = var.compartment_id
  db_version     = var.oneoff_patch_db_version
  display_name   = var.oneoff_patch_display_name
  release_update = var.oneoff_patch_release_update

  #Optional
  defined_tags    = var.oneoff_patch_defined_tags_value
  freeform_tags   = var.oneoff_patch_freeform_tags
  one_off_patches = var.oneoff_patch_one_off_patches
}

data "oci_database_oneoff_patches" "test_oneoff_patches" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = var.oneoff_patch_display_name
  state        = var.oneoff_patch_state
}


