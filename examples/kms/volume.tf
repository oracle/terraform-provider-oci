// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_core_volume" "my_volume" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_id
  display_name        = "-tf-volume"
  size_in_gbs         = var.volume_size
  kms_key_id          = oci_kms_key.test_key.id
}

