// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "generate_on_prem_connector_configuration_password" {
  default = "BEstrO0ng_#1111"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_data_safe_generate_on_prem_connector_configuration" "test_generate_on_prem_connector_configuration" {
  #Required
  on_prem_connector_id = oci_data_safe_on_prem_connector.test_on_prem_connector.id
  password             = var.generate_on_prem_connector_configuration_password
}