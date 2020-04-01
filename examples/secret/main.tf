// Copyright (c) 2017, 2019, 2020, Oracle and/or its affiliates. All rights reserved.

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "secret_id" {}
variable "vault_id" {}

provider "oci" {
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region           = "${var.region}"
}

data "oci_vault_secret" "test_secret" {
  secret_id = "${var.secret_id}"
}

data "oci_vault_secret" "test_secret" {
  compartment_id = "${var.compartment_ocid}"
  state          = "Active"
  vault_id       = "${var.vault_id}"
}
