# Creating a Virtual Cloud Network (VCN)
resource "oci_core_virtual_network" "MyVcn" {
  cidr_block = "${var.VPC-CIDR}"
  dns_label = "myvcn"
  compartment_id = "${var.compartment_ocid}"
  display_name = "MyVcn"
}

# Setting up an Internet Gateway.
# This is used to send-receive traffic between the Internet and (compute) instances hosted inside our VCN.
resource "oci_core_internet_gateway" "MyInternetGateway" {
  compartment_id = "${var.compartment_ocid}"
  display_name = "MyInternetGateway"
  vcn_id = "${oci_core_virtual_network.MyVcn.id}"
}

# Setting up a Route Table.
# Network traffic intended to head out to the general Internet (0.0.0.0/0) is being routed to our Internet Gateway.
resource "oci_core_route_table" "MyRouteTable" {
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${oci_core_virtual_network.MyVcn.id}"
  display_name = "MyRouteTable"
  route_rules {
    cidr_block = "0.0.0.0/0"
    network_entity_id = "${oci_core_internet_gateway.MyInternetGateway.id}"
  }
}

# Setting up Security Lists.
# The allow/deny rules in the security list determine if a subnet is public or private.
resource "oci_core_security_list" "MyPublicSubnetSecurityList" {
  compartment_id = "${var.compartment_ocid}"
  display_name = "MyPublicSubnetSecurityList"
  vcn_id = "${oci_core_virtual_network.MyVcn.id}"
  egress_security_rules = [{
    destination = "0.0.0.0/0"
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
      protocol = "6"
      source = "${var.VPC-CIDR}"
    }]
}

resource "oci_core_security_list" "MyPrivateSubnetSecurityList" {
  compartment_id = "${var.compartment_ocid}"
  display_name = "MyPrivateSubnetSecurityList"
  vcn_id = "${oci_core_virtual_network.MyVcn.id}"
  egress_security_rules = [{
    protocol = "6"
    destination = "${var.VPC-CIDR}"
  }]
  ingress_security_rules = [{
    protocol = "6"
    source = "${var.VPC-CIDR}"
  }]
}

# Creating a public subnet within our VCN.
# Instances hosted in public-subnets can be accessed from the Internet.
resource "oci_core_subnet" "MyPublicSubnetAD1" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0], "name")}"
  cidr_block = "10.0.1.0/24"
  display_name = "MyPublicSubnetAD1"
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${oci_core_virtual_network.MyVcn.id}"
  route_table_id = "${oci_core_route_table.MyRouteTable.id}"
  security_list_ids = ["${oci_core_security_list.MyPublicSubnetSecurityList.id}"]
  dhcp_options_id = "${oci_core_virtual_network.MyVcn.default_dhcp_options_id}"
}

# Creating another public subnet within our VCN, but in a different availability domain.
# (Compute) Instances running in different availability domain make our infra more reliable and available.
resource "oci_core_subnet" "MyWebSubnetAD2" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[1],"name")}"
  cidr_block = "10.0.2.0/24"
  display_name = "MyWebSubnetAD2"
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${oci_core_virtual_network.MyVcn.id}"
  route_table_id = "${oci_core_route_table.MyRouteTable.id}"
  security_list_ids = ["${oci_core_security_list.MyPublicSubnetSecurityList.id}"]
  dhcp_options_id = "${oci_core_virtual_network.MyVcn.default_dhcp_options_id}"
}

# Creating a private subnet within our VCN.
# Instances hosted in private subnets can only talk to other instances hosted in the same VCN.
resource "oci_core_subnet" "MyPrivateSubnetAD1" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
  cidr_block = "10.0.4.0/24"
  display_name = "MyPrivateSubnetAD1"
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${oci_core_virtual_network.MyVcn.id}"
  route_table_id = "${oci_core_route_table.MyRouteTable.id}"
  security_list_ids = ["${oci_core_security_list.MyPrivateSubnetSecurityList.id}"]
  dhcp_options_id = "${oci_core_virtual_network.MyVcn.default_dhcp_options_id}"
  prohibit_public_ip_on_vnic = "true"
}

# Creating another private subnet within our VCN, but in a different availability domain.
# (Compute) Instances running in different availability domain make our infra more reliable and available.
resource "oci_core_subnet" "MyPrivateSubnetAD2" {
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[1],"name")}"
  cidr_block = "10.0.5.0/24"
  display_name = "MyPrivateSubnetAD2"
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${oci_core_virtual_network.MyVcn.id}"
  route_table_id = "${oci_core_route_table.MyRouteTable.id}"
  security_list_ids = ["${oci_core_security_list.MyPrivateSubnetSecurityList.id}"]
  dhcp_options_id = "${oci_core_virtual_network.MyVcn.default_dhcp_options_id}"
  prohibit_public_ip_on_vnic = "true"

}
