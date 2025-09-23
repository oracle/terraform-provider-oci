// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "fleet_library_managed_instance_host_name" {
  default = "hostName"
}

variable "fleet_library_managed_instance_hostname_contains" {
  default = "hostnameContains"
}

variable "fleet_library_managed_instance_library_key" {
  default = "libraryKey"
}

data "oci_jms_fleet_library_managed_instances" "test_fleet_library_managed_instances" {
  #Required
  fleet_id    = var.fleet_ocid
  library_key = var.fleet_library_managed_instance_library_key

  #Optional
  application_id      = var.application_id
  host_name           = var.fleet_library_managed_instance_host_name
  hostname_contains   = var.fleet_library_managed_instance_hostname_contains
  managed_instance_id = var.managed_instance_ocid
  time_end            = var.time_end
  time_start          = var.time_start
}