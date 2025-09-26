variable "compartment_ocid" {}

variable "region" {}

variable config_file_profile {}

variable "imported_model_capabilities" {
  default = ["TEXT_TO_TEXT"]
}

variable "imported_model_capability" {
  default = ["TEXT_TO_TEXT"]
}

variable "hf_access_token" {
  type = string
}

variable "imported_model_data_source_source_type" {
  default = "HUGGING_FACE_MODEL"
}

variable "imported_model_description" {
  default = "description"
}

variable "imported_model_display_name" {
  default = "displayName"
}

variable "imported_model_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "imported_model_model_id" {
  default = "google/gemma-3-12b-it"
}

variable "imported_model_vendor" {
  default = "vendor"
}

variable "imported_model_version" {
  default = "4.0"
}


