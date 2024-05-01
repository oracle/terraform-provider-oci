// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

variable "protected_database_database_size" {
  default = "XS"
}

variable "protected_database_db_unique_name" {
  default = "dbUniqueName"
}

variable "protected_database_defined_tags_value" {
  default = "value"
}

variable "protected_database_display_name" {
  default = "displayName"
}

variable "database_id" {
  default = "database-id"
}

variable "protected_database_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "protected_database_id" {
  default = "id"
}

variable "protected_database_is_redo_logs_shipped" {
  default = false
}

variable "protected_database_password" {
  default = "BEstrO0ng_secret#11"
}

variable "protected_database_state" {
  default = "ACTIVE"
}

variable "recovery_service_subnet_id" {
}

variable "protection_policy_id" {
  default = "ocid1.recoveryservicepolicy.region1..aaaaaaaam22xkw32t524xvst7dbxz4qsxtwetmfnnxfsgslbq664vya5jbkq"
}

resource "oci_recovery_protected_database" "test_protected_database" {
  #Required
  compartment_id       = var.compartment_id
  db_unique_name       = var.protected_database_db_unique_name
  display_name         = var.protected_database_display_name
  password             = var.protected_database_password
  protection_policy_id = var.protection_policy_id
  recovery_service_subnets {
    #Required
    recovery_service_subnet_id = var.recovery_service_subnet_id
  }

  #Optional
  database_id          = var.database_id
  deletion_schedule    = "DELETE_AFTER_72_HOURS"
  database_size        = var.protected_database_database_size
  freeform_tags        = var.protected_database_freeform_tags
  is_redo_logs_shipped = var.protected_database_is_redo_logs_shipped
}

data "oci_recovery_protected_databases" "test_protected_databases" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name               = var.protected_database_display_name
  id                         = var.protected_database_id
  protection_policy_id       = var.protection_policy_id
  recovery_service_subnet_id = var.recovery_service_subnet_id
  state                      = var.protected_database_state
}

data "oci_recovery_protected_database_fetch_configuration" "test_protected_database_fetch_configuration" {
  #Required
  protected_database_id = oci_recovery_protected_database.test_protected_database.id

  #Optional
  configuration_type    = "ALL"
  base64_encode_content = true
}
