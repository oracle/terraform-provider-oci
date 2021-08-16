// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "managed_database_group_description" {
  default = "Sales test database group"
}

variable "managed_database_group_id" {
  default = "id"
}

variable "managed_database_group_name" {
  default = "TestGroup"
}

variable "managed_database_group_state" {
  default = "ACTIVE"
}

variable "managed_database_id" {
  default = "testManagedDatabase0"
}

variable "managed_databases_database_parameter_credentials_username" {
  default = "sys"
}

variable "managed_databases_database_parameter_credentials_password" {
  default = "sys"
}

variable "managed_databases_database_parameter_credentials_role" {
  default = "NORMAL"
}

variable "managed_databases_database_parameter_parameters_name" {
  default = "open_cursors"
}

variable "managed_databases_database_parameter_parameters_value" {
  default = "305"
}

variable "managed_databases_database_parameter_update_comment" {
  default = "Terraform update of open cursors"
}

variable "managed_databases_database_parameter_scope" {
  default = "BOTH"
}

variable "managed_databases_database_parameter_is_allowed_values_included" {
  default = "false"
}

variable "managed_databases_database_parameter_source" {
  default = "CURRENT"
}

variable "db_management_private_endpoint_name" {
  default = "TestPrivateEndpoint"
}

variable "db_management_private_endpoint_description" {
  default = "Test private endpoint"
}

variable "db_management_private_endpoint_state" {
  default = "ACTIVE"
}

provider "oci" {
  tenancy_ocid = var.tenancy_ocid
  user_ocid = var.user_ocid
  fingerprint = var.fingerprint
  private_key_path = var.private_key_path
  region = var.region
}

resource "oci_database_management_managed_database_group" "test_managed_database_group" {
  #Required
  compartment_id = var.compartment_id
  name = var.managed_database_group_name

  #Optional
  description = var.managed_database_group_description
  managed_databases {
    id = var.managed_database_id
  }
}

data "oci_database_management_managed_database_groups" "test_managed_database_groups_with_id" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  id = oci_database_management_managed_database_group.test_managed_database_group.id
  state = var.managed_database_group_state
}

data "oci_database_management_managed_database_groups" "test_managed_database_groups_with_name" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  name = var.managed_database_group_name
  state = var.managed_database_group_state
}

resource "oci_database_management_managed_databases_change_database_parameter" "test_managed_databases_change_database_parameter" {
  #Required
  credentials {

    #Optional
    password = var.managed_databases_database_parameter_credentials_password
    role = var.managed_databases_database_parameter_credentials_role
    user_name = var.managed_databases_database_parameter_credentials_username
  }
  managed_database_id = var.managed_database_id
  parameters {
    #Required
    name = var.managed_databases_database_parameter_parameters_name
    value = var.managed_databases_database_parameter_parameters_value

    #Optional
    update_comment = var.managed_databases_database_parameter_update_comment
  }
  scope = var.managed_databases_database_parameter_scope
}

resource "oci_database_management_managed_databases_reset_database_parameter" "test_managed_databases_reset_database_parameter" {
  #Required
  credentials {

    #Optional
    password = var.managed_databases_database_parameter_credentials_password
    role = var.managed_databases_database_parameter_credentials_role
    user_name = var.managed_databases_database_parameter_credentials_username
  }
  managed_database_id = var.managed_database_id
  parameters = [var.managed_databases_database_parameter_parameters_name]
  scope = var.managed_databases_database_parameter_scope
}

data "oci_database_management_managed_databases_database_parameter" "test_managed_databases_database_parameter" {
  #Required
  managed_database_id = var.managed_database_id

  #Optional
  is_allowed_values_included = var.managed_databases_database_parameter_is_allowed_values_included
  name = var.managed_databases_database_parameter_parameters_name
  source = var.managed_databases_database_parameter_source
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_id
  display_name   = "exampleVCN"
  dns_label      = "tfexamplevcn"
}

resource "oci_core_subnet" "test_subnet" {
  cidr_block     = "10.0.0.0/24"
  display_name   = "regionalSubnet"
  dns_label      = "regionalsubnet"
  compartment_id = var.compartment_id
  vcn_id         = oci_core_vcn.test_vcn.id
}

resource "oci_core_network_security_group" "test_network_security_group" {
  #Required
  compartment_id = var.compartment_id
  vcn_id         = oci_core_vcn.test_vcn.id
}

resource "oci_database_management_db_management_private_endpoint" "test_db_management_private_endpoint" {
  #Required
  compartment_id = var.compartment_id
  name = var.db_management_private_endpoint_name
  subnet_id = oci_core_subnet.test_subnet.id

  #Optional
  description = var.db_management_private_endpoint_description
  nsg_ids   = [oci_core_network_security_group.test_network_security_group.id]
}

data "oci_database_management_db_management_private_endpoint" "test_db_management_private_endpoint" {
  db_management_private_endpoint_id = oci_database_management_db_management_private_endpoint.test_db_management_private_endpoint.id
}

data "oci_database_management_db_management_private_endpoints" "test_db_management_private_endpoints" {
  #Required
  compartment_id = var.compartment_id
}

data "oci_database_management_db_management_private_endpoints" "test_db_management_private_endpoints_with_name" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  name = var.db_management_private_endpoint_name
  vcn_id = oci_core_vcn.test_vcn.id
  state = var.db_management_private_endpoint_state
}