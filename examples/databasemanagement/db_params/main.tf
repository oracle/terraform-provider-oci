// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

provider "oci" {
  tenancy_ocid = var.tenancy_ocid
  user_ocid = var.user_ocid
  fingerprint = var.fingerprint
  private_key_path = var.private_key_path
  region = var.region
}

variable "managed_database_id" {
   default = "<database.ocid>"
}

variable "managed_databases_database_parameter_credentials_password" {
  default = "DBPassword"
}

variable "managed_databases_database_parameter_credentials_username" {
  default = "SYS"
}

variable "managed_databases_database_parameter_credentials_role" {
  default = "SYSDBA"
}

variable "managed_databases_database_parameter_parameters_name" {
  default = "open_cursors"
}

variable "managed_databases_database_parameter_parameters_value" {
  default = "305"
}

variable "managed_databases_database_parameter_update_comment" {
  default = "Terraform update of open cursors"
}

variable "managed_databases_database_parameter_scope" {
  default = "BOTH"
}

variable "managed_databases_database_parameter_is_allowed_values_included" {
  default = "false"
}

variable "managed_databases_database_parameter_source" {
  default = "CURRENT"
}

# Change a database parameter value.
resource "oci_database_management_managed_databases_change_database_parameter" "test_managed_databases_change_database_parameter" {
  #Required
  credentials {

    #Optional
    password = var.managed_databases_database_parameter_credentials_password
    role = var.managed_databases_database_parameter_credentials_role
    user_name = var.managed_databases_database_parameter_credentials_username
  }
  managed_database_id = var.managed_database_id
  parameters {
    #Required
    name = var.managed_databases_database_parameter_parameters_name
    value = var.managed_databases_database_parameter_parameters_value

    #Optional
    update_comment = var.managed_databases_database_parameter_update_comment
  }
  scope = var.managed_databases_database_parameter_scope
}

# Reset a database parameter value.
resource "oci_database_management_managed_databases_reset_database_parameter" "test_managed_databases_reset_database_parameter" {
  #Required
  credentials {

    #Optional
    password = var.managed_databases_database_parameter_credentials_password
    role = var.managed_databases_database_parameter_credentials_role
    user_name = var.managed_databases_database_parameter_credentials_username
  }
  managed_database_id = var.managed_database_id
  parameters = [var.managed_databases_database_parameter_parameters_name]
  scope = var.managed_databases_database_parameter_scope
}

# List database parameters filtered by their name and source.
data "oci_database_management_managed_databases_database_parameters" "test_managed_databases_database_parameter" {
  #Required
  managed_database_id = var.managed_database_id

  #Optional
  is_allowed_values_included = var.managed_databases_database_parameter_is_allowed_values_included
  name = var.managed_databases_database_parameter_parameters_name
  source = var.managed_databases_database_parameter_source
}
