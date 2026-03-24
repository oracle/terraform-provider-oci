// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

// Required to match the existing key
variable "oracle_db_azure_connector_id" {
  type = string
  default = ""
}

variable "oracle_db_azure_vault_id" {
  type = string
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_dbmulticloud_oracle_db_azure_vault" "existing" {
azure_vault_id = "PrasannaHSM2"
compartment_id = var.compartment_ocid
display_name = "Tersi_Example_Vault"
location = "eastus2"
oracle_db_azure_resource_group = "Prasanna.RG"
oracle_db_connector_id = var.oracle_db_azure_connector_id
type = "managedHSMs"
  // Replication
  action        = "DELETE"
  target_region = "us-boardman-1"
}

data "oci_dbmulticloud_oracle_db_azure_vault" "existing" {
  oracle_db_azure_vault_id =var.oracle_db_azure_vault_id
}

output "azure_vault_id" {
  value = oci_dbmulticloud_oracle_db_azure_vault.existing.id
}

