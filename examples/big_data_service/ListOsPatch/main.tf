
	# Need to have this block even though it's empty; for import testing
	provider "oci" {
	}
	
	variable "tenancy_ocid" {
		}

	variable "ssh_public_key" {
		}

	variable "region" {
		default = "us-ashburn-1"
	}

	
data "oci_bds_bds_instance_list_os_patches" "test_bds_instance_list_os_patches" {
bds_instance_id = "${oci_bds_bds_instance.test_bds_instance.id}"
}
variable "compartment_id" { }
variable "subnet_id" {  }

resource "oci_bds_bds_instance" "test_bds_instance" {
cluster_admin_password = "T3JhY2xlVGVhbVVTQSExMjM="
cluster_public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDpUa4zUZKyU3AkW9yoJTBDO550wpWZOXdHswfRq75gbJ2ZYlMtifvwiO3qUL/RIZSC6e1wA5OL2LQ97UaHrLLPXgjvKGVIDRHqPkzTOayjJ4ZA7NPNhcu6f/OxhKkCYF3TAQObhMJmUSMrWSUeufaRIujDz1HHqazxOgFk09fj4i2dcGnfPcm32t8a9MzlsHSmgexYCUwxGisuuWTsnMgxbqsj6DaY51l+SEPi5tf10iFmUWqziF0eKDDQ/jHkwLJ8wgBJef9FSOmwJReHcBY+NviwFTatGj7Cwtnks6CVomsFD+rAMJ9uzM8SCv5agYunx07hnEXbR9r/TXqgXGfN bdsclusterkey@oracleoci.com"
cluster_version = "ODH1"
compartment_id = "${var.compartment_id}"
compute_only_worker_node {
block_volume_size_in_gbs = "150"
number_of_nodes = "2"
shape = "VM.Standard.E4.Flex"
shape_config {
memory_in_gbs = "32"
ocpus = "3"
}
subnet_id = "${var.subnet_id}"
}
display_name = "displayName"
edge_node {
block_volume_size_in_gbs = "150"
number_of_nodes = "2"
shape = "VM.Standard.E4.Flex"
shape_config {
memory_in_gbs = "32"
ocpus = "3"
}
subnet_id = "${var.subnet_id}"
}
is_high_availability = "true"
is_secure = "true"
master_node {
block_volume_size_in_gbs = "150"
number_of_nodes = "2"
shape = "VM.Standard.E4.Flex"
shape_config {
memory_in_gbs = "32"
ocpus = "3"
}
subnet_id = "${var.subnet_id}"
}
util_node {
block_volume_size_in_gbs = "150"
number_of_nodes = "2"
shape = "VM.Standard.E4.Flex"
shape_config {
memory_in_gbs = "32"
ocpus = "3"
}
subnet_id = "${var.subnet_id}"
}
worker_node {
block_volume_size_in_gbs = "150"
number_of_nodes = "3"
shape = "VM.Standard2.4"
subnet_id = "${var.subnet_id}"
}
}

resource "oci_core_subnet" "test_subnet" {
cidr_block = "10.0.0.0/24"
compartment_id = "${var.compartment_id}"
lifecycle {
ignore_changes = ["defined_tags"]
}
vcn_id = "${oci_core_vcn.test_vcn.id}"
}

resource "oci_core_vcn" "test_vcn" {
cidr_block = "10.0.0.0/16"
compartment_id = "${var.compartment_id}"
lifecycle {
ignore_changes = ["defined_tags"]
}
}
