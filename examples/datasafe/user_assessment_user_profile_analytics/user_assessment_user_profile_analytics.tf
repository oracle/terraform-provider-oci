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


variable "description" {
  default = "description"
}

variable "display_name" {
  default = "UA_1"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_data_safe_user_assessment" "test_user_assessment" {
  #Required
  compartment_id  = var.compartment_id
  target_id       = var.target_id

  #Optional
  description  = var.description
  display_name = var.display_name
}

data "oci_data_safe_user_assessment_profile_analytics" "test_user_assessment_profile_analytics" {
  #Required
  user_assessment_id  = var.user_assessment_id
  compartment_id      = var.compartment_id

}