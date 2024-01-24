// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "workspace_project_description" {
  default = "description1"
}

variable "workspace_project_fields" {
  default = []
}

variable "workspace_project_identifier" {
  default = ["TESTWORKSPACEPROJECT"]
}

variable "workspace_project_key" {
}

variable "workspace_project_model_version" {
  default = "20220913"
}

variable "workspace_project_name" {
  default = "TestWorkspaceProject"
}

variable "workspace_project_name_contains" {
  default = "TestWorkspaceProject"
}

variable "workspace_project_object_status" {
  default = 0
}

variable "workspace_project_registry_metadata_aggregator_key" {
}

variable "workspace_project_registry_metadata_is_favorite" {
  default = false
}

variable "workspace_project_registry_metadata_key" {
}

variable "workspace_project_registry_metadata_labels" {
  default = []
}

variable "workspace_project_registry_metadata_registry_version" {
  default = 0
}

variable "compartment_ocid" {
}

resource "oci_dataintegration_workspace" "test_workspace" {
  #Required
  display_name = "TfTestWorkspace"
  compartment_id = var.compartment_ocid
  is_private_network_enabled = false
}


resource "oci_dataintegration_workspace_project" "test_workspace_project" {
  #Required
  identifier   = element(var.workspace_project_identifier, 0)
  name         = var.workspace_project_name
  workspace_id = oci_dataintegration_workspace.test_workspace.id

  #Optional
  description   = var.workspace_project_description
  key           = var.workspace_project_key
  model_version = var.workspace_project_model_version
  object_status = var.workspace_project_object_status
  registry_metadata {

    #Optional
    aggregator_key   = var.workspace_project_registry_metadata_aggregator_key
    is_favorite      = var.workspace_project_registry_metadata_is_favorite
    key              = var.workspace_project_registry_metadata_key
    labels           = var.workspace_project_registry_metadata_labels
    registry_version = var.workspace_project_registry_metadata_registry_version
  }
}

data "oci_dataintegration_workspace_projects" "test_workspace_projects" {
  #Required
  workspace_id = oci_dataintegration_workspace.test_workspace.id

  #Optional
  fields        = var.workspace_project_fields
  identifier    = var.workspace_project_identifier
  name          = var.workspace_project_name
  name_contains = var.workspace_project_name_contains
}

data "oci_dataintegration_workspace_project" "test_workspace_project" {
  #Required
  workspace_id = oci_dataintegration_workspace.test_workspace.id
  project_key           = oci_dataintegration_workspace_project.test_workspace_project.key

}