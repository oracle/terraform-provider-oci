// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "java_installation_path" {}
variable "java_installation_path_contains" {}
variable "jre_version" {}
variable "jre_vendor" {}
variable "jre_distribution" {}
variable "jre_security_status" {
  default = "UNKNOWN"
}
variable "os_family" {
  default = []
}

data "oci_jms_fleet_installation_sites" "test_fleet_installation_sites" {
  #Required
  fleet_id = var.fleet_ocid

  #Optional
  application_id      = var.application_id
  installation_path   = var.java_installation_path
  path_contains       = var.java_installation_path_contains
  jre_version         = var.jre_version
  jre_vendor          = var.jre_vendor
  jre_distribution    = var.jre_distribution
  jre_security_status = var.jre_security_status
  managed_instance_id = var.managed_instance_ocid
  os_family           = var.os_family
  time_end            = var.time_end
  time_start          = var.time_start
}

