// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0


variable "media_workflow_media_workflow_configuration_ids" {
  default = []
}

variable "media_workflow_parameters" {
  default = "{\"inputs\":{\"namespace\":\"namespace\"}}"
}

variable "media_workflow_tasks_key" {
  default = "move"
}

variable "media_workflow_tasks_parameters" {
  default = "{\"taskParameters\":[{\"bucketName\":\"inputBucket\",\"namespaceName\":\"namespaceName\",\"objectName\":\"$${/videos/inputObject}\",\"storageType\":\"objectStorage\",\"target\":\"video.mp4\"}]}"
}

variable "media_workflow_tasks_prerequisites" {
  default = []
}

variable "media_workflow_tasks_type" {
  default = "getFiles"
}

variable "media_workflow_tasks_version" {
  default = 1
}

resource "oci_media_services_media_workflow" "test_media_workflow" {
  #Required
  compartment_id = var.compartment_id
  display_name   = var.display_name

  #Optional
  defined_tags                     = map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.defined_tags_value}")
  freeform_tags                    = var.freeform_tags
  media_workflow_configuration_ids = var.media_workflow_media_workflow_configuration_ids
  parameters                       = var.media_workflow_parameters
  tasks {
    #Required
    key           = var.media_workflow_tasks_key
    parameters    = var.media_workflow_tasks_parameters
    type          = var.media_workflow_tasks_type
    version       = var.media_workflow_tasks_version

    #Optional
    prerequisites              = var.media_workflow_tasks_prerequisites
  }
  locks {
    #Required
    compartment_id = var.compartment_id
    type = var.locks_type

    #Optional
    message = var.locks_message
  }
  is_lock_override = var.is_lock_override
  lifecycle {
    ignore_changes = [defined_tags, locks, is_lock_override]
  }
}

data "oci_media_services_media_workflows" "test_media_workflows" {

  #Optional
  compartment_id = var.compartment_id
  display_name   = var.display_name
  id             = var.id
  state          = var.active_state
}

