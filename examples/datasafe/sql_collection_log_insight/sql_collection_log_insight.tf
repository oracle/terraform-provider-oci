// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "data_safe_sql_collection_generate_fp_ocid" {}

variable "sql_collection_log_insight_time_ended" {
  default = "2024-01-01T00:00:00.000Z"
}

variable "sql_collection_log_insight_time_started" {
  default = "2023-01-01T00:00:00.000Z"
}

variable "sql_collection_log_insight_group_by" {
  default = "clientIp"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_data_safe_sql_collection_log_insights" "test_sql_collection_log_insights" {
  #Required
  sql_collection_id = var.data_safe_sql_collection_generate_fp_ocid
  time_ended = var.sql_collection_log_insight_time_ended
  time_started = var.sql_collection_log_insight_time_started

  #Optional
  group_by = var.sql_collection_log_insight_group_by
}