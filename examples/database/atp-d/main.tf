// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_database_autonomous_container_database" "test_autonomous_container_database" {
  #Required
  cloud_autonomous_vm_cluster_id       = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id
  display_name                         = "example-container-database"
  patch_model                          = "RELEASE_UPDATES"
  db_version                           = "19.23.0.1.0"
  db_name                              = "ACDNAME"

  #Optional
  backup_config {
    #Optional
    recovery_window_in_days = var.autonomous_container_database_backup_config_recovery_window_in_days
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

resource "random_string" "db_unique_name_adsi_acd" {
  length = 8
  special = false
  number = false
}

resource "oci_database_autonomous_container_database" "autonomous_container_database_from_adsi" {
  cloud_autonomous_vm_cluster_id       = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id
  database_software_image_id = oci_database_autonomous_database_software_image.autonomous_database_software_image.id
  backup_config {
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
variable "cloud_exadata_infrastructure_un_allocated_resource_db_servers" {
  default = []
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

resource "oci_database_autonomous_database" "test_autonomous_database_character_set_support" {
  #Required
  admin_password           = random_string.autonomous_database_admin_password.result
  compartment_id           = var.compartment_ocid
  compute_count            = "2"
  data_storage_size_in_tbs = "1"
  db_name                  = "atpdb2"

  #Optional
  autonomous_container_database_id = oci_database_autonomous_container_database.test_autonomous_container_database.id
  db_workload                      = "OLTP"
  display_name                     = "example_autonomous_db-2"
  is_dedicated                     = "true"
  character_set                    = "AL32UTF8"
  ncharacter_set                   = "AL16UTF16"
}

resource "oci_database_autonomous_database" "test_autonomous_database_developer" {
  #Required
  admin_password           = random_string.autonomous_database_admin_password.result
  compartment_id           = var.compartment_ocid
  compute_count            = "4"
  data_storage_size_in_gb  = "32"
  db_name                  = "atpdb3"

  #Optional
  autonomous_container_database_id = oci_database_autonomous_container_database.test_autonomous_container_database.id
  db_workload                      = "OLTP"
  display_name                     = "example_autonomous_db-developer"
  is_dedicated                     = "true"
  character_set                    = "AL32UTF8"
  ncharacter_set                   = "AL16UTF16"
  is_dev_tier                      = "true"
}

data "oci_database_autonomous_container_databases" "test_autonomous_container_databases" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  cloud_autonomous_vm_cluster_id       = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id
  availability_domain                  = data.oci_identity_availability_domain.ad.name
  display_name                         = "example-container-database"
  state                                = "AVAILABLE"
}

data "oci_database_autonomous_databases" "autonomous_databases" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  autonomous_container_database_id = oci_database_autonomous_container_database.test_autonomous_container_database.id
  display_name                     = oci_database_autonomous_database.test_autonomous_database.display_name
  db_workload                      = "OLTP"
}

output "autonomous_database_admin_password" {
  value = random_string.autonomous_database_admin_password.result
}

output "autonomous_database_high_connection_string" {
  value = lookup(
    oci_database_autonomous_database.test_autonomous_database.connection_strings[0].all_connection_strings,
    "high",
    "unavailable",
  )
}

output "autonomous_databases" {
  value = data.oci_database_autonomous_databases.autonomous_databases.autonomous_databases
}

output "autonomous_container_databases" {
  value = data.oci_database_autonomous_container_databases.test_autonomous_container_databases.autonomous_container_databases
}

data "oci_database_cloud_exadata_infrastructure_un_allocated_resource" "test_cloud_exadata_infrastructure_un_allocated_resources" {
  #Required
  cloud_exadata_infrastructure_id = oci_database_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure.id

  #Optional
  db_servers = var.cloud_exadata_infrastructure_un_allocated_resource_db_servers
}

data "oci_database_autonomous_container_database_resource_usage" "test_autonomous_container_database_resource_usages" {
  #Required
  autonomous_container_database_id = oci_database_autonomous_container_database.test_autonomous_container_database.id
}

data "oci_database_cloud_autonomous_vm_cluster_acd_resource_usages" "test_cloud_autonomous_vm_cluster_acd_resource_usages" {
  #Required
  cloud_autonomous_vm_cluster_id = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id

  #Optional
  compartment_id = var.compartment_ocid
}

data "oci_database_cloud_autonomous_vm_cluster_resource_usage" "test_cloud_autonomous_vm_cluster_resource_usages" {
  #Required
  cloud_autonomous_vm_cluster_id = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id
}

resource "oci_database_autonomous_container_database" "test_autonomous_container_database_primary" {
  #Required
  cloud_autonomous_vm_cluster_id       = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster_primary.id
  display_name                         = "PrimaryACD"
  patch_model                          = "RELEASE_UPDATES"
  db_version                           = "19.22.0.1.0"
  db_name                              = "PRIMARY"

  #Optional
  backup_config {
    #Optional
    recovery_window_in_days = var.autonomous_container_database_backup_config_recovery_window_in_days
  }

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
  version_preference = "LATEST_RELEASE_UPDATE"

    lifecycle {
      ignore_changes = [
          peer_autonomous_container_database_display_name,
          peer_autonomous_exadata_infrastructure_id,
          peer_autonomous_vm_cluster_id,
          peer_cloud_autonomous_vm_cluster_id,
          peer_db_unique_name,
          service_level_agreement_type,
          protection_mode,
          peer_autonomous_container_database_backup_config,
      ]
    }

}

resource "oci_database_autonomous_container_database_dataguard_association" "test_autonomous_container_database_dataguard_association" {
  #Required
  autonomous_container_database_id                  = oci_database_autonomous_container_database.test_autonomous_container_database_primary.id
  is_automatic_failover_enabled                     = false
  protection_mode                                   = "MAXIMUM_AVAILABILITY"
  peer_cloud_autonomous_vm_cluster_id               = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster_standby.id
  peer_autonomous_container_database_display_name   = "StandbyACD"
  peer_autonomous_container_database_compartment_id = var.compartment_ocid
}

