# If you receive a service error indicating: Failed to create Autonomous Transaction Processing manual backup because Object Storage credentials and tenancy URL are not defined. Refer to:
# https://docs.cloud.oracle.com/iaas/Content/Database/Tasks/atpbackingup.htm#creatingbucket
resource "oci_database_autonomous_data_warehouse_backup" "test_autonomous_data_warehouse_backup" {
  # Create the backup only if the autonomous database id is provided and has been configured for backups.
  count = "${var.autonomous_data_warehouse_id == "" ? 0 : 1}"

  #Required
  autonomous_data_warehouse_id = "${var.autonomous_data_warehouse_id}"
  display_name                 = "${var.autonomous_data_warehouse_backup_display_name}"
}

data "oci_database_autonomous_data_warehouse_backups" "autonomous_data_warehouse_backups" {
  #Optional
  autonomous_data_warehouse_id = "${var.autonomous_data_warehouse_id}"

  #compartment_id = "${var.compartment_ocid}"
  display_name = "${var.autonomous_data_warehouse_backup_display_name}"

  #state = "${var.autonomous_data_warehouse_backup_state}"
}

output "autonomous_data_warehouse_backups" {
  value = "${data.oci_database_autonomous_data_warehouse_backups.autonomous_data_warehouse_backups.autonomous_data_warehouse_backups}"
}
