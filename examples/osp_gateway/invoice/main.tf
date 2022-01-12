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

variable "compartment_ocid" {
}

variable "region" {
}

provider "oci" {
  region           = "${var.region}"
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
}

// invoices
data "oci_osp_gateway_invoices" "test_invoices" {
  compartment_id = "${var.tenancy_ocid}"
  osp_home_region = "${var.region}"
}

data "oci_osp_gateway_invoice" "test_invoice" {
  compartment_id = "${var.tenancy_ocid}"
  osp_home_region = "${var.region}"
  internal_invoice_id = "${lookup(data.oci_osp_gateway_invoices.test_invoices.invoice_collection.0.items[3], "internal_invoice_id")}"
}

data "oci_osp_gateway_invoices_invoice_lines" "test_invoice_lines" {
  compartment_id = "${var.tenancy_ocid}"
  osp_home_region = "${var.region}"
  internal_invoice_id = "${data.oci_osp_gateway_invoice.test_invoice.internal_invoice_id}"
}
