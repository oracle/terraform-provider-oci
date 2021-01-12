// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
}

variable "oda_instance_state" {
  default = "INACTIVE"
}

variable "compartment_ocid" {
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_oda_oda_instance" "TFOdaInstance" {
  compartment_id = var.compartment_ocid
  shape_name     = "DEVELOPMENT"
  description    = "test instance"
  display_name   = "TestInstance"

  #Optional
  state = var.oda_instance_state
}

data "oci_oda_oda_instances" "TFOdaInstances" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name = oci_oda_oda_instance.TFOdaInstance.display_name
  #state        = var.oda_instance_state
}

data "oci_oda_oda_instance" "TFOdaInstance" {
  #Required
  oda_instance_id = oci_oda_oda_instance.TFOdaInstance.id
}

