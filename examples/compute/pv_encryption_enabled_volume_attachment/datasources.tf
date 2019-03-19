// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

data "oci_identity_availability_domain" "ad" {
  compartment_id = "${var.tenancy_ocid}"
  ad_number      = 1
}

# Gets a list of all Oracle Linux 7.5 images that support a given Instance shape
data "oci_core_images" "TFSupportedShapeImages" {
  compartment_id   = "${var.tenancy_ocid}"
  shape            = "${var.instance_shape}"
  operating_system = "${var.ImageOS}"

  filter {
    name   = "launch_options.is_pv_encryption_in_transit_enabled"
    values = ["true"]
  }
}
