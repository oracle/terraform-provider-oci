// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "batch_task_environment_defined_tags_value" {
  default = "value"
}

variable "batch_task_environment_description" {
  default = "description"
}

variable "batch_task_environment_display_name" {
  default = "displayName"
}

variable "batch_task_environment_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "batch_task_environment_id" {
  default = "id"
}

variable "batch_task_environment_image_url" {
  default = "imageUrl"
}

variable "batch_task_environment_security_context_fs_group" {
  default = 10
}

variable "batch_task_environment_security_context_run_as_group" {
  default = 10
}

variable "batch_task_environment_security_context_run_as_user" {
  default = 10
}

variable "batch_task_environment_state" {
  default = "AVAILABLE"
}

variable "batch_task_environment_volumes_local_mount_directory_path" {
  default = "localMountDirectoryPath"
}

variable "batch_task_environment_volumes_mount_target_export_path" {
  default = "mountTargetExportPath"
}

variable "batch_task_environment_volumes_mount_target_fqdn" {
  default = "mountTargetFqdn"
}

variable "batch_task_environment_volumes_name" {
  default = "name"
}

variable "batch_task_environment_volumes_type" {
  default = "NFS"
}

variable "batch_task_environment_working_directory" {
  default = "workingDirectory"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_batch_batch_task_environment" "test_batch_task_environment" {
  #Required
  compartment_id = var.compartment_id
  image_url      = var.batch_task_environment_image_url

  #Optional
  defined_tags  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.batch_task_environment_defined_tags_value)
  description   = var.batch_task_environment_description
  display_name  = var.batch_task_environment_display_name
  freeform_tags = var.batch_task_environment_freeform_tags
  security_context {

    #Optional
    fs_group     = var.batch_task_environment_security_context_fs_group
    run_as_group = var.batch_task_environment_security_context_run_as_group
    run_as_user  = var.batch_task_environment_security_context_run_as_user
  }
  volumes {
    #Required
    local_mount_directory_path = var.batch_task_environment_volumes_local_mount_directory_path
    mount_target_export_path   = var.batch_task_environment_volumes_mount_target_export_path
    mount_target_fqdn          = var.batch_task_environment_volumes_mount_target_fqdn
    name                       = var.batch_task_environment_volumes_name
    type                       = var.batch_task_environment_volumes_type
  }
  working_directory = var.batch_task_environment_working_directory
}

data "oci_batch_batch_task_environments" "test_batch_task_environments" {

  #Optional
  compartment_id = var.compartment_id
  display_name   = var.batch_task_environment_display_name
  id             = var.batch_task_environment_id
  state          = var.batch_task_environment_state
}

