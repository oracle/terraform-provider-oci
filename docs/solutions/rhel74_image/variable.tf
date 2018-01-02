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
variable "region" {}

variable "ipxe_instance" {
	type = "map"
	default = {
		name = "ipxe-rhel74"
		hostname = "ipxe-rhel74"
		shape = "VM.Standard1.1"
	}
}

variable "ipxe_instance_image_ocid" {
	type = "map"
	default = {
		// Use the image "Oracle-Linux-7.4-2017.09.29-0"
		us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaa3g2xpzlbrrdknqcjtzv2tvxcofjc55vdcmpxdlbohmtt7encpana"
		us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaaawy2hh3nreaesyqcdp4m6csg4lwen6ya2njgiyjeu5sodiahlaxq"
		eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt1.aaaaaaaaydqkfzrcejyllfiwhcfqob2yyvkmytghwki6zcmhyciyruinokva"
	}
}

variable "region_all_zeros_ocid" {
	type = "map"
	default = {
		us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaay27pdopotkapf2ahjlsn2wxndui5hn5w37hd2wss4ses4ol5xs6a"
		us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaaqftkoa5web2r7w4ls3wekgqmqy5f7untloetfiozyqbv2ql6qidq"
		eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaah4rggbyglst25peqd7vnyjzl6n5lwogiyllb6jaircakom46nswq"
	}
}