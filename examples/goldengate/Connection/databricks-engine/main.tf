variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_id" {}
variable "region" {}

variable "password" {
  description = "Databricks personal access token."
  type        = string
  sensitive   = true
}

variable "workspace_url" {
  description = "Databricks workspace REST API URL."
  type        = string
  default     = "https://example.cloud.databricks.com/api/2.0/"
}

variable "artifact_path" {
  description = "Unity Catalog volume path used for application artifacts."
  type        = string
  default     = "/Volumes/main/default/ggs-artifacts"
}

variable "display_name" {
  type    = string
  default = "DatabricksEngine_TerraformExample"
}

variable "freeform_tags" {
  type    = map(string)
  default = { "bar-key" = "value" }
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_golden_gate_connection" "databricks_engine" {
  compartment_id      = var.compartment_id
  connection_type     = "DATABRICKS_ENGINE"
  technology_type     = "DATABRICKS_ENGINE"
  display_name        = var.display_name
  description         = "Databricks Engine connection created by Terraform"
  freeform_tags       = var.freeform_tags
  routing_method      = "SHARED_SERVICE_ENDPOINT"
  does_use_secret_ids = false

  authentication_type = "PERSONAL_ACCESS_TOKEN"
  auth_token          = var.password
  workspace_url       = var.workspace_url

  artifact_location {
    location_type             = "UC_VOLUME"
    path                      = var.artifact_path
    should_overwrite_existing = true
  }
}

data "oci_golden_gate_connection" "databricks_engine" {
  connection_id = oci_golden_gate_connection.databricks_engine.id
}

data "oci_golden_gate_connections" "databricks_engines" {
  compartment_id  = var.compartment_id
  connection_type = ["DATABRICKS_ENGINE"]
  technology_type = ["DATABRICKS_ENGINE"]

  filter {
    name   = "id"
    values = [oci_golden_gate_connection.databricks_engine.id]
  }
}

output "connection_id" {
  value = oci_golden_gate_connection.databricks_engine.id
}

output "connection_display_name" {
  value = data.oci_golden_gate_connection.databricks_engine.display_name
}

output "connection_count" {
  value = length(data.oci_golden_gate_connections.databricks_engines.connection_collection[0].items)
}
