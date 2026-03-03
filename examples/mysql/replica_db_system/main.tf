// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.

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
  # uncomment to run backwards compatibility testing
  # to avoid compatibility issues use the lastest version released:
  # https://github.com/oracle/terraform-provider-oci/releases
  # version = "8.7.0"
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_core_subnet" "test_subnet" {
  cidr_block     = "10.0.0.0/24"
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.test_vcn.id
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_ocid
}

resource "oci_mysql_mysql_configuration" "test_mysql_configuration" {
  #Required
  compartment_id = var.compartment_ocid
  shape_name = "MySQL.4"

  #Optional
  description = "test configuration created by terraform"
  display_name = "terraform test configuration"
  parent_configuration_id = data.oci_mysql_mysql_configurations.test_mysql_configurations.configurations[0].id

  #Optional
  options {
    name  = "binlog_expire_logs_seconds"
    value = "345600"
  }
}

resource "oci_mysql_mysql_db_system" "test_mysql_db_system" {
  display_name        = "db-system-source"
  admin_password      = "BEstrO0ng_#11"
  admin_username      = "adminUser"
  availability_domain = data.oci_identity_availability_domains.test_availability_domains.availability_domains[0].name
  compartment_id      = var.compartment_ocid
  configuration_id    = oci_mysql_mysql_configuration.test_mysql_configuration.id
  shape_name          = "MySQL.4"
  subnet_id           = oci_core_subnet.test_subnet.id
  data_storage_size_in_gb = "50"

  #Optional
  backup_policy {
    is_enabled        = "false"
    retention_in_days = "10"
    window_start_time = "01:00-00:00"
  }

  database_management = "DISABLED"
}

data "oci_mysql_mysql_configurations" "test_mysql_configurations" {
  compartment_id = var.compartment_ocid

  #Optional
  state        = "ACTIVE"
  shape_name   = "MySQL.4"
}

data "oci_identity_availability_domains" "test_availability_domains" {
  compartment_id = var.tenancy_ocid
}

resource "oci_mysql_mysql_db_system" "test_mysql_db_system_replica" {
  display_name        = "db-system-replica"
  admin_password      = "BEstrO0ng_#11"
  admin_username      = "adminUser"
  availability_domain = data.oci_identity_availability_domains.test_availability_domains.availability_domains[0].name
  compartment_id      = var.compartment_ocid
  configuration_id    = data.oci_mysql_mysql_configurations.test_mysql_configurations.configurations[0].id
  shape_name          = "MySQL.4"
  subnet_id           = oci_core_subnet.test_subnet.id

  #Optional
  backup_policy {
    is_enabled        = "false"
    retention_in_days = "10"
    window_start_time = "01:00-00:00"
  }

  database_management = "DISABLED"

  source {
    #Required
    source_type = "DBSYSTEM"
    db_system_id   = oci_mysql_mysql_db_system.test_mysql_db_system.id

    #Optional
    channel {
      source_password = "BEstrO0ng_#11"
      source_username = "adminUser"
      ssl_mode = "REQUIRED"
    }
  }
}

output "replica_db_system_id" {
  value = oci_mysql_mysql_db_system.test_mysql_db_system_replica.id
}