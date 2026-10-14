// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "assessment_assessor_action_assessor_action" {
  default = "assessorAction"
}

variable "assessment_assessor_action_assessor_name" {
  default = "assessorName"
}

variable "assessment_assessor_action_items_name" {
  default = "name"
}

variable "assessment_assessor_action_items_value" {
  default = "value"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_database_migration_assessment_assessor_action" "test_assessment_assessor_action" {
  #Required
  assessment_id   = oci_database_migration_assessment.test_assessment.id
  assessor_action = var.assessment_assessor_action_assessor_action
  assessor_name   = var.assessment_assessor_action_assessor_name
  items {
    #Required
    name  = var.assessment_assessor_action_items_name
    value = var.assessment_assessor_action_items_value
  }
}


