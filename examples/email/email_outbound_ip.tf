// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "email_outbound_ip_assignment_state" {
  default = "AVAILABLE"
}

#dependent on region and provide this Ip before executing tests
variable "email_outbound_ip_outbound_ip" {
  default = ""
}

variable "email_outbound_ip_state" {
  default = "ACTIVE"
}

data "oci_email_email_outbound_ips" "test_email_outbound_ips" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  outbound_ip      = var.email_outbound_ip_outbound_ip
  state            = var.email_outbound_ip_state
}
