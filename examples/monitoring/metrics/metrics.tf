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

variable "compartment_id" {
}

variable "metric_compartment_id_in_subtree" {
  default = false
}

variable "metric_dimension_filters" {
  default = "dimensionFilters"
}

variable "metric_group_by" {
  default = []
}

variable "metric_name" {
  default = "AcceptedConnections"
}

variable "metric_namespace" {
  default = "oci_lbaas"
}

variable "metric_resource_group" {
  default = ""
}

variable "metric_data_resource_group" {
  default = ""
}

variable "metric_data_compartment_id_in_subtree" {
  default = false
}

variable "metric_data_end_time" {
  default = "metricDataEndTimeStr"
}

variable "metric_data_namespace" {
  default = "oci_vcn"
}

variable "metric_data_query" {
  default = "VnicToNetworkPackets[4m].max()"
}

variable "metric_data_resolution" {
  default = "2m"
}

variable "metric_data_start_time" {
  default = "metricDataStartTimeStr"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_monitoring_metric_data" "test_metric_data" {
  #Required
  compartment_id = var.compartment_id
  namespace      = var.metric_data_namespace
  query          = var.metric_data_query

  #Optional
  compartment_id_in_subtree = var.metric_data_compartment_id_in_subtree
  resolution                = var.metric_data_resolution
  resource_group            = var.metric_data_resource_group
}

data "oci_monitoring_metrics" "test_metrics" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  compartment_id_in_subtree = var.metric_compartment_id_in_subtree
  name                      = var.metric_name
  namespace                 = var.metric_namespace
  resource_group            = var.metric_resource_group
}

output "metricData" {
  value = data.oci_monitoring_metric_data.test_metric_data.metric_data
}

output "metric" {
  value = data.oci_monitoring_metrics.test_metrics.metrics
}

