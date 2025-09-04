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
  // Define the region where destination backup will be created.
 }

variable "compartment_ocid" {
}

provider "oci" {
  # un-ignore to run backwards compatibility testing
  #version = "6.32.0"
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

resource "oci_core_network_security_group" "test_network_security_group" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.test_vcn.id
}

resource "oci_mysql_mysql_backup" "test_mysql_backup" {
  db_system_id = oci_mysql_mysql_db_system.test_mysql_backup_db_system.id

  #Optional
  # To trigger backup validation, set the validate_trigger field to 1 or any higher value.
  # Each increment to validate_trigger will re-run the backup validation.
  # Note: Validation is supported only via the update backup resource.
  # validate_trigger = "1"

  # Set this value to true or false before performing backup validation.
  # validate_backup_details {
  #   is_prepared_backup_required = "false"
  # }
}

resource "oci_mysql_mysql_db_system" "test_mysql_backup_db_system" {
  #Required
  admin_password      = "BEstrO0ng_#11"
  admin_username      = "adminUser"
  availability_domain = data.oci_identity_availability_domains.test_availability_domains.availability_domains[0].name
  compartment_id      = var.compartment_ocid
  configuration_id    = data.oci_mysql_mysql_configurations.test_mysql_configurations.configurations[0].id
  shape_name          = "MySQL.VM.Standard.E3.1.8GB"
  subnet_id           = oci_core_subnet.test_subnet.id

  #Optional
  data_storage_size_in_gb = "50"
}

resource "oci_mysql_mysql_configuration" "test_mysql_configuration" {
	#Required
	compartment_id = var.compartment_ocid
	shape_name = "MySQL.VM.Standard.E3.1.8GB"

	#Optional
	description = "test configuration created by terraform"
	display_name = "terraform test configuration"
	parent_configuration_id = data.oci_mysql_mysql_configurations.test_mysql_configurations.configurations[0].id
	variables {

		#Optional
		max_connections = "501"
	}
}

resource "oci_mysql_mysql_db_system" "test_mysql_db_system" {
  #Required
  admin_password      = "BEstrO0ng_#11"
  admin_username      = "adminUser"
  availability_domain = data.oci_identity_availability_domains.test_availability_domains.availability_domains[0].name
  compartment_id      = var.compartment_ocid
  configuration_id    = oci_mysql_mysql_configuration.test_mysql_configuration.id
  shape_name          = "MySQL.VM.Standard.E3.1.8GB"
  subnet_id           = oci_core_subnet.test_subnet.id

  #Optional
  backup_policy {
    is_enabled        = "false"
    retention_in_days = "10"
    window_start_time = "01:00-00:00"
    copy_policies {
        backup_copy_retention_in_days = "2"
        copy_to_region                = "us-phoenix-1"
    }
  }

  #Optional
  read_endpoint {
    is_enabled       = "false"
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

  nsg_ids       = [oci_core_network_security_group.test_network_security_group.id]
  port          = "3306"
  port_x        = "33306"

  # Creating DB System using a backup
  source {
    backup_id   = oci_mysql_mysql_backup.test_mysql_backup.id
    source_type = "BACKUP"
  }

  #Optional
  access_mode = "UNRESTRICTED"
  database_mode = "READ_WRITE"

  #Optional
  crash_recovery = "ENABLED"
  database_management = "DISABLED"
  secure_connections {
    certificate_generation_type = "SYSTEM"
  }
  encrypt_data {
    key_generation_type = "SYSTEM"
  }

  #Optional
  deletion_policy {
    automatic_backup_retention = "RETAIN"
    final_backup = "SKIP_FINAL_BACKUP"
    is_delete_protected = "false"
  }

  #Optional
  customer_contacts {
    email = "email@discard.oracle.com"
  }

  #Optional
  data_storage {
    is_auto_expand_storage_enabled = "false"
    max_storage_size_in_gbs = "100"
  }

  #Optional
  rest {
    configuration = "DISABLED"
    port = "443"
  }
}

data "oci_mysql_mysql_configurations" "test_mysql_configurations" {
  compartment_id = var.compartment_ocid

  #Optional
  state        = "ACTIVE"
  shape_name   = "MySQL.VM.Standard.E3.1.8GB"
}

data "oci_mysql_mysql_backups" "test_mysql_backups" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  backup_id = oci_mysql_mysql_backup.test_mysql_backup.id
}

data "oci_mysql_shapes" "test_shapes" {
  compartment_id = var.compartment_ocid
  availability_domain = lower(
    data.oci_identity_availability_domains.test_availability_domains.availability_domains[0].name,
  )
}

data "oci_identity_availability_domains" "test_availability_domains" {
  compartment_id = var.tenancy_ocid
}

data "oci_mysql_mysql_db_system" "test_mysql_db_system" {
  #Required
  db_system_id = oci_mysql_mysql_db_system.test_mysql_backup_db_system.id
}

data "oci_mysql_mysql_backup" "test_mysql_backup" {
  #Required
  backup_id = oci_mysql_mysql_backup.test_mysql_backup.id
}

output "configuration_id" {
  value = data.oci_mysql_mysql_configurations.test_mysql_configurations.configurations[0].id
}

