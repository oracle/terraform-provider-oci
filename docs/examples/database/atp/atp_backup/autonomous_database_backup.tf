# If you receive a service error indicating backup cannot be created. Refer to:
# https://docs.cloud.oracle.com/iaas/Content/Database/Tasks/adwbackingup.htm#CreatingaBuckettoStoreManualBackups
resource "oci_database_autonomous_database_backup" "autonomous_database_backup" {
  # Create the backup only if the autonomous database id is provided and has been configured for backups.
  count = "${var.autonomous_database_id == "" ? 0 : 1}"

  #Required
  autonomous_database_id = "${var.autonomous_database_id}"
  display_name           = "${var.autonomous_database_backup_display_name}"
}

data "oci_database_autonomous_database_backups" "test_autonomous_database_backups" {
  #Optional
  compartment_id = "${var.compartment_ocid}"
  display_name   = "${var.autonomous_database_backup_display_name}"
}

output "autonomous_database_backups" {
  value = "${data.oci_database_autonomous_database_backups.test_autonomous_database_backups.autonomous_database_backups}"
}
