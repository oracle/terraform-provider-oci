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

# variable "subnet_id" {
#   description = "OCID of the subnet used for the Distributed Database Private Endpoint"
#   type        = string
# }

variable "private_endpoint_display_name" {
  description = "Display name for the Distributed Database Private Endpoint"
  type        = string
  default     = "tf-ddb-private-endpoint-exadb"
}

# --- Distributed Database (EXADB_XS) ---

variable "display_name" {
  description = "Display name for the Distributed Database"
  type        = string
}

variable "db_deployment_type" {
  description = "Deployment type discriminator"
  type        = string
  default     = "EXADB_XS"
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
  default     = "19c"
}

variable "chunks" {
  description = "Chunks"
  type        = number
  default     = 120
}

variable "listener_port" {
  description = "Listener port"
  type        = number
  default     = 1521
}

variable "ons_port_local" {
  description = "ONS local port"
  type        = number
  default     = 6200
}

variable "ons_port_remote" {
  description = "ONS remote port"
  type        = number
  default     = 6200
}

variable "prefix" {
  description = "Prefix for the distributed database"
  type        = string
}

# Catalog details (EXADB_XS)
variable "catalog_admin_password" {
  description = "Catalog admin password (input-only)"
  type        = string
}

variable "catalog_vm_cluster_id" {
  description = "OCID of the Exadata VM Cluster used by the catalog"
  type        = string
}

# Shard details (EXADB_XS)
variable "shard_admin_password" {
  description = "Shard admin password (input-only)"
  type        = string
}

variable "shard_vm_cluster_id1" {
  description = "OCID of the Exadata VM Cluster used by the shard"
  type        = string
}
variable "shard_vm_cluster_id2" {
  description = "OCID of the Exadata VM Cluster used by the shard"
  type        = string
}
variable "shard_vm_cluster_id3" {
  description = "OCID of the Exadata VM Cluster used by the shard"
  type        = string
}

# Optional
variable "freeform_tags" {
  description = "Freeform tags"
  type        = map(string)
  default     = {}
}

variable "private_endpoint_ids" {
  type        = list(string)
  description = "List of private endpoint OCIDs used by the distributed autonomous database. In this example it will include the PE created in this stack."
  default     = []
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
Configure Distributed Database Sharding action.

This value is passed to the ConfigureSharding action API and has no effect
unless configure_sharding_trigger is incremented.
EOT

  type    = bool
  default = null
}

variable "start_database_trigger" {
  type        = number
  description = "Optional trigger to run configure sharding."
  default     = null
}
variable "stop_database_trigger" {
  type        = number
  description = "Optional trigger to run configure sharding."
  default     = null
}

variable "validate_network_trigger" {
  type        = number
  description = "Optional trigger to run validate network."
  default     = null
}

variable "validate_network_details" {
  description = "Optional parameters for ValidateDistributedDatabaseNetwork action"
  type = object({
    is_surrogate  = bool
    resource_name = string
    shard_group   = string
  })
  default = null
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

variable "generate_gsm_certificate_signing_request_trigger" {
  type        = number
  description = "Optional trigger to generate GSM certificate signing request."
  default     = null
}

variable "download_gsm_certificate_signing_request_trigger" {
  type        = number
  description = "Optional trigger to download GSM certificate signing request."
  default     = null
}

variable "sharding_method" {
  default = "SYSTEM"
}

variable "kms_key_id" {
  description = "KMS Key Id"
}

variable "vault_id" {
  description = "vault id"
}
