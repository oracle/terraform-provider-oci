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

###
## CONFIGURE BELOW VARIABLES TO CUSTOMIZE DEPLOYMENT
###

## Specify number of Worker Hosts here
variable "nodecount" { default = "5" }

## Specify number of Master Hosts here
variable "MasterNodeCount" { default = "2" }

## Specify the size of each Block Volume attached to Worker Hosts
variable "blocksize_in_gbs" { default = "1024" }

## Set the shape to be used for Bastion Host
variable "BastionInstanceShape" {
  default = "VM.Standard2.8"
}

## Set the shape to be used for Master Hosts
variable "MasterInstanceShape" {
  default = "VM.Standard2.8"
}

## Set the shape to be used for Worker Hosts
variable "WorkerInstanceShape" {
  default = "BM.DenseIO2.52"
}

###
## End Configuration Customization
###

variable "3TB" {
  default = "3145728"
}
variable "2TB" { 
  default = "2097152"
}
variable "1TB" {
  default = "1024000"
}
variable "500GB" {
  default = "512000"
}
variable "256GB" {
  default = "262144"
}
variable "50GB" {
  default = "51200"
}

