// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "workspace_application_patch_application_key" {
  default = "applicationKey"
}

variable "workspace_application_patch_description" {
  default = "description"
}

variable "workspace_application_patch_fields" {
  default = []
}

variable "workspace_application_patch_identifier" {
  default = "NAME"
}

variable "workspace_application_patch_name" {
  default = "Name"
}

variable "workspace_application_patch_object_keys" {
  default = []
}

variable "workspace_application_patch_object_status" {
  default = 8
}

variable "workspace_application_patch_patch_type" {
  default = "PUBLISH"
}

variable "workspace_application_patch_registry_metadata_aggregator_key" {
  default = "aggregatorKey"
}

variable "workspace_application_patch_registry_metadata_is_favorite" {
  default = false
}

variable "workspace_application_patch_registry_metadata_labels" {
  default = []
}

variable "workspace_id" {
}

resource "oci_dataintegration_workspace_application_patch" "test_workspace_application_patch" {
  #Required
  application_key = var.workspace_application_patch_application_key
  identifier      = var.workspace_application_patch_identifier
  name            = var.workspace_application_patch_name
  object_keys     = var.workspace_application_patch_object_keys
  patch_type      = var.workspace_application_patch_patch_type
  workspace_id    = var.workspace_id

  #Optional
  description   = var.workspace_application_patch_description
  registry_metadata {

    #Optional
    aggregator_key   = var.workspace_application_patch_registry_metadata_aggregator_key
    is_favorite      = var.workspace_application_patch_registry_metadata_is_favorite
    labels           = var.workspace_application_patch_registry_metadata_labels
  }
}

data "oci_dataintegration_workspace_application_patches" "test_workspace_application_patches" {
  #Required
  application_key = var.workspace_application_patch_application_key
  workspace_id    = var.workspace_id

  #Optional
  fields     = var.workspace_application_patch_fields
  identifier = var.workspace_application_patch_identifier
  name       = var.workspace_application_patch_name
}