variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_ocid" {}
variable "data_safe_target_ocid" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

variable "relation_type" {
  default = "APP_DEFINED"
}

variable "is_sensitive" {
  default = false
}

variable "sensitive_data_model_display_name" {
  default = "displayName"
}
variable "sensitive_data_model_schemas_for_discovery" {
  default = []
}

variable "sensitive_data_model_sensitive_type_ids_for_discovery" {
  default = []
}

resource "oci_data_safe_sensitive_data_model" "test_sensitive_data_model" {
  #Required
  compartment_id = var.compartment_ocid
  target_id      = var.data_safe_target_ocid

  #Optional
  display_name                              = var.sensitive_data_model_display_name
  schemas_for_discovery                     = var.sensitive_data_model_schemas_for_discovery
  sensitive_type_ids_for_discovery          = var.sensitive_data_model_sensitive_type_ids_for_discovery
}

resource "oci_data_safe_sensitive_data_model_referential_relation" "test_sensitive_data_model_referential_relation" {
  #Required
  sensitive_data_model_id = oci_data_safe_sensitive_data_model.test_sensitive_data_model.id
  relation_type = var.relation_type
  parent {
    app_name = "HR"
    column_group = ["EMAIL"]
    object = "EMPLOYEES"
    object_type = "TABLE"
    schema_name = "HR"
  }
  child {
    app_name = "HR"
    column_group = ["FIRST_NAME"]
    object = "EMPLOYEES"
    object_type = "TABLE"
    schema_name = "HR"
  }
  is_sensitive = var.is_sensitive
}

data "oci_data_safe_sensitive_data_model_referential_relations" "test_sensitive_data_model_referential_relations" {
  #Required
  sensitive_data_model_id = oci_data_safe_sensitive_data_model.test_sensitive_data_model.id

  #Optional
  relation_type = [oci_data_safe_sensitive_data_model_referential_relation.test_sensitive_data_model_referential_relation.relation_type]
}