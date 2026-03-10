// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0
variable "email_ip_pool_defined_tags_value" {
  default = "value"
}

variable "email_ip_pool_description" {
  default = "description"
}

variable "email_ip_pool_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "email_ip_pool_id" {
  default = "id"
}

variable "email_ip_pool_name" {
  default = "name"
}

variable "email_ip_pool_state" {
  default = "ACTIVE"
}

#dependent on region and provide this Ip before executing tests
variable "email_ip_pool_outbound_ips" {
  type    = list(string)
  default = [""]
}

resource "oci_email_email_ip_pool" "test_email_ip_pool" {
  #Required
  compartment_id = var.compartment_ocid
  name           = var.email_ip_pool_name
  outbound_ips   = var.email_ip_pool_outbound_ips

  #Optional
  description   = var.email_ip_pool_description
  freeform_tags = var.email_ip_pool_freeform_tags
}

data "oci_email_email_ip_pools" "test_email_ip_pools" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  id    = var.email_ip_pool_id
  name  = var.email_ip_pool_name
  state = var.email_ip_pool_state
}
