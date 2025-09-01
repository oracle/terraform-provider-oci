// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "java_migration_analysi_id" {
  default = "id"
}

variable "java_migration_analysi_project_name" {
  default = "project_name"
}


data "oci_jms_utils_java_migration_analysis" "test_java_migration_analysis" {

  #Optional
  analysis_project_name = var.java_migration_analysi_project_name
  compartment_id        = var.tenancy_ocid
  id                    = var.java_migration_analysi_id
}
