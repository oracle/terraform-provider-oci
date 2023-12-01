// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "managed_instance_id" {
}

variable "managed_instance_module_name" {
}

variable "managed_instance_module_stream_name" {
}

variable "managed_instance_module_stream_status" {
  default = "ENABLED"
}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_osmanagement_managed_instance_module_streams" "test_managed_instance_module_streams" {
  #Required
  managed_instance_id = var.managed_instance_id

  #Optional
  compartment_id = var.compartment_id
  module_name    = var.managed_instance_module_name
  stream_name    = var.managed_instance_module_stream_name
  stream_status  = var.managed_instance_module_stream_status
}

output "test_managed_instance_module_streams" {
  value = {
    module_streams_on_managed_instance = data.oci_osmanagement_managed_instance_module_streams.test_managed_instance_module_streams.module_stream_on_managed_instances
  }  
}

