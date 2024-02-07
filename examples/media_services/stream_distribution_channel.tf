// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0


resource "oci_media_services_stream_distribution_channel" "test_stream_distribution_channel" {
  #Required
  compartment_id = var.compartment_id
  display_name   = var.display_name

  #Optional
  defined_tags  = map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.defined_tags_value}")
  freeform_tags = var.freeform_tags
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

data "oci_media_services_stream_distribution_channels" "test_stream_distribution_channels" {

  #Optional
  compartment_id = var.compartment_id
  display_name   = var.display_name
  id             = var.id
  state          = var.active_state
}

