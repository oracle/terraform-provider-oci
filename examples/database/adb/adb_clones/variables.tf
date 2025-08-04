// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
}

variable "compartment_ocid" {
}

variable "autonomous_database_freeform_tags" {
  default = {
    "Department" = "Finance"
  }
}

variable "autonomous_database_license_model" {
  default = "LICENSE_INCLUDED"
}

variable "autonomous_database_db_workload" {
  default = "OLTP"
}

variable "autonomous_database_autonomous_maintenance_schedule_type" {
  default = "EARLY"
}