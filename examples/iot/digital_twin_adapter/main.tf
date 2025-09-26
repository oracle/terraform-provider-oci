// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "iot_domain_ocid" {}
variable "digital_twin_model_ocid" {}

variable "digital_twin_adapter_defined_tags_value" {
  default = "value"
}

variable "digital_twin_adapter_description" {
  default = "description"
}

variable "digital_twin_adapter_digital_twin_model_spec_uri" {
  default = "dtmi:example:device;1"
}

variable "digital_twin_adapter_display_name" {
  default = "displayName"
}

variable "digital_twin_adapter_freeform_tags" {
  default = { "Protocol" = "MQTT" }
}

variable "digital_twin_adapter_id" {
  default = "id"
}

variable "digital_twin_adapter_inbound_envelope_envelope_mapping_time_observed" {
  default = "$.time"
}

variable "digital_twin_adapter_inbound_envelope_reference_endpoint" {
  default = "telemetry/temperature"
}

variable "digital_twin_adapter_inbound_envelope_reference_payload_data" {
  default = {
    "time" = "<timestamp>"
    "temperature" = 98.6
  }
}

variable "digital_twin_adapter_inbound_envelope_reference_payload_data_format" {
  default = "JSON"
}

variable "digital_twin_adapter_inbound_routes_condition" {
  default = "*"
}

variable "digital_twin_adapter_inbound_routes_description" {
  default = "description"
}

variable "digital_twin_adapter_inbound_routes_payload_mapping" {
  default = {
    "$.temperature" = "$.temperature"
  }
}

variable "digital_twin_adapter_inbound_routes_reference_payload_data" {
  default = {
    "temperature" = 98.6
  }
}

variable "digital_twin_adapter_inbound_routes_reference_payload_data_format" {
  default = "JSON"
}

variable "digital_twin_adapter_state" {
  default = "ACTIVE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_iot_digital_twin_adapter" "test_digital_twin_adapter" {
  #Required
  iot_domain_id = var.iot_domain_ocid

  #Optional
  #defined_tags                = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.digital_twin_adapter_defined_tags_value)
  description                 = var.digital_twin_adapter_description
  digital_twin_model_id       = var.digital_twin_model_ocid
  #digital_twin_model_spec_uri = var.digital_twin_adapter_digital_twin_model_spec_uri ## either digital_twin_model_id or digital_twin_model_spec_uri
  display_name                = var.digital_twin_adapter_display_name
  freeform_tags               = var.digital_twin_adapter_freeform_tags
  inbound_envelope {
    #Required
    reference_endpoint = var.digital_twin_adapter_inbound_envelope_reference_endpoint

    #Optional
    envelope_mapping {

      #Optional
      time_observed = var.digital_twin_adapter_inbound_envelope_envelope_mapping_time_observed
    }
    reference_payload {
      #Required
      data        = var.digital_twin_adapter_inbound_envelope_reference_payload_data
      data_format = var.digital_twin_adapter_inbound_envelope_reference_payload_data_format
    }
  }
  inbound_routes {
    #Required
    condition = var.digital_twin_adapter_inbound_routes_condition

    #Optional
    description     = var.digital_twin_adapter_inbound_routes_description
    payload_mapping = var.digital_twin_adapter_inbound_routes_payload_mapping
    reference_payload {
      #Required
      data        = var.digital_twin_adapter_inbound_routes_reference_payload_data
      data_format = var.digital_twin_adapter_inbound_routes_reference_payload_data_format
    }
  }
}

data "oci_iot_digital_twin_adapters" "test_digital_twin_adapters" {
  #Required
  iot_domain_id = var.iot_domain_ocid

  #Optional
  digital_twin_model_id       = var.digital_twin_model_ocid
  digital_twin_model_spec_uri = var.digital_twin_adapter_digital_twin_model_spec_uri
  display_name                = var.digital_twin_adapter_display_name
  id                          = oci_iot_digital_twin_adapter.test_digital_twin_adapter.id
  state                       = var.digital_twin_adapter_state
}

