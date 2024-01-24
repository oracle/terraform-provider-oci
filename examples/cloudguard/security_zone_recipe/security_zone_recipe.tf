// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

//Has to be unique
variable "security_recipe_display_name" {
  default = "displayName"
}

//Acceptable values come from LifecycleStateEnum
variable "lifecycle_state_active" {
  default = "ACTIVE"
}

variable "security_recipe_description" {
  default = "description"
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

/*
This data sources is used to get the security policy id to attach to security zone.
*/
data "oci_cloud_guard_security_policies" "test_security_policies" {
  #Required
  compartment_id = "${var.tenancy_ocid}"
  #Optional
  state          = "${var.lifecycle_state_active}"
}

/*
When CloudGuard is enabled, a Maximum Security Zone Recipe with all the default policies enabled is
available to be used when creating Security Zone.
If an user wants to create a custom Security Zone Recipe, then they can  perform a create recipe
by listing all the relevant policies.

In this example, we will list all the available security policies and then pick the first item in
the collection for creating new recipe.
*/
resource "oci_cloud_guard_security_recipe" "test_security_recipe" {
  #Required
  compartment_id    = "${var.compartment_id}"
  display_name      = "${var.security_recipe_display_name}"
  security_policies = [
    "${data.oci_cloud_guard_security_policies.test_security_policies.security_policy_collection[0].items[0].id}",
  ]
  #Optional
  description       = "${var.security_recipe_description}"

}