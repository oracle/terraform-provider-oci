variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "compartment_ocid" {}
variable "ssh_public_key" {}
variable "ssh_private_key" {}

variable "VPC-CIDR" {
  default = "10.0.0.0/16"
}
variable "BlockVolumeSize" {
  default = "50" // size in GBs, min: 50, max 16384
}

variable "InstanceImageDisplayName" {
  default = "Oracle-Linux-7.4-2017.10.25-0"
}

variable "InstanceShape" {
  default = "VM.Standard1.2"
}
