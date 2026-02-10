// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
  // Define the region where configuration will be created.
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

resource "oci_mysql_mysql_configuration" "test_mysql_configuration_options" {
  #Required
  compartment_id = var.compartment_ocid
  shape_name     = "MySQL.32"

  #Optional
  description                = "example configuration created using options"
  display_name               = "terraform test configuration options"
  parent_configuration_id    = data.oci_mysql_mysql_configurations.test_mysql_configurations.configurations[0].id

  # Only `options` OR `variables` can be specified (NOT both)
  options {
    name  = "binlog_expire_logs_seconds"
    value = "3601"
  }

  options {
    name  = "autocommit"
    value = "OFF"
  }
  # Add further options as needed
}

data "oci_mysql_mysql_configurations" "test_mysql_configurations" {
  compartment_id = var.compartment_ocid

  #Optional
  state        = "ACTIVE"
  shape_name   = "MySQL.32"
  type         = ["DEFAULT"]
}

data "oci_identity_availability_domains" "test_availability_domains" {
  compartment_id = var.tenancy_ocid
}
