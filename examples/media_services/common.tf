// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0


variable "tenancy_ocid" {}

variable "user_ocid" {}

variable "fingerprint" {}

variable "private_key_path" {}

variable "region" {}

variable "compartment_id" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

variable "media_asset_state" {
  default = "ACTIVE"
}

variable "media_asset_type" {
  default = "AUDIO"
}

variable "defined_tags_value" {
  default = "value"
}

variable "display_name" {
  default = "displayName"
}

variable "freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "id" {
  default = "id"
}

variable "active_state" {
  default = "ACTIVE"
}

variable "accepted_state" {
  default = "ACCEPTED"
}

variable "locks_type" {
  default = "FULL"
}

variable "locks_message" {
  default = "message"
}

variable "is_lock_override" {
  default = true
}

resource "oci_identity_tag_namespace" "tag-namespace1" {
  compartment_id = var.tenancy_ocid
  description    = "example tag namespace"
  name           = "examples-tag-namespace-all"
  is_retired = false
}

resource "oci_identity_tag" "tag1" {
  description      = "example tag"
  name             = "example-tag"
  tag_namespace_id = oci_identity_tag_namespace.tag-namespace1.id
  is_cost_tracking = false
  is_retired       = false
}
