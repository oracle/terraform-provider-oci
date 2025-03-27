// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "random_string" "autonomous_database_admin_password" {
  length      = 16
  min_numeric = 1
  min_lower   = 1
  min_upper   = 1
  min_special = 1
}

data "oci_database_autonomous_db_versions" "test_autonomous_db_versions" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  db_workload = var.autonomous_database_db_workload

  filter {
    name   = "version"
    values = ["19c"]
  }
}

resource "oci_database_autonomous_database" "autonomous_database" {
  #Required
  admin_password           = random_string.autonomous_database_admin_password.result
  compartment_id           = var.compartment_ocid
  cpu_core_count           = "1"
  data_storage_size_in_tbs = "1"
  db_name                  = "adbDatabaseName123"

  #Optional
  db_version                                     = data.oci_database_autonomous_db_versions.test_autonomous_db_versions.autonomous_db_versions[0].version
  db_workload                                    = var.autonomous_database_db_workload
  display_name                                   = "example_autonomous_database"
  freeform_tags                                  = var.autonomous_database_freeform_tags
  is_auto_scaling_enabled                        = "true"
  is_auto_scaling_for_storage_enabled            = "true"
  license_model                                  = var.autonomous_database_license_model
  is_preview_version_with_service_terms_accepted = "false"
  whitelisted_ips                                = ["1.1.1.1/28"]
  character_set                                  = "AL32UTF8"
  ncharacter_set                                 = "AL16UTF16"
}

resource "oci_database_autonomous_database" "test_autonomous_database" {
  admin_password                       = random_string.autonomous_database_admin_password.result
  compartment_id                       = var.compartment_ocid
  cpu_core_count                       = "1"
  data_storage_size_in_tbs             = "1"
  db_name                              = "adbdb11f"
  db_version                           = "19c"
  db_workload                          = "AJD"
  license_model                        = "LICENSE_INCLUDED"
  is_free_tier                         = "false"
  autonomous_maintenance_schedule_type = var.autonomous_database_autonomous_maintenance_schedule_type
}



resource "oci_database_autonomous_database" "test_autonomous_database1" {
  admin_password                       = random_string.autonomous_database_admin_password.result
  compartment_id                       = var.compartment_ocid
  cpu_core_count                       = "1"
  data_storage_size_in_tbs             = "1"
  db_name                              = "adbdb15f"
  db_version                           = "19c"
  db_workload                          = "OLTP"
  license_model                        = "BRING_YOUR_OWN_LICENSE"
  is_free_tier                         = "false"
  autonomous_maintenance_schedule_type = var.autonomous_database_autonomous_maintenance_schedule_type
  open_mode                            = "READ_ONLY"
  permission_level                     = "RESTRICTED"
  data_safe_status                     = "REGISTERED"
  database_edition                     = "STANDARD_EDITION"
  operations_insights_status           = "ENABLED"
  database_management_status           = "ENABLED"
}

resource "oci_database_autonomous_database" "test_autonomous_database_ecpu" {
  admin_password           = random_string.autonomous_database_admin_password.result
  compartment_id           = var.compartment_ocid
  compute_count            = "2.0"
  compute_model            = "ECPU"
  data_storage_size_in_tbs = "1"
  db_name                  = "adbEcpu"
  db_version               = "19c"
  license_model            = "LICENSE_INCLUDED"
  is_free_tier             = "false"
}

resource "oci_database_autonomous_database" "test_autonomous_database_developer" {
  admin_password           = random_string.autonomous_database_admin_password.result
  compartment_id           = var.compartment_ocid
  compute_count            = "4.0"
  compute_model            = "ECPU"
  data_storage_size_in_gb = "20"
  db_name                  = "adbDeveloper"
  db_version               = "19c"
  license_model            = "LICENSE_INCLUDED"
  is_free_tier             = "false"
}

resource "oci_database_autonomous_database" "test_autonomous_database_local_adg_failover_data_loss_limit" {
  admin_password           = random_string.autonomous_database_admin_password.result
  compartment_id           = var.compartment_ocid
  cpu_core_count           = "1"
  data_storage_size_in_tbs = "1"
  db_name                  = "adbLocalAdg"
  is_local_data_guard_enabled = "true"
  local_adg_auto_failover_max_data_loss_limit = "30"
  db_version               = "19c"
  db_workload              = "OLTP"
  license_model            = "LICENSE_INCLUDED"
  is_free_tier             = "false"
}

resource "oci_database_autonomous_database" "test_autonomous_database_apex" {
  admin_password           = random_string.autonomous_database_admin_password.result
  compartment_id           = var.compartment_ocid
  cpu_core_count           = "1"
  data_storage_size_in_tbs = "1"
  db_name                  = "adbapex"
  db_version               = "19c"
  db_workload              = "APEX"
  license_model            = "LICENSE_INCLUDED"
  is_free_tier             = "false"
}

resource "oci_database_autonomous_database" "test_autonomous_database_bck_ret_lock" {
  admin_password           = random_string.autonomous_database_admin_password.result
  compartment_id           = var.compartment_ocid
  compute_count            = "2.0"
  compute_model            = "ECPU"
  data_storage_size_in_tbs = "1"
  db_name                  = "adbBckRetLock"
  db_version               = "19c"
  license_model            = "LICENSE_INCLUDED"
  is_free_tier             = "false"
  is_backup_retention_locked = "false"
}


resource "oci_database_autonomous_database" "test_autonomous_database_bck_ret_days" {
  admin_password                       = random_string.autonomous_database_admin_password.result
  compartment_id                       = var.compartment_ocid
  cpu_core_count                       = "1"
  data_storage_size_in_tbs             = "1"
  backup_retention_period_in_days      = "15"
  db_name                              = "adbBckRetDays"
  db_version                           = "19c"
  db_workload                          = "AJD"
  license_model                        = "LICENSE_INCLUDED"
  is_free_tier                         = "false"
}

resource "oci_database_autonomous_database" "autonomous_database_private_ip_with_acls" {
  #Required
  admin_password           = random_string.autonomous_database_admin_password.result
  compartment_id           = var.compartment_ocid
  compute_count            = "2.0"
  compute_model            = "ECPU"
  data_storage_size_in_tbs = "1"
  db_name                  = "adbdbpeacl"

  #Optional
  db_version                                     = data.oci_database_autonomous_db_versions.test_autonomous_db_versions.autonomous_db_versions[0].version
  db_workload                                    = var.autonomous_database_db_workload
  display_name                                   = "example_autonomous_database"
  freeform_tags                                  = var.autonomous_database_freeform_tags
  is_auto_scaling_enabled                        = "true"
  is_auto_scaling_for_storage_enabled            = "true"
  license_model                                  = var.autonomous_database_license_model
  is_preview_version_with_service_terms_accepted = "false"
  character_set                                  = "AL32UTF8"
  ncharacter_set                                 = "AL16UTF16"
  subnet_id = oci_core_subnet.test_subnet.id
  nsg_ids = ["test-bn-nsg-id-1"]
  whitelisted_ips = ["1.1.1.28"]
}

resource "oci_database_autonomous_database" "test_autonomous_database_db_tools" {
  admin_password                       = random_string.autonomous_database_admin_password.result
  compartment_id                       = var.compartment_ocid
  cpu_core_count                       = "1"
  data_storage_size_in_tbs             = "1"
  backup_retention_period_in_days      = "15"
  db_name                              = "adbdbtoolsL"
  db_version                           = "19c"
  db_workload                          = "AJD"
  license_model                        = "LICENSE_INCLUDED"
  is_free_tier                         = "false"
  db_tools_details {
              name = "APEX"
                is_enabled = true
                compute_count = 0
                max_idle_time_in_minutes = 0
              }
              db_tools_details {
                name = "DATA_TRANSFORMS"
                is_enabled = true
                compute_count = 0
                max_idle_time_in_minutes = 0
              }
              db_tools_details {
                name = "DATABASE_ACTIONS"
                is_enabled = true
                compute_count = 0
                max_idle_time_in_minutes = 0
              }
              db_tools_details {
                name = "GRAPH_STUDIO"
                is_enabled = false
                compute_count = 0
                max_idle_time_in_minutes = 0
              }
              db_tools_details {
                name = "MONGODB_API"
                is_enabled = true
                compute_count = 0
                max_idle_time_in_minutes = 0
              }
              db_tools_details {
                name = "OML"
                is_enabled = true
                compute_count = 0
                max_idle_time_in_minutes = 0
              }
              db_tools_details {
                name = "ORDS"
                is_enabled = true
                compute_count = 0
                max_idle_time_in_minutes = 0
              }
    }

resource "oci_database_autonomous_database" "test_autonomous_database_actions" {
  admin_password           = random_string.autonomous_database_admin_password.result
  compartment_id           = var.compartment_ocid
  cpu_core_count           = "1"
  data_storage_size_in_tbs = "1"
  db_name                  = "adbdb11ff"
  db_version               = "19c"
  db_workload              = "AJD"
  license_model            = "LICENSE_INCLUDED"
  is_free_tier             = "false"
  state                    = "STOPPED"
}


resource "oci_database_autonomous_database" "autonomous_database_oneway_tls_connection" {
  admin_password              = random_string.autonomous_database_admin_password.result
  compartment_id              = var.compartment_ocid
  cpu_core_count              = "1"
  data_storage_size_in_tbs    = "1"
  db_name                     = "adbOneWay"

  whitelisted_ips             = ["1.1.1.1"]
  is_mtls_connection_required = "true"
}

resource "oci_database_autonomous_database" "autonomous_database_dbms_status" {
  admin_password              = random_string.autonomous_database_admin_password.result
  compartment_id              = var.compartment_ocid
  cpu_core_count              = "1"
  data_storage_size_in_tbs    = "1"
  db_name                     = "adbdbms"
  database_management_status  = "ENABLED"
}

// Per service, we need to pass in a back up that is at least 2 hours old
/*
resource "oci_database_autonomous_database" "autonomous_database_from_backup_id" {
  #Required
  admin_password           = random_string.autonomous_database_admin_password.result
  compartment_id           = var.compartment_ocid
  cpu_core_count           = "1"
  data_storage_size_in_tbs = "1"
  db_name                  = "adbdb2"

  clone_type                    = "FULL"
  source                        = "BACKUP_FROM_ID"
  autonomous_database_backup_id = oci_database_autonomous_database_backup.autonomous_database_backup.id
}

resource "oci_database_autonomous_database" "autonomous_database_from_backup_timestamp" {
  #Required
  admin_password           = random_string.autonomous_database_admin_password.result
  compartment_id           = var.compartment_ocid
  cpu_core_count           = "1"
  data_storage_size_in_tbs = "1"
  db_name                  = "adbdb3"

  clone_type             = "FULL"
  source                 = "BACKUP_FROM_TIMESTAMP"
  autonomous_database_id = oci_database_autonomous_database_backup.autonomous_database_backup.autonomous_database_id
  timestamp              = oci_database_autonomous_database_backup.autonomous_database_backup.time_ended
}
*/

/*
resource "oci_database_autonomous_database" "autonomous_database_from_backup_latest" {
  #Required
  admin_password           = random_string.autonomous_database_admin_password.result
  compartment_id           = var.compartment_ocid
  cpu_core_count           = "1"
  data_storage_size_in_tbs = "1"
  db_name                  = "adbdb3"

  clone_type             = "FULL"
  source                 = "BACKUP_FROM_TIMESTAMP"
  autonomous_database_id = oci_database_autonomous_database_backup.autonomous_database_backup.autonomous_database_id
  use_latest_available_backup_time_stamp              = "true"
  whitelisted_ips             = ["1.1.1.1/28"]
}
*/
resource "oci_core_vcn" "test_vcn" {
  cidr_block = "10.0.0.0/16"
  compartment_id = var.compartment_ocid
  display_name = "test_vcn"
}

resource "oci_core_subnet" "test_subnet" {
  cidr_block = "10.0.0.0/24"
  compartment_id = var.compartment_ocid
  vcn_id = oci_core_vcn.test_vcn.id
  display_name = "test_vcn"
}

resource "oci_database_autonomous_database" "autonomous_database_private_ip" {
  #Required
  admin_password           = random_string.autonomous_database_admin_password.result
  compartment_id           = var.compartment_ocid
  cpu_core_count           = "1"
  data_storage_size_in_tbs = "1"
  db_name                  = "adbdbip"

  #Optional
  db_version                                     = data.oci_database_autonomous_db_versions.test_autonomous_db_versions.autonomous_db_versions[0].version
  db_workload                                    = var.autonomous_database_db_workload
  display_name                                   = "example_autonomous_database"
  freeform_tags                                  = var.autonomous_database_freeform_tags
  is_auto_scaling_enabled                        = "true"
  is_auto_scaling_for_storage_enabled            = "true"
  license_model                                  = var.autonomous_database_license_model
  is_preview_version_with_service_terms_accepted = "false"
  character_set                                  = "AL32UTF8"
  ncharacter_set                                 = "AL16UTF16"
  private_endpoint_ip = "10.0.0.97"
  subnet_id = oci_core_subnet.test_subnet.id
  nsg_ids = ["test-bn-nsg-id-1"]
}

data "oci_database_autonomous_database_character_sets" "autonomous_databases_character_sets" {
  #Optional
  character_set_type = var.autonomous_database_character_set_character_set_type
  is_shared = var.autonomous_database_character_set_is_shared
  is_dedicated =  var.autonomous_database_character_set_is_dedicated
}

data "oci_database_autonomous_databases" "autonomous_databases" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name = oci_database_autonomous_database.autonomous_database.display_name
  db_workload  = var.autonomous_database_db_workload
}

data "oci_database_autonomous_database_refreshable_clones" "autonomous_database_refreshable_clones" {
  #Required
  autonomous_database_id = oci_database_autonomous_database.autonomous_database.id
}

output "autonomous_database_admin_password" {
  value = random_string.autonomous_database_admin_password.result
}

output "autonomous_database_high_connection_string" {
  value = lookup(
    oci_database_autonomous_database.autonomous_database.connection_strings[0],
    "high",
    "unavailable",
  )
}

output "autonomous_databases" {
  value = data.oci_database_autonomous_databases.autonomous_databases.autonomous_databases
}

resource "oci_database_autonomous_database" "test_autonomous_database_shrink" {
  admin_password           = random_string.autonomous_database_admin_password.result
  compartment_id           = var.compartment_ocid
  cpu_core_count           = "16"
  data_storage_size_in_tbs = "16"
  db_name                  = "adbapex18"
  db_version               = "19c"
  db_workload              = "OLTP"
  license_model            = "LICENSE_INCLUDED"
  shrink_adb_trigger       = "2"
  is_auto_scaling_for_storage_enabled = "true"
}
