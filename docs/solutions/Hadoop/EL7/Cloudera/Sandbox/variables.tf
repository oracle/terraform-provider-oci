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

variable "boot_volume_size" { default = "256" }

variable "image_ocid" { }

variable "blocksize_in_gbs" { default = "1024" }
