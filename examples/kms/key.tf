// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_kms_key" "test_key" {
  #Required
  compartment_id      = var.compartment_id
  display_name        = var.key_display_name
  management_endpoint = data.oci_kms_vault.test_vault.management_endpoint

  key_shape {
    #Required
    algorithm = var.key_key_shape_algorithm
    length    = var.key_key_shape_length
  }

}


