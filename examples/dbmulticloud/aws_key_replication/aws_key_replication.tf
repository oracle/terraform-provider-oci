// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

// Required to match the existing key
variable "oracle_db_aws_connector_id" {
  type = string
}

variable "oracle_db_aws_key_id" {
  type = string
}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}


resource "oci_dbmulticloud_oracle_db_aws_key" "existing" {
  compartment_id         = var.compartment_ocid
  display_name           = "MockResourceName"
  oracle_db_connector_id = var.oracle_db_aws_connector_id
  aws_key_arn            = "arn:aws:kms:us-east-2:895395310091:key/mrk-276f7f63c95046e88797b19d8949f814"
  is_aws_key_enabled     = false
  // Replication
  action        = "DELETE"
  target_region = "us-boardman-1"
}


data "oci_dbmulticloud_oracle_db_aws_key" "existing" {
  oracle_db_aws_key_id = var.oracle_db_aws_key_id
}

output "aws_key_id" {
  value = oci_dbmulticloud_oracle_db_aws_key.existing.id
}

