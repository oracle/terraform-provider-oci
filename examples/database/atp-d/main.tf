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

resource "oci_database_autonomous_container_database" "test_autonomous_container_database_for_add_standby" {
  #Required
  cloud_autonomous_vm_cluster_id       = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id
  display_name                         = "add-standby-container-database"
  patch_model                          = "RELEASE_UPDATES"
  db_version                           = var.acd_db_version
  db_name                              = "ACD${random_string.db_unique_name.result}"

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

  version_preference = "LATEST_RELEASE_UPDATE"
  is_dst_file_update_enabled = false

  lifecycle {
    ignore_changes = [
      service_level_agreement_type,
      protection_mode,
      peer_autonomous_container_database_display_name,
      peer_autonomous_exadata_infrastructure_id,
      peer_autonomous_vm_cluster_id,
      peer_cloud_autonomous_vm_cluster_id,
      peer_db_unique_name,
      peer_autonomous_container_database_backup_config,
    ]
  }
}

resource "oci_database_autonomous_database_software_image" "autonomous_database_software_image" {
  compartment_id = var.compartment_ocid
  display_name = "ADSI-TFTest"
  image_shape_family = "EXADATA_SHAPE"
  source_cdb_id = oci_database_autonomous_container_database.test_autonomous_container_database.id
  freeform_tags = {
      "Department" = "Finance"
    }
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

# # Helper to build valid DB identifier.
# locals {
#   db_unique_name = "${random_string.db_unique_name_letter.result}${random_string.db_unique_name_rest.result}"
# }

resource "oci_database_autonomous_container_database" "autonomous_container_database_from_adsi" {
  cloud_autonomous_vm_cluster_id       = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id
  database_software_image_id = oci_database_autonomous_database_software_image.autonomous_database_software_image.id
  backup_config {
    backup_destination_details {
      type = "OBJECT_STORE"
      backup_retention_policy_on_terminate = "RETAIN_FOR_72_HOURS"
      is_retention_lock_enabled = false
    }
    recovery_window_in_days = "7"
  }
  compartment_id = var.compartment_ocid
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

# resource "oci_database_autonomous_database" "test_autonomous_database_character_set_support" {
#   #Required
#   admin_password           = random_string.autonomous_database_admin_password.result
#   compartment_id           = var.compartment_ocid
#   compute_count            = "2"
#   data_storage_size_in_tbs = "1"
#   db_name                  = "atpdb2"
#
#   #Optional
#   autonomous_container_database_id = oci_database_autonomous_container_database.test_autonomous_container_database.id
#   db_workload                      = "OLTP"
#   display_name                     = "example_autonomous_db-2"
#   is_dedicated                     = "true"
#   character_set                    = "AL32UTF8"
#   ncharacter_set                   = "AL16UTF16"
# }

# resource "oci_database_autonomous_database" "test_autonomous_database_developer" {
#   #Required
#   admin_password           = random_string.autonomous_database_admin_password.result
#   compartment_id           = var.compartment_ocid
#   compute_count            = "4"
#   data_storage_size_in_gb  = "32"
#   db_name                  = "atpdb3"
#
#   #Optional
#   autonomous_container_database_id = oci_database_autonomous_container_database.test_autonomous_container_database.id
#   db_workload                      = "OLTP"
#   display_name                     = "example_autonomous_db-developer"
#   is_dedicated                     = "true"
#   character_set                    = "AL32UTF8"
#   ncharacter_set                   = "AL16UTF16"
#   is_dev_tier                      = "true"
# }

# resource "oci_database_autonomous_container_database" "test_autonomous_container_database_primary" {
#   #Required
#   cloud_autonomous_vm_cluster_id       = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster_primary.id
#   display_name                         = "PrimaryACD"
#   patch_model                          = "RELEASE_UPDATES"
#   db_version                           = var.acd_db_version
#   db_name                              = "PRIM${random_string.db_unique_name.result}"
#
#   #Optional
#   backup_config {
#     #Optional
#     recovery_window_in_days = var.autonomous_container_database_backup_config_recovery_window_in_days
#   }
#
#   compartment_id               = var.compartment_ocid
#   freeform_tags                = var.autonomous_database_freeform_tags
#   service_level_agreement_type = "STANDARD"
#
#   maintenance_window_details {
#     preference = "CUSTOM_PREFERENCE"
#
#     days_of_week {
#       name = "MONDAY"
#     }
#
#     hours_of_day = ["4"]
#
#     months {
#       name = "JANUARY"
#     }
#
#     months {
#       name = "APRIL"
#     }
#
#     months {
#       name = "JULY"
#     }
#
#     months {
#       name = "OCTOBER"
#     }
#
#     weeks_of_month = ["2"]
#   }
#   version_preference = "LATEST_RELEASE_UPDATE"
#
#     lifecycle {
#       ignore_changes = [
#           peer_autonomous_container_database_display_name,
#           peer_autonomous_exadata_infrastructure_id,
#           peer_autonomous_vm_cluster_id,
#           peer_cloud_autonomous_vm_cluster_id,
#           peer_db_unique_name,
#           service_level_agreement_type,
#           protection_mode,
#           peer_autonomous_container_database_backup_config,
#       ]
#     }
#
# }

#depends on parameter
# resource "oci_database_autonomous_container_database_dataguard_association" "test_autonomous_container_database_dataguard_association" {
#   #Required
#   autonomous_container_database_id                  = oci_database_autonomous_container_database.test_autonomous_container_database_primary.id
#   is_automatic_failover_enabled                     = false
#   protection_mode                                   = "MAXIMUM_AVAILABILITY"
#   peer_cloud_autonomous_vm_cluster_id               = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster_standby.id
#   peer_autonomous_container_database_display_name   = "StandbyACD"
#   peer_autonomous_container_database_compartment_id = var.compartment_ocid
# }

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

resource "oci_database_autonomous_container_database_add_standby" "example_acd_add_standby" {
  autonomous_container_database_id      = oci_database_autonomous_container_database.test_autonomous_container_database_for_add_standby.id
  protection_mode                       = "MAXIMUM_AVAILABILITY"
  is_automatic_failover_enabled         = false
  peer_cloud_autonomous_vm_cluster_id   = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster_standby.id

  peer_autonomous_container_database_backup_config {
    recovery_window_in_days = 7
    backup_destination_details {
      type                         = "OBJECT_STORE"
      backup_retention_policy_on_terminate = "RETAIN_FOR_72_HOURS"
      is_retention_lock_enabled           = false
    }
  }
}

# Example for oci_database_autonomous_container_database_snapshot_standby resource.
# This creates a snapshot standby for the standby ACD created above.
resource "oci_database_autonomous_container_database_snapshot_standby" "example_acd_snapshot_standby" {
  autonomous_container_database_id = oci_database_autonomous_container_database_add_standby.example_acd_add_standby.dataguard_group_members[1].autonomous_container_database_id
  role = "STANDBY"
}
