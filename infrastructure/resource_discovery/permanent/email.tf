// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_email_sender" "email_sender_rd" {
  #Required
  compartment_id = "${var.compartment_ocid}"
  email_address  = "${var.sender_email_address}"

  timeouts {
    create = "10m"
  }
}

resource "oci_email_suppression" "email_suppression_rd" {
  #Required
  compartment_id = "${var.tenancy_ocid}"
  email_address  = "${var.suppression_email_address}"
}
