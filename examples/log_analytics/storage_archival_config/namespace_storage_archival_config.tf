// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

variable "namespace_storage_archival_config_archiving_configuration_active_storage_duration" {
  default = "P60D"
}

variable "namespace_storage_archival_config_archiving_configuration_archival_storage_duration" {
  default = "-1"
}

data "oci_log_analytics_namespaces" "test_namespaces" {
  compartment_id = var.tenancy_ocid
}

resource "oci_log_analytics_namespace_storage_archival_config" "test_namespace_storage_archival_config" {
  #Required
  archiving_configuration {

    #Optional
    active_storage_duration   = var.namespace_storage_archival_config_archiving_configuration_active_storage_duration
    archival_storage_duration = var.namespace_storage_archival_config_archiving_configuration_archival_storage_duration
  }
  namespace = data.oci_log_analytics_namespaces.test_namespaces.namespace_collection.0.items.0.namespace
}

data "oci_log_analytics_namespace_storage_archival_config" "test_namespace_storage_archival_config" {
  #Required
  namespace = data.oci_log_analytics_namespaces.test_namespaces.namespace_collection.0.items.0.namespace
}