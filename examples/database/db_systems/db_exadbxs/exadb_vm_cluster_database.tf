// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_database_db_home" "test_db_home" {
  display_name = "ExampleExaDbVmDbHome"
  db_system_id = oci_database_exadb_vm_cluster.test_exadb_vm_cluster.id
  db_version   = "23.4.0.24.05"
}

resource "oci_database_database" "test_db1" {
  database {
    admin_password = var.test_db_password
    db_name        = "TFDB1"
  }
  db_home_id = oci_database_db_home.test_db_home.id
  source     = "NONE"
}

resource "oci_database_database" "test_db2" {
  database {
    admin_password = var.test_db_password
    db_name        = "TFDB2"
  }
  db_home_id = oci_database_db_home.test_db_home.id
  source     = "NONE"
}

resource "oci_database_pluggable_database" "test_db1_pdb" {
  container_database_id = oci_database_database.test_db1.id
  pdb_name              = "DB1PDB"
  pdb_admin_password    = var.test_db_password
  tde_wallet_password   = var.test_db_password
}

resource "oci_database_pluggable_database" "test_db1_local_cloned_pdb" {
  container_database_id = oci_database_database.test_db1.id
  pdb_name = "DB1LocalThinClonedPDB"
  pdb_admin_password = var.test_db_password
  tde_wallet_password = var.test_db_password
  pdb_creation_type_details {
    creation_type = "LOCAL_CLONE_PDB"
    source_pluggable_database_id = oci_database_pluggable_database.test_db1_pdb.id
    is_thin_clone = true
  }
}

# resource "oci_database_pluggable_database" "test_db2_remote_cloned_pdb" {
#   container_database_id = oci_database_database.test_db2.id
#   pdb_name = "DB2RemoteThinClonedPDB"
#   pdb_admin_password = var.test_db_password
#   tde_wallet_password = var.test_db_password
#   pdb_creation_type_details {
#     creation_type = "REMOTE_CLONE_PDB"
#     source_container_database_admin_password = var.test_db_password
#     source_pluggable_database_id = oci_database_pluggable_database.test_db1_pdb.id
#     is_thin_clone = true
#   }
# }

data "oci_database_pluggable_databases" "test_pdbs" {
  compartment_id        = var.compartment_ocid
  state = "AVAILABLE"
}

data "oci_database_pluggable_database" "test_db1_local_cloned_pdb" {
  pluggable_database_id = oci_database_pluggable_database.test_db1_local_cloned_pdb.id
}