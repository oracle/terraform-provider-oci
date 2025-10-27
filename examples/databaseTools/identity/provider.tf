provider "oci" {
  region           = var.region
  auth = "SecurityToken"
  config_file_profile = "terraform-federation-test"
}