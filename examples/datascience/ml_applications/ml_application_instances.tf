resource "oci_datascience_ml_application_instance" "test_ml_application_instance" {
  #Required
  compartment_id = var.compartment_id
  ml_application_id = oci_datascience_ml_application.test_ml_application.id

  #Optional
  auth_configuration {
    #Required
    type = var.ml_application_instance_auth_configuration_type
  }
  dynamic "configuration" {
    for_each = var.configuration_map
    content {
      key = configuration.key
      value = configuration.value
    }
  }
  display_name = var.ml_application_instance_display_name
  freeform_tags = var.ml_application_instance_freeform_tags
  is_enabled = var.ml_application_instance_is_enabled
  ml_application_implementation_id = oci_datascience_ml_application_implementation.test_ml_application_implementation.id
}

data "oci_datascience_ml_application_instance" "test_ml_application_instance" {
  #Required
  ml_application_instance_id = oci_datascience_ml_application_instance.test_ml_application_instance.id
}

data "oci_datascience_ml_application_instances" "test_ml_application_instances" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = var.ml_application_instance_display_name
  ml_application_id = oci_datascience_ml_application.test_ml_application.id
  state = var.ml_application_instance_state
}

variable "ml_application_instance_auth_configuration_access_token" {
  default = "accessToken"
}

variable "ml_application_instance_auth_configuration_audience" {
  default = "audience"
}

variable "ml_application_instance_auth_configuration_role_name" {
  default = "roleName"
}

variable "ml_application_instance_auth_configuration_scope" {
  default = "scope"
}

variable "ml_application_instance_auth_configuration_type" {
  default = "IAM"
}

variable "ml_application_instance_defined_tags_value" {
  default = "value"
}

variable "ml_application_instance_display_name" {
  default = "ml-app-instance-name"
}

variable "ml_application_instance_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "ml_application_instance_is_enabled" {
  default = true
}

variable "ml_application_instance_state" {
  default = "ACTIVE"
}

variable "configuration_map" {
  default = {
    ingestion_bucket_name = "ingestion_bucket_name",
  }
}