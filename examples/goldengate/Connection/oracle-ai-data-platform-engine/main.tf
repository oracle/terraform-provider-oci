variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_id" {}
variable "region" {}

variable "workspace_url" {
  description = "Oracle AI Data Platform workspace URL."
  type        = string
  default     = "https://ai-data-platform.example.com/"
}

variable "display_name" {
  type    = string
  default = "OracleAiDataPlatformEngine_TerraformExample"
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

resource "oci_golden_gate_connection" "oracle_ai_data_platform_engine" {
  compartment_id  = var.compartment_id
  connection_type = "ORACLE_AI_DATA_PLATFORM_ENGINE"
  technology_type = "ORACLE_AI_DATA_PLATFORM_ENGINE"
  display_name    = var.display_name
  description     = "Oracle AI Data Platform Engine connection created by Terraform"
  freeform_tags   = var.freeform_tags
  routing_method  = "SHARED_SERVICE_ENDPOINT"
  workspace_url   = var.workspace_url
}

data "oci_golden_gate_connection" "oracle_ai_data_platform_engine" {
  connection_id = oci_golden_gate_connection.oracle_ai_data_platform_engine.id
}

data "oci_golden_gate_connections" "oracle_ai_data_platform_engines" {
  compartment_id  = var.compartment_id
  connection_type = ["ORACLE_AI_DATA_PLATFORM_ENGINE"]
  technology_type = ["ORACLE_AI_DATA_PLATFORM_ENGINE"]

  filter {
    name   = "id"
    values = [oci_golden_gate_connection.oracle_ai_data_platform_engine.id]
  }
}

output "connection_id" {
  value = oci_golden_gate_connection.oracle_ai_data_platform_engine.id
}

output "connection_display_name" {
  value = data.oci_golden_gate_connection.oracle_ai_data_platform_engine.display_name
}

output "connection_count" {
  value = length(data.oci_golden_gate_connections.oracle_ai_data_platform_engines.connection_collection[0].items)
}
