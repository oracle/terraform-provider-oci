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
}

resource "oci_psql_db_system" "test_flexdb_system" {
  #Required
  db_version          = "14"
  display_name = "tf-flex-test-dbSystem"
  network_details {
    subnet_id = oci_core_subnet.test_subnet.id
  }
  shape = "PostgreSQL.VM.Standard.E4.Flex"
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
  instance_ocpu_count = "2"
  instance_memory_size_in_gbs = "10"
  system_type = "OCI_OPTIMIZED_STORAGE"
  config_id = oci_psql_configuration.test_flexible_configuration.id
}

# Creating a dbSystem configuration
resource "oci_psql_configuration" "test_configuration" {
	#Required
	compartment_id = var.compartment_ocid
	shape = "VM.Standard.E4.Flex"
  db_configuration_overrides {
    items {
      config_key = "effective_io_concurrency"
      overriden_config_value = "1"
    }
  }
  db_version = "14"
	display_name = "terraform test configuration"

	#Optional
  instance_memory_size_in_gbs = "64"
  instance_ocpu_count = "4"
	description = "test configuration created by terraform"
}

# Creating a dbSystem configuration
resource "oci_psql_configuration" "test_flexible_configuration" {
        #Required
        compartment_id = var.compartment_ocid
        shape = "VM.Standard.E4.Flex"
  db_configuration_overrides {
    items {
      config_key = "effective_io_concurrency"
      overriden_config_value = "1"
    }
  }
  db_version = "14"
        display_name = "terraform test flex configuration"
        #Optional
  instance_memory_size_in_gbs = "0"
  instance_ocpu_count = "0"
  is_flexible = true
        description = "test configuration created by terraform"
}

data "oci_psql_configurations" "test_configurations" {
  compartment_id = var.compartment_ocid

}
