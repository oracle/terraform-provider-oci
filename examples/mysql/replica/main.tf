// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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

resource "oci_core_subnet" "test_subnet" {
  cidr_block     = "10.0.0.0/24"
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.test_vcn.id
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_ocid
}

resource "oci_mysql_mysql_db_system" "test_mysql_db_system" {
  #Required
  admin_password      = "BEstrO0ng_#11"
  admin_username      = "adminUser"
  availability_domain = data.oci_identity_availability_domains.test_availability_domains.availability_domains[0].name
  compartment_id      = var.compartment_ocid
  configuration_id    = data.oci_mysql_mysql_configurations.test_mysql_configurations.configurations[0].id
  shape_name          = "MySQL.VM.Standard.E3.4.64GB"
  subnet_id           = oci_core_subnet.test_subnet.id
  data_storage_size_in_gb = "50"

  #Optional
  backup_policy {
    is_enabled        = "false"
    retention_in_days = "10"
    window_start_time = "01:00-00:00"
  }

  #defined_tags  = {"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "${var.mysql_defined_tags_value}"}
  #freeform_tags = var.mysql_freeform_tags
  description = "MySQL Database Service"

  display_name   = "DBSystem001"
  fault_domain   = "FAULT-DOMAIN-1"
  hostname_label = "hostnameLabel"
  ip_address     = "10.0.0.8"

  maintenance {
    window_start_time = "sun 01:00"
  }

  port          = "3306"
  port_x        = "33306"

}

data "oci_mysql_mysql_configurations" "test_mysql_configurations" {
  compartment_id = var.compartment_ocid

  #Optional
  state        = "ACTIVE"
  shape_name   = "MySQL.VM.Standard.E3.4.64GB"
}

data "oci_mysql_shapes" "test_shapes" {
  compartment_id = var.compartment_ocid
}

data "oci_identity_availability_domains" "test_availability_domains" {
  compartment_id = var.tenancy_ocid
}

output "configuration_id" {
  value = data.oci_mysql_mysql_configurations.test_mysql_configurations.configurations[0].id
}

resource "oci_mysql_replica" "test_replica" {
  #Required
  db_system_id = oci_mysql_mysql_db_system.test_mysql_db_system.id

  #Optional
  #defined_tags        = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, "value")
  description          = "description"
  display_name         = "displayName"
  #freeform_tags       = { "bar-key" = "value" }
  is_delete_protected  = false
  replica_overrides {
    configuration_id   = data.oci_mysql_mysql_configurations.test_mysql_configurations.configurations[0].id
    mysql_version      = "8.1.0"
    shape_name         = "MySQL.VM.Standard.E3.4.64GB"
  }
}

data "oci_mysql_replicas" "test_replicas" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  db_system_id = oci_mysql_mysql_db_system.test_mysql_db_system.id
  display_name = "displayName"
  replica_id   = oci_mysql_replica.test_replica.id
  state        = "ACTIVE"
}

data "oci_mysql_replica" "test_replica" {
  #Required
  replica_id   = oci_mysql_replica.test_replica.id
}