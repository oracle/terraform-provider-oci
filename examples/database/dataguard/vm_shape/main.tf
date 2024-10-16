# $Header$
#
# Copyright (c) 2024, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      main.tf - Shepherd Main file
#
#    USAGE
#      Use the following path for Example Test & Backward-Compatibility-Test: database/dataguard/vm_shape
#    NOTES
#      Terraform Example: TestDatabaseDataGuardAssociationResource_basic
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
  cidr_block = "10.0.0.0/16"
  compartment_id = var.compartment_id
  display_name = "displayName"
  dns_label = "dnslabel"
}

resource "oci_core_network_security_group" "test_network_security_group" {
  compartment_id = var.compartment_id
  display_name = "displayName2"
  freeform_tags = {
    "Department" = "Accounting"
  }

  vcn_id = oci_core_vcn.test_vcn.id
}

resource "oci_core_service_gateway" "test_service_gateway" {
  compartment_id = var.compartment_id
  display_name = "test_service_gateway"
  services {
    service_id = data.oci_core_services.test_services.services.0.id
  }
  vcn_id = oci_core_vcn.test_vcn.id
}

resource "oci_core_default_route_table" "test_vcn_default_route_table" {
  manage_default_resource_id = oci_core_vcn.test_vcn.default_route_table_id
  route_rules {
    description = "Internal traffic for OCI Services"
    destination = data.oci_core_services.test_services.services[0].cidr_block
    destination_type = "SERVICE_CIDR_BLOCK"
    network_entity_id = oci_core_service_gateway.test_service_gateway.id
  }
}

resource "oci_core_route_table" "test_route_table" {
  compartment_id = var.compartment_id
  display_name = "test_subnet_rt"
  route_rules {
    description = "Internal traffic for OCI Services"
    destination = data.oci_core_services.test_services.services[0].cidr_block
    destination_type = "SERVICE_CIDR_BLOCK"
    network_entity_id = oci_core_service_gateway.test_service_gateway.id
  }
  vcn_id = oci_core_vcn.test_vcn.id
}

resource "oci_core_security_list" "test_security_list" {
  compartment_id = var.compartment_id
  display_name = "test_security_list"
  egress_security_rules {
    destination = "0.0.0.0/0"
    protocol = "6"
  }
  ingress_security_rules {
    protocol = "6"
    source = "0.0.0.0/0"
  }
  vcn_id = oci_core_vcn.test_vcn.id
}

resource "oci_core_subnet" "test_subnet" {
  cidr_block = "10.0.2.0/24"
  compartment_id = var.compartment_id
  dhcp_options_id = oci_core_vcn.test_vcn.default_dhcp_options_id
  display_name = "test_subnet"
  dns_label = "tftestsubnet"
  prohibit_public_ip_on_vnic = "true"
  route_table_id = oci_core_route_table.test_route_table.id
  security_list_ids = [oci_core_security_list.test_security_list.id]
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
      db_name = "tfDbName"
    }
    db_version = "12.1.0.2"
    display_name = "TFTestDbHome1"
  }
  disk_redundancy = "NORMAL"
  display_name = "tfExampleDbSystemDataguardAssociationPrimary"
  domain = oci_core_subnet.test_subnet.subnet_domain_name
  hostname = "myOracleDB"
  license_model = "LICENSE_INCLUDED"
  node_count = "1"
  shape = "VM.Standard2.2"
  ssh_public_keys = ["ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"]
  subnet_id = oci_core_subnet.test_subnet.id
}

resource "oci_database_data_guard_association" "test_data_guard_association" {
  availability_domain = data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name
  backup_network_nsg_ids = [oci_core_network_security_group.test_network_security_group.id]
  cpu_core_count = "10"
  creation_type = "NewDbSystem"
  data_collection_options {
    is_diagnostics_events_enabled = "false"
    is_health_monitoring_enabled = "false"
    is_incident_logs_enabled = "false"
  }
  database_admin_password = "BEstrO0ng_#11"
  database_defined_tags = map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "databaseDefinedTags1")
  database_freeform_tags = {
    "databaseFreeformTagsK" = "databaseFreeformTagsV"
  }
  database_id = data.oci_database_databases.db.databases.0.id
  db_system_defined_tags = map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "dbSystemDefinedTags1")
  db_system_freeform_tags = {
    "dbSystemFreeformTagsK" = "dbSystemFreeformTagsV"
  }
  db_system_security_attributes = {
    "oracle-zpr.maxegresscount.mode" = "enforce"
    "oracle-zpr.maxegresscount.value" = "42"
  }
  delete_standby_db_home_on_delete = "true"
  depends_on = [oci_database_db_system.test_db_system]
  display_name = "tfExampleDbSystemDataguardAssociationStandby"
  domain = "tftestsubnet.dnslabel.oraclevcn.com"
  fault_domains = ["FAULT-DOMAIN-3"]
  hostname = "hostname"
  is_active_data_guard_enabled = "false"
  license_model = "BRING_YOUR_OWN_LICENSE"
  node_count = "1"
  nsg_ids = [oci_core_network_security_group.test_network_security_group.id]
  private_ip = "10.0.2.223"
  protection_mode = "MAXIMUM_PERFORMANCE"
  shape = "VM.Standard2.2"
  storage_volume_performance_mode = "BALANCED"
  subnet_id = oci_core_subnet.test_subnet.id
  time_zone = "US/Pacific"
  transport_type = "ASYNC"
}
