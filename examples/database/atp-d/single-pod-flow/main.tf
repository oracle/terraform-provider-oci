// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_database_autonomous_container_database" "test_autonomous_container_database" {
  #Required
  cloud_autonomous_vm_cluster_id       = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id
  display_name                         = "container-database"
  patch_model                          = "RELEASE_UPDATES"
  db_version                           = var.acd_db_version
  db_name                              = "NAME${random_string.db_unique_name.result}"

  #Optional
  backup_config {
    #Optional
    recovery_window_in_days = var.autonomous_container_database_backup_config_recovery_window_in_days
    backup_destination_details {
      type = "OBJECT_STORE"
      backup_retention_policy_on_terminate = "RETAIN_FOR_72_HOURS"
      is_retention_lock_enabled = false
    }
  }

  #Optional
  db_split_threshold           = 12
  vm_failover_reservation      = 25
  distribution_affinity        = "MINIMUM_DISTRIBUTION"
  net_services_architecture    = "DEDICATED"

  compartment_id               = var.compartment_ocid
  freeform_tags                = var.autonomous_database_freeform_tags
  service_level_agreement_type = "STANDARD"

  maintenance_window_details {
    preference = "CUSTOM_PREFERENCE"

    days_of_week {
      name = "MONDAY"
    }

    hours_of_day = ["4"]

    months {
      name = "JANUARY"
    }

    months {
      name = "APRIL"
    }

    months {
      name = "JULY"
    }

    months {
      name = "OCTOBER"
    }

    weeks_of_month = ["2"]
  }
  rotate_key_trigger = "true"
  version_preference = "LATEST_RELEASE_UPDATE"
  is_dst_file_update_enabled = false

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

resource "oci_database_autonomous_container_database" "test_autonomous_container_database_no_backup" {
  #Required
  cloud_autonomous_vm_cluster_id       = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id
  display_name                         = "container-database-no-backup"
  patch_model                          = "RELEASE_UPDATES"
  db_version                           = var.acd_db_version
  db_name                              = "NAME${random_string.db_unique_name_2.result}"

  #Optional
  backup_config {
    recovery_window_in_days = 0
  }

  maintenance_window_details {
      preference = "NO_PREFERENCE"
  }

  compartment_id               = var.compartment_ocid
  db_split_threshold           = 12
  vm_failover_reservation      = 25
  service_level_agreement_type = "STANDARD"
  distribution_affinity        = "MINIMUM_DISTRIBUTION"
  net_services_architecture    = "DEDICATED"
  version_preference = "LATEST_RELEASE_UPDATE"
  is_dst_file_update_enabled = false
}


# # First character must be a letter.
# resource "random_string" "db_unique_name_letter" {
#   length  = 1
#   upper   = true
#   lower   = false
#   special = false
#   numeric = false
# }

# Next three can be alnum or underscore.
resource "random_string" "db_unique_name" {
  length  = 3
  upper   = true
  lower   = false
  numeric  = true
  special = false
}

resource "random_string" "db_unique_name_2" {
  length  = 3
  upper   = true
  lower   = false
  numeric  = true
  special = false
}

resource "random_string" "autonomous_database_admin_password" {
  length      = 16
  min_numeric = 1
  min_lower   = 1
  min_upper   = 1
  min_special = 1
}

resource "oci_database_autonomous_database" "test_autonomous_database" {
  #Required
  admin_password           = random_string.autonomous_database_admin_password.result
  compartment_id           = var.compartment_ocid
  compute_count            = 16
  data_storage_size_in_tbs = "1"
  db_name                  = "atpdb1"

  #Optional
  autonomous_container_database_id = oci_database_autonomous_container_database.test_autonomous_container_database.id
  db_workload                      = "OLTP"
  display_name                     = "example_autonomous_db-1"
  freeform_tags                    = var.autonomous_database_freeform_tags
  is_dedicated                     = "true"
  rotate_key_trigger               = "true"
  in_memory_percentage             = 50
  compute_model                    = "ECPU"
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

