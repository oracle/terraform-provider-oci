terraform {
  required_providers {
    oci = {
      source  = "oracle/oci"
      version = "8.22.0"
    }
  }
}

variable "compartment_id" {
  type = string
}

variable "datasafe_private_endpoint_id" {
  type = string
}

variable "target_database_user_name" {
  type = string
}

variable "target_database_password" {
  description = "Password to set on the imported target database."
  type        = string
  sensitive   = true
}

resource "oci_data_safe_target_database" "credential_update" {
  compartment_id = var.compartment_id

  # Keep this block unchanged during a credential-only update. In particular,
  # do not add db_system_id or vm_cluster_id when they are not returned by GET.
  database_details {
    database_type       = "DATABASE_CLOUD_SERVICE"
    infrastructure_type = "ORACLE_CLOUD"
  }

  connection_option {
    connection_type              = "PRIVATE_ENDPOINT"
    datasafe_private_endpoint_id = var.datasafe_private_endpoint_id
  }

  credentials {
    user_name = var.target_database_user_name
    password  = var.target_database_password
  }

  display_name = var.target_database_display_name
}

variable "target_database_display_name" {
  type = string
}
