variable "region" {}

variable "config_file_profile" {}

variable "compartment_ocid" {}

variable "source_connection_oracle_id" {}

variable "source_connection_container_oracle_id" {}

variable "target_connection_oracle_id" {}

provider "oci" {
  auth                = "SecurityToken"
  config_file_profile = var.config_file_profile
  region              = var.region
}

resource "oci_database_migration_migration" "backcompat" {
  compartment_id                          = var.compartment_ocid
  database_combination                    = "ORACLE"
  source_database_connection_id           = var.source_connection_oracle_id
  source_container_database_connection_id = var.source_connection_container_oracle_id
  target_database_connection_id           = var.target_connection_oracle_id
  type                                    = "OFFLINE"
  display_name                            = "TF_backcompat_migration"

  data_transfer_medium_details {
    type = "OBJECT_STORAGE"

    object_storage_bucket {
      bucket    = "terraform-provider-acceptance-test"
      namespace = "terraform-provider-acceptance-test"
    }
  }

  initial_load_settings {
    job_mode = "SCHEMA"

    export_directory_object {
      name = "DATA_PUMP_DIR"
      path = "/u01/app/oracle/dumpdir"
    }
  }
}
