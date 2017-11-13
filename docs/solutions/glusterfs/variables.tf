// These settings can be populated here or read from your env-vars settings

// Settings for authentication
variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "region" {}
variable "private_key_path" {}
variable "private_key_password" {}

variable "compartment_ocid" {}

// The SSH public key for connecting to the compute instances
variable "ssh_public_key" {}

// The name DNS label to use for the VCN
variable "DnsLabel" {}

variable "ServerInstanceShape" {
  default = "BM.DenseIO1.36"
}

variable "ClientInstanceShape" {
  default = "VM.Standard1.2"
}

variable "ServerInstanceImage" {
  default = "CentOS-7-2017.07.17-0"
}

variable "ClientInstanceImage" {
  default = "Oracle-Linux-7.4-2017.10.25-0"
}

variable "ServerBootStrapFile" {
  default = "./userdata/bootstrap-server.sh"
}

variable "ClientBootStrapFile" {
  default = "./userdata/bootstrap-client.sh"
}
