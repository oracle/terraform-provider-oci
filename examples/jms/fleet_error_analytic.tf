// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

data "oci_jms_fleet_error_analytics" "test_fleet_error_analytics" {
  compartment_id = var.compartment_ocid
  #Optional
  compartment_id_in_subtree = false
}
