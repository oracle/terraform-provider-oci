// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_dbmulticloud_oracle_db_aws_key_pool" "test_oracle_db_aws_key_pool" {
  compartment_id           = var.compartment_ocid
  display_name             = "TestDbAwsKeyPoolFromExample"
  is_multi_region_key_pool = false
  is_reuse_allowed         = false
}

data "oci_dbmulticloud_oracle_db_aws_key_pool" "test_oracle_db_aws_key_pool" {
  oracle_db_aws_key_pool_id = oci_dbmulticloud_oracle_db_aws_key_pool.test_oracle_db_aws_key_pool.id
}

output "aws_key_pool_id" {
  value = oci_dbmulticloud_oracle_db_aws_key_pool.test_oracle_db_aws_key_pool.id
}
