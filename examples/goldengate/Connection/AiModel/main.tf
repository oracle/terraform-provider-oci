variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "deployment_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_id" {}
variable "region" {}

variable "description" {
  default = "Created as example for AI_MODEL GoldenGate connection"
}
variable "freeform_tags" {
  default = { "bar-key" = "value" }
}
variable "connection_type" {
  default = "AI_MODEL"
}
variable "display_name" {
  default = "AiModel_TerraformTest"
}
variable "technology_type" {
  default = "AI_MODEL"
}
variable "provider_type" {
  default = "OPEN_AI"
}
variable "model_key" {
  default = "text-embedding-3-small"
}
variable "max_input_chars" {
  default = 20000
}
variable "does_use_secret_ids" {
  default = false
}
variable "base_url" {
  default = "https://api.openai.com/v1/embeddings"
}
variable "auth_type" {
  default = "API_KEY"
}
variable "api_key" {
  default = "1234"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_golden_gate_connection" "test_connection" {
  compartment_id       = var.compartment_id
  connection_type      = var.connection_type
  technology_type      = var.technology_type
  display_name         = var.display_name
  description          = var.description
  freeform_tags        = var.freeform_tags
  provider_type        = var.provider_type
  model_key            = var.model_key
  max_input_chars      = var.max_input_chars
  does_use_secret_ids  = var.does_use_secret_ids

  auth_details {
    auth_type = var.auth_type
    base_url  = var.base_url
    api_key   = var.api_key
  }
}

resource "oci_golden_gate_connection_assignment" "test_connection_assignment" {
  connection_id     = oci_golden_gate_connection.test_connection.id
  deployment_id     = var.deployment_ocid
}

data "oci_golden_gate_ai_providers" "test_ai_providers" {
  compartment_id = var.compartment_id
}

data "oci_golden_gate_connection" "fetched_connection" {
  connection_id = oci_golden_gate_connection.test_connection.id
}

data "oci_golden_gate_connections" "connections_by_type" {
  compartment_id   = var.compartment_id
  connection_type  = [var.connection_type]
  technology_type  = [var.technology_type]
}

data "oci_golden_gate_connections" "connections_not_oracle" {
  compartment_id                 = var.compartment_id
  connection_type_not_equal_to   = ["ORACLE"]
  technology_type                = [var.technology_type]
}

data "oci_golden_gate_connection_assignments" "assignments_by_type" {
  compartment_id   = var.compartment_id
  deployment_id    = var.deployment_ocid
  connection_type  = [var.connection_type]
}

data "oci_golden_gate_connection_assignments" "assignments_not_oracle" {
  compartment_id                = var.compartment_id
  deployment_id                 = var.deployment_ocid
  connection_type_not_equal_to  = ["ORACLE"]
}

output "connection_id" {
  value = oci_golden_gate_connection.test_connection.id
}

output "connection_display_name" {
  value = data.oci_golden_gate_connection.fetched_connection.display_name
}

output "connection_provider_type" {
  value = data.oci_golden_gate_connection.fetched_connection.provider_type
}

output "connections_by_type_count" {
  value = length(data.oci_golden_gate_connections.connections_by_type.connection_collection[0].items)
}

output "connections_not_oracle_count" {
  value = length(data.oci_golden_gate_connections.connections_not_oracle.connection_collection[0].items)
}

output "assignments_by_type_count" {
  value = length(data.oci_golden_gate_connection_assignments.assignments_by_type.connection_assignment_collection[0].items)
}

output "assignments_not_oracle_count" {
  value = length(data.oci_golden_gate_connection_assignments.assignments_not_oracle.connection_assignment_collection[0].items)
}

output "ai_provider_count" {
  value = length(data.oci_golden_gate_ai_providers.test_ai_providers.ai_provider_collection[0].items)
}
