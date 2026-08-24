// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "job_defined_tags_value" {
  default = "value"
}

variable "job_display_name" {
  default = "displayName"
}

variable "job_freeform_tags" {
  default = "freeformTags"
}

variable "job_job_id_not_equal_to" {
  default = "jobIdNotEqualTo"
}

variable "job_state" {
  default = "AVAILABLE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_database_migration_job" "test_job" {
  #Required
  job_id = oci_database_migration_job.test_job.id

  #Optional
  defined_tags  = tomap({ oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name = var.job_defined_tags_value })
  display_name  = var.job_display_name
  freeform_tags = var.job_freeform_tags
}

data "oci_database_migration_jobs" "test_jobs" {
  #Required
  migration_id = oci_database_migration_migration.test_migration.id

  #Optional
  display_name        = var.job_display_name
  job_id_not_equal_to = var.job_job_id_not_equal_to
  state               = var.job_state
}

