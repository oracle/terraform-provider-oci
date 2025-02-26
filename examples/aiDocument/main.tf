provider "oci" {
#   version = "6.21.0"
}

variable "tenancy_ocid" {
  default = ""
}

variable "region" {
  default = "us-phoenix-1"
}

variable "model_model_version" {
  default = "V1.0"
}

variable "compartment_id" { default = "" }

variable defined_tag_namespace_name { default = "" }


resource "oci_ai_document_project" "test_project" {
  compartment_id = var.compartment_id
}

resource "oci_ai_document_model" "test_model" {
  #Required
  compartment_id = var.compartment_id
  model_type = "KEY_VALUE_EXTRACTION"
  project_id = oci_ai_document_project.test_project.id
  inference_units = 1
  language = "ENG"

  training_dataset {
    bucket = "canary_test"
    dataset_type = "OBJECT_STORAGE"
    namespace = "axylfvgphoea"
    object = "canary-aadhar-dataset_1686632830312.jsonl"
  }

  #Optional
  display_name = "test_tf_model"
  is_quick_mode = "false"
  model_version              = var.model_model_version
}

variable "object_locations" {
  type = list(map(string))
  default = [
    {
      bucket    = "canary_test"
      namespace = "axylfvgphoea"
      object    = "dus_test.pdf"
    }
  ]
}


variable "features" {
  type = list(map(string))
  default = [
    {
      feature_type = "KEY_VALUE_EXTRACTION"
      selection_mark_detection = false
    }
  ]
}

resource "oci_ai_document_processor_job" "test_processor_job" {
  compartment_id = var.compartment_id
  display_name = "test_tf_processor_job"
  input_location {

    dynamic "object_locations" {
      for_each = var.object_locations
      content {
        bucket    = object_locations.value["bucket"]
        namespace = object_locations.value["namespace"]
        object    = object_locations.value["object"]
      }
    }

    source_type = "OBJECT_STORAGE_LOCATIONS"
  }
  output_location {
    bucket = "canary_test"
    namespace = "axylfvgphoea"
    prefix = "test"
  }
  processor_config {

    dynamic "features" {
      for_each = var.features
      content {
        feature_type    = features.value["feature_type"]
        selection_mark_detection = features.value["selection_mark_detection"]
      }
    }

    processor_type = "GENERAL"
    document_type = "INVOICE"
    language = "ENG"
  }
}