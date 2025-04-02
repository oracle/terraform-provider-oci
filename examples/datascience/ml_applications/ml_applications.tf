resource "oci_datascience_ml_application" "test_ml_application" {
  #Required
  compartment_id = var.compartment_id
  name = var.ml_application_name

  #Optional
  description = var.ml_application_description
  freeform_tags = var.ml_application_freeform_tags
}

data "oci_datascience_ml_application" "test_ml_application" {
  #Required
  ml_application_id = oci_datascience_ml_application.test_ml_application.id
}

data "oci_datascience_ml_applications" "test_ml_applications" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  compartment_id_in_subtree = var.ml_application_compartment_id_in_subtree
  ml_application_id = oci_datascience_ml_application.test_ml_application.id
  name = var.ml_application_name
  state = var.ml_application_state
}

variable "ml_application_compartment_id_in_subtree" {
  default = false
}

variable "ml_application_defined_tags_value" {
  default = "value"
}

variable "ml_application_description" {
  default = "Test description"
}

variable "ml_application_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "ml_application_name" {
  default = "ml-app-name_T6"
}

variable "ml_application_state" {
  default = "ACTIVE"
}