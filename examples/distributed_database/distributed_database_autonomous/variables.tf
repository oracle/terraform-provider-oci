variable "tenancy_ocid" {
  type        = string
  description = "OCID of the OCI tenancy."
}

variable "user_ocid" {
  type        = string
  description = "OCID of the OCI user used to authenticate Terraform."
}

variable "fingerprint" {
  type        = string
  description = "Fingerprint of the OCI API signing key associated with the user."
}

variable "private_key_path" {
  type        = string
  description = "Path to the private API signing key file used for OCI authentication."
}

variable "private_key_password" {
  type        = string
  description = "Optional passphrase for the private API signing key. Set to null if the key is not encrypted."
  default     = null
}

variable "region" {
  type        = string
  description = "OCI region identifier (for example: eu-frankfurt-1, us-ashburn-1)."
}

variable "compartment_id" {
  description = "OCID of the compartment where resources will be created"
  type        = string
}

variable "private_endpoint_display_name" {
  description = "Display name for the Distributed Database Private Endpoint"
  type        = string
  default     = "tf-ddb-private-endpoint-adb"
}

# --- Distributed Autonomous Database (ADB_D) ---

variable "display_name" {
  description = "Display name for the distributed autonomous database"
  type        = string
}

variable "db_deployment_type" {
  description = "Deployment type discriminator"
  type        = string
  default     = "ADB_D"
}

variable "db_workload" {
  description = "Database workload type"
  type        = string
  default     = "OLTP"
}

variable "character_set" {
  description = "Database character set"
  type        = string
  default     = "AL32UTF8"
}

variable "ncharacter_set" {
  description = "Database national character set"
  type        = string
  default     = "AL16UTF16"
}

variable "database_version" {
  description = "Database version"
  type        = string
  default     = "26ai"
}

variable "chunks" {
  description = "Chunks"
  type        = number
  default     = 120
}

variable "listener_port" {
  description = "Listener port"
  type        = number
  default     = 11241
}

variable "listener_port_tls" {
  description = "Listener TLS port"
  type        = number
  default     = 11244
}

variable "ons_port_local" {
  description = "ONS local port"
  type        = number
  default     = 11242
}

variable "ons_port_remote" {
  description = "ONS remote port"
  type        = number
  default     = 11243
}

variable "prefix" {
  description = "Prefix for the distributed autonomous database"
  type        = string
}

variable "sharding_method" {
  description = "Sharding method"
  type        = string
  default     = "SYSTEM"
}

# Catalog details (ADB_D)
variable "catalog_admin_password" {
  description = "Catalog admin password (input-only)"
  type        = string
}

variable "catalog_cloud_autonomous_vm_cluster_id" {
  description = "OCID of the cloud autonomous VM cluster used by the catalog"
  type        = string
}

variable "catalog_compute_count" {
  description = "Catalog compute count"
  type        = number
  default     = 4
}

variable "catalog_data_storage_size_in_gbs" {
  description = "Catalog data storage size in GBs"
  type        = number
  default     = 128
}

variable "catalog_is_auto_scaling_enabled" {
  description = "Whether auto scaling is enabled for the catalog"
  type        = bool
  default     = false
}

# Shard details (ADB_D)
variable "shard_admin_password" {
  description = "Shard admin password (input-only)"
  type        = string
}

variable "shard_cloud_autonomous_vm_cluster_id1" {
  description = "OCID of the cloud autonomous VM cluster used by shard 1"
  type        = string
}

variable "shard_cloud_autonomous_vm_cluster_id2" {
  description = "OCID of the cloud autonomous VM cluster used by shard 2"
  type        = string
}

variable "shard_cloud_autonomous_vm_cluster_id3" {
  description = "OCID of the cloud autonomous VM cluster used by shard 3"
  type        = string
}

variable "shard_compute_count" {
  description = "Shard compute count"
  type        = number
  default     = 4
}

variable "shard_data_storage_size_in_gbs" {
  description = "Shard data storage size in GBs"
  type        = number
  default     = 128
}

variable "shard_is_auto_scaling_enabled" {
  description = "Whether auto scaling is enabled for shards"
  type        = bool
  default     = false
}

variable "kms_key_id" {
  description = "KMS key OCID used for encryption"
  type        = string
  default     = null
}

variable "vault_id" {
  description = "Vault OCID used for encryption"
  type        = string
  default     = null
}

# Optional
variable "freeform_tags" {
  description = "Freeform tags"
  type        = map(string)
  default     = {}
}

variable "private_endpoint_ids" {
  type        = list(string)
  description = "List of private endpoint OCIDs used by the distributed autonomous database."
  default     = []
}

variable "replication_factor" {
  description = "Replication factor"
  type        = number
  default     = 3
}

variable "replication_method" {
  description = "Replication method"
  type        = string
  default     = "RAFT"
}

variable "replication_unit" {
  description = "Replication unit"
  type        = number
  default     = 6
}

variable "state" {
  description = "Desired lifecycle state (ACTIVE or INACTIVE)"
  type        = string
  default     = null
}

# Triggers (optional)
variable "configure_sharding_trigger" {
  type        = number
  description = "Optional trigger to run configure sharding."
  default     = null
}

variable "configure_sharding_is_rebalance_required" {
  description = <<EOT
Indicates whether shard rebalancing should be performed as part of the
Configure Distributed Autonomous Database Sharding action.

This value is passed to the ConfigureSharding action API and has no effect
unless configure_sharding_trigger is incremented.
EOT

  type    = bool
  default = null
}

variable "start_database_trigger" {
  type        = number
  description = "Optional trigger to run start database action."
  default     = null
}

variable "stop_database_trigger" {
  type        = number
  description = "Optional trigger to run stop database action."
  default     = null
}

variable "validate_network_trigger" {
  type        = number
  description = "Optional trigger to run validate network."
  default     = null
}

variable "generate_wallet_trigger" {
  type        = number
  description = "Optional trigger to generate wallet."
  default     = null
}

variable "generate_wallet_password" {
  type        = string
  description = "Optional password for trigger to generate wallet."
  default     = null
}

variable "change_db_backup_config_trigger" {
  type        = number
  description = "Optional trigger to change db backup config."
  default     = null
}

variable "upload_signed_certificate_and_generate_wallet_trigger" {
  type        = number
  description = "Optional trigger to upload signed certificate and generate wallet."
  default     = null
}

variable "upload_ca_signed_certificate" {
  type        = string
  description = "Optional CA-signed certificate payload used by upload signed certificate action."
  default     = null
}

variable "download_gsm_certificate_signing_request_trigger" {
  type        = number
  description = "Optional trigger to download GSM certificate signing request."
  default     = null
}

variable "generate_gsm_certificate_signing_request_trigger" {
  type        = number
  description = "Optional trigger to generate GSM certificate signing request."
  default     = null
}

variable "generate_gsm_certificate_signing_request_trigger_ca_bundle_id" {
  type        = string
  description = "Optional CA bundle OCID used during GSM CSR generation."
  default     = null
}
