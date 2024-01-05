// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
}

variable "compartment_ocid" {
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

/*
resource "oci_streaming_stream" "stream" {
  compartment_id     = var.compartment_ocid
  name               = "stream1"
  partitions         = "1"
  retention_in_hours = "24"
}

data "oci_streaming_stream" "stream1" {
  stream_id = oci_streaming_stream.stream.id
}

# Output the result
output "stream" {
  value = <<EOF

  id = data.oci_streaming_stream.stream1.id
  compartment_id = data.oci_streaming_stream.stream1.compartment_id
  messages_endpoint = data.oci_streaming_stream.stream1.messages_endpoint
  name = data.oci_streaming_stream.stream1.name
  partitions = data.oci_streaming_stream.stream1.partitions
  retention_in_hours = data.oci_streaming_stream.stream1.retention_in_hours
  state = data.oci_streaming_stream.stream1.state
  time_created = data.oci_streaming_stream.stream1.time_created
EOF

  # This value is not always present--when state is FAILED it may contain an explanation.
  #lifecycle_state_details = data.oci_streaming_stream.stream1.lifecycle_state_details
}

data "oci_streaming_streams" "streams" {
  compartment_id = oci_streaming_stream.stream.compartment_id

  # optional
  state = "ACTIVE"

  //  id    = oci_streaming_stream.stream.id
  //  name  = oci_streaming_stream.stream.name
}

resource "oci_streaming_connect_harness" "test_connect_harness" {
  #Required
  compartment_id = var.compartment_ocid
  name           = "TFConnectHarness"
}

data "oci_streaming_connect_harness" "test_connect_harness" {
  connect_harness_id = oci_streaming_connect_harness.test_connect_harness.id
}
*/

resource "oci_streaming_stream_pool" "test_stream_pool" {
  #Required
  compartment_id = var.compartment_ocid
  name           = "TFStreamPool"

  #Optional
  private_endpoint_settings {
    nsg_ids             = [oci_core_network_security_group.test_nsg.id]
    private_endpoint_ip = "10.0.0.5"
    subnet_id           = oci_core_subnet.test_subnet.id
  }

  kafka_settings {
    #Optional
    auto_create_topics_enable = true
    log_retention_hours       = 24
    num_partitions            = 1
  }
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "testvcn"
  dns_label      = "dnslabel"
}

resource "oci_core_subnet" "test_subnet" {
  cidr_block     = "10.0.0.0/24"
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.test_vcn.id
}

resource "oci_core_network_security_group" "test_nsg" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.test_vcn.id
}

