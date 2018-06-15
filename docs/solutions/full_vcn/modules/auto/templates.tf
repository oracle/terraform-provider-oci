### Alter only select portions of this file to meet your naming standards ###
# templates.tf - templates for using in generating names and subnet information

# Generate the CIDR for the subnet.  DO NOT ALTER THIS SECTION.
# This generates the template defined subnet CIDR. It takes the number of subnets and 
# evenly splits the VCN CIDR range across them.  On odd numbers of subnets, it excludes the
# last possible subnet in the range.
data "template_file" "subnet" {
	count = "${length(var.subnet["public"])}"
	template = "$${subnet_cidr}"
	vars = {
		subnet_cidr = "${cidrsubnet(var.vcn_state["cidr_block"], length(var.subnet["public"]) % 2 == 0 ? length(var.subnet["public"])/2 : (floor(length(var.subnet["public"]))/2) + 1 , count.index)}"
	}
}