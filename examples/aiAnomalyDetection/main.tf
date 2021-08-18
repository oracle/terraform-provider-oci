// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

// These variables would commonly be defined as environment variables or sourced in a .env file

variable "tenancy_ocid" {
  default = ""
}
variable "user_ocid" {
  default = ""
}
variable "fingerprint" {
  default = ""
}
variable "private_key_path" {
  default = ""
}
variable "region" {
  default = ""
}
variable "compartment_id" {
  default = ""
}

// PROJECT
variable "project_defined_tags_value" {
  default = "value"
}

variable "project_description" {
  default = "description"
}

variable "project_display_name" {
  default = "displayName"
}

variable "project_freeform_tags" {
  default = {
    "bar-key" = "value"
  }
}

variable "project_state" {
  default = "ACTIVE"
}

// PRIVATE ENDPOINT
variable "ai_private_endpoint_defined_tags_value" {
  default = "value"
}

variable "ai_private_endpoint_display_name" {
  default = "displayName"
}

variable "ai_private_endpoint_dns_zones" {
  default = []
}

variable "ai_private_endpoint_freeform_tags" {
  default = {
    "bar-key" = "value"
  }
}

variable "ai_private_endpoint_id" {
  default = "id"
}

variable "ai_private_endpoint_state" {
  default = "ACTIVE"
}

// DATA ASSET
variable "data_asset_data_source_details_bucket" {
  default = "mset-idp-test-datasets"
}

variable "data_asset_data_source_details_data_source_type" {
  default = "ORACLE_OBJECT_STORAGE"
}


variable "data_asset_data_source_details_namespace" {
  default = "ax3dvjxgkemg"
}

variable "data_asset_data_source_details_object" {
  default = "latest_training_data.json"
}

variable "data_asset_defined_tags_value" {
  default = "value"
}

variable "data_asset_description" {
  default = "description"
}

variable "data_asset_display_name" {
  default = "displayName"
}

variable "data_asset_freeform_tags" {
  default = {
    "bar-key" = "value"
  }
}

variable "data_asset_state" {
  default = "ACTIVE"
}

// MODEL
variable "model_defined_tags_value" {
  default = "value"
}

variable "model_description" {
  default = "description"
}

variable "model_display_name" {
  default = "displayName"
}

variable "model_freeform_tags" {
  default = {
    "bar-key" = "value"
  }
}

variable "model_model_training_details_data_asset_ids" {
  default = []
}

variable "model_model_training_details_target_fap" {
  default = 0.01
}

variable "model_model_training_details_training_fraction" {
  default = 0.7
}

variable "model_state" {
  default = "ACTIVE"
}


provider "oci" {
  tenancy_ocid = var.tenancy_ocid
  user_ocid = var.user_ocid
  fingerprint = var.fingerprint
  private_key_path = var.private_key_path
  region = var.region
}

//DEPENDENCIES
variable defined_tag_namespace_name {
  default = ""
}
resource "oci_identity_tag_namespace" "tag-namespace1" {
  #Required
  compartment_id = var.tenancy_ocid
  description = "example tag namespace"
  name = var.defined_tag_namespace_name != "" ? var.defined_tag_namespace_name : "example-tag-namespace-all"

  is_retired = false
}

resource "oci_identity_tag" "tag1" {
  #Required
  description = "example tag"
  name = "example-tag"
  tag_namespace_id = oci_identity_tag_namespace.tag-namespace1.id

  is_retired = false
}

resource "oci_core_subnet" "test_subnet" {
  availability_domain = lower(data.oci_identity_availability_domains.test_availability_domains.availability_domains[0].name)
  cidr_block = "10.0.0.0/24"
  compartment_id = var.compartment_id
  //defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}"
  dhcp_options_id = oci_core_vcn.test_vcn.default_dhcp_options_id
  display_name = "MySubnet"
  dns_label = "dnslabel"
  freeform_tags = {
    "Department" = "Finance"
  }
  lifecycle {
    ignore_changes = [
      defined_tags]
  }
  prohibit_internet_ingress = "false"
  prohibit_public_ip_on_vnic = "false"
  route_table_id = oci_core_vcn.test_vcn.default_route_table_id
  security_list_ids = [
    oci_core_vcn.test_vcn.default_security_list_id]
  vcn_id = oci_core_vcn.test_vcn.id
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block = "10.0.0.0/16"
  compartment_id = var.compartment_id
  //defined_tags = "${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}"
  display_name = "displayName"
  dns_label = "dnslabel"
  freeform_tags = {
    "Department" = "Finance"
  }
  lifecycle {
    ignore_changes = [
      defined_tags]
  }
}

data "oci_identity_availability_domains" "test_availability_domains" {
  compartment_id = var.tenancy_ocid
}

//PROJECT
resource "oci_ai_anomaly_detection_project" "test_project" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  //defined_tags  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.project_defined_tags_value)
  description = var.project_description
  display_name = var.project_display_name
  freeform_tags = var.project_freeform_tags
}

data "oci_ai_anomaly_detection_projects" "test_projects" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = var.project_display_name
  state = var.project_state
}

//PRIVATE ENDPOINT
resource "oci_ai_anomaly_detection_ai_private_endpoint" "test_ai_private_endpoint" {
  #Required
  compartment_id = var.compartment_id
  dns_zones = [
    oci_core_subnet.test_subnet.subnet_domain_name]
  subnet_id = oci_core_subnet.test_subnet.id

  #Optional
  //defined_tags  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.ai_private_endpoint_defined_tags_value)
  display_name = var.ai_private_endpoint_display_name
  freeform_tags = var.ai_private_endpoint_freeform_tags
}

data "oci_ai_anomaly_detection_ai_private_endpoints" "test_ai_private_endpoints" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = var.ai_private_endpoint_display_name
  state = var.ai_private_endpoint_state
}

//DATA ASSET
resource "oci_ai_anomaly_detection_data_asset" "test_data_asset" {
  #Required
  compartment_id = var.compartment_id
  data_source_details {
    #Required
    data_source_type = var.data_asset_data_source_details_data_source_type
    bucket = var.data_asset_data_source_details_bucket
    namespace = var.data_asset_data_source_details_namespace
    object = var.data_asset_data_source_details_object
  }
  project_id = oci_ai_anomaly_detection_project.test_project.id

  #Optional
  //defined_tags        = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.data_asset_defined_tags_value)
  description = var.data_asset_description
  display_name = var.data_asset_display_name
  freeform_tags = var.data_asset_freeform_tags
}

data "oci_ai_anomaly_detection_data_assets" "test_data_assets" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = var.data_asset_display_name
  state = var.data_asset_state
}

//MODEL
resource "oci_ai_anomaly_detection_model" "test_model" {
  #Required
  compartment_id = var.compartment_id
  model_training_details {
    #Required
    data_asset_ids = [
      oci_ai_anomaly_detection_data_asset.test_data_asset.id]

    #Optional
    target_fap = var.model_model_training_details_target_fap
    training_fraction = var.model_model_training_details_training_fraction
  }
  project_id = oci_ai_anomaly_detection_data_asset.test_data_asset.project_id

  #Optional
  //defined_tags  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.model_defined_tags_value)
  description = var.model_description
  display_name = var.model_display_name
  //freeform_tags = var.model_freeform_tags
}

data "oci_ai_anomaly_detection_models" "test_models" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = var.model_display_name
  project_id = oci_ai_anomaly_detection_project.test_project.id
  state = var.model_state
}


