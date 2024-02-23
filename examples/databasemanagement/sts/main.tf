// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "compartment_id" {  default = "<compartment.ocid>"}

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

variable "managed_database_sql_tuning_set_name_contains" {
  default = "SYS"
}

variable "managed_database_sql_tuning_set_owner" {
  default = "SYS"
}

data "oci_database_management_managed_database_sql_tuning_sets" "test_managed_database_sql_tuning_sets" {
	#Required
	managed_database_id = var.managed_database_id

	#Optional
	name_contains = var.managed_database_sql_tuning_set_name_contains
	owner = var.managed_database_sql_tuning_set_owner
}