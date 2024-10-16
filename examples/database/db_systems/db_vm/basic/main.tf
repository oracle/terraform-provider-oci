# $Header$
#
# Copyright (c) 2024, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      main.tf - Shepherd Main file
#
#    USAGE
#      Use the following path for Example Test & Backward-Compatibility-Test: database/db_systems/db_vm/basic
#    NOTES
#      Terraform Example: TestResourceDatabaseDBSystemBasic
#    FILES
#
#    DESCRIPTION
#
#    MODIFIED   MM/DD/YY
#    escabrer   10/14/2024 - Created



resource "oci_identity_tag_namespace" "tag-namespace1" {
  #Required
  compartment_id = var.tenancy_ocid
  description = "example tag namespace"
  name = var.defined_tag_namespace_name != "" ? var.defined_tag_namespace_name : "example-tag-namespace-all"

  is_retired = false
}

resource "oci_identity_tag" "tag1" {
  #Required
  description = "example tag"
  name = "example-tag"
  tag_namespace_id = oci_identity_tag_namespace.tag-namespace1.id

  is_retired = false
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block = "10.1.0.0/16"
  compartment_id = var.compartment_id
  display_name = "tfVcn"
  dns_label = "tfvcn"
}

resource "oci_core_route_table" "test_route_table" {
  compartment_id = var.compartment_id
  display_name = "tfRouteTable"
  route_rules {
    cidr_block = "0.0.0.0/0"
    description = "Internal traffic for OCI Services"
    network_entity_id = oci_core_internet_gateway.test_internet_gateway.id
  }
  vcn_id = oci_core_vcn.test_vcn.id
}

resource "oci_core_internet_gateway" "test_internet_gateway" {
  compartment_id = var.compartment_id
  display_name = "tfInternetGateway"
  vcn_id = oci_core_vcn.test_vcn.id
}

resource "oci_core_subnet" "test_subnet" {
  cidr_block = "10.1.20.0/24"
  compartment_id = var.compartment_id
  dhcp_options_id = oci_core_vcn.test_vcn.default_dhcp_options_id
  display_name = "tfSubnet"
  dns_label = "tfsubnet"
  route_table_id = oci_core_route_table.test_route_table.id
  security_list_ids = [oci_core_vcn.test_vcn.default_security_list_id]
  vcn_id = oci_core_vcn.test_vcn.id
}

resource "oci_database_db_system" "test_db_system" {
  availability_domain = data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name
  compartment_id = var.compartment_id
  cpu_core_count = "2"
  data_storage_size_in_gb = "256"
  database_edition = "ENTERPRISE_EDITION"
  db_home {
    database {
      admin_password = "BEstrO0ng_#11"
      character_set = "AL32UTF8"
      db_name = "tfDb"
      db_workload = "OLTP"
      kms_key_id = var.kms_key_id
      kms_key_version_id = var.kms_key_version_id
      ncharacter_set = "AL16UTF16"
      pdb_name = "tfPdb"
      vault_id = var.vault_id
    }
    db_version = "19.25.0.0"
    display_name = "tfDbHome"
  }
  disk_redundancy = "NORMAL"
  display_name = "tfDbSystem"
  domain = oci_core_subnet.test_subnet.subnet_domain_name
  fault_domains = ["FAULT-DOMAIN-1"]
  hostname = "tfOracleDb"
  license_model = "LICENSE_INCLUDED"
  node_count = "1"
  security_attributes = {
    "oracle-zpr.maxegresscount.mode" = "enforce"
    "oracle-zpr.maxegresscount.value" = "42"
  }
  shape = "VM.Standard2.2"
  ssh_public_keys = ["ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"]
  subnet_id = oci_core_subnet.test_subnet.id
}
