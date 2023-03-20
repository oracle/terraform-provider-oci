resource "random_string" "db_unique_name" {
  length = 8
  special = false
  number = false
}

resource "oci_database_autonomous_container_database" "autonomous_container_database" {
  autonomous_vm_cluster_id = oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster.id
  db_version = "19.18.0.1.0"
  backup_config {
    backup_destination_details {
      type = "LOCAL"
    }
    recovery_window_in_days = "7"
  }
  compartment_id = var.compartment_ocid
  db_unique_name = random_string.db_unique_name.result
  display_name = "ACD-TFTest"
  freeform_tags = {
    "Department" = "Finance"
  }
  maintenance_window_details {
    preference = "NO_PREFERENCE"
  }
  patch_model = "RELEASE_UPDATES"
  service_level_agreement_type = "STANDARD"
  version_preference = "LATEST_RELEASE_UPDATE"
}