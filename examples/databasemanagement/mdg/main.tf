// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

provider "oci" {
  tenancy_ocid = var.tenancy_ocid
  user_ocid = var.user_ocid
  fingerprint = var.fingerprint
  private_key_path = var.private_key_path
  region = var.region
}

variable "compartment_id" {
  default = "ocid1.test.oc1..<unique_ID>EXAMPLE-compartmentId-Value"
}

variable "managed_database_group_name" {
  default = "EXAMPLE-managedDatabaseGroupName-Value"
}

variable "managed_database_id" {
   default = "ocid1.test.oc1..<unique_ID>EXAMPLE-managedDatabaseId-Value"
}

variable "managed_database_group_state" {
  default = "ACTIVE"
}

variable "managed_database_group_description" {
  default = "Sales test database group"
}

variable "managed_db_group_defined_tags_value" {
  default = "managed_db_group_tag_value"
}

variable "managed_db_group_freeform_tags" {
  default = { "bar-key" = "value" }
}

# Create a new Tag Namespace.
resource "oci_identity_tag_namespace" "tag_namespace1" {
  #Required
  compartment_id = var.tenancy_ocid
  description    = "example tag namespace"
  name           = "example-tag-namespace-all"
}

# Create a new Tag definition in the above Tag Namespace.
resource "oci_identity_tag" "tag1" {
  #Required
  description      = "example tag"
  name             = "example-tag"
  tag_namespace_id = oci_identity_tag_namespace.tag_namespace1.id
}

# Create a new Managed Database Group resource and optionally add a Managed Database to it.
resource "oci_database_management_managed_database_group" "test_managed_database_group" {
  #Required
  compartment_id = var.compartment_id
  name = var.managed_database_group_name

  #Optional
  description = var.managed_database_group_description
  managed_databases {
    id = var.managed_database_id
  }
  defined_tags  = {
    "${oci_identity_tag_namespace.tag_namespace1.name}.${oci_identity_tag.tag1.name}" = var.managed_db_group_defined_tags_value
  }
  freeform_tags = var.managed_db_group_freeform_tags
}

# List Managed Database Group resources filtered by OCID and lifecycle state.
data "oci_database_management_managed_database_groups" "test_managed_database_groups_with_id" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  id = oci_database_management_managed_database_group.test_managed_database_group.id
  state = var.managed_database_group_state
}

# List Managed Database Group resources filtered by their name and lifecycle state.
data "oci_database_management_managed_database_groups" "test_managed_database_groups_with_name" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  name = var.managed_database_group_name
  state = var.managed_database_group_state
}
