// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}
variable "ssh_public_key" {}

#ocid of database_software_image, refer to database_software_image example on how to create a database software image
variable "test_database_software_image_ocid" {}

provider "oci" {
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region           = "${var.region}"
}

resource "oci_core_subnet" "test_subnet" {
  availability_domain        = "${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}"
  cidr_block                 = "10.0.0.0/24"
  compartment_id             = "${var.compartment_id}"
  dhcp_options_id            = "${oci_core_vcn.test_vcn.default_dhcp_options_id}"
  display_name               = "MySubnet"
  dns_label                  = "dnslabel"
  prohibit_public_ip_on_vnic = "false"
  route_table_id             = "${oci_core_route_table.test_route_table.id}"
  security_list_ids          = ["${oci_core_vcn.test_vcn.default_security_list_id}"]
  vcn_id                     = "${oci_core_vcn.test_vcn.id}"
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = "${var.compartment_id}"
  display_name   = "displayName"
  dns_label      = "dnslabel"
}

resource "oci_core_route_table" "test_route_table" {
  compartment_id = "${var.compartment_id}"
  display_name   = "MyRouteTable"

  route_rules {
    description       = "description"
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = "${oci_core_internet_gateway.test_internet_gateway.id}"
  }

  vcn_id = "${oci_core_vcn.test_vcn.id}"
}

resource "oci_core_internet_gateway" "test_internet_gateway" {
  compartment_id = "${var.compartment_id}"
  display_name   = "MyInternetGateway"
  enabled        = "false"
  vcn_id         = "${oci_core_vcn.test_vcn.id}"
}

data "oci_identity_availability_domains" "test_availability_domains" {
  compartment_id = "${var.tenancy_ocid}"
}

resource "oci_database_db_system" "test_db_system" {
  availability_domain     = "${lower("${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}")}"
  compartment_id          = "${var.compartment_id}"
  subnet_id               = "${oci_core_subnet.test_subnet.id}"
  database_edition        = "ENTERPRISE_EDITION"
  disk_redundancy         = "NORMAL"
  shape                   = "BM.DenseIO1.36"
  cpu_core_count          = "2"
  ssh_public_keys         = ["${var.ssh_public_key}"]
  domain                  = "${oci_core_subnet.test_subnet.subnet_domain_name}"
  hostname                = "myOracleDB"
  data_storage_size_in_gb = "256"
  license_model           = "LICENSE_INCLUDED"
  node_count              = "1"
  display_name            = "tfDbSystemTest"

  db_home {
    db_version   = "12.1.0.2"
    display_name = "dbHome1"

    database {
      admin_password = "BEstrO0ng_#11"
      db_name        = "tfDbName"
    }
  }
}

data "oci_database_db_homes" "t" {
  compartment_id = "${var.compartment_id}"
  db_system_id   = "${oci_database_db_system.test_db_system.id}"

  filter {
    name   = "display_name"
    values = ["dbHome1"]
  }
}

data "oci_database_databases" "db" {
  compartment_id = "${var.compartment_id}"
  db_home_id     = "${data.oci_database_db_homes.t.db_homes.0.db_home_id}"
}

resource "oci_database_db_home" "test_db_home_source_none" {
  database {
    admin_password = "BEstrO0ng_#11"
    character_set  = "AL32UTF8"

    db_backup_config {
      auto_backup_enabled     = "true"
      auto_backup_window      = "SLOT_TWO"
      recovery_window_in_days = "10"
    }

    database_software_image_id = "${var.test_database_software_image_ocid}"
    db_name                    = "dbNonegk"
    db_workload                = "OLTP"
    ncharacter_set             = "AL16UTF16"
    pdb_name                   = "pdbName"
  }

  database_software_image_id = "${var.test_database_software_image_ocid}"
  db_system_id               = "${oci_database_db_system.test_db_system.id}"
  db_version                 = "12.1.0.2"
  display_name               = "createdDbHomeNone"
  source                     = "NONE"
}

resource "oci_database_db_home" "test_db_home_source_database" {
  database {
    admin_password             = "BEstrO0ng_#11"
    backup_tde_password        = "BEstrO0ng_#11"
    database_id                = "${data.oci_database_databases.db.databases.0.id}"
    db_name                    = "dbDbgrk"
    database_software_image_id = "${var.test_database_software_image_ocid}"
  }

  database_software_image_id = "${var.test_database_software_image_ocid}"
  db_system_id               = "${oci_database_db_system.test_db_system.id}"
  display_name               = "createdDbHomeDatabase"
  source                     = "DATABASE"
}

resource "oci_database_backup" "test_backup" {
  database_id  = "${data.oci_database_databases.db.databases.0.id}"
  display_name = "Monthly Backup"
}

resource "oci_database_db_home" "test_db_home_source_db_backup" {
  database {
    admin_password             = "BEstrO0ng_#11"
    backup_id                  = "${oci_database_backup.test_backup.id}"
    backup_tde_password        = "BEstrO0ng_#11"
    db_name                    = "dbBackup"
    database_software_image_id = "${var.test_database_software_image_ocid}"
  }

  db_system_id               = "${oci_database_db_system.test_db_system.id}"
  display_name               = "createdDbHomeBackup"
  source                     = "DB_BACKUP"
  database_software_image_id = "${var.test_database_software_image_ocid}"
}
