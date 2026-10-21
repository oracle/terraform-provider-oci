variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_id" {}
variable "region" {}


provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

variable "deployment_id" {
  description = "OCID of an existing OCI GoldenGate deployment."
  type        = string
}

variable "certificate_key" {
  description = "Unique certificate key within the deployment."
  type        = string
  default     = "GGS23708Repro"
}

resource "oci_golden_gate_deployment_certificate" "repro" {
  deployment_id       = var.deployment_id
  key                 = var.certificate_key
  certificate_content = filebase64("${path.module}/certificate.pem")

  lifecycle {
    ignore_changes = [certificate_content]
  }
}
