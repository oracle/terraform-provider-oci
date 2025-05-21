// Copyright (c) 2017, 2025, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

// This example illustrates how to update a NoSQL configuration. Conceptually, a
// configuration serves as a centralized repository for global parameters that
// affect the NoSQL service. Currently, there is only one such parameter: a
// customer-provided key for encrypting NoSQL data at rest.

// The Customer-Managed Encryption Keys (CMEK) feature is exclusively available
// in private NoSQL environments dedicated to a single tenancy, where the CMEK
// option has been enabled. Updating the configuration of the default, regional,
// multi-tenancy NoSQL service is not supported. Therefore, this example focuses
// on updating the configuration for a private NoSQL environment that supports
// this operation.

// To specify the dedicated environment, set the environment variable
// CLIENT_HOST_OVERRIDES=oci_nosql.NosqlClient=$ENDPOINT
// Where $ENDPOINT is the endpoint of the dedicated NoSQL environment.
// For example:
// $ export CLIENT_HOST_OVERRIDES=oci_nosql.NosqlClient=https://acme-widgets.nosql.oci.oraclecloud.com

// The key must be stored in an OCI vault, and is referenced by its vault and
// key OCIDs. Note that the configuration cannot be deleted. To revert to the
// default encryption key, set the TF_VAR_vault_ocid and TF_VAR_key_ocid
// environment variables to be empty, or assign the literal value of null
// to kms_key.id and kms_key.kms_vault_id.

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "key_ocid" {}
variable "vault_ocid" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_nosql_configuration" "example_configuration" {
  compartment_id = var.tenancy_ocid
  environment = "HOSTED"
  kms_key {
    kms_vault_id = var.vault_ocid
    id = var.key_ocid
  }
}

data "oci_nosql_configuration" "test_configurations" {
  compartment_id = var.tenancy_ocid
}
