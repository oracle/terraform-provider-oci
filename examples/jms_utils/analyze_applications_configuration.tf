// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "analyze_applications_configuration_bucket" {
  default = null
}

variable "analyze_applications_configuration_namespace" {
  default = null
}

resource "oci_jms_utils_analyze_applications_configuration" "test_analyze_applications_configuration" {

  compartment_id = var.tenancy_ocid

  #Optional
  bucket         = var.analyze_applications_configuration_bucket
  namespace      = var.analyze_applications_configuration_namespace
}

data "oci_jms_utils_analyze_applications_configuration" "test_analyze_applications_configurations" {

  compartment_id = var.tenancy_ocid
}