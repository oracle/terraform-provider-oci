// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0


variable "media_workflow_job_media_workflow_configuration_ids" {
  default = []
}

variable "media_workflow_job_parameters" {
  default = "{\"videos\":{\"inputObject\":\"inputObject.mp4\",\"outputObject\":\"outputObject.mp4\"}}"
}

variable "media_workflow_job_workflow_identifier_type" {
  default = "ID"
}

resource "oci_media_services_media_workflow_job" "test_media_workflow_job" {
  #Required
  compartment_id           = var.compartment_id
  workflow_identifier_type = var.media_workflow_job_workflow_identifier_type

  #Optional
  defined_tags                     = map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.defined_tags_value}")
  display_name                     = var.display_name
  freeform_tags                    = var.freeform_tags
  media_workflow_configuration_ids = [oci_media_services_media_workflow_configuration.test_media_workflow_configuration.id]
  media_workflow_id                = oci_media_services_media_workflow.test_media_workflow.id
  media_workflow_name              = oci_media_services_media_workflow.test_media_workflow.display_name
  parameters                       = var.media_workflow_job_parameters
}

data "oci_media_services_media_workflow_jobs" "test_media_workflow_jobs" {

  #Optional
  compartment_id    = var.compartment_id
  display_name      = var.display_name
  id                = var.id
  media_workflow_id = oci_media_services_media_workflow.test_media_workflow.id
  state             = var.accepted_state
}

