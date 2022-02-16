provider "oci" {
}

variable "tenancy_ocid" {
  default = "ocid1.tenancy.oc1..aaaaaaaa6xo4q4r2l2nvcr3sl657pwla5k3xtbk2s6vgyrvxfuh4p66frooq"
}

variable "region" {
  default = "us-phoenix-1"
}

variable "compartment_id" { default = "ocid1.compartment.oc1..aaaaaaaawzg4jwgam76fgmkq6fqbo6pmnlctp6mmt4k5qzh5xxwga3daqlbq" }

variable defined_tag_namespace_name { default = "" }


resource "oci_ai_vision_project" "test_project" {
  compartment_id = var.compartment_id
}

resource "oci_ai_vision_model" "test_model" {
  compartment_id = var.compartment_id
  is_quick_mode = "false"
  max_training_duration_in_hours = "0.01"
  model_type = "IMAGE_CLASSIFICATION"
  project_id = oci_ai_vision_project.test_project.id
  training_dataset {
    bucket = "golden_dataset"
    dataset_type = "OBJECT_STORAGE"
    namespace_name = "axhheqi2ofpb"
    object = "a_hymenoptera_v3.json"
  }
}
