// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

resource "oci_file_storage_mount_target" "my_mount_target_1" {
  #Required
  availability_domain = "${data.oci_identity_availability_domain.ad.name}"
  compartment_id      = "${var.compartment_ocid}"
  subnet_id           = "${oci_core_subnet.my_subnet.id}"

  #Optional
  display_name = "${var.mount_target_1_display_name}"
}

resource "oci_file_storage_mount_target" "my_mount_target_2" {
  #Required
  availability_domain = "${data.oci_identity_availability_domain.ad.name}"
  compartment_id      = "${var.compartment_ocid}"
  subnet_id           = "${oci_core_subnet.my_subnet.id}"

  #Optional
  display_name = "${var.mount_target_2_display_name}"
}
