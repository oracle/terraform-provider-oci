// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0
variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "authz_compartment_id" {
}

variable "api_metadata_by_entity_type_display_name" {
  default = "displayName"
}

variable "api_metadata_by_entity_type_resource_type" {
  default = "EXADATAINFRASTRUCTURE"
}

variable "api_metadata_by_entity_type_state" {
  default = "ACTIVE"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_apiaccesscontrol_api_metadata_by_entity_types" "test_api_metadata_by_entity_types" {
  compartment_id = var.authz_compartment_id
  #Optional
  display_name   = var.api_metadata_by_entity_type_display_name
  resource_type  = var.api_metadata_by_entity_type_resource_type
  state          = var.api_metadata_by_entity_type_state
}

