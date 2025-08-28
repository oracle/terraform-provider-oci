// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
}

variable "compartment_ocid" {
}

variable "target_database_group_display_name" {
}

variable "access_level" {
  default = "ACCESSIBLE"
}
variable "display_name" {
  default = "TestGroup"
}
variable "state" {
  default = "ACTIVE"
}
variable "time_created_greater_than_or_equal_to" {
}
variable "time_created_less_than" {

}

variable "target_database_group_id" {
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

variable "target_database_id" {
  default = ""
}

variable "freeformTagsExample" {
  default = ""
}

# Create a Data Safe target database group
resource "oci_data_safe_target_database_group" "test_target_database_group" {
  #Required
  compartment_id = var.compartment_ocid
  display_name   = var.target_database_group_display_name

  matching_criteria {
    #Optional
    include {
      compartments {
        id                 = var.compartment_ocid
        is_include_subtree = true
      }
      freeform_tags = {
        Department = var.freeformTagsExample
      }
      target_database_ids = [var.target_database_id]
    }
  }
}

# Get Data Safe target database group
data "oci_data_safe_target_database_group" "test_target_database_group" {
  #Required
  target_database_group_id = oci_data_safe_target_database_group.test_target_database_group.id
}

# List Data Safe target database groups
data "oci_data_safe_target_database_groups" "test_target_database_groups" {
  #Required
  compartment_id = var.compartment_ocid

  # Optional
  access_level              = var.access_level
  compartment_id_in_subtree = true
  display_name              = var.display_name
  state                     = var.state
  target_database_group_id  = oci_data_safe_target_database_group.test_target_database_group.id
  time_created_greater_than_or_equal_to = var.time_created_greater_than_or_equal_to
  time_created_less_than                = var.time_created_less_than
}