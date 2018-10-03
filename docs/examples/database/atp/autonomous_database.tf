resource "random_string" "autonomous_database_admin_password" {
  length  = 16
  special = true
}

resource "oci_database_autonomous_database" "autonomous_database" {
  #Required
  admin_password           = "${random_string.autonomous_database_admin_password.result}"
  compartment_id           = "${var.compartment_ocid}"
  cpu_core_count           = "${var.autonomous_database_cpu_core_count}"
  data_storage_size_in_tbs = "${var.autonomous_database_data_storage_size_in_tbs}"
  db_name                  = "${var.autonomous_database_db_name}"

  #Optional
  display_name  = "${var.autonomous_database_display_name}"
  freeform_tags = "${var.autonomous_database_freeform_tags}"
  license_model = "${var.autonomous_database_license_model}"
}

data "oci_database_autonomous_databases" "autonomous_databases" {
  #Required
  compartment_id = "${var.compartment_ocid}"

  #Optional
  display_name = "${var.autonomous_database_display_name}"
}

output "autonomous_database_admin_password" {
  value = "${random_string.autonomous_database_admin_password.result}"
}

output "autonomous_databases" {
  value = "${data.oci_database_autonomous_databases.autonomous_databases.autonomous_databases}"
}

output "parallel_connection_string" {
  value = ["${lookup(oci_database_autonomous_database.autonomous_database.connection_strings.0.all_connection_strings, "PARALLEL", "Unavailable")}"]
}
