// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0
//
// We will create a Postgresql database_tools_connection

variable "secret_ocid" {
}

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "compartment_ocid" {
}

variable "region" {
}

provider "oci" {
  region           = var.region
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
}

### Connection
# Connection - Resource
resource "oci_database_tools_database_tools_connection" "dbtools_connection_postgresql" {
  compartment_id    = var.compartment_ocid
  display_name      = "My Postgresql connection"
  type              = "POSTGRESQL"
  connection_string = "postgresql://example.com:5432/db"
  user_name         = "john.doe@oracle.com"
  runtime_support   = "UNSUPPORTED"
  user_password {
    value_type = "SECRETID"
    # The user password to use mus exist as a secret in an OCI Vault
    secret_id  = var.secret_ocid
  }

  # Optional
  freeform_tags = {
    my-Freeform-tag1 = "value f1",
    my-Freeform-tag2 = "value f2"
  }

  key_stores {
    key_store_content {
      value_type = "SECRETID"
      secret_id = var.secret_ocid
    }
    key_store_type = "CA_CERTIFICATE_PEM"
  }
}

output "connection_r_postgresql" {
  value = oci_database_tools_database_tools_connection.dbtools_connection_postgresql
}

# Connection - Data Sources
data "oci_database_tools_database_tools_connections" "test_database_tools_connections_postgresql" {
  compartment_id  = var.compartment_ocid
  display_name    = oci_database_tools_database_tools_connection.dbtools_connection_postgresql.display_name
  state           = "ACTIVE"
  runtime_support = ["UNSUPPORTED"]
  type            = ["POSTGRESQL"]
}

output "connections_d_postgresql" {
  value = data.oci_database_tools_database_tools_connections.test_database_tools_connections_postgresql
}