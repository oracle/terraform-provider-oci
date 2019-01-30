// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

data "oci_core_cross_connect_locations" "cross_connect_locations" {
  #Required
  compartment_id = "${var.compartment_ocid}"
}

output "cross_connect_locations" {
  value = "${data.oci_core_cross_connect_locations.cross_connect_locations.cross_connect_locations}"
}
