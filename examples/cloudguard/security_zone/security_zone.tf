// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "security_zone_access_level" {
  default = "ACCESSIBLE"
}

variable "security_zone_description" {
  default = "description"
}

//Has to be unique
variable "security_zone_display_name" {
  default = "displayName"
}

//Acceptable values come from LifecycleStateEnum
variable "lifecycle_state_active" {
  default = "ACTIVE"
}

provider "oci" {
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region           = "${var.region}"
}


data "oci_cloud_guard_security_recipes" "test_security_recipes" {
  #Required
  compartment_id = "${var.tenancy_ocid}"
  #Optional
  state          = "${var.lifecycle_state_active}"
}

data "oci_cloud_guard_security_zones" "test_security_zones" {
  #Required
  compartment_id = "${var.compartment_id}"
  #Optional
  state          = "${var.lifecycle_state_active}"
}

/*
This data sources is used to get the security_policy_id to attach to security zone.
*/
data "oci_cloud_guard_security_policies" "test_security_policies" {
  #Required
  compartment_id = "${var.tenancy_ocid}"
  #Optional
  state          = "${var.lifecycle_state_active}"
}

/*
When CloudGuard is enabled, a Maximum Security Zone Recipe with all the default policies enabled is
made available. If a user wants to create a custom Security Zone Recipe,
then they can  perform a create recipe with all the relevant policies enabled.

In this example, we will list all the available security recipes and then pick the first item in
the collection for creating a security zone.
*/
resource "oci_cloud_guard_security_zone" "test_security_zone" {
  #Required
  compartment_id          = "${var.compartment_id}"
  display_name            = "${var.security_zone_display_name}"
  security_zone_recipe_id = "${data.oci_cloud_guard_security_recipes.test_security_recipes.security_recipe_collection.0.items.0.id}"


  #Optional
  description   = "${var.security_zone_description}"

}
