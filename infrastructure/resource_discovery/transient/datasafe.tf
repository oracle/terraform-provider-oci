// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_data_safe_data_safe_private_endpoint" "test_data_safe_private_endpoint_rd" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "datasafePrivateEndpointRD"
  subnet_id      = "${oci_core_subnet.tf_subnet.id}"
  vcn_id         = "${oci_core_vcn.vcn2_rd.id}"
  description    = "description"
}
