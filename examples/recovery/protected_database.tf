// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

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


resource "oci_recovery_protected_database" "test_protected_database" {
  #Required
  compartment_id       = var.compartment_id
  db_unique_name       = var.protected_database_db_unique_name
  display_name         = var.protected_database_display_name
  password             = var.protected_database_password
  protection_policy_id = oci_recovery_protection_policy.test_protection_policy.id
  recovery_service_subnets {
    #Required
    recovery_service_subnet_id = oci_recovery_recovery_service_subnet.test_recovery_service_subnet.id
  }

  #Optional
  database_id          = var.database_id
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
  protection_policy_id       = oci_recovery_protection_policy.test_protection_policy.id
  recovery_service_subnet_id = oci_recovery_recovery_service_subnet.test_recovery_service_subnet.id
  state                      = var.protected_database_state
}

