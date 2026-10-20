// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

terraform {
  required_providers {
    oci = {
      source = "oracle/oci"
    }
  }
}

variable "region" {
  type = string
}

variable "config_file_profile" {
  type = string
}

variable "compartment_id" {
  type = string
}

variable "vault_secret_id" {
  type = string
}

provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = var.config_file_profile
  region              = var.region
}

resource "oci_core_vcn" "opensearch" {
  compartment_id = var.compartment_id
  cidr_block     = "10.91.0.0/16"
  display_name   = "opensearch-vault-example-vcn"
}

resource "oci_core_subnet" "opensearch" {
  compartment_id = var.compartment_id
  vcn_id         = oci_core_vcn.opensearch.id
  cidr_block     = "10.91.1.0/24"
  display_name   = "opensearch-vault-example-subnet"
}

resource "oci_opensearch_opensearch_cluster" "vault_auth" {
  compartment_id                     = var.compartment_id
  display_name                       = "opensearch-vault-auth-example"
  software_version                   = "2.11.0"
  vcn_id                             = oci_core_vcn.opensearch.id
  vcn_compartment_id                 = var.compartment_id
  subnet_id                          = oci_core_subnet.opensearch.id
  subnet_compartment_id              = var.compartment_id
  master_node_count                  = 1
  master_node_host_type              = "FLEX"
  master_node_host_ocpu_count        = 1
  master_node_host_memory_gb         = 20
  data_node_count                    = 3
  data_node_host_type                = "FLEX"
  data_node_host_ocpu_count          = 2
  data_node_host_memory_gb           = 20
  data_node_storage_gb               = 50
  opendashboard_node_count           = 1
  opendashboard_node_host_ocpu_count = 2
  opendashboard_node_host_memory_gb  = 10
  security_mode                      = "ENFORCING"

  security_master_user_auth_details {
    user_details {
      username = "testuser"

      password_details {
        source_type     = "OCI_VAULT_SECRET"
        vault_secret_id = var.vault_secret_id
      }
    }
  }
}

output "cluster_id" {
  value = oci_opensearch_opensearch_cluster.vault_auth.id
}
