// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "data_safe_target_ocid" {}


variable "sensitive_data_models_sensitive_schema_schema_name_list" {
  type = list(string)
  default = ["schemaName"]
}

variable "sensitive_data_models_sensitive_schema_schema_name_var" {
  default = "schemaName"
}


variable "sensitive_data_models_sensitive_column_object_var" {
  default = "object"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

variable "sensitive_data_model_display_name" {
  default = "displayName"
}

variable "sensitive_data_model_schemas_for_discovery" {
  default = []
}

variable "sensitive_data_model_sensitive_type_ids_for_discovery" {
  default = []
}

variable "sensitive_data_models_sensitive_column_column_name_var" {
  default = "columnName"
}

variable "sensitive_data_models_sensitive_column_schema_name_var" {
  default = "schemaName"
}

resource "oci_data_safe_sensitive_data_model" "test_sensitive_data_model" {
  #Required
  compartment_id = var.compartment_ocid
  target_id      = var.data_safe_target_ocid

  #Optional
  display_name                              = var.sensitive_data_model_display_name
  schemas_for_discovery                     = var.sensitive_data_model_schemas_for_discovery
  sensitive_type_ids_for_discovery          = var.sensitive_data_model_sensitive_type_ids_for_discovery
}

resource "oci_data_safe_sensitive_data_models_sensitive_column" "test_sensitive_data_models_sensitive_column" {
  #Required
  column_name             = var.sensitive_data_models_sensitive_column_column_name_var
  object                  = var.sensitive_data_models_sensitive_column_object_var
  schema_name             = var.sensitive_data_models_sensitive_column_schema_name_var
  sensitive_data_model_id = oci_data_safe_sensitive_data_model.test_sensitive_data_model.id

}

data "oci_data_safe_sensitive_data_model_sensitive_schemas" "test_sensitive_data_model_sensitive_schemas" {
  #Required
  sensitive_data_model_id = oci_data_safe_sensitive_data_model.test_sensitive_data_model.id
  schema_name                           = var.sensitive_data_models_sensitive_schema_schema_name_list
}

