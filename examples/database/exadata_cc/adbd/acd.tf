resource "random_string" "db_unique_name" {
  length = 8
  special = false
  numeric = false
}

resource "oci_database_autonomous_container_database" "autonomous_container_database" {
  autonomous_vm_cluster_id = oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster.id
  db_version = "19.28.0.1.0"
  backup_config {
    backup_destination_details {
      type = "NFS"
      id = oci_database_backup_destination.test_backup_destination.id
      backup_retention_policy_on_terminate = "RETAIN_FOR_72_HOURS"
      is_retention_lock_enabled = false
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
  is_dst_file_update_enabled = false

  #Optional
  // OKV related
  key_store_id = oci_database_key_store.test_key_store.id
  okv_end_point_group_name = "DUMMY_OKV_EPG_GROUP"

  customer_contacts {
    email = "contact1@example.com"
  }

  customer_contacts {
    email = "contact2@example.com"
  }
}

resource "oci_database_key_store" "test_key_store" {
  compartment_id           = var.compartment_ocid
  display_name             = "example-key-store"
  type_details {
    admin_username = "example-username"
    connection_ips = ["192.1.1.1"]
    secret_id      = var.okv_secret
    type           = "ORACLE_KEY_VAULT"
    vault_id       = var.kms_vault_ocid
  }
}

data "oci_database_autonomous_container_database_resource_usage" "test_autonomous_container_database_resource_usages" {
  #Required
  autonomous_container_database_id = oci_database_autonomous_container_database.autonomous_container_database.id
}