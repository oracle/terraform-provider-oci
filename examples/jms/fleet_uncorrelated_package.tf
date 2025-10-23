// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

data "oci_jms_fleet_uncorrelated_packages" "test_fleet_uncorrelated_packages" {
  #Required
  fleet_id = var.fleet_ocid

  #Optional
  application_id      = var.application_id
  managed_instance_id = var.managed_instance_ocid
  package_name        = "example"
  time_end            = var.time_end
  time_start          = var.time_start
}