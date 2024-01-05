// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "workspace_export_request_are_references_included" {
  default = false
}

variable "workspace_export_request_bucket" {
  default = "bucket"
}

variable "workspace_export_request_export_status" {
  default = "SUCCESSFUL"
}

variable "workspace_export_request_file_name" {
  default = "MyExportObjects.zip"
}

variable "workspace_export_request_filters" {
  default = []
}

variable "workspace_export_request_is_object_overwrite_enabled" {
  default = true
}

variable "workspace_export_request_name" {
  default = "name"
}

variable "workspace_export_request_object_keys" {
  default = []
}

variable "workspace_export_request_object_storage_region" {
  default = "us-ashburn-1"
}

variable "workspace_export_request_projection" {
  default = "SUMMARY"
}

variable "workspace_export_request_time_ended_in_millis" {
}

variable "workspace_export_request_time_started_in_millis" {
}

variable "workspace_id" {
}

variable "tenancy_id" {
}

resource "oci_dataintegration_workspace_export_request" "test_workspace_export_request" {
  #Required
  bucket       = var.workspace_export_request_bucket
  workspace_id = var.workspace_id

  #Optional
  are_references_included     = var.workspace_export_request_are_references_included
  file_name                   = var.workspace_export_request_file_name
  filters                     = var.workspace_export_request_filters
  is_object_overwrite_enabled = var.workspace_export_request_is_object_overwrite_enabled
  object_keys                 = var.workspace_export_request_object_keys
  object_storage_region       = var.workspace_export_request_object_storage_region
  object_storage_tenancy_id   = var.tenancy_id
}

data "oci_dataintegration_workspace_export_requests" "test_workspace_export_requests" {
  #Required
  workspace_id = var.workspace_id

  #Optional
  export_status          = var.workspace_export_request_export_status
  name                   = var.workspace_export_request_name
  projection             = var.workspace_export_request_projection
  time_ended_in_millis   = var.workspace_export_request_time_ended_in_millis
  time_started_in_millis = var.workspace_export_request_time_started_in_millis
}

