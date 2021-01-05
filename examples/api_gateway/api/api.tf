// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "api_content" {
  default = "openapi: 3.0.0\ninfo:\n  version: 1.0.0\n  title: Test API\n  license:\n    name: MIT\npaths:\n  /ping:\n    get:\n      responses:\n        '200':\n          description: OK"
}

variable "api_state" {
  default = "ACTIVE"
}

variable "api_display_name" {
  default = "test API definiton"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_apigateway_api" "test_api" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  content       = var.api_content
  defined_tags  = { "example-tag-namespace-all.example-tag" = "value" }

  display_name  = var.api_display_name
  freeform_tags = { "Department" = "Finance" }
}

data "oci_apigateway_apis" "test_apis" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = var.api_display_name
  state        = "ACTIVE"
}

# Validation result details

data "oci_apigateway_api_validation" "test_api_validation" {
  #Required
  api_id = oci_apigateway_api.test_api.id
}

# content of the Api definiton

data "oci_apigateway_api_content" "test_api_content" {
  #Required
  api_id = oci_apigateway_api.test_api.id
}

# generated deployment specification 

data "oci_apigateway_api_deployment_specification" "test_api_deployment_specification" {
  #Required
  api_id = oci_apigateway_api.test_api.id
}
