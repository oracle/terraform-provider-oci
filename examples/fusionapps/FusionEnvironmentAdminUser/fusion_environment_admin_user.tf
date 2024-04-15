// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "fusion_environment_admin_user_email_address" {
  default = "JohnSmithnew@example.com"
}

variable "fusion_environment_admin_user_first_name" {
  default = "firstName"
}

variable "fusion_environment_admin_user_last_name" {
  default = "lastName"
}

variable "fusion_environment_admin_user_password" {
  default = "BEstrO0ng_#11"
}

variable "fusion_environment_admin_user_username" {
  default = "username_test"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_fusion_apps_fusion_environment_admin_user" "test_fusion_environment_admin_user" {
  #Required
  email_address         = var.fusion_environment_admin_user_email_address
  first_name            = var.fusion_environment_admin_user_first_name
  fusion_environment_id = oci_fusion_apps_fusion_environment.test_fusion_environment.id
  last_name             = var.fusion_environment_admin_user_last_name 
  username              = var.fusion_environment_admin_user_username
  #Optional
  password              = var.fusion_environment_admin_user_password
}

data "oci_fusion_apps_fusion_environment_admin_users" "test_fusion_environment_admin_users" {
  #Required
  fusion_environment_id = oci_fusion_apps_fusion_environment.test_fusion_environment.id
}
