locals {
  create_private_endpoint = try(var.create_private_endpoint, false)
  private_endpoint_ids = local.create_private_endpoint ? concat(
    [module.private_endpoint[0].id],
    try(var.additional_private_endpoint_ids, [])
  ) : try(var.additional_private_endpoint_ids, [])
}

module "private_endpoint" {
  count  = local.create_private_endpoint ? 1 : 0
  source = "./modules/private_endpoint"

  compartment_id = var.compartment_id
  display_name   = var.private_endpoint.display_name
  subnet_id      = var.private_endpoint.subnet_id
  vcn_id         = var.private_endpoint.vcn_id
  description    = try(var.private_endpoint.description, null)
  nsg_ids        = try(var.private_endpoint.nsg_ids, [])
  freeform_tags  = try(var.private_endpoint.freeform_tags, {})
}

module "distributed_database" {
  source = "./modules/distributed_database"

  operational_state                        = var.operational_state
  compartment_id                           = var.compartment_id
  display_name                             = var.display_name
  database_version                         = var.database_version
  character_set                            = var.character_set
  ncharacter_set                           = var.ncharacter_set
  prefix                                   = var.prefix
  listener_port                            = var.listener_port
  listener_port_tls                        = var.listener_port_tls
  ons_port_local                           = var.ons_port_local
  ons_port_remote                          = var.ons_port_remote
  scan_listener_port                       = var.scan_listener_port
  freeform_tags                            = var.freeform_tags
  system_chunk_count                       = var.system_chunk_count
  system_raft_replication_unit_count       = var.system_raft_replication_unit_count
  db_backup_config                         = var.db_backup_config
  catalog_details                          = var.catalog_details
  system_raft_clusters                     = var.system_raft_clusters
  validate_network_trigger                 = var.validate_network_trigger
  change_db_backup_config_trigger          = var.change_db_backup_config_trigger
  configure_sharding_enabled               = var.configure_sharding_enabled
  configure_sharding_is_rebalance_required = var.configure_sharding_is_rebalance_required
  private_endpoint_ids                     = local.private_endpoint_ids

  depends_on = [module.private_endpoint]
}
