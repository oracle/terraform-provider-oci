// Copyright (c) 2017, 2025, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
}

variable "cluster_profile" {
  default = "HADOOP"
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
}

variable "secret_id" {
}

variable "kms_key_id" {
}

variable "bootstrap_script_url" {
}

variable "bds_instance_cluster_public_key" {
}

variable "bds_instance_cluster_version" {
  default = "ODH2_0"
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
  default = true
}

variable "bds_instance_is_secure" {
  default = true
}

variable "bds_instance_network_config_cidr_block" {
  default = "111.112.0.0/16"
}

variable "bds_instance_network_config_is_nat_gateway_required" {
  default = true
}

variable "bds_instance_nodes_block_volume_size_in_gbs" {
  default = 150
}

variable "bds_instance_worker_nodes_block_volume_size_in_gbs" {
  default = 150
}

variable "bds_instance_nodes_shape" {
  default = "VM.Standard.Generic"
}

variable "bds_instance_worker_node_shape" {
  default = "VM.Standard.Generic"
}

variable "bds_instance_compute_only_worker_node_shape" {
  default = "VM.Standard.Generic"
}

variable "bds_instance_compute_only_worker_memory_per_node" {
  default = 32
}

variable "bds_instance_compute_only_worker_ocpu_per_node" {
  default = 3
}

variable "bds_instance_edge_node_shape" {
  default = "VM.Standard.Generic"
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

variable "nodes_instance_id" {
  description = "List of BDS nodes instance OCID to be removed."
  type = list(string)
  default = ["<instanceOcid1>", "<instanceOcid2>"]
}

variable "is_force_remove_enabled" {
  description = "Force removal even if graceful removal fails."
  type = bool
  default = false
}

data "oci_core_services" "test_bds_services" {
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_bds_bds_instance" "test_bds_instance" {
  #Required
  cluster_admin_password = var.bds_instance_cluster_admin_password // Comment this if secret_id usage is preferred.
  // secret_id              = var.secret_id
  cluster_public_key   = var.bds_instance_cluster_public_key
  cluster_version      = var.bds_instance_cluster_version
  compartment_id       = var.compartment_id
  display_name         = var.bds_instance_display_name
  is_high_availability = var.bds_instance_is_high_availability
  is_secure            = var.bds_instance_is_secure
  cluster_profile      = var.cluster_profile
  // kms_key_id             = var.kms_key_id
  // bootstrap_script_url = var.bootstrap_script_url
  // remove_nodes = var.nodes_instance_id
  // is_force_remove_enabled = var.is_force_remove_enabled

  master_node {
    #Required
    shape = var.bds_instance_nodes_shape

    subnet_id                = var.subnet_id
    block_volume_size_in_gbs = var.bds_instance_nodes_block_volume_size_in_gbs
    number_of_nodes          = 2
    shape_config {
      memory_in_gbs = 96
      ocpus         = 8
    }
  }

  util_node {
    #Required
    shape = var.bds_instance_nodes_shape

    subnet_id                = var.subnet_id
    block_volume_size_in_gbs = var.bds_instance_nodes_block_volume_size_in_gbs
    number_of_nodes          = 2
    shape_config {
      memory_in_gbs = 96
      ocpus         = 8
    }
  }

  worker_node {
    #Required
    shape                    = var.bds_instance_worker_node_shape
    block_volume_size_in_gbs = var.bds_instance_nodes_block_volume_size_in_gbs
    subnet_id                = var.subnet_id
    number_of_nodes          = 5
    shape_config {
      memory_in_gbs = 96
      ocpus         = 8
    }
  }
  edge_node {
    #Required
    shape = var.bds_instance_edge_node_shape

    subnet_id                = var.subnet_id
    block_volume_size_in_gbs = var.bds_instance_worker_nodes_block_volume_size_in_gbs
    number_of_nodes          = 1
    shape_config {
      memory_in_gbs = var.bds_instance_edge_memory_per_node
      ocpus         = var.bds_instance_edge_ocpu_per_node
    }
  }

  compute_only_worker_node {
    #Required
    shape = var.bds_instance_compute_only_worker_node_shape

    subnet_id                = var.subnet_id
    block_volume_size_in_gbs = var.bds_instance_worker_nodes_block_volume_size_in_gbs
    number_of_nodes          = 1
    shape_config {
      memory_in_gbs = var.bds_instance_compute_only_worker_memory_per_node
      ocpus         = var.bds_instance_compute_only_worker_ocpu_per_node
    }
  }

  ignore_existing_nodes_shape = ["worker", "master", "utility"]

  is_cloud_sql_configured = false


  # Change value to true for use of Kafka cluster
  is_kafka_configured = false

  # Uncomment kafka_broker_node block for use of Kafka cluster
  /*
   kafka_broker_node {
     #Required
     shape = var.bds_instance_compute_only_worker_node_shape

     subnet_id                = var.subnet_id
     block_volume_size_in_gbs = var.bds_instance_worker_nodes_block_volume_size_in_gbs
     number_of_nodes          = 1
     shape_config {
       memory_in_gbs = var.bds_instance_compute_only_worker_memory_per_node
       ocpus         = var.bds_instance_compute_only_worker_ocpu_per_node
     }
   }
   */

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