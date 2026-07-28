variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "deployment_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_id" {}
variable "region" {}
variable "test_subnet_id" {}

variable "connection_type" {
  default = "GOLDENGATE"
}
variable "display_name" {
  default = "RcT_TerraformTest"
}
variable "technology_type" {
  default = "GOLDENGATE"
}
variable "host"{
  default = "10.0.0.129"
}
variable "port" {
  default = "14"
}
variable "is_lock_override" {
  default = false
}
provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_golden_gate_connection_assignment" "test_connection_assignment"{
  #Required
  connection_id = oci_golden_gate_connection.test_connection.id
  deployment_id = var.deployment_ocid
  is_lock_override = var.is_lock_override
  lifecycle {
    ignore_changes = [is_lock_override]
  }
}

resource "oci_golden_gate_connection" "test_connection"{
  #Required
  compartment_id = var.compartment_id
  connection_type = var.connection_type
  display_name = var.display_name
  technology_type = var.technology_type
  #Optional

  host = var.host
  port = var.port

}

data "oci_golden_gate_connection" "fetched_connection" {
  connection_id = oci_golden_gate_connection.test_connection.id
}

data "oci_golden_gate_connections" "fetched_connections" {
  compartment_id = var.compartment_id

  filter {
    name   = "id"
    values = [oci_golden_gate_connection.test_connection.id]
  }
}

data "oci_golden_gate_connection_assignment" "fetched_connection_assignment" {
  connection_assignment_id = oci_golden_gate_connection_assignment.test_connection_assignment.id
}

data "oci_golden_gate_connection_assignments" "fetched_connection_assignments" {
  compartment_id = var.compartment_id
  deployment_id  = var.deployment_ocid

  filter {
    name   = "id"
    values = [oci_golden_gate_connection_assignment.test_connection_assignment.id]
  }
}

output "connection_display_name" {
  value = data.oci_golden_gate_connection.fetched_connection.display_name
}

output "connections_count" {
  value = length(data.oci_golden_gate_connections.fetched_connections.connection_collection[0].items)
}

output "connection_assignment_id" {
  value = data.oci_golden_gate_connection_assignment.fetched_connection_assignment.id
}

output "connection_assignments_count" {
  value = length(data.oci_golden_gate_connection_assignments.fetched_connection_assignments.connection_assignment_collection[0].items)
}
