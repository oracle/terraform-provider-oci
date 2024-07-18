// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "sensitive_data_model_id" {}
variable "sensitive_type_id" {}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_data_safe_sensitive_data_model_sensitive_types" "test_sensitive_data_model_sensitive_types" {
  #Required
  sensitive_data_model_id               = var.sensitive_data_model_id

  #Optional
  sensitive_type_id               = var.sensitive_type_id
}