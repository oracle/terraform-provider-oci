// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "primary_db_system_display_name" {
  description = "Primary OCI Database with PostgreSQL DB system"
  type        = string
  default     = "example-postgresqldb-primary"
}

variable "standby_db_system_display_name" {
  description = "Standby OCI Database with PostgreSQL DB system"
  type        = string
  default     = "example-postgresqldb-standby"
}

data "oci_psql_db_systems" "test_db_systems_primary" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name = var.primary_db_system_display_name
  state        = "ACTIVE"
}

data "oci_psql_db_systems" "test_db_systems_standby" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name = var.standby_db_system_display_name
  state        = "ACTIVE"
}
