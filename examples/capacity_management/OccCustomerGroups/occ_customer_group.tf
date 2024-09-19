// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "occ_customer_group_customers_list_description" {
  default = "customer group"
}

variable "occ_customer_group_customers_list_display_name" {
  default = "customerGroup"
}

variable "occ_customer_group_customers_list_status" {
  default = "ENABLED"
}

variable "occ_customer_group_defined_tags_value" {
  default = "value"
}

variable "occ_customer_group_description" {
  default = "customer group"
}

variable "occ_customer_group_display_name" {
  default = "customerGroup"
}

variable "occ_customer_group_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "occ_customer_group_id" {
  default = "id"
}

variable "occ_customer_group_lifecycle_details" {
  default = "lifecycleDetails"
}

variable "occ_customer_group_status" {
  default = "ENABLED"
}

variable "occ_customer_group-occ_customer_tenancy_id" {
  default = "customerId"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_capacity_management_occ_customer_group" "test_occ_customer_group" {
  #Required
  compartment_id = var.compartment_id
  display_name   = var.occ_customer_group_display_name

  #Optional
  customers_list {
    #Required
    display_name = var.occ_customer_group_customers_list_display_name
    tenancy_id   = var.occ_customer_group-occ_customer_tenancy_id

    #Optional
    description = var.occ_customer_group_customers_list_description
    status      = var.occ_customer_group_customers_list_status
  }
  defined_tags      = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.occ_customer_group_defined_tags_value)
  description       = var.occ_customer_group_description
  freeform_tags     = var.occ_customer_group_freeform_tags
  lifecycle_details = var.occ_customer_group_lifecycle_details
  status            = var.occ_customer_group_status
}

data "oci_capacity_management_occ_customer_groups" "test_occ_customer_groups" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = var.occ_customer_group_display_name
  id           = var.occ_customer_group_id
  status       = var.occ_customer_group_status
}