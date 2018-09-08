
resource "oci_database_autonomous_data_warehouse_backup" "test_autonomous_data_warehouse_backup" {
	#Required
	autonomous_data_warehouse_id = "${var.autonomous_data_warehouse_id}"
	display_name = "${var.autonomous_data_warehouse_backup_display_name}"
}

data "oci_database_autonomous_data_warehouse_backups" "test_autonomous_data_warehouse_backups" {

	#Optional
	autonomous_data_warehouse_id = "${var.autonomous_data_warehouse_id}"
	#compartment_id = "${var.compartment_ocid}"
	display_name = "${var.autonomous_data_warehouse_backup_display_name}"
	#state = "${var.autonomous_data_warehouse_backup_state}"
}
