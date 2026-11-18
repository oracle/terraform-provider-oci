variable "compartment_ocid" {}
variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "stack_display_name" {
  default = "example-rm-stack"
}

variable "stack_description" {
  default = "Example Resource Manager stack managed by Terraform"
}

variable "stack_zip_path" {
  description = "Absolute or module-relative path to the ZIP archive containing the Terraform configuration for the stack."
}

variable "stack_working_directory" {
  description = "Optional directory inside the ZIP archive from which Terraform runs."
  default     = "env/dev"
}
