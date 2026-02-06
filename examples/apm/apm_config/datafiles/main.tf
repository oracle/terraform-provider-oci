// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" { }

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


variable "data_file_content_disposition" {
  default = "contentDisposition"
}

variable "data_file_content_encoding" {
  default = "contentEncoding"
}

variable "data_file_content_language" {
  default = "contentLanguage"
}

variable "data_file_content_md5" {
  default = "CY9rzUYh03PK3k6DJie09g=="
}

variable "data_file_content_type" {
  default = "text/plain"
}

variable "data_file_apm_type" {
  default = "apmType"
}

variable "content" {
  default = "test"
}

variable "data_file_metadata" {
  default = {
    "key": "value"
  }
}

variable "data_file_name" {
  default = "name"
}

variable "data_file_time_last_modified_after" {
  default = "2006-01-02T15:04:05Z"
}

variable "data_file_time_last_modified_before" {
  default = "4001-01-02T15:04:05Z"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
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

resource "oci_apm_config_data_file" "test_data_file" {
  #Required
  content = var.content
  apm_domain_id      = oci_apm_apm_domain.test_apm_domain.id
  apm_type           = var.data_file_apm_type
  data_file_name     = var.data_file_name

  #Optional
  content_disposition = var.data_file_content_disposition
  content_encoding    = var.data_file_content_encoding
  content_language    = var.data_file_content_language
  content_md5         = var.data_file_content_md5
  content_type        = var.data_file_content_type
  metadata            = var.data_file_metadata
}

data "oci_apm_config_data_files" "test_data_files" {
  #Required
  apm_domain_id = oci_apm_apm_domain.test_apm_domain.id

  #Optional
  apm_type                  = var.data_file_apm_type
  metadata                  = var.data_file_metadata
  name                      = var.data_file_name
  time_last_modified_after  = var.data_file_time_last_modified_after
  time_last_modified_before = var.data_file_time_last_modified_before
}

