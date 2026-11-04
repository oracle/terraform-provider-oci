// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "container_repository_compartment_id_in_subtree" {
  default = false
}

variable "container_repository_defined_tags_value" {
  default = "value"
}

variable "container_repository_display_name" {
  default = "displayName"
}

variable "container_repository_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "container_repository_is_immutable" {
  default = false
}

variable "container_repository_is_public" {
  default = false
}

variable "container_repository_is_top_level" {
  default = false
}

variable "container_repository_multi_region_peers" {
  default = []
}

variable "container_repository_readme_content" {
  default = "content"
}

variable "container_repository_readme_format" {
  default = "TEXT_MARKDOWN"
}

variable "container_repository_repository_type" {
  default = "STANDARD"
}

variable "container_repository_state" {
  default = "AVAILABLE"
}

variable "container_repository_upstream_url" {
  default = "upstreamUrl"
}

variable "container_repository_upstream_username" {
  default = "upstreamUsername"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_artifacts_container_repository" "test_container_repository" {
  #Required
  compartment_id = var.compartment_id
  display_name   = var.container_repository_display_name

  #Optional
  defined_tags       = tomap({ oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name = var.container_repository_defined_tags_value })
  freeform_tags      = var.container_repository_freeform_tags
  is_immutable       = var.container_repository_is_immutable
  is_public          = var.container_repository_is_public
  is_top_level       = var.container_repository_is_top_level
  multi_region_peers = var.container_repository_multi_region_peers
  readme {
    #Required
    content = var.container_repository_readme_content
    format  = var.container_repository_readme_format
  }
  repository_type    = var.container_repository_repository_type
  upstream_secret_id = oci_vault_secret.test_secret.id
  upstream_url       = var.container_repository_upstream_url
  upstream_username  = var.container_repository_upstream_username
}

data "oci_artifacts_container_repositories" "test_container_repositories" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  compartment_id_in_subtree = var.container_repository_compartment_id_in_subtree
  display_name              = var.container_repository_display_name
  is_public                 = var.container_repository_is_public
  is_top_level              = var.container_repository_is_top_level
  repository_id             = oci_artifacts_repository.test_repository.id
  repository_type           = var.container_repository_repository_type
  state                     = var.container_repository_state
}

