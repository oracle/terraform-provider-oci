
variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "private_key_password" {
  default = null
}

variable "region" {
}

variable "compartment_ocid" {
}

variable "subnet_ocid" {
}

variable "nsg_ocids" {
  type    = list(string)
  default = null
}

variable "ogdd_sharded_database_catalog_details_admin_password" {
}

variable "ogdd_sharded_database_catalog_details_cloud_autonomous_vm_cluster_id" {
}

variable "ogdd_sharded_database_catalog_details_compute_count" {
  default = 2
}

variable "ogdd_sharded_database_catalog_details_data_storage_size_in_gbs" {
  default = 32
}

variable "ogdd_sharded_database_catalog_details_is_auto_scaling_enabled" {
  type    = bool
  default = false
}

variable "ogdd_sharded_database_catalog_details_peer_cloud_autonomous_vm_cluster_id" {
  default = null
}

variable "ogdd_sharded_database_character_set" {
  default = "AL32UTF8"
}

variable "ogdd_sharded_database_db_deployment_type" {
  default = "DEDICATED"
}

variable "ogdd_sharded_database_db_version" {
  default = "19c"
}

variable "ogdd_sharded_database_db_workload" {
  default = "OLTP"
}

variable "ogdd_sharded_database_listener_port" {
  type        = number
  default     = 37381
  description = "Needs to be updated/incremented for each new service provisioning on the same environment, even if the previous services ware destroyed."
}

variable "listener_port_tls" {
  type        = number
  default     = 37382
  description = "Needs to be updated/incremented for each new service provisioning on the same environment, even if the previous services ware destroyed."
}

variable "ogdd_sharded_database_ncharacter_set" {
  default = "AL16UTF16"
}

variable "ogdd_sharded_database_ons_port_local" {
  type        = number
  default     = 37384
  description = "Needs to be updated/incremented for each new service provisioning on the same environment, even if the previous services ware destroyed."
}

variable "ogdd_sharded_database_ons_port_remote" {
  type        = number
  default     = 37385
  description = "Needs to be updated/incremented for each new service provisioning on the same environment, even if the previous services ware destroyed."
}

variable "ogdd_sharded_database_prefix" {
  default     = "p1"
  description = "Needs to be updated to a unqiue value for each new service provisioning on the same environment, even if the previous services ware destroyed."
}


variable "ogdd_sharded_database_shard_details_admin_password" {
}

variable "ogdd_sharded_database_shard_details_cloud_autonomous_vm_cluster_id" {
}

variable "ogdd_sharded_database_shard_details_compute_count" {
  default = 2
}

variable "ogdd_sharded_database_shard_details_data_storage_size_in_gbs" {
  default = 32
}

variable "ogdd_sharded_database_shard_details_is_auto_scaling_enabled" {
  type    = bool
  default = false
}

variable "ogdd_sharded_database_shard_details_peer_cloud_autonomous_vm_cluster_id" {
  default = null
}

variable "ogdd_sharded_database_shard_details_shard_space" {
  default = "my-shard-space"
}

variable "ogdd_sharded_database_sharding_method" {
  default = "SYSTEM"
}

variable "ogdd_sharded_database_chunks" {
  default = 120
}

variable "ogdd_sharded_database_cluster_certificate_common_name" {
  default = "gdad_test"
}

variable "autonomous_data_warehouse_db_workload" {
  default = "DW"
}

variable "oci_globally_distributed_database_defined_tags_value" {
  default = {
    "foo-namespace.bar-key" = "value"
  }
}

variable "oci_globally_distributed_database_freeform_tags" {
  default = {
    "Department" = "Finance"
  }
}

variable "oci_globally_distributed_database_replication_method" {
  default = "RAFT"
}

variable "oci_globally_distributed_database_replication_factor" {
  default = 3
}

variable "oci_globally_distributed_database_replication_unit" {
  default = 6
}

variable "vault_id" {
  default = null
}

variable "kms_key_id" {
  default = null
}

variable "kms_key_version_id" {
  default = null
}


variable "oci_globally_distributed_database_configure_gsms_trigger" {
  description = "POST operations will only be performed if the corresponding trigger is not NULL, is != 0 and is greater than the previous value."
  default     = null
}

variable "oci_globally_distributed_database_configure_gsms_trigger_old_gsm_names" {
  type        = list(string)
  description = "configure gsms trigger old gsm names"
  default     = null
}

variable "oci_globally_distributed_database_configure_gsms_trigger_is_latest_gsm_image" {
  type        = bool
  description = "is latest gsm image"
  default     = null
}

variable "oci_globally_distributed_database_configure_sharding_trigger" {
  description = "POST operations will only be performed if the corresponding trigger is not NULL, is != 0 and is greater than the previous value."
  default     = null
}

variable "oci_globally_distributed_database_download_gsm_certificate_signing_request_trigger" {
  description = "POST operations will only be performed if the corresponding trigger is not NULL, is != 0 and is greater than the previous value."
  default     = null
}

variable "oci_globally_distributed_database_generate_gsm_certificate_signing_request_trigger" {
  description = "POST operations will only be performed if the corresponding trigger is not NULL, is != 0 and is greater than the previous value."
  default     = null
}

variable "oci_globally_distributed_database_upload_signed_certificate_and_generate_wallet" {
  description = "POST operations will only be performed if the corresponding trigger is not NULL, is != 0 and is greater than the previous value."
  default     = null
}

variable "oci_globally_distributed_database_generate_wallet" {
  description = "POST operations will only be performed if the corresponding trigger is not NULL, is != 0 and is greater than the previous value."
  default     = null
}

variable "oci_globally_distributed_database_generate_wallet_password" {
  description = "Wallet password."
  default     = null
}

variable "oci_globally_distributed_database_ca_signed_certificate" {
  description = "ca signed certificate"
  default     = null
}


variable "oci_globally_distributed_database_get_connection_string_trigger" {
  description = "POST operations will only be performed if the corresponding trigger is not NULL, is != 0 and is greater than the previous value."
  default     = null
}

variable "oci_globally_distributed_database_start_database_trigger" {
  description = "POST operations will only be performed if the corresponding trigger is not NULL, is != 0 and is greater than the previous value."
  default     = null
}

variable "oci_globally_distributed_database_stop_database_trigger" {
  description = "POST operations will only be performed if the corresponding trigger is not NULL, is != 0 and is greater than the previous value."
  default     = null
}

variable "oci_globally_distributed_database_validate_network_trigger" {
  description = "POST operations will only be performed if the corresponding trigger is not NULL, is != 0 and is greater than the previous value."
  default     = null
}

variable "oci_globally_distributed_database_patch_operation_operation" {
  description = "(Required) (Updatable) The operation can be one of these values: INSERT, MERGE, REMOVE"
  default     = null
}

variable "oci_globally_distributed_database_patch_operation_selection" {
  description = "(Required) (Updatable)"
  default     = null
}

variable "oci_globally_distributed_database_patch_operation_value" {
  description = "(Required when operation=INSERT | MERGE) (Updatable)"
  type        = string
  default     = null
}