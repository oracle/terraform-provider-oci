// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "opensearch_cluster_pipeline_data_prepper_configuration_body" {
  default = "dataPrepperConfigurationBody"
}

variable "opensearch_cluster_pipeline_defined_tags_value" {
  default = "value"
}

variable "opensearch_cluster_pipeline_display_name" {
  default = "displayName"
}

variable "opensearch_cluster_pipeline_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "opensearch_cluster_pipeline_id" {
  default = "id"
}

variable "opensearch_cluster_pipeline_memory_gb" {
  default = 10
}

variable "opensearch_cluster_pipeline_node_count" {
  default = 10
}

variable "opensearch_cluster_pipeline_node_shape" {
  default = "nodeShape"
}

variable "opensearch_cluster_pipeline_ocpu_count" {
  default = 10
}

variable "opensearch_cluster_pipeline_opc_dry_run" {
  default = false
}

variable "opensearch_cluster_pipeline_pipeline_configuration_body" {
  default = "pipelineConfigurationBody"
}

variable "opensearch_cluster_pipeline_reverse_connection_endpoints_customer_fqdn" {
  default = "customerFqdn"
}

variable "opensearch_cluster_pipeline_reverse_connection_endpoints_customer_ip" {
  default = "customerIp"
}

variable "opensearch_cluster_pipeline_state" {
  default = "AVAILABLE"
}



provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

resource "oci_opensearch_opensearch_cluster_pipeline" "test_opensearch_cluster_pipeline" {
  #Required
  compartment_id                  = var.compartment_id
  data_prepper_configuration_body = var.opensearch_cluster_pipeline_data_prepper_configuration_body
  display_name                    = var.opensearch_cluster_pipeline_display_name
  memory_gb                       = var.opensearch_cluster_pipeline_memory_gb
  node_count                      = var.opensearch_cluster_pipeline_node_count
  ocpu_count                      = var.opensearch_cluster_pipeline_ocpu_count
  pipeline_configuration_body     = var.opensearch_cluster_pipeline_pipeline_configuration_body

  #Optional
  defined_tags  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.opensearch_cluster_pipeline_defined_tags_value)
  freeform_tags = var.opensearch_cluster_pipeline_freeform_tags
  node_shape    = var.opensearch_cluster_pipeline_node_shape
  nsg_id        = oci_opensearch_nsg.test_nsg.id
  opc_dry_run   = var.opensearch_cluster_pipeline_opc_dry_run
  reverse_connection_endpoints {
    #Required
    customer_fqdn = var.opensearch_cluster_pipeline_reverse_connection_endpoints_customer_fqdn
    customer_ip   = var.opensearch_cluster_pipeline_reverse_connection_endpoints_customer_ip
  }
  subnet_compartment_id = oci_identity_compartment.test_compartment.id
  subnet_id             = oci_core_subnet.test_subnet.id
  vcn_compartment_id    = oci_identity_compartment.test_compartment.id
  vcn_id                = oci_core_vcn.test_vcn.id
}

data "oci_opensearch_opensearch_cluster_pipelines" "test_opensearch_cluster_pipelines" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name          = var.opensearch_cluster_pipeline_display_name
  id                    = var.opensearch_cluster_pipeline_id
  pipeline_component_id = oci_opensearch_pipeline_component.test_pipeline_component.id
  state                 = var.opensearch_cluster_pipeline_state
}

