// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

###########################################################################
# This example assumes the following prerequisites:
# - The DTDL Models and Adapter have already been created.
#   You can refer to the model example for guidance.
# - The source and target Digital Twin instances have already been created.
###########################################################################

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "iot_domain_ocid" {}
variable "digital_twin_source_instance_ocid" {}
variable "digital_twin_target_instance_ocid" {}

variable "digital_twin_relationship_content" {
  default = {
    "connectionStrength" = 98
  }
}

variable "digital_twin_relationship_content_path" {
  default = "connectedHumidity"
}

variable "digital_twin_relationship_defined_tags_value" {
  default = "value"
}

variable "digital_twin_relationship_description" {
  default = "description"
}

variable "digital_twin_relationship_display_name" {
  default = "displayName"
}

variable "digital_twin_relationship_freeform_tags" {
  default = { "Protocol" = "MQTT" }
}

variable "digital_twin_relationship_id" {
  default = "id"
}

variable "digital_twin_relationship_state" {
  default = "ACTIVE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_iot_digital_twin_relationship" "test_digital_twin_relationship" {
  #Required
  content_path                    = var.digital_twin_relationship_content_path
  iot_domain_id                   = var.iot_domain_ocid
  source_digital_twin_instance_id = var.digital_twin_source_instance_ocid
  target_digital_twin_instance_id = var.digital_twin_target_instance_ocid

  #Optional
  content       = jsonencode(var.digital_twin_relationship_content)
  #defined_tags  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.digital_twin_relationship_defined_tags_value)
  description   = var.digital_twin_relationship_description
  display_name  = var.digital_twin_relationship_display_name
  freeform_tags = var.digital_twin_relationship_freeform_tags
}

data "oci_iot_digital_twin_relationships" "test_digital_twin_relationships" {
  #Required
  iot_domain_id = var.iot_domain_ocid

  #Optional
  content_path                    = var.digital_twin_relationship_content_path
  display_name                    = var.digital_twin_relationship_display_name
  id                              = var.digital_twin_relationship_id
  source_digital_twin_instance_id = var.digital_twin_source_instance_ocid
  state                           = var.digital_twin_relationship_state
  target_digital_twin_instance_id = var.digital_twin_target_instance_ocid
}

