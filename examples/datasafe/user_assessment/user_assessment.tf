// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0


variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}
variable "user_assessment_id" {}
variable "target_id" {}

variable "display_name" {
  default = "UA_1"
}

variable "description" {
  default = "description"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_data_safe_user_assessment" "oci_data_safe_user_assessment" {
  #Required
  compartment_id = var.compartment_id
  target_id = var.target_id

  #Optional
  display_name = var.display_name
  description   = var.description
}

data "oci_data_safe_user_assessments" "test_user_assessments" {
  #Required
  compartment_id = var.compartment_id
  target_id = var.target_id

}