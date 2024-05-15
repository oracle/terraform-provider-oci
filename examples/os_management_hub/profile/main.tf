// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}
variable "lifecycle_stage_id" {} 


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

#----------------
# Software Source
#----------------
# Assumption: "ol8_baseos_latest-x86_64" exists as a Vendor Software Source in the Tenancy 
data "oci_os_management_hub_software_sources" "ol8_baseos_latest_x86_64" {
  arch_type = ["X86_64"]
  availability = ["SELECTED"]
  compartment_id = "${var.compartment_id}"
  display_name = "ol8_baseos_latest-x86_64"
  os_family = ["ORACLE_LINUX_8"]
  software_source_type = ["VENDOR"]
  state = ["ACTIVE"]
  vendor_name = "ORACLE"
}

#-------------------
# Management Station
#-------------------
resource "oci_os_management_hub_management_station" "test_management_station" {
  compartment_id = "${var.compartment_id}"
  display_name = "displayName"
  hostname = "hostname"
  mirror {
    directory = "/directory"
    port = "50001"
    sslport = "50002"
  }
  proxy {
    forward = "https://forward.com"
    hosts = ["host"]
    is_enabled = "true"
    port = "1029"
  }
}

#-----------------------
# Managed Instance Group
#-----------------------
resource "oci_os_management_hub_managed_instance_group" "test_managed_instance_group" {
  arch_type = "X86_64"
  compartment_id = "${var.compartment_id}"
  display_name = "displayName"
  os_family = "ORACLE_LINUX_8"
  software_source_ids = ["${data.oci_os_management_hub_software_sources.ol8_baseos_latest_x86_64.software_source_collection[0].items[0].id}"]
  vendor_name = "ORACLE"
}

#----------------------------------
# Software Source Profile (non-OCI)
#----------------------------------
resource "oci_os_management_hub_profile" "test_ss_profile" {
  #Required
  compartment_id = "${var.compartment_id}"
  display_name = "SSdisplayName"
  profile_type = "SOFTWARESOURCE"

  #Optional
  arch_type = "X86_64"
  defined_tags = {"Operations.CostCenter"= "42"}
  description = "description"
  is_default_profile = "false"
  freeform_tags = {"Department"= "Finance"}
  os_family = "ORACLE_LINUX_8"
  software_source_ids = ["${data.oci_os_management_hub_software_sources.ol8_baseos_latest_x86_64.software_source_collection[0].items[0].id}"]
  vendor_name = "ORACLE"
  management_station_id = "${oci_os_management_hub_management_station.test_management_station.id}"
}

data "oci_os_management_hub_profile" "test_ss_profile" {
  #Required
  profile_id = oci_os_management_hub_profile.test_ss_profile.id
}

data "oci_os_management_hub_profiles" "test_profiles" {
  #Optional
  compartment_id = var.compartment_id
  arch_type = "X86_64"
  display_name = ["SSdisplayName"]
  display_name_contains = "displayName"
  filter {
    name = "id"
    values = ["${oci_os_management_hub_profile.test_ss_profile.id}"]
  }
  is_default_profile = "false"
  is_service_provided_profile = "false"
  os_family = "ORACLE_LINUX_8"
  profile_id = "${oci_os_management_hub_profile.test_ss_profile.id}"
  profile_type = ["SOFTWARESOURCE"]
  registration_type = ["NON_OCI_LINUX"]
  state = "ACTIVE"
  vendor_name = "ORACLE"
}

#------------------------
# Group Profile (non-OCI)
#------------------------
resource "oci_os_management_hub_profile" "test_grp_profile" {
  compartment_id = "${var.compartment_id}"
  display_name = "GRPdisplayName"
  description = "description"
  managed_instance_group_id = "${oci_os_management_hub_managed_instance_group.test_managed_instance_group.id}"
  management_station_id = "${oci_os_management_hub_management_station.test_management_station.id}"
  profile_type = "GROUP"
  arch_type = "X86_64"
  freeform_tags = {
    "Department" = "Finance"
  }
  os_family = "ORACLE_LINUX_8"
  vendor_name = "ORACLE"
}

data "oci_os_management_hub_profile" "test_grp_profile" {
  #Required
  profile_id = "${oci_os_management_hub_profile.test_grp_profile.id}"
}

#------------------------------------
# Lifecycle Environment Profile (OCI)
#------------------------------------
resource "oci_os_management_hub_profile" "test_le_profile" {
  compartment_id = "${var.compartment_id}"
  display_name = "LEdisplayName"
  lifecycle_stage_id = "${var.lifecycle_stage_id}" 
  registration_type = "OCI_LINUX" 
  #management_station_id = "${oci_os_management_hub_management_station.test_management_station.id}"
  profile_type = "LIFECYCLE"
}

data "oci_os_management_hub_profile" "test_le_profile" {
  #Required
  profile_id = "${oci_os_management_hub_profile.test_le_profile.id}"
}

