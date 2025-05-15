// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
  type = string
  default = "ocid1.tenancy.oc1..aaaaaaaajoua5f4hwv5dtcwq43hhk7d55m4uxqfg4pwj5uipmjcqpht4upgq"
}
variable "user_ocid" {
  type = string
  default = "TODO"
}
variable "fingerprint" {
  type = string
  default = "TODO"
}
variable "private_key_path" {
  type = string
  default = "TODO"
}
variable "region" {
  type = string
  default = "us-ashburn-1"
}
variable "compartment_id" {
  type = string
  default = "ocid1.compartment.oc1..aaaaaaaayxkbos7zkio4jk7sawovt7phmz3plakrsvfyxcseewtbqbgvzrxq"
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_id
}

resource "oci_core_subnet" "test_subnet" {
  cidr_block     = "10.0.0.0/24"
  compartment_id = var.compartment_id
  vcn_id         = oci_core_vcn.test_vcn.id
}

resource "oci_core_network_security_group" "test_nsg" {
  compartment_id = var.compartment_id
  display_name   = "tfNsgForPipeline"
  vcn_id         = oci_core_vcn.test_vcn.id
}

variable "opensearch_cluster_pipeline_data_prepper_configuration_body" {
  default = "source_coordination:\n  store:\n    oci-object-bucket:\n      name: data-prepper-source-coordination-testing\n      namespace: idv3bncjikjv"
}


variable "opensearch_cluster_pipeline_display_name" {
  default = "OpensearchExamplePipeline"
}

variable "opensearch_cluster_pipeline_freeform_tags" {
  default = { "bar-key" = "value" }
}

variable "opensearch_cluster_pipeline_id" {
  default = "id"
}

variable "opensearch_cluster_pipeline_memory_gb" {
  default = 8
}

variable "opensearch_cluster_pipeline_node_count" {
  default = 1
}

variable "opensearch_cluster_pipeline_ocpu_count" {
  default = 1
}

variable "opensearch_cluster_pipeline_opc_dry_run" {
  default = false
}

variable "opensearch_cluster_pipeline_node_shape" {
  default = "VM.Standard.E3.Flex"
}

locals  {
  pipeline_configuration_body = "version: 2\npipeline_configurations:\n  oci:\n    secrets:\n      opensearch-username:\n        secret_id: {{username-vaultsecret}}\n        refresh_interval: PT2H\n      opensearch-password:\n        secret_id: {{password-vaultsecret}}\n        refresh_interval: PT2H\najapraka-log-pipeline:\n  source:\n    oci-object:\n      acknowledgments: true\n      codec:\n        newline:\n      compression: none\n      scan:\n        scheduling:\n          interval: PT30S\n        buckets:\n          - bucket:\n              namespace: idv3bncjikjv\n              name: data_prepper_integration_test_object_storage_source_bucket_0\n              region: us-ashburn-1\n  sink:\n    - opensearch:\n        hosts: [ {{clusterOCID}} ]\n        username: $${{oci_secrets:opensearch-username}}\n        password: $${{oci_secrets:opensearch-password}}\n        insecure: false\n        index: pipeline-stage-testing-index-1"
}

variable "opensearch_cluster_pipeline_reverse_connection_endpoints_customer_fqdn" {
  default = "gzltoensueoa.streaming.us-ashburn-1.oci.oraclecloud.com"
}

variable "opensearch_cluster_pipeline_reverse_connection_endpoints_customer_ip" {
  default = "10.0.0.211"
}

variable "opensearch_cluster_pipeline_state" {
  default = "ACTIVE"
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
  pipeline_configuration_body     = local.pipeline_configuration_body

  #Optional
  //defined_tags  = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.opensearch_cluster_pipeline_defined_tags_value)
  freeform_tags = var.opensearch_cluster_pipeline_freeform_tags
  nsg_id        = oci_core_network_security_group.test_nsg.id
  opc_dry_run   = var.opensearch_cluster_pipeline_opc_dry_run
  reverse_connection_endpoints {
    #Required
    customer_fqdn = var.opensearch_cluster_pipeline_reverse_connection_endpoints_customer_fqdn
    customer_ip   = var.opensearch_cluster_pipeline_reverse_connection_endpoints_customer_ip
  }
  subnet_compartment_id = var.compartment_id
  subnet_id             = oci_core_subnet.test_subnet.id
  vcn_compartment_id    = var.compartment_id
  vcn_id                = oci_core_vcn.test_vcn.id
}

data "oci_opensearch_opensearch_cluster_pipelines" "test_opensearch_cluster_pipelines" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name          = var.opensearch_cluster_pipeline_display_name
  //id                    = var.opensearch_cluster_pipeline_id
  //pipeline_component_id = oci_opensearch_pipeline_component.test_pipeline_component.id
  //state                 = var.opensearch_cluster_pipeline_state
}
