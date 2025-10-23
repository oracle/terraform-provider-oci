// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

data "oci_jms_fleet_library_applications" "test_fleet_library_applications" {
  #Required
  fleet_id    = var.fleet_ocid
  library_key = "example-library-key"

  #Optional
  application_id            = var.application_id
  application_name          = var.application_name
  application_name_contains = "example"
  managed_instance_id       = var.managed_instance_ocid
  time_end                  = var.time_end
  time_start                = var.time_start
}