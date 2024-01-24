// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * This example shows how to manage log analytics log groups
 */

# Create a log group with required parameters
resource "oci_log_analytics_log_analytics_log_group" "logGroupRequired" {
  compartment_id             = var.compartment_ocid
  namespace                  = data.oci_objectstorage_namespace.ns.namespace
  display_name               = "tf-loggroup-example-req"
}

# Get details of above created log group with required parameters
data "oci_log_analytics_log_analytics_log_group" "logGroupRequiredDetails" {
  namespace                  = data.oci_objectstorage_namespace.ns.namespace
  log_analytics_log_group_id = oci_log_analytics_log_analytics_log_group.logGroupRequired.id
}

# Create a log group with optional parameters
resource "oci_log_analytics_log_analytics_log_group" "logGroupOptional" {
  compartment_id             = var.compartment_ocid
  namespace                  = data.oci_objectstorage_namespace.ns.namespace
  display_name               = "tf-loggroup-example-opt"
  description                = "Log group with optional parameters"
  freeform_tags              = tomap({"servicegroup" = "test", "Dept" = "Devops"})
}

# Get details of above created log group with optional parameters
data "oci_log_analytics_log_analytics_log_group" "logGroupOptionalDetails" {
  namespace                  = data.oci_objectstorage_namespace.ns.namespace
  log_analytics_log_group_id = oci_log_analytics_log_analytics_log_group.logGroupOptional.id
}

# Get the list of log groups in a compartment
data "oci_log_analytics_log_analytics_log_groups" "logGroupsList" {
  compartment_id             = var.compartment_ocid
  namespace                  = data.oci_objectstorage_namespace.ns.namespace
}

# Get the count of log groups in a compartment
data "oci_log_analytics_log_analytics_log_groups_summary" "logGroupsSummary" {
  compartment_id             = var.compartment_ocid
  namespace                  = data.oci_objectstorage_namespace.ns.namespace
}