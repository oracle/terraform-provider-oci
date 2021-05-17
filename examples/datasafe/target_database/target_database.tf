// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
}

variable "compartment_ocid" {
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "exampleVCN"
  dns_label      = "tfexamplevcn"
}

resource "oci_core_subnet" "test_subnet" {
  cidr_block     = "10.0.1.0/24"
  display_name   = "regionalSubnet"
  dns_label      = "regionalsubnet"
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.test_vcn.id
}

resource "oci_data_safe_data_safe_configuration" "test_data_safe_configuration" {
  is_enabled = "true"
}

resource "oci_data_safe_data_safe_private_endpoint" "test_data_safe_private_endpoint" {
  compartment_id = var.compartment_ocid
  display_name   = "PE2"
  subnet_id      = oci_core_subnet.test_subnet.id
  vcn_id         = oci_core_vcn.test_vcn.id
}

variable "target_database_description" {
  default = "description"
}

variable "target_database_display_name" {
  default = "targetDatabase1"
}

resource "random_string" "autonomous_database_admin_password" {
  length      = 16
  min_numeric = 1
  min_lower   = 1
  min_upper   = 1
  min_special = 1
}
variable "autonomous_database_db_workload" {
  default = "OLTP"
}

variable "autonomous_database_freeform_tags" {
  default = {
    "Department" = "Finance"
  }
}

variable "autonomous_database_license_model" {
  default = "LICENSE_INCLUDED"
}

resource "oci_database_autonomous_database" "autonomous_database" {
  #Required
  admin_password           = random_string.autonomous_database_admin_password.result
  compartment_id           = var.compartment_ocid
  cpu_core_count           = "1"
  data_storage_size_in_tbs = "1"
  db_name                  = "adbdb1"

  #Optional
  db_workload                                    = var.autonomous_database_db_workload
  display_name                                   = "example_autonomous_database"
  freeform_tags                                  = var.autonomous_database_freeform_tags
  is_auto_scaling_enabled                        = "true"
  license_model                                  = var.autonomous_database_license_model
  is_preview_version_with_service_terms_accepted = "false"
}

resource "oci_data_safe_target_database" "test_target_database" {
#Required
    compartment_id = var.compartment_ocid
    display_name = var.target_database_display_name

    database_details {
        database_type = "AUTONOMOUS_DATABASE"
        infrastructure_type = "ORACLE_CLOUD"
        autonomous_database_id = oci_database_autonomous_database.autonomous_database.id
    }

   #Optional
     connection_option {
     connection_type = "PRIVATE_ENDPOINT"
     datasafe_private_endpoint_id = oci_data_safe_data_safe_private_endpoint.test_data_safe_private_endpoint.id
     }
     description = var.target_database_description
}

data "oci_data_safe_target_databases" "test_target_databases" {
    compartment_id = var.compartment_ocid
    display_name = var.target_database_display_name
    target_database_id = oci_data_safe_target_database.test_target_database.id
}
