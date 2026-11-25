terraform {
  required_version = ">= 1.3.0"

  required_providers {
    oci = {
      source = "oracle/oci"
    }
  }
}

provider "oci" {
  auth                = try(var.auth, "SecurityToken")
  config_file_profile = try(var.config_file_profile, "DEFAULT")
  region              = var.region
}
