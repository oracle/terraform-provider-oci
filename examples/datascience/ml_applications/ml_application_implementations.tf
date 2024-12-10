resource "oci_datascience_ml_application_implementation" "test_ml_application_implementation" {
  #Required
  compartment_id    = var.compartment_id
  ml_application_id = oci_datascience_ml_application.test_ml_application.id
  name              = var.ml_application_implementation_name

  #Optional
  /*ml_application_package = {
    source_type = "local"
    path = "file://${path.root}/ml-app-package.zip"
  }*/
  ml_application_package = {
    source_type = "object_storage_download"
    uri = "https://objectstorage.us-ashburn-1.oraclecloud.com/n/ociodscdev/b/Artifact/o/windows.zip"
  }
  opc_ml_app_package_args        = var.opc_ml_app_package_args
  allowed_migration_destinations = var.ml_application_implementation_allowed_migration_destinations
  freeform_tags                  = var.ml_application_implementation_freeform_tags
}

data "oci_datascience_ml_application_implementation" "test_ml_application_implementation" {
  #Required
  ml_application_implementation_id = oci_datascience_ml_application_implementation.test_ml_application_implementation.id
}

data "oci_datascience_ml_application_implementations" "test_ml_application_implementations" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  compartment_id_in_subtree        = var.ml_application_implementation_compartment_id_in_subtree
  ml_application_id                = oci_datascience_ml_application.test_ml_application.id
  ml_application_implementation_id = oci_datascience_ml_application_implementation.test_ml_application_implementation.id
  name                             = var.ml_application_implementation_name
  state                            = var.ml_application_implementation_state
}

variable "ml_application_implementation_allowed_migration_destinations" {
  default = []
}

variable "ml_application_implementation_compartment_id_in_subtree" {
  default = false
}

variable "ml_application_implementation_defined_tags_value" {
  default = "value"
}

variable "ml_application_implementation_freeform_tags" {
  default = { "Department" = "Finance", "Purpose" : "Test" }
}

variable "ml_application_implementation_name" {
  default = "ml-app-impl-name"
}

variable "ml_application_implementation_state" {
  default = "ACTIVE"
}

variable "opc_ml_app_package_args" {
  default = {
    "bucket_namespace" : "idtlxnfdweil"
  }
}