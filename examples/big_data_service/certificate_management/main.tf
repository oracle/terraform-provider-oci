// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
}

variable "compartment_id" {
}

variable "subnet_id" {
}

variable "bds_instance_cluster_admin_password" {
	default = "T3JhY2xlVGVhbVVTQSExMjM="
}

variable "kms_key_id" {
}

variable "cluster_profile" {
}

variable "bds_instance_cluster_public_key" {
	default = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDpUa4zUZKyU3AkW9yoJTBDO550wpWZOXdHswfRq75gbJ2ZYlMtifvwiO3qUL/RIZSC6e1wA5OL2LQ97UaHrLLPXgjvKGVIDRHqPkzTOayjJ4ZA7NPNhcu6f/OxhKkCYF3TAQObhMJmUSMrWSUeufaRIujDz1HHqazxOgFk09fj4i2dcGnfPcm32t8a9MzlsHSmgexYCUwxGisuuWTsnMgxbqsj6DaY51l+SEPi5tf10iFmUWqziF0eKDDQ/jHkwLJ8wgBJef9FSOmwJReHcBY+NviwFTatGj7Cwtnks6CVomsFD+rAMJ9uzM8SCv5agYunx07hnEXbR9r/TXqgXGfN bdsclusterkey@oracleoci.com"
}

variable "bds_instance_cluster_version" {
	default = "ODH1"
}

variable "bds_instance_defined_tags_value" {
	default = "value"
}

variable "bds_instance_display_name" {
	default = "displayName2"
}

variable "bds_instance_freeform_tags" {
	default = {
		"bar-key" = "value"
	}
}

variable "bds_instance_is_high_availability" {
	default = false
}

variable "bds_instance_is_secure" {
	default = false
}

variable "bds_instance_network_config_cidr_block" {
	default = "111.112.0.0/16"
}

variable "bds_instance_network_config_is_nat_gateway_required" {
	default = false
}

variable "bds_instance_nodes_block_volume_size_in_gbs" {
	default = 150
}

variable "bds_instance_worker_nodes_block_volume_size_in_gbs" {
	default = 150
}

variable "bds_instance_nodes_shape" {
	default = "VM.Standard2.4"
}

variable "bds_instance_worker_node_shape" {
	default = "VM.Standard2.4"
}

variable "bds_instance_compute_only_worker_node_shape" {
	default = "VM.Standard.E4.Flex"
}

variable "bds_instance_compute_only_worker_memory_per_node" {
	default = 32
}

variable "bds_instance_compute_only_worker_ocpu_per_node" {
	default = 3
}

variable "bds_instance_edge_node_shape" {
	default = "VM.Standard.E4.Flex"
}

variable "bds_instance_edge_memory_per_node" {
	default = 32
}

variable "bds_instance_edge_ocpu_per_node" {
	default = 3
}

variable "bds_instance_state" {
	default = "ACTIVE"
}

data "oci_core_services" "test_bds_services" {
}

#Uncomment this when running in home region (PHX)
#variable "tag_namespace_description" {
#  default = "Just a test"
#}

#Uncomment this when running in home region (PHX)
#variable "tag_namespace_name" {
#  default = "testexamples-tag-namespace"
#}

provider "oci" {
	tenancy_ocid     = var.tenancy_ocid
#	user_ocid        = var.user_ocid
#	fingerprint      = var.fingerprint
#	private_key_path = var.private_key_path
	auth             = "SecurityToken"
	config_file_profile = "terraform-federation-test"
	region           = var.region
}

#Uncomment this when running in home region (PHX)
#resource "oci_identity_tag_namespace" "tag-namespace1" {
#  #Required
#  compartment_id = var.tenancy_ocid
#  description    = var.tag_namespace_description
#  name           = var.tag_namespace_name
#}

#Uncomment this when running in home region (PHX)
#resource "oci_identity_tag" "tag1" {
#  #Required
#  description      = "tf example tag"
#  name             = "tf-example-tag"
#  tag_namespace_id = oci_identity_tag_namespace.tag-namespace1.id
#}

resource "oci_core_vcn" "vcn_bds" {
	cidr_block     = "111.111.0.0/16"
	compartment_id = var.compartment_id
	display_name   = "BDS_VCN"
	dns_label      = "bdsvcn"
}

resource "oci_core_service_gateway" "export_sgw" {
	compartment_id = var.compartment_id
	display_name   = "sgw"

	services {
		service_id = data.oci_core_services.test_bds_services.services[0]["id"]
	}

	vcn_id = oci_core_vcn.vcn_bds.id
}

resource "oci_core_route_table" "private_rt" {
	compartment_id = var.compartment_id
	display_name   = "private-rt"

	route_rules {
		destination       = data.oci_core_services.test_bds_services.services[0]["cidr_block"]
		destination_type  = "SERVICE_CIDR_BLOCK"
		network_entity_id = oci_core_service_gateway.export_sgw.id
	}

	vcn_id = oci_core_vcn.vcn_bds.id
}

resource "oci_core_subnet" "regional_subnet_bds" {
	cidr_block        = "111.111.0.0/24"
	display_name      = "regionalSubnetBds"
	dns_label         = "regionalbds"
	compartment_id    = var.compartment_id
	vcn_id            = oci_core_vcn.vcn_bds.id
	security_list_ids = [oci_core_vcn.vcn_bds.default_security_list_id]
	route_table_id    = oci_core_vcn.vcn_bds.default_route_table_id
	dhcp_options_id   = oci_core_vcn.vcn_bds.default_dhcp_options_id
}

resource "oci_bds_bds_instance" "test_bds_instance" {
	#Required
	cluster_admin_password = var.bds_instance_cluster_admin_password
	cluster_public_key     = var.bds_instance_cluster_public_key
	cluster_version        = var.bds_instance_cluster_version
	compartment_id         = var.compartment_id
	display_name           = var.bds_instance_display_name
	is_high_availability   = var.bds_instance_is_high_availability
	is_secure              = var.bds_instance_is_secure
	kms_key_id             = var.kms_key_id
	cluster_profile        = var.cluster_profile
	#bootstrap_script_url = "https://objectstorage.us-ashburn-1.oraclecloud.com/p/5cwJG7AZBnrZsbTCOVc2TacwIcvND6pCxRwSm-Te7GmOY5TkS8zV9FpVgSuLCH4W/n/idpbwtq1b3ta/b/bucket-20230214-1316/o/execute_bootstrap_script.sh"

	master_node {
		#Required
		shape = var.bds_instance_nodes_shape

		subnet_id                = var.subnet_id
		block_volume_size_in_gbs = var.bds_instance_nodes_block_volume_size_in_gbs
		number_of_nodes          = 1
	}

	util_node {
		#Required
		shape = var.bds_instance_nodes_shape

		subnet_id                = var.subnet_id
		block_volume_size_in_gbs = var.bds_instance_nodes_block_volume_size_in_gbs
		number_of_nodes          = 1
	}

	worker_node {
		#Required
		shape = var.bds_instance_worker_node_shape

		subnet_id                = var.subnet_id
		block_volume_size_in_gbs = var.bds_instance_worker_nodes_block_volume_size_in_gbs
		number_of_nodes          = 3
	}

	#edge_node {
		#Required
	#	shape = var.bds_instance_edge_node_shape

	#	subnet_id                = var.subnet_id
	#	block_volume_size_in_gbs = var.bds_instance_worker_nodes_block_volume_size_in_gbs
	#	number_of_nodes          = 1
	#	shape_config {
	#memory_in_gbs = var.bds_instance_edge_memory_per_node
	#		ocpus         = var.bds_instance_edge_ocpu_per_node
	#	}
	#}

	#compute_only_worker_node {
		#Required
	#	shape = var.bds_instance_compute_only_worker_node_shape

	#	subnet_id                = var.subnet_id
	#	block_volume_size_in_gbs = var.bds_instance_worker_nodes_block_volume_size_in_gbs
	#	number_of_nodes          = 1
	#	shape_config {
	#		memory_in_gbs = var.bds_instance_compute_only_worker_memory_per_node
	#		ocpus         = var.bds_instance_compute_only_worker_ocpu_per_node
	#	}
	#}




	#   cloud_sql_details {
	#     shape                    = "VM.Standard2.4"
	#     block_volume_size_in_gbs = 1000
	#   }

	is_cloud_sql_configured = false


	#Change value to true for use of Kafka cluster
	is_kafka_configured = false

	#Uncomment kafka_broker_node block for use of Kafka cluster
	#kafka_broker_node {
	#Required
	#  shape = var.bds_instance_compute_only_worker_node_shape

	#  subnet_id                = var.subnet_id
	#  block_volume_size_in_gbs = var.bds_instance_worker_nodes_block_volume_size_in_gbs
	#  number_of_nodes          = 1
	#  shape_config {
	#    memory_in_gbs = var.bds_instance_compute_only_worker_memory_per_node
	#    ocpus         = var.bds_instance_compute_only_worker_ocpu_per_node
	#  }
	#}

	#Optional
	#Uncomment this when running in home region (PHX)
	#  defined_tags = {
	#    "${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = var.bds_instance_defined_tags_value
	#  }
	freeform_tags = var.bds_instance_freeform_tags
	network_config {
		#Optional
		cidr_block              = var.bds_instance_network_config_cidr_block
		is_nat_gateway_required = var.bds_instance_network_config_is_nat_gateway_required
	}
}

data "oci_bds_bds_instances" "test_bds_instances" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = oci_bds_bds_instance.test_bds_instance.display_name
	state        = "ACTIVE"
}

data "oci_bds_bds_instance" "test_bds_instance" {
	#Required
	bds_instance_id = oci_bds_bds_instance.test_bds_instance.id
}

resource "oci_bds_bds_instance_operation_certificate_managements_management" "test_bds_instance_operation_certificate_managements_management" {
	bds_instance_id = oci_bds_bds_instance.test_bds_instance.id
	cluster_admin_password = var.bds_instance_cluster_admin_password
	enable_operation_certificate_management = "true"
	renew_operation_certificate_management = "false"
	services = ["OOZIE"]
}
