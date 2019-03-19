// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

/*
 * This example file shows how to configure the oci provider to target a single region.
 */
provider "oci" {
  region           = "${var.region}"
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
}
