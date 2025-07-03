data "oci_identity_availability_domain" "ad" {
  compartment_id = var.compartment_ocid
  ad_number      = 1
}

data "oci_database_maintenance_run" "test_maintenance_run" {
  maintenance_run_id = oci_database_maintenance_run.test_maintenance_run.id
}

data "oci_database_maintenance_runs" "test_maintenance_runs" {
  compartment_id = var.compartment_ocid
  target_resource_id = oci_database_autonomous_container_database.test_autonomous_container_database.id
}

data "oci_database_maintenance_run_history" "test_maintenance_run_history" {
  maintenance_run_history_id = var.maintenance_run_history_id
}

data "oci_database_maintenance_run_histories" "test_maintenance_run_histories" {
  compartment_id = var.tenancy_ocid
}