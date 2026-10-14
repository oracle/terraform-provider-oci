// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_database_migration_job_advisor_report_check_objects" "test_job_advisor_report_check_objects" {
  #Required
  advisor_report_check_id = oci_database_migration_advisor_report_check.test_advisor_report_check.id
  job_id                  = oci_database_migration_job.test_job.id
}

