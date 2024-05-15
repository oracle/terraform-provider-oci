// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}
variable "software_source_name" {}

provider "oci" {
    tenancy_ocid     = var.tenancy_ocid
    user_ocid        = var.user_ocid
    fingerprint      = var.fingerprint
    private_key_path = var.private_key_path
    region           = var.region
}

### software_source_name
# Software source name must be the full display name of an existing vendor software source in the tenancy.
# The software source must not be unavailable or restricted for OCI or on-premise, and must
# not be in use by any other resource (e.g. profiles, custom software sources, etc).
#
# Example: ol6_u0_base-x86_64

#############################################
# Software Sources
#############################################

# Vendor Software Sources
data "oci_os_management_hub_software_sources" "ol8_baseos_latest_x86_64" {
    #Optional
    arch_type = ["X86_64"]
    availability = ["SELECTED"]
    compartment_id = var.compartment_id
    display_name = "ol8_baseos_latest-x86_64"
    os_family = ["ORACLE_LINUX_8"]
    software_source_type = ["VENDOR"]
    state = ["ACTIVE"]
    vendor_name = "ORACLE"
}

data "oci_os_management_hub_software_sources" "ol8_appstream_x86_64" {
    #Optional
    arch_type = ["X86_64"]
    availability = ["SELECTED"]
    compartment_id = var.compartment_id
    display_name = "ol8_appstream-x86_64"
    os_family = ["ORACLE_LINUX_8"]
    software_source_type = ["VENDOR"]
    state = ["ACTIVE"]
    vendor_name = "ORACLE"
}

# Custom Software Sources
resource "oci_os_management_hub_software_source" "test_software_source_filter" {
    #Required
    compartment_id = var.compartment_id
    display_name = "displayName"
    software_source_type = "CUSTOM"
    vendor_software_sources {
        #Required
        display_name = "ol8_appstream-x86_64"
        id = data.oci_os_management_hub_software_sources.ol8_appstream_x86_64.software_source_collection[0].items[0].id
    }
    vendor_software_sources {
        display_name = "ol8_baseos_latest-x86_64"
        id = data.oci_os_management_hub_software_sources.ol8_baseos_latest_x86_64.software_source_collection[0].items[0].id
    }

    #Optional
    custom_software_source_filter {
        #Optional
        module_stream_profile_filters {
            filter_type = "INCLUDE"
            module_name = "php"
            profile_name = "common"
            stream_name = "8.0"
        }
        package_filters {
            #Required
            filter_type = "INCLUDE"

            #Optional
            package_name = "ed"
        }
        package_group_filters {
            filter_type = "INCLUDE"
            package_groups = ["base"]
        }
    }
    defined_tags = {"Operations.CostCenter"= "42"}
    description = "description"
    freeform_tags = {"Department"= "Finance"}
    is_auto_resolve_dependencies = "false"
    is_automatically_updated = "true"
    is_created_from_package_list = "false"
}

resource "oci_os_management_hub_software_source" "test_software_source_list" {
    #Required
    compartment_id = var.compartment_id
    display_name = "displayName"
    software_source_type = "CUSTOM"
    vendor_software_sources {
        #Required
        display_name = "ol8_appstream-x86_64"
        id = data.oci_os_management_hub_software_sources.ol8_appstream_x86_64.software_source_collection[0].items[0].id
    }
    vendor_software_sources {
        display_name = "ol8_baseos_latest-x86_64"
        id = data.oci_os_management_hub_software_sources.ol8_baseos_latest_x86_64.software_source_collection[0].items[0].id
    }

    #Optional
    defined_tags = {"Operations.CostCenter"= "42"}
    description = "description"
    freeform_tags = {"Department"= "Finance"}
    is_auto_resolve_dependencies = "false"
    is_automatically_updated = "true"
    is_created_from_package_list = "true"
}

# Get software source
data "oci_os_management_hub_software_source" "test_software_source" {
    #Required
    software_source_id = oci_os_management_hub_software_source.test_software_source_filter.id
}

# List software sources
data "oci_os_management_hub_software_sources" "test_software_sources" {
    #Optional
    software_source_id = data.oci_os_management_hub_software_sources.ol8_baseos_latest_x86_64.software_source_collection[0].items[0].id
    compartment_id = var.compartment_id
    display_name = "ol8_baseos_latest-x86_64"
    display_name_contains = "ol8_baseos_latest-x86_64"
    display_name_not_equal_to = ["displayNameNotEqualTo"]
    software_source_type = ["VENDOR"]
    arch_type = ["X86_64"]
    availability = ["SELECTED"]
    availability_anywhere = ["SELECTED"]
    availability_at_oci = ["SELECTED"]
    is_mandatory_for_autonomous_linux = "false"
    os_family = ["ORACLE_LINUX_8"]
    state = ["ACTIVE"]
    vendor_name = "ORACLE"
    filter {
        name = "id"
        values = [data.oci_os_management_hub_software_sources.ol8_baseos_latest_x86_64.software_source_collection[0].items[0].id]
    }
}

# List software source vendors
data "oci_os_management_hub_software_source_vendors" "test_software_source_vendors" {
    #Required
    compartment_id = var.compartment_id
}

#############################################
# Software Packages
#############################################

# Get software package by name
data "oci_os_management_hub_software_package" "test_software_package" {
    #Required
    software_package_name = "ModemManager-glib-devel-1.10.4-1.el8.x86_64.rpm"
}

# List all software packages
data "oci_os_management_hub_software_packages" "test_software_packages" {
    #Optional
    architecture = "X86_64"
    display_name_contains = "ModemManager"
    os_family = "ORACLE_LINUX_8"
    version = "1.10.4-1.el8"
}

# List software sources for a given package
data "oci_os_management_hub_software_package_software_source" "test_software_package_software_source" {
    #Required
    compartment_id = var.compartment_id
    software_package_name = "ModemManager-glib-devel-1.10.4-1.el8.x86_64.rpm"

    #Optional
    arch_type = ["X86_64"]
    availability = ["SELECTED"]
    availability_anywhere = ["SELECTED"]
    availability_at_oci = ["SELECTED"]
    display_name = "ol8_codeready_builder-x86_64"
    display_name_contains = "ol8_codeready_builder-x86_64"
    os_family = ["ORACLE_LINUX_8"]
    software_source_type = ["VENDOR"]
    state = ["ACTIVE"]
}

# Add software package
resource "oci_os_management_hub_software_source_add_packages_management" "test_software_source_add_packages_management" {
    packages = ["ModemManager-glib-1.10.4-1.el8.x86_64.rpm"]
    software_source_id = oci_os_management_hub_software_source.test_software_source_list.id
}

# Get software source package
data "oci_os_management_hub_software_source_software_package" "test_software_source_software_package" {
    #Required
    software_package_name = "zsh-5.5.1-10.el8.x86_64.rpm"
    software_source_id = data.oci_os_management_hub_software_sources.ol8_baseos_latest_x86_64.software_source_collection[0].items[0].id
}

# List software source packages
data "oci_os_management_hub_software_source_software_packages" "test_software_source_software_packages" {
    #Required
    software_source_id = data.oci_os_management_hub_software_sources.ol8_baseos_latest_x86_64.software_source_collection[0].items[0].id

    #Optional
    display_name = "zsh"
    display_name_contains = "zsh"
    is_latest = "true"
}

# Get software source package group
data "oci_os_management_hub_software_source_package_group" "test_software_source_package_group" {
    #Required
    package_group_id = "base"
    software_source_id = data.oci_os_management_hub_software_sources.ol8_baseos_latest_x86_64.software_source_collection[0].items[0].id
}

# List software source package groups
data "oci_os_management_hub_software_source_package_groups" "test_software_source_package_groups" {
    #Required
    software_source_id = data.oci_os_management_hub_software_sources.ol8_baseos_latest_x86_64.software_source_collection[0].items[0].id

    #Optional
    compartment_id = var.compartment_id
    group_type = ["GROUP"]
    name = "Base"
    name_contains = "Base"
}

#############################################
# Module Streams
#############################################

# Get module stream
data "oci_os_management_hub_software_source_module_stream" "test_software_source_module_stream" {
    #Required
    module_name = "php"
    software_source_id = data.oci_os_management_hub_software_sources.ol8_appstream_x86_64.software_source_collection[0].items[0].id
    stream_name = "8.0"
}

# List module streams
data "oci_os_management_hub_software_source_module_streams" "test_software_source_module_streams" {
    #Required
    software_source_id = data.oci_os_management_hub_software_sources.ol8_appstream_x86_64.software_source_collection[0].items[0].id

    #Optional
    is_latest = "true"
    module_name = "php"
    module_name_contains = "php"
    name = "8.0"
}

# Get module stream profile
data "oci_os_management_hub_software_source_module_stream_profile" "test_software_source_module_stream_profile" {
    #Required
    software_source_id = data.oci_os_management_hub_software_sources.ol8_appstream_x86_64.software_source_collection[0].items[0].id
    module_name = "php"
    profile_name = "common"
    stream_name = "8.0"
}

# List module stream profiles
data "oci_os_management_hub_software_source_module_stream_profiles" "test_software_source_module_stream_profiles" {
    #Required
    software_source_id = data.oci_os_management_hub_software_sources.ol8_appstream_x86_64.software_source_collection[0].items[0].id

    #Optional
    module_name = "php"
    name = "common"
    stream_name = "8.0"
}

#############################################
# Change Software Source Availabilities
#############################################

data "oci_os_management_hub_software_sources" "test_software_sources_change_availability" {
    compartment_id = var.compartment_id
    display_name = var.software_source_name
    software_source_type = ["VENDOR"]
}

resource "oci_os_management_hub_software_source_change_availability_management" "test_software_source_change_availability_management_selected" {
    software_source_availabilities {
        availability = "SELECTED"
        availability_at_oci = "SELECTED"
        software_source_id = data.oci_os_management_hub_software_sources.test_software_sources_change_availability.software_source_collection[0].items[0].id
    }
}

resource "oci_os_management_hub_software_source_change_availability_management" "test_software_source_change_availability_management_onprem_selected" {
    software_source_availabilities {
        availability = "SELECTED"
        availability_at_oci = "AVAILABLE"
        software_source_id = data.oci_os_management_hub_software_sources.test_software_sources_change_availability.software_source_collection[0].items[0].id
    }
}

resource "oci_os_management_hub_software_source_change_availability_management" "test_software_source_change_availability_management_oci_selected" {
    software_source_availabilities {
        availability = "AVAILABLE"
        availability_at_oci = "SELECTED"
        software_source_id = data.oci_os_management_hub_software_sources.test_software_sources_change_availability.software_source_collection[0].items[0].id
    }
}

resource "oci_os_management_hub_software_source_change_availability_management" "test_software_source_change_availability_management_available" {
    software_source_availabilities {
        availability = "AVAILABLE"
        availability_at_oci = "AVAILABLE"
        software_source_id = data.oci_os_management_hub_software_sources.test_software_sources_change_availability.software_source_collection[0].items[0].id
    }
}

