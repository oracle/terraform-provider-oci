// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

data "oci_jms_fleet_containers" "test_fleet_containers" {
  #Required
  fleet_id = var.fleet_ocid

  #Optional
  application_name                      = var.application_name
  display_name                          = "example"
  jre_security_status                   = "EARLY_ACCESS"
  jre_version                           = "17.0.0"
  managed_instance_id                   = var.managed_instance_ocid
  time_started_greater_than_or_equal_to = var.time_end
  time_started_less_than_or_equal_to    = var.time_start
}