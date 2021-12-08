// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * This example shows how to get a log analytics entity topology details
 */

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}
variable "region" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

# Fetch namespace name from object store GET /n
data "oci_objectstorage_namespace" "ns" {
  compartment_id             = var.compartment_ocid
}

# Sample create entity with required parameters.
resource "oci_log_analytics_log_analytics_entity" "entityRequired" {
  compartment_id          = var.compartment_ocid
  namespace               = data.oci_objectstorage_namespace.ns.namespace
  name                    = "tf-entity-example-topo-req"
  entity_type_name        = "Host (Linux)"
}

# Get entity topo details of above created entity with required parameters
data "oci_log_analytics_log_analytics_entity_topology" "entityTopoRequiredDetails" {
  namespace               = data.oci_objectstorage_namespace.ns.namespace
  log_analytics_entity_id = oci_log_analytics_log_analytics_entity.entityRequired.id
}

# Get entity topo details of above created entity with required parameters
data "oci_log_analytics_log_analytics_entity_topology" "entityTopoOptionalDetails" {
  namespace               = data.oci_objectstorage_namespace.ns.namespace
  log_analytics_entity_id = oci_log_analytics_log_analytics_entity.entityRequired.id
  state                   = "ACTIVE"
}
