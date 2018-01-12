# Set your information here:

# iso_url - specify the URL for the RHEL ISO.  This can be any publically accessible URL.
# The URL must contain the name of the ISO image with an '.iso' extension.

variable "iso_url" {
	default = "<URL of ISO>"
}

# RHEL account variables:
# user_name - The user name of the account that holds the subscription
# password - password for said user name

variable "rhel_account" {
	type = "map"
	default = {
		user_name = "<RHEL subscription username>"
		password = "<RHEL subscription password>"
	}
}

# Build environment variables:
# compartment - your compartment name
# ad - which Availability Domain to use. Format in either AD-x or ad-x where 'x' is the AD number
# vcn - the display name of the vcn to use
# subnet - display name of the subnet to use

variable "build_env" {
	type = "map"
	default = {
		compartment = "<Compartment display name>"
		ad = "<AD specified as ad-x or AD-x>"
		vcn = "<VCN display name>"
		subnet = "<Subnet display name>"
	}
}
