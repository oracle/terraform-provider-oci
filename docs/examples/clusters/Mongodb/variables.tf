variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "compartment_ocid" {}
variable "ssh_public_key" {}
variable "ssh_private_key" {}


variable "BastionShape" {
    default = "VM.Standard1.1"
}

variable "MongoDBShape" {
    default = "BM.DenseIO1.36"
}

variable "InstanceImageDisplayName" {
    default = "Oracle-Linux-7.4-2017.10.25-0"
}

variable "VPC-CIDR" {
    default = "10.0.0.0/26"
}

variable "PubSubnetAD1CIDR" {
    default = "10.0.0.0/28"
}

variable "PrivSubnetAD1CIDR" {
    default = "10.0.0.16/28"
}

variable "PrivSubnetAD2CIDR" {
    default = "10.0.0.32/28"
}

variable "BastSubnetAD1CIDR" {
    default = "10.0.0.48/28"
}

variable "BastionBootStrap" {
    default = "./userdata/bastion"
}

variable "MongoDBBootStrap" {
    default = "./userdata/MongoDB"
}

