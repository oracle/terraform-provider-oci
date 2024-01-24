// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_stack_monitoring_process_set" "test_process_set" {
  #Required
  compartment_id = var.compartment_ocid
  display_name   = "javaProcessSet"
  specification {
    #Required
    items {
      #Required
      process_command            = "java"
      
      #Optional
      label                      = "javaProcesses"
    }
  }
}

data "oci_stack_monitoring_process_sets" "test_process_sets" {
  #Required
  compartment_id = var.compartment_ocid
}

data "oci_stack_monitoring_process_set" "test_process_set" {
  #Required
  process_set_id = oci_stack_monitoring_process_set.test_process_set.id
}
