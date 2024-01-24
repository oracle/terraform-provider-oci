// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "software_source_id" {
}

variable "software_source_module_name" {
}

variable "software_source_module_stream_name" {
}

variable "software_source_module_stream_profile_name" {
}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_osmanagement_software_source_stream_profiles" "test_software_source_stream_profiles" {
  #Required
  software_source_id = var.software_source_id

  #Optional
  compartment_id = var.compartment_id
  module_name    = var.software_source_module_name
  profile_name   = var.software_source_module_stream_profile_name
  stream_name    = var.software_source_module_stream_name
}

output "test_software_source_stream_profiles" {
  value = {
    module_stream_profiles = data.oci_osmanagement_software_source_stream_profiles.test_software_source_stream_profiles.module_stream_profiles
  }  
}

