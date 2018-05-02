###
## Variables here are sourced from env, but still need to be initialized for Terraform
###

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" { default = "us-phoenix-1" }

variable "compartment_ocid" {}
variable "ssh_public_key" {}
variable "ssh_private_key" {}

variable "AD" { default = "2" }

variable "image_ocid" {
  default = " "
}
variable "nodecount" { default = "6" }

variable "MasterNodeCount" { default = "2" }

variable "blocksize_in_gbs" { default = "1024" }

variable "BastionInstanceShape" {
  default = "VM.Standard2.8"
}

variable "MasterInstanceShape" {
  default = "VM.Standard2.8"
}

variable "WorkerInstanceShape" {
  default = "BM.DenseIO2.52"
}


