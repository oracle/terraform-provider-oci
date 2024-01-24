// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "compartment_ocid" {}

variable "function_image" {}

variable "function_image_digest" {}

variable "private_key_data" {}

variable "instance_image_ocid" {
  type = "map"

  default = {
    # See https://docs.us-phoenix-1.oraclecloud.com/images/
    # Oracle-provided image "Oracle-Linux-7.5-2018.10.16-0"
    us-phoenix-1 = "ocid1.image.oc1.phx.aaaaaaaaoqj42sokaoh42l76wsyhn3k2beuntrh5maj3gmgmzeyr55zzrwwa"

    us-ashburn-1   = "ocid1.image.oc1.iad.aaaaaaaageeenzyuxgia726xur4ztaoxbxyjlxogdhreu3ngfj2gji3bayda"
    eu-frankfurt-1 = "ocid1.image.oc1.eu-frankfurt-1.aaaaaaaaitzn6tdyjer7jl34h2ujz74jwy5nkbukbh55ekp6oyzwrtfa4zma"
    uk-london-1    = "ocid1.image.oc1.uk-london-1.aaaaaaaa32voyikkkzfxyo4xbdmadc2dmvorfxxgdhpnk6dw64fa3l4jh7wa"
  }
}

variable "instance_shape" {
  default = "VM.Standard2.1"
}

/* DB SYSTEMS */
variable "db_system_shape" {
  default = "VM.Standard2.1"
}

variable "db_edition" {
  default = "ENTERPRISE_EDITION"
}

variable "db_admin_password" {
  default = "BEstrO0ng_#12"
}

variable "db_version" {
  default = "19.0.0.0"
}

variable "db_disk_redundancy" {
  default = "NORMAL"
}

variable "sparse_diskgroup" {
  default = true
}

variable "hostname" {
  default = "myoracledb"
}

variable "host_user_name" {
  default = "opc"
}

variable "n_character_set" {
  default = "AL16UTF16"
}

variable "character_set" {
  default = "AL32UTF8"
}

variable "db_workload" {
  default = "OLTP"
}

variable "pdb_name" {
  default = "pdbName"
}

variable "data_storage_size_in_gb" {
  default = "256"
}

variable "license_model" {
  default = "LICENSE_INCLUDED"
}

/* ADB */

variable "autonomous_database_backup_display_name" {
  default = "Monthly Backup"
}

variable "autonomous_database_db_workload" {
  default = "OLTP"
}

variable "autonomous_data_warehouse_db_workload" {
  default = "DW"
}

variable "autonomous_database_defined_tags_value" {
  default = "value"
}

variable "autonomous_database_freeform_tags" {
  default = {
    "Department" = "Finance"
  }
}

variable "autonomous_database_license_model" {
  default = "LICENSE_INCLUDED"
}

variable "autonomous_database_is_dedicated" {
  default = false
}

/* ATP-D */

variable "autonomous_exadata_infrastructure_domain" {
  default = "subnetexadatard.vcnrd.oraclevcn.com"
}

variable "autonomous_container_database_backup_config_recovery_window_in_days" {
  default = 10
}

/* Osmanagement */

variable "ssh_private_key" {}

/*  apigateway  */

variable "gateway_state" {
  default = "ACTIVE"
}

variable "deployment_state" {
  default = "ACTIVE"
}

variable "gateway_display_name" {
  default = "apigatewayGatewayRd"
}

variable "gateway_endpoint_type" {
  default = "PUBLIC"
}

variable "deployment_path_prefix" {
  default = "/v1"
}

variable "deployment_specification_logging_policies_access_log_is_enabled" {
  default = false
}

variable "deployment_specification_logging_policies_execution_log_is_enabled" {
  default = false
}

variable "deployment_specification_logging_policies_execution_log_log_level" {
  default = "INFO"
}

variable "deployment_specification_request_policies_authentication_type" {
  default = "CUSTOM_AUTHENTICATION"
}

variable "deployment_specification_request_policies_authentication_is_anonymous_access_allowed" {
  default = false
}

variable "deployment_specification_request_policies_authentication_token_header" {
  default = "Authorization"
}

variable "deployment_specification_request_policies_cors_allowed_origins" {
  default = ["*"]
}

variable "deployment_specification_request_policies_cors_allowed_headers" {
  default = ["*"]
}

variable "deployment_specification_request_policies_cors_allowed_methods" {
  default = ["*"]
}

variable "deployment_specification_request_policies_cors_exposed_headers" {
  default = ["*"]
}

variable "deployment_specification_request_policies_cors_is_allow_credentials_enabled" {
  default = false
}

variable "deployment_specification_request_policies_cors_max_age_in_seconds" {
  default = "600"
}

variable "deployment_specification_request_policies_rate_limiting_rate_in_requests_per_second" {
  default = 10
}

variable "deployment_specification_request_policies_rate_limiting_rate_key" {
  default = "CLIENT_IP"
}

variable "deployment_specification_routes_backend_type" {
  default = "HTTP_BACKEND"
}

variable "deployment_specification_routes_backend_url" {
  default = "https://api.weather.gov"
}

variable "deployment_specification_routes_path" {
  default = "/hello"
}

variable "deployment_specification_routes_logging_policies_access_log_is_enabled" {
  default = false
}

variable "deployment_specification_routes_logging_policies_execution_log_is_enabled" {
  default = false
}

variable "deployment_specification_routes_logging_policies_execution_log_log_level" {
  default = "INFO"
}

variable "deployment_specification_routes_methods" {
  default = ["GET"]
}

variable "deployment_specification_routes_request_policies_authorization_type" {
  default = "AUTHENTICATION_ONLY"
}

variable "deployment_specification_routes_request_policies_cors_allowed_origins" {
  default = ["*"]
}

variable "deployment_specification_routes_request_policies_cors_allowed_headers" {
  default = ["*"]
}

variable "deployment_specification_routes_request_policies_cors_allowed_methods" {
  default = ["GET"]
}

variable "deployment_specification_routes_request_policies_cors_exposed_headers" {
  default = ["*"]
}

variable "deployment_specification_routes_request_policies_cors_is_allow_credentials_enabled" {
  default = false
}

variable "deployment_specification_routes_request_policies_cors_max_age_in_seconds" {
  default = "600"
}

/* function */
variable "config" {
  default = {
    "MY_FUNCTION_CONFIG" = "ConfVal"
  }
}

variable "function_memory_in_mbs" {
  default = 128
}

variable "function_timeout_in_seconds" {
  default = 30
}

variable "application_state" {
  default = "AVAILABLE"
}

/* Identity */

variable "authentication_policy_password_policy_is_lowercase_characters_required" {
  default = true
}

variable "authentication_policy_password_policy_is_numeric_characters_required" {
  default = true
}

variable "authentication_policy_password_policy_is_special_characters_required" {
  default = true
}

variable "authentication_policy_password_policy_is_uppercase_characters_required" {
  default = true
}

variable "authentication_policy_password_policy_is_username_containment_allowed" {
  default = false
}

variable "authentication_policy_password_policy_minimum_password_length" {
  default = 11
}

variable "network_source_defined_tags_value" {
  default = "value"
}

variable "network_source_description" {
  default = "corporate ip ranges to be used for ip based authorization"
}

variable "network_source_freeform_tags" {
  default = {
    "Department" = "Finance"
  }
}

variable "network_source_public_source_list" {
  default = ["128.2.13.5"]
}

variable "network_source_services" {
  default = ["all"]
}

variable "network_source_virtual_source_list" {
  default = []
}

/* datascience */

variable "datascience_notebook_session_shape" {
  default = "VM.Standard.E2.2"
}

variable "datascience_model_artifact_content_length" {
  default = "9"
}

variable "datascience_model_artifact" {
  default = "./resources/model_artifact.xml"
}

/* Core */
variable "freeform_tags" {
  type = "map"

  default = {
    Department = "Finance"
  }
}

variable "volume_backup_type" {
  default = "FULL"
}
