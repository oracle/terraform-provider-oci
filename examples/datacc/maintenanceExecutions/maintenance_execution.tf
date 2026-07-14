// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "maintenance_execution_display_name" {
  default = "displayName"
}

variable "maintenance_execution_maintenance_subtype" {
  default = "YEARLY"
}

variable "maintenance_execution_maintenance_type" {
  default = "PLANNED"
}

variable "maintenance_execution_state" {
  default = "AVAILABLE"
}

variable "maintenance_execution_target_resource_type" {
  default = "DB_CC_INFRASTRUCTURE"
}

variable "maintenance_execution_time_accepted_greater_than_or_equal_to" {
  default = "timeAcceptedGreaterThanOrEqualTo"
}

variable "maintenance_execution_time_accepted_less_than_or_equal_to" {
  default = "timeAcceptedLessThanOrEqualTo"
}

variable "maintenance_execution_type" {
  default = "NOTIFY"
}

variable "maintenance_run_id" {
  default = "maintenanceRunId"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_datacc_maintenance_executions" "test_maintenance_executions" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name                           = var.maintenance_execution_display_name
  infrastructure_id                      = oci_datacc_infrastructure.test_infrastructure.id
  maintenance_run_id                     = var.maintenance_run_id
  maintenance_subtype                    = var.maintenance_execution_maintenance_subtype
  maintenance_type                       = var.maintenance_execution_maintenance_type
  state                                  = var.maintenance_execution_state
  target_resource_type                   = var.maintenance_execution_target_resource_type
  time_accepted_greater_than_or_equal_to = var.maintenance_execution_time_accepted_greater_than_or_equal_to
  time_accepted_less_than_or_equal_to    = var.maintenance_execution_time_accepted_less_than_or_equal_to
  type                                   = var.maintenance_execution_type
}

