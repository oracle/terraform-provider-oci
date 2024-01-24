// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

// These variables would commonly be defined as environment variables or sourced in a .env file

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "compartment_ocid" {
}

variable "region" {
}

provider "oci" {
  region           = var.region
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
}

data "oci_datascience_notebook_session_shapes" "tf_notebook_session_shapes" {
  compartment_id = var.compartment_ocid
}

variable "artifact_content_length" {
  default = "6954"
}

variable "model_artifact" {
}

variable "shape" {
  default = "VM.Standard2.1"
}

resource "oci_datascience_project" "tf_project" {
  #Required
  compartment_id = var.compartment_ocid
  #Optional
  #defined_tags  = {"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.project_defined_tags_value}"}
  #description   = var.project_description
  #display_name  = var.project_display_name
  #freeform_tags = var.project_freeform_tags
}

data "oci_datascience_projects" "tf_projects" {
  #Required
  compartment_id = var.compartment_ocid
  #Optional
  #created_by   = var.project_created_by
  #display_name = var.project_display_name
  #id           = var.project_id
  #state        = var.project_state
}

resource "oci_datascience_notebook_session" "tf_notebook_session" {
  #Required
  compartment_id = var.compartment_ocid

  notebook_session_configuration_details {
    #Required
    shape     = var.shape
    subnet_id = oci_core_subnet.tf_subnet.id
    #Optional
    #block_storage_size_in_gbs = var.notebook_session_notebook_session_configuration_details_block_storage_size_in_gbs
    #notebook_session_shape_config_details {
    #  #Required
    #  ocpus = var.notebook_session_notebook_session_configuration_details_notebook_session_shape_config_details_opcus
    #  memory_in_gbs = var.notebook_session_notebook_session_configuration_details_notebook_session_shape_config_details_memory_in_gbs
    #}
  }

  project_id = oci_datascience_project.tf_project.id
  #Optional
  #defined_tags  = {"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.notebook_session_defined_tags_value}"}
  #display_name  = var.notebook_session_display_name
  #freeform_tags = var.notebook_session_freeform_tags
}

data "oci_datascience_notebook_sessions" "tf_notebook_sessions" {
  #Required
  compartment_id = var.compartment_ocid
  #Optional
  #created_by   = var.notebook_session_created_by
  #display_name = var.notebook_session_display_name
  #id           = var.notebook_session_id
  #project_id   = oci_datascience_project.tf_project.id
  #state        = var.notebook_session_state
}

resource "oci_core_subnet" "tf_subnet" {
  cidr_block     = "10.0.1.0/24"
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.tf_vcn.id
}

resource "oci_core_vcn" "tf_vcn" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_ocid
}

