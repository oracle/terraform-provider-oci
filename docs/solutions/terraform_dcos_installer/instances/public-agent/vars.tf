#OCI variables
variable "tenancy_ocid" {}
variable "compartment_ocid" {}
variable "dcos_bootstrap_instance_id" {}

variable "ssh_public_key" {}
variable "ssh_private_key" {}
variable "display_name_prefix" {}

variable "region" {
  default = "us-phoenix-1"
}

# Choose an Availability Domain
variable "availability_domain" {}
variable "subnet_id" {}
variable "shape" {}
variable "image" {}
variable "count" {}
variable "dcos_cluster_name" {}


