// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

data "oci_jms_fleet_errors" "test_fleet_errors" {

  compartment_id = var.compartment_ocid
  #Optional
  compartment_id_in_subtree                = false
  fleet_id                                 = var.fleet_ocid
  time_first_seen_greater_than_or_equal_to = var.time_start
  time_first_seen_less_than_or_equal_to    = var.time_end
  time_last_seen_greater_than_or_equal_to  = var.time_start
  time_last_seen_less_than_or_equal_to     = var.time_end
}
