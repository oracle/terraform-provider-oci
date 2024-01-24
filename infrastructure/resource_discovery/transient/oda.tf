// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_oda_oda_instance" "oda_rd" {
  compartment_id = "${var.compartment_ocid}"
  shape_name     = "DEVELOPMENT"
  description    = "test instance"
  display_name   = "OdaInstanceRd"
}
