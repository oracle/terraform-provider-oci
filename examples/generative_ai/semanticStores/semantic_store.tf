// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "semantic_store_data_source_connection_type" {
  default = "DATABASE_TOOLS_CONNECTION"
}

variable "semantic_store_defined_tags_value" {
  default = "value"
}

variable "semantic_store_description" {
  default = "description"
}

variable "semantic_store_display_name" {
  default = "displayName"
}

variable "semantic_store_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "semantic_store_id" {
  default = "id"
}

variable "semantic_store_refresh_schedule_type" {
  default = "INTERVAL"
}

variable "semantic_store_refresh_schedule_value" {
  default = "value"
}

variable "semantic_store_schema_name" {
  default = "schemaName"
}

variable "semantic_store_state" {
  default = []
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_generative_ai_semantic_store" "test_semantic_store" {
  #Required
  compartment_id = var.compartment_id
  data_source {
    #Required
    connection_type          = var.semantic_store_data_source_connection_type
    enrichment_connection_id = oci_database_tools_database_tools_connection.test_connection.id
    querying_connection_id   = oci_database_tools_database_tools_connection.test_connection.id
  }
  display_name = var.semantic_store_display_name

  schemas {
    #Required
    connection_type = var.semantic_store_data_source_connection_type
    schemas {
      #Required
      name = var.semantic_store_schema_name
    }
  }

  #Optional
  defined_tags  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.semantic_store_defined_tags_value)
  description   = var.semantic_store_description
  freeform_tags = var.semantic_store_freeform_tags
  refresh_schedule {
    #Required
    type = var.semantic_store_refresh_schedule_type

    #Optional
    value = var.semantic_store_refresh_schedule_value
  }
}

data "oci_generative_ai_semantic_stores" "test_semantic_stores" {

  #Optional
  compartment_id                     = var.compartment_id
  data_source_querying_connection_id = oci_database_migration_connection.test_connection.id
  display_name                       = var.semantic_store_display_name
  id                                 = var.semantic_store_id
  state                              = var.semantic_store_state
}

