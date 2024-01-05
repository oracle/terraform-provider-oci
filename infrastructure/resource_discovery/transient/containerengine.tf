// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_core_internet_gateway" "internet_gateway_containerengine_rd" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "internetGatewayContainerengineRD"
  vcn_id         = "${oci_core_vcn.vcn2_rd.id}"
}

resource "oci_core_route_table" "test_route_table_containerengine_rd" {
  compartment_id = "${var.compartment_ocid}"
  vcn_id         = "${oci_core_vcn.vcn2_rd.id}"
  display_name   = "tfClustersRouteTableContainerengineRD"

  route_rules {
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = "${oci_core_internet_gateway.internet_gateway_containerengine_rd.id}"
  }
}

resource "oci_core_subnet" "clusterSubnet1_RD" {
  #Required
  availability_domain = "${data.oci_identity_availability_domain.ad1.name}"
  cidr_block          = "10.0.20.0/24"
  compartment_id      = "${var.compartment_ocid}"
  vcn_id              = "${oci_core_vcn.vcn2_rd.id}"

  # Provider code tries to maintain compatibility with old versions.
  security_list_ids = ["${oci_core_vcn.vcn2_rd.default_security_list_id}"]
  display_name      = "tfSubNet1ForClusters1RD"
  route_table_id    = "${oci_core_route_table.test_route_table_containerengine_rd.id}"
}

resource "oci_core_subnet" "clusterSubnet2_rd" {
  #Required
  availability_domain = "${data.oci_identity_availability_domain.ad2.name}"
  cidr_block          = "10.0.21.0/24"
  compartment_id      = "${var.compartment_ocid}"
  vcn_id              = "${oci_core_vcn.vcn2_rd.id}"
  display_name        = "tfSubNet2ForClustersRD"

  # Provider code tries to maintain compatibility with old versions.
  security_list_ids = ["${oci_core_vcn.vcn2_rd.default_security_list_id}"]
  route_table_id    = "${oci_core_route_table.test_route_table_containerengine_rd.id}"
}

resource "oci_core_subnet" "nodePool_Subnet1_rd" {
  #Required
  availability_domain = "${data.oci_identity_availability_domain.ad1.name}"
  cidr_block          = "10.0.22.0/24"
  compartment_id      = "${var.compartment_ocid}"
  vcn_id              = "${oci_core_vcn.vcn2_rd.id}"

  # Provider code tries to maintain compatibility with old versions.
  security_list_ids = ["${oci_core_vcn.vcn2_rd.default_security_list_id}"]
  display_name      = "tfSubNet1ForNodePool1RD"
  route_table_id    = "${oci_core_route_table.test_route_table_containerengine_rd.id}"
}

resource "oci_core_subnet" "nodePool_Subnet2_rd" {
  #Required
  availability_domain = "${data.oci_identity_availability_domain.ad2.name}"
  cidr_block          = "10.0.23.0/24"
  compartment_id      = "${var.compartment_ocid}"
  vcn_id              = "${oci_core_vcn.vcn2_rd.id}"

  # Provider code tries to maintain compatibility with old versions.
  security_list_ids = ["${oci_core_vcn.vcn2_rd.default_security_list_id}"]
  display_name      = "tfSubNet2ForNodePool2RD"
  route_table_id    = "${oci_core_route_table.test_route_table_containerengine_rd.id}"
}

resource "oci_containerengine_cluster" "test_cluster_rd" {
  #Required
  compartment_id     = "${var.compartment_ocid}"
  kubernetes_version = "${data.oci_containerengine_cluster_option.test_cluster_option.kubernetes_versions.0}"
  name               = "tfTestClusterRD"
  vcn_id             = "${oci_core_vcn.vcn2_rd.id}"

  #Optional
  options {
    service_lb_subnet_ids = ["${oci_core_subnet.clusterSubnet1_RD.id}", "${oci_core_subnet.clusterSubnet2_rd.id}"]

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
  }
}

resource "oci_containerengine_node_pool" "test_node_pool_rd" {
  #Required
  cluster_id         = "${oci_containerengine_cluster.test_cluster_rd.id}"
  compartment_id     = "${var.compartment_ocid}"
  kubernetes_version = "${data.oci_containerengine_node_pool_option.test_node_pool_option.kubernetes_versions.0}"
  name               = "tfPoolRD"
  node_shape         = "VM.Standard2.1"
  subnet_ids         = ["${oci_core_subnet.nodePool_Subnet1_rd.id}", "${oci_core_subnet.nodePool_Subnet2_rd.id}"]

  #Optional
  initial_node_labels {
    #Optional
    key   = "key"
    value = "value"
  }

  node_source_details {
    #Required
    image_id    = "${data.oci_containerengine_node_pool_option.test_node_pool_option.sources.0.image_id}"
    source_type = "${data.oci_containerengine_node_pool_option.test_node_pool_option.sources.0.source_type}"
  }

  quantity_per_subnet = 2
  ssh_public_key      = "${var.ssh_public_key}"
}

data "oci_containerengine_cluster_option" "test_cluster_option" {
  cluster_option_id = "all"
}

data "oci_containerengine_node_pool_option" "test_node_pool_option" {
  node_pool_option_id = "all"
}
