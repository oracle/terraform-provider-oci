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
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_core_subnet" "test_subnet" {
  cidr_block      = "10.0.0.0/24"
  ipv6cidr_blocks = [cidrsubnet(oci_core_vcn.test_vcn.ipv6cidr_blocks[0], 8, 0)]
  compartment_id  = var.compartment_ocid
  vcn_id          = oci_core_vcn.test_vcn.id
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.0.0.0/16"
  is_ipv6enabled = "true"
  compartment_id = var.compartment_ocid
}

resource "oci_mysql_mysql_db_system" "test_mysql_db_system_ipv6" {
  display_name            = "db-system-ipv6"
  admin_password          = "BEstrO0ng_#11"
  admin_username          = "adminUser"
  availability_domain     = data.oci_identity_availability_domains.test_availability_domains.availability_domains[0].name
  compartment_id          = var.compartment_ocid
  configuration_id        = data.oci_mysql_mysql_configurations.test_mysql_configurations.configurations[0].id
  shape_name              = "MySQL.4"
  subnet_id               = oci_core_subnet.test_subnet.id
  data_storage_size_in_gb = "50"
  is_ipv6enabled          = "true"

  ipv6address_ipv6subnet_cidr_pair_details {
    ipv6subnet_cidr = oci_core_subnet.test_subnet.ipv6cidr_blocks[0]
    ipv6address = cidrhost(oci_core_subnet.test_subnet.ipv6cidr_blocks[0], 1000)
  }

  lifecycle {
    // Ignore changes in the generated IPV6 IP address.
    ignore_changes = [
      ipv6address_ipv6subnet_cidr_pair_details[0].ipv6address
    ]
  }

  #Optional
  backup_policy {
    is_enabled = "false"
  }

  database_management = "DISABLED"
  hostname_label      = "hostnameLabel"
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

output "test_mysql_db_system_ipv6_id" {
  value = oci_mysql_mysql_db_system.test_mysql_db_system_ipv6.id
}

resource "oci_mysql_channel" "test_channel_ipv6" {
  #Required
  source {
    #Required
    hostname    = oci_mysql_mysql_db_system.test_mysql_db_system_ipv6.endpoints.0.hostname
    password    = "BEstrO0ng_#11"
    source_type = "MYSQL"
    username    = "adminUser"
    ssl_mode    = "REQUIRED"

    #Optional
    must_use_ipv6on_dual_stack = "true"
  }

  target {
    #Required
    db_system_id = oci_mysql_mysql_db_system.test_mysql_db_system_ipv6.id
    target_type  = "DBSYSTEM"
    channel_name = "channelipv6"
  }
}

data "oci_mysql_channels" "test_channels" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  channel_id   = oci_mysql_channel.test_channel_ipv6.id
  db_system_id = oci_mysql_mysql_db_system.test_mysql_db_system_ipv6.id
}

data "oci_mysql_channel" "test_channel" {
  #Required
  channel_id   = oci_mysql_channel.test_channel_ipv6.id
}

output "test_channel_ipv6_id" {
  value = data.oci_mysql_channel.test_channel.id
}
