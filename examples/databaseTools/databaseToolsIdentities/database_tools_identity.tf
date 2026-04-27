// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}
variable "database_tools_connection_ocid" {}

variable "database_tools_identity_credential_key" {
  default = "Key1"
}

variable "database_tools_identity_defined_tags_value" {
  default = "value"
}

variable "database_tools_identity_display_name" {
  default = "Identity1"
}

variable "database_tools_identity_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "database_tools_identity_locks_message" {
  default = "message"
}

variable "database_tools_identity_locks_time_created" {
  default = "timeCreated"
}

variable "database_tools_identity_locks_type" {
  default = "ORACLE_DATABASE_RESOURCE_PRINCIPAL"
}

variable "database_tools_identity_state" {
  default = "ACTIVE"
}

variable "database_tools_identity_type" {
  default = "ORACLE_DATABASE_RESOURCE_PRINCIPAL"
}


provider "oci" {
  region              = var.region
  auth                = "SecurityToken"
  config_file_profile = "terraform-federation-test"
}

resource "oci_database_tools_database_tools_identity" "test_database_tools_identity" {
  #Required
  compartment_id               = var.compartment_id
  credential_key               = var.database_tools_identity_credential_key
  database_tools_connection_id = var.database_tools_connection_ocid
  display_name                 = var.database_tools_identity_display_name
  type                         = var.database_tools_identity_type

  #Optional
  defined_tags  = {}
  freeform_tags = var.database_tools_identity_freeform_tags
}

data "oci_database_tools_database_tools_identities" "test_database_tools_identities" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  database_tools_connection_id = var.database_tools_connection_ocid
  display_name                 = var.database_tools_identity_display_name
  state                        = var.database_tools_identity_state
  type                         = [var.database_tools_identity_type]
}

