// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

provider "oci" {
  tenancy_ocid = var.tenancy_ocid
  user_ocid = var.user_ocid
  fingerprint = var.fingerprint
  private_key_path = var.private_key_path
  region = var.region
}

variable "compartment_id" {
  default = "ocid1.test.oc1..<unique_ID>EXAMPLE-compartmentId-Value"
}

variable "managed_database_id" {
   default = "ocid1.test.oc1..<unique_ID>EXAMPLE-managedDatabaseId-Value"
}

variable "managed_database_deployment_type" {
  default = "ONPREMISE"
}

variable "managed_database_management_option" {
  default = "ADVANCED"
}

variable "managed_database_name" {
  default = "EXAMPLE-managedDatabaseName-Value"
}

variable "managed_db_defined_tags_value" {
  default = "managed_db_tag_value"
}

variable "managed_db_freeform_tags" {
  default = { "bar-key" = "value" }
}

# Create a new Tag Namespace.
resource "oci_identity_tag_namespace" "tag_namespace1" {
  #Required
  compartment_id = var.tenancy_ocid
  description    = "example tag namespace"
  name           = "example-tag-namespace-all"
}

# Create a new Tag definition in the above Tag Namespace.
resource "oci_identity_tag" "tag1" {
  #Required
  description      = "example tag"
  name             = "example-tag"
  tag_namespace_id = oci_identity_tag_namespace.tag_namespace1.id
}

data "oci_database_management_managed_databases" "test_managed_databases" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	deployment_type = var.managed_database_deployment_type
	#external_exadata_infrastructure_id = oci_database_management_external_exadata_infrastructure.test_external_exadata_infrastructure.id
	id = var.managed_database_id
	#management_option = var.managed_database_management_option
	#name = var.managed_database_name
}

# Update tags on a Managed Database resource.
resource "oci_database_management_managed_database" "test_managed_database" {
  managed_database_id = var.managed_database_id
  defined_tags  = {
    "${oci_identity_tag_namespace.tag_namespace1.name}.${oci_identity_tag.tag1.name}" = var.managed_db_defined_tags_value
  }
  freeform_tags = var.managed_db_freeform_tags
}

# External CDB
variable "external_cdb_id" {
  default = "ocid1.externalcontainerdatabase"
}
variable "connector_id" {
  default = "ocid1.externaldatabaseconnector"
}

resource "oci_database_management_externalcontainerdatabase_external_container_dbm_features_management" "test_externalcontainerdatabase_external_container_dbm_features_management" {
  feature_details {
    connector_details {
      connector_type = "EXTERNAL"
      database_connector_id = var.connector_id
    }
    feature = "DIAGNOSTICS_AND_MANAGEMENT"
    license_model = "LICENSE_INCLUDED"
  }
  external_container_database_id = var.external_cdb_id
  enable_external_container_dbm_feature = "true"
}

# Uncomment PDB enable APIs only after CDB enablement is done
# External PDB
/*variable "external_pdb_id" {
  default = "ocid1.externalpluggabledatabase"
}
variable "pdb_connector_id" {
  default = "ocid1.externaldatabaseconnector"
}

resource "oci_database_management_externalpluggabledatabase_external_pluggable_dbm_features_management" "test_externalpluggabledatabase_external_pluggable_dbm_features_management" {
  feature_details {
    connector_details {
      connector_type = "EXTERNAL"
      database_connector_id = var.pdb_connector_id
    }
    feature = "DIAGNOSTICS_AND_MANAGEMENT"
  }
  external_pluggable_database_id = var.external_pdb_id
  enable_external_pluggable_dbm_feature = "true"
}*/


# External Non-CDB
variable "external_non_cdb_id" {
  default = "ocid1.externalnoncontainerdatabase"
}
variable "non_cdb_connector_id" {
  default = "ocid1.externaldatabaseconnector"
}

resource "oci_database_management_externalnoncontainerdatabase_external_non_container_dbm_features_management" "test_externalnoncontainerdatabase_external_non_container_dbm_features_management" {
  feature_details {
    connector_details {
      connector_type = "EXTERNAL"
      database_connector_id = var.non_cdb_connector_id
    }
    feature = "DIAGNOSTICS_AND_MANAGEMENT"
    license_model = "LICENSE_INCLUDED"
  }
  external_non_container_database_id = var.external_non_cdb_id
  enable_external_non_container_dbm_feature = "true"
}

# Cloud CDB
variable "cloud_cdb_id" {
  default = "ocid1.database"
}
variable "cdb_pe_id" {
  default = "ocid1.dbmgmtprivateendpoint"
}
variable "vault_secret_id" {
  default = "ocid1.vaultsecret"
}
variable "cdb_user_role" {
  default = "SYSDBA"
}
variable "cdb_user" {
  default = "dbsnmp"
}
variable "cdb_service" {
  default = "cdb_service"
}
resource "oci_database_management_database_dbm_features_management" "test_database_dbm_features_management" {
  feature_details {
    connector_details {
      connector_type = "PE"
      private_end_point_id = var.cdb_pe_id
    }
    database_connection_details {
      connection_credentials {
        credential_type = "DETAILS"
        password_secret_id = var.vault_secret_id
        role = var.cdb_user_role
        user_name = var.cdb_user
      }
      connection_string {
        connection_type = "BASIC"
        port = "1521"
        protocol = "TCP"
        service = var.cdb_service
      }
    }
    feature = "DIAGNOSTICS_AND_MANAGEMENT"
    management_type = "ADVANCED"
  }
  database_id = var.cloud_cdb_id
  enable_database_dbm_feature = "true"
}

# Uncomment PDB enable APIs only after CDB enablement is done
# Cloud PDB
/*variable "cloud_pdb_id" {
  default = "ocid1.pluggabledatabase"
}
variable "pdb_pe_id" {
  default = "ocid1.dbmgmtprivateendpoint"
}
variable "vault_secret_id" {
  default = "ocid1.vaultsecret"
}
variable "pdb_user_role" {
  default = "SYSDBA"
}
variable "pdb_user" {
  default = "dbsnmp"
}
variable "db_service" {
  default = "pdb_service"
}
resource "oci_database_management_pluggabledatabase_pluggable_database_dbm_features_management" "test_pluggabledatabase_pluggable_database_dbm_features_management" {
  feature_details {
    connector_details {
      connector_type = "PE"
      private_end_point_id = var.pdb_pe_id
    }
    database_connection_details {
      connection_credentials {
        credential_type = "DETAILS"
        password_secret_id = var.vault_secret_id
        role = var.pdb_user_role
        user_name = var.pdb_user
      }
      connection_string {
        connection_type = "BASIC"
        port = "1521"
        protocol = "TCP"
        service = var.db_service
      }
    }
    feature = "DIAGNOSTICS_AND_MANAGEMENT"
    management_type = "ADVANCED"
  }
  pluggable_database_id = var.cloud_pdb_id
  enable_pluggable_database_dbm_feature = "true"
}*/
