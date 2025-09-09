// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "ai_data_platform_ai_data_platform_type" {
default = "aiDataPlatformType"
}

variable "ai_data_platform_display_name" {
default = "displayName"
}

variable "ai_data_platform_workspace_name" {
  default = "workspaceName"
}

variable "ai_data_platform_exclude_lifecycle_state" {
default = "CREATING"
}

variable "ai_data_platform_freeform_tags" {
default = { "Department" = "Finance" }
}

variable "ai_data_platform_include_legacy" {
default = "true"
}

variable "ai_data_platform_state" {
default = "ACTIVE"
}

variable "ai_data_platform_system_tags" {
default = "value"
}

provider "oci" {
tenancy_ocid     = var.tenancy_ocid
user_ocid        = var.user_ocid
fingerprint      = var.fingerprint
private_key_path = var.private_key_path
region           = var.region
}

resource "oci_ai_data_platform_ai_data_platform" "test_ai_data_platform" {
#Required
compartment_id = var.compartment_id

#Optional
ai_data_platform_type  = var.ai_data_platform_ai_data_platform_type
default_workspace_name = var.ai_data_platform_workspace_name
#defined_tags           = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.ai_data_platform_defined_tags_value)
display_name           = var.ai_data_platform_display_name
freeform_tags          = var.ai_data_platform_freeform_tags
}

data "oci_ai_data_platform_ai_data_platforms" "test_ai_data_platforms" {

depends_on = [oci_ai_data_platform_ai_data_platform.test_ai_data_platform]

#Optional
compartment_id          = var.compartment_id
display_name            = var.ai_data_platform_display_name
exclude_lifecycle_state = var.ai_data_platform_exclude_lifecycle_state
id                      = oci_ai_data_platform_ai_data_platform.test_ai_data_platform.id
include_legacy          = var.ai_data_platform_include_legacy
state                   = var.ai_data_platform_state
}

