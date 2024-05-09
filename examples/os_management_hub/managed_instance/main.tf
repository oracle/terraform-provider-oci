// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0
variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}
variable "osmh_managed_instance_ocid" {}
variable "osmh_managed_instance_failed_ocid" {}
variable "osmh_managed_instance_windows_ocid" {}
variable "osmh_windows_update_id" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

### Instances
# OL8 instance
resource "oci_os_management_hub_managed_instance" "test_managed_instance" {
  managed_instance_id = var.osmh_managed_instance_ocid
}

# OL8 instance in registration failure and agent turned off
resource "oci_os_management_hub_managed_instance" "test_registration_failed_managed_instance" {
  managed_instance_id = var.osmh_managed_instance_failed_ocid
}

# Windows 2022 instance
resource "oci_os_management_hub_managed_instance" "test_managed_instance_windows" {
  managed_instance_id = var.osmh_managed_instance_windows_ocid
}

################################
# Oracle Linux Instance
################################

# 1. Attach profile to instance
# Reference OL8 software source
data "oci_os_management_hub_software_sources" "ol8_baseos_latest_x86_64" {
  arch_type = ["X86_64"]
  availability = ["SELECTED"]
  compartment_id = var.compartment_id
  display_name = "ol8_baseos_latest-x86_64"
  os_family = ["ORACLE_LINUX_8"]
  software_source_type = ["VENDOR"]
  state = ["ACTIVE"]
  vendor_name = "ORACLE"
}

# Create OL8 software source profile
resource "oci_os_management_hub_profile" "test_profile" {
  arch_type = "X86_64"
  compartment_id = var.compartment_id
  display_name = "displayNameExample"
  os_family = "ORACLE_LINUX_8"
  profile_type = "SOFTWARESOURCE"
  registration_type = "OCI_LINUX"
  software_source_ids = [
    data.oci_os_management_hub_software_sources.ol8_baseos_latest_x86_64.software_source_collection[0].items[0].id
  ]
  vendor_name = "ORACLE"
}
resource "oci_os_management_hub_managed_instance_attach_profile_management" "test_managed_instance_attach_profile_management" {
  managed_instance_id = oci_os_management_hub_managed_instance.test_registration_failed_managed_instance.id
  profile_id = oci_os_management_hub_profile.test_profile.id
}

# 2. Detach profile from instance
resource "oci_os_management_hub_managed_instance_detach_profile_management" "test_managed_instance_detach_profile_management" {
  managed_instance_id = oci_os_management_hub_managed_instance.test_registration_failed_managed_instance.id
  depends_on = [oci_os_management_hub_managed_instance_attach_profile_management.test_managed_instance_attach_profile_management]
}

# 3. List available packages
data "oci_os_management_hub_managed_instance_available_packages" "test_managed_instance_available_packages" {
  # Required
  managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id

  # Optional
  compartment_id = var.compartment_id
  display_name_contains = "389-ds-base"
}

# 4. List installed packages
data "oci_os_management_hub_managed_instance_installed_packages" "test_managed_instance_installed_packages" {
  # Required
  managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id

  # Optional
  compartment_id = var.compartment_id
}

# 5. List updatable packages
data "oci_os_management_hub_managed_instance_updatable_packages" "test_managed_instance_updatable_packages" {
  # Required
  managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id

  # Optional
  compartment_id = var.compartment_id
}

# 6. List errata
data "oci_os_management_hub_managed_instance_errata" "test_managed_instance_errata" {
  # Required
  managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id

  # Optional
  compartment_id = var.compartment_id
}

# 7. List modules
data "oci_os_management_hub_managed_instance_modules" "test_managed_instance_modules" {
  # Required
  managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id

  # Optional
  compartment_id = var.compartment_id
}

# 8. Available software source
data "oci_os_management_hub_managed_instance_available_software_sources" "test_managed_instance_available_software_sources" {
  # Required
  managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id

  # Optional
  compartment_id = var.compartment_id
}

################################
# Windows Instance / Update
################################

# 1. List available windows update
data "oci_os_management_hub_managed_instance_available_windows_updates" "test_managed_instance_available_windows_updates" {
  # Required
  managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance_windows.id

  # Optional
  compartment_id = var.compartment_id
}

# 2. List installed windows update
data "oci_os_management_hub_managed_instance_installed_windows_updates" "test_managed_instance_installed_windows_updates" {
  # Required
  managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance_windows.id

  # Optional
  compartment_id = var.compartment_id
}

# 3. List windows updates
data "oci_os_management_hub_windows_updates" "test_windows_updates" {
  # Required
  compartment_id = var.compartment_id
}

# 4. Get one windows update
data "oci_os_management_hub_windows_update" "test_windows_update" {
  # Required
  windows_update_id = var.osmh_windows_update_id
}

# 5. Install windows update
resource "oci_os_management_hub_managed_instance_install_windows_updates_management" "test_managed_instance_install_windows_updates_management_update_type" {
  # Required
  managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance_windows.id

  # Optional
  windows_update_types = ["OTHER"]
  work_request_details {
    description = "description"
    display_name = "displayName"
  }
}
