variable "operational_state" {
  description = "Operational state."
  type        = string
  default     = null
}

variable "compartment_id" {
  description = "Compartment id."
  type        = string
  nullable    = false
}

variable "display_name" {
  description = "Display name."
  type        = string
  nullable    = false
}

variable "database_version" {
  description = "Database version."
  type        = string
  default     = "26ai"
  nullable    = false
}

variable "character_set" {
  description = "Character set."
  type        = string
  default     = "AL32UTF8"
  nullable    = false
}

variable "ncharacter_set" {
  description = "Ncharacter set."
  type        = string
  default     = "AL16UTF16"
  nullable    = false
}

variable "prefix" {
  description = "Prefix."
  type        = string
  nullable    = false
}

variable "listener_port" {
  description = "Listener port."
  type        = number
  nullable    = false
}

variable "listener_port_tls" {
  description = "Listener port tls."
  type        = number
  default     = null
}

variable "ons_port_local" {
  description = "Ons port local."
  type        = number
  nullable    = false
}

variable "ons_port_remote" {
  description = "Ons port remote."
  type        = number
  nullable    = false
}

variable "scan_listener_port" {
  description = "Scan listener port."
  type        = number
  default     = null
}

variable "freeform_tags" {
  description = "Freeform tags."
  type        = map(string)
  default     = {}
  nullable    = false
}

variable "system_chunk_count" {
  description = "System chunk count."
  type        = number
  default     = null
}

variable "system_raft_replication_unit_count" {
  description = "System raft replication unit count."
  type        = number
  default     = null
}

variable "db_backup_config" {
  description = "Db backup config."
  type = object({
    is_auto_backup_enabled        = optional(bool)
    auto_backup_window            = optional(string)
    auto_full_backup_day          = optional(string)
    auto_full_backup_window       = optional(string)
    backup_deletion_policy        = optional(string)
    recovery_window_in_days       = optional(number)
    can_run_immediate_full_backup = optional(bool)
    is_remote_backup_enabled      = optional(bool)
    remote_region                 = optional(string)
    backup_destination_details = optional(list(object({
      dbrs_policy_id = optional(string)
      id             = optional(string)
      internet_proxy = optional(string)
      is_remote      = optional(bool)
      remote_region  = optional(string)
      type           = string
    })), [])
  })
  default = null
}

variable "catalog_details" {
  description = "Catalog details."
  type = list(object({
    original_replica = object({
      name = string
      database = object({
        admin_password                       = optional(string)
        admin_password_secret_id             = optional(string)
        admin_password_secret_version_number = optional(number)
        source                               = string
        vm_cluster_id                        = optional(string)
        db_home_id                           = optional(string)
        vault_id                             = optional(string)
        kms_key_id                           = optional(string)
        protection_mode                      = optional(string)
        transport_type                       = optional(string)
      })
    })
    data_guard_replicas = optional(list(object({
      name = string
      database = object({
        admin_password                       = optional(string)
        admin_password_secret_id             = optional(string)
        admin_password_secret_version_number = optional(number)
        source                               = string
        vm_cluster_id                        = optional(string)
        db_home_id                           = optional(string)
        vault_id                             = optional(string)
        kms_key_id                           = optional(string)
        protection_mode                      = optional(string)
        transport_type                       = optional(string)
      })
    })), [])
  }))
  nullable = false
}

variable "composite_data_guard_shard_spaces" {
  description = "Composite data guard shard spaces."
  type = list(object({
    name        = string
    chunk_count = number
    original_replica = object({
      name = string
      databases = list(object({
        admin_password                       = optional(string)
        admin_password_secret_id             = optional(string)
        admin_password_secret_version_number = optional(number)
        source                               = string
        vm_cluster_id                        = optional(string)
        db_home_id                           = optional(string)
        vault_id                             = optional(string)
        kms_key_id                           = optional(string)
        protection_mode                      = optional(string)
        transport_type                       = optional(string)
      }))
    })
    data_guard_replicas = optional(list(object({
      name = string
      databases = list(object({
        admin_password                       = optional(string)
        admin_password_secret_id             = optional(string)
        admin_password_secret_version_number = optional(number)
        source                               = string
        vm_cluster_id                        = optional(string)
        db_home_id                           = optional(string)
        vault_id                             = optional(string)
        kms_key_id                           = optional(string)
        protection_mode                      = optional(string)
        transport_type                       = optional(string)
      }))
    })), [])
  }))
  default  = []
  nullable = false
}

variable "validate_network_trigger" {
  description = "Validate network trigger."
  type        = number
  default     = null
}

variable "configure_sharding_enabled" {
  description = "Configure sharding enabled."
  type        = bool
  default     = true
  nullable    = false
}

variable "configure_sharding_is_rebalance_required" {
  description = "Configure sharding is rebalance required."
  type        = bool
  default     = true
  nullable    = false
}

variable "create_private_endpoint" {
  description = "Create private endpoint."
  type        = bool
  default     = false
  nullable    = false
}

variable "additional_private_endpoint_ids" {
  description = "Additional private endpoint ids."
  type        = list(string)
  default     = []
  nullable    = false
}

variable "private_endpoint" {
  description = "Private endpoint."
  type = object({
    display_name  = string
    subnet_id     = string
    vcn_id        = string
    description   = optional(string)
    nsg_ids       = optional(list(string), [])
    freeform_tags = optional(map(string), {})
  })
  default = null
}

variable "gds_control_nodes" {
  description = "Gds control nodes."
  type = list(object({
    public_ssh_key = string
    subnet_id      = optional(string)
  }))
  default  = []
  nullable = false
}

variable "actions" {
  description = "Actions."
  type = object({
    insert = optional(list(object({
      selection = string
      value     = string
    })), [])
    insert_dg = optional(list(object({
      selection    = string
      replica_name = string
      database = object({
        admin_password                       = optional(string)
        admin_password_secret_id             = optional(string)
        admin_password_secret_version_number = optional(number)
        availability_domain                  = optional(string)
        db_home_id                           = optional(string)
        exascale_db_storage_vault_id         = optional(string)
        protection_mode                      = optional(string)
        source                               = string
        transport_type                       = optional(string)
        vm_cluster_id                        = optional(string)
        db_storage_vault_details = optional(object({
          additional_flash_cache_in_percent = optional(number)
          autoscale_limit_in_gbs            = optional(number)
          compartment_id                    = optional(string)
          high_capacity_database_storage    = optional(number)
          is_autoscale_enabled              = optional(bool)
          subscription_id                   = optional(string)
        }))
        vm_cluster_details = optional(object({
          backup_network_nsg_ids        = optional(list(string))
          backup_subnet_id              = optional(string)
          compartment_id                = optional(string)
          domain                        = optional(string)
          enabled_ecpu_count            = optional(number)
          is_diagnostics_events_enabled = optional(bool)
          is_health_monitoring_enabled  = optional(bool)
          is_incident_logs_enabled      = optional(bool)
          license_model                 = optional(string)
          nsg_ids                       = optional(list(string))
          private_zone_id               = optional(string)
          ssh_public_keys               = optional(list(string))
          subnet_id                     = optional(string)
          subscription_id               = optional(string)
          total_ecpu_count              = optional(number)
          vm_file_system_storage_size   = optional(number)
        }))
      })
    })), [])
    remove_dg = optional(list(object({
      selection    = string
      replica_name = string
    })), [])
    remove = optional(list(object({
      selection = string
    })), [])
  })
  default  = {}
  nullable = false
}

variable "auth" {
  description = "Auth."
  type        = string
  default     = "SecurityToken"
  nullable    = false
}

variable "config_file_profile" {
  description = "Config file profile."
  type        = string
  default     = "DEFAULT"
  nullable    = false
}

variable "region" {
  description = "Region."
  type        = string
  nullable    = false
}
