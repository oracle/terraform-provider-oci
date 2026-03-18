resource "oci_distributed_database_distributed_autonomous_database" "dadb" {
  compartment_id       = var.compartment_id
  display_name         = var.display_name
  db_deployment_type   = var.db_deployment_type
  db_workload          = var.db_workload
  character_set        = var.character_set
  ncharacter_set       = var.ncharacter_set
  database_version     = var.database_version
  chunks               = var.chunks
  prefix               = var.prefix
  listener_port        = var.listener_port
  listener_port_tls    = var.listener_port_tls
  ons_port_local       = var.ons_port_local
  ons_port_remote      = var.ons_port_remote
  sharding_method      = var.sharding_method
  private_endpoint_ids = var.private_endpoint_ids

  catalog_details {
    admin_password                 = var.catalog_admin_password
    source                         = "ADB_D"
    cloud_autonomous_vm_cluster_id = var.catalog_cloud_autonomous_vm_cluster_id
    compute_count                  = var.catalog_compute_count
    data_storage_size_in_gbs       = var.catalog_data_storage_size_in_gbs
    is_auto_scaling_enabled        = var.catalog_is_auto_scaling_enabled
    kms_key_id                     = var.kms_key_id
    vault_id                       = var.vault_id
  }

  shard_details {
    admin_password                 = var.shard_admin_password
    source                         = "ADB_D"
    cloud_autonomous_vm_cluster_id = var.shard_cloud_autonomous_vm_cluster_id1
    compute_count                  = var.shard_compute_count
    data_storage_size_in_gbs       = var.shard_data_storage_size_in_gbs
    is_auto_scaling_enabled        = var.shard_is_auto_scaling_enabled
    kms_key_id                     = var.kms_key_id
    vault_id                       = var.vault_id
  }

  shard_details {
    admin_password                 = var.shard_admin_password
    source                         = "ADB_D"
    cloud_autonomous_vm_cluster_id = var.shard_cloud_autonomous_vm_cluster_id2
    compute_count                  = var.shard_compute_count
    data_storage_size_in_gbs       = var.shard_data_storage_size_in_gbs
    is_auto_scaling_enabled        = var.shard_is_auto_scaling_enabled
    kms_key_id                     = var.kms_key_id
    vault_id                       = var.vault_id
  }

  shard_details {
    admin_password                 = var.shard_admin_password
    source                         = "ADB_D"
    cloud_autonomous_vm_cluster_id = var.shard_cloud_autonomous_vm_cluster_id3
    compute_count                  = var.shard_compute_count
    data_storage_size_in_gbs       = var.shard_data_storage_size_in_gbs
    is_auto_scaling_enabled        = var.shard_is_auto_scaling_enabled
    kms_key_id                     = var.kms_key_id
    vault_id                       = var.vault_id
  }

  /*db_backup_config {
    recovery_window_in_days = 10

    backup_destination_details {
      type = "OBJECT_STORE"
    }
  }*/

  replication_factor = var.replication_factor
  replication_method = var.replication_method
  replication_unit   = var.replication_unit

  state         = var.state
  freeform_tags = var.freeform_tags

  configure_sharding_trigger                                    = var.configure_sharding_trigger
  configure_sharding_is_rebalance_required                      = var.configure_sharding_is_rebalance_required
  download_gsm_certificate_signing_request_trigger              = var.download_gsm_certificate_signing_request_trigger
  generate_gsm_certificate_signing_request_trigger              = var.generate_gsm_certificate_signing_request_trigger
  generate_gsm_certificate_signing_request_trigger_ca_bundle_id = var.generate_gsm_certificate_signing_request_trigger_ca_bundle_id
  start_database_trigger                                        = var.start_database_trigger
  stop_database_trigger                                         = var.stop_database_trigger
  validate_network_trigger                                      = var.validate_network_trigger
  generate_wallet_trigger                                       = var.generate_wallet_trigger
  generate_wallet_password                                      = var.generate_wallet_password
  change_db_backup_config_trigger                               = var.change_db_backup_config_trigger
  upload_signed_certificate_and_generate_wallet_trigger         = var.upload_signed_certificate_and_generate_wallet_trigger
  upload_ca_signed_certificate                                  = var.upload_ca_signed_certificate
}
