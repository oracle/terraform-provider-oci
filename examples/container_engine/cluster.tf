// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

resource "oci_containerengine_cluster" "test_cluster" {
  #Required
  compartment_id     = "${var.compartment_ocid}"
  kubernetes_version = "${data.oci_containerengine_cluster_option.test_cluster_option.kubernetes_versions.0}"
  name               = "${var.cluster_name}"
  vcn_id             = "${oci_core_virtual_network.test_vcn.id}"

  #Optional
  options {
    service_lb_subnet_ids = ["${oci_core_subnet.clusterSubnet_1.id}", "${oci_core_subnet.clusterSubnet_2.id}"]

    #Optional
    add_ons {
      #Optional
      is_kubernetes_dashboard_enabled = "${var.cluster_options_add_ons_is_kubernetes_dashboard_enabled}"
      is_tiller_enabled               = "${var.cluster_options_add_ons_is_tiller_enabled}"
    }

    kubernetes_network_config {
      #Optional
      pods_cidr     = "${var.cluster_options_kubernetes_network_config_pods_cidr}"
      services_cidr = "${var.cluster_options_kubernetes_network_config_services_cidr}"
    }
  }
}

resource "oci_containerengine_node_pool" "test_node_pool" {
  #Required
  cluster_id         = "${oci_containerengine_cluster.test_cluster.id}"
  compartment_id     = "${var.compartment_ocid}"
  kubernetes_version = "${data.oci_containerengine_node_pool_option.test_node_pool_option.kubernetes_versions.0}"
  name               = "${var.node_pool_name}"
  node_image_name    = "${var.node_pool_node_image_name}"
  node_shape         = "${var.node_pool_node_shape}"
  subnet_ids         = ["${oci_core_subnet.nodePool_Subnet_1.id}", "${oci_core_subnet.nodePool_Subnet_2.id}"]

  #Optional
  initial_node_labels {
    #Optional
    key   = "${var.node_pool_initial_node_labels_key}"
    value = "${var.node_pool_initial_node_labels_value}"
  }

  quantity_per_subnet = "${var.node_pool_quantity_per_subnet}"
  ssh_public_key      = "${var.node_pool_ssh_public_key}"
}

output "cluster" {
  value = {
    id                 = "${oci_containerengine_cluster.test_cluster.id}"
    kubernetes_version = "${oci_containerengine_cluster.test_cluster.kubernetes_version}"
    name               = "${oci_containerengine_cluster.test_cluster.name}"
  }
}

output "node_pool" {
  value = {
    id                 = "${oci_containerengine_node_pool.test_node_pool.id}"
    kubernetes_version = "${oci_containerengine_node_pool.test_node_pool.kubernetes_version}"
    name               = "${oci_containerengine_node_pool.test_node_pool.name}"
    subnet_ids         = "${oci_containerengine_node_pool.test_node_pool.subnet_ids}"
  }
}
