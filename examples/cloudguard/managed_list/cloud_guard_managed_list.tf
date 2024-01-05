// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

//Acceptable values are ACCESSIBLE and RESTRICTED.
variable "managed_list_access_level" {
  default = "ACCESSIBLE"
}

variable "managed_list_compartment_id_in_subtree" {
  default = true
}
/*
If compartment_id_in_subtree is true and access_level is ACCESSIBLE then all managedLists under the provided compartment and child compartments will be shows
provided the user has permissions for it. If a user is lacking permission for one of the child compartments then the managedList from that child compartment will
not be visible in the result set. However if the access_level is RESTRICTED then the user needs to have permissions for all the child compartments of the passed in
compartment , else the operation will be unauthorized.
*/

variable "managed_list_defined_tags_value" {
  default = "value"
}

variable "managed_list_description" {
  default = "description"
}

//Has to be unique
variable "managed_list_display_name" {
  default = "displayName"
}

variable "managed_list_freeform_tags" {
  default = {
    "bar-key" = "value"
  }
}

variable "managed_list_list_items" {
  default = ["test-user"]
}

//Acceptable values should come from ManagedListTypeEnum.
variable "managed_list_list_type" {
  default = "USERS"
}

//Acceptable values come from LifecycleStateEnum
variable "managed_list_state" {
  default = "ACTIVE"
}

provider "oci" {
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region           = "${var.region}"
}

resource "oci_cloud_guard_managed_list" "test_managed_list" {
  #Required
  compartment_id = "${var.compartment_id}"
  display_name   = "${var.managed_list_display_name}"

  #Optional
  description   = "${var.managed_list_description}"
  list_items     = "${var.managed_list_list_items}"
  list_type      = "${var.managed_list_list_type}"
  /*
  When CloudGuard is Enabled, an Oracle Managed Managed List is made available having all the defaults.
  However if user wants to create and manage the ORACLE MANAGED Entity itself, it will have to enable the cloudguard with
  selfManageResources flag set to true. (More details in cloud_guard_configuration.tf)
  If user chooses to create its own ORACLE Managed Entity, below value should have non-nil value and have id of the ORACLE
  MANAGED Entity which is always available.
  */
  //source_managed_list_id
}

//Plural Data Source Representation
data "oci_cloud_guard_managed_lists" "test_managed_lists" {
  #Required
  compartment_id = "${var.tenancy_ocid}"

  #Optional
  access_level              = "${var.managed_list_access_level}"
  compartment_id_in_subtree = "${var.managed_list_compartment_id_in_subtree}"
  display_name              = "${var.managed_list_display_name}"
  list_type                 = "${var.managed_list_list_type}"
  state                     = "${var.managed_list_state}"
}