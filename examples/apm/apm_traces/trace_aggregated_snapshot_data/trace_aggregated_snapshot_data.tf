// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "apm_domain_id" {}

variable "trace_aggregated_snapshot_data_trace_key" {
  default = "traceKey"
}

variable "trace_aggregated_snapshot_data_server_name" {
  default = "serverName"
}

variable "trace_aggregated_snapshot_data_service_name" {
  default = "serviceName"
}

variable "trace_aggregated_snapshot_data_span_key" {
  default = "spanKey"
}

variable "trace_aggregated_snapshot_data_span_name" {
  default = "spanName"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_apm_traces_trace_aggregated_snapshot_data" "test_trace_aggregated_snapshot_data" {
  #Required
  apm_domain_id = var.apm_domain_id
  trace_key     = var.trace_aggregated_snapshot_data_trace_key

  #Optional
  server_name = var.trace_aggregated_snapshot_data_server_name
  service_name = var.trace_aggregated_snapshot_data_service_name
  span_key = var.trace_aggregated_snapshot_data_span_key
  span_name = var.trace_aggregated_snapshot_data_span_name
}

