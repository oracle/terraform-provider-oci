// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

resource "oci_core_volume" "TFBlock0" {
  availability_domain = "${data.oci_identity_availability_domain.ad.name}"
  compartment_id      = "${var.compartment_ocid}"
  display_name        = "2TB NFS"
  size_in_gbs         = "${var.2TB}"
}

resource "oci_core_volume_attachment" "TFBlock0Attach" {
  attachment_type = "iscsi"
  compartment_id  = "${var.compartment_ocid}"
  instance_id     = "${oci_core_instance.TFInstance.id}"
  volume_id       = "${oci_core_volume.TFBlock0.id}"
}
