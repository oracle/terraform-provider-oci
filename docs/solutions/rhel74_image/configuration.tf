# Set your information here:

variable "iso_location" {
	type = "map"
	default = {
		bucket_name = "<your bucket name>"
		iso_name = "<RHEL iso file name>"
	}
}

variable "rhel_account" {
	type = "map"
	default = {
		user_name = "<RHEL Subscription Username>"
		password = "<RHEL Subscription Password"
	}
}

variable "build_env" {
	type = "map"
	default = {
		compartment = "<Compartment display name>"
		ad = "<The number of the availability domain to use - 1, 2, or 3>"
		vcn = "<VCN name>"
		subnet = "<subnet name in compartment in ad in vcn>"
	}
}
