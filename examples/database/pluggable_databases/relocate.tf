// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

data "oci_identity_availability_domains" "ADsRelocate" {
  compartment_id = "${var.compartment_id}"
}

resource "oci_core_virtual_network" "relocate" {
  compartment_id = "${var.compartment_id}"
  cidr_block = "10.1.0.0/16"
  display_name = "-tf-vcn-clone"
  dns_label = "tfvcnclone"
}

resource "oci_core_route_table" "relocate" {
  compartment_id = "${var.compartment_id}"
  vcn_id = "${oci_core_virtual_network.relocate.id}"
  route_rules {
    cidr_block = "0.0.0.0/0"
    network_entity_id = "${oci_core_internet_gateway.relocate.id}"
  }
}
resource "oci_core_internet_gateway" "relocate" {
  compartment_id = "${var.compartment_id}"
  vcn_id = "${oci_core_virtual_network.relocate.id}"
  display_name = "-tf-internet-gateway-clone"
}

resource "oci_core_subnet" "relocate" {
  availability_domain = "${data.oci_identity_availability_domains.ADsRelocate.availability_domains.0.name}"
  cidr_block          = "10.1.20.0/24"
  display_name        = "TFSubnetClone1"
  compartment_id      = "${var.compartment_id}"
  vcn_id              = "${oci_core_virtual_network.relocate.id}"
  route_table_id      = "${oci_core_route_table.relocate.id}"
  dhcp_options_id     = "${oci_core_virtual_network.relocate.default_dhcp_options_id}"
  security_list_ids   = ["${oci_core_virtual_network.relocate.default_security_list_id}"]
  dns_label           = "tfsubnetclone"
}
resource "oci_core_subnet" "t2Clone" {
  availability_domain = "${data.oci_identity_availability_domains.ADsRelocate.availability_domains.0.name}"
  cidr_block          = "10.1.21.0/24"
  display_name        = "TFSubnetClone2"
  compartment_id      = "${var.compartment_id}"
  vcn_id              = "${oci_core_virtual_network.relocate.id}"
  route_table_id      = "${oci_core_route_table.relocate.id}"
  dhcp_options_id     = "${oci_core_virtual_network.relocate.default_dhcp_options_id}"
  security_list_ids   = ["${oci_core_virtual_network.relocate.default_security_list_id}"]
  dns_label           = "tfsubnetclone2"
}
resource "oci_core_network_security_group" "test_network_security_group_clone" {
  compartment_id  = "${var.compartment_id}"
  vcn_id            = "${oci_core_virtual_network.relocate.id}"
  display_name      =  "displayName"
}

resource "oci_core_network_security_group" "test_network_security_group_clone2" {
  compartment_id = "${var.compartment_id}"
  vcn_id            = "${oci_core_virtual_network.relocate.id}"
}

resource "oci_database_db_system" "relocate" {
  compartment_id = "${var.compartment_id}"
  subnet_id = "${oci_core_subnet.relocate.id}"
  database_edition = "ENTERPRISE_EDITION"
  availability_domain = "${data.oci_identity_availability_domains.ADsRelocate.availability_domains.0.name}"
  disk_redundancy = "NORMAL"
  shape = "VM.Standard2.1"
  ssh_public_keys = ["ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"]
  display_name = "-tf-dbSystem-clone-001"
  domain = "${oci_core_subnet.relocate.dns_label}.${oci_core_virtual_network.relocate.dns_label}.oraclevcn.com"
  hostname = "myOracleDB" // this will be lowercased server side
  data_storage_size_in_gb = "256"
  license_model = "LICENSE_INCLUDED"
  node_count = "1"
  cpu_core_count =  "${var.cpu_core_count}"
  fault_domains = ["FAULT-DOMAIN-1"]
  db_home {
    db_version = "21.8.0.0"
    display_name = "-tf-db-home-clone"
    database {
      admin_password = "BEstrO0ng_#11"
      db_name = "aTFdbC"
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
  nsg_ids = ["${oci_core_network_security_group.test_network_security_group_clone.id}"]
  lifecycle {
    ignore_changes = [
      db_home.0.db_version,
      defined_tags,
      db_home.0.database.0.defined_tags,
    ]
  }
}
data "oci_database_db_systems" "relocate" {
  compartment_id = "${var.compartment_id}"
  filter {
    name   = "id"
    values = ["${oci_database_db_system.relocate.id}"]
  }
}
data "oci_database_db_homes" "relocate" {
  compartment_id = "${var.compartment_id}"
  db_system_id = "${oci_database_db_system.relocate.id}"
  filter {
    name   = "db_system_id"
    values = ["${oci_database_db_system.relocate.id}"]
  }
}
data "oci_database_db_home" "relocate" {
  db_home_id = "${data.oci_database_db_homes.relocate.db_homes.0.db_home_id}"
}
data "oci_database_databases" "relocate" {
  compartment_id = "${var.compartment_id}"
  db_home_id = "${data.oci_database_db_homes.relocate.db_homes.0.id}"
  filter {
    name   = "db_name"
    values = ["${oci_database_db_system.relocate.db_home.0.database.0.db_name}"]
  }
}
data "oci_database_database" "relocate" {
  database_id = "${data.oci_database_databases.relocate.databases.0.id}"
}

resource "oci_database_pluggable_databases_local_clone" "test_pluggable_databases_relocate" {
  pdb_name = "pdbRelocate"
  container_database_id = "${data.oci_database_database.relocate.id}"
  tde_wallet_password = "BEstrO0ng_#11"
  pdb_admin_password = "BEstrO0ng_#11"
  pdb_creation_type_details {
    creation_type = "RELOCATE"
    source_pluggable_database_id = "${oci_database_pluggable_database.test_pluggable_database2.id}"
    source_container_database_admin_password = "BEstrO0ng_#11"
  }
}