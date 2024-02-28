// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "occ_capacity_request_availability_domain" {
  default = "availabilityDomain"
}

variable "occ_capacity_request_date_expected_capacity_handover" {
  default = "2023/12/17"
}

variable "occ_capacity_request_defined_tags_value" {
  default = "value"
}

variable "occ_capacity_request_description" {
  default = "This is the test request created for UI"
}

variable "occ_capacity_request_details_actual_handover_quantity" {
  default = 10
}

variable "occ_capacity_request_details_date_actual_handover" {
  default = "dateActualHandover"
}

variable "occ_capacity_request_details_date_expected_handover" {
  default = "dateExpectedHandover"
}

variable "occ_capacity_request_details_demand_quantity" {
  default = 10
}

variable "occ_capacity_request_details_expected_handover_quantity" {
  default = 10
}

variable "occ_capacity_request_details_resource_type" {
  default = "SERVER_HW"
}

variable "occ_capacity_request_details_workload_type" {
  default = "GENERIC"
}

variable "occ_capacity_request_display_name" {
  default = "UI test request"
}

variable "occ_capacity_request_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "occ_capacity_request_id" {
  default = "id"
}

variable "occ_capacity_request_lifecycle_details" {
  default = "lifecycleDetails"
}

variable "occ_capacity_request_namespace" {
  default = "COMPUTE"
}

variable "occ_capacity_request_region" {
  default = "US-ASHBURN-1"
}

variable "occ_capacity_request_request_state" {
  default = "SUBMITTED"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_capacity_management_occ_capacity_request" "test_occ_capacity_request" {
  #Required
  availability_domain             = var.occ_capacity_request_availability_domain
  compartment_id                  = var.compartment_id
  date_expected_capacity_handover = var.occ_capacity_request_date_expected_capacity_handover
  details {
    #Required
    demand_quantity = var.occ_capacity_request_details_demand_quantity
    resource_name   = oci_usage_proxy_resource.test_resource.name
    resource_type   = var.occ_capacity_request_details_resource_type
    workload_type   = var.occ_capacity_request_details_workload_type

    #Optional
    actual_handover_quantity   = var.occ_capacity_request_details_actual_handover_quantity
    date_actual_handover       = var.occ_capacity_request_details_date_actual_handover
    date_expected_handover     = var.occ_capacity_request_details_date_expected_handover
    expected_handover_quantity = var.occ_capacity_request_details_expected_handover_quantity
  }
  display_name                = var.occ_capacity_request_display_name
  namespace                   = var.occ_capacity_request_namespace
  occ_availability_catalog_id = oci_capacity_management_occ_availability_catalog.test_occ_availability_catalog.id
  occ_capacity_request_id     = var.occ_capacity_request_occ_capacity_request_id
  region                      = var.occ_capacity_request_region

  #Optional
  defined_tags      = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.occ_capacity_request_defined_tags_value)
  description       = var.occ_capacity_request_description
  freeform_tags     = var.occ_capacity_request_freeform_tags
  lifecycle_details = var.occ_capacity_request_lifecycle_details
  patch_operations {
    #Required
    operation = var.occ_capacity_request_patch_operations_operation
    selection = var.occ_capacity_request_patch_operations_selection

    #Optional
    from          = var.occ_capacity_request_patch_operations_from
    position      = var.occ_capacity_request_patch_operations_position
    selected_item = var.occ_capacity_request_patch_operations_selected_item
    value         = var.occ_capacity_request_patch_operations_value
    values        = var.occ_capacity_request_patch_operations_values
  }
  request_state = var.occ_capacity_request_request_state
}

data "oci_capacity_management_occ_capacity_requests" "test_occ_capacity_requests" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name                = var.occ_capacity_request_display_name
  id                          = var.occ_capacity_request_id
  namespace                   = var.occ_capacity_request_namespace
  occ_availability_catalog_id = oci_capacity_management_occ_availability_catalog.test_occ_availability_catalog.id
}