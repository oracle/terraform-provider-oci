// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}
variable "profile_software_source_ids" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_os_management_hub_profile" "test_profile" {
    #Required
    compartment_id = var.compartment_id
    display_name = "displayName"
    profile_type = "SOFTWARESOURCE"

    #Optional
    arch_type = "X86_64"
    defined_tags = {"Operations.CostCenter"= "42"}
    description = "description"
    freeform_tags = {"Department"= "Finance"}
    os_family = "ORACLE_LINUX_8"
    software_source_ids = var.profile_software_source_ids
    vendor_name = "ORACLE"
}

data "oci_os_management_hub_profile" "test_profile" {
    #Required
    profile_id = oci_os_management_hub_profile.test_profile.id
}

data "oci_os_management_hub_profiles" "test_profiles" {
    #Optional
       compartment_id = var.compartment_id
}

