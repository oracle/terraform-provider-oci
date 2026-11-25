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

variable "db_workload_type" {
  description = "Db workload type."
  type        = string
  default     = "OLTP"
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
  nullable    = false
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

variable "catalog_details" {
  description = "Catalog details."
  type = list(object({
    original_replica = object({
      name = string
      database = object({
        admin_password                       = optional(string)
        admin_password_secret_id             = optional(string)
        admin_password_secret_version_number = optional(number)
        cloud_autonomous_vm_cluster_id       = string
        compute_count                        = number
        data_storage_size_in_gbs             = number
        source                               = string
        is_auto_scaling_enabled              = optional(bool)
        vault_id                             = optional(string)
        kms_key_id                           = optional(string)
        protection_mode                      = optional(string)
      })
    })
    data_guard_replicas = optional(list(object({
      name = string
      database = object({
        cloud_autonomous_vm_cluster_id       = string
        compute_count                        = number
        data_storage_size_in_gbs             = number
        admin_password                       = optional(string)
        admin_password_secret_id             = optional(string)
        admin_password_secret_version_number = optional(number)
        source                               = string
        is_auto_scaling_enabled              = optional(bool)
        vault_id                             = optional(string)
        kms_key_id                           = optional(string)
        protection_mode                      = optional(string)
      })
    })), [])
  }))
  nullable = false
}

variable "system_data_guard_databases" {
  description = "System data guard databases."
  type = object({
    original_replica = object({
      name = string
      databases = list(object({
        admin_password                       = optional(string)
        admin_password_secret_id             = optional(string)
        admin_password_secret_version_number = optional(number)
        cloud_autonomous_vm_cluster_id       = string
        compute_count                        = number
        data_storage_size_in_gbs             = number
        source                               = string
        is_auto_scaling_enabled              = optional(bool)
        vault_id                             = optional(string)
        kms_key_id                           = optional(string)
        protection_mode                      = optional(string)
      }))
    })
    data_guard_replicas = optional(list(object({
      name = string
      databases = list(object({
        cloud_autonomous_vm_cluster_id       = string
        compute_count                        = number
        data_storage_size_in_gbs             = number
        admin_password                       = optional(string)
        admin_password_secret_id             = optional(string)
        admin_password_secret_version_number = optional(number)
        source                               = string
        is_auto_scaling_enabled              = optional(bool)
        vault_id                             = optional(string)
        kms_key_id                           = optional(string)
        protection_mode                      = optional(string)
      }))
    })), [])
  })
  nullable = false
}

variable "validate_network_trigger" {
  description = "Validate network trigger."
  type        = number
  default     = null
}

variable "change_db_backup_config_trigger" {
  description = "Change db backup config trigger."
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

variable "private_endpoint_ids" {
  description = "Private endpoint ids."
  type        = list(string)
  nullable    = false
}

variable "certificate_id" {
  description = "Certificate id."
  type        = string
}

variable "ca_bundle_id" {
  description = "Ca bundle id."
  type        = string
}

variable "db_backup_config" {
  description = "Db backup config."
  type = object({
    recovery_window_in_days = optional(number)
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
