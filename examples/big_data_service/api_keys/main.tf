// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "bds_instance_id" {}

variable "bds_instance_api_key_default_region" {
  default = "us-ashburn-1"
}

variable "bds_instance_api_key_display_name" {
  default = "keyAlias"
}

variable "bds_instance_api_key_key_alias" {
  default = "keyAlias"
}

variable "bds_instance_api_key_passphrase" {
  default = "V2VsY29tZTE="
}

variable "bds_instance_api_key_state" {
  default = "ACTIVE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_bds_bds_instance_api_key" "test_bds_instance_api_key" {
  #Required
  bds_instance_id = var.bds_instance_id
  key_alias       = var.bds_instance_api_key_key_alias
  passphrase      = var.bds_instance_api_key_passphrase
  user_id         = var.user_ocid

  #Optional
  default_region = var.bds_instance_api_key_default_region
}

data "oci_bds_bds_instance_api_keys" "test_bds_instance_api_keys" {
  #Required
  bds_instance_id = var.bds_instance_id

  #Optional
  state        = var.bds_instance_api_key_state
}
