// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

data "oci_identity_availability_domain" "ad1" {
  compartment_id = "${var.compartment_ocid}"
  ad_number      = 1
}

data "oci_identity_availability_domain" "ad2" {
  compartment_id = "${var.compartment_ocid}"
  ad_number      = 2
}
