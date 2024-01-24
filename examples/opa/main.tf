// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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

variable "compartment_id" {
}

variable "idcs_access_token" {

}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_opa_opa_instance" "test_opa_instance" {
  compartment_id = "${var.compartment_id}"
  consumption_model = "UCM"
  display_name = "displayName"
  idcs_at = "${var.idcs_access_token}"
  metering_type = "EXECUTION_PACK"
  shape_name = "PRODUCTION"
}

data "oci_opa_opa_instances" "test_opa_instances" {
  compartment_id = "${var.compartment_id}"
  display_name = "displayName2"
  filter {
    name = "id"
    values = ["${oci_opa_opa_instance.test_opa_instance.id}"]
  }
  id = "${oci_opa_opa_instance.test_opa_instance.id}"
  state = "ACTIVE"
}

data "oci_opa_opa_instance" "test_opa_instance" {
  opa_instance_id = "${oci_opa_opa_instance.test_opa_instance.id}"
}
