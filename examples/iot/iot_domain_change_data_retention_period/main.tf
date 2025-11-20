// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "iot_domain_ocid" {}

variable "iot_domain_change_data_retention_period_data_retention_period_in_days" {
  default = 90
}

variable "iot_domain_change_data_retention_period_type" {
  default = "RAW_DATA"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_iot_iot_domain_change_data_retention_period" "test_iot_domain_change_data_retention_period" {
  #Required
  data_retention_period_in_days = var.iot_domain_change_data_retention_period_data_retention_period_in_days
  iot_domain_id                 = var.iot_domain_ocid
  type                          = var.iot_domain_change_data_retention_period_type
}


