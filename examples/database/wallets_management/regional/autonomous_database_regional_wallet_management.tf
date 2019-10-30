// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

provider "oci" {
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region           = "${var.region}"
}

resource "oci_database_autonomous_database_regional_wallet_management" "test_autonomous_database_regional_wallet_management" {
  should_rotate = false
}

data "oci_database_autonomous_database_regional_wallet_management" "test_autonomous_database_regional_wallet_management" {}
