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

variable "compartment_id" {
}

variable "ssh_public_key" {
}

variable "ssl_secret_id" {
  default = "secret_id"
}

variable "ssl_user_name" {
  default = "dbsnmp"
}

variable "pluggable_database_pluggabledatabasemanagements_management_port" {
  default = 10
}

variable defined_tag_namespace_name {
    default = ""
}

variable "pluggable_database_pluggabledatabasemanagements_management_protocol" {
  default = "TCP"
}

variable "pluggable_database_pluggabledatabasemanagements_management_role" {
  default = "NORMAL"
}

variable enable_pluggabledatabasemanagement {
  default = "true"
}

variable "enable_management" {
  default = "true"
}

variable "protocol" {
  default = "TCP"
}

variable "port" {
  default = "1521"
}

variable "role" {
  default = "NORMAL"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}


data "oci_identity_availability_domains" "test_availability_domains" {
    compartment_id = "${var.tenancy_ocid}"
}


data "oci_identity_availability_domains" "ADs" {
    compartment_id = "${var.compartment_id}"
}

resource "oci_core_virtual_network" "t" {
    compartment_id = "${var.compartment_id}"
    cidr_block = "10.1.0.0/16"
    display_name = "-tf-vcn"
    dns_label = "tfvcn"
}

resource "oci_core_route_table" "t" {
    compartment_id = "${var.compartment_id}"
    vcn_id = "${oci_core_virtual_network.t.id}"
    route_rules {
            cidr_block = "0.0.0.0/0"
            network_entity_id = "${oci_core_internet_gateway.t.id}"
    }
}
resource "oci_core_internet_gateway" "t" {
    compartment_id = "${var.compartment_id}"
    vcn_id = "${oci_core_virtual_network.t.id}"
    display_name = "-tf-internet-gateway"
}

resource "oci_core_subnet" "t" {
    availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
    cidr_block          = "10.1.20.0/24"
    display_name        = "TFSubnet1"
    compartment_id      = "${var.compartment_id}"
    vcn_id              = "${oci_core_virtual_network.t.id}"
    route_table_id      = "${oci_core_route_table.t.id}"
    dhcp_options_id     = "${oci_core_virtual_network.t.default_dhcp_options_id}"
    security_list_ids   = ["${oci_core_virtual_network.t.default_security_list_id}"]
    dns_label           = "tfsubnet"
}
resource "oci_core_subnet" "t2" {
    availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
    cidr_block          = "10.1.21.0/24"
    display_name        = "TFSubnet2"
    compartment_id      = "${var.compartment_id}"
    vcn_id              = "${oci_core_virtual_network.t.id}"
    route_table_id      = "${oci_core_route_table.t.id}"
    dhcp_options_id     = "${oci_core_virtual_network.t.default_dhcp_options_id}"
    security_list_ids   = ["${oci_core_virtual_network.t.default_security_list_id}"]
    dns_label           = "tfsubnet2"
}
resource "oci_core_network_security_group" "test_network_security_group" {
    compartment_id  = "${var.compartment_id}"
    vcn_id            = "${oci_core_virtual_network.t.id}"
    display_name      =  "displayName"
}

resource "oci_core_network_security_group" "test_network_security_group2" {
    compartment_id = "${var.compartment_id}"
    vcn_id            = "${oci_core_virtual_network.t.id}"
}
resource "oci_identity_tag_namespace" "tag-namespace1" {
        #Required
        compartment_id = "${var.tenancy_ocid}"
        description = "example tag namespace"
        name = "${var.defined_tag_namespace_name != "" ? var.defined_tag_namespace_name : "example-tag-namespace-all"}"

        is_retired = false
}

resource "oci_identity_tag" "tag1" {
        #Required
        description = "example tag"
        name = "example-tag"
        tag_namespace_id = "${oci_identity_tag_namespace.tag-namespace1.id}"

        is_retired = false
}

resource "oci_database_db_system" "t" {
        compartment_id = "${var.compartment_id}"
        subnet_id = "${oci_core_subnet.t.id}"
        database_edition = "ENTERPRISE_EDITION"
        availability_domain = "${data.oci_identity_availability_domains.ADs.availability_domains.0.name}"
        disk_redundancy = "NORMAL"
        shape = "VM.Standard1.1"
        ssh_public_keys = ["ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"]
        display_name = "-tf-dbSystem-001"
        domain = "${oci_core_subnet.t.dns_label}.${oci_core_virtual_network.t.dns_label}.oraclevcn.com"
        hostname = "myOracleDB" // this will be lowercased server side
        data_storage_size_in_gb = "256"
        license_model = "LICENSE_INCLUDED"
        node_count = "1"
        fault_domains = ["FAULT-DOMAIN-1"]
        db_home {
                db_version = "19.12.0.0"
                display_name = "-tf-db-home"
                database {
                        admin_password = "BEstrO0ng_#11"
                        db_name = "aTFdb"
                        character_set = "AL32UTF8"
                        defined_tags = "${map("example-tag-namespace-all.example-tag", "originalValue")}"
                        freeform_tags = {"Department" = "Finance"}
                        ncharacter_set = "AL16UTF16"
                        db_workload = "OLTP"
                        pdb_name = "pdbName"
                }
        }
        db_system_options {
                storage_management = "LVM"
        }
        defined_tags = "${map("example-tag-namespace-all.example-tag", "originalValue")}"
        freeform_tags = {"Department" = "Finance"}
        nsg_ids = ["${oci_core_network_security_group.test_network_security_group.id}"]
        lifecycle {
                ignore_changes = [
                        db_home.0.db_version,
                        defined_tags,
                        db_home.0.database.0.defined_tags,
                ]
        }
}
data "oci_database_db_systems" "t" {
        compartment_id = "${var.compartment_id}"
        filter {
                name   = "id"
                values = ["${oci_database_db_system.t.id}"]
        }
}
data "oci_database_db_homes" "t" {
        compartment_id = "${var.compartment_id}"
        db_system_id = "${oci_database_db_system.t.id}"
        filter {
                name   = "db_system_id"
                values = ["${oci_database_db_system.t.id}"]
        }
}
data "oci_database_db_home" "t" {
        db_home_id = "${data.oci_database_db_homes.t.db_homes.0.db_home_id}"
}
data "oci_database_databases" "t" {
        compartment_id = "${var.compartment_id}"
        db_home_id = "${data.oci_database_db_homes.t.db_homes.0.id}"
        filter {
                name   = "db_name"
                values = ["${oci_database_db_system.t.db_home.0.database.0.db_name}"]
        }
}
data "oci_database_database" "t" {
          database_id = "${data.oci_database_databases.t.databases.0.id}"
}
resource "oci_database_pluggable_database" "test_pluggable_database" {
        container_database_id = "${data.oci_database_database.t.id}"
        lifecycle {
        ignore_changes = ["defined_tags"]
        }
        pdb_admin_password = "BEstrO0ng_#11"
        pdb_name = "SalesPdb"
        tde_wallet_password = "BEstrO0ng_#11"
}

resource "oci_database_management_db_management_private_endpoint" "test_db_management_private_endpoint" {
  compartment_id        = var.compartment_id
  name                  = "${oci_database_db_system.t.db_home.0.database.0.db_name}"
  subnet_id             = "${oci_core_subnet.t.id}"
}

resource "oci_database_cloud_database_management" "test" {
  database_id           = data.oci_database_databases.t.databases.0.id
  management_type       = "ADVANCED"
  private_end_point_id  = oci_database_management_db_management_private_endpoint.test_db_management_private_endpoint.id
  service_name          = join("", [data.oci_database_databases.t.databases.0.db_unique_name, oci_core_subnet.t.subnet_domain_name])
  credentialdetails {
    user_name           = var.ssl_user_name
    password_secret_id  = var.ssl_secret_id
  }
  enable_management     = var.enable_management
  protocol              = var.protocol
  port                  = var.port
  role                  = var.role
  ssl_secret_id         = var.ssl_secret_id
}

resource "oci_database_pluggable_database_pluggabledatabasemanagements_management" "test_pluggable_database_pluggabledatabasemanagements_management" {
  pluggable_database_id              = oci_database_pluggable_database.test_pluggable_database.id
  enable_pluggabledatabasemanagement = var.enable_pluggabledatabasemanagement
  credential_details {
    password_secret_id = var.ssl_secret_id
    user_name          = var.ssl_user_name
  }
  private_end_point_id = oci_database_management_db_management_private_endpoint.test_db_management_private_endpoint.id
  service_name         = join("", [data.oci_database_databases.t.databases.0.db_unique_name, oci_core_subnet.t.subnet_domain_name])
  port                 = var.pluggable_database_pluggabledatabasemanagements_management_port
  protocol             = var.pluggable_database_pluggabledatabasemanagements_management_protocol
  role                 = var.pluggable_database_pluggabledatabasemanagements_management_role
  ssl_secret_id        = var.ssl_secret_id
}

