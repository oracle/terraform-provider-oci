// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0
//
// We will create a database_tools_connection for a MySQL database
// The database_tools_connection will not a database_tools_private_endpoint.

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

variable "secret_ocid" {
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
resource "oci_database_tools_database_tools_connection" "dbtools_connection_mysql" {
  compartment_id    = var.compartment_ocid
  display_name      = "My Connection MySQL"
  type              = "MYSQL"
  connection_string = "mysql://example.com:3306/db"
  user_name         = "john.doe@oracle.com"
  user_password {
    value_type = "SECRETID"
    # The user password to use exists as a secret in an OCI Vault
    secret_id  = var.secret_ocid
  }

  # Optional
  freeform_tags = { my-Freeform-tag1 = "value f1", my-Freeform-tag2 = "value f2"}
  advanced_properties = {
    "sslMode": "VERIFY_CA"
  }
  key_stores {
    key_store_type = "CLIENT_CERTIFICATE_PEM"
    key_store_content {
      value_type = "SECRETID"
      secret_id = var.secret_ocid
    }
  }
  key_stores {
    key_store_type = "CLIENT_PRIVATE_KEY_PEM"
    key_store_content {
      value_type = "SECRETID"
      secret_id = var.secret_ocid
    }
    key_store_password {
      value_type = "SECRETID"
      secret_id = var.secret_ocid
    }
  }
  key_stores {
    key_store_type = "CA_CERTIFICATE_PEM"
    key_store_content {
      value_type = "SECRETID"
      secret_id = var.secret_ocid
    }
  }

  related_resource {
    entity_type = "MYSQLDBSYSTEM"
    identifier  = "ocid1.mysqldbsystem.oc1.phx.exampletksujfufl4bhe5sqkfgn7t7lcrkkpy7km5iwzvg6ycls7r5dlfff1"
  }
}

output "connection_r_mysql" {
  value = oci_database_tools_database_tools_connection.dbtools_connection_mysql
}

# Connection - Data Sources
data "oci_database_tools_database_tools_connections" "test_database_tools_connections_mysql" {
  compartment_id = var.compartment_ocid
  display_name   = oci_database_tools_database_tools_connection.dbtools_connection_mysql.display_name
  state          = "ACTIVE"
}

output "connections_d_mysql" {
  value = data.oci_database_tools_database_tools_connections.test_database_tools_connections_mysql
}