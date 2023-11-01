// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}
variable "software_source_vendor_software_sources_id" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_os_management_hub_software_source" "test_software_source" {
    #Required
    compartment_id = var.compartment_id
    display_name = "displayName"
    software_source_type = "CUSTOM"
    vendor_software_sources {
        #Required
        display_name = "ol8_appstream-x86_64"
        id = var.software_source_vendor_software_sources_id
    }

    #Optional
    custom_software_source_filter {
        #Optional
        package_filters {
            #Required
            filter_type = "INCLUDE"

            #Optional
            package_name = "ed"
        }
    }
    defined_tags = {"Operations.CostCenter"= "42"}
    description = "description"
    freeform_tags = {"Department"= "Finance"}
    is_automatically_updated = "true"
}

data "oci_os_management_hub_software_source" "test_software_source" {
    #Required
    software_source_id = oci_os_management_hub_software_source.test_software_source.id
}

data "oci_os_management_hub_software_sources" "test_software_sources" {
    #Optional
    compartment_id = var.compartment_id
}

data "oci_os_management_hub_software_source_module_stream" "test_software_source_module_stream" {
    #Required
    module_name = "php"
    software_source_id = var.software_source_vendor_software_sources_id
    stream_name = "8.0"
}

data "oci_os_management_hub_software_source_module_streams" "test_software_source_module_streams" {
    #Required
    software_source_id = var.software_source_vendor_software_sources_id

    #Optional
    is_latest = "true"
    module_name = "php"
}

data "oci_os_management_hub_software_source_module_stream_profile" "test_software_source_module_stream_profile" {
    #Required
    module_name = "php"
    profile_name = "common"
    software_source_id = var.software_source_vendor_software_sources_id
    stream_name = "8.0"
}

data "oci_os_management_hub_software_source_module_stream_profiles" "test_software_source_module_stream_profiles" {
    #Required
    software_source_id = var.software_source_vendor_software_sources_id

    #Optional
    module_name = "php"
    stream_name = "8.0"
}

data "oci_os_management_hub_software_source_package_group" "test_software_source_package_group" {
    #Required
    package_group_id = "base"
    software_source_id = var.software_source_vendor_software_sources_id
}

data "oci_os_management_hub_software_source_package_groups" "test_software_source_package_groups" {
    #Required
    software_source_id = var.software_source_vendor_software_sources_id

    #Optional
    compartment_id = var.compartment_id
}

data "oci_os_management_hub_software_source_software_package" "test_software_source_software_package" {
    #Required
    software_package_name = "zsh-5.5.1-10.el8.x86_64.rpm"
    software_source_id = var.software_source_vendor_software_sources_id
}

data "oci_os_management_hub_software_source_software_packages" "test_software_source_software_packages" {
    #Required
    software_source_id = var.software_source_vendor_software_sources_id

    #Optional
    display_name = "zsh"
    is_latest = "true"
}

data "oci_os_management_hub_software_source_vendors" "test_software_source_vendors" {
    #Required
    compartment_id = var.compartment_id

    #Optional
    name = "ORACLE"
}

