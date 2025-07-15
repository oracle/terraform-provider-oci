// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "fleet_uncorrelated_package_application_package_name" {
  default = "packageName"
}

data "oci_jms_fleet_uncorrelated_package_applications" "test_fleet_uncorrelated_package_applications" {
  #Required
  fleet_id     = var.fleet_ocid
  package_name = var.fleet_uncorrelated_package_application_package_name

  #Optional
  application_id      = var.application_id
  managed_instance_id = var.managed_instance_ocid
  time_end            = var.time_end
  time_start          = var.time_start
}