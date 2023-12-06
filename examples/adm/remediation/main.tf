// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "devops_code_repository_ocid" {}
variable "devops_build_pipeline_ocid" {}
variable "core_subnet_ocid" {}

variable "remediation_recipe_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "remediation_recipe_defined_tags" {
  default  = { "example-tag-namespace-all.example-tag" = "value" }
}

variable "remediation_recipe_id" {
  default = "id"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_adm_knowledge_base" "example_knowledge_base" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  display_name  = "Example_Knowledge_Base"
}

resource "oci_adm_remediation_recipe" "example_remediation_recipe" {
  compartment_id = var.compartment_ocid
  display_name = "example_remediation_recipe"
  knowledge_base_id = oci_adm_knowledge_base.example_knowledge_base.id

  is_run_triggered_on_kb_change = false
  detect_configuration {
    exclusions = []
    max_permissible_cvss_v2score = 1.5
    max_permissible_cvss_v3score = 1.5
    upgrade_policy = "NEAREST"
  }
  network_configuration {
    subnet_id = var.core_subnet_ocid
  }
  scm_configuration {
    branch = "main"
    is_automerge_enabled = false
    scm_type = "OCI_CODE_REPOSITORY"
    build_file_location = "pom.xml"
    oci_code_repository_id = var.devops_code_repository_ocid
  }
  verify_configuration {
    build_service_type = "OCI_DEVOPS_BUILD"
    pipeline_id = var.devops_build_pipeline_ocid
  }
}

resource "oci_adm_remediation_run" "example_remediation_run" {
  remediation_recipe_id = oci_adm_remediation_recipe.example_remediation_recipe.id
  display_name = "example_remediation_run"
}

data "oci_adm_remediation_runs" "example_remediation_runs_data" {
  remediation_recipe_id = oci_adm_remediation_recipe.example_remediation_recipe.id
  compartment_id = var.compartment_ocid
}

data "oci_adm_remediation_run_stages" "example_run_stages" {
  remediation_run_id = oci_adm_remediation_run.example_remediation_run.id
}
