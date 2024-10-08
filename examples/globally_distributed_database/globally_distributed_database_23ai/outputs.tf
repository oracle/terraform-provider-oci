output "oci_globally_distributed_database_private_endpoints" {
  value = {
    oci_globally_distributed_database_private_endpoint = oci_globally_distributed_database_private_endpoint.this
    data_source_private_endpoint                       = data.oci_globally_distributed_database_private_endpoint.this
    data_source_private_endpoints                      = data.oci_globally_distributed_database_private_endpoints.this
  }
}

output "oci_globally_distributed_database_sharded_database_non_secret" {
  value = {
    oci_globally_distributed_database_sharded_database = {
      catalog_details = [
        for catalog_detail in oci_globally_distributed_database_sharded_database.this.catalog_details : {
          //admin_password                      = nonsensitive(catalog_detail.admin_password)
          cloud_autonomous_vm_cluster_id      = catalog_detail.cloud_autonomous_vm_cluster_id
          compute_count                       = catalog_detail.compute_count
          container_database_id               = catalog_detail.container_database_id
          container_database_parent_id        = catalog_detail.container_database_parent_id
          data_storage_size_in_gbs            = catalog_detail.container_database_parent_id
          encryption_key_details              = catalog_detail.encryption_key_details
          is_auto_scaling_enabled             = catalog_detail.is_auto_scaling_enabled
          metadata                            = catalog_detail.metadata
          name                                = catalog_detail.name
          peer_cloud_autonomous_vm_cluster_id = catalog_detail.peer_cloud_autonomous_vm_cluster_id
          shard_group                         = catalog_detail.shard_group
          status                              = catalog_detail.status
          supporting_resource_id              = catalog_detail.supporting_resource_id
          time_created                        = catalog_detail.time_created
          time_ssl_certificate_expires        = catalog_detail.time_ssl_certificate_expires
          time_updated                        = catalog_detail.time_updated
        }
      ]
      character_set                                    = oci_globally_distributed_database_sharded_database.this.character_set
      chunks                                           = oci_globally_distributed_database_sharded_database.this.chunks
      cluster_certificate_common_name                  = oci_globally_distributed_database_sharded_database.this.cluster_certificate_common_name
      compartment_id                                   = oci_globally_distributed_database_sharded_database.this.compartment_id
      configure_sharding_trigger                       = oci_globally_distributed_database_sharded_database.this.configure_sharding_trigger
      connection_strings                               = oci_globally_distributed_database_sharded_database.this.connection_strings
      db_deployment_type                               = oci_globally_distributed_database_sharded_database.this.db_deployment_type
      db_version                                       = oci_globally_distributed_database_sharded_database.this.db_version
      db_workload                                      = oci_globally_distributed_database_sharded_database.this.db_workload
      defined_tags                                     = oci_globally_distributed_database_sharded_database.this.defined_tags
      display_name                                     = oci_globally_distributed_database_sharded_database.this.display_name
      download_gsm_certificate_signing_request_trigger = oci_globally_distributed_database_sharded_database.this.download_gsm_certificate_signing_request_trigger
      freeform_tags                                    = oci_globally_distributed_database_sharded_database.this.freeform_tags
      generate_gsm_certificate_signing_request_trigger = oci_globally_distributed_database_sharded_database.this.generate_gsm_certificate_signing_request_trigger
      //fetched_all_connection_strings                   = oci_globally_distributed_database_sharded_database.this.fetched_all_connection_strings
      gsms                    = oci_globally_distributed_database_sharded_database.this.gsms
      id                      = oci_globally_distributed_database_sharded_database.this.id
      lifecycle_state_details = oci_globally_distributed_database_sharded_database.this.lifecycle_state_details
      listener_port           = oci_globally_distributed_database_sharded_database.this.listener_port
      listener_port_tls       = oci_globally_distributed_database_sharded_database.this.listener_port_tls
      ncharacter_set          = oci_globally_distributed_database_sharded_database.this.ncharacter_set
      ons_port_local          = oci_globally_distributed_database_sharded_database.this.ons_port_local
      ons_port_remote         = oci_globally_distributed_database_sharded_database.this.ons_port_remote
      patch_operations        = oci_globally_distributed_database_sharded_database.this.patch_operations
      prefix                  = oci_globally_distributed_database_sharded_database.this.prefix
      private_endpoint        = oci_globally_distributed_database_sharded_database.this.private_endpoint
      shard_details = [
        for shard_detail in oci_globally_distributed_database_sharded_database.this.shard_details : {
          //admin_password                      = nonsensitive(shard_detail.admin_password)
          cloud_autonomous_vm_cluster_id      = shard_detail.cloud_autonomous_vm_cluster_id
          compute_count                       = shard_detail.compute_count
          container_database_id               = shard_detail.container_database_id
          container_database_parent_id        = shard_detail.container_database_parent_id
          data_storage_size_in_gbs            = shard_detail.data_storage_size_in_gbs
          encryption_key_details              = shard_detail.encryption_key_details
          is_auto_scaling_enabled             = shard_detail.is_auto_scaling_enabled
          metadata                            = shard_detail.metadata
          name                                = shard_detail.name
          peer_cloud_autonomous_vm_cluster_id = shard_detail.peer_cloud_autonomous_vm_cluster_id
          shard_group                         = shard_detail.shard_group
          shard_space                         = shard_detail.shard_space
          status                              = shard_detail.status
          supporting_resource_id              = shard_detail.supporting_resource_id
          time_created                        = shard_detail.time_created
          time_ssl_certificate_expires        = shard_detail.time_ssl_certificate_expires
          time_updated                        = shard_detail.time_updated
        }
      ]
      sharded_database_id           = oci_globally_distributed_database_sharded_database.this.sharded_database_id
      sharding_method               = oci_globally_distributed_database_sharded_database.this.sharding_method
      start_database_trigger        = oci_globally_distributed_database_sharded_database.this.start_database_trigger
      state                         = oci_globally_distributed_database_sharded_database.this.state
      stop_database_trigger         = oci_globally_distributed_database_sharded_database.this.stop_database_trigger
      system_tags                   = oci_globally_distributed_database_sharded_database.this.system_tags
      time_created                  = oci_globally_distributed_database_sharded_database.this.time_created
      time_updated                  = oci_globally_distributed_database_sharded_database.this.time_updated
      time_zone                     = oci_globally_distributed_database_sharded_database.this.time_zone
      timeouts                      = oci_globally_distributed_database_sharded_database.this.timeouts
      get_connection_string_trigger = oci_globally_distributed_database_sharded_database.this.get_connection_string_trigger
      validate_network_trigger      = oci_globally_distributed_database_sharded_database.this.validate_network_trigger
    },
    data_source_sharded_database  = data.oci_globally_distributed_database_sharded_database.this
    data_source_sharded_databases = data.oci_globally_distributed_database_sharded_databases.this
  }
  sensitive = true
}

output "oci_globally_distributed_database_sharded_database" {
  value = {
    oci_globally_distributed_database_sharded_database = oci_globally_distributed_database_sharded_database.this,
    data_source_sharded_database                       = data.oci_globally_distributed_database_sharded_database.this
    data_source_sharded_databases                      = data.oci_globally_distributed_database_sharded_databases.this
  }
  sensitive = true
}