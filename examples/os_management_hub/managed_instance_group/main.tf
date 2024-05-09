// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}
variable "managed_instances" { type = list(string) }

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
  // version = "5.35.0"
}

data "oci_os_management_hub_software_sources" "ol8_baseos_latest_x86_64" {
  #Optional
  arch_type            = ["X86_64"]
  compartment_id       = var.compartment_id
  display_name         = "ol8_baseos_latest-x86_64"
  os_family            = ["ORACLE_LINUX_8"]
  software_source_type = ["VENDOR"]
  state                = ["ACTIVE"]
  vendor_name          = "ORACLE"
}

data "oci_os_management_hub_software_sources" "ol8_appstream_x86_64" {
  #Optional
  arch_type            = ["X86_64"]
  compartment_id       = var.compartment_id
  display_name         = "ol8_appstream-x86_64"
  os_family            = ["ORACLE_LINUX_8"]
  software_source_type = ["VENDOR"]
  state                = ["ACTIVE"]
  vendor_name          = "ORACLE"
}

resource "oci_os_management_hub_managed_instance_group" "test_managed_instance_group" {
  #Required
  arch_type           = "X86_64"
  compartment_id      = var.compartment_id
  display_name        = "displayName"
  os_family           = "ORACLE_LINUX_8"
  software_source_ids = [data.oci_os_management_hub_software_sources.ol8_baseos_latest_x86_64.software_source_collection[0].items[0].id]
  vendor_name         = "ORACLE"

  #Optional
  defined_tags  = { "Operations.CostCenter" = "42" }
  description   = "description"
  freeform_tags = { "Department" = "Finance" }
  location      = "OCI_COMPUTE"

  lifecycle {
    ignore_changes = [defined_tags]
  }
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
  compartment_id        = var.compartment_id
  display_name_contains = "tmux"
}

data "oci_os_management_hub_managed_instance_group_available_software_sources" "test_managed_instance_group_available_software_sources" {
  #Required
  managed_instance_group_id = oci_os_management_hub_managed_instance_group.test_managed_instance_group.id

  #Optional
  compartment_id = var.compartment_id
}

resource "oci_os_management_hub_managed_instance_group_attach_managed_instances_management" "test_managed_instance_group_attach_managed_instances_management" {
  #Required
  managed_instance_group_id = oci_os_management_hub_managed_instance_group.test_managed_instance_group.id
  managed_instances         = var.managed_instances

  #Optional
  work_request_details {
    #Optional
    description  = "description"
    display_name = "display_name"
  }
}

resource "time_sleep" "wait_for_attach_managed_instances" {
  depends_on = [oci_os_management_hub_managed_instance_group_attach_managed_instances_management.test_managed_instance_group_attach_managed_instances_management]

  create_duration = "30s"
}

resource "oci_os_management_hub_managed_instance_group_detach_managed_instances_management" "test_managed_instance_group_detach_managed_instances_management" {
  #Required
  managed_instance_group_id = oci_os_management_hub_managed_instance_group.test_managed_instance_group.id
  managed_instances         = var.managed_instances

  depends_on = [
    time_sleep.wait_for_attach_managed_instances
  ]
}

resource "oci_os_management_hub_managed_instance_group_attach_software_sources_management" "test_managed_instance_group_attach_software_sources_management" {
  #Required
  managed_instance_group_id = oci_os_management_hub_managed_instance_group.test_managed_instance_group.id
  software_sources          = [data.oci_os_management_hub_software_sources.ol8_appstream_x86_64.software_source_collection[0].items[0].id]

  #Optional
  work_request_details {
    #Optional
    description  = "description"
    display_name = "display_name"
  }
}

resource "time_sleep" "wait_for_attach_software_sources" {
  depends_on = [oci_os_management_hub_managed_instance_group_attach_software_sources_management.test_managed_instance_group_attach_software_sources_management]

  create_duration = "30s"
}

resource "oci_os_management_hub_managed_instance_group_detach_software_sources_management" "test_managed_instance_group_detach_software_sources_management" {
  #Required
  managed_instance_group_id = oci_os_management_hub_managed_instance_group.test_managed_instance_group.id
  software_sources          = [data.oci_os_management_hub_software_sources.ol8_appstream_x86_64.software_source_collection[0].items[0].id]

  #Optional
  work_request_details {
    #Optional
    description  = "description"
    display_name = "display_name"
  }

  depends_on = [
    time_sleep.wait_for_attach_software_sources
  ]
}

resource "oci_os_management_hub_managed_instance_group_install_packages_management" "test_managed_instance_group_install_packages_management" {
  #Required
  managed_instance_group_id = oci_os_management_hub_managed_instance_group.test_managed_instance_group.id
  package_names             = ["ed-1.14.2-4.el8.x86_64.rpm"]

  #Optional
  work_request_details {
    #Optional
    description  = "description"
    display_name = "display_name"
  }
}

resource "oci_os_management_hub_managed_instance_group_remove_packages_management" "test_managed_instance_group_remove_packages_management" {
  #Required
  managed_instance_group_id = oci_os_management_hub_managed_instance_group.test_managed_instance_group.id
  package_names             = ["ed-1.14.2-4.el8.x86_64.rpm"]

  #Optional
  work_request_details {
    #Optional
    description  = "description"
    display_name = "display_name"
  }

  depends_on = [
    oci_os_management_hub_managed_instance_group_install_packages_management.test_managed_instance_group_install_packages_management
  ]
}

resource "oci_os_management_hub_managed_instance_group_install_windows_updates_management" "test_managed_instance_group_install_windows_updates_management" {
  #Required
  managed_instance_group_id = oci_os_management_hub_managed_instance_group.test_managed_instance_group.id
  windows_update_types      = ["OTHER"]

  #Optional
  work_request_details {
    #Optional
    description  = "description"
    display_name = "display_name"
  }
}

resource "oci_os_management_hub_managed_instance_group_update_all_packages_management" "test_managed_instance_group_update_all_packages_management" {
  #Required
  managed_instance_group_id = oci_os_management_hub_managed_instance_group.test_managed_instance_group.id
  update_types              = ["OTHER"]

  #Optional
  work_request_details {
    #Optional
    description  = "description"
    display_name = "display_name"
  }
}

resource "oci_os_management_hub_managed_instance_group_manage_module_streams_management" "test_managed_instance_group_manage_module_streams_management" {
  #Required
  managed_instance_group_id = oci_os_management_hub_managed_instance_group.test_managed_instance_group.id

  #Optional
  disable {
    #Required
    module_name = "php"
    stream_name = "7.2"
  }

  is_dry_run = "true"

  work_request_details {
    #Optional
    description  = "description"
    display_name = "display_name"
  }
}
