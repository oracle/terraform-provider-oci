# Set your information here:

variable "iso_location" {
	type = "map"
	default = {
		bucket_name = "c2-isos"
		iso_name = "rhel-server-7.4-x86_64-dvd.iso"
	}
}

variable "rhel_account" {
	type = "map"
	default = {
		user_name = "nelsonse1"
		password = "sp1Tfir3"
	}
}

variable "build_env" {
	type = "map"
	default = {
		compartment = "c2"
		ad = "ad-1"
		vcn = "c2-vcn1"
		subnet = "c2-vcn1-ad1-sn1"
	}
}