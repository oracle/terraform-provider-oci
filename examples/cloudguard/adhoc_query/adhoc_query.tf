// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "adhoc_query_access_level" {
  default = "ACCESSIBLE"
}

variable "adhoc_query_adhoc_query_details_adhoc_query_resources_region" {
  default = "us-phoenix-1"
}

variable "adhoc_query_adhoc_query_details_adhoc_query_resources_resource_ids" {
  default = ["ocid1.tenancy.oc1..aaaaaaaaqoggzsjut2u64wqliyd4eyd3dl4ipsu26lgqx4bihofnve5li5hq"]
}

variable "adhoc_query_adhoc_query_details_adhoc_query_resources_resource_type" {
  default = "TENANCY"
}

variable "adhoc_query_adhoc_query_details_query" {
  default = "select pid from processes"
}

variable "adhoc_query_adhoc_query_status" {
  default = "CREATING"
}

variable "adhoc_query_compartment_id_in_subtree" {
  default = true
}

variable "adhoc_query_defined_tags_value" {
  default = "value"
}

variable "adhoc_query_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "adhoc_query_time_ended_filter_query_param" {
  default = "2024-05-03T12:52:59.817Z"
}

variable "adhoc_query_time_started_filter_query_param" {
  default = "2024-05-03T10:52:59.817Z"
}



provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = "terraform-federation-test"
  region              = var.region
  //version             = "5.39.0"
  /*tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region*/
}

resource "oci_cloud_guard_adhoc_query" "test_adhoc_query" {
  #Required
  adhoc_query_details {
    #Required
    adhoc_query_resources {

      #Optional
      region        = var.adhoc_query_adhoc_query_details_adhoc_query_resources_region
      resource_ids  = var.adhoc_query_adhoc_query_details_adhoc_query_resources_resource_ids
      resource_type = var.adhoc_query_adhoc_query_details_adhoc_query_resources_resource_type
    }
    query = var.adhoc_query_adhoc_query_details_query
  }
  compartment_id = var.compartment_id

  #Optional
  defined_tags  = { "example-tag-namespace-all.example-tag" = var.adhoc_query_defined_tags_value}
  freeform_tags = var.adhoc_query_freeform_tags
}

data "oci_cloud_guard_adhoc_queries" "test_adhoc_queries" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  access_level                    = var.adhoc_query_access_level
  adhoc_query_status              = var.adhoc_query_adhoc_query_status
  compartment_id_in_subtree       = var.adhoc_query_compartment_id_in_subtree
  time_ended_filter_query_param   = var.adhoc_query_time_ended_filter_query_param
  time_started_filter_query_param = var.adhoc_query_time_started_filter_query_param
}
