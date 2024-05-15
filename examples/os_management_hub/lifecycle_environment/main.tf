// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}
variable "osmh_managed_instance_ocid" {}

provider "oci" {
    tenancy_ocid     = var.tenancy_ocid
    user_ocid        = var.user_ocid
    fingerprint      = var.fingerprint
    private_key_path = var.private_key_path
    region           = var.region
}

### Dependency resources
# OL8 instance to attach
resource "oci_os_management_hub_managed_instance" "test_managed_instance" {
    managed_instance_id = var.osmh_managed_instance_ocid
}

# Referencing OL8 Versioned Custom Software Source
data "oci_os_management_hub_software_sources" "versioned_ol8_addons-x86_64" {
    arch_type = ["X86_64"]
    availability = ["SELECTED"]
    compartment_id = var.compartment_id
    display_name = "tf-vcss-to-promote-1"
    os_family = ["ORACLE_LINUX_8"]
    software_source_type = ["VERSIONED"]
    state = ["ACTIVE"]
}

# 1. Create a Lifecycle Environment with two Lifecycle Stages (min 2, max 5, cannot be created individually)
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
    location = "OCI_COMPUTE"
}

# 2. Get Lifecycle Environment by id
data "oci_os_management_hub_lifecycle_environment" "test_lifecycle_environment" {
    #Required
    lifecycle_environment_id = oci_os_management_hub_lifecycle_environment.test_lifecycle_environment.id
}

# 3. List Lifecycle Environments
data "oci_os_management_hub_lifecycle_environments" "test_lifecycle_environments" {
    #Optional
    compartment_id = var.compartment_id
}

# 4. Get Lifecycle Stage by id
data "oci_os_management_hub_lifecycle_stage" "test_lifecycle_stage" {
    #Required
    lifecycle_stage_id = oci_os_management_hub_lifecycle_environment.test_lifecycle_environment.stages[0].id
}

# 5. List Lifecycle Stages
data "oci_os_management_hub_lifecycle_stages" "test_lifecycle_stages" {
    #Optional
    compartment_id = var.compartment_id
}

# 6. Attach Managed Instance to a Lifecycle Stage
resource "oci_os_management_hub_lifecycle_stage_attach_managed_instances_management" "test_lifecycle_stage_attach_managed_instances_management" {
    lifecycle_stage_id = data.oci_os_management_hub_lifecycle_stage.test_lifecycle_stage.lifecycle_stage_id
    managed_instance_details {
        managed_instances = [
            oci_os_management_hub_managed_instance.test_managed_instance.id
        ]
    }
    depends_on = [oci_os_management_hub_lifecycle_environment.test_lifecycle_environment]
}

# 7. Detach Managed Instance from a Lifecycle Stage
resource "oci_os_management_hub_lifecycle_stage_detach_managed_instances_management" "test_lifecycle_stage_detach_managed_instances_management" {
    lifecycle_stage_id = data.oci_os_management_hub_lifecycle_stage.test_lifecycle_stage.lifecycle_stage_id
    managed_instance_details {
        managed_instances = [
            oci_os_management_hub_managed_instance.test_managed_instance.id
        ]
    }
    depends_on = [oci_os_management_hub_lifecycle_stage_attach_managed_instances_management.test_lifecycle_stage_attach_managed_instances_management]
}

# 8. Promote Versioned Custom Software Source to a Lifecycle Stage
resource "oci_os_management_hub_lifecycle_stage_promote_software_source_management" "test_lifecycle_stage_promote_software_source_management" {
    lifecycle_stage_id = data.oci_os_management_hub_lifecycle_stage.test_lifecycle_stage.lifecycle_stage_id
    software_source_id = data.oci_os_management_hub_software_sources.versioned_ol8_addons-x86_64.software_source_collection[0].items[0].id
    depends_on = [oci_os_management_hub_lifecycle_stage_detach_managed_instances_management.test_lifecycle_stage_detach_managed_instances_management]
}