#OCI variables
variable "tenancy_ocid" {}
variable "compartment_ocid" {}

variable "ssh_public_key" {}
variable "ssh_private_key" {}
variable "display_name_prefix" {}

variable "region" {
  default = "us-phoenix-1"
}

# Choose an Availability Domain
variable "availability_domain" {}
variable "subnet_id" {}
variable "image" {}

variable "shape" {
    default = "VM.Standard1.4"
}

variable "dcos_cluster_name" {
  description = "Name of your cluster. Alpha-numeric and hyphens only, please."
  default     = "oci-dcos"
}

variable "count" {
  description = "Number of master nodes. 1, 3, or 5."
  default     = "1"
}

