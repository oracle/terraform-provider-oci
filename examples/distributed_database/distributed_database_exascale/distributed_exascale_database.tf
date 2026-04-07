
resource "oci_distributed_database_distributed_database" "ddb" {
  compartment_id = var.compartment_id
  display_name   = var.display_name
  #  distributed_database_id = var.distributed_database_id
  db_deployment_type = var.db_deployment_type
  character_set      = var.character_set
  ncharacter_set     = var.ncharacter_set
  database_version   = var.database_version
  # db_workload        = var.db_workload
  chunks          = 120
  prefix          = var.prefix
  listener_port   = var.listener_port
  ons_port_local  = var.ons_port_local
  ons_port_remote = var.ons_port_remote
  #sharding_method      = var.sharding_method
  sharding_method = var.sharding_method
#  private_endpoint_ids = local.dedb_private_endpoint_ids
  private_endpoint_ids = ["ocid1.osddistributeddbprivateendpoint.oc1.eu-frankfurt-1.amaaaaaaymy7j7yajs4rmo2yc2auzfw2t7wowj5rk6ae6gx6h7vnesx6fknq"]
  catalog_details {
    admin_password = var.catalog_admin_password
    source         = "EXADB_XS"
    vm_cluster_id  = var.catalog_vm_cluster_id
    kms_key_id     = var.kms_key_id
    vault_id       = var.vault_id
  }
  shard_details {
    admin_password = var.shard_admin_password
    source         = "EXADB_XS"
    vm_cluster_id  = var.shard_vm_cluster_id1
    kms_key_id     = var.kms_key_id
    vault_id       = var.vault_id
  }
  shard_details {
    admin_password = var.shard_admin_password
    source         = "EXADB_XS"
    vm_cluster_id  = var.shard_vm_cluster_id2
    kms_key_id     = var.kms_key_id
    vault_id       = var.vault_id
  }
  shard_details {
    admin_password = var.shard_admin_password
    source         = "EXADB_XS"
    vm_cluster_id  = var.shard_vm_cluster_id3
    kms_key_id     = var.kms_key_id
    vault_id       = var.vault_id
  }
 /* shard_details {
    admin_password = var.shard_admin_password
    source         = "EXADB_XS"
    vm_cluster_id  = var.shard_vm_cluster_id1
    kms_key_id     = var.kms_key_id
    vault_id       = var.vault_id
  }*/

  /*db_backup_config {
    auto_backup_window= "SLOT_FIVE"
    auto_full_backup_day= "MONDAY"
    auto_full_backup_window= "SLOT_FIVE"
    backup_deletion_policy= "DELETE_IMMEDIATELY"
    backup_destination_details {
      type= "OBJECT_STORE"
    }

    can_run_immediate_full_backup= true
    is_auto_backup_enabled= true
    recovery_window_in_days= 30
  }*/

  # Optional top-level fields
  /*replication_factor = var.replication_factor
  replication_method = var.replication_method
  replication_unit   = var.replication_unit*/
  replication_factor = 3
  replication_method = "RAFT"
  replication_unit   = 6
  #state              = var.state
  freeform_tags = var.freeform_tags
  #defined_tags  = var.defined_tags
  # Optional triggers
  configure_sharding_trigger      = var.configure_sharding_trigger
  validate_network_trigger        = var.validate_network_trigger
  start_database_trigger          = var.start_database_trigger
  stop_database_trigger           = var.stop_database_trigger
  change_db_backup_config_trigger = var.change_db_backup_config_trigger
  db_backup_config {
    recovery_window_in_days = 45
    is_auto_backup_enabled  = true
    backup_destination_details {
      type = "OBJECT_STORE"
    }
  }

# patch_operations {
#   operation = "INSERT"
#   selection = "shardDetails"

#   value = jsonencode({
#     source        = "EXADB_XS"
#     adminPassword = var.shard_admin_password
#     vmClusterId   = var.shard_vm_cluster_id3

#     # Optional for EXADB_XS
#     # shardSpace      = var.shard_space
#     # vaultId         = var.vault_id
#     # kmsKeyId        = var.kms_key_id
#     # kmsKeyVersionId = var.kms_key_version_id
#     # peerVmClusterIds = [var.peer_vm_cluster_id]
#   })
# }

# patch_operations {
#   operation = "REMOVE"
#   selection = "shardDetails[?name == 'mt200004x']"
#
#   # Provider quirk: schema still requires value even for REMOVE,
#   # but the REMOVE code path ignores it.
#   value = jsonencode({})
# }



  upload_signed_certificate_and_generate_wallet_trigger = var.upload_signed_certificate_and_generate_wallet_trigger
  generate_gsm_certificate_signing_request_trigger      = var.generate_gsm_certificate_signing_request_trigger
  download_gsm_certificate_signing_request_trigger      = var.download_gsm_certificate_signing_request_trigger

  configure_sharding_is_rebalance_required = var.configure_sharding_is_rebalance_required

  # dynamic "validate_network_details" {
  #   for_each = (
  #     length(keys(local.vnd)) == 0
  #     ? []
  #     : [local.vnd]
  #   )

  #   content {
  #     # Use empty defaults instead of null for TF 0.12 friendliness
  #     is_surrogate  = lookup(validate_network_details.value, "is_surrogate", false)
  #     resource_name = lookup(validate_network_details.value, "resource_name", "")
  #     shard_group   = lookup(validate_network_details.value, "shard_group", "")
  #   }
  # }
}
