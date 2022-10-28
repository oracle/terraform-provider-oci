// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0


variable "stream_packaging_config_encryption_algorithm" {
  default = "AES"
}

variable "stream_packaging_config_segment_time_in_seconds" {
  default = 10
}

variable "stream_packaging_config_stream_packaging_format" {
  default = "HLS"
}

resource "oci_media_services_stream_packaging_config" "test_stream_packaging_config" {
  #Required
  display_name            = var.display_name
  distribution_channel_id = oci_media_services_stream_distribution_channel.test_stream_distribution_channel.id
  segment_time_in_seconds = var.stream_packaging_config_segment_time_in_seconds
  stream_packaging_format = var.stream_packaging_config_stream_packaging_format

  #Optional
  defined_tags = map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.defined_tags_value}")
  encryption {
    #Required
    algorithm = var.stream_packaging_config_encryption_algorithm

    #Optional
    kms_key_id = data.oci_kms_keys.test_keys_dependency_RSA.keys[0].id
  }
  freeform_tags = var.freeform_tags
}

data "oci_media_services_stream_packaging_configs" "test_stream_packaging_configs" {
  #Required
  distribution_channel_id = oci_media_services_stream_distribution_channel.test_stream_distribution_channel.id

  #Optional
  display_name               = var.display_name
  state                      = var.active_state
  stream_packaging_config_id = oci_media_services_stream_packaging_config.test_stream_packaging_config.id
}

