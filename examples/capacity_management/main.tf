// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "occ_customer_group_display_name" {
  default = "displayName"
}

variable "occ_customer_group_id" {
  default = "id"
}

variable "occ_customer_group_status" {
  default = "ENABLED"
}

variable "occ_availability_catalog_base64encoded_catalog_details" {
  default = "RmluYWwgQ3VzdG9tZXIgT3JkZXIgRGF0ZSxDYXBhY2l0eSBIYW5kb3ZlciBEYXRlLFJlc291cmNlIFR5cGUsV29ya2xvYWQgVHlwZSxOYW1lLEF2YWlsYWJsZSBRdWFudGl0eSxVbml0CjIwMjUtMDItMjMsMjAyNS0wNC0wNSxDQVBBQ0lUWV9DT05TVFJBSU5ULFVTX1BST0QsVVMtQVNIQlVSTi0xLUFELTIsMTc1NSxTZXJ2ZXJzCjIwMjUtMDItMjMsMjAyNS0wNC0wNSxDQVBBQ0lUWV9DT05TVFJBSU5ULFJPVyxVUy1BU0hCVVJOLTEtQUQtMiwxNzU1LFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxVU19QUk9ELEJNLlN0YW5kYXJkMy42NCwxMDgsU2VydmVycwoyMDI1LTAyLTIzLDIwMjUtMDQtMDUsU0VSVkVSX0hXLFJPVyxCTS5TdGFuZGFyZDMuNjQsMTA4LFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxST1csQk0uU3RhbmRhcmQzLVdCLjY0LDE2MzA4LFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxST1csQk0uU3RhbmRhcmQyVC5FNC1XQi4xMjgsNzM4LFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxST1csQk0uU3RhbmRhcmQyVC5BMS1XQi4xNjAsNTgxNCxTZXJ2ZXJzCjIwMjUtMDItMjMsMjAyNS0wNC0wNSxTRVJWRVJfSFcsVVNfUFJPRCxCTS5TdGFuZGFyZDIuNTIsODQsU2VydmVycwoyMDI1LTAyLTIzLDIwMjUtMDQtMDUsU0VSVkVSX0hXLFJPVyxCTS5TdGFuZGFyZDIuNTIsODQsU2VydmVycwoyMDI1LTAyLTIzLDIwMjUtMDQtMDUsU0VSVkVSX0hXLFVTX1BST0QsQk0uU3RhbmRhcmQuRTVULUxNLjE5Miw4NDAsU2VydmVycwoyMDI1LTAyLTIzLDIwMjUtMDQtMDUsU0VSVkVSX0hXLFJPVyxCTS5TdGFuZGFyZC5FNVQtTE0uMTkyLDg0MCxTZXJ2ZXJzCjIwMjUtMDItMjMsMjAyNS0wNC0wNSxTRVJWRVJfSFcsVVNfUFJPRCxCTS5TdGFuZGFyZC5FNS4xOTIsMCxTZXJ2ZXJzCjIwMjUtMDItMjMsMjAyNS0wNC0wNSxTRVJWRVJfSFcsUk9XLEJNLlN0YW5kYXJkLkU1LjE5MiwwLFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxVU19QUk9ELEJNLlN0YW5kYXJkLkU0LjEyOCwzNTQ2LFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxST1csQk0uU3RhbmRhcmQuRTQuMTI4LDM1NDYsU2VydmVycwoyMDI1LTAyLTIzLDIwMjUtMDQtMDUsU0VSVkVSX0hXLFVTX1BST0QsQk0uU3RhbmRhcmQuQTEuMTYwLDMyNTgsU2VydmVycwoyMDI1LTAyLTIzLDIwMjUtMDQtMDUsU0VSVkVSX0hXLFJPVyxCTS5TdGFuZGFyZC5BMS4xNjAsMzI1OCxTZXJ2ZXJzCjIwMjUtMDItMjMsMjAyNS0wNC0wNSxTRVJWRVJfSFcsUk9XLEJNLlN0YW5kYXJkLkExLVdCLjE2MCw2ODA0LFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxVU19QUk9ELEJNLkRlbnNlSU8uRTVULjEyOCw0MzUsU2VydmVycwoyMDI1LTAyLTIzLDIwMjUtMDQtMDUsU0VSVkVSX0hXLFJPVyxCTS5EZW5zZUlPLkU1VC4xMjgsNDM1LFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxVU19QUk9ELEJNLkRlbnNlSU8uRTQuMTI4LDI3MDAsU2VydmVycwoyMDI1LTAyLTIzLDIwMjUtMDQtMDUsU0VSVkVSX0hXLFJPVyxCTS5EZW5zZUlPLkU0LjEyOCwyNzAwLFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxVU19QUk9ELEJNLkJpZ0RhdGEyLkU0LjEyOCAxNCBUQiwxNTAwLFNlcnZlcnMKMjAyNS0wMi0yMywyMDI1LTA0LTA1LFNFUlZFUl9IVyxST1csQk0uQmlnRGF0YTIuRTQuMTI4IDE0IFRCLDE1MDAsU2VydmVycw=="
}

variable "occ_availability_catalog_display_name" {
  default = "testAvailabilityCatalog"
}

variable "occ_availability_catalog_namespace" {
  default = "COMPUTE"
}

variable "occ_availability_catalog_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "occ_availability_catalog_description" {
  default = "This availability catalog is created via terraform provider"
}

variable "customer_compartment_id" {
  default = "ocid1.tenancy.oc1..aaaaaaaa3gmchdtrnbtbdxt23e4vg6teuxdz5p533353tpo3k3jmebhzoklq"
}

variable "occ_availability_catalog_metadata_details_format_version" {
  default = "V1"
}

variable "occ_availability_catalog_catalog_state" {
  default = "NOT_UPLOADED"
}

variable "occ_availability_catalog_id" {
  default = "id"
}

variable "occ_capacity_request_availability_domain" {
  default = "US-ASHBURN-1-AD-1"
}

variable "occ_capacity_request_date_expected_capacity_handover" {
  default = "2024/06/03"
}

variable "occ_capacity_request_defined_tags_value" {
  default = "value"
}

variable "occ_capacity_request_description" {
  default = "This is a test capacity request created via terraform provider"
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
  default = "US_PROD"
}

variable "occ_capacity_request_display_name" {
  default = "Test Request"
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
  default = "CREATED"
}

variable "occ_availability_catalog_occ_availability_resource_type" {
  default = "SERVER_HW"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_capacity_management_occ_customer_groups" "test_occ_customer_groups" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  #  display_name = var.occ_customer_group_display_name
  #  id           = var.occ_customer_group_id
  #  status       = var.occ_customer_group_status
}

resource "oci_capacity_management_occ_availability_catalog" "test_occ_availability_catalog" {
  #Required
  base64encoded_catalog_details = var.occ_availability_catalog_base64encoded_catalog_details
  compartment_id                = var.compartment_id
  display_name                  = var.occ_availability_catalog_display_name
  namespace                     = var.occ_availability_catalog_namespace
  occ_customer_group_id         = lookup(data.oci_capacity_management_occ_customer_groups.test_occ_customer_groups.occ_customer_group_collection.0.items[0], "id")

  #Optional
  #  defined_tags  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.occ_availability_catalog_defined_tags_value)
  description   = var.occ_availability_catalog_description
  freeform_tags = var.occ_availability_catalog_freeform_tags
  metadata_details {
    #Required
    format_version = var.occ_availability_catalog_metadata_details_format_version
  }
}


data "oci_capacity_management_occ_availability_catalogs" "published_catalog_collection" {
  compartment_id = var.customer_compartment_id
}

data "oci_capacity_management_occ_availability_catalog_occ_availabilities" "published_occ_availability_catalog_occ_availabilities" {
  #Required
  occ_availability_catalog_id = lookup(data.oci_capacity_management_occ_availability_catalogs.published_catalog_collection.occ_availability_catalog_collection.0.items[0], "id")

  #Optional
#  date_expected_capacity_handover = var.occ_availability_catalog_occ_availability_date_expected_capacity_handover
#  resource_name                   = oci_usage_proxy_resource.test_resource.name
  resource_type                   = var.occ_availability_catalog_occ_availability_resource_type
#  workload_type                   = var.occ_availability_catalog_occ_availability_workload_type
}

resource "oci_capacity_management_occ_capacity_request" "test_occ_capacity_request" {
  #Required
  availability_domain             = var.occ_capacity_request_availability_domain
  compartment_id                  = var.customer_compartment_id
#  date_expected_capacity_handover = lookup(data.oci_capacity_management_occ_availability_catalog_occ_availabilities.published_occ_availability_catalog_occ_availabilities.occ_availability_collection.0.items[0], "date_expected_capacity_handover")
#  date_expected_capacity_handover = formatdate("YYYY-MM-DDTHH:MM:SSZ",data.oci_capacity_management_occ_availability_catalog_occ_availabilities.published_occ_availability_catalog_occ_availabilities.occ_availability_collection.0.items[0].date_expected_capacity_handover)
   date_expected_capacity_handover = "2124-06-15T00:00:00.00Z"

  details {
    #Required
    demand_quantity = var.occ_capacity_request_details_demand_quantity
    resource_name   = lookup(data.oci_capacity_management_occ_availability_catalog_occ_availabilities.published_occ_availability_catalog_occ_availabilities.occ_availability_collection.0.items[0], "resource_name")
    resource_type   = lookup(data.oci_capacity_management_occ_availability_catalog_occ_availabilities.published_occ_availability_catalog_occ_availabilities.occ_availability_collection.0.items[0], "resource_type")
    workload_type   = lookup(data.oci_capacity_management_occ_availability_catalog_occ_availabilities.published_occ_availability_catalog_occ_availabilities.occ_availability_collection.0.items[0], "workload_type")

    #Optional
#    actual_handover_quantity   = var.occ_capacity_request_details_actual_handover_quantity
#    date_actual_handover       = var.occ_capacity_request_details_date_actual_handover
#    date_expected_handover     = var.occ_capacity_request_details_date_expected_handover
#    expected_handover_quantity = var.occ_capacity_request_details_expected_handover_quantity
  }
  display_name                = var.occ_capacity_request_display_name
  namespace                   = var.occ_capacity_request_namespace
  occ_availability_catalog_id = lookup(data.oci_capacity_management_occ_availability_catalogs.published_catalog_collection.occ_availability_catalog_collection.0.items[0], "id")
#  occ_capacity_request_id     = var.occ_capacity_request_occ_capacity_request_id
  region                      = var.occ_capacity_request_region

  #Optional
#  defined_tags      = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.occ_capacity_request_defined_tags_value)
  description       = var.occ_capacity_request_description
  freeform_tags     = var.occ_capacity_request_freeform_tags
  lifecycle_details = var.occ_capacity_request_lifecycle_details
#  patch_operations {
#    #Required
#    operation = var.occ_capacity_request_patch_operations_operation
#    selection = var.occ_capacity_request_patch_operations_selection
#
#    #Optional
#    from          = var.occ_capacity_request_patch_operations_from
#    position      = var.occ_capacity_request_patch_operations_position
#    selected_item = var.occ_capacity_request_patch_operations_selected_item
#    value         = var.occ_capacity_request_patch_operations_value
#    values        = var.occ_capacity_request_patch_operations_values
#  }
  request_state = var.occ_capacity_request_request_state
}
