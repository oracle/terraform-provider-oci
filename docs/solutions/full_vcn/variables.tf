### variables.tf - Settings for VCN creation
# Set all variables here only

# vcn_state - General settings for the VCN environment
# cidr_block: Desired address space for the VCN 
# compartment: Compartment name for the VCN
variable "vcn_state" {
	type="map"
	default = {
		cidr_block = "<ip address space>/<netmask>"
		compartment = "<desired compartment by name>"
	}
}

# subnet - Set desired subnet layout
# public: Per subnet to create, indicate whether to allow Public IP addresses or deny
#         'true' - allow, 'false' - deny
# ad: Which AD to create the subnet.  Set all subnet entries to '0' to have the template auto-select
# subnet_masks: What subnet mask to assign to the subnet.  Must be less than the CIDR block 
#               address, and the total address space must be less than the CIDR block.  See
#               the README.md for more details on how to use this.
variable "subnet" {
	type = "map"
	default = {
		public = [ false, false, false, false, false, false, true, true, true ]
		ad = [ 1, 1, 2, 2, 3, 3, 1, 2, 3 ]
		subnet_masks = [ 27, 27, 27, 27, 27, 27, 29, 30, 30 ]
	}
}

# Set module for use.  If you want the template to calculate the subnet masks evenly for the 
# number of subnets to create, use the "./modules/auto" line.  Otherwise, specify the subnet
# masks above in the "subnet" variable, and use the "./modules/calc" line. ONLY USE ONE LINE - 
# DO NOT COMMENT OUT BOTH LINES.  See the README.md for more details.
module "subnet" {
	source = "./modules/auto" 
#	source = "./modules/calc"
	subnet = "${var.subnet}"
	vcn_state = "${var.vcn_state}"
}
