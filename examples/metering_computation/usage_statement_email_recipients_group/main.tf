variable "tenancy_ocid" {
}

variable "region" {
  default = "r1"
}

variable "subscription_id" {
  default = "10153310"
}


provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = "terraform-federation-test"
  region              = var.region
}

resource "oci_metering_computation_usage_statement_email_recipients_group" "test_usage_statement_email_recipients_group" {
  subscription_id = var.subscription_id
  compartment_id  = var.tenancy_ocid
  recipients_list {
    email_id = "test@example.com"
    state = "ACTIVE"
  }
}