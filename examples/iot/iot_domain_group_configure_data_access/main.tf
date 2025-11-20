// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "iot_domain_group_ocid" {}

variable "iot_domain_group_configure_data_access_db_allow_listed_vcn_ids" {
  default = []
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_iot_iot_domain_group_configure_data_access" "test_iot_domain_group_configure_data_access" {
  #Required
  db_allow_listed_vcn_ids = var.iot_domain_group_configure_data_access_db_allow_listed_vcn_ids
  iot_domain_group_id     = var.iot_domain_group_ocid
}


