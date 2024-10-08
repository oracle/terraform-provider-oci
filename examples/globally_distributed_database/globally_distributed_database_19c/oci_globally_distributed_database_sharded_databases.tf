
// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

resource "oci_globally_distributed_database_sharded_database" "this" {
  #Required
  catalog_details {
    #Required
    admin_password                      = var.ogdd_sharded_database_catalog_details_admin_password
    cloud_autonomous_vm_cluster_id      = var.ogdd_sharded_database_catalog_details_cloud_autonomous_vm_cluster_id
    compute_count                       = var.ogdd_sharded_database_catalog_details_compute_count
    data_storage_size_in_gbs            = var.ogdd_sharded_database_catalog_details_data_storage_size_in_gbs
    is_auto_scaling_enabled             = var.ogdd_sharded_database_catalog_details_is_auto_scaling_enabled
    peer_cloud_autonomous_vm_cluster_id = var.ogdd_sharded_database_catalog_details_peer_cloud_autonomous_vm_cluster_id
  }

  character_set      = var.ogdd_sharded_database_character_set
  compartment_id     = var.compartment_ocid
  db_deployment_type = var.ogdd_sharded_database_db_deployment_type
  db_version         = var.ogdd_sharded_database_db_version
  db_workload        = var.ogdd_sharded_database_db_workload
  display_name       = "GloballyDistributedDB-Sharded-Database-Example"
  listener_port      = var.ogdd_sharded_database_listener_port
  listener_port_tls  = var.listener_port_tls
  ncharacter_set     = var.ogdd_sharded_database_ncharacter_set
  ons_port_local     = var.ogdd_sharded_database_ons_port_local
  ons_port_remote    = var.ogdd_sharded_database_ons_port_remote
  prefix             = var.ogdd_sharded_database_prefix

  shard_details {
    #Required
    admin_password                      = var.ogdd_sharded_database_shard_details_admin_password
    cloud_autonomous_vm_cluster_id      = var.ogdd_sharded_database_shard_details_cloud_autonomous_vm_cluster_id
    compute_count                       = var.ogdd_sharded_database_shard_details_compute_count
    data_storage_size_in_gbs            = var.ogdd_sharded_database_shard_details_data_storage_size_in_gbs
    is_auto_scaling_enabled             = var.ogdd_sharded_database_shard_details_is_auto_scaling_enabled
    peer_cloud_autonomous_vm_cluster_id = var.ogdd_sharded_database_shard_details_peer_cloud_autonomous_vm_cluster_id
    shard_space                         = var.ogdd_sharded_database_shard_details_shard_space
  }

  sharding_method                 = var.ogdd_sharded_database_sharding_method
  chunks                          = var.ogdd_sharded_database_chunks
  cluster_certificate_common_name = var.ogdd_sharded_database_cluster_certificate_common_name
  #defined_tags                    = var.oci_globally_distributed_database_defined_tags_value
  #freeform_tags                   = var.oci_globally_distributed_database_freeform_tags

  # POST OPERATIONS
  # POST operations will only be performed if the corresponding trigger is not NULL, is != 0 and is greater than the previous value.

  configure_sharding_trigger                       = var.oci_globally_distributed_database_configure_sharding_trigger
  download_gsm_certificate_signing_request_trigger = var.oci_globally_distributed_database_download_gsm_certificate_signing_request_trigger
  generate_gsm_certificate_signing_request_trigger = var.oci_globally_distributed_database_generate_gsm_certificate_signing_request_trigger
  get_connection_string_trigger                    = var.oci_globally_distributed_database_get_connection_string_trigger
  start_database_trigger                           = var.oci_globally_distributed_database_start_database_trigger
  stop_database_trigger                            = var.oci_globally_distributed_database_stop_database_trigger
  validate_network_trigger                         = var.oci_globally_distributed_database_validate_network_trigger

  # PATCH Operations
  /*
  patch_operations {
    #Required
    operation = var.oci_globally_distributed_database_patch_operation_operation
    selection = var.oci_globally_distributed_database_patch_operation_selection

    #Optional
    value = var.oci_globally_distributed_database_patch_operation_value
  }
  */

  depends_on = [oci_globally_distributed_database_private_endpoint.this]
}

