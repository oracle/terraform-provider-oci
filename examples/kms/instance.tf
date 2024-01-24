// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_core_instance" "my_instance" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_id
  display_name        = "my instance with FSS access"
  shape               = var.instance_shape

  metadata = {
    ssh_authorized_keys = var.ssh_public_key
  }

  create_vnic_details {
    subnet_id      = oci_core_subnet.my_subnet.id
    hostname_label = "myinstance"
  }

  source_details {
    source_type = "image"
    source_id   = var.instance_image_ocid[var.region]
    kms_key_id  = oci_kms_key.test_key.id
  }

  timeouts {
    create = "60m"
  }
}

