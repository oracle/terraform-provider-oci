variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

variable "compartment_ocid" {}
variable "ssh_public_key" {}
variable "ssh_private_key" {}

# Choose an Availability Domain
variable "AD" {
    default = "1"
}

variable "InstanceShape" {
    default = "VM.Standard1.2"
}

variable "InstanceImageDisplayName" {
    default = "Oracle-Linux-7.4-2017.10.25-0"
}

variable "DBSize" {
    default = "50" // size in GBs
}

variable "BootStrapFile" {
    default = "./userdata/bootstrap"
}
