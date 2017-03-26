resource "baremetal_core_virtual_network" "MongoDB" {
  cidr_block = "${var.VPC-CIDR}"
  compartment_id = "${var.compartment_ocid}"
 	 display_name = "MongoDB"
}

resource "baremetal_core_internet_gateway" "MongoDB" {
    compartment_id = "${var.compartment_ocid}"
    display_name = "MongoIG"
    vcn_id = "${baremetal_core_virtual_network.MongoDB.id}"
}

resource "baremetal_core_route_table" "MongoDB" {
    compartment_id = "${var.compartment_ocid}"
    vcn_id = "${baremetal_core_virtual_network.MongoDB.id}"
    display_name = "MongoDB"
    route_rules {
        cidr_block = "0.0.0.0/0"
        network_entity_id = "${baremetal_core_internet_gateway.MongoDB.id}"
    }
}

resource "baremetal_core_security_list" "PublicSubnet" {
    compartment_id = "${var.compartment_ocid}"
    display_name = "Public"
    vcn_id = "${baremetal_core_virtual_network.MongoDB.id}"
    egress_security_rules = [{
	tcp_options {
	     "max" = 27017
	     "min" = 27017
	} 
        destination = "${var.VPC-CIDR}"
        protocol = "6"
    }]
    ingress_security_rules = [{
        tcp_options {
            "max" = 80
            "min" = 80
        }
        protocol = "6"
        source = "0.0.0.0/0"
    },
	{
        tcp_options {
            "max" = 443
            "min" = 443
        }
        protocol = "6"
        source = "0.0.0.0/0"
    }]
}

resource "baremetal_core_security_list" "PrivateSubnets" {
    compartment_id = "${var.compartment_ocid}"
    display_name = "Private"
    vcn_id = "${baremetal_core_virtual_network.MongoDB.id}"
    egress_security_rules = [{
	tcp_options {
		"max" = 27017
		"min" = 27017
		}
	protocol = "6"
	destination = "${var.VPC-CIDR}"
    }]
    ingress_security_rules = [{
        tcp_options {
                "max" = 27017
                "min" = 27017
                }
	protocol = "6"
        source = "${var.VPC-CIDR}"
    },
	{
        tcp_options {
                "max" = 22
                "min" = 22
                }
        protocol = "6"
        source = "${var.VPC-CIDR}"
     }]
}

resource "baremetal_core_security_list" "BastionSubnet" {
    compartment_id = "${var.compartment_ocid}"
    display_name = "Bastion"
    vcn_id = "${baremetal_core_virtual_network.MongoDB.id}"
    egress_security_rules = [{
	protocol = "6"
        destination = "0.0.0.0/0"
    }]
    ingress_security_rules = [{
        tcp_options {
            "max" = 22
            "min" = 22
        }
        protocol = "6"
        source = "0.0.0.0/0"
    }]
}

resource "baremetal_core_subnet" "PublicSubnetAD1" {
  availability_domain = "${lookup(data.baremetal_identity_availability_domains.ADs.availability_domains[0],"name")}" 
  cidr_block = "${var.PubSubnetAD1CIDR}"
  display_name = "PublicSubnetAD1"
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${baremetal_core_virtual_network.MongoDB.id}"
  route_table_id = "${baremetal_core_route_table.MongoDB.id}"
  security_list_ids = ["${baremetal_core_security_list.PublicSubnet.id}"]
}

resource "baremetal_core_subnet" "PrivSubnetAD1" {
  availability_domain = "${lookup(data.baremetal_identity_availability_domains.ADs.availability_domains[0],"name")}"
  cidr_block = "${var.PrivSubnetAD1CIDR}"
  display_name = "PrivateSubnetAD1"
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${baremetal_core_virtual_network.MongoDB.id}"
  route_table_id = "${baremetal_core_route_table.MongoDB.id}"
  security_list_ids = ["${baremetal_core_security_list.PrivateSubnets.id}"]
}

resource "baremetal_core_subnet" "PrivSubnetAD2" {
  availability_domain = "${lookup(data.baremetal_identity_availability_domains.ADs.availability_domains[1],"name")}"
  cidr_block = "${var.PrivSubnetAD2CIDR}"
  display_name = "PrivateSubnetAD2"
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${baremetal_core_virtual_network.MongoDB.id}"
  route_table_id = "${baremetal_core_route_table.MongoDB.id}"
  security_list_ids = ["${baremetal_core_security_list.PrivateSubnets.id}"]
}

resource "baremetal_core_subnet" "BastionSubnetAD1" {
  availability_domain = "${lookup(data.baremetal_identity_availability_domains.ADs.availability_domains[0],"name")}"
  cidr_block = "${var.BastSubnetAD1CIDR}"
  display_name = "BastionSubnetAD1"
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${baremetal_core_virtual_network.MongoDB.id}"
  route_table_id = "${baremetal_core_route_table.MongoDB.id}"
  security_list_ids = ["${baremetal_core_security_list.BastionSubnet.id}"]
}

