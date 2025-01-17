// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "primary_cluster_name" {
  description = "Primary OKE cluster"
  type        = string
  default     = "myCluster"
}

variable "peer_cluster_name" {
  description = "Peer OKE cluster"
  type        = string
  default     = "peerCluster"
}

data "oci_containerengine_clusters" "test_clusters" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  name = var.primary_cluster_name
}

data "oci_containerengine_clusters" "peer_clusters" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  name = var.peer_cluster_name
}
