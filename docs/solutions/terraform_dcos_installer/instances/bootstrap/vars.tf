#OCI variables
variable "tenancy_ocid" {}
variable "compartment_ocid" {}

variable "ssh_public_key" {}
variable "ssh_private_key" {}

variable "region" {
  default = "us-phoenix-1"
}

# Choose an Availability Domain
variable "BootstrapAD" {}

variable "shape" {
    default = "VM.Standard1.4"
}

variable "subnets" {type = "list"}
variable "image" {}
variable "dcos_cluster_name" {} 

variable "dcos_installer_url" {
  description = "Path to get DCOS"
  default     = "https://downloads.dcos.io/dcos/EarlyAccess/dcos_generate_config.sh"
}

