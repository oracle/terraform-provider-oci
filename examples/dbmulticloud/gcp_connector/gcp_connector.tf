// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

# variable "tenancy_ocid" {}
# variable "user_ocid" {}
# variable "fingerprint" {}
# variable "private_key_path" {}
# variable "region" {}
variable "compartment_id" {}


variable "oracle_db_gcp_identity_connector_display_name" {
  default = "Tersi example"
}

variable "oracle_db_gcp_identity_connector_gcp_location" {
  default = "global"
}

variable "oracle_db_gcp_identity_connector_issuer_url" {
  default = "https://idcs-test.identity.oraclecloud.com"
}

variable "oracle_db_gcp_resource_service_agent_id" {
  default = "test agent"
}

variable "oracle_db_gcp_workload_identity_pool_id" {
  default = "dbmci"
}

variable "oracle_db_gcp_workload_identity_provider_id" {
  default = "test provider"
}

variable "project_id" {
  default = "test project"
}

variable "resource_id" {
  default = "ocid1.cloudvmcluster.test..tersitestexamplegcp"
}

# provider "oci" {
#   tenancy_ocid     = var.tenancy_ocid
#   user_ocid        = var.user_ocid
#   fingerprint      = var.fingerprint
#   private_key_path = var.private_key_path
#   region           = var.region
# }

resource "oci_dbmulticloud_oracle_db_gcp_identity_connector" "test_oracle_db_gcp_identity_connector" {
  #Required
  compartment_id                    = var.compartment_id
  display_name                      = var.oracle_db_gcp_identity_connector_display_name
  gcp_location                      = var.oracle_db_gcp_identity_connector_gcp_location
  gcp_resource_service_agent_id     = var.oracle_db_gcp_resource_service_agent_id
  gcp_workload_identity_pool_id     = var.oracle_db_gcp_workload_identity_pool_id
  gcp_workload_identity_provider_id = var.oracle_db_gcp_workload_identity_provider_id
  issuer_url                        = var.oracle_db_gcp_identity_connector_issuer_url
  project_id                        = var.project_id
  resource_id                       = var.resource_id
}

data "oci_dbmulticloud_oracle_db_gcp_identity_connectors" "test_oracle_db_gcp_identity_connectors" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = var.oracle_db_gcp_identity_connector_display_name
  resource_id  = var.resource_id
}


output "gcp_connector_id" {
  value = oci_dbmulticloud_oracle_db_gcp_identity_connector.test_oracle_db_gcp_identity_connector.id
}
