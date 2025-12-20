// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "source_connection_oracle_id" {
  default = ""
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  auth             = "SecurityToken"
  region           = var.region

}

data "oci_database_migration_connection_databaseconnectiontypes" "test_connection_databaseconnectiontypes" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  connection_type      = ["ORACLE"]
  source_connection_id = var.source_connection_oracle_id
  technology_type      = ["OCI_AUTONOMOUS_DATABASE"]
}
