// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_datascience_project" "datascience_project_rd" {
  #Required
  compartment_id = "${var.compartment_ocid}"

  #Optional
  description  = "description"
  display_name = "datascienceProjectRD"
}

data "oci_datascience_projects" "datascience_projects_rd" {
  #Required
  compartment_id = "${var.compartment_ocid}"

  #Optional
  display_name = "datascienceProjectsRD"
}

resource "oci_datascience_notebook_session" "datascience_notebook_session_rd" {
  #Required
  compartment_id = "${var.compartment_ocid}"

  notebook_session_configuration_details {
    #Required
    shape     = "${var.datascience_notebook_session_shape}"
    subnet_id = "${oci_core_subnet.tf_subnet.id}"

    #Optional
    block_storage_size_in_gbs = "50"
  }

  project_id = "${oci_datascience_project.datascience_project_rd.id}"

  #Optional
  display_name = "datascienceNotebookSessionRD"
}

data "oci_datascience_notebook_sessions" "datascience_notebook_sessions_rd" {
  #Required
  compartment_id = "${var.compartment_ocid}"

  #Optional
  #created_by   = "${var.notebook_session_created_by}"
  #display_name = "${var.notebook_session_display_name}"
  #id           = "${var.notebook_session_id}"
  #project_id   = "${oci_datascience_project.tf_project.id}"
  #state        = "${var.notebook_session_state}"
}

resource "oci_datascience_model" "datascience_model_rd" {
  #Required
  artifact_content_length = "${var.datascience_model_artifact_content_length}"
  model_artifact          = "${var.datascience_model_artifact}"
  compartment_id          = "${var.compartment_ocid}"
  project_id              = "${oci_datascience_project.datascience_project_rd.id}"

  #Optional
  description  = "description2"
  display_name = "datascienceModelRD"
}

data "oci_datascience_models" "tf_models_rd" {
  #Required
  compartment_id = "${var.compartment_ocid}"

  #Optional
  #created_by   = "${var.model_created_by}"
  #display_name = "${var.model_display_name}"
  #id           = "${var.model_id}"
  #project_id   = "${oci_datascience_project.tf_project.id}"
  #state        = "${var.model_state}"
}

resource "oci_datascience_model_provenance" "datascience_model_provenance_rd" {
  #Required
  model_id = "${oci_datascience_model.datascience_model_rd.id}"

  #Optional
  git_branch      = "gitBranch"
  git_commit      = "gitCommit"
  repository_url  = "repositoryURL"
  script_dir      = "scriptDir"
  training_script = "trainingScript"
}

data "oci_datascience_model_provenance" "tf_model_provenance" {
  #Required
  model_id = "${oci_datascience_model.datascience_model_rd.id}"
}

resource "oci_core_subnet" "tf_subnet" {
  cidr_block     = "10.0.1.0/24"
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_vcn.vcn2_rd.id}"
}
