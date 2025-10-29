// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0
//
// We will create a Generic JDBC database_tools_connection

### Connection
# Connection - Resource
resource "oci_database_tools_database_tools_connection" "dbtools_connection_generic_jdbc" {
  compartment_id    = var.compartment_ocid
  display_name      = "My Generic JDBC connection"
  type              = "GENERIC_JDBC"
  url               = "jdbc:mysql://example.com:3306/db"
  user_name         = "john.doe@oracle.com"
  runtime_support   = "UNSUPPORTED"
  user_password {
    value_type = "SECRETID"
    # The user password to use must exist as a secret in an OCI Vault
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
      secret_id = var.ca_certificate_pem_secret_ocid
    }
    key_store_type = "CA_CERTIFICATE_PEM"
  }
}

output "connection_r_generic_jdbc" {
  value = oci_database_tools_database_tools_connection.dbtools_connection_generic_jdbc
}

# Connection - Data Sources
data "oci_database_tools_database_tools_connections" "test_database_tools_connections_generic_jdbc" {
  compartment_id  = var.compartment_ocid
  display_name    = oci_database_tools_database_tools_connection.dbtools_connection_generic_jdbc.display_name
  state           = "ACTIVE"
  runtime_support = ["UNSUPPORTED"]
  type            = ["GENERIC_JDBC"]
}

output "connections_d_generic_jdbc" {
  value = data.oci_database_tools_database_tools_connections.test_database_tools_connections_generic_jdbc
}