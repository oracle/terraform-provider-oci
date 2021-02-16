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

variable "compartment_id" {
}

variable "integration_instance_idcs_access_token" {

}

variable "integration_instance_consumption_model" {
  default = "UCM"
}

variable allow_listed_http_vcn {
  default = ""
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_integration_integration_instance" "test_integration_instance" {
  #Required
  compartment_id            = var.compartment_id
  display_name              = "displayName"
  integration_instance_type = "STANDARD"
  is_byol                   = "false"
  message_packs             = "10"

  #Optional
  consumption_model = "${var.integration_instance_consumption_model}"
  custom_endpoint {
    hostname = "hostname.com"
  }
  freeform_tags = {
    "bar-key" = "value"
  }

  idcs_at                = var.integration_instance_idcs_access_token
  is_file_server_enabled = true
  state                  = "ACTIVE"

  network_endpoint_details {
    allowlisted_http_ips = ["172.16.0.239/32"]
    allowlisted_http_vcns {
      allowlisted_ips = ["172.16.0.239/32"]
      id = "${var.allow_listed_http_vcn}"
    }
    is_integration_vcn_allowlisted = "false"
    network_endpoint_type = "PUBLIC"
  }

}

data "oci_integration_integration_instances" "test_integration_instances" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = "displayName"
  state        = "Active"
}

data "oci_integration_integration_instance" "test_integration_instance" {
  #Required
  integration_instance_id = oci_integration_integration_instance.test_integration_instance.id
}
