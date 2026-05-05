# This example clones an Autonomous Container Database from an existing ACD
# backup OCID. It reuses the source ACD, source ADB, backup destination, and
# peer AVM cluster from the other files in this example directory.
#
# Execution order:
# 1. The base example creates the Exadata infrastructure, VM cluster networks,
#    source AVM cluster, peer AVM cluster, source ACD, source ADB, and NFS
#    backup destination.
# 2. This data source lists active local ACD backups for the source ACD after
#    the source ADB exists.
# 3. The clone resource uses the first returned backup OCID to create a full
#    ACD clone on the peer AVM cluster.
# 4. The final data source lists ADBs inside the cloned ACD for validation.

data "oci_database_autonomous_container_database_backups" "acd_clone_from_backup_backups" {
  autonomous_container_database_id = oci_database_autonomous_container_database.autonomous_container_database.id
  infrastructure_type              = "CLOUD_AT_CUSTOMER"
  is_remote                        = false
  state                            = "ACTIVE"

  depends_on = [
    oci_database_autonomous_database.test_autonomous_database
  ]
}

resource "oci_database_autonomous_container_database" "acd_clone_from_backup" {
  source                                  = "BACKUP_FROM_ID"
  autonomous_container_database_backup_id = data.oci_database_autonomous_container_database_backups.acd_clone_from_backup_backups.autonomous_container_database_backup_collection.0.items.0.id
  autonomous_vm_cluster_id                = oci_database_autonomous_vm_cluster.autonomous_vm_cluster_2.id
  clone_type                              = "FULL"
  compartment_id                          = var.compartment_ocid
  display_name                            = "acdCloneFromBackup"
  patch_model                             = "RELEASE_UPDATES"
  service_level_agreement_type            = "STANDARD"

  backup_config {
    backup_destination_details {
      type                                 = "NFS"
      id                                   = oci_database_backup_destination.test_backup_destination.id
      backup_retention_policy_on_terminate = "RETAIN_FOR_72_HOURS"
      is_retention_lock_enabled            = false
    }

    recovery_window_in_days = "10"
  }

  lifecycle {
    ignore_changes = [
      source,
      clone_type,
      autonomous_container_database_backup_id
    ]
  }
}

data "oci_database_autonomous_databases" "acd_clone_from_backup_databases" {
  autonomous_container_database_id = oci_database_autonomous_container_database.acd_clone_from_backup.id
  compartment_id                   = var.compartment_ocid
  display_name                     = oci_database_autonomous_database.test_autonomous_database.display_name

  depends_on = [
    oci_database_autonomous_container_database.acd_clone_from_backup
  ]
}
