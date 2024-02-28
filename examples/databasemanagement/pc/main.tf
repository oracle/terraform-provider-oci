// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

provider "oci" {
  tenancy_ocid = var.tenancy_ocid
  user_ocid = var.user_ocid
  fingerprint = var.fingerprint
  private_key_path = var.private_key_path
  region = var.region
}

variable "managed_database_id" {
   default = "<database.ocid>"
}

variable "managed_database_preferred_credential_credential_name" {
  default = "PC_WRITE"
}

data "oci_database_management_managed_database_preferred_credentials" "test_managed_database_preferred_credentials" {
  #Required
  managed_database_id = var.managed_database_id
}

data "oci_database_management_managed_database_preferred_credential" "test_managed_database_preferred_credential" {
  #Required
  credential_name = var.managed_database_preferred_credential_credential_name
  managed_database_id = var.managed_database_id
}