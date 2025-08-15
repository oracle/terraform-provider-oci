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

resource "oci_dbmulticloud_oracle_db_azure_vault" "test_oracle_db_azure_vault" {
azure_vault_id = "PrasannaHSM2"
compartment_id = var.compartment_ocid
display_name = "Discover_Tersi_Test"
location = "eastus2"
oracle_db_azure_resource_group = "Prasanna.RG"
oracle_db_connector_id = var.oracle_db_azure_connector_id
type = "managedHSMs"
}

output "azure_vault_id" {
  value = oci_dbmulticloud_oracle_db_azure_vault.test_oracle_db_azure_vault.id
}
