### Alter only select portions of this file to meet your naming standards ###
# templates.tf - templates for using in generating names and subnet information

# Generate the VCN name.  Set your own VCN naming standard here as desired.
data "template_file" "vcn_name" {
	template = "$${vcn_name}"
	vars = {
		vcn_name = "${format("%s-vcn%d",var.vcn_state["compartment"], (length(data.oci_core_virtual_networks.vcn.virtual_networks) + 1))}"
	}
}

# Subnet name.  Generate the display name for each subnet.  Based on the VCN name
# by default.
data "template_file" "sn_name" {
	count = "${length(var.subnet["public"])}"
	template = "$${sn_name}"
	vars = {
		sn_name = "${format("%s-ad%s-sn%d", data.template_file.vcn_name.rendered, data.template_file.ad_number.*.rendered[count.index], (count.index + 1))}"
	}
}

# Generate the AD name to place the subnet in.  This either takes the AD number and finds the 
# official AD name for the subnet, or calculates it based on the relative position of the subnet
# in the list of ADs.  This works by simply calculating the position of the entry relative to
# the number of entries.  See the README.md for details.  DO NOT ALTER.
data "template_file" "ad_name" {
	count = "${length(var.subnet["public"])}"
	template = "$${ad}"
	vars = {
		ad = "${format("ad%d", (element(var.subnet["ad"], count.index) != 0 ? element(var.subnet["ad"], count.index) : (count.index + 1) % 3 == 0 ? 3 : (count.index + 1) % 3))}"
	}
}

# Calculate the AD number for use in other resources.  DO NOT ALTER.
data "template_file" "ad_number" {
	count = "${length(var.subnet["public"])}"
	template = "$${ad}"
	vars = {
		ad = "${(element(var.subnet["ad"], count.index) != 0 ? element(var.subnet["ad"], count.index) : (count.index + 1) % 3 == 0 ? 3 : (count.index + 1) % 3)}"
	}
}