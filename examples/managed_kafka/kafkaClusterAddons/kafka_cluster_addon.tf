// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "region" {}

variable "kafka_cluster_addon_addon_type" {
  default = "PUBLICCONNECTIVITY"
}

variable "kafka_cluster_addon_authentication_mechanism" {
  default = "SASL"
}

variable "kafka_cluster_addon_description" {
  default = "description"
}

variable "kafka_cluster_addon_name" {
  default = "test-terraform-name-santosh"
}

variable "kafka_cluster_addon_network_cidrs" {
  default = ["160.34.0.0/16"]
}

variable "kafka_cluster_addon_state" {
  default = "ACTIVE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  region           = var.region
}

resource "oci_managed_kafka_kafka_cluster_addon" "test_kafka_cluster_addon" {
  #Required
  addon_type               = var.kafka_cluster_addon_addon_type
  authentication_mechanism = var.kafka_cluster_addon_authentication_mechanism
  kafka_cluster_id         = oci_managed_kafka_kafka_cluster.test_kafka_cluster.id
  name                     = var.kafka_cluster_addon_name
  network_cidrs            = var.kafka_cluster_addon_network_cidrs

  #Optional
  description = var.kafka_cluster_addon_description

  timeouts {
    create = "120m"
    delete = "120m"
  }
}

data "oci_managed_kafka_kafka_cluster_addons" "test_kafka_cluster_addons" {
  #Required
  kafka_cluster_id = oci_managed_kafka_kafka_cluster.test_kafka_cluster.id

  #Optional
  name  = var.kafka_cluster_addon_name
  state = var.kafka_cluster_addon_state
}

