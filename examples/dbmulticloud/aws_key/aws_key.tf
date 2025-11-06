// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

variable "oracle_db_aws_connector_id" {
  type = string
  default = ""
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_dbmulticloud_oracle_db_aws_key" "test_oracle_db_aws_key" {
  compartment_id         = var.compartment_ocid
  display_name           = "AWS_Key_Tersi_Test"
  oracle_db_connector_id = var.oracle_db_aws_connector_id
  aws_key_arn            = "arn:aws:iam::867344470629:role/OracleDatabaseKMS"
  is_aws_key_enabled     = false
}  

data "oci_dbmulticloud_oracle_db_aws_key" "test_oracle_db_aws_key" {
  oracle_db_aws_key_id = oci_dbmulticloud_oracle_db_aws_key.test_oracle_db_aws_key.id
}

output "aws_key_id" {
  value = oci_dbmulticloud_oracle_db_aws_key.test_oracle_db_aws_key.id
}
