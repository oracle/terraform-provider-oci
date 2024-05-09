// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}
variable "management_station_ocid" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

# Reference OL8 software source
data "oci_os_management_hub_software_sources" "ol8_baseos_latest_x86_64" {
  arch_type            = ["X86_64"]
  availability         = ["SELECTED"]
  compartment_id       = var.compartment_id
  display_name         = "ol8_baseos_latest-x86_64"
  os_family            = ["ORACLE_LINUX_8"]
  software_source_type = ["VENDOR"]
  state                = ["ACTIVE"]
  vendor_name          = "ORACLE"
}

resource "oci_os_management_hub_management_station" "test_management_station" {
  #Required
  compartment_id = var.compartment_id
  display_name   = "displayName"
  hostname       = "hostname"
  mirror {
    #Required
    directory = "/directory"
    port      = "50001"
    sslport   = "50002"

    #Optional
    sslcert = "/etc/ssl/cert"
  }
  proxy {
    #Required
    is_enabled = "true"

    #Optional
    forward = "https://example.com/forward"
    hosts   = ["host"]
    port    = "80"
  }

  #Optional
  defined_tags  = { "Operations.CostCenter" = "42" }
  description   = "description"
  freeform_tags = { "Department" = "Finance" }

  lifecycle {
    ignore_changes = [defined_tags]
  }
}

resource "oci_os_management_hub_management_station_mirror_synchronize_management" "test_management_station_mirror_synchronize_management" {
  #Required
  management_station_id = var.management_station_ocid
  mirror_id             = data.oci_os_management_hub_software_sources.ol8_baseos_latest_x86_64.software_source_collection[0].items[0].id
}

resource "oci_os_management_hub_management_station_refresh_management" "test_management_station_refresh_management" {
  #Required
  management_station_id = var.management_station_ocid
}

resource "oci_os_management_hub_management_station_synchronize_mirrors_management" "test_management_station_synchronize_mirrors_management" {
  #Required
  management_station_id = var.management_station_ocid
  software_source_list  = [data.oci_os_management_hub_software_sources.ol8_baseos_latest_x86_64.software_source_collection[0].items[0].id]
}

data "oci_os_management_hub_management_station" "test_management_station" {
  #Required
  management_station_id = oci_os_management_hub_management_station.test_management_station.id
}

data "oci_os_management_hub_management_stations" "test_management_stations" {
  #Optional
  compartment_id = var.compartment_id
}

data "oci_os_management_hub_management_station_mirrors" "test_management_station_mirrors" {
  #Required
  management_station_id = oci_os_management_hub_management_station.test_management_station.id

  #Optional
  mirror_states         = ["SYNCED"]
  display_name          = "displayName"
  display_name_contains = "displayNameContains"
}