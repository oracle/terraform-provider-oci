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

# NOTE: If a different all zeros is required for a specific region, specify it here, then change
# the datasource.tf to match (specify var.region between the [] vs. "all")
variable "region_all_zeros_ocid" {
	type = "map"
	default = {
		all = "ocid1.image.oc1..aaaaaaaadevqufnklkexuu6z62f7riocqigz6zng5mxhuhghy3e6zurwct2a"
	}
}

# The images here represent the latest OL 7.x images - currently the latest 
# OL 7.5.  As new major versions are released, these should be updated.
variable "ipxe_image_ocid" {
	type = "map"
	default = {
		us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaaozjbzisykoybkppaiwviyfzusjzokq7jzwxi7nvwdiopk7ligoia"
		us-ashburn-1 = "ocid1.image.oc1.iad.aaaaaaaa6ybn2lkqp2ejhijhehf5i65spqh3igt53iyvncyjmo7uhm5235ca"
		eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaazregkysspxnktw35k4r5vzwurxk6myu44umqthjeakbkvxvxdlkq"
		uk-london-1 = "ocid1.image.oc1.uk-london-1.aaaaaaaayodsld656eh5stds5mo4hrmwuhk2ugin4eyfpgoiiskqfxll6a4a"
	}
}
