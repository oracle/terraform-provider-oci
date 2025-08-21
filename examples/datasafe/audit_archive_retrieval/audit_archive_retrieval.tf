// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "target_id" {}
variable "start_date" {}
variable "end_date" {}

variable "audit_archive_retrieval_access_level" {
  default = "RESTRICTED"
}

variable "audit_archive_retrieval_compartment_id_in_subtree" {
  default = false
}

variable "audit_archive_retrieval_defined_tags_value" {
  default = "value"
}

variable "audit_archive_retrieval_description" {
  default = "Archival example test"
}

variable "audit_archive_retrieval_display_name" {
  default = "Audit_archival"
}

variable "audit_archive_retrieval_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "audit_archive_retrieval_state" {
  default = "ACTIVE"
}

variable "audit_archive_retrieval_start_date" {
  default = "2024-02-01T00:00:00Z"
}

variable "audit_archive_retrieval_end_date" {
  default = "2024-03-01T00:00:00Z"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_data_safe_audit_archive_retrieval" "test_audit_archive_retrieval" {
  #Required
  compartment_id = var.compartment_ocid
  target_id      = var.target_id
  start_date     = var.audit_archive_retrieval_start_date
  end_date       = var.audit_archive_retrieval_end_date

  #Optional
  description    = var.audit_archive_retrieval_description
  display_name   = var.audit_archive_retrieval_display_name
  freeform_tags  = var.audit_archive_retrieval_freeform_tags
}

data "oci_data_safe_audit_archive_retrieval" "test_audit_archive_retrieval" {
  #Optional
  audit_archive_retrieval_id           = oci_data_safe_audit_archive_retrieval.test_audit_archive_retrieval.id
}

data "oci_data_safe_audit_archive_retrievals" "test_audit_archive_retrievals" {
  #Required
  compartment_id = var.compartment_ocid
}

