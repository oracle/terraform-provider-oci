// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
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

data "oci_os_management_hub_software_sources" "ol8_baseos_latest_x86_64" {
  #Optional
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
  #Optional
  arch_type            = ["X86_64"]
  availability         = ["SELECTED"]
  compartment_id       = var.compartment_ocid
  display_name         = "ol8_appstream-x86_64"
  os_family            = ["ORACLE_LINUX_8"]
  software_source_type = ["VENDOR"]
  state                = ["ACTIVE"]
  vendor_name          = "ORACLE"
}

resource "oci_os_management_hub_software_source" "test_software_source" {
  #Required
  compartment_id       = var.compartment_ocid
  display_name         = "displayNameVersioned"
  software_source_type = "VERSIONED"
  vendor_software_sources {
    display_name = "ol8_baseos_latest-x86_64"
    id           = data.oci_os_management_hub_software_sources.ol8_baseos_latest_x86_64.software_source_collection[0].items[0].id
  }
  vendor_software_sources {
    #Required
    display_name = "ol8_appstream-x86_64"
    id           = data.oci_os_management_hub_software_sources.ol8_appstream_x86_64.software_source_collection[0].items[0].id
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
  description                  = "description"
  is_auto_resolve_dependencies = "false"
  is_automatically_updated     = "true"
  is_created_from_package_list = "false"
  is_latest_content_only       = "true"
  software_source_version      = "1.0"

  lifecycle {
    ignore_changes = [defined_tags]
  }
}

# 1. Create a Lifecycle Environment with two Lifecycle Stages (min 2, max 5, cannot be created individually)
resource "oci_os_management_hub_lifecycle_environment" "test_lifecycle_environment" {
  #Required
  arch_type      = "X86_64"
  compartment_id = var.compartment_ocid
  display_name   = "displayName"
  os_family      = "ORACLE_LINUX_8"
  stages {
    #Required
    display_name = "test"
    rank         = "1"

    #Optional
    defined_tags  = { "Operations.CostCenter" = "42" }
    freeform_tags = { "Department" = "Finance" }
  }
  stages {
    #Required
    display_name = "prod"
    rank         = "2"

    #Optional
    defined_tags  = { "Operations.CostCenter" = "42" }
    freeform_tags = { "Department" = "Finance" }
  }
  vendor_name = "ORACLE"

  #Optional
  defined_tags  = { "Operations.CostCenter" = "42" }
  description   = "description"
  freeform_tags = { "Department" = "Finance" }
  location      = "OCI_COMPUTE"
}

# 2. Get Lifecycle Environment by id
data "oci_os_management_hub_lifecycle_environment" "test_lifecycle_environment" {
  #Required
  lifecycle_environment_id = oci_os_management_hub_lifecycle_environment.test_lifecycle_environment.id
}

# 3. List Lifecycle Environments
data "oci_os_management_hub_lifecycle_environments" "test_lifecycle_environments" {
  #Optional
  compartment_id = var.compartment_ocid
}

# 4. Get Lifecycle Stage by id
data "oci_os_management_hub_lifecycle_stage" "test_lifecycle_stage" {
  #Required
  lifecycle_stage_id = oci_os_management_hub_lifecycle_environment.test_lifecycle_environment.stages[0].id
}

# 5. List Lifecycle Stages
data "oci_os_management_hub_lifecycle_stages" "test_lifecycle_stages" {
  #Optional
  compartment_id = var.compartment_ocid
}

# 6. Attach Managed Instance to a Lifecycle Stage
resource "oci_os_management_hub_lifecycle_stage_attach_managed_instances_management" "test_lifecycle_stage_attach_managed_instances_management" {
  lifecycle_stage_id = data.oci_os_management_hub_lifecycle_stage.test_lifecycle_stage.lifecycle_stage_id
  managed_instance_details {
    managed_instances = [
      oci_os_management_hub_managed_instance.test_managed_instance.id
    ]
    work_request_details {
      description  = "description"
      display_name = "displayName"
    }
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
    work_request_details {
      description  = "description"
      display_name = "displayName"
    }
  }
  depends_on = [oci_os_management_hub_lifecycle_stage_attach_managed_instances_management.test_lifecycle_stage_attach_managed_instances_management]
}

# 8. Promote Versioned Custom Software Source to a Lifecycle Stage
resource "oci_os_management_hub_lifecycle_stage_promote_software_source_management" "test_lifecycle_stage_promote_software_source_management" {
  #Required
  lifecycle_stage_id = data.oci_os_management_hub_lifecycle_stage.test_lifecycle_stage.lifecycle_stage_id
  #Optional
  software_source_id = oci_os_management_hub_software_source.test_software_source.id
  depends_on         = [oci_os_management_hub_lifecycle_stage_detach_managed_instances_management.test_lifecycle_stage_detach_managed_instances_management]
}

# 9. Reboot Lifecycle Stage
resource "oci_os_management_hub_lifecycle_stage_reboot_management" "test_lifecycle_stage_reboot_management" {
  #Required
  lifecycle_stage_id = data.oci_os_management_hub_lifecycle_stage.test_lifecycle_stage.lifecycle_stage_id
  #Optional
  reboot_timeout_in_mins = "5"
}
