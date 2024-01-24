// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.

# Vars provided when configuring tf for oci
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

# Creating a private subnet to used to access the dbSystem
resource "oci_core_subnet" "test_subnet" {
  cidr_block     = "10.0.0.0/24"
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.test_vcn.id
  prohibit_public_ip_on_vnic = true
}

# Creating a VCN for the private subnet
resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_ocid
}

# Creating DbSystem Backup
resource "oci_psql_backup" "test_backup" {
  db_system_id = oci_psql_db_system.test_db_system.id
  compartment_id = var.compartment_ocid
  display_name = "tf-test-dbSystem-backup"
}

# Creating DbSystem
resource "oci_psql_db_system" "test_db_system" {
  #Required
  db_version          = "14"
  display_name = "tf-test-dbSystem"
  network_details {
    subnet_id = oci_core_subnet.test_subnet.id
  }
  shape = "PostgreSQL.VM.Standard.E4.Flex.2.32GB"
  storage_details {
    is_regionally_durable = true
    system_type = "OCI_OPTIMIZED_STORAGE"
  }
  credentials {
    username = "adminUser"
    password_details {
      password_type = "PLAIN_TEXT"
      password = "BEstrO0ng_#11"
    }
  }
  compartment_id      = var.compartment_ocid
  instance_count = "1"
  system_type = "OCI_OPTIMIZED_STORAGE"

  # Specify patch operations after creating the dbSystem resource and update instance_count with it
  ## Add Replica
  patch_operations {
    operation = "INSERT"
    selection = "instances"
    value {
      displayName = "my-db-instance"
      description = "my description"
    }
  }

  ## Remove Replica
  patch_operations {
    operation = "REMOVE"
    selection = "instancesinstances[?id == '${var.db_instance_id}']"
  }
}