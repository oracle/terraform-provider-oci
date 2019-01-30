// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

data "oci_identity_availability_domain" "ad" {
  compartment_id = "${var.tenancy_ocid}"
  ad_number      = 1
}

# Gets the boot volume attachments for each instance
data "oci_core_boot_volume_attachments" "TFBootVolumeAttachments" {
  depends_on          = ["oci_core_instance.TFInstance"]
  count               = "${var.NumInstances}"
  availability_domain = "${oci_core_instance.TFInstance.*.availability_domain[count.index]}"
  compartment_id      = "${var.compartment_ocid}"

  instance_id = "${oci_core_instance.TFInstance.*.id[count.index]}"
}

data "oci_core_instance_devices" "TFInstanceDevices" {
  count       = "${var.NumInstances}"
  instance_id = "${oci_core_instance.TFInstance.*.id[count.index]}"
}

data "oci_core_volume_backup_policies" "test_predefined_volume_backup_policies" {
  filter {
    name = "display_name"

    values = [
      "silver",
    ]
  }
}
