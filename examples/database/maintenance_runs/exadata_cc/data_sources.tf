data "oci_identity_availability_domain" "ad" {
  compartment_id = var.compartment_ocid
  ad_number      = 1
}

data "oci_database_maintenance_run_history" "test_maintenance_run_history" {
  maintenance_run_history_id = var.maintenance_run_history_id
}

data "oci_database_maintenance_run_histories" "test_maintenance_run_histories" {
  compartment_id = var.tenancy_ocid
}