variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_id" {}
variable "region" {}

variable "api_key" {
  description = "API key used to authenticate with the AI provider."
  type        = string
  sensitive   = true
  default     = "1234"
}

variable "base_url" {
  description = "Base URL of the AI provider API."
  type        = string
  default     = "https://api.openai.com/v1"
}

variable "display_name" {
  type    = string
  default = "AiProvider_TerraformExample"
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

resource "oci_golden_gate_connection" "ai_provider" {
  compartment_id      = var.compartment_id
  connection_type     = "AI_PROVIDER"
  technology_type     = "OPEN_AI"
  display_name        = var.display_name
  description         = "AI Provider connection created by Terraform"
  freeform_tags       = var.freeform_tags
  routing_method      = "SHARED_SERVICE_ENDPOINT"
  does_use_secret_ids = false

  auth_details {
    auth_type = "API_KEY"
    api_key   = var.api_key
    base_url  = var.base_url
  }
}

data "oci_golden_gate_ai_providers" "available" {
  compartment_id = var.compartment_id
}

data "oci_golden_gate_connection" "ai_provider" {
  connection_id = oci_golden_gate_connection.ai_provider.id
}

data "oci_golden_gate_connections" "ai_providers" {
  compartment_id  = var.compartment_id
  connection_type = ["AI_PROVIDER"]
  technology_type = ["OPEN_AI"]

  filter {
    name   = "id"
    values = [oci_golden_gate_connection.ai_provider.id]
  }
}

output "connection_id" {
  value = oci_golden_gate_connection.ai_provider.id
}

output "connection_display_name" {
  value = data.oci_golden_gate_connection.ai_provider.display_name
}

output "connection_count" {
  value = length(data.oci_golden_gate_connections.ai_providers.connection_collection[0].items)
}

output "available_ai_provider_count" {
  value = length(data.oci_golden_gate_ai_providers.available.ai_provider_collection[0].items)
}
