// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

data "oci_core_cross_connect_locations" "cross_connect_locations" {
  #Required
  compartment_id = var.compartment_ocid
}

output "cross_connect_locations" {
  value = data.oci_core_cross_connect_locations.cross_connect_locations.cross_connect_locations
}

