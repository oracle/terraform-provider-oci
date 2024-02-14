// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * This example shows how to fetch data related to recalls
 */

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}
variable "region" {}

provider "oci" {
  tenancy_ocid      = var.tenancy_ocid
  user_ocid         = var.user_ocid
  fingerprint       = var.fingerprint
  private_key_path  = var.private_key_path
  region            = var.region
}

locals {
  namespace         = data.oci_objectstorage_namespace.ns.namespace
  time_data_started = "2020-06-05T00:00:00.000Z"
  time_data_ended   = "2020-06-25T00:00:00.000Z"
}

# Fetch namespace name from object store GET /n
data "oci_objectstorage_namespace" "ns" {
  compartment_id    = var.compartment_ocid
}

# Fetch the recall count details for the tenancy
data "oci_log_analytics_namespace_storage_recall_count" "recall_count" {
  namespace         = local.namespace
}

# Fetch the recalled data size across all time
data "oci_log_analytics_namespace_storage_recalled_data_size" "recalled_data_size_all" {
  namespace         = local.namespace
}

# Fetch the recalled data size for a specific time range
data "oci_log_analytics_namespace_storage_recalled_data_size" "recalled_data_size_range" {
  namespace         = local.namespace
  time_data_started = local.time_data_started
  time_data_ended   = local.time_data_ended
}

# Fetch the overlapping recalls across all time
data "oci_log_analytics_namespace_storage_overlapping_recalls" "overlapping_recalls_all" {
  namespace         = local.namespace
}

# Fetch the overlapping recalls for a specific time range
data "oci_log_analytics_namespace_storage_overlapping_recalls" "overlapping_recalls_range" {
  namespace         = local.namespace
  time_data_started = local.time_data_started
  time_data_ended   = local.time_data_ended
}

output "recall_count_succeeded" {
  value = data.oci_log_analytics_namespace_storage_recall_count.recall_count.recall_succeeded
}

output "recalled_data_size_all_bytes" {
  value = data.oci_log_analytics_namespace_storage_recalled_data_size.recalled_data_size_all.recalled_data_in_bytes
}

output "recalled_data_size_range_bytes" {
  value = data.oci_log_analytics_namespace_storage_recalled_data_size.recalled_data_size_range.recalled_data_in_bytes
}

output "overlapping_recalls_all_status" {
  value = data.oci_log_analytics_namespace_storage_overlapping_recalls.overlapping_recalls_all.overlapping_recall_collection[0].items[0].status
}

output "overlapping_recalls_all_collection_id" {
  value = data.oci_log_analytics_namespace_storage_overlapping_recalls.overlapping_recalls_all.overlapping_recall_collection[0].items[0].collection_id
}

output "overlapping_recalls_all_recall_id" {
  value = data.oci_log_analytics_namespace_storage_overlapping_recalls.overlapping_recalls_all.overlapping_recall_collection[0].items[0].recall_id
}

output "overlapping_recalls_range_status" {
  value = data.oci_log_analytics_namespace_storage_overlapping_recalls.overlapping_recalls_range.overlapping_recall_collection[0].items[0].status
}

output "overlapping_recalls_range_collection_id" {
  value = data.oci_log_analytics_namespace_storage_overlapping_recalls.overlapping_recalls_range.overlapping_recall_collection[0].items[0].collection_id
}

output "overlapping_recalls_range_recall_id" {
  value = data.oci_log_analytics_namespace_storage_overlapping_recalls.overlapping_recalls_range.overlapping_recall_collection[0].items[0].recall_id
}
