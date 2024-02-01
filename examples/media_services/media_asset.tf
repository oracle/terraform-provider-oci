// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0


variable "media_asset_bucket" {
  default = "bucket"
}

variable "media_asset_media_asset_tags_type" {
  default = "USER"
}

variable "media_asset_media_asset_tags_value" {
  default = "value"
}

variable "media_asset_metadata_metadata" {
  default = "{\"some\":\"json\"}"
}

variable "media_asset_namespace" {
  default = "namespace"
}

variable "media_asset_object" {
  default = "object"
}

variable "media_asset_object_etag" {
  default = "objectEtag"
}

variable "media_asset_segment_range_end_index" {
  default = 10
}

variable "media_asset_segment_range_start_index" {
  default = 10
}

variable "media_asset_source_media_workflow_version" {
  default = 10
}

resource "oci_media_services_media_asset" "test_media_asset" {
  #Required
  compartment_id = var.compartment_id
  type           = var.media_asset_type

  #Optional
  bucket                = var.media_asset_bucket
  defined_tags          = map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.defined_tags_value}")
  display_name          = var.display_name
  freeform_tags         = var.freeform_tags
  media_asset_tags {
    #Required
    value = var.media_asset_media_asset_tags_value

    #Optional
    type = var.media_asset_media_asset_tags_type
  }
  media_workflow_job_id = oci_media_services_media_workflow_job.test_media_workflow_job.id
  metadata {
    #Required
    metadata = var.media_asset_metadata_metadata
  }
  namespace                     = var.media_asset_namespace
  object                        = var.media_asset_object
  object_etag                   = var.media_asset_object_etag
  segment_range_end_index       = var.media_asset_segment_range_end_index
  segment_range_start_index     = var.media_asset_segment_range_start_index
  source_media_workflow_id      = oci_media_services_media_workflow.test_media_workflow.id
  source_media_workflow_version = var.media_asset_source_media_workflow_version
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

data "oci_media_services_media_assets" "test_media_assets" {

  #Optional
  bucket                        = var.media_asset_bucket
  compartment_id                = var.compartment_id
  display_name                  = var.display_name
  media_workflow_job_id         = oci_media_services_media_workflow_job.test_media_workflow_job.id
  object                        = var.media_asset_object
  source_media_workflow_id      = oci_media_services_media_workflow.test_media_workflow.id
  source_media_workflow_version = var.media_asset_source_media_workflow_version
  state                         = var.media_asset_state
  type                          = var.media_asset_type
}

