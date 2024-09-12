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

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
  version          = "6.0.0"
}

resource "oci_datascience_project" "tf_project" {
  compartment_id = var.compartment_id
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
