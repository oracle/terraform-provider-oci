# DO NOT ALTER THIS FILE

provider "oci" {
  tenancy_ocid         = "${var.tenancy_ocid}"
  user_ocid            = "${var.user_ocid}"
  fingerprint          = "${var.fingerprint}"
  private_key_path     = "${var.private_key_path}"
  private_key_password = "${var.private_key_password}"
  region               = "${var.region}"
}

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "private_key_password" {}
variable "ssh_public_key" {}
variable "profile" {}
variable "region" {}

variable "ipxe_instance" {
	type = "map"
	default = {
		name = "ipxe-rhel74"
		hostname = "ipxe-rhel74"
		image = "Oracle-Linux-7.4-2017.09.29-0"
	}
}