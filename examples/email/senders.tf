// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "sender_email_address" {
  default = "johnsmithtester@example.com"
}

variable "sender_state" {
  default = "ACTIVE"
}

resource "oci_email_sender" "test_sender" {
  #Required
  compartment_id = var.compartment_ocid
  email_address  = var.sender_email_address

  timeouts {
    create = "10m"
  }
}

data "oci_email_senders" "test_senders" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  email_address = var.sender_email_address
  state         = var.sender_state

  filter {
    name   = "id"
    values = [oci_email_sender.test_sender.id]
  }
}

output "emailSenders" {
  value = data.oci_email_senders.test_senders.senders
}
