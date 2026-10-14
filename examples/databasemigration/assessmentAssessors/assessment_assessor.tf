// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "assessment_assessor_display_name" {
  default = "displayName"
}

variable "assessment_assessor_state" {
  default = "AVAILABLE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_database_migration_assessment_assessors" "test_assessment_assessors" {
  #Required
  assessment_id = oci_database_migration_assessment.test_assessment.id

  #Optional
  display_name = var.assessment_assessor_display_name
  state        = var.assessment_assessor_state
}

