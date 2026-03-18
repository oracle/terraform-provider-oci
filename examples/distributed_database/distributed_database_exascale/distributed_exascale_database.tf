resource "oci_distributed_database_distributed_database" "ddb" {
  compartment_id = var.compartment_id
  display_name   = var.display_name
  db_deployment_type = var.db_deployment_type
  character_set      = var.character_set
  ncharacter_set     = var.ncharacter_set
  database_version   = var.database_version
  chunks          = var.chunks
  prefix          = var.prefix
  listener_port   = var.listener_port
  ons_port_local  = var.ons_port_local
  ons_port_remote = var.ons_port_remote
  sharding_method = var.sharding_method
  private_endpoint_ids = var.private_endpoint_ids
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
  freeform_tags = var.freeform_tags

  # patch_operations {
  #   operation = "INSERT"
  #   selection = "shardDetails"

  #   value = jsonencode({
  #     adminPassword        = "WElcomeHome1234##"
  #     computeCount         = 2
  #     source               = "EXADB_XS"
  #     dataStorageSizeInGbs = 128
  #     isAutoScalingEnabled = "false"
  #     vmClusterId          = "ocid1.exadbvmcluster.oc1.eu-frankfurt-1.antheljrymy7j7yabstq4pfy5gnfdjw5wdv4cmpayjjjhmlvee2hsnu2nx3q"
  #     vaultId              = var.vault_id
  #   })
  # }

  configure_sharding_trigger      = var.configure_sharding_trigger
  start_database_trigger          = var.start_database_trigger
  stop_database_trigger           = var.stop_database_trigger
  change_db_backup_config_trigger = var.change_db_backup_config_trigger
}
