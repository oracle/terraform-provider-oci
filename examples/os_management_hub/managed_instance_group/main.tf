// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}
variable "software_source_ids" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_os_management_hub_managed_instance_group" "test_managed_instance_group" {
    #Required
    arch_type = "X86_64"
    compartment_id = var.compartment_id
    display_name = "displayName"
    os_family = "ORACLE_LINUX_8"
    software_source_ids = var.software_source_ids
    vendor_name = "ORACLE"

    #Optional
    defined_tags = {"Operations.CostCenter"= "42"}
    description = "description"
    freeform_tags = {"Department"= "Finance"}
}

data "oci_os_management_hub_managed_instance_group" "test_managed_instance_group" {
    #Required
    managed_instance_group_id = oci_os_management_hub_managed_instance_group.test_managed_instance_group.id
}

data "oci_os_management_hub_managed_instance_groups" "test_managed_instance_groups" {
    #Optional
    compartment_id = var.compartment_id
}

data "oci_os_management_hub_managed_instance_group_available_modules" "test_managed_instance_group_available_modules" {
    #Required
    managed_instance_group_id = oci_os_management_hub_managed_instance_group.test_managed_instance_group.id

    #Optional
    compartment_id = var.compartment_id
}

data "oci_os_management_hub_managed_instance_group_available_packages" "test_managed_instance_group_available_packages" {
    #Required
    managed_instance_group_id = oci_os_management_hub_managed_instance_group.test_managed_instance_group.id

    #Optional
    compartment_id = var.compartment_id
}

data "oci_os_management_hub_managed_instance_group_available_software_sources" "test_managed_instance_group_available_software_sources" {
    #Required
    managed_instance_group_id = oci_os_management_hub_managed_instance_group.test_managed_instance_group.id

    #Optional
    compartment_id = var.compartment_id
}

