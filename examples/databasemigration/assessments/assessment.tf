// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "assessment_acceptable_downtime" {
  default = "LESS_THAN_10_MINUTES"
}

variable "assessment_bulk_include_exclude_data" {
  default = "bulkIncludeExcludeData"
}

variable "assessment_creation_type" {
  default = "CREATE_ONLY"
}

variable "assessment_database_combination" {
  default = "MYSQL"
}

variable "assessment_database_data_size" {
  default = "LESS_THAN_1GB"
}

variable "assessment_ddl_expectation" {
  default = "DDL_EXPECTED"
}

variable "assessment_defined_tags_value" {
  default = "value"
}

variable "assessment_description" {
  default = "description"
}

variable "assessment_display_name" {
  default = "displayName"
}

variable "assessment_exclude_objects_is_omit_excluded_table_from_replication" {
  default = false
}

variable "assessment_exclude_objects_object" {
  default = "object"
}

variable "assessment_exclude_objects_owner" {
  default = "owner"
}

variable "assessment_exclude_objects_schema" {
  default = "schema"
}

variable "assessment_exclude_objects_type" {
  default = "type"
}

variable "assessment_freeform_tags" {
  default = "freeformTags"
}

variable "assessment_include_objects_is_omit_excluded_table_from_replication" {
  default = false
}

variable "assessment_include_objects_object" {
  default = "object"
}

variable "assessment_include_objects_owner" {
  default = "owner"
}

variable "assessment_include_objects_schema" {
  default = "schema"
}

variable "assessment_include_objects_type" {
  default = "type"
}

variable "assessment_lifecycle_details" {
  default = "READY"
}

variable "assessment_migration_scope" {
  default = "SCHEMA"
}

variable "assessment_network_speed_megabit_per_second" {
  default = "MBPS_10"
}

variable "assessment_source_database_connection_id" {
  default = "id"
}

variable "assessment_state" {
  default = "AVAILABLE"
}

variable "assessment_target_database_connection_connection_type" {
  default = "MYSQL"
}

variable "assessment_target_database_connection_database_version" {
  default = "databaseVersion"
}

variable "assessment_target_database_connection_id" {
  default = "id"
}

variable "assessment_target_database_connection_technology_sub_type" {
  default = "technologySubType"
}

variable "assessment_target_database_connection_technology_type" {
  default = "OCI_AUTONOMOUS_DATABASE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_database_migration_assessment" "test_assessment" {
  #Required
  acceptable_downtime              = var.assessment_acceptable_downtime
  compartment_id                   = var.compartment_id
  database_combination             = var.assessment_database_combination
  database_data_size               = var.assessment_database_data_size
  ddl_expectation                  = var.assessment_ddl_expectation
  migration_scope                  = var.assessment_migration_scope
  network_speed_megabit_per_second = var.assessment_network_speed_megabit_per_second
  source_database_connection {
    #Required
    id = var.assessment_source_database_connection_id
  }
  target_database_connection {

    #Optional
    connection_type     = var.assessment_target_database_connection_connection_type
    database_version    = var.assessment_target_database_connection_database_version
    id                  = var.assessment_target_database_connection_id
    technology_sub_type = var.assessment_target_database_connection_technology_sub_type
    technology_type     = var.assessment_target_database_connection_technology_type
  }

  #Optional
  bulk_include_exclude_data = var.assessment_bulk_include_exclude_data
  creation_type             = var.assessment_creation_type
  defined_tags              = tomap({ oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name = var.assessment_defined_tags_value })
  description               = var.assessment_description
  display_name              = var.assessment_display_name
  exclude_objects {
    #Required
    object = var.assessment_exclude_objects_object

    #Optional
    is_omit_excluded_table_from_replication = var.assessment_exclude_objects_is_omit_excluded_table_from_replication
    owner                                   = var.assessment_exclude_objects_owner
    schema                                  = var.assessment_exclude_objects_schema
    type                                    = var.assessment_exclude_objects_type
  }
  freeform_tags = var.assessment_freeform_tags
  include_objects {
    #Required
    object = var.assessment_include_objects_object

    #Optional
    is_omit_excluded_table_from_replication = var.assessment_include_objects_is_omit_excluded_table_from_replication
    owner                                   = var.assessment_include_objects_owner
    schema                                  = var.assessment_include_objects_schema
    type                                    = var.assessment_include_objects_type
  }
}

data "oci_database_migration_assessments" "test_assessments" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name      = var.assessment_display_name
  lifecycle_details = var.assessment_lifecycle_details
  state             = var.assessment_state
}

