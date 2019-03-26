// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

resource "oci_core_instance" "my_instance" {
  availability_domain = "${data.oci_identity_availability_domain.ad.name}"
  compartment_id      = "${var.compartment_id}"
  display_name        = "my instance with FSS access"
  hostname_label      = "myinstance"
  shape               = "${var.instance_shape}"
  subnet_id           = "${oci_core_subnet.my_subnet.id}"

  metadata {
    ssh_authorized_keys = "${var.ssh_public_key}"
  }

  source_details {
    source_type = "image"
    source_id   = "${var.instance_image_ocid[var.region]}"
    kms_key_id  = "${oci_kms_key.test_key.id}"
  }

  timeouts {
    create = "60m"
  }
}
