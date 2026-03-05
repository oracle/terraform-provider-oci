# $Header$
#
# Copyright (c) 2024, Oracle and/or its affiliates. All rights reserved.
#    NAME
#      main.tf - Shepherd Main file
#
#    USAGE
#      Use the following path for the Example & Backward Compatibility tests: database/db_systems/db_vm/multiple_standby
#    NOTES
#      Terraform Example: TestResourceDbSystemDataGuardAssociation


resource "oci_database_db_system" "test_db_system" {
  availability_domain = data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name
  compartment_id = var.compartment_id
  cpu_core_count = "2"
  data_storage_size_in_gb = "256"
  database_edition = "ENTERPRISE_EDITION"
  db_home {
    database {
      admin_password = "BEstrO0ng_#11"
      db_name = "tfDbP"
    }
    db_version = "19.0.0.0"
    display_name = "TFTestDbHomePrimary"
  }
  db_system_options {
    storage_management = "LVM"
  }
  disk_redundancy = "NORMAL"
  display_name = "tfExampleDbSystemDataguardAssociationPrimary"
  domain = oci_core_subnet.test_subnet.subnet_domain_name
  hostname = "myOracleDBPrm"
  license_model = "LICENSE_INCLUDED"
  node_count = "1"
  shape = "VM.Standard.E4.Flex"
  ssh_public_keys = [var.ssh_public_key]
  subnet_id = oci_core_subnet.test_subnet.id
}

resource "oci_database_db_system" "test_multiple_standby_1" {
  availability_domain = data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name
  compartment_id = var.compartment_id
  cpu_core_count = "2"
  data_storage_size_in_gb = "256"
  database_edition = "ENTERPRISE_EDITION"
  source = "DATAGUARD"
  primary_db_system_id = oci_database_db_system.test_db_system.id
  db_home {
    database {
      admin_password = "BEstrO0ng_#11"
      db_name = "tfDbS"
      protection_mode = "MAXIMUM_PERFORMANCE"
      transport_type = "ASYNC"
    }
    db_version = "19.0.0.0"
    display_name = "TFTestDbHomeStandby1"
  }
  db_system_options {
    storage_management = "LVM"
  }
  disk_redundancy = "NORMAL"
  display_name = "tfExampleDbSystemDataguardAssociationStandby1"
  domain = oci_core_subnet.test_subnet.subnet_domain_name
  hostname = "myOracleDBStd1"
  license_model = "LICENSE_INCLUDED"
  node_count = "1"
  shape = "VM.Standard.E4.Flex"
  ssh_public_keys = [var.ssh_public_key]
  subnet_id = oci_core_subnet.test_subnet.id
}

resource "oci_database_db_system" "test_multiple_standby_2" {
  depends_on = [oci_database_db_system.test_multiple_standby_1]
  availability_domain = data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name
  compartment_id = var.compartment_id
  cpu_core_count = "2"
  data_storage_size_in_gb = "256"
  database_edition = "ENTERPRISE_EDITION"
  source = "DATAGUARD"
  primary_db_system_id = oci_database_db_system.test_db_system.id
  db_home {
    database {
      admin_password = "BEstrO0ng_#11"
      db_name = "tfDbSm"
      protection_mode = "MAXIMUM_PERFORMANCE"
      transport_type = "ASYNC"
    }
    db_version = "19.0.0.0"
    display_name = "TFTestDbHomeStandby2"
  }
  db_system_options {
    storage_management = "LVM"
  }
  disk_redundancy = "NORMAL"
  display_name = "tfExampleDbSystemDataguardAssociationStandby2"
  domain = oci_core_subnet.test_subnet.subnet_domain_name
  hostname = "myOracleDBStd2"
  license_model = "LICENSE_INCLUDED"
  node_count = "1"
  shape = "VM.Standard.E4.Flex"
  ssh_public_keys = [var.ssh_public_key]
  subnet_id = oci_core_subnet.test_subnet.id
}