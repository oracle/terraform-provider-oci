// for provider
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "tenancy_ocid" {}
variable "region" {}

// for deployment
variable "compartment_id" {}
variable "test_subnet_id" {}
variable "password" {}

// for certificate - base64 encoded
variable "certificate_content" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_golden_gate_deployment" "test_deployment" {
  compartment_id          = var.compartment_id
  cpu_core_count          = 1
  deployment_type         = "DATABASE_ORACLE"
  display_name            = "TF_CERTIFICATE_EXAMPLE"
  is_auto_scaling_enabled = false
  license_model           = "LICENSE_INCLUDED"
  subnet_id               = var.test_subnet_id
  ogg_data {
    admin_password  = var.password
    admin_username  = "oggadmin"
    deployment_name = "tf-certificate-example"
  }
}

resource "oci_golden_gate_deployment_certificate" "certificate" {
  deployment_id = oci_golden_gate_deployment.test_deployment.id
  key = "certificate"
  certificate_content = var.certificate_content
}