// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_file_storage_mount_target" "my_mount_target_1" {
  #Required
  availability_domain = "${data.oci_identity_availability_domain.ad.name}"
  compartment_id      = "${var.compartment_ocid}"
  subnet_id           = "${oci_core_subnet.my_subnet.id}"

  #Optional
  display_name = "${var.mount_target_1_display_name}"
  defined_tags = "${map("example-tag-namespace-all.example-tag", "value")}"

  freeform_tags = {
    "Department" = "Finance"
  }

  nsg_ids = ["${oci_core_network_security_group.test_network_security_group.id}"]
}

resource "oci_file_storage_mount_target" "my_mount_target_2" {
  #Required
  availability_domain = "${data.oci_identity_availability_domain.ad.name}"
  compartment_id      = "${var.compartment_ocid}"
  subnet_id           = "${oci_core_subnet.my_subnet.id}"

  #Optional
  display_name = "${var.mount_target_2_display_name}"
  defined_tags = "${map("example-tag-namespace-all.example-tag", "value")}"

  freeform_tags = {
    "Department" = "Accounting"
  }

  nsg_ids = ["${oci_core_network_security_group.test_network_security_group.id}"]
}

# Use export_set.tf config to update the size for a mount target

