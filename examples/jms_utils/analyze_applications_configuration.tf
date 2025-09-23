// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "analyze_applications_configuration_bucket" {
  default = null
}

variable "analyze_applications_configuration_namespace" {
  default = null
}

data "oci_jms_utils_analyze_applications_configuration" "test_analyze_applications_configuration" {

  #Optional
  bucket         = var.analyze_applications_configuration_bucket
  compartment_id = var.tenancy_ocid
  namespace      = var.analyze_applications_configuration_namespace
}
