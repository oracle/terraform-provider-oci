// Copyright (c) 2017, 2025, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "apm_domain_id" {}

variable "attribute_auto_activate_status_data_key_type" {
  default = "PRIVATE_DATA_KEY"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_apm_traces_attribute_auto_activate_status" "test_attribute_auto_activate_status" {
  #Required
  apm_domain_id = var.apm_domain_id
  data_key_type = var.attribute_auto_activate_status_data_key_type
}
