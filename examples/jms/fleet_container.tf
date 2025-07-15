// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "fleet_container_display_name" {
  default = "displayName"
}

variable "fleet_container_jre_security_status" {
  default = "EARLY_ACCESS"
}

variable "fleet_container_jre_version" {
  default = "jreVersion"
}

variable "fleet_container_time_started_greater_than_or_equal_to" {
  default = "timeStartedGreaterThanOrEqualTo"
}

variable "fleet_container_time_started_less_than_or_equal_to" {
  default = "timeStartedLessThanOrEqualTo"
}

data "oci_jms_fleet_containers" "test_fleet_containers" {
  #Required
  fleet_id = var.fleet_ocid

  #Optional
  application_name                      = var.application_name
  display_name                          = var.fleet_container_display_name
  jre_security_status                   = var.fleet_container_jre_security_status
  jre_version                           = var.fleet_container_jre_version
  managed_instance_id                   = var.managed_instance_ocid
  time_started_greater_than_or_equal_to = var.time_end
  time_started_less_than_or_equal_to    = var.time_start
}