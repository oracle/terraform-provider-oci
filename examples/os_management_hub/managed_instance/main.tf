// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0
variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "osmh_managed_instance_ocid" {}
variable "osmh_managed_instance_unregistered_ocid" {}
variable "osmh_managed_instance_windows_ocid" {}
variable "osmh_managed_instance_ubuntu_ocid" {}

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
  arch_type            = ["X86_64"]
  availability         = ["SELECTED"]
  compartment_id       = var.compartment_ocid
  display_name         = "ol8_baseos_latest-x86_64"
  os_family            = ["ORACLE_LINUX_8"]
  software_source_type = ["VENDOR"]
  state                = ["ACTIVE"]
  vendor_name          = "ORACLE"
}

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

# Create OL8 software source profile
resource "oci_os_management_hub_profile" "test_profile" {
  arch_type         = "X86_64"
  compartment_id    = var.compartment_ocid
  display_name      = "displayNameExample"
  os_family         = "ORACLE_LINUX_8"
  profile_type      = "SOFTWARESOURCE"
  registration_type = "OCI_LINUX"
  software_source_ids = [
    data.oci_os_management_hub_software_sources.ol8_baseos_latest_x86_64.software_source_collection[0].items[0].id
  ]
  vendor_name = "ORACLE"
}

resource "oci_os_management_hub_managed_instance_attach_profile_management" "test_managed_instance_attach_profile_management" {
  managed_instance_id = var.osmh_managed_instance_unregistered_ocid
  profile_id          = oci_os_management_hub_profile.test_profile.id
}

# 2. Detach profile from instance
resource "oci_os_management_hub_managed_instance_detach_profile_management" "test_managed_instance_detach_profile_management" {
  managed_instance_id = var.osmh_managed_instance_unregistered_ocid
  depends_on          = [oci_os_management_hub_managed_instance_attach_profile_management.test_managed_instance_attach_profile_management]
}

# 3. List available packages
data "oci_os_management_hub_managed_instance_available_packages" "test_managed_instance_available_packages" {
  # Required
  managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id

  # Optional
  compartment_id        = var.compartment_ocid
  display_name_contains = "389-ds-base"
}

# 4. List installed packages
data "oci_os_management_hub_managed_instance_installed_packages" "test_managed_instance_installed_packages" {
  # Required
  managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id

  # Optional
  compartment_id = var.compartment_ocid
}

# 5. List updatable packages
data "oci_os_management_hub_managed_instance_updatable_packages" "test_managed_instance_updatable_packages" {
  # Required
  managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id

  # Optional
  compartment_id = var.compartment_ocid
}

# 6. List errata
data "oci_os_management_hub_managed_instance_errata" "test_managed_instance_errata" {
  # Required
  managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id

  # Optional
  compartment_id = var.compartment_ocid
}

# 7. List modules
data "oci_os_management_hub_managed_instance_modules" "test_managed_instance_modules" {
  # Required
  managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id

  # Optional
  compartment_id = var.compartment_ocid
}

# 8. Available software source
data "oci_os_management_hub_managed_instance_available_software_sources" "test_managed_instance_available_software_sources" {
  # Required
  managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id

  # Optional
  compartment_id = var.compartment_ocid
}

# 9. Install packages
resource "oci_os_management_hub_managed_instance_install_packages_management" "test_managed_instance_install_packages_management" {
  managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id
  package_names       = ["ModemManager-1.20.2-1.el8.x86_64"]
}

# 10. Remove packages
resource "oci_os_management_hub_managed_instance_remove_packages_management" "test_managed_instance_remove_packages_management" {
  depends_on          = [oci_os_management_hub_managed_instance_install_packages_management.test_managed_instance_install_packages_management]
  managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance.id
  package_names       = ["ModemManager-1.20.2-1.el8.x86_64"]
  # Optional
  work_request_details {
    description  = "description"
    display_name = "displayName"
  }
}

# 11. test reboot
resource "oci_os_management_hub_managed_instance_reboot_management" "test_managed_instance_reboot_management" {
  managed_instance_id = var.osmh_managed_instance_ocid

  #optional
  reboot_timeout_in_mins = "15"
  depends_on             = [oci_os_management_hub_managed_instance_remove_packages_management.test_managed_instance_remove_packages_management, oci_os_management_hub_managed_instance_detach_software_sources_management.test_managed_instance_detach_software_sources_management]
}

# 12. Attach / Detach Software Sources
resource "oci_os_management_hub_managed_instance_detach_software_sources_management" "test_managed_instance_detach_software_sources_management" {
  depends_on          = [oci_os_management_hub_managed_instance_attach_software_sources_management.test_managed_instance_attach_software_sources_management]
  managed_instance_id = var.osmh_managed_instance_ocid
  software_sources    = ["${data.oci_os_management_hub_software_sources.ol8_appstream_x86_64.software_source_collection[0].items[0].id}"]
  # Optional
  work_request_details {
    description  = "Detach SS"
    display_name = "displayName"
  }
}

resource "oci_os_management_hub_managed_instance_attach_software_sources_management" "test_managed_instance_attach_software_sources_management" {
  managed_instance_id = var.osmh_managed_instance_ocid
  software_sources    = ["${data.oci_os_management_hub_software_sources.ol8_appstream_x86_64.software_source_collection[0].items[0].id}"]
  # Optional
  work_request_details {
    description  = "Attach SS"
    display_name = "displayName"
  }
}

# 13. Refresh Software
resource "oci_os_management_hub_managed_instance_refresh_software_management" "test_managed_instance_refresh_software_management" {
  managed_instance_id = var.osmh_managed_instance_ubuntu_ocid
}

################################
# Windows Instance / Update
################################

# 1. List available windows update
data "oci_os_management_hub_managed_instance_available_windows_updates" "test_managed_instance_available_windows_updates" {
  # Required
  managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance_windows.id

  # Optional
  compartment_id = var.compartment_ocid
}

# 2. List installed windows update
data "oci_os_management_hub_managed_instance_installed_windows_updates" "test_managed_instance_installed_windows_updates" {
  # Required
  managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance_windows.id

  # Optional
  compartment_id = var.compartment_ocid
}

# 3. List windows updates
data "oci_os_management_hub_windows_updates" "test_windows_updates" {
  # Required
  compartment_id = var.compartment_ocid
}


# 5. Install windows update
resource "oci_os_management_hub_managed_instance_install_windows_updates_management" "test_managed_instance_install_windows_updates_management_update_type" {
  # Required
  managed_instance_id = oci_os_management_hub_managed_instance.test_managed_instance_windows.id

  # Optional
  windows_update_types = ["OTHER"]
  work_request_details {
    description  = "description"
    display_name = "displayName"
  }
}

################################
# Ubuntu instances
################################

# Install Snaps on Ubuntu

resource "oci_os_management_hub_managed_instance_install_snaps_management" "test_managed_instance_install_snaps_management_1" {
  managed_instance_id = var.osmh_managed_instance_ubuntu_ocid
  snap_details {
    channel = "stable"
    name    = "speedtest-cli"
  }
}

resource "oci_os_management_hub_managed_instance_install_snaps_management" "test_managed_instance_install_snaps_management" {
  managed_instance_id = var.osmh_managed_instance_ubuntu_ocid
  snap_details {
    name    = "hello-world"
    channel = "stable"
    # Optional
    is_signed = "false"
    mode      = "CLASSIC"
    revision  = "27"
  }
  # Optional
  work_request_details {
    description  = "description"
    display_name = "displayName"
  }
}

# Remove Snaps on Ubuntu
resource "oci_os_management_hub_managed_instance_remove_snaps_management" "test_managed_instance_remove_snaps_management" {
  depends_on          = [oci_os_management_hub_managed_instance_install_snaps_management.test_managed_instance_install_snaps_management]
  managed_instance_id = var.osmh_managed_instance_ubuntu_ocid
  snap_details {
    name     = "hello-world"
    revision = "27"
  }
  # Optional
  work_request_details {
    description  = "description"
    display_name = "displayName"
  }
}

# Switch Snap Channel on Ubuntu
resource "oci_os_management_hub_managed_instance_switch_snap_channel_management" "test_managed_instance_switch_snap_channel_management" {
  depends_on          = [oci_os_management_hub_managed_instance_install_snaps_management.test_managed_instance_install_snaps_management_1]
  managed_instance_id = var.osmh_managed_instance_ubuntu_ocid
  snap_details {
    channel = "beta"
    name    = "speedtest-cli"
  }
  # Optional
  work_request_details {
    description  = "description"
    display_name = "displayName"
  }
}

# Installed snaps

data "oci_os_management_hub_managed_instance_snaps" "test_managed_instance_snaps" {
  managed_instance_id = var.osmh_managed_instance_ubuntu_ocid
}

data "oci_os_management_hub_managed_instance_snaps" "test_managed_instance_snaps_with__name_filters" {
  managed_instance_id = var.osmh_managed_instance_ubuntu_ocid
  # Optional
  name          = "oracle-cloud-agent"
  name_contains = "oracle-cloud-agent"
}