# Example: Create an Autonomous Container Database (ACD) using encryption_location_details, after registerPkcs is called on Cloud Autonomous VM Cluster.
#
# Assumes register_pkcs.tf example has been run and the VM Cluster is managed as resource "oci_database_cloud_autonomous_vm_cluster.external_pkcs".
# You may adjust encryption_location_details block fields to match your key provider (example below uses AWS).

# variable "cloud_autonomous_vm_cluster_id" {
#   description = "OCID of Cluster (should match external_pkcs resource in register_pkcs.tf)"
#   type        = string
# }
#

variable "cloud_autonomous_vm_cluster_id" {
  description = "The OCID of the Cloud Autonomous VM Cluster."
  type        = string
}
variable "autonomous_container_database_display_name" {
  description = "Display name for the Autonomous Container Database"
  type        = string
  default     = "example-acd"
}

variable "db_version" {
  description = "Database version for Autonomous Container Database"
  type        = string
  default     = "19.29.0.1.0"
}

variable "db_name" {
  description = "Database name for the Autonomous Container Database"
  type        = string
  default     = "ACDTEST"
}

# Encryption location variables—adjust for your use case (example: AWS)
variable "encryption_location_type" {
  description = "Type of encryption location ('AWS', 'ORACLE_KEY_VAULT', etc.)"
  type        = string
  default     = "AWS"
}

variable "aws_kms_key_arn" {
  description = "AWS KMS Key ARN for TDE (required if using AWS as key store)"
  type        = string
  default     = ""
}

resource "oci_database_autonomous_container_database" "test_acd_encryption_location" {
   cloud_autonomous_vm_cluster_id = var.cloud_autonomous_vm_cluster_id
   display_name                   = var.autonomous_container_database_display_name
   db_version                     = var.db_version
   db_name                        = var.db_name
   patch_model                    = "RELEASE_UPDATES"

  encryption_key_location_details {
    provider_type      = "AWS"     # e.g., "AWS", "ORACLE_KEY_VAULT", etc.
    aws_encryption_key_id   = "ocid1.keyarn"              # For AWS: supply the KMS key ARN
    # For other providers, add fields as required, e.g.:
    # vault_id, compartment_id, endpoint, username, secret_id, etc.
  }

#   depends_on = [oci_database_cloud_autonomous_vm_cluster.external_pkcs]
}

# USAGE:
# 1. Complete register_pkcs.tf steps; VM Cluster must have PKCS registered.
# 2. Update terraform.tfvars with your cluster OCID, encryption location type, and key details.
# 3. terraform apply to create the ACD using the specified encryption location.
#
# Example terraform.tfvars:
# cloud_autonomous_vm_cluster_id    = ""
# encryption_location_type         = "AWS"
# aws_kms_key_arn                  = ""
