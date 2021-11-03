// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0
//
// We will create a database_tools_connection.
// The database_tools_connection will use a database_tools_private_endpoint.
// The database_tools_private_endpoint will need to be in a subnet, so we will create a vcn and a subnet.
// We will also create an oci_core_network_security_group to be used by the database_tools_private_endpoint.

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

### Network resources - vcn, subnet and network security group
# vcn
resource "oci_core_vcn" "tf_vcn" {
  cidr_block     = "10.0.3.0/24"
  compartment_id = var.compartment_ocid
  display_name   = "test databaseTools Vcn"
}

# subnet
resource "oci_core_subnet" "tf_subnet" {
  compartment_id = var.compartment_ocid
  vcn_id = oci_core_vcn.tf_vcn.id
  cidr_block = "10.0.3.0/26"
  display_name = "test databaseTools Subnet"
}

# network security group
resource "oci_core_network_security_group" "test_network_security_group" {
  #Required
  compartment_id = var.compartment_ocid
  vcn_id = oci_core_vcn.tf_vcn.id

  #Optional
  display_name  = "nsg1"
  freeform_tags = {"Department"= "Finance"}
}

### Endpoints services
# Endpoints services - Data Sources
data "oci_database_tools_database_tools_endpoint_services" "test_database_tools_endpoint_services" {
  compartment_id = var.compartment_ocid
  state = "ACTIVE"
}

data "oci_database_tools_database_tools_endpoint_service" "test_database_tools_endpoint_service" {
  database_tools_endpoint_service_id = data.oci_database_tools_database_tools_endpoint_services.test_database_tools_endpoint_services.database_tools_endpoint_service_collection.0.items.0.id
}

output "endpoint_service" {
  value = data.oci_database_tools_database_tools_endpoint_service.test_database_tools_endpoint_service
}

### Private Endpoint
# Private Endpoint - Resource
resource "oci_database_tools_database_tools_private_endpoint" "test_database_tools_private_endpoint" {
  #Required
  compartment_id      = var.compartment_ocid
  display_name        = "My private endpoint"
  endpoint_service_id = data.oci_database_tools_database_tools_endpoint_service.test_database_tools_endpoint_service.id
  subnet_id           = oci_core_subnet.tf_subnet.id

  #Optional
  description         = "Private Endpoint used by connection"
  nsg_ids             = [oci_core_network_security_group.test_network_security_group.id]
  private_endpoint_ip = "10.0.3.4"
}

# Private Endpoint - Data Sources
data "oci_database_tools_database_tools_private_endpoints" "test_database_tools_private_endpoints" {
  compartment_id  = var.compartment_ocid
  state           = "ACTIVE"
  subnet_id       = oci_core_subnet.tf_subnet.id
  display_name    = oci_database_tools_database_tools_private_endpoint.test_database_tools_private_endpoint.display_name
}

output "private_endpoints_d" {
  value = data.oci_database_tools_database_tools_private_endpoints.test_database_tools_private_endpoints
}

data "oci_database_tools_database_tools_private_endpoint" "test_database_tools_private_endpoint" {
  database_tools_private_endpoint_id = data.oci_database_tools_database_tools_private_endpoints.test_database_tools_private_endpoints.database_tools_private_endpoint_collection.0.items.0.id
}

output "private_endpoint_d" {
  value = data.oci_database_tools_database_tools_private_endpoint.test_database_tools_private_endpoint
}

### Connection
# Connection - Resource
resource "oci_database_tools_database_tools_connection" "dbtools_connection" {
  compartment_id    = var.compartment_ocid
  display_name      = "My Connection"
  type              = "ORACLE_DATABASE"
  connection_string = "tcps://adb-prod.us-phoenix-1.oraclecloud.com:1522/exampleb2baffff_db20210323ffff_low.adb.oraclecloud.com"
  user_name         = "john.doe@oracle.com"
  user_password {
    value_type = "SECRETID"

    # Here, we assume that the user password to use exists as a secret in an OCI Vault
    secret_id  = "ocid1.vaultsecret.oc1.phx.exampleaihuofciaiazy2u5ko3uyz3sspwd6hf7oqhqmlk5xu3xdetkpffff"
  }

  # Optional
  freeform_tags = { my-Freeform-tag1 = "value f1", my-Freeform-tag2 = "value f2"}
  advanced_properties = {
    "oracle.jdbc.loginTimeout": "0"
  }
  related_resource {
    entity_type = "DATABASE"
    identifier  = "ocid1.database.oc1.phx.exampletksujfufl4bhe5sqkfgn7t7lcrkkpy7km5iwzvg6ycls7r5dlffff"
  }
  private_endpoint_id = oci_database_tools_database_tools_private_endpoint.test_database_tools_private_endpoint.id
}

output "connection_r" {
  value = oci_database_tools_database_tools_connection.dbtools_connection
}

# Connection - Data Sources
data "oci_database_tools_database_tools_connections" "test_database_tools_connections" {
  compartment_id = var.compartment_ocid
  display_name   = oci_database_tools_database_tools_connection.dbtools_connection.display_name
  state          = "ACTIVE"
}

output "connections_d" {
  value = data.oci_database_tools_database_tools_connections.test_database_tools_connections
}