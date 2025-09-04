// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

provider "oci" {
  version          = "7.13.0"
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}


resource "oci_dbmulticloud_oracle_db_azure_blob_container" "test_oracle_db_azure_blob_container" {
azure_storage_account_name = "ADBDAppStorageAccounts"
azure_storage_container_name = "ADBDContainers"
compartment_id =  var.compartment_ocid
display_name = "TestDBAzureBlobContainerUpdate"
}
