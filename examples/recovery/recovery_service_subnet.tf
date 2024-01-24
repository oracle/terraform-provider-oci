// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "recovery_service_subnet_defined_tags_value" {
  default = "value"
}

variable "recovery_service_subnet_display_name" {
  default = "displayName"
}

variable "recovery_service_subnet_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "recovery_service_subnet_id" {
  default = "id"
}

variable "recovery_service_subnet_state" {
  default = "ACTIVE"
}


resource "oci_recovery_recovery_service_subnet" "test_recovery_service_subnet" {
  #Required
  compartment_id = var.compartment_id
  display_name   = var.recovery_service_subnet_display_name
  subnet_id      = oci_core_subnet.test_subnet.id
  vcn_id         = oci_core_vcn.test_vcn.id

  #Optional
  freeform_tags = var.recovery_service_subnet_freeform_tags
}

data "oci_recovery_recovery_service_subnets" "test_recovery_service_subnets" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = var.recovery_service_subnet_display_name
  id           = var.recovery_service_subnet_id
  state        = var.recovery_service_subnet_state
  vcn_id       = oci_core_vcn.test_vcn.id
}

