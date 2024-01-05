// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_streaming_stream_pool" "stream_pool_rd" {
  #Required
  compartment_id = "${var.compartment_ocid}"
  name           = "TFStreamPoolRd"

  kafka_settings {
    #Optional
    auto_create_topics_enable = true
    log_retention_hours       = 24
    num_partitions            = 1
  }
}

resource "oci_streaming_stream" "stream_rd" {
  compartment_id     = "${var.compartment_ocid}"
  name               = "streamRd"
  partitions         = "1"
  retention_in_hours = "24"
}
