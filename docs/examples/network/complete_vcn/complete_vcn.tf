/*
 * Create a default Virtual Cloud Network (VCN) with CIDR 10.0.0.0/16 spanning 
 * across 3 availability domains (PHX-AD1/AD2/AD3)
 * 
 * Bastion Subnet (with security list allowing SSH traffic from Internet)
 * Create 3 subnets (one in each availability domain) to host compute instances 
 * that have SSH access from the Internet. These compute instances provide the
 * Bastion Host functionality.See https://en.wikipedia.org/wiki/Bastion_host
 * These bastion hosts provide the front-end for your compute instances 
 *
 * Private Subnet (with security list allowing traffic only within the VCN)
 * Create 3 subnets (one in each availability domain) to host compute instances 
 * that has SSH access only from the Bastion Hosts. The compute instances in
 * these subnets can host your web, application and database services.
 *
 * LB Subnet (with security list allowing internet traffic for Port 80)
 * Create 2 subnets (one in each availability domain) so that you can use it to 
 * run your load balancer and distribute internet traffic to your applications.
 *
 * Note: Currently, when you create a VCN, VCN creates DefaultSecurityList and 
 * DefaultRouteTable resources. These resources are not used in this configuration.
 *
 * Next Steps: 
 *   - Create Compute instances within the Bastion Subnet and harden these compute
 *     instances
 *   - Create Compute instances within the Private Subnet and configure your web,
 *     application and databases
 *   - Create Load Balancer within the LB Subnet and front-end Internet traffic
 *   - (port 80) to your compute instances within the Private Subnet
 */
variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "compartment_ocid" {}

provider "baremetal" {
  tenancy_ocid = "${var.tenancy_ocid}"
  user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  private_key = "${var.private_key_path}"
}

variable "VPC-CIDR" {
  default = "10.0.0.0/16"
}

variable "ADs" {
  default = ["Uocm:PHX-AD-1", "Uocm:PHX-AD-2", "Uocm:PHX-AD-3"]
}

resource "baremetal_core_virtual_network" "CompleteVCN" {
  cidr_block = "${var.VPC-CIDR}"
  compartment_id = "${var.compartment_ocid}"
  display_name = "CompleteVCN"
}

resource "baremetal_core_internet_gateway" "CompleteIG" {
    compartment_id = "${var.compartment_ocid}"
    display_name = "CompleteIG"
    vcn_id = "${baremetal_core_virtual_network.CompleteVCN.id}"
}

resource "baremetal_core_route_table" "RouteForComplete" {
    compartment_id = "${var.compartment_ocid}"
    vcn_id = "${baremetal_core_virtual_network.CompleteVCN.id}"
    display_name = "RouteTableForComplete"
    route_rules {
        cidr_block = "0.0.0.0/0"
        network_entity_id = "${baremetal_core_internet_gateway.CompleteIG.id}"
    }
}

resource "baremetal_core_security_list" "LBSubnet" {
    compartment_id = "${var.compartment_ocid}"
    display_name = "Public"
    vcn_id = "${baremetal_core_virtual_network.CompleteVCN.id}"
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

resource "baremetal_core_security_list" "PrivateSubnet" {
    compartment_id = "${var.compartment_ocid}"
    display_name = "Private"
    vcn_id = "${baremetal_core_virtual_network.CompleteVCN.id}"
    egress_security_rules = [{
	protocol = "6"
	destination = "${var.VPC-CIDR}"
    }]
    ingress_security_rules = [{
        protocol = "6"
        source = "${var.VPC-CIDR}"
    }]
}

resource "baremetal_core_security_list" "BastionSubnet" {
    compartment_id = "${var.compartment_ocid}"
    display_name = "Bastion"
    vcn_id = "${baremetal_core_virtual_network.CompleteVCN.id}"
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
    },
	{
	protocol = "6"
        source = "${var.VPC-CIDR}"
    }]	
}

resource "baremetal_core_subnet" "LBSubnetAD1" {
  availability_domain = "${var.ADs[0]}"
  cidr_block = "10.0.1.0/24"
  display_name = "LBSubnetAD1"
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${baremetal_core_virtual_network.CompleteVCN.id}"
  route_table_id = "${baremetal_core_route_table.RouteForComplete.id}"
  security_list_ids = ["${baremetal_core_security_list.LBSubnet.id}"]
}

resource "baremetal_core_subnet" "LBSubnetAD2" {
  availability_domain = "${var.ADs[1]}"
  cidr_block = "10.0.2.0/24"
  display_name = "LBSubnetAD2"
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${baremetal_core_virtual_network.CompleteVCN.id}"
  route_table_id = "${baremetal_core_route_table.RouteForComplete.id}"
  security_list_ids = ["${baremetal_core_security_list.LBSubnet.id}"]
}

resource "baremetal_core_subnet" "PrivateSubnetAD1" {
  availability_domain = "${var.ADs[0]}"
  cidr_block = "10.0.4.0/24"
  display_name = "PrivateSubnetAD1"
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${baremetal_core_virtual_network.CompleteVCN.id}"
  route_table_id = "${baremetal_core_route_table.RouteForComplete.id}"
  security_list_ids = ["${baremetal_core_security_list.PrivateSubnet.id}"]
}

resource "baremetal_core_subnet" "PrivateSubnetAD2" {
  availability_domain = "${var.ADs[1]}"
  cidr_block = "10.0.5.0/24"
  display_name = "PrivateSubnetAD2"
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${baremetal_core_virtual_network.CompleteVCN.id}"
  route_table_id = "${baremetal_core_route_table.RouteForComplete.id}"
  security_list_ids = ["${baremetal_core_security_list.PrivateSubnet.id}"]
}

resource "baremetal_core_subnet" "PrivateSubnetAD3" {
  availability_domain = "${var.ADs[2]}"
  cidr_block = "10.0.6.0/24"
  display_name = "PrivateSubnetAD3"
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${baremetal_core_virtual_network.CompleteVCN.id}"
  route_table_id = "${baremetal_core_route_table.RouteForComplete.id}"
  security_list_ids = ["${baremetal_core_security_list.PrivateSubnet.id}"]
}

resource "baremetal_core_subnet" "BastionSubnetAD1" {
  availability_domain = "${var.ADs[0]}"
  cidr_block = "10.0.7.0/24"
  display_name = "BastionSubnetAD1"
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${baremetal_core_virtual_network.CompleteVCN.id}"
  route_table_id = "${baremetal_core_route_table.RouteForComplete.id}"
  security_list_ids = ["${baremetal_core_security_list.BastionSubnet.id}"]
}

resource "baremetal_core_subnet" "BastionSubnetAD2" {
  availability_domain = "${var.ADs[1]}"
  cidr_block = "10.0.8.0/24"
  display_name = "BastionSubnetAD2"
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${baremetal_core_virtual_network.CompleteVCN.id}"
  route_table_id = "${baremetal_core_route_table.RouteForComplete.id}"
  security_list_ids = ["${baremetal_core_security_list.BastionSubnet.id}"]
}

resource "baremetal_core_subnet" "BastionSubnetAD3" {
  availability_domain = "${var.ADs[2]}"
  cidr_block = "10.0.9.0/24"
  display_name = "BastionSubnetAD3"
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${baremetal_core_virtual_network.CompleteVCN.id}"
  route_table_id = "${baremetal_core_route_table.RouteForComplete.id}"
  security_list_ids = ["${baremetal_core_security_list.BastionSubnet.id}"]
}
