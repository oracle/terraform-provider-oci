// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0


variable "media_workflow_configuration_parameters" {
  default = "{\"storage\":{\"inputbucket\":\"myinputbucket\",\"outputbucket\":\"myoutputBucket\"}}"
}

resource "oci_media_services_media_workflow_configuration" "test_media_workflow_configuration" {
  #Required
  compartment_id = var.compartment_id
  display_name   = var.display_name
  parameters     = var.media_workflow_configuration_parameters

  #Optional
  defined_tags  = map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.defined_tags_value}")
  freeform_tags = var.freeform_tags
}

data "oci_media_services_media_workflow_configurations" "test_media_workflow_configurations" {

  #Optional
  compartment_id = var.compartment_id
  display_name   = var.display_name
  id             = var.id
  state          = var.active_state
}

