// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "data_safe_user_assessment_id" {}

variable "ua_password_expiry_analytic_compartment_id_in_subtree" {
  default = true
}

variable "ua_password_expiry_analytic_access_level" {
  default = "RESTRICTED"
}

variable "ua_password_expiry_analytic_time_less_than"{
  default = "2038-01-01T00:00:00.000Z"
}

variable "ua_password_expiry_analytic_user_category" {
  default = "HIGH"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_data_safe_user_assessment_password_expiry_date_analytics" "test_user_assessment_password_expiry_date_analytics" {
  #Required
  user_assessment_id = var.data_safe_user_assessment_id

  #Optional
  compartment_id_in_subtree = var.ua_password_expiry_analytic_compartment_id_in_subtree
  access_level = var.ua_password_expiry_analytic_access_level
  time_password_expiry_less_than = var.ua_password_expiry_analytic_time_less_than
  user_category = var.ua_password_expiry_analytic_user_category
}