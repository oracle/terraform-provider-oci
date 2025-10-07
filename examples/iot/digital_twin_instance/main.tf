// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

###########################################################################
# This example assumes the following prerequisites:
# - The DTDL Models and Adapter have already been created.
#   You can refer to the model example for guidance.
###########################################################################

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "iot_domain_ocid" {}
variable "digital_twin_model_ocid" {}
variable "digital_twin_adapter_ocid" {}
variable "auth_ocid" {}

variable "digital_twin_instance_defined_tags_value" {
  default = "value"
}

variable "digital_twin_instance_description" {
  default = "description"
}

variable "digital_twin_instance_digital_twin_model_spec_uri" {
  // Refer model example
  default = "dtmi:example:device;1"
}

variable "digital_twin_instance_display_name" {
  default = "displayName"
}

variable "digital_twin_instance_external_key" {
  default = "externalKey"
}

variable "digital_twin_instance_freeform_tags" {
  default = { "Protocol" = "MQTT" }
}

variable "digital_twin_instance_id" {
  default = "id"
}

variable "digital_twin_instance_state" {
  default = "ACTIVE"
}

variable "digital_twin_instance_content_should_include_metadata" {
  default = true
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_iot_digital_twin_instance" "test_digital_twin_instance" {
  #Required
  auth_id       = var.auth_ocid
  iot_domain_id = var.iot_domain_ocid

  #Optional
  #defined_tags            = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.digital_twin_instance_defined_tags_value)
  description             = var.digital_twin_instance_description
  digital_twin_adapter_id = var.digital_twin_adapter_ocid
  display_name            = var.digital_twin_instance_display_name
  external_key            = var.digital_twin_instance_external_key
  freeform_tags           = var.digital_twin_instance_freeform_tags
}

data "oci_iot_digital_twin_instances" "test_digital_twin_instances" {
  #Required
  iot_domain_id = var.iot_domain_ocid

  #Optional
  digital_twin_model_id       = var.digital_twin_model_ocid
  digital_twin_model_spec_uri = var.digital_twin_instance_digital_twin_model_spec_uri
  display_name                = var.digital_twin_instance_display_name
  id                          = var.digital_twin_instance_id
  state                       = var.digital_twin_instance_state
}

data "oci_iot_digital_twin_instance_content" "test_digital_twin_instance_content" {
  #Required
  digital_twin_instance_id = oci_iot_digital_twin_instance.test_digital_twin_instance.id

  #Optional
  should_include_metadata = var.digital_twin_instance_content_should_include_metadata
}
