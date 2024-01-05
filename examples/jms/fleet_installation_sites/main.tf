// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "fleet_id" {
  default = "example-fleet-id"
}
variable "application_id" {}
variable "fleet_installation_site_installation_path" {}
variable "fleet_installation_site_jre_distribution" {}
variable "fleet_installation_site_jre_security_status" {
  default = "UNKNOWN"
}
variable "fleet_installation_site_jre_vendor" {}
variable "fleet_installation_site_jre_version" {}
variable "managed_instance_id"{}
variable "fleet_installation_site_os_family" {
  default = []
}
variable "fleet_installation_site_path_contains" {}
variable "fleet_installation_site_time_start" {}
variable "fleet_installation_site_time_end" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_jms_fleet_installation_sites" "test_fleet_installation_sites" {
  #Required
  fleet_id = var.fleet_id

  #Optional
  application_id      = var.application_id
  installation_path   = var.fleet_installation_site_installation_path
  jre_distribution    = var.fleet_installation_site_jre_distribution
  jre_security_status = var.fleet_installation_site_jre_security_status
  jre_vendor          = var.fleet_installation_site_jre_vendor
  jre_version         = var.fleet_installation_site_jre_version
  managed_instance_id = var.managed_instance_id
  os_family           = var.fleet_installation_site_os_family
  path_contains       = var.fleet_installation_site_path_contains
  time_end            = var.fleet_installation_site_time_end
  time_start          = var.fleet_installation_site_time_start
}

