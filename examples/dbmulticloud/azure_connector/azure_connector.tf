// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

variable "db_cluster_resource_id" {
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

resource "oci_dbmulticloud_oracle_db_azure_connector" "test_oracle_db_azure_connector" {
  access_token = "AzureAccessToken"
  azure_identity_mechanism = "ARC_AGENT"
  azure_resource_group = "Prasanna.RG"
  azure_subscription_id = "7080446f-ee76-4aa2-b9dd-c2625f63cab0"
  azure_tenant_id = "5b743bc7-c1e2-4d46-b4b5-a32eddac0286"
  compartment_id  = var.compartment_ocid
  db_cluster_resource_id = var.db_cluster_resource_id
  display_name = "AzureConnectorTest-Tersi"
}

data "oci_dbmulticloud_oracle_db_azure_connector" "test_oracle_db_azure_connector" {
  oracle_db_azure_connector_id = oci_dbmulticloud_oracle_db_azure_connector.test_oracle_db_azure_connector.id
}

output "azure_connector_id" {
  value = oci_dbmulticloud_oracle_db_azure_connector.test_oracle_db_azure_connector.id
}
