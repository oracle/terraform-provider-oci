// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

# This example clones an Autonomous Container Database from a backup timestamp
# in public cloud. It creates a third Cloud AVM cluster on the existing Cloud
# Exadata Infrastructure so this partial timestamp clone can run in parallel
# with the full clone-from-backup example.
#
# Execution order:
# 1. The base example creates the VCN, subnet, source Cloud Exadata
#    Infrastructure, source Cloud AVM cluster, source ACD, and source ADB.
# 2. This file creates a third Cloud AVM cluster on the same Cloud Exadata
#    Infrastructure. This third AVM is the timestamp clone target.
# 3. This data source lists active local public-cloud ACD backups for the source
#    ACD after the source ADB exists.
# 4. The clone resource uses the first returned backup end timestamp to create a
#    partial ACD clone on the third Cloud AVM cluster.
# 5. The final data sources list cloned ADBs inside the timestamp-cloned ACD.

resource "oci_database_cloud_autonomous_vm_cluster" "timestamp_peer_cloud_autonomous_vm_cluster" {
  cloud_exadata_infrastructure_id = oci_database_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure.id
  compartment_id                  = var.compartment_ocid
  display_name                    = "TimestampPeerCloudAutonomousVmCluster"
  freeform_tags                   = var.autonomous_database_freeform_tags
  license_model                   = "LICENSE_INCLUDED"
  subnet_id                       = oci_core_subnet.exadata_subnet.id
  compute_model                   = "ECPU"

  lifecycle {
    ignore_changes = [
      autonomous_data_storage_size_in_tbs,
      db_servers,
    ]
  }
}

data "oci_database_autonomous_container_database_backups" "acd_clone_from_backup_timestamp_backups" {
  autonomous_container_database_id = oci_database_autonomous_container_database.test_autonomous_container_database.id
//  compartment_id                   = var.compartment_ocid
  display_name                     = "Automatic Backup"
  infrastructure_type              = "CLOUD"
  is_remote                        = false
  state                            = "ACTIVE"

  depends_on = [
    oci_database_autonomous_database.test_autonomous_database
  ]
}

resource "oci_database_autonomous_container_database" "acd_clone_from_backup_timestamp" {
  source                                  = "BACKUP_FROM_TIMESTAMP"
  source_autonomous_container_database_id = oci_database_autonomous_container_database.test_autonomous_container_database.id
  time_stamp_to_use_for_cloning           = data.oci_database_autonomous_container_database_backups.acd_clone_from_backup_timestamp_backups.autonomous_container_database_backup_collection.0.items.0.time_ended
  autonomous_databases_to_clone           = [oci_database_autonomous_database.test_autonomous_database.display_name]
  cloud_autonomous_vm_cluster_id          = oci_database_cloud_autonomous_vm_cluster.timestamp_peer_cloud_autonomous_vm_cluster.id
  clone_type                              = "PARTIAL"
  compartment_id                          = var.compartment_ocid
  display_name                            = "acdCloneFromBackupTimestamp"
  patch_model                             = "RELEASE_UPDATES"
  service_level_agreement_type            = "STANDARD"

  backup_config {
    backup_destination_details {
      type                                 = "OBJECT_STORE"
      backup_retention_policy_on_terminate = "RETAIN_FOR_72_HOURS"
      is_retention_lock_enabled            = false
    }

    recovery_window_in_days = var.autonomous_container_database_backup_config_recovery_window_in_days
  }

  lifecycle {
    ignore_changes = [
      source,
      clone_type,
      source_autonomous_container_database_id,
      time_stamp_to_use_for_cloning,
      autonomous_databases_to_clone
    ]
  }
}

data "oci_database_autonomous_databases" "acd_clone_from_backup_timestamp_databases" {
  autonomous_container_database_id = oci_database_autonomous_container_database.acd_clone_from_backup_timestamp.id
  compartment_id                   = var.compartment_ocid
  display_name                     = oci_database_autonomous_database.test_autonomous_database.display_name

  depends_on = [
    oci_database_autonomous_container_database.acd_clone_from_backup_timestamp
  ]
}

data "oci_database_autonomous_databases" "all_acd_clone_from_backup_timestamp_databases" {
  autonomous_container_database_id = oci_database_autonomous_container_database.acd_clone_from_backup_timestamp.id
  compartment_id                   = var.compartment_ocid

  depends_on = [
    oci_database_autonomous_container_database.acd_clone_from_backup_timestamp
  ]
}
