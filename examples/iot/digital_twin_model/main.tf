// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "iot_domain_ocid" {}

variable "digital_twin_model_defined_tags_value" {
  default = "value"
}

variable "digital_twin_model_description" {
  default = "description"
}

variable "spec_uri_starts_with" {
  default = "dtmi:example"
}

variable "digital_twin_model_display_name" {
  default = "displayName"
}

variable "digital_twin_model_freeform_tags" {
  default = { "Protocol" = "MQTT" }
}

variable "digital_twin_model_id" {
  default = "id"
}

variable "digital_twin_model_spec" {
  default = {
    "@id"      = "dtmi:example:device;1"
    "@type"    = "Interface"
    "@context" = "dtmi:dtdl:context;3"
    "displayName" = "IoT Device Model"
    "description" = "Represents a simple IoT device with temperature property."
    "contents" = [
      {
        "@type"       = "Property"
        "name"        = "temperature"
        "displayName" = "Temperature"
        "description" = "The current temperature reading of the device."
        "schema"      = "double"
      },
      {
        "@type" = "Relationship"
        "name"  = "connectedHumidity"
        "target"  = "dtmi:example:device;2"
        "properties" = [
          {
            "@type" = "Property"
            "name" = "connectionStrength"
            "schema" = "integer"
            "description" = "Strength of the connection between devices (0–100)."
          }
        ]
      }
    ]
  }
}

variable "digital_twin_model_state" {
  default = "ACTIVE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_iot_digital_twin_model" "test_digital_twin_model" {
  #Required
  iot_domain_id = var.iot_domain_ocid
  spec          = jsonencode(var.digital_twin_model_spec)

  #Optional
  #defined_tags  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.digital_twin_model_defined_tags_value)
  description   = var.digital_twin_model_description
  display_name  = var.digital_twin_model_display_name
  freeform_tags = var.digital_twin_model_freeform_tags
}

data "oci_iot_digital_twin_models" "test_digital_twin_models" {
  #Required
  iot_domain_id = var.iot_domain_ocid

  #Optional
  spec_uri_starts_with                    = var.spec_uri_starts_with
  display_name                            = var.digital_twin_model_display_name
  id                                      = oci_iot_digital_twin_model.test_digital_twin_model.id
  state                                   = var.digital_twin_model_state
}

data "oci_iot_digital_twin_model_spec" "test_digital_twin_model_spec" {
  #Required
  digital_twin_model_id = oci_iot_digital_twin_model.test_digital_twin_model.id
}

