// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "email_outbound_ip_assignment_state" {
  default = "AVAILABLE"
}

variable "email_outbound_ip_outbound_ip" {
  default = "outboundIp"
}

variable "email_outbound_ip_state" {
  default = "AVAILABLE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_email_email_outbound_ips" "test_email_outbound_ips" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  assignment_state = var.email_outbound_ip_assignment_state
  outbound_ip      = var.email_outbound_ip_outbound_ip
  state            = var.email_outbound_ip_state
}

