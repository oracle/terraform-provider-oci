// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "workspace_folder_aggregator_key" {
}

variable "workspace_folder_description" {
  default = "description1"
}

variable "workspace_folder_fields" {
  default = []
}

variable "workspace_folder_identifier" {
  default = ["IDENTIFIER1"]
}

variable "workspace_folder_key" {
}

variable "workspace_folder_model_version" {
  default = "20220913"
}

variable "workspace_folder_name" {
  default = "WorkspaceFolderName"
}

variable "workspace_folder_name_contains" {
  default = "WorkspaceFolderName"
}

variable "workspace_folder_object_status" {
  default = 0
}

variable "workspace_folder_registry_metadata_aggregator_key" {
}

variable "workspace_folder_registry_metadata_is_favorite" {
  default = false
}

variable "workspace_folder_registry_metadata_key" {
}

variable "workspace_folder_registry_metadata_labels" {
  default = []
}

variable "workspace_folder_registry_metadata_registry_version" {
  default = 0
}

variable "compartment_ocid" {
}

variable "workspace_project_key" {
}

resource "oci_dataintegration_workspace" "test_workspace" {
  #Required
  display_name = "TfTestWorkspace"
  compartment_id = var.compartment_ocid
  is_private_network_enabled = false
}

resource "oci_dataintegration_workspace_project" "test_workspace_project" {
  #Required
  identifier   = "IDENTIFIER11"
  name         = "TestWorkspaceProject"
  workspace_id = oci_dataintegration_workspace.test_workspace.id

  #Optional
  key          = var.workspace_project_key
}

resource "oci_dataintegration_workspace_folder" "test_workspace_folder" {
  #Required
  identifier = element(var.workspace_folder_identifier, 0)
  name       = var.workspace_folder_name
  registry_metadata {

    #Optional
    aggregator_key   = oci_dataintegration_workspace_project.test_workspace_project.key
    is_favorite      = var.workspace_folder_registry_metadata_is_favorite
    key              = var.workspace_folder_registry_metadata_key
    labels           = var.workspace_folder_registry_metadata_labels
    registry_version = var.workspace_folder_registry_metadata_registry_version
  }
  workspace_id = oci_dataintegration_workspace.test_workspace.id

  #Optional
  description   = var.workspace_folder_description
  key           = var.workspace_folder_key
  model_version = var.workspace_folder_model_version
  object_status = var.workspace_folder_object_status
}

data "oci_dataintegration_workspace_folders" "test_workspace_folders" {
  #Required
  workspace_id = oci_dataintegration_workspace.test_workspace.id

  #Optional
  aggregator_key = var.workspace_folder_aggregator_key
  fields         = var.workspace_folder_fields
  identifier     = var.workspace_folder_identifier
  name           = var.workspace_folder_name
  name_contains  = var.workspace_folder_name_contains
}

data "oci_dataintegration_workspace_folder" "test_workspace_folder" {
  #Required
  workspace_id = oci_dataintegration_workspace.test_workspace.id
  folder_key          = oci_dataintegration_workspace_folder.test_workspace_folder.key

}
