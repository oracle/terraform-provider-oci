// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "data_source_access_level" {
  default = "ACCESSIBLE"
}

variable "data_source_compartment_id_in_subtree" {
  default = true
}

variable "data_source_defined_tags_value" {
  default = "value"
}

variable "data_source_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "data_source_status" {
  default = "DISABLED"
}

//Has to be unique
variable "data_source_display_name" {
  default = "displayName"
}

variable "data_source_data_source_feed_provider" {
  default = "LOGGINGQUERY"
}

//Acceptable values come from LifecycleStateEnum
variable "data_source_state" {
  default = "ACTIVE"
}

variable "data_source_data_source_details_logging_query_type" {
  default = "INSIGHT"
}

variable "data_source_data_source_details_operator" {
  default = "GREATERTHANEQUALTO"
}

variable "data_source_data_source_details_additional_entities_count" {
  default = "2"
}

variable "data_source_data_source_details_interval_in_minutes" {
  default = "10"
}

variable "data_source_data_source_details_threshold" {
  default = "0"
}

variable "data_source_data_source_details_query" {
  default = "search \"ocid1.tenancy.oc1..aaaaaaaaqoggzsjut2u64wqliyd4eyd3dl4ipsu26lgqx4bihofnve5li5hq\" | isNotNull(data.eventName) | limit 1 | select data.eventName as cgkey01, data.message as cg01, data.resourceId as cg02"
}

variable "data_source_data_source_details_regions" {
  default = ["us-phoenix-1"]
}

variable "data_source_data_source_details_logging_query_details_key_entities_count" {
  default = "1"
}

variable "data_source_data_source_details_query_start_time_start_policy_type" {
  default = "ABSOLUTE_TIME_START_POLICY"
}

variable "data_source_data_source_details_query_start_time_query_start_time" {
  default = "2024-05-02T12:52:59.817Z"
}

provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = "terraform-federation-test"
  region              = var.region
  //version             = "5.39.0"
  /*
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  */
}

data "oci_cloud_guard_data_sources" "test_data_sources" {
  #Required
  compartment_id = var.tenancy_ocid

  #Optional
  state        = var.data_source_state
  display_name = var.data_source_display_name
}

resource "oci_cloud_guard_data_source" "test_data_source" {
  #Required
  compartment_id            = var.compartment_id
  display_name              = var.data_source_display_name
  data_source_feed_provider = var.data_source_data_source_feed_provider

  #Required
  data_source_details {
    data_source_feed_provider = var.data_source_data_source_feed_provider
    additional_entities_count = var.data_source_data_source_details_additional_entities_count
    interval_in_minutes       = var.data_source_data_source_details_interval_in_minutes
    logging_query_type        = var.data_source_data_source_details_logging_query_type
    operator                  = var.data_source_data_source_details_operator
    query                     = var.data_source_data_source_details_query
    regions                   = var.data_source_data_source_details_regions
    threshold                 = var.data_source_data_source_details_threshold

    logging_query_details {
      logging_query_type  = var.data_source_data_source_details_logging_query_type
      key_entities_count = var.data_source_data_source_details_logging_query_details_key_entities_count
    }

    query_start_time {
      start_policy_type = var.data_source_data_source_details_query_start_time_start_policy_type
      query_start_time = var.data_source_data_source_details_query_start_time_query_start_time
    }
  }

  #Optional
  defined_tags  = { "example-tag-namespace-all.example-tag" = var.data_source_defined_tags_value }
  freeform_tags = var.data_source_freeform_tags
  status        = var.data_source_status
}