// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "assessment_assessor_check_affected_object_assessor_name" {
  default = "assessorName"
}

variable "assessment_assessor_check_affected_object_check_name" {
  default = "checkName"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_database_migration_assessment_assessor_check_affected_objects" "test_assessment_assessor_check_affected_objects" {
  #Required
  assessment_id = oci_database_migration_assessment.test_assessment.id
  assessor_name = var.assessment_assessor_check_affected_object_assessor_name
  check_name    = var.assessment_assessor_check_affected_object_check_name
}

