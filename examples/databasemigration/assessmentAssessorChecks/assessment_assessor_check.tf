// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "assessment_assessor_check_assessor_name" {
  default = "assessorName"
}

variable "assessment_assessor_check_display_name" {
  default = "displayName"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_database_migration_assessment_assessor_checks" "test_assessment_assessor_checks" {
  #Required
  assessment_id  = oci_database_migration_assessment.test_assessment.id
  assessor_name  = var.assessment_assessor_check_assessor_name
  compartment_id = var.compartment_id

  #Optional
  display_name = var.assessment_assessor_check_display_name
}

