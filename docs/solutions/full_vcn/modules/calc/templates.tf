### Alter only select portions of this file to meet your naming standards ###
# templates.tf - templates for using in generating names and subnet information

# Generate the netnumber for custom subnet ranges. 
# ONLY CHANGE THE LOCATION OF YOUR PYTHON 2.7.x binary! 
# This executes the call to the included Python 2.7.x script to calculate the needed 
# net numbers for the cidrsubnet call below.
data "template_file" "nnlist" {
        template = "$${nnlist}"
        vars = {
                nnlist = "${data.external.netnum.result["nnStrList"]}"
        }
}

# Generate the list of subnets based on the desired CIDR specifcation on a per subnet basis.
# DO NOT ALTER THIS SECTION.
# Takes the data from the subnet definition in the variable.tf file, gathers the net number list
# and calculates the new subnet and range. 
data "template_file" "subnet" {
	count = "${length(var.subnet["subnet_masks"])}"
	template = "$${subnet_cidr}"
	vars = {
		subnet_cidr = "${cidrsubnet(var.vcn_state["cidr_block"], element(var.subnet["subnet_masks"], count.index) - element(split("/", var.vcn_state["cidr_block"]),1), element(split(",", data.template_file.nnlist.rendered), count.index))}"
	}	
}