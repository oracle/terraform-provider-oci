### Alter this file only where indicated ###
# resource.tf - create the VCN resources

# Create the VCN.  Note a couple of items:
# 1. The display name is generated using a defined naming standard in the templates.tf
#    Change this if desired to meet your own naming standard.
# 2. The dns_label is also generate using template.  Alter to meet your naming standard.
resource "oci_core_virtual_network" "vcn" {
	cidr_block = "${var.vcn_state["cidr_block"]}"
	compartment_id = "${data.oci_identity_compartments.compartment.compartments.0.id}"
	display_name = "${data.template_file.vcn_name.rendered}"
	dns_label = "${replace(data.template_file.vcn_name.rendered,"-","")}"
}

# Create an Internet Gateway.  Used for the public subnets.  Also used as a placeholder
# for private subnets, but is non-functional on those subnets.
resource "oci_core_internet_gateway" "vcn_ig" {
	compartment_id = "${data.oci_identity_compartments.compartment.compartments.0.id}"
	display_name = "${format("%s-ig", data.template_file.vcn_name.rendered)}"
	vcn_id = "${oci_core_virtual_network.vcn.id}"
}

# Create a Global security list.  This security list holds rules that apply across the
# VCN. If you have specific rules you would like 
# apply on a per-subnet basis, other than the generic one here, create an entry.  
# Understand that the rules will apply to all subnets you create here.
resource "oci_core_security_list" "global_seclist" {
	compartment_id = "${data.oci_identity_compartments.compartment.compartments.0.id}"
	vcn_id = "${oci_core_virtual_network.vcn.id}"
	display_name = "${format("%s-global-sl", data.template_file.vcn_name.rendered)}"
	egress_security_rules {
        	protocol = "all"
        	destination = "0.0.0.0/0"
    }

    	ingress_security_rules {
        	protocol = "1"
        	source = "0.0.0.0/0"

        	icmp_options {
            	"type" = 3
            	"code" = 4
        	}
    	}
}

# Create a per-subnet security list.  These security lists are applied on a per subnet
# basis and are used for subnet specific rules.  If you have specific rules you would like 
# apply on a per-subnet basis, other than the generic one here, create an entry.  Be sure
# that the rule will apply to each subnet you create here.
resource "oci_core_security_list" "subnet_seclist" {
	count = "${length(var.subnet["public"])}"
	compartment_id = "${data.oci_identity_compartments.compartment.compartments.0.id}"
	vcn_id = "${oci_core_virtual_network.vcn.id}"
	display_name = "${format("%s-sl1", data.template_file.sn_name.*.rendered[count.index])}"

	ingress_security_rules {
		protocol = "6"
		source = "${element(var.subnet["public"], count.index) ? "0.0.0.0/0" : var.vcn_state["cidr_block"]}"
		stateless = false
		tcp_options {
			"min" = 22
			"max" = 22
		}
	}
}

# Route table for private subnets.  The private subnets (non-public IP ones) have a route table
# separate from the public ones to allow for Proxy/NAT routing.  The IG is listed here as a 
# placeholder to allow the creation of the route table, but is non-functional in this case.  If you # have pre-defined routing instructions insert them here.
resource "oci_core_route_table" "private_rt" {
	compartment_id = "${data.oci_identity_compartments.compartment.compartments.0.id}"
	display_name = "Internal"
	route_rules {
		cidr_block = "0.0.0.0/0"
		network_entity_id = "${oci_core_internet_gateway.vcn_ig.id}"
	}
	vcn_id = "${oci_core_virtual_network.vcn.id}"
}

# Route table for public subnets.  All public subnets use this routing table.  Generally holds
# the IG only.
resource "oci_core_route_table" "public_rt" {
	compartment_id = "${data.oci_identity_compartments.compartment.compartments.0.id}"
	display_name = "External"
	route_rules {
		cidr_block = "0.0.0.0/0"
		network_entity_id = "${oci_core_internet_gateway.vcn_ig.id}"
	}
	vcn_id = "${oci_core_virtual_network.vcn.id}"
}

# Create each subnet.  Be careful in altering this section.
# The subnet is created using the above resources.  Things like the display_name and dns_label
# are generated using templates which set the naming standards.  Change the template to make 
# changes.  The dhcp_options_id uses the default, but a resource could be created to make a custom
# one if so desired using the design pattern for either route tables or security lists.
resource "oci_core_subnet" "subnets" {
	count = "${length(var.subnet["public"])}"
	compartment_id = "${data.oci_identity_compartments.compartment.compartments.0.id}"
	vcn_id = "${oci_core_virtual_network.vcn.id}"
	display_name = "${data.template_file.sn_name.*.rendered[count.index]}"
	availability_domain = "${lookup(data.oci_identity_availability_domains.ad.availability_domains[data.template_file.ad_number.*.rendered[count.index] - 1],"name")}"
	security_list_ids = ["${oci_core_security_list.global_seclist.id}", "${oci_core_security_list.subnet_seclist.*.id[count.index]}"]
	route_table_id = "${element(var.subnet["public"], count.index) ? oci_core_route_table.public_rt.id : oci_core_route_table.private_rt.id}"
	dhcp_options_id = "${oci_core_virtual_network.vcn.default_dhcp_options_id}"
	cidr_block = "${element(split(",", module.subnet.cidr_list), count.index)}"
	dns_label = "${format("ad%ssn%d", data.template_file.ad_number.*.rendered[count.index], (count.index + 1))}"
	prohibit_public_ip_on_vnic = "${!(element(var.subnet["public"], count.index))}"
}
