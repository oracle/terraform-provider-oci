// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_database_db_home" "test_db_home" {
  display_name = "ExampleExaDbVmDbHome"
  db_system_id = oci_database_exadb_vm_cluster.test_exadb_vm_cluster.id
  db_version   = data.oci_database_db_versions.test_db_versions.db_versions[0].version

  depends_on = [oci_database_db_node_snapshot.test_db_node_snapshot]
}

resource "oci_database_database" "test_db1" {
  database {
    admin_password = var.test_db_password
    db_name        = "TFDB1"
    dynamic "storage_size_details" {
      for_each = local.include_storage_size_details ? [1] : []
      content {
        data_storage_size_in_gb  = 60
        reco_storage_size_in_gbs = 52
      }
    }
  }
  db_home_id = oci_database_db_home.test_db_home.id
  source     = "NONE"
}

resource "oci_database_database_snapshot" "test_db1_snapshot" {
  database_id = oci_database_database.test_db1.id
  name        = "TFDB1Snapshot"
}

resource "oci_database_db_home" "test_snapshot_clone_db_home" {
  display_name = "ExampleExaDbVmSnapshotCloneDbHome"
  db_system_id = oci_database_exadb_vm_cluster.test_exadb_vm_cluster.id
  db_version   = data.oci_database_db_versions.test_db_versions.db_versions[0].version

  depends_on = [oci_database_db_node_snapshot.test_db_node_snapshot]
}

resource "oci_database_database" "test_db1_snapshot_clone" {
  database {
    admin_password        = var.test_db_password
    db_name               = "TFDBSN"
    is_thin_clone         = true
    source_admin_password = var.test_db_password
    source_snapshot_id    = oci_database_database_snapshot.test_db1_snapshot.id
    source_tde_password   = var.test_db_password
    tde_wallet_password   = var.test_db_password
    dynamic "storage_size_details" {
      for_each = local.include_storage_size_details ? [1] : []
      content {
        data_storage_size_in_gb  = 60
        reco_storage_size_in_gbs = 52
      }
    }
  }
  db_home_id = oci_database_db_home.test_snapshot_clone_db_home.id
  source     = "DB_SNAPSHOT"
}

resource "oci_database_db_home" "test_database_clone_db_home" {
  display_name = "ExampleExaDbVmDatabaseCloneDbHome"
  db_system_id = oci_database_exadb_vm_cluster.test_exadb_vm_cluster.id
  db_version   = data.oci_database_db_versions.test_db_versions.db_versions[0].version

  depends_on = [oci_database_db_node_snapshot.test_db_node_snapshot]
}

resource "oci_database_database" "test_db1_database_clone" {
  database {
    admin_password        = var.test_db_password
    db_name               = "TFDBCL"
    is_thin_clone         = true
    source_admin_password = var.test_db_password
    source_database_id    = oci_database_database.test_db1.id
    source_tde_password   = var.test_db_password
    tde_wallet_password   = var.test_db_password
    dynamic "storage_size_details" {
      for_each = local.include_storage_size_details ? [1] : []
      content {
        data_storage_size_in_gb  = 60
        reco_storage_size_in_gbs = 52
      }
    }
  }
  db_home_id = oci_database_db_home.test_database_clone_db_home.id
  source     = "DB_CLONE"
}

# resource "oci_database_database" "test_db2" {
#   database {
#     admin_password = var.test_db_password
#     db_name        = "TFDB2"
#   }
#   db_home_id = oci_database_db_home.test_db_home.id
#   source     = "NONE"
# }

resource "oci_database_pluggable_database" "test_db1_pdb" {
  container_database_id = oci_database_database.test_db1.id
  pdb_name              = "DB1PDB"
  pdb_admin_password    = var.test_db_password
  tde_wallet_password   = var.test_db_password
}

locals {
  # Block storage: runs CDB snapshot/clone examples, skips service-unsupported PDB snapshot/clone.
  # Smart storage: omits CDB storage sizing and enables both CDB and PDB snapshot/clone paths.
  supports_pdb_snapshots       = local.shape_attribute == "SMART_STORAGE"
  include_storage_size_details = local.shape_attribute == "BLOCK_STORAGE"
}

resource "oci_database_pluggable_database_snapshot" "test_db1_pdb_snapshot" {
  count = local.supports_pdb_snapshots ? 1 : 0

  name                  = "DB1PdbSnapshot"
  pluggable_database_id = oci_database_pluggable_database.test_db1_pdb.id
}

resource "oci_database_pluggable_database" "test_db1_local_cloned_pdb" {
  count = local.supports_pdb_snapshots ? 1 : 0

  container_database_id = oci_database_database.test_db1.id
  pdb_name              = "DB1LocalThinClonedPDB"
  pdb_admin_password    = var.test_db_password
  tde_wallet_password   = var.test_db_password
  pdb_creation_type_details {
    creation_type                         = "LOCAL_CLONE_PDB"
    source_pluggable_database_id          = oci_database_pluggable_database.test_db1_pdb.id
    source_pluggable_database_snapshot_id = oci_database_pluggable_database_snapshot.test_db1_pdb_snapshot[0].id
    is_thin_clone                         = true
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
#     source_pluggable_database_snapshot_id = oci_database_pluggable_database_snapshot.test_db1_pdb_snapshot[0].id
#     is_thin_clone = true
#   }
# }

data "oci_database_database_snapshots" "test_db1_snapshots" {
  cluster_id  = oci_database_exadb_vm_cluster.test_exadb_vm_cluster.id
  database_id = oci_database_database.test_db1.id
  name        = oci_database_database_snapshot.test_db1_snapshot.name
  state       = "AVAILABLE"
}

data "oci_database_database_snapshot" "test_db1_snapshot" {
  database_snapshot_id = oci_database_database_snapshot.test_db1_snapshot.id
}

data "oci_database_pluggable_database_snapshots" "test_pdb_snapshots" {
  count = local.supports_pdb_snapshots ? 1 : 0

  compartment_id = var.compartment_ocid
  cluster_id     = oci_database_exadb_vm_cluster.test_exadb_vm_cluster.id
}

data "oci_database_pluggable_database_snapshot" "test_db1_pdb_snapshot" {
  count = local.supports_pdb_snapshots ? 1 : 0

  pluggable_database_snapshot_id = oci_database_pluggable_database_snapshot.test_db1_pdb_snapshot[0].id
}

data "oci_database_pluggable_databases" "test_pdbs" {
  compartment_id = var.compartment_ocid
  state          = "AVAILABLE"
}

data "oci_database_pluggable_database" "test_db1_local_cloned_pdb" {
  count = local.supports_pdb_snapshots ? 1 : 0

  pluggable_database_id = oci_database_pluggable_database.test_db1_local_cloned_pdb[0].id
}

data "oci_database_db_homes" "test_db_homes" {
  compartment_id = var.compartment_ocid
  vm_cluster_id  = oci_database_exadb_vm_cluster.test_exadb_vm_cluster.id
  state          = "AVAILABLE"
}

data "oci_database_db_home" "test_db_home" {
  db_home_id = oci_database_db_home.test_db_home.id
}

data "oci_database_databases" "test_databases" {
  compartment_id = var.compartment_ocid
  db_home_id     = oci_database_db_home.test_db_home.id
}

data "oci_database_database" "test_db1" {
  database_id = oci_database_database.test_db1.id
}

data "oci_database_database" "test_db1_snapshot_clone" {
  database_id = oci_database_database.test_db1_snapshot_clone.id
}

data "oci_database_database" "test_db1_database_clone" {
  database_id = oci_database_database.test_db1_database_clone.id
}
