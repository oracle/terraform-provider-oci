// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}

variable "container_repository_compartment_id_in_subtree" {
  default = false
}

variable "container_repository_is_immutable" {
  default = false
}

variable "container_repository_is_public" {
  default = false
}

variable "container_repository_readme_content" {
  default = "content"
}

variable "container_repository_readme_format" {
  default = "TEXT_MARKDOWN"
}

variable "container_repository_state" {
  default = "AVAILABLE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

// repository displayName needs to be unique within a tenant, so generate random string here to avoid collision
resource "random_string" "container_repository_display_name" {
  length  = 5
  number  = false
  special = false
  upper = false
}

resource "oci_artifacts_container_repository" "test_container_repository" {
  #Required
  compartment_id = var.compartment_ocid
  display_name   = random_string.container_repository_display_name.result

  #Optional
  is_immutable = var.container_repository_is_immutable
  is_public    = var.container_repository_is_public
  readme {
    #Required
    content = var.container_repository_readme_content
    format  = var.container_repository_readme_format
  }
}

data "oci_artifacts_container_repositories" "test_container_repositories" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  compartment_id_in_subtree = var.container_repository_compartment_id_in_subtree
  is_public                 = var.container_repository_is_public
  repository_id             = oci_artifacts_container_repository.test_container_repository.id
  state                     = var.container_repository_state
}
