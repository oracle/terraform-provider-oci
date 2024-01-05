// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "software_source_id" {
}

variable "software_source_module_name" {
}

variable "software_source_module_stream_name" {
}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_osmanagement_software_source_module_stream" "test_software_source_module_stream" {
  #Required
  module_name        = var.software_source_module_name
  software_source_id = var.software_source_id
  stream_name        = var.software_source_module_stream_name
}

output "test_software_source_module_stream" {
  value = {
    software_source_id = data.oci_osmanagement_software_source_module_stream.test_software_source_module_stream.software_source_id
    module_name        = data.oci_osmanagement_software_source_module_stream.test_software_source_module_stream.module_name
    stream_name        = data.oci_osmanagement_software_source_module_stream.test_software_source_module_stream.stream_name
    architecture       = data.oci_osmanagement_software_source_module_stream.test_software_source_module_stream.architecture
    description        = data.oci_osmanagement_software_source_module_stream.test_software_source_module_stream.description
    is_default         = data.oci_osmanagement_software_source_module_stream.test_software_source_module_stream.is_default
    profiles           = data.oci_osmanagement_software_source_module_stream.test_software_source_module_stream.profiles
    packages           = data.oci_osmanagement_software_source_module_stream.test_software_source_module_stream.packages
  }
}

