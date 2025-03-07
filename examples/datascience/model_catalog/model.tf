variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {
  default = "ocid1.tenancy.oc1..aaaaaaaahzy3x4boh7ipxyft2rowu2xeglvanlfewudbnueugsieyuojkldq"
}

variable "model_artifact_export_artifact_source_type" {
  default = "artifactSourceType"
}

variable "model_artifact_export_namespace" {
  default = "namespace"
}

variable "model_artifact_export_source_bucket" {
  default = "sourceBucket"
}

variable "model_artifact_export_source_object_name" {
  default = "sourceObjectName"
}

variable "model_artifact_export_source_region" {
  default = "us-ashburn-1"
}

variable "artifact_content_length" {
  default = "6954"
}

variable "model_artifact" {
}

variable "content_disposition" {
}

variable "shape" {
}

variable "model_defined_tags" {
}

variable "model_freeform_tag" {
}

variable "model_state" {
}


variable "model_display_name" {
  default = "terraform-testing-model"
}

variable "model_description" {
  default = "Model for terraform testing"
}

variable "model_category" {
  default = "USER"
}

variable "model_custom_metadata_list_category" {
  default = "category"
}

variable "model_custom_metadata_list_description" {
  default = "description"
}

variable "model_custom_metadata_list_has_artifact" {
  default = false
}

variable "model_custom_metadata_list_key" {
  default = "key"
}

variable "model_custom_metadata_list_keywords" {
  default = []
}

variable "model_custom_metadata_list_value" {
  default = "value"
}

variable "model_defined_metadata_list_category" {
  default = "category"
}

variable "model_defined_metadata_list_description" {
  default = "description"
}

variable "model_defined_metadata_list_has_artifact" {
  default = false
}

variable "model_defined_metadata_list_key" {
  default = "key"
}

variable "model_defined_metadata_list_keywords" {
  default = []
}

variable "model_defined_metadata_list_value" {
  default = "value"
}

variable "model_custom_metadata_artifact_model_custom_metadatum_artifact" {
  default = "modelCustomMetadatumArtifact"
}

variable "model_custom_metadata_artifact_content_disposition" {
  default = "contentDisposition"
}

variable "model_custom_metadata_artifact_content_length" {
  default = 10
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_datascience_project" "tf_project" {
  compartment_id = var.compartment_id
}

variable "model_custom_metadata_artifact_content_range" {
  default = "range"
}

variable "model_defined_metadata_artifact_content_range" {
  default = "range"
}

variable "model_defined_metadata_artifact_model_defined_metadatum_artifact" {
  default = "modelDefinedMetadatumArtifact"
}

variable "model_defined_metadata_artifact_content_disposition" {
  default = "contentDisposition"
}

variable "model_defined_metadata_artifact_content_length" {
  default = 10
}

# A model resource configurations for creating a new model
resource "oci_datascience_model" "tf_model" {
  # Required
  artifact_content_length = var.artifact_content_length
  model_artifact          = "${path.root}/artifact.zip"
  compartment_id          = var.compartment_id
  project_id              = oci_datascience_project.tf_project.id
  #Optional
  #defined_tags  = {"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "${var.model_defined_tags_value}"}
  artifact_content_disposition = var.content_disposition
  description                  = var.model_description
  display_name                 = var.model_display_name
}

resource "oci_datascience_model" "test_model_new" {
  #Required
  compartment_id = var.compartment_id
  project_id     = oci_datascience_project.tf_project.id

  custom_metadata_list {

    #Optional
    category    = var.model_custom_metadata_list_category
    description = var.model_custom_metadata_list_description
    key         = var.model_custom_metadata_list_key
    value       = var.model_custom_metadata_list_value
    category     = var.model_custom_metadata_list_category
    description  = var.model_custom_metadata_list_description
    has_artifact = var.model_custom_metadata_list_has_artifact
    key          = var.model_custom_metadata_list_key
    keywords     = var.model_custom_metadata_list_keywords
    value        = var.model_custom_metadata_list_value
  }
  defined_metadata_list {

    #Optional
    category    = var.model_defined_metadata_list_category
    description = var.model_defined_metadata_list_description
    key         = var.model_defined_metadata_list_key
    value       = var.model_defined_metadata_list_value
    category     = var.model_defined_metadata_list_category
    description  = var.model_defined_metadata_list_description
    has_artifact = var.model_defined_metadata_list_has_artifact
    key          = var.model_defined_metadata_list_key
    keywords     = var.model_defined_metadata_list_keywords
    value        = var.model_defined_metadata_list_value
  }
  description          = var.model_description
  display_name         = var.model_display_name
}

data "oci_datascience_model_custom_metadata_artifact_contents" "test_model_custom_metadata_artifact_contents" {
  #Required
  metadatum_key_name = "metadatumkeyname"
  model_id           = oci_datascience_model.tf_model.id

  #Optional
  range = var.model_custom_metadata_artifact_content_range
}

resource "oci_datascience_model_custom_metadata_artifact" "test_model_custom_metadata_artifact" {
  #Required
  model_custom_metadatum_artifact = var.model_custom_metadata_artifact_model_custom_metadatum_artifact
  content_length                  = var.model_custom_metadata_artifact_content_length
  metadatum_key_name              = "metadatumkeyname"
  model_id                        = oci_datascience_model.tf_model.id

  #Optional
  content_disposition = var.model_custom_metadata_artifact_content_disposition
}

data "oci_datascience_model_defined_metadata_artifact_contents" "test_model_defined_metadata_artifact_contents" {
  #Required
  metadatum_key_name = "metadatumkeyname"
  model_id           = oci_datascience_model.tf_model.id

  #Optional
  range = var.model_defined_metadata_artifact_content_range
}

resource "oci_datascience_model_defined_metadata_artifact" "test_model_defined_metadata_artifact" {
  #Required
  model_defined_metadatum_artifact = var.model_defined_metadata_artifact_model_defined_metadatum_artifact
  content_length                   = var.model_defined_metadata_artifact_content_length
  metadatum_key_name               = "metadatumkeyname"
  model_id                         = oci_datascience_model.tf_model.id

  #Optional
  content_disposition = var.model_defined_metadata_artifact_content_disposition
}


