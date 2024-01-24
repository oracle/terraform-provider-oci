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

variable "compartment_ocid" {
}

variable "kms_vault_id" {
}

variable "region" {
  default = "us-ashburn-1"
}

# Provide the SSH public key to be set on each node in the node pool on launch.
variable "node_pool_ssh_public_key" {
  
}

variable "node_pool_node_config_details_size" {
    default = 1
}

variable "cluster_defined_tags_value" {
  default = "value"
}

variable "cluster_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "cluster_options_persistent_volume_config_defined_tags_value" {
  default = "value"
}

variable "cluster_options_persistent_volume_config_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "cluster_options_service_lb_config_defined_tags_value" {
  default = "value"
}

variable "cluster_options_service_lb_config_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "node_pool_defined_tags_value" {
  default = "value"
}

variable "node_pool_freeform_tags" {
  default = { "Department" = "Finance" }
}

variable "node_pool_node_config_details_defined_tags_value" {
  default = "value"
}

variable "node_pool_node_config_details_freeform_tags" {
  default = { "Department" = "Finance" }
}

provider "oci" {
  region           = var.region
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
}

data "oci_identity_availability_domain" "ad1" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}

data "oci_identity_availability_domain" "ad2" {
  compartment_id = var.tenancy_ocid
  ad_number      = 2
}

//DEPENDENCIES
variable defined_tag_namespace_name {
  default = "test"
}
resource "oci_identity_tag_namespace" "tag-namespace1" {
  #Required
  compartment_id = var.tenancy_ocid
  description = "example tag namespace"
  name = var.defined_tag_namespace_name != "" ? var.defined_tag_namespace_name : "example-tag-namespace-all"

  is_retired = false
}

resource "oci_identity_tag" "tag1" {
  #Required
  description = "example tag"
  name = "example-tag"
  tag_namespace_id = oci_identity_tag_namespace.tag-namespace1.id

  is_retired = false
}

resource "oci_kms_vault" "test_vault" {
  #Required
  compartment_id = var.compartment_ocid
  display_name = "tf_test"
  vault_type     = "DEFAULT"
}

resource "oci_kms_key" "test_key" {
  #Required
  compartment_id = var.compartment_ocid
  display_name = "tf-test-key"
  key_shape {
    #Required
    algorithm = "AES"
    length = 32
  }
  management_endpoint = oci_kms_vault.test_vault.management_endpoint
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_ocid
  display_name   = "tfVcnForClusters"
}

resource "oci_core_internet_gateway" "test_ig" {
  compartment_id = var.compartment_ocid
  display_name   = "tfClusterInternetGateway"
  vcn_id         = oci_core_vcn.test_vcn.id
}

resource "oci_core_route_table" "test_route_table" {
  compartment_id = var.compartment_ocid
  vcn_id         = oci_core_vcn.test_vcn.id
  display_name   = "tfClustersRouteTable"

  route_rules {
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = oci_core_internet_gateway.test_ig.id
  }
}

resource "oci_core_subnet" "nodePool_Subnet_1" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad1.name
  cidr_block          = "10.0.22.0/24"
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.test_vcn.id

  # Provider code tries to maintain compatibility with old versions.
  security_list_ids = [oci_core_vcn.test_vcn.default_security_list_id]
  display_name      = "tfSubNet1ForNodePool"
  route_table_id    = oci_core_route_table.test_route_table.id
}

resource "oci_core_subnet" "clusterSubnet_1" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad1.name
  cidr_block          = "10.0.20.0/24"
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.test_vcn.id

  # Provider code tries to maintain compatibility with old versions.
  security_list_ids = [oci_core_vcn.test_vcn.default_security_list_id]
  display_name      = "tfSubNet1ForClusters"
  route_table_id    = oci_core_route_table.test_route_table.id
}

resource "oci_core_subnet" "clusterSubnet_2" {
  #Required
  availability_domain = data.oci_identity_availability_domain.ad2.name
  cidr_block          = "10.0.21.0/24"
  compartment_id      = var.compartment_ocid
  vcn_id              = oci_core_vcn.test_vcn.id
  display_name        = "tfSubNet1ForClusters"

  # Provider code tries to maintain compatibility with old versions.
  security_list_ids = [oci_core_vcn.test_vcn.default_security_list_id]
  route_table_id    = oci_core_route_table.test_route_table.id
}

resource "oci_containerengine_cluster" "test_cluster" {
  #Required
  compartment_id     = var.compartment_ocid
  kubernetes_version = "v1.20.11"
  name               = "tfTestCluster"
  vcn_id             = oci_core_vcn.test_vcn.id

  #Optional
  #defined_tags = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.cluster_defined_tags_value)
  defined_tags = {"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "${var.cluster_defined_tags_value}"}
  freeform_tags = var.cluster_freeform_tags
  options {
    service_lb_subnet_ids = [oci_core_subnet.clusterSubnet_1.id, oci_core_subnet.clusterSubnet_2.id]

    #Optional
    add_ons {
      #Optional
      is_kubernetes_dashboard_enabled = "true"
      is_tiller_enabled               = "true"
    }

    admission_controller_options {
      #Optional
      is_pod_security_policy_enabled = true
    }

    kubernetes_network_config {
      #Optional
      pods_cidr     = "10.1.0.0/16"
      services_cidr = "10.2.0.0/16"
    }

    persistent_volume_config {

      #Optional
      #defined_tags  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.cluster_options_persistent_volume_config_defined_tags_value)
      defined_tags  = {"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "${var.cluster_defined_tags_value}"}
      freeform_tags = var.cluster_options_persistent_volume_config_freeform_tags
    }
    service_lb_config {

      #Optional
      #defined_tags  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.cluster_options_service_lb_config_defined_tags_value)
      defined_tags  = {"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "${var.cluster_defined_tags_value}"}
      freeform_tags = var.cluster_options_service_lb_config_freeform_tags
    }
  }
}

resource "oci_containerengine_node_pool" "test_node_pool" {
  #Required
  cluster_id         = oci_containerengine_cluster.test_cluster.id
  compartment_id     = var.compartment_ocid
  kubernetes_version = "v1.20.11"
  name               = "tfPool"
  node_shape         = "VM.Standard2.1"

  #Optional
  defined_tags  = {"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "${var.node_pool_defined_tags_value}"}
  freeform_tags = var.node_pool_freeform_tags
  initial_node_labels {
    #Optional
    key   = "key"
    value = "value"
  }

  node_source_details {
    #Required
    image_id    = local.image_id
    source_type = "IMAGE"
  }

  node_config_details {
    #Required
    placement_configs {
      #Required
      availability_domain = data.oci_identity_availability_domain.ad1.name
      subnet_id           = oci_core_subnet.nodePool_Subnet_1.id
      #optional
      fault_domains = ["FAULT-DOMAIN-1", "FAULT-DOMAIN-3"]
    }
    size = var.node_pool_node_config_details_size

    defined_tags  = {"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "${var.node_pool_defined_tags_value}"}
    freeform_tags = var.node_pool_node_config_details_freeform_tags
  }

  ssh_public_key      = var.node_pool_ssh_public_key
}

output "node_pool" {
  value = {
    id                 = oci_containerengine_node_pool.test_node_pool.id
    kubernetes_version = oci_containerengine_node_pool.test_node_pool.kubernetes_version
    name               = oci_containerengine_node_pool.test_node_pool.name
    subnet_ids         = oci_containerengine_node_pool.test_node_pool.subnet_ids
  }
}

data "oci_containerengine_node_pool_option" "test_node_pool_option" {
  node_pool_option_id = "all"
}

data "oci_core_images" "shape_specific_images" {
  #Required
  compartment_id = var.tenancy_ocid
  shape = "VM.Standard2.1"
}

locals {
  all_images = "${data.oci_core_images.shape_specific_images.images}"
  all_sources = "${data.oci_containerengine_node_pool_option.test_node_pool_option.sources}"

  compartment_images = [for image in local.all_images : image.id if length(regexall("Oracle-Linux-[0-9]*.[0-9]*-20[0-9]*",image.display_name)) > 0 ]

  oracle_linux_images = [for source in local.all_sources : source.image_id if length(regexall("Oracle-Linux-[0-9]*.[0-9]*-20[0-9]*",source.source_name)) > 0]

  image_id = tolist(setintersection( toset(local.compartment_images), toset(local.oracle_linux_images)))[0]

}