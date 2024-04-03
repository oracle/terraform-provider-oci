// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

data "oci_email_configuration" "test_configuration" {
  #Required
  compartment_id = var.compartment_ocid
}