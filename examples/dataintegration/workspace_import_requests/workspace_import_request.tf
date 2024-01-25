// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0


variable "workspace_import_request_are_data_asset_references_included" {
  default = false
}

variable "workspace_import_request_bucket" {
  default = "bucket"
}

variable "workspace_import_request_file_name" {
  default = "MyExportObjects.zip"
}

variable "workspace_import_request_import_conflict_resolution_duplicate_prefix" {
  default = "duplicatePrefix"
}

variable "workspace_import_request_import_conflict_resolution_duplicate_suffix" {
  default = "duplicateSuffix"
}

variable "workspace_import_request_import_conflict_resolution_import_conflict_resolution_type" {
  default = "REPLACE"
}

variable "workspace_import_request_import_status" {
  default = "SUCCESSFUL"
}

variable "workspace_import_request_name" {
  default = "name"
}

variable "workspace_import_request_object_key_for_import" {
  default = ""
}

variable "workspace_import_request_object_storage_region" {
  default = "us-ashburn-1"
}

variable "workspace_import_request_projection" {
  default = "SUMMARY"
}

variable "workspace_import_request_time_ended_in_millis" {
}

variable "workspace_import_request_time_started_in_millis" {
}

variable "workspace_id" {
}

variable "tenancy_id" {
}

resource "oci_dataintegration_workspace_import_request" "test_workspace_import_request" {
  #Required
  bucket       = var.workspace_import_request_bucket
  file_name    = var.workspace_import_request_file_name
  workspace_id = var.workspace_id

  #Optional
  are_data_asset_references_included = var.workspace_import_request_are_data_asset_references_included
  import_conflict_resolution {
    #Required
    import_conflict_resolution_type = var.workspace_import_request_import_conflict_resolution_import_conflict_resolution_type

    #Optional
    duplicate_prefix = var.workspace_import_request_import_conflict_resolution_duplicate_prefix
    duplicate_suffix = var.workspace_import_request_import_conflict_resolution_duplicate_suffix
  }
  object_key_for_import     = var.workspace_import_request_object_key_for_import
  object_storage_region     = var.workspace_import_request_object_storage_region
  object_storage_tenancy_id = var.tenancy_id
}

data "oci_dataintegration_workspace_import_requests" "test_workspace_import_requests" {
  #Required
  workspace_id = var.workspace_id

  #Optional
  import_status          = var.workspace_import_request_import_status
  name                   = var.workspace_import_request_name
  projection             = var.workspace_import_request_projection
  time_ended_in_millis   = var.workspace_import_request_time_ended_in_millis
  time_started_in_millis = var.workspace_import_request_time_started_in_millis
}

