resource "oci_database_autonomous_data_warehouse" "test_autonomous_data_warehouse" {
	#Required
	admin_password = "${var.autonomous_data_warehouse_admin_password}"
	compartment_id = "${var.compartment_ocid}"
	cpu_core_count = "${var.autonomous_data_warehouse_cpu_core_count}"
	data_storage_size_in_tbs = "${var.autonomous_data_warehouse_data_storage_size_in_tbs}"
	db_name = "${var.autonomous_data_warehouse_db_name}"

	#Optional
	display_name = "${var.autonomous_data_warehouse_display_name}"
	license_model = "${var.autonomous_data_warehouse_license_model}"
}

data "oci_database_autonomous_data_warehouses" "test_autonomous_data_warehouses" {
	#Required
	compartment_id = "${var.compartment_ocid}"

	#Optional
	display_name = "${var.autonomous_data_warehouse_display_name}"
	#state = "${var.autonomous_data_warehouse_state}"
}

output "autonomous_data_warehouses_name" {
    value = "${data.oci_database_autonomous_data_warehouses.test_autonomous_data_warehouses.display_name}"
}