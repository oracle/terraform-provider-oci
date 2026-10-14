// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "migration_advanced_parameters_data_type" {
  default = "STRING"
}

variable "migration_advanced_parameters_name" {
  default = "name"
}

variable "migration_advanced_parameters_value" {
  default = "value"
}

variable "migration_advisor_settings_is_ignore_errors" {
  default = false
}

variable "migration_advisor_settings_is_skip_advisor" {
  default = false
}

variable "migration_bulk_include_exclude_data" {
  default = "bulkIncludeExcludeData"
}

variable "migration_data_transfer_medium_details_name" {
  default = "name"
}

variable "migration_data_transfer_medium_details_object_storage_bucket_bucket" {
  default = "bucket"
}

variable "migration_data_transfer_medium_details_object_storage_bucket_namespace" {
  default = "namespace"
}

variable "migration_data_transfer_medium_details_region" {
  default = "region"
}

variable "migration_data_transfer_medium_details_secret_access_key" {
  default = "secretAccessKey"
}

variable "migration_data_transfer_medium_details_source_kind" {
  default = "CURL"
}

variable "migration_data_transfer_medium_details_source_oci_home" {
  default = "ociHome"
}

variable "migration_data_transfer_medium_details_source_wallet_location" {
  default = "walletLocation"
}

variable "migration_data_transfer_medium_details_target_kind" {
  default = "CURL"
}

variable "migration_data_transfer_medium_details_target_oci_home" {
  default = "ociHome"
}

variable "migration_data_transfer_medium_details_target_wallet_location" {
  default = "walletLocation"
}

variable "migration_data_transfer_medium_details_type" {
  default = "DBLINK"
}

variable "migration_database_combination" {
  default = "MYSQL"
}

variable "migration_defined_tags_value" {
  default = "value"
}

variable "migration_description" {
  default = "description"
}

variable "migration_display_name" {
  default = "displayName"
}

variable "migration_exclude_objects_is_omit_excluded_table_from_replication" {
  default = false
}

variable "migration_exclude_objects_object" {
  default = "object"
}

variable "migration_exclude_objects_owner" {
  default = "owner"
}

variable "migration_exclude_objects_schema" {
  default = "schema"
}

variable "migration_exclude_objects_type" {
  default = "type"
}

variable "migration_freeform_tags" {
  default = "freeformTags"
}

variable "migration_ggs_details_acceptable_lag" {
  default = 10
}

variable "migration_ggs_details_extract_long_trans_duration" {
  default = 10
}

variable "migration_ggs_details_extract_performance_profile" {
  default = "LOW"
}

variable "migration_ggs_details_replicat_performance_profile" {
  default = "LOW"
}

variable "migration_hub_details_acceptable_lag" {
  default = 10
}

variable "migration_hub_details_extract_long_trans_duration" {
  default = 10
}

variable "migration_hub_details_extract_performance_profile" {
  default = "LOW"
}

variable "migration_hub_details_replicat_performance_profile" {
  default = "LOW"
}

variable "migration_hub_details_rest_admin_credentials_password" {
  default = "BEstrO0ng_#11"
}

variable "migration_hub_details_rest_admin_credentials_username" {
  default = "username"
}

variable "migration_hub_details_url" {
  default = "url"
}

variable "migration_include_objects_is_omit_excluded_table_from_replication" {
  default = false
}

variable "migration_include_objects_object" {
  default = "object"
}

variable "migration_include_objects_owner" {
  default = "owner"
}

variable "migration_include_objects_schema" {
  default = "schema"
}

variable "migration_include_objects_type" {
  default = "type"
}

variable "migration_initial_load_settings_compatibility" {
  default = []
}

variable "migration_initial_load_settings_data_pump_parameters_estimate" {
  default = "BLOCKS"
}

variable "migration_initial_load_settings_data_pump_parameters_exclude_parameters" {
  default = []
}

variable "migration_initial_load_settings_data_pump_parameters_export_parallelism_degree" {
  default = 10
}

variable "migration_initial_load_settings_data_pump_parameters_import_parallelism_degree" {
  default = 10
}

variable "migration_initial_load_settings_data_pump_parameters_is_cluster" {
  default = false
}

variable "migration_initial_load_settings_data_pump_parameters_table_exists_action" {
  default = "TRUNCATE"
}

variable "migration_initial_load_settings_export_directory_object_name" {
  default = "name"
}

variable "migration_initial_load_settings_export_directory_object_path" {
  default = "path"
}

variable "migration_initial_load_settings_handle_grant_errors" {
  default = "ABORT"
}

variable "migration_initial_load_settings_import_directory_object_name" {
  default = "name"
}

variable "migration_initial_load_settings_import_directory_object_path" {
  default = "path"
}

variable "migration_initial_load_settings_is_consistent" {
  default = false
}

variable "migration_initial_load_settings_is_ignore_existing_objects" {
  default = false
}

variable "migration_initial_load_settings_is_tz_utc" {
  default = false
}

variable "migration_initial_load_settings_job_mode" {
  default = "FULL"
}

variable "migration_initial_load_settings_metadata_remaps_new_value" {
  default = "newValue"
}

variable "migration_initial_load_settings_metadata_remaps_old_value" {
  default = "oldValue"
}

variable "migration_initial_load_settings_metadata_remaps_type" {
  default = "SCHEMA"
}

variable "migration_initial_load_settings_primary_key_compatibility" {
  default = "NONE"
}

variable "migration_initial_load_settings_tablespace_details_block_size_in_kbs" {
  default = "SIZE_2K"
}

variable "migration_initial_load_settings_tablespace_details_extend_size_in_mbs" {
  default = 10
}

variable "migration_initial_load_settings_tablespace_details_is_auto_create" {
  default = false
}

variable "migration_initial_load_settings_tablespace_details_is_big_file" {
  default = false
}

variable "migration_initial_load_settings_tablespace_details_remap_target" {
  default = "remapTarget"
}

variable "migration_initial_load_settings_tablespace_details_target_type" {
  default = "ADB_S_REMAP"
}

variable "migration_lifecycle_details" {
  default = "READY"
}

variable "migration_state" {
  default = "AVAILABLE"
}

variable "migration_type" {
  default = "ONLINE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_database_migration_migration" "test_migration" {
  #Required
  compartment_id       = var.compartment_id
  database_combination = var.migration_database_combination
  type                 = var.migration_type

  #Optional
  advanced_parameters {

    #Optional
    data_type = var.migration_advanced_parameters_data_type
    name      = var.migration_advanced_parameters_name
    value     = var.migration_advanced_parameters_value
  }
  advisor_settings {

    #Optional
    is_ignore_errors = var.migration_advisor_settings_is_ignore_errors
    is_skip_advisor  = var.migration_advisor_settings_is_skip_advisor
  }
  assessment_id             = oci_database_migration_assessment.test_assessment.id
  bulk_include_exclude_data = var.migration_bulk_include_exclude_data
  data_transfer_medium_details {
    #Required
    type = var.migration_data_transfer_medium_details_type

    #Optional
    access_key_id = oci_kms_key.test_key.id
    name          = var.migration_data_transfer_medium_details_name
    object_storage_bucket {

      #Optional
      bucket    = var.migration_data_transfer_medium_details_object_storage_bucket_bucket
      namespace = var.migration_data_transfer_medium_details_object_storage_bucket_namespace
    }
    region                         = var.migration_data_transfer_medium_details_region
    secret_access_key              = var.migration_data_transfer_medium_details_secret_access_key
    shared_storage_mount_target_id = oci_file_storage_mount_target.test_mount_target.id
    source {
      #Required
      kind = var.migration_data_transfer_medium_details_source_kind

      #Optional
      oci_home        = var.migration_data_transfer_medium_details_source_oci_home
      wallet_location = var.migration_data_transfer_medium_details_source_wallet_location
    }
    target {
      #Required
      kind = var.migration_data_transfer_medium_details_target_kind

      #Optional
      oci_home        = var.migration_data_transfer_medium_details_target_oci_home
      wallet_location = var.migration_data_transfer_medium_details_target_wallet_location
    }
  }
  defined_tags = tomap({ oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name = var.migration_defined_tags_value })
  description  = var.migration_description
  display_name = var.migration_display_name
  exclude_objects {
    #Required
    object = var.migration_exclude_objects_object

    #Optional
    is_omit_excluded_table_from_replication = var.migration_exclude_objects_is_omit_excluded_table_from_replication
    owner                                   = var.migration_exclude_objects_owner
    schema                                  = var.migration_exclude_objects_schema
    type                                    = var.migration_exclude_objects_type
  }
  freeform_tags = var.migration_freeform_tags
  ggs_details {

    #Optional
    acceptable_lag = var.migration_ggs_details_acceptable_lag
    extract {

      #Optional
      long_trans_duration = var.migration_ggs_details_extract_long_trans_duration
      performance_profile = var.migration_ggs_details_extract_performance_profile
    }
    replicat {

      #Optional
      performance_profile = var.migration_ggs_details_replicat_performance_profile
    }
  }
  hub_details {
    #Required
    key_id = oci_kms_key.test_key.id
    rest_admin_credentials {
      #Required
      password = var.migration_hub_details_rest_admin_credentials_password
      username = var.migration_hub_details_rest_admin_credentials_username
    }
    url      = var.migration_hub_details_url
    vault_id = oci_kms_vault.test_vault.id

    #Optional
    acceptable_lag = var.migration_hub_details_acceptable_lag
    compute_id     = oci_database_migration_compute.test_compute.id
    extract {

      #Optional
      long_trans_duration = var.migration_hub_details_extract_long_trans_duration
      performance_profile = var.migration_hub_details_extract_performance_profile
    }
    replicat {

      #Optional
      performance_profile = var.migration_hub_details_replicat_performance_profile
    }
  }
  include_objects {
    #Required
    object = var.migration_include_objects_object

    #Optional
    is_omit_excluded_table_from_replication = var.migration_include_objects_is_omit_excluded_table_from_replication
    owner                                   = var.migration_include_objects_owner
    schema                                  = var.migration_include_objects_schema
    type                                    = var.migration_include_objects_type
  }
  initial_load_settings {

    #Optional
    compatibility = var.migration_initial_load_settings_compatibility
    data_pump_parameters {

      #Optional
      estimate                  = var.migration_initial_load_settings_data_pump_parameters_estimate
      exclude_parameters        = var.migration_initial_load_settings_data_pump_parameters_exclude_parameters
      export_parallelism_degree = var.migration_initial_load_settings_data_pump_parameters_export_parallelism_degree
      import_parallelism_degree = var.migration_initial_load_settings_data_pump_parameters_import_parallelism_degree
      is_cluster                = var.migration_initial_load_settings_data_pump_parameters_is_cluster
      table_exists_action       = var.migration_initial_load_settings_data_pump_parameters_table_exists_action
    }
    export_directory_object {

      #Optional
      name = var.migration_initial_load_settings_export_directory_object_name
      path = var.migration_initial_load_settings_export_directory_object_path
    }
    handle_grant_errors = var.migration_initial_load_settings_handle_grant_errors
    import_directory_object {

      #Optional
      name = var.migration_initial_load_settings_import_directory_object_name
      path = var.migration_initial_load_settings_import_directory_object_path
    }
    is_consistent              = var.migration_initial_load_settings_is_consistent
    is_ignore_existing_objects = var.migration_initial_load_settings_is_ignore_existing_objects
    is_tz_utc                  = var.migration_initial_load_settings_is_tz_utc
    job_mode                   = var.migration_initial_load_settings_job_mode
    metadata_remaps {

      #Optional
      new_value = var.migration_initial_load_settings_metadata_remaps_new_value
      old_value = var.migration_initial_load_settings_metadata_remaps_old_value
      type      = var.migration_initial_load_settings_metadata_remaps_type
    }
    primary_key_compatibility = var.migration_initial_load_settings_primary_key_compatibility
    tablespace_details {
      #Required
      target_type = var.migration_initial_load_settings_tablespace_details_target_type

      #Optional
      block_size_in_kbs  = var.migration_initial_load_settings_tablespace_details_block_size_in_kbs
      extend_size_in_mbs = var.migration_initial_load_settings_tablespace_details_extend_size_in_mbs
      is_auto_create     = var.migration_initial_load_settings_tablespace_details_is_auto_create
      is_big_file        = var.migration_initial_load_settings_tablespace_details_is_big_file
      remap_target       = var.migration_initial_load_settings_tablespace_details_remap_target
    }
  }
  source_container_database_connection_id = oci_database_migration_connection.test_connection.id
  source_database_connection_id           = oci_database_migration_connection.test_connection.id
  source_standby_database_connection_id   = oci_database_migration_connection.test_connection.id
  target_database_connection_id           = oci_database_migration_connection.test_connection.id
}

data "oci_database_migration_migrations" "test_migrations" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name      = var.migration_display_name
  lifecycle_details = var.migration_lifecycle_details
  state             = var.migration_state
}

