# $Header$
#
# Copyright (c) 2024, Oracle and/or its affiliates. All rights reserved.
#
#    NAME
#      main.tf
#
#    USAGE
#      Use the following path for the Example & Backward Compatibility tests: database/db_systems/db_vm/db_vm_std_x86
#    NOTES
#      Terraform Integration Test: TestResourceDatabaseDBSystemVMStdx86
#
#    FILE(S)
#      database_db_system_resource_vm_std_x86_test.go
#
#    MODIFIED   MM/DD/YY
#    aavadhan   08/18/2025 - Created

data "oci_identity_availability_domain" "ad" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}

# Get DB node list
data "oci_database_db_nodes" "db_nodes" {
  compartment_id = var.compartment_ocid
  db_system_id   = oci_database_db_system.test_db_system.id
}

# Get DB node details
data "oci_database_db_node" "db_node_details" {
  db_node_id = data.oci_database_db_nodes.db_nodes.db_nodes[0]["id"]
}


data "oci_database_db_homes" "db_homes" {
  compartment_id = var.compartment_ocid
  db_system_id   = oci_database_db_system.test_db_system.id
}

data "oci_database_databases" "databases" {
  compartment_id = var.compartment_ocid
  db_home_id     = data.oci_database_db_homes.db_homes.db_homes[0].db_home_id
}

data "oci_database_db_versions" "test_db_versions_by_db_system_id" {
  compartment_id = var.compartment_ocid
  db_system_id   = oci_database_db_system.test_db_system.id
}


data "oci_database_db_system_shapes" "test_db_system_shapes" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid

  filter {
    name   = "shape"
    values = [var.db_system_shape]
  }
}

data "oci_database_db_systems" "db_systems" {
  compartment_id = var.compartment_ocid

  filter {
    name   = "id"
    values = [oci_database_db_system.test_db_system.id]
  }
}


