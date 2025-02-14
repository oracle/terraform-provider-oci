// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}


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
  arch_type            = ["X86_64"]
  availability         = ["SELECTED"]
  compartment_id       = var.compartment_ocid
  display_name         = "ol8_baseos_latest-x86_64"
  os_family            = ["ORACLE_LINUX_8"]
  software_source_type = ["VENDOR"]
  state                = ["ACTIVE"]
  vendor_name          = "ORACLE"
}

# Assumption: "ol8_appstream_x86_64" exists as a Vendor Software Source in the Tenancy
data "oci_os_management_hub_software_sources" "ol8_appstream_x86_64" {
  arch_type            = ["X86_64"]
  availability         = ["SELECTED"]
  compartment_id       = var.compartment_ocid
  display_name         = "ol8_appstream-x86_64"
  os_family            = ["ORACLE_LINUX_8"]
  software_source_type = ["VENDOR"]
  state                = ["ACTIVE"]
  vendor_name          = "ORACLE"
}

#-------------------
# Management Station
#-------------------
resource "oci_os_management_hub_management_station" "test_management_station" {
  compartment_id = var.compartment_ocid
  display_name   = "displayName"
  hostname       = "hostname"
  mirror {
    directory = "/directory"
    port      = "50001"
    sslport   = "50002"
  }
  proxy {
    forward    = "https://forward.com"
    hosts      = ["host"]
    is_enabled = "true"
    port       = "1029"
  }
}

#---------------------------------
# Managed Instance Group (non-oci)
#---------------------------------
resource "oci_os_management_hub_managed_instance_group" "test_managed_instance_group" {
  arch_type           = "X86_64"
  compartment_id      = var.compartment_ocid
  display_name        = "displayName"
  os_family           = "ORACLE_LINUX_8"
  software_source_ids = [data.oci_os_management_hub_software_sources.ol8_baseos_latest_x86_64.software_source_collection[0].items[0].id]
  vendor_name         = "ORACLE"
}

#-----------------------------
# Managed Instance Group (OCI)
#-----------------------------
resource "oci_os_management_hub_managed_instance_group" "test_managed_instance_group_oci" {
  arch_type           = "X86_64"
  compartment_id      = var.compartment_ocid
  display_name        = "displayName"
  location            = "OCI_COMPUTE"
  os_family           = "ORACLE_LINUX_8"
  software_source_ids = [data.oci_os_management_hub_software_sources.ol8_baseos_latest_x86_64.software_source_collection[0].items[0].id]
  vendor_name         = "ORACLE"
}

#----------------------------------
# Software Source Profile (non-OCI)
#----------------------------------
resource "oci_os_management_hub_profile" "test_ss_profile" {
  #Required
  compartment_id = var.compartment_ocid
  display_name   = "SSdisplayName"
  profile_type   = "SOFTWARESOURCE"

  #Optional
  arch_type             = "X86_64"
  defined_tags          = { "Operations.CostCenter" = "42" }
  description           = "description"
  is_default_profile    = "false"
  freeform_tags         = { "Department" = "Finance" }
  os_family             = "ORACLE_LINUX_8"
  software_source_ids   = [data.oci_os_management_hub_software_sources.ol8_baseos_latest_x86_64.software_source_collection[0].items[0].id]
  vendor_name           = "ORACLE"
  management_station_id = oci_os_management_hub_management_station.test_management_station.id
}

data "oci_os_management_hub_profile" "test_ss_profile" {
  #Required
  profile_id = oci_os_management_hub_profile.test_ss_profile.id
}

# Profile Version
data "oci_os_management_hub_profile_version" "test_profile_version" {
  profile_id      = oci_os_management_hub_profile.test_profile_oci.id
  profile_version = oci_os_management_hub_profile.test_profile_oci.profile_version
}

data "oci_os_management_hub_profiles" "test_profiles" {
  #Optional
  compartment_id        = var.compartment_ocid
  arch_type             = "X86_64"
  display_name          = ["SSdisplayName"]
  display_name_contains = "displayName"
  filter {
    name   = "id"
    values = [oci_os_management_hub_profile.test_ss_profile.id]
  }
  is_default_profile          = "false"
  is_service_provided_profile = "false"
  os_family                   = "ORACLE_LINUX_8"
  profile_id                  = oci_os_management_hub_profile.test_ss_profile.id
  profile_type                = ["SOFTWARESOURCE"]
  registration_type           = ["NON_OCI_LINUX"]
  state                       = "ACTIVE"
  vendor_name                 = "ORACLE"
}

#------------------------
# Group Profile (non-OCI)
#------------------------
resource "oci_os_management_hub_profile" "test_grp_profile" {
  compartment_id            = var.compartment_ocid
  display_name              = "GRPdisplayName"
  description               = "description"
  managed_instance_group_id = oci_os_management_hub_managed_instance_group.test_managed_instance_group.id
  management_station_id     = oci_os_management_hub_management_station.test_management_station.id
  profile_type              = "GROUP"
  arch_type                 = "X86_64"
  freeform_tags = {
    "Department" = "Finance"
  }
  os_family   = "ORACLE_LINUX_8"
  vendor_name = "ORACLE"
}

data "oci_os_management_hub_profile" "test_grp_profile" {
  #Required
  profile_id = oci_os_management_hub_profile.test_grp_profile.id
}

#------------------------
# Group Profile (OCI)
#------------------------
resource "oci_os_management_hub_profile" "test_profile_oci" {
  compartment_id            = var.compartment_ocid
  display_name              = "displayName"
  managed_instance_group_id = oci_os_management_hub_managed_instance_group.test_managed_instance_group_oci.id
  profile_type              = "GROUP"
  registration_type         = "OCI_LINUX"
}

#------------------------------------
# Lifecycle Environment Profile (OCI)
#------------------------------------
resource "oci_os_management_hub_lifecycle_environment" "test_lifecycle_environment" {
  arch_type      = "X86_64"
  compartment_id = var.compartment_ocid
  description    = "description2"
  display_name   = "displayName2"
  freeform_tags = {
    "Department" = "Accounting"
  }
  location  = "OCI_COMPUTE"
  os_family = "ORACLE_LINUX_8"
  stages {
    display_name = "displayName2"
    rank         = "1"
  }
  stages {
    display_name = "prod2"
    rank         = "2"
  }
  vendor_name = "ORACLE"
}

resource "oci_os_management_hub_profile" "test_le_profile" {
  compartment_id     = var.compartment_ocid
  display_name       = "LEdisplayName"
  lifecycle_stage_id = oci_os_management_hub_lifecycle_environment.test_lifecycle_environment.stages[0].id
  registration_type  = "OCI_LINUX"
  profile_type       = "LIFECYCLE"
}

data "oci_os_management_hub_profile" "test_le_profile" {
  #Required
  profile_id = oci_os_management_hub_profile.test_le_profile.id
}

#---------------------------------------
# Profile Attach Managed Instances Group
#---------------------------------------
resource "oci_os_management_hub_profile_attach_managed_instance_group_management" "test_profile_attach_managed_instance_group_management" {
  managed_instance_group_id = oci_os_management_hub_managed_instance_group.test_managed_instance_group_oci.id
  profile_id                = oci_os_management_hub_profile.test_profile_oci.id
}

#----------------------------------
# Profile Attach Management Station
#----------------------------------
resource "oci_os_management_hub_profile_attach_management_station_management" "test_profile_attach_management_station_management" {
  management_station_id = oci_os_management_hub_management_station.test_management_station.id
  profile_id            = oci_os_management_hub_profile.test_ss_profile.id
  depends_on            = [oci_os_management_hub_management_station.test_management_station]
}

#--------------------------------
# Profile Attach Software Sources
#--------------------------------
resource "oci_os_management_hub_profile_attach_software_sources_management" "test_profile_attach_software_sources_management" {
  profile_id       = oci_os_management_hub_profile.test_ss_profile.id
  software_sources = [data.oci_os_management_hub_software_sources.ol8_appstream_x86_64.software_source_collection[0].items[0].id]
}

#--------------------------------
# Profile Detach Software Sources
#--------------------------------
resource "oci_os_management_hub_profile_detach_software_sources_management" "test_profile_detach_software_sources_management" {
  profile_id       = oci_os_management_hub_profile.test_ss_profile.id
  software_sources = [data.oci_os_management_hub_software_sources.ol8_appstream_x86_64.software_source_collection[0].items[0].id]
  depends_on       = [oci_os_management_hub_profile_attach_software_sources_management.test_profile_attach_software_sources_management]
}

#-----------------------------------
# Profile Available Software Sources
#-----------------------------------
data "oci_os_management_hub_profile_available_software_sources" "test_profile_available_software_sources" {
  profile_id = oci_os_management_hub_profile.test_ss_profile.id
}

