// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "workspace_application_defined_tags_value" {
  default = "value"
}

variable "workspace_application_description" {
  default = "description"
}

variable "workspace_application_display_name" {
  default = "displayName"
}

variable "workspace_application_fields" {
  default = []
}

variable "workspace_application_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "workspace_application_identifier" {
  default = "identifier"
}

variable "workspace_application_key" {
}

variable "workspace_application_model_type" {
  default = "INTEGRATION_APPLICATION"
}

variable "workspace_application_model_version" {
  default = "20230426"
}

variable "workspace_application_name" {
  default = "name"
}

variable "workspace_application_name_contains" {
  default = "name"
}

variable "workspace_application_object_status" {
  default = 8
}

variable "workspace_application_registry_metadata_aggregator_key" {
}

variable "workspace_application_registry_metadata_is_favorite" {
  default = false
}

variable "workspace_application_registry_metadata_key" {
}

variable "workspace_application_registry_metadata_labels" {
  default = []
}

variable "workspace_application_registry_metadata_registry_version" {
}

variable "workspace_application_source_application_info_application_key" {
}

variable "workspace_application_source_application_info_copy_type" {
}

variable "workspace_application_state" {
}

variable "compartment_ocid" {
}

resource "oci_dataintegration_workspace" "test_workspace" {
  #Required
  display_name = "TfTestWorkspace"
  compartment_id = var.compartment_ocid
  is_private_network_enabled = false
}

resource "oci_dataintegration_workspace_application" "test_workspace_application" {
  #Required
  identifier   = var.workspace_application_identifier
  model_type   = var.workspace_application_model_type
  name         = var.workspace_application_name
  workspace_id = oci_dataintegration_workspace.test_workspace.id

  #Optional
  defined_tags  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.workspace_application_defined_tags_value)
  description   = var.workspace_application_description
  display_name  = var.workspace_application_display_name
  freeform_tags = var.workspace_application_freeform_tags
  key           = var.workspace_application_key
  model_version = var.workspace_application_model_version
  object_status = var.workspace_application_object_status
  registry_metadata {

    #Optional
    aggregator_key   = var.workspace_application_registry_metadata_aggregator_key
    is_favorite      = var.workspace_application_registry_metadata_is_favorite
    key              = var.workspace_application_registry_metadata_key
    labels           = var.workspace_application_registry_metadata_labels
    registry_version = var.workspace_application_registry_metadata_registry_version
  }
  source_application_info {

    #Optional
    application_key = var.workspace_application_source_application_info_application_key
    copy_type       = var.workspace_application_source_application_info_copy_type
    workspace_id    = oci_dataintegration_workspace.test_workspace.id
  }
  state = var.workspace_application_state
}

data "oci_dataintegration_workspace_applications" "test_workspace_applications" {
  #Required
  workspace_id = oci_dataintegration_workspace.test_workspace.id

  #Optional
  fields        = var.workspace_application_fields
  identifier    = var.workspace_application_identifier
  name          = var.workspace_application_name
  name_contains = var.workspace_application_name_contains
}
