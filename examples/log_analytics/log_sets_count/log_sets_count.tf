// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * This example shows how to query the count of log sets
 */

variable "compartment_ocid" {}

# Fetch namespace name from object store GET /n
data "oci_objectstorage_namespace" "ns" {
  compartment_id             = var.compartment_ocid
}

# Get the count of total log sets by namespace
data "oci_log_analytics_log_sets_count" "logSetsCount" {
  namespace                  = data.oci_objectstorage_namespace.ns.namespace
}