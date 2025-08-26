# $Header$
#
# Copyright (c) 2024, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      main.tf - Shepherd Main file
#
#    USAGE
#      Use the following path for Example Test & Backward-Compatibility-Test: database/dataguard/vm_shape/vm_std_x86
#    NOTES
#      Terraform Example: TestDatabaseDataGuardAssociationResourceVmStdx86_basic
#    FILES
#
#    DESCRIPTION
#
#    MODIFIED   MM/DD/YY
#    aavadhan   08/18/2025 - Created


resource "oci_database_db_system" "test_db_system" {
  availability_domain = data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name
  compartment_id = var.compartment_id
  compute_model = "ECPU"
  compute_count = "4"
  data_storage_size_in_gb = "256"
  database_edition = "ENTERPRISE_EDITION"
  db_home {
    database {
      admin_password = "BEstrO0ng_#11"
      db_name = "tfDbName"
    }
    db_version = "19.0.0.0"
    display_name = "TFTestDbHome1"
  }
  db_system_options {
    storage_management = "LVM"
  }
  disk_redundancy = "NORMAL"
  display_name = "tfExampleDbSystemDataguardAssociationPrimary"
  domain = oci_core_subnet.test_subnet.subnet_domain_name
  hostname = "myOracleDB"
  license_model = "LICENSE_INCLUDED"
  node_count = "1"
  shape = "VM.Standard.x86"
  ssh_public_keys = [var.ssh_public_key]
  subnet_id = oci_core_subnet.test_subnet.id
}

resource "oci_database_data_guard_association" "test_data_guard_association" {
  availability_domain = data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name
  backup_network_nsg_ids = [oci_core_network_security_group.test_network_security_group.id]
  compute_model = "ECPU"
  compute_count = "4"
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
  shape = "VM.Standard.x86"
  storage_volume_performance_mode = "BALANCED"
  subnet_id = oci_core_subnet.test_subnet.id
  time_zone = "US/Pacific"
  transport_type = "ASYNC"
}
