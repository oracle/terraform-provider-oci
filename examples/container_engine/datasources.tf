// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

data "oci_identity_availability_domain" "ad1" {
  compartment_id = "${var.tenancy_ocid}"
  ad_number      = 1
}

data "oci_identity_availability_domain" "ad2" {
  compartment_id = "${var.tenancy_ocid}"
  ad_number      = 2
}

data "oci_containerengine_cluster_option" "test_cluster_option" {
  cluster_option_id = "all"
}

data "oci_containerengine_node_pool_option" "test_node_pool_option" {
  node_pool_option_id = "all"
}

output "cluster_kubernetes_versions" {
  value = ["${data.oci_containerengine_cluster_option.test_cluster_option.kubernetes_versions}"]
}

output "node_pool_kubernetes_version" {
  value = ["${data.oci_containerengine_node_pool_option.test_node_pool_option.kubernetes_versions}"]
}
