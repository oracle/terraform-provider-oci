variable "tenancy_ocid" {
  default = ""
}

variable "region" {
  default = ""
}

variable "compartment_ocid" { 
  default = "" 
}

variable "fingerprint" {
  default = ""
}

variable "private_key_path" {
  default = ""
}

provider "oci" {
  tenancy_ocid = var.tenancy_ocid
  fingerprint = var.fingerprint
  private_key_path = var.private_key_path
  region = var.region

}

variable defined_tag_namespace_name { default = "" }


resource "oci_ai_language_project" "test_project" {
  compartment_id = var.compartment_ocid
}

resource "oci_ai_language_model" "test_model" {
  compartment_id = var.compartment_ocid
  project_id = oci_ai_language_project.test_project.id
  description = "Creating test model"
  model_details {
    model_type = "NAMED_ENTITY_RECOGNITION"
    language_code = "en"
  }
  training_dataset {
    dataset_type = "OBJECT_STORAGE"
    location_details {
      location_type = "OBJECT_LIST"
      bucket = "TERSI-Test"
      namespace = "idngwwc5ajp5"
      object_names = ["test.jsonl"]
    }
  }
}

resource "oci_ai_language_endpoint" "test_endpoint" {
  compartment_id = var.compartment_ocid
  model_id = oci_ai_language_model.test_model.id
  inference_units = 1
}