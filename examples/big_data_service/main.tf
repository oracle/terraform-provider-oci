// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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
  default = "T3JhY2xlVGVhbVVTQSExMjM="
}

variable "kms_key_id" {
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
  default = true
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
  default = "VM.Standard.Generic"
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
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
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
  bootstrap_script_url = "https://objectstorage.us-ashburn-1.oraclecloud.com/p/Lk5JT9tnUIOG4yLm6S21QVR7m3Rm2uj1RAS2Olx5v14onLU2Y-b0lIc_N0RuUIge/n/idpbwtq1b3ta/b/bucket-20230214-1316/o/execute_bootstrap_script.sh"

  master_node {
    #Required
    shape = "VM.Standard.E4.Flex"

    subnet_id                = var.subnet_id
    block_volume_size_in_gbs = var.bds_instance_nodes_block_volume_size_in_gbs
    number_of_nodes          = 1
    shape_config {
            memory_in_gbs = 120
           ocpus         = 8
        }
  }

  util_node {
    #Required
    shape = "VM.Standard.E4.Flex"

    subnet_id                = var.subnet_id
    block_volume_size_in_gbs = var.bds_instance_nodes_block_volume_size_in_gbs
    number_of_nodes          = 1
    shape_config {
            memory_in_gbs = 120
            ocpus         = 8
        }
  }

  worker_node {
    #Required
    shape = var.bds_instance_worker_node_shape
    block_volume_size_in_gbs = var.bds_instance_nodes_block_volume_size_in_gbs
    subnet_id                = var.subnet_id
    number_of_nodes          = 3
       shape_config {
              memory_in_gbs = 120
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

