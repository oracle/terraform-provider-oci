// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0


variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "topic_id" {}
variable "compartment_ocid" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_identity_tag_namespace" "tag-namespace1" {
  compartment_id = var.tenancy_ocid
  description    = "example tag namespace"
  name           = "examples-tag-namespace-all"
  is_retired = false
}

resource "oci_identity_tag" "tag1" {
  description      = "example tag"
  name             = "example-tag"
  tag_namespace_id = oci_identity_tag_namespace.tag-namespace1.id
  is_cost_tracking = false
  is_retired       = false
}

variable "news_frequency" {
  default = "WEEKLY"
}

variable "news_locale" {
  default = "EN" 
}

variable "news_report_name" {
  default = "Example_Report"
}

variable "news_report_description" {
  default = "Example Report Description"
}

variable "cp_resources" {
 default =  [ "HOST","DATABASE","EXADATA" ]
}


variable "news_report_defined_tags_value" {
  default = "value"
}

variable "news_report_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "resource_status" {
  default = "ENABLED"
}

// To Create a News Report
resource "oci_opsi_news_report" "test_news_report" {
  compartment_id    = var.compartment_ocid
  locale            = var.news_locale
  name              = var.news_report_name
  description       = var.news_report_description
  news_frequency    = var.news_frequency
  content_types {
    capacity_planning_resources = var.cp_resources 
  }
  ons_topic_id	    = var.topic_id
  freeform_tags     = var.news_report_freeform_tags
  status            = var.resource_status
}

variable "news_report_state" {
  default = ["ACTIVE"]
}

variable "news_report_status" {
  default = ["ENABLED"]
}

// List news reports
data  "oci_opsi_news_reports" "test_news_reports" {
  compartment_id = var.compartment_ocid
  state          = var.news_report_state
  status         = var.news_report_status
}

// Get a news report
data "oci_opsi_news_report" "test_news_report" {
  news_report_id = oci_opsi_news_report.test_news_report.id
}

