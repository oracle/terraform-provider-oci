// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "ar_customer_transaction_id" {}
variable "invoice_line_id" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_onesubscription_invoices" "test_invoices" {
  #Required
  compartment_id  = var.compartment_id
  ar_customer_transaction_id = var.ar_customer_transaction_id
}

data "oci_onesubscription_invoice_line_computed_usages" "test_invoice_line_computed_usages" {
  #Required
  compartment_id  = var.compartment_id
  invoice_line_id = var.invoice_line_id
}
