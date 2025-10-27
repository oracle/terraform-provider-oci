// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0
//
// We will create an oracle_database_resource_principal identity for a database tools connection to an Oracle database.
// The database_tools_connection will use a database_tools_private_endpoint.
// The database_tools_private_endpoint will need to be in a subnet, so we will create a vcn and a subnet.
// We will also create an oci_core_network_security_group to be used by the database_tools_private_endpoint.

resource "oci_database_tools_database_tools_connection" "example_connection" {
  compartment_id = var.compartment_ocid
  display_name = "Test Connection"
  type = "ORACLE_DATABASE"
  connection_string = var.connection_string
  user_name = "admin"
  # Identity resources can only be created against connections using RESOURCE_PRINCIPAL as the runtime_identity
  runtime_identity = "RESOURCE_PRINCIPAL"
  user_password {
    secret_id  = var.user_password_secret_ocid
    value_type = "SECRETID"
  }
  key_stores {
    key_store_type = "SSO"
    key_store_content {
      value_type = "SECRETID"
      secret_id = var.database_wallet_secret_ocid
    }
  }
}

resource "oci_database_tools_database_tools_identity" "example_identity" {
  type = "ORACLE_DATABASE_RESOURCE_PRINCIPAL"
  display_name = "My Identity"
  database_tools_connection_id = oci_database_tools_database_tools_connection.example_connection.id
  compartment_id = var.compartment_ocid
  credential_key = "MY_TEST_CREDENTIAL"
}