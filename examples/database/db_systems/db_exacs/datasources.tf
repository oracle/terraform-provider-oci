// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

data "oci_identity_availability_domain" "ad" {
  compartment_id = var.tenancy_ocid
  ad_number      = 1
}

data "oci_database_db_system_shapes" "test_db_system_shapes" {
  availability_domain = data.oci_identity_availability_domain.ad.name
  compartment_id      = var.compartment_ocid

  filter {
    name   = "shape"
    values = [var.db_system_shape]
  }
}

data "oci_database_cloud_vm_cluster_iorm_config" "test_cloud_vm_cluster_iorm_config" {
  cloud_vm_cluster_id = oci_database_cloud_vm_cluster_iorm_config.test_cloud_vm_cluster_iorm_config.cloud_vm_cluster_id
}

data "oci_identity_tenancy" "tenancy" {
  tenancy_id = var.tenancy_ocid
}

data "oci_database_db_servers" "test_cloud_db_servers" {
  #Required
  compartment_id            = var.compartment_ocid
  exadata_infrastructure_id = oci_database_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure.id
}

data "oci_database_application_vips" "test_application_vips" {
  #Required
  cloud_vm_cluster_id = oci_database_cloud_vm_cluster.test_cloud_vm_cluster.id
  compartment_id      = var.compartment_ocid

  #Optional
  state               = "AVAILABLE"
}

data "oci_database_db_nodes" "db_nodes" {
  compartment_id = var.compartment_ocid
  vm_cluster_id  = oci_database_cloud_vm_cluster.test_cloud_vm_cluster.id
}