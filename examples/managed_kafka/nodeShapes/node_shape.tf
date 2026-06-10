// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "region" {}
variable "compartment_id" {}

variable "node_shape_name" {
  default = "name"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  region           = var.region
}

data "oci_managed_kafka_node_shapes" "test_node_shapes" {

  #Optional
  compartment_id = var.compartment_id
  name           = var.node_shape_name
}

