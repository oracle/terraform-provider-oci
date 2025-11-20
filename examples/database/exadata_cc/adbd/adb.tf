resource "random_string" "autonomous_database_admin_password" {
  length      = 16
  min_numeric = 1
  min_lower   = 1
  min_upper   = 1
  min_special = 1
}

resource "oci_database_autonomous_database" "test_autonomous_database" {
  #Required
  admin_password           = random_string.autonomous_database_admin_password.result
  compartment_id           = var.compartment_ocid
  compute_count            = 8
  data_storage_size_in_tbs = "1"
  db_name                  = "atpdb2"

  #Optional
  autonomous_container_database_id = oci_database_autonomous_container_database.autonomous_container_database.id
  db_workload                      = "OLTP"
  display_name                     = "exacc_tf-adb"
  is_dedicated                     = "true"
}

resource "oci_database_autonomous_database_backup" "test_autonomous_database_backup" {
  autonomous_database_id = oci_database_autonomous_database.test_autonomous_database.id
  display_name = "DbBackupName"
  is_long_term_backup = "true"
  retention_period_in_days = 90
  backup_destination_details {
    type = "NFS"
    id = oci_database_backup_destination.test_backup_destination.id
  }
}

resource "oci_database_backup_destination" "test_backup_destination" {
  compartment_id = var.compartment_ocid
  display_name = "NFS1"
  type = "NFS"
  mount_type_details {
    mount_type = "AUTOMATED_MOUNT"
    nfs_server = ["98.56.65.88", "101.67.98.66"]
    nfs_server_export = "/mount/export"
  }

}

resource "oci_database_autonomous_database" "test_autonomous_database_from_backup" {
  admin_password           = random_string.autonomous_database_admin_password.result
  compartment_id           = var.compartment_ocid
  compute_count            = 8
  data_storage_size_in_tbs = "1"
  compute_model = "ECPU"


  autonomous_container_database_id = oci_database_autonomous_container_database.autonomous_container_database.id
  db_workload                      = "OLTP"
  display_name                     = "partial-adb"
  is_dedicated                     = "true"

  db_name                       = "cloneadb"
  clone_type                    = "FULL"
  source                        = "BACKUP_FROM_ID"
  autonomous_database_backup_id = oci_database_autonomous_database_backup.test_autonomous_database_backup.id
}
