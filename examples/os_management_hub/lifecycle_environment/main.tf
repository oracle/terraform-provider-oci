// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_os_management_hub_lifecycle_environment" "test_lifecycle_environment" {
    #Required
    arch_type = "X86_64"
       compartment_id = var.compartment_id
    display_name = "displayName"
    os_family = "ORACLE_LINUX_8"
    stages {
        #Required
        display_name = "test"
        rank = "1"

        #Optional
        defined_tags = {"Operations.CostCenter"= "42"}
        freeform_tags = {"Department"= "Finance"}
    }
    stages {
        #Required
        display_name = "prod"
        rank = "2"

        #Optional
        defined_tags = {"Operations.CostCenter"= "42"}
        freeform_tags = {"Department"= "Finance"}
    }
    vendor_name = "ORACLE"

    #Optional
    defined_tags = {"Operations.CostCenter"= "42"}
    description = "description"
    freeform_tags = {"Department"= "Finance"}
}

data "oci_os_management_hub_lifecycle_environment" "test_lifecycle_environment" {
    #Required
    lifecycle_environment_id = oci_os_management_hub_lifecycle_environment.test_lifecycle_environment.id
}

data "oci_os_management_hub_lifecycle_environments" "test_lifecycle_environments" {
    #Optional
    compartment_id = var.compartment_id
}

data "oci_os_management_hub_lifecycle_stage" "test_lifecycle_stage" {
    #Required
    lifecycle_stage_id = oci_os_management_hub_lifecycle_environment.test_lifecycle_environment.stages[0].id
}

data "oci_os_management_hub_lifecycle_stages" "test_lifecycle_stages" {
    #Optional
    compartment_id = var.compartment_id
}

