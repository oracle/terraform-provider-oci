// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

variable "oracle_db_azure_connector_id" {
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

resource "oci_dbmulticloud_multi_cloud_resource_discovery" "test_multi_cloud_resource_discovery" {
compartment_id = var.compartment_ocid
display_name = "Tersi_Discover_Test"
oracle_db_connector_id = var.oracle_db_azure_connector_id
resource_type = "VAULTS"
}

data "oci_dbmulticloud_multi_cloud_resource_discovery" "test_multi_cloud_resource_discovery" {
  multi_cloud_resource_discovery_id = oci_dbmulticloud_multi_cloud_resource_discovery.test_multi_cloud_resource_discovery.id
}

data "oci_dbmulticloud_multi_cloud_resource_discoveries" "test_multi_cloud_resource_discovery" {
  compartment_id = var.compartment_ocid
  resource_type = "VAULTS"
}
