// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0


variable "media_workflow_task_declaration_is_current" {
  default = false
}

variable "media_workflow_task_declaration_name" {
  default = "name"
}

variable "media_workflow_task_declaration_version" {
  default = 10
}

data "oci_media_services_media_workflow_task_declaration" "test_media_workflow_task_declaration" {

  #Optional
  compartment_id = var.compartment_id
  is_current     = var.media_workflow_task_declaration_is_current
  name           = var.media_workflow_task_declaration_name
  version        = var.media_workflow_task_declaration_version
}

