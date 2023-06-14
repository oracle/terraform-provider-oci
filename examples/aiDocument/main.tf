provider "oci" {
}

variable "tenancy_ocid" {
  default = "ocid1.tenancy.oc1..aaaaaaaaikox5b3adi2w237m2fwomzxybp52i7byjrj5fxradayjqxum7bsq"
}

variable "region" {
  default = "us-phoenix-1"
}

variable "model_model_version" {
  default = "modelVersion"
}

variable "compartment_id" { default = "ocid1.compartment.oc1..aaaaaaaa3jewat7ub6yf257bsxvfcfz5zt46fruduji37ekbsefwmcmzvgxq" }

variable defined_tag_namespace_name { default = "" }


resource "oci_ai_document_project" "test_project" {
  compartment_id = var.compartment_id
}

resource "oci_ai_document_model" "test_model" {
  #Required
  compartment_id = var.compartment_id
  model_type = "KEY_VALUE_EXTRACTION"
  project_id = oci_ai_document_project.test_project.id

  training_dataset {
    bucket = "tf_test_bucket"
    dataset_type = "OBJECT_STORAGE"
    namespace = "axgexwaxnm7k"
    object = "tf_test_dataset_1680065500556.jsonl"
  }

  #Optional
  display_name = "test_tf_model"
  is_quick_mode = "false"
  max_training_time_in_hours = "0.5"
  model_version              = var.model_model_version
}
