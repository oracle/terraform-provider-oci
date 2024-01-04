// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

variable "apm_domain_description" {
  default = "description"
}

variable "apm_domain_display_name" {
  default = "displayName"
}

variable "apm_domain_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "apm_domain_is_free_tier" {
  default = false
}

variable "apm_domain_state" {
  default = "ACTIVE"
}

variable "data_key_data_key_type" {
  default = "PRIVATE"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}


variable "on_premise_vantage_point_display_name" {
  default = "displayName"
}

variable "on_premise_vantage_point_name" {
  default = "OPVP-name"
}


resource "oci_apm_apm_domain" "test_apm_domain" {
  #Required
  compartment_id = var.compartment_ocid
  display_name   = var.apm_domain_display_name

  #Optional
  description   = var.apm_domain_description
  freeform_tags = var.apm_domain_freeform_tags
  is_free_tier  = var.apm_domain_is_free_tier
}

resource "oci_apm_synthetics_on_premise_vantage_point" "test_on_premise_vantage_point" {
  #Required
  apm_domain_id = oci_apm_apm_domain.test_apm_domain.id
  name         = var.on_premise_vantage_point_name
}


data "oci_apm_synthetics_on_premise_vantage_points" "test_on_premise_vantage_points" {
  #Required
  apm_domain_id = oci_apm_apm_domain.test_apm_domain.id

  #Optional
  name         = var.on_premise_vantage_point_name
}

variable "on_premise_vantage_point_worker_name" {
  default = "worker"
}

variable "on_premise_vantage_point_worker_version" {
  default = "1.2.4"
}

variable "on_premise_vantage_point_worker_resource_principal_token_public_key" {
  default = "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA0GuJMVpuYE3l2pAe4mwrB73pExN42hG5OkbiJimrSbSHBffng6NacHz4hX8Ri9WhuQSq51nXsGjixnVvjnI3RzgenAaLVrf48a8RmS5D0pwrjshkf5Vt/hSXYL2lVUToGTUdOzXb5ZAH6BN9SE+LPEeBl6QnXn90teMXeVPnarg9WE1LMf8eNoD3PRaXEa9i3Q0Q2/3cfXVX1MhHk5wi/fUKsnbTjy67a49vvC3DKbakw76q4lrdtvp7M5EKN+paD0qg64wRKn8/bCYvI/tjM+LufvSLJJSj7KQs83t5xKBK60FVRUK/d3bRdilb8XnezBSGSdPDY9fL6yn0z8UORQIDAQAB\n-----END PUBLIC KEY-----"
}


resource "oci_apm_synthetics_on_premise_vantage_point_worker" "test_on_premise_vantage_point_worker" {
  #Required
  apm_domain_id = oci_apm_apm_domain.test_apm_domain.id
  on_premise_vantage_point_id = oci_apm_synthetics_on_premise_vantage_point.test_on_premise_vantage_point.id
  name         = var.on_premise_vantage_point_worker_name
  version = var.on_premise_vantage_point_worker_version
  resource_principal_token_public_key = var.on_premise_vantage_point_worker_resource_principal_token_public_key
}

data "oci_apm_synthetics_on_premise_vantage_point_workers" "test_on_premise_vantage_point_workers" {
  #Required
  apm_domain_id = oci_apm_apm_domain.test_apm_domain.id
  on_premise_vantage_point_id = oci_apm_synthetics_on_premise_vantage_point.test_on_premise_vantage_point.id
  name         = var.on_premise_vantage_point_name

  #Optional
}
