// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "java_download_token_description" {
  default = "Example token description for script friendly download"
}

variable "java_download_token_display_name" {
  default = "Unique-displayName-in-a-tenancy"
}

variable "java_download_token_id" {
  default = "id"
}

variable "java_download_token_is_default" {
  default = false
}

variable "java_download_token_java_version" {
  default = "11"
}

variable "java_download_token_license_type" {
  default = ["OTN"]
}

variable "java_download_token_state" {
  default = "ACTIVE"
}

variable "java_download_token_time_expires" {
  default = "2024-09-26T15:58:25.748Z"
}

variable "java_download_token_value" {
  default = "value"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_jms_java_downloads_java_download_token" "test_java_download_token" {
  #Required
  compartment_id = var.tenancy_ocid
  description    = var.java_download_token_description
  display_name   = var.java_download_token_display_name
  java_version   = var.java_download_token_java_version
  license_type   = var.java_download_token_license_type
  time_expires   = var.java_download_token_time_expires

  #Optional
  is_default    = var.java_download_token_is_default
}

data "oci_jms_java_downloads_java_download_tokens" "test_java_download_tokens" {
  #Required
  compartment_id = var.tenancy_ocid

  #Optional
  display_name   = var.java_download_token_display_name
  family_version = var.java_download_token_java_version
  id             = var.java_download_token_id
  search_by_user = var.user_ocid
  state          = var.java_download_token_state
  value          = var.java_download_token_value
}

