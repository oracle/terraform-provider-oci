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

variable "integration_instance_idcs_access_token" {

}

variable "integration_instance_consumption_model" {
  default = "UCM"
}

variable allow_listed_http_vcn {
  default = ""
}

variable certificate_secret_id {
  default = ""
}

variable domain_id {
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
  display_name              = "instance4643"
  integration_instance_type = "STANDARDX"
  shape                     = "DEVELOPMENT"
  # shape                     = "PRODUCTION"
  is_byol                   = "false"
  message_packs             = "10"
  domain_id                 = var.domain_id
  # idcs_at                 = var.integration_instance_idcs_access_token

  #Optional
# For stand / enterprise type only
#  consumption_model = "${var.integration_instance_consumption_model}"
#  custom_endpoint {
#    certificate_secret_id = var.certificate_secret_id
#    hostname = "hostname.com"
#  }
#  freeform_tags = {
#    "bar-key" = "value"
#  }

#  is_file_server_enabled = true
#  is_visual_builder_enabled = true
#  state                  = "ACTIVE"

# For stand / enterprise type only
#  network_endpoint_details {
#    allowlisted_http_ips = ["10.0.0.0/28"]
#    allowlisted_http_vcns {
#      allowlisted_ips = ["0.0.0.0/0"]
#      id = "${var.allow_listed_http_vcn}"
#    }
#    is_integration_vcn_allowlisted = "false"
#    network_endpoint_type = "PUBLIC"
#  }

}

data "oci_integration_integration_instances" "test_integration_instances" {
  #Required
  compartment_id = var.compartment_id

  #Optional
#  display_name = "displayName"
#  state        = "Active"
}

data "oci_integration_integration_instance" "test_integration_instance" {
  #Required
  integration_instance_id = oci_integration_integration_instance.test_integration_instance.id
}

resource "oci_integration_integration_instance" "test_integration_instance_idcs" {
  #Required
  compartment_id            = var.compartment_id
  display_name              = "instance4643_idcs"
  integration_instance_type = "STANDARDX"
  shape                     = "DEVELOPMENT"
  # shape                     = "PRODUCTION"
  is_byol                   = "false"
  message_packs             = "10"
  idcs_at                   = var.integration_instance_idcs_access_token
}
