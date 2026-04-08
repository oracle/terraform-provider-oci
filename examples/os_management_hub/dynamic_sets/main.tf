// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0
variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "osmh_managed_instance_ocid" {}
variable "osmh_managed_instance_group_ocid" {}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_os_management_hub_dynamic_set" "test_dynamic_set" {
  compartment_id = var.compartment_ocid
  description    = "description2"
  display_name   = "displayName2"
  freeform_tags = {
    "Department" = "Accounting"
  }
  match_type = "ANY"
  matching_rule {
    architectures              = ["AARCH64"]
    display_names              = ["displayNames2"]
    is_reboot_required         = "true"
    locations                  = ["OCI_COMPUTE"]
    managed_instance_group_ids = ["${var.osmh_managed_instance_group_ocid}"]
    managed_instance_ids       = ["${var.osmh_managed_instance_ocid}"]
    os_families                = ["ORACLE_LINUX_8"]
    os_names                   = ["ORACLE_LINUX"]
    statuses                   = ["NORMAL"]
    tags {
      key       = "key2"
      namespace = "namespace2"
      type      = "DEFINED"
      value     = "value2"
    }
  }
  target_compartments {
    compartment_id        = var.compartment_ocid
    does_include_children = "true"
  }
}


resource "oci_os_management_hub_dynamic_set" "test_dynamic_set_ol" {
  compartment_id = var.compartment_ocid
  display_name   = "ol8-dynamic-Set"
  match_type     = "ALL"
  matching_rule {
    architectures = ["X86_64"]
    locations     = ["OCI_COMPUTE"]
    os_families   = ["ORACLE_LINUX_8"]
  }
  target_compartments {
    compartment_id        = var.compartment_ocid
    does_include_children = "false"
  }
}

resource "oci_os_management_hub_dynamic_set_install_packages_management" "test_dynamic_set_install_packages_management" {
  dynamic_set_id    = oci_os_management_hub_dynamic_set.test_dynamic_set_ol.id
  managed_instances = ["${var.osmh_managed_instance_ocid}"]
  package_names     = ["ed"]
  # Optional
  work_request_details {
    description  = "description"
    display_name = "displayName"
  }
}

resource "oci_os_management_hub_dynamic_set_remove_packages_management" "test_dynamic_set_remove_packages_management" {
  dynamic_set_id    = oci_os_management_hub_dynamic_set.test_dynamic_set_ol.id
  managed_instances = ["${var.osmh_managed_instance_ocid}"]
  package_names     = ["ed"]
  depends_on        = [oci_os_management_hub_dynamic_set_install_packages_management.test_dynamic_set_install_packages_management]
  # Optional
  work_request_details {
    description  = "description"
    display_name = "displayName"
  }
}


resource "oci_os_management_hub_dynamic_set_update_packages_management" "test_dynamic_set_update_packages_management" {
  dynamic_set_id = oci_os_management_hub_dynamic_set.test_dynamic_set_ol.id
  update_types   = ["SECURITY"]
  # Optional
  managed_instances = ["${var.osmh_managed_instance_ocid}"]
  work_request_details {
    description  = "description"
    display_name = "displayName"
  }
}


resource "oci_os_management_hub_dynamic_set_reboot_management" "test_dynamic_set_reboot_management" {
  dynamic_set_id         = oci_os_management_hub_dynamic_set.test_dynamic_set_ol.id
  reboot_timeout_in_mins = "10"
  depends_on             = [oci_os_management_hub_dynamic_set_remove_packages_management.test_dynamic_set_remove_packages_management, oci_os_management_hub_dynamic_set_update_packages_management.test_dynamic_set_update_packages_management]
  # Optional
  managed_instances = ["${var.osmh_managed_instance_ocid}"]
  work_request_details {
    description  = "description"
    display_name = "displayName"
  }
}
