# This configuration generally implements this - https://community.oracle.com/community/cloud_computing/bare-metal/blog/2017/01/12/secure-mongodb-on-oracle-bare-metal-cloud-services

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key" {}
variable "private_key_path" {}
variable "compartment_ocid" {}
variable "ssh_key_4_metadata" {}

provider "baremetal" {
  tenancy_ocid = "${var.tenancy_ocid}"
  user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  private_key = "${var.private_key}"
}

variable "VPC-CIDR" {
  default = "10.0.0.0/26"
}

variable "ADs" {
  default = ["Uocm:PHX-AD-1", "Uocm:PHX-AD-2", "Uocm:PHX-AD-3"]
}

variable "PubSubnetAD1CIDR" {
  default = "10.0.0.0/28"
}

variable "PrivSubnetAD1CIDR" {
  default = "10.0.0.16/28"
}

variable "PrivSubnetAD2CIDR" {
  default = "10.0.0.32/28"
}

variable "BastSubnetAD1CIDR" {
  default = "10.0.0.48/28"
}

variable "Oracle-Linux-7_3" {
  default = "ocid1.image.oc1.phx.aaaaaaaaifdnkw5d7xvmwfsfw2rpjpxe56viepslmmisuyy64t3q4aiquema"
}

variable "MongoDBShape" {
  default = "BM.DenseIO1.36"
}

variable "BastionShape" {
  default = "VM.Standard1.1"
}

variable "BastionBootStrap" {
  default = "./bastion.sh"
}

variable "MongoDBBootStrap" {
  default = "./MongoDB.sh"
}

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
  availability_domain = "${var.ADs[0]}"
  cidr_block = "${var.PubSubnetAD1CIDR}"
  display_name = "PublicSubnetAD1"
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${baremetal_core_virtual_network.MongoDB.id}"
  route_table_id = "${baremetal_core_route_table.MongoDB.id}"
  security_list_ids = ["${baremetal_core_security_list.PublicSubnet.id}"]
}

resource "baremetal_core_subnet" "PrivSubnetAD1" {
  availability_domain = "${var.ADs[0]}"
  cidr_block = "${var.PrivSubnetAD1CIDR}"
  display_name = "PrivateSubnetAD1"
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${baremetal_core_virtual_network.MongoDB.id}"
  route_table_id = "${baremetal_core_route_table.MongoDB.id}"
  security_list_ids = ["${baremetal_core_security_list.PrivateSubnets.id}"]
}

resource "baremetal_core_subnet" "PrivSubnetAD2" {
  availability_domain = "${var.ADs[1]}"
  cidr_block = "${var.PrivSubnetAD2CIDR}"
  display_name = "PrivateSubnetAD2"
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${baremetal_core_virtual_network.MongoDB.id}"
  route_table_id = "${baremetal_core_route_table.MongoDB.id}"
  security_list_ids = ["${baremetal_core_security_list.PrivateSubnets.id}"]
}

resource "baremetal_core_subnet" "BastionSubnetAD1" {
  availability_domain = "${var.ADs[0]}"
  cidr_block = "${var.BastSubnetAD1CIDR}"
  display_name = "BastionSubnetAD1"
  compartment_id = "${var.compartment_ocid}"
  vcn_id = "${baremetal_core_virtual_network.MongoDB.id}"
  route_table_id = "${baremetal_core_route_table.MongoDB.id}"
  security_list_ids = ["${baremetal_core_security_list.BastionSubnet.id}"]
}

resource "baremetal_core_instance" "MongoDBBast" {
    availability_domain = "${var.ADs[0]}"
    compartment_id = "${var.compartment_ocid}"
    display_name = "MongoDB-Bastion"
    image = "${var.Oracle-Linux-7_3}"
    shape = "${var.BastionShape}"
    subnet_id = "${baremetal_core_subnet.BastionSubnetAD1.id}"
    metadata {
        ssh_authorized_keys = "${var.ssh_key_4_metadata}"
	user_data = "${base64encode(file(var.BastionBootStrap))}" 
    }
}

resource "baremetal_core_instance" "MongoDBAD1" {
    availability_domain = "${var.ADs[0]}"
    compartment_id = "${var.compartment_ocid}"
    display_name = "MongoDBAD1"
    image = "${var.Oracle-Linux-7_3}"
    shape = "${var.MongoDBShape}"
    subnet_id = "${baremetal_core_subnet.PrivSubnetAD1.id}"
    metadata {
        ssh_authorized_keys = "${var.ssh_key_4_metadata}"
        user_data = "${base64encode(file(var.MongoDBBootStrap))}"
    }
}

resource "baremetal_core_instance" "MongoDBAD2" {
    availability_domain = "${var.ADs[1]}"
    compartment_id = "${var.compartment_ocid}"
    display_name = "MongoDBAD2"
    image = "${var.Oracle-Linux-7_3}"
    shape = "${var.MongoDBShape}"
    subnet_id = "${baremetal_core_subnet.PrivSubnetAD2.id}"
    metadata {
        ssh_authorized_keys = "${var.ssh_key_4_metadata}"
        user_data = "${base64encode(file(var.MongoDBBootStrap))}"
    }
}
