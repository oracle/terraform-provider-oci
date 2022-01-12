// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * This example file shows how to create a group and add a user to it. 
 */

resource "oci_identity_group" "group1" {
  name           = "tf-example-group"
  description    = "group created by terraform"
  compartment_id = var.tenancy_ocid
}

resource "oci_identity_user_group_membership" "user-group-mem1" {
  compartment_id = var.tenancy_ocid
  user_id        = oci_identity_user.user1.id
  group_id       = oci_identity_group.group1.id
}

data "oci_identity_groups" "groups1" {
  compartment_id = oci_identity_group.group1.compartment_id

  filter {
    name   = "name"
    values = ["tf-example-group"]
  }
}

output "groups" {
  value = data.oci_identity_groups.groups1.groups
}

/*
 * Some more directives to show dynamic groups and policy for it
 */

variable "dynamic_group_defined_tags_value" {
  default = "test_value"
}

variable "dynamic_group_freeform_tags" {
  default = {
    "Department" = "Finance"
  }
}

resource "oci_identity_dynamic_group" "dynamic-group-1" {
  compartment_id = var.tenancy_ocid
  name           = "tf-example-dynamic-group"
  description    = "dynamic group created by terraform"
  matching_rule  = "ANY {instance.compartment.id = '${data.oci_identity_compartments.compartments1.compartments[0].id}'}"

  #Optional
  defined_tags = {
    "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = var.dynamic_group_defined_tags_value
  }
  freeform_tags = var.dynamic_group_freeform_tags
}

data "oci_identity_dynamic_groups" "dynamic-groups-1" {
  compartment_id = var.tenancy_ocid

  filter {
    name   = "id"
    values = [oci_identity_dynamic_group.dynamic-group-1.id]
  }
}

output "dynamicGroups" {
  value = data.oci_identity_dynamic_groups.dynamic-groups-1.dynamic_groups
}

