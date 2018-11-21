resource "random_string" "autonomous_data_warehouse_admin_password" {
  length  = 16
  special = true
}

resource "oci_database_autonomous_data_warehouse" "autonomous_data_warehouse" {
  #Required
  admin_password           = "${random_string.autonomous_data_warehouse_admin_password.result}"
  compartment_id           = "${var.compartment_ocid}"
  cpu_core_count           = "${var.autonomous_data_warehouse_cpu_core_count}"
  data_storage_size_in_tbs = "${var.autonomous_data_warehouse_data_storage_size_in_tbs}"
  db_name                  = "${var.autonomous_data_warehouse_db_name}"

  #Optional
  display_name  = "${var.autonomous_data_warehouse_display_name}"
  license_model = "${var.autonomous_data_warehouse_license_model}"
}

data "oci_database_autonomous_data_warehouses" "autonomous_data_warehouses" {
  #Required
  compartment_id = "${var.compartment_ocid}"

  #Optional
  display_name = "${oci_database_autonomous_data_warehouse.autonomous_data_warehouse.display_name}"

  #state = "${var.autonomous_data_warehouse_state}"
}

output "autonomous_data_warehouse_admin_password" {
  value = "${random_string.autonomous_data_warehouse_admin_password.result}"
}

output "autonomous_data_warehouses" {
  value = "${data.oci_database_autonomous_data_warehouses.autonomous_data_warehouses.autonomous_data_warehouses}"
}

output "parallel_connection_string" {
  value = ["${lookup(oci_database_autonomous_data_warehouse.autonomous_data_warehouse.connection_strings.0.all_connection_strings, "PARALLEL", "Unavailable")}"]
}
