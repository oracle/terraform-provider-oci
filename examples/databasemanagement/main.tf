// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "managed_database_group_description" {
  default = "Sales test database group"
}

variable "managed_database_group_id" {
  default = "id"
}

variable "managed_database_group_name" {
  default = "TestGroup"
}

variable "managed_database_group_state" {
  default = "ACTIVE"
}

variable "managed_database_id" {
  default = "testManagedDatabase0"
}

variable "managed_databases_database_parameter_credentials_username" {
  default = "sys"
}

variable "managed_databases_database_parameter_credentials_password" {
  default = "sys"
}

variable "managed_databases_database_parameter_credentials_role" {
  default = "NORMAL"
}

variable "managed_databases_database_parameter_parameters_name" {
  default = "open_cursors"
}

variable "managed_databases_database_parameter_parameters_value" {
  default = "305"
}

variable "managed_databases_database_parameter_update_comment" {
  default = "Terraform update of open cursors"
}

variable "managed_databases_database_parameter_scope" {
  default = "BOTH"
}

variable "managed_databases_database_parameter_is_allowed_values_included" {
  default = "false"
}

variable "managed_databases_database_parameter_source" {
  default = "CURRENT"
}

variable "db_management_private_endpoint_name" {
  default = "TestPrivateEndpoint"
}

variable "db_management_private_endpoint_description" {
  default = "Test private endpoint"
}

variable "db_management_private_endpoint_state" {
  default = "ACTIVE"
}

variable "db_management_private_endpoint_is_cluster" {
  default = false
}

variable "managed_database_sql_tuning_advisor_task_name" {
  default = "name"
}

variable "managed_database_sql_tuning_advisor_task_status" {
  default = "INITIAL"
}

variable "managed_database_sql_tuning_advisor_task_time_greater_than_or_equal_to" {
  default = "timeGreaterThanOrEqualTo"
}

variable "managed_database_sql_tuning_advisor_task_time_less_than_or_equal_to" {
  default = "timeLessThanOrEqualTo"
}

variable "managed_database_sql_tuning_advisor_tasks_finding_finding_filter" {
  default = "none"
}

variable "managed_database_sql_tuning_advisor_tasks_finding_index_hash_filter" {
  default = "indexHashFilter"
}

variable "managed_database_sql_tuning_advisor_tasks_finding_search_period" {
  default = "LAST_24HR"
}

variable "managed_database_sql_tuning_advisor_tasks_finding_stats_hash_filter" {
  default = "statsHashFilter"
}

variable "managed_database_sql_tuning_advisor_tasks_sql_execution_plan_attribute" {
  default = "ORIGINAL"
}

variable "managed_database_sql_tuning_advisor_tasks_summary_report_begin_exec_id_greater_than_or_equal_to" {
  default = 10
}

variable "managed_database_sql_tuning_advisor_tasks_summary_report_end_exec_id_less_than_or_equal_to" {
  default = 10
}

variable "managed_database_sql_tuning_advisor_tasks_summary_report_search_period" {
  default = "LAST_24HR"
}

variable "managed_database_sql_tuning_advisor_tasks_summary_report_time_greater_than_or_equal_to" {
  default = "timeGreaterThanOrEqualTo"
}

variable "managed_database_sql_tuning_advisor_tasks_summary_report_time_less_than_or_equal_to" {
  default = "timeLessThanOrEqualTo"
}

variable "managed_databases_asm_property_name" {
  default = "DATA"
}

variable "managed_database_optimizer_statistics_advisor_execution_end_time_less_than_or_equal_to" {
  default = "endTimeLessThanOrEqualTo"
}

variable "managed_database_optimizer_statistics_advisor_execution_start_time_greater_than_or_equal_to" {
  default = "startTimeGreaterThanOrEqualTo"
}

variable "managed_database_optimizer_statistics_advisor_execution_script_execution_name" {
  default = "executionName"
}

variable "managed_database_optimizer_statistics_advisor_execution_script_task_name" {
  default = "taskName"
}

variable "managed_database_optimizer_statistics_collection_aggregation_end_time_less_than_or_equal_to" {
  default = "endTimeLessThanOrEqualTo"
}

variable "managed_database_optimizer_statistics_collection_aggregation_group_type" {
  default = "TASK_STATUS"
}

variable "managed_database_optimizer_statistics_collection_aggregation_start_time_greater_than_or_equal_to" {
  default = "startTimeGreaterThanOrEqualTo"
}

variable "managed_database_optimizer_statistics_collection_aggregation_task_type" {
  default = "ALL"
}

variable "managed_database_optimizer_statistics_collection_operation_end_time_less_than_or_equal_to" {
  default = "endTimeLessThanOrEqualTo"
}

variable "managed_database_optimizer_statistics_collection_operation_filter_by" {
  default = "filterBy"
}

variable "managed_database_optimizer_statistics_collection_operation_start_time_greater_than_or_equal_to" {
  default = "startTimeGreaterThanOrEqualTo"
}

variable "managed_database_optimizer_statistics_collection_operation_task_type" {
  default = "ALL"
}

variable "managed_database_preferred_credential_credential_name" {
  default = "credentialName"
}

provider "oci" {
  tenancy_ocid = var.tenancy_ocid
  user_ocid = var.user_ocid
  fingerprint = var.fingerprint
  private_key_path = var.private_key_path
  region = var.region
}

resource "oci_database_management_managed_database_group" "test_managed_database_group" {
  #Required
  compartment_id = var.compartment_id
  name = var.managed_database_group_name

  #Optional
  description = var.managed_database_group_description
  managed_databases {
    id = var.managed_database_id
  }
}

data "oci_database_management_managed_database_sql_tuning_sets" "test_managed_database_sql_tuning_sets" {
	#Required
	managed_database_id = oci_database_management_managed_database.test_managed_database.id

	#Optional
	name_contains = var.managed_database_sql_tuning_set_name_contains
	owner = var.managed_database_sql_tuning_set_owner
}

data "oci_database_management_managed_database_groups" "test_managed_database_groups_with_id" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  id = oci_database_management_managed_database_group.test_managed_database_group.id
  state = var.managed_database_group_state
}

data "oci_database_management_managed_database_groups" "test_managed_database_groups_with_name" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  name = var.managed_database_group_name
  state = var.managed_database_group_state
}

resource "oci_database_management_managed_databases_change_database_parameter" "test_managed_databases_change_database_parameter" {
  #Required
  credentials {

    #Optional
    password = var.managed_databases_database_parameter_credentials_password
    role = var.managed_databases_database_parameter_credentials_role
    user_name = var.managed_databases_database_parameter_credentials_username
  }
  managed_database_id = var.managed_database_id
  parameters {
    #Required
    name = var.managed_databases_database_parameter_parameters_name
    value = var.managed_databases_database_parameter_parameters_value

    #Optional
    update_comment = var.managed_databases_database_parameter_update_comment
  }
  scope = var.managed_databases_database_parameter_scope
}

resource "oci_database_management_managed_databases_reset_database_parameter" "test_managed_databases_reset_database_parameter" {
  #Required
  credentials {

    #Optional
    password = var.managed_databases_database_parameter_credentials_password
    role = var.managed_databases_database_parameter_credentials_role
    user_name = var.managed_databases_database_parameter_credentials_username
  }
  managed_database_id = var.managed_database_id
  parameters = [var.managed_databases_database_parameter_parameters_name]
  scope = var.managed_databases_database_parameter_scope
}

data "oci_database_management_managed_databases_database_parameter" "test_managed_databases_database_parameter" {
  #Required
  managed_database_id = var.managed_database_id

  #Optional
  is_allowed_values_included = var.managed_databases_database_parameter_is_allowed_values_included
  name = var.managed_databases_database_parameter_parameters_name
  source = var.managed_databases_database_parameter_source
}

resource "oci_core_vcn" "test_vcn" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = var.compartment_id
  display_name   = "exampleVCN"
  dns_label      = "tfexamplevcn"
}

resource "oci_core_subnet" "test_subnet" {
  cidr_block     = "10.0.0.0/24"
  display_name   = "regionalSubnet"
  dns_label      = "regionalsubnet"
  compartment_id = var.compartment_id
  vcn_id         = oci_core_vcn.test_vcn.id
}

resource "oci_core_network_security_group" "test_network_security_group" {
  #Required
  compartment_id = var.compartment_id
  vcn_id         = oci_core_vcn.test_vcn.id
}

data "oci_database_management_managed_databases" "test_managed_databases" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	deployment_type = var.managed_database_deployment_type
	external_exadata_infrastructure_id = oci_database_management_external_exadata_infrastructure.test_external_exadata_infrastructure.id
	id = var.managed_database_id
	management_option = var.managed_database_management_option
	name = var.managed_database_name
}

resource "oci_database_management_db_management_private_endpoint" "test_db_management_private_endpoint" {
  #Required
  compartment_id = var.compartment_id
  name = var.db_management_private_endpoint_name
  subnet_id = oci_core_subnet.test_subnet.id

  #Optional
  description = var.db_management_private_endpoint_description
  nsg_ids   = [oci_core_network_security_group.test_network_security_group.id]
  is_cluster  = var.db_management_private_endpoint_is_cluster
}

data "oci_database_management_db_management_private_endpoint" "test_db_management_private_endpoint" {
  db_management_private_endpoint_id = oci_database_management_db_management_private_endpoint.test_db_management_private_endpoint.id
}

data "oci_database_management_db_management_private_endpoints" "test_db_management_private_endpoints" {
  #Required
  compartment_id = var.compartment_id
}

data "oci_database_management_db_management_private_endpoints" "test_db_management_private_endpoints_with_name" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  name = var.db_management_private_endpoint_name
  vcn_id = oci_core_vcn.test_vcn.id
  state = var.db_management_private_endpoint_state
  is_cluster = var.db_management_private_endpoint_is_cluster
}

data "oci_database_management_job_executions_status" "test_job_executions_status" {
  #Required
  compartment_id = var.compartment_id
  start_time = formatdate("YYYY-MM-DD'T'hh:mm:ss'.000'Z", timeadd(timestamp(), "-12h"))
  end_time = formatdate("YYYY-MM-DD'T'hh:mm:ss'.000'Z", timestamp())

  #Optional
  managed_database_id = var.managed_database_id
}

data "oci_database_management_managed_database_sql_tuning_advisor_tasks" "test_managed_database_sql_tuning_advisor_tasks" {
  #Required
  managed_database_id = oci_database_management_managed_database.test_managed_database.id

  #Optional
  name                          = var.managed_database_sql_tuning_advisor_task_name
  status                        = var.managed_database_sql_tuning_advisor_task_status
  time_greater_than_or_equal_to = var.managed_database_sql_tuning_advisor_task_time_greater_than_or_equal_to
  time_less_than_or_equal_to    = var.managed_database_sql_tuning_advisor_task_time_less_than_or_equal_to
}

data "oci_database_management_managed_database_sql_tuning_advisor_tasks_execution_plan_stats_comparision" "test_managed_database_sql_tuning_advisor_tasks_execution_plan_stats_comparision" {
  #Required
  execution_id               = oci_database_management_execution.test_execution.id
  managed_database_id        = oci_database_management_managed_database.test_managed_database.id
  sql_object_id              = oci_objectstorage_object.test_object.id
  sql_tuning_advisor_task_id = oci_database_management_sql_tuning_advisor_task.test_sql_tuning_advisor_task.id
}

data "oci_database_management_managed_database_sql_tuning_advisor_tasks_findings" "test_managed_database_sql_tuning_advisor_tasks_findings" {
  #Required
  managed_database_id        = oci_database_management_managed_database.test_managed_database.id
  sql_tuning_advisor_task_id = oci_database_management_sql_tuning_advisor_task.test_sql_tuning_advisor_task.id

  #Optional
  begin_exec_id     = oci_database_management_begin_exec.test_begin_exec.id
  end_exec_id       = oci_database_management_end_exec.test_end_exec.id
  finding_filter    = var.managed_database_sql_tuning_advisor_tasks_finding_finding_filter
  index_hash_filter = var.managed_database_sql_tuning_advisor_tasks_finding_index_hash_filter
  search_period     = var.managed_database_sql_tuning_advisor_tasks_finding_search_period
  stats_hash_filter = var.managed_database_sql_tuning_advisor_tasks_finding_stats_hash_filter
}

data "oci_database_management_managed_database_sql_tuning_advisor_tasks_recommendations" "test_managed_database_sql_tuning_advisor_tasks_recommendations" {
  #Required
  execution_id               = oci_database_management_execution.test_execution.id
  managed_database_id        = oci_database_management_managed_database.test_managed_database.id
  sql_object_id              = oci_objectstorage_object.test_object.id
  sql_tuning_advisor_task_id = oci_database_management_sql_tuning_advisor_task.test_sql_tuning_advisor_task.id
}

data "oci_database_management_managed_database_sql_tuning_advisor_tasks_sql_execution_plan" "test_managed_database_sql_tuning_advisor_tasks_sql_execution_plan" {
  #Required
  attribute                  = var.managed_database_sql_tuning_advisor_tasks_sql_execution_plan_attribute
  managed_database_id        = oci_database_management_managed_database.test_managed_database.id
  sql_object_id              = oci_objectstorage_object.test_object.id
  sql_tuning_advisor_task_id = oci_database_management_sql_tuning_advisor_task.test_sql_tuning_advisor_task.id
}

data "oci_database_management_managed_database_sql_tuning_advisor_tasks_summary_report" "test_managed_database_sql_tuning_advisor_tasks_summary_report" {
  #Required
  managed_database_id        = oci_database_management_managed_database.test_managed_database.id
  sql_tuning_advisor_task_id = oci_database_management_sql_tuning_advisor_task.test_sql_tuning_advisor_task.id

  #Optional
  begin_exec_id_greater_than_or_equal_to = var.managed_database_sql_tuning_advisor_tasks_summary_report_begin_exec_id_greater_than_or_equal_to
  end_exec_id_less_than_or_equal_to      = var.managed_database_sql_tuning_advisor_tasks_summary_report_end_exec_id_less_than_or_equal_to
  search_period                          = var.managed_database_sql_tuning_advisor_tasks_summary_report_search_period
  time_greater_than_or_equal_to          = var.managed_database_sql_tuning_advisor_tasks_summary_report_time_greater_than_or_equal_to
  time_less_than_or_equal_to             = var.managed_database_sql_tuning_advisor_tasks_summary_report_time_less_than_or_equal_to
}

data "oci_database_management_managed_databases_asm_properties" "test_managed_databases_asm_properties" {
  #Required
  managed_database_id = oci_database_management_managed_database.test_managed_database.id

  #Optional
  name = var.managed_databases_asm_property_name
}

data "oci_database_management_managed_databases_asm_property" "test_managed_databases_asm_property" {
  #Required
  managed_database_id = oci_database_management_managed_database.test_managed_database.id

  #Optional
  name = var.managed_databases_asm_property_name
}

data "oci_database_management_managed_database_optimizer_statistics_advisor_executions" "test_managed_database_optimizer_statistics_advisor_executions" {
  #Required
  managed_database_id = oci_database_management_managed_database.test_managed_database.id

  #Optional
  end_time_less_than_or_equal_to      = var.managed_database_optimizer_statistics_advisor_execution_end_time_less_than_or_equal_to
  start_time_greater_than_or_equal_to = var.managed_database_optimizer_statistics_advisor_execution_start_time_greater_than_or_equal_to
}

data "oci_database_management_managed_database_optimizer_statistics_advisor_execution_scripts" "test_managed_database_optimizer_statistics_advisor_execution_scripts" {
  #Required
  execution_name      = var.managed_database_optimizer_statistics_advisor_execution_script_execution_name
  managed_database_id = oci_database_management_managed_database.test_managed_database.id
  task_name           = var.managed_database_optimizer_statistics_advisor_execution_script_task_name
}

data "oci_database_management_managed_database_optimizer_statistics_collection_aggregations" "test_managed_database_optimizer_statistics_collection_aggregations" {
  #Required
  group_type          = var.managed_database_optimizer_statistics_collection_aggregation_group_type
  managed_database_id = oci_database_management_managed_database.test_managed_database.id

  #Optional
  end_time_less_than_or_equal_to      = var.managed_database_optimizer_statistics_collection_aggregation_end_time_less_than_or_equal_to
  start_time_greater_than_or_equal_to = var.managed_database_optimizer_statistics_collection_aggregation_start_time_greater_than_or_equal_to
  task_type                           = var.managed_database_optimizer_statistics_collection_aggregation_task_type
}

data "oci_database_management_managed_database_optimizer_statistics_collection_operations" "test_managed_database_optimizer_statistics_collection_operations" {
  #Required
  managed_database_id = oci_database_management_managed_database.test_managed_database.id

  #Optional
  end_time_less_than_or_equal_to      = var.managed_database_optimizer_statistics_collection_operation_end_time_less_than_or_equal_to
  filter_by                           = var.managed_database_optimizer_statistics_collection_operation_filter_by
  start_time_greater_than_or_equal_to = var.managed_database_optimizer_statistics_collection_operation_start_time_greater_than_or_equal_to
  task_type                           = var.managed_database_optimizer_statistics_collection_operation_task_type
}

data "oci_database_management_managed_database_table_statistics" "test_managed_database_table_statistics" {
  #Required
  managed_database_id = oci_database_management_managed_database.test_managed_database.id
}

data "oci_database_management_managed_database_preferred_credentials" "test_managed_database_preferred_credentials" {
  #Required
  managed_database_id = oci_database_management_managed_database.test_managed_database.id
}

data "oci_database_management_managed_database_preferred_credential" "test_managed_database_preferred_credential" {
  #Required
  credential_name = var.managed_database_preferred_credential_credential_name
  managed_database_id = oci_database_management_managed_database.test_managed_database.id
}

####################### External DB System #########################
variable "external_db_system_discovery_display_name" {
  default = "EXAMPLE-displayName-Value"
}

variable "agent_id" {
  default = "ocid1.test.oc1..<unique_ID>EXAMPLE-agentId-Value"
}

variable "external_db_system_database_management_config_license_model" {
  default = "LICENSE_INCLUDED"
}

variable "external_db_system_display_name" {
  default = "EXAMPLE-displayName-Value"
}

variable "db_host_name" {
  default = "EXAMPLE-hostName-Value"
}

variable "db_port" {
  default = "1521"
}

variable "db_service_name" {
  default = "EXAMPLE-service-Value"
}

variable "db_user_name" {
  default = "EXAMPLE-userName-Value"
}

variable "db_password_secret_id" {
  default = "ocid1.test.oc1..<unique_ID>EXAMPLE-passwordSecretId-Value"
}

variable "db_credential_name" {
  default = "EXAMPLE-dbCredName-Value"
}

variable "asm_host_name" {
  default = "EXAMPLE-hostName-Value"
}

variable "asm_port" {
  default = "1521"
}

variable "asm_service_name" {
  default = "EXAMPLE-service-Value"
}

variable "asm_user_name" {
  default = "EXAMPLE-userName-Value"
}

variable "asm_password_secret_id" {
  default = "ocid1.test.oc1..<unique_ID>EXAMPLE-passwordSecretId-Value"
}

variable "asm_credential_name" {
  default = "EXAMPLE-asmCredName-Value"
}

variable "external_asm_connector_display_name" {
  default = "asmConnectorName"
}

variable "external_listener_connector_display_name" {
  default = "listenerConnectorName"
}

variable "connector_agent_id" {
  default = "ocid1.test.oc1..<unique_ID>EXAMPLE-agentId-Value"
}

# Create a new ExternalDbSystemDiscovery resource and discover an external DB System and its components.
# Also add a connector to the discovered Oracle Container Database (CDB).
resource "oci_database_management_external_db_system_discovery" "test_external_db_system_discovery" {
  #Required
  agent_id                        = var.agent_id
  compartment_id                  = var.compartment_id

  #Optional
  display_name = var.external_db_system_discovery_display_name

  # Patch the Discovery resource and add connector to the database component
  patch_operations {
    operation = "MERGE"
    selection = "discoveredComponents[?componentType == 'DATABASE'] | [0]"
    value {
      connector {
        agent_id = var.agent_id
        connection_info {
          component_type = "DATABASE"
          connection_credentials {
            credential_name = var.db_credential_name
            credential_type = "DETAILS"
            password_secret_id = var.db_password_secret_id
            role = "NORMAL"
            user_name = var.db_user_name
          }
          connection_string {
            host_name = var.db_host_name
            port = var.db_port
            protocol = "TCP"
            service = var.db_service_name
          }
        }
        connector_type = "MACS"
        display_name = "EXAMPLE-displayName-Value"
      }
      is_selected_for_monitoring = "true"
    }
  }

  # Deselect the PDBs
  patch_operations {
    operation = "MERGE"
    selection = "discoveredComponents[?componentType == 'DATABASE'].pluggableDatabases"
    value {
      is_selected_for_monitoring = "false"
    }
  }
}

# List ExternalDbSystemDiscovery resources
data "oci_database_management_external_db_system_discoveries" "test_external_db_system_discoveries" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = var.external_db_system_discovery_display_name
}

# Create a new ExternalDbSystem resource from ExternalDbSystemDiscovery resource
resource "oci_database_management_external_db_system" "test_external_db_system" {
  #Required
  compartment_id         = var.compartment_id
  db_system_discovery_id = oci_database_management_external_db_system_discovery.test_external_db_system_discovery.id

  #Optional
  database_management_config {
    #Required
    license_model = var.external_db_system_database_management_config_license_model
  }
  display_name = var.external_db_system_display_name
}

# List ExternalDbSystem resources
data "oci_database_management_external_db_systems" "test_external_db_systems" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = var.external_db_system_display_name
}

# Add connector to an ExternalAsm resource
resource "oci_database_management_external_asm" "test_external_asm" {
  external_asm_id = data.oci_database_management_external_asms.test_external_asms.external_asm_collection.0.items.0.id
  external_connector_id = oci_database_management_external_db_system_connector.test_external_asm_connector.id
}

# List ExternalAsms in ExternalDbSystem
data "oci_database_management_external_asms" "test_external_asms" {
  #Required
  external_db_system_id = oci_database_management_external_db_system.test_external_db_system.id

  #Optional
  compartment_id        = var.compartment_id
}

# Get configuration details for ExternalAsm
data "oci_database_management_external_asm_configuration" "test_external_asm_configuration" {
  #Required
  external_asm_id = oci_database_management_external_asm.test_external_asm.id
}

# List ASM disk groups in ExternalAsm
data "oci_database_management_external_asm_disk_groups" "test_external_asm_disk_groups" {
  #Required
  external_asm_id = oci_database_management_external_asm.test_external_asm.id
}

# List ASM users in ExternalAsm
data "oci_database_management_external_asm_users" "test_external_asm_users" {
  #Required
  external_asm_id = oci_database_management_external_asm.test_external_asm.id
}

# List ExternalAsmInstances in ExternalAsm
data "oci_database_management_external_asm_instances" "test_external_asm_instances" {
  #Required
  external_asm_id = oci_database_management_external_asm.test_external_asm.id

  #Optional
  compartment_id  = var.compartment_id
}

# List ExternalClusters in ExternalDbSystem
data "oci_database_management_external_clusters" "test_external_clusters" {
  #Required
  external_db_system_id = oci_database_management_external_db_system.test_external_db_system.id

  #Optional
  compartment_id        = var.compartment_id
}

# List ExternalClusterInstances in ExternalCluster
data "oci_database_management_external_cluster_instances" "test_external_cluster_instances" {
  #Required
  external_cluster_id = data.oci_database_management_external_clusters.test_external_clusters.external_cluster_collection.0.items.0.id

  #Optional
  compartment_id      = var.compartment_id
}

# List ExternalDatabases in ExternalDbSystem
data "oci_database_management_external_databases" "test_external_databases" {
  #Required
  external_db_system_id = oci_database_management_external_db_system.test_external_db_system.id

  #Optional
  compartment_id        = var.compartment_id
}

# List ExternalDbHomes in ExternalDbSystem
data "oci_database_management_external_db_homes" "test_external_db_homes" {
  #Required
  external_db_system_id = oci_database_management_external_db_system.test_external_db_system.id

  #Optional
  compartment_id        = var.compartment_id
}

# List ExternalDbNodes in ExternalDbSystem
data "oci_database_management_external_db_nodes" "test_external_db_nodes" {
  #Required
  external_db_system_id = oci_database_management_external_db_system.test_external_db_system.id

  #Optional
  compartment_id        = var.compartment_id
}

# List ExternalListeners in ExternalDbSystem
data "oci_database_management_external_listeners" "test_external_listeners" {
  #Required
  external_db_system_id = oci_database_management_external_db_system.test_external_db_system.id

  #Optional
  compartment_id        = var.compartment_id
}

resource "oci_database_management_external_listener" "test_external_listener" {
  #Required
  external_listener_id = data.oci_database_management_external_listeners.test_external_listeners.external_listener_collection.0.items.0.id
  external_connector_id = oci_database_management_external_db_system_connector.test_external_listener_connector.id
}

# List ExternalListenerServices for ExternalListener
data "oci_database_management_external_listener_services" "test_external_listener_services" {
  #Required
  external_listener_id = oci_database_management_external_listener.test_external_listener.id
  managed_database_id  = oci_database_management_external_listener.test_external_listener.serviced_databases.0.id
}

# Create a new Management Agent based ExternalDbSystemConnector
resource "oci_database_management_external_db_system_connector" "test_external_asm_connector" {
  #Required
  connector_type        = "MACS"
  external_db_system_id = oci_database_management_external_db_system.test_external_db_system.id
  agent_id              = var.connector_agent_id

  #Optional
  display_name = var.external_asm_connector_display_name
  connection_info {
    component_type = "ASM"
    connection_credentials {
      credential_name = var.asm_credential_name
      credential_type = "DETAILS"
      password_secret_id = var.asm_password_secret_id
      role = "SYSASM"
      user_name = var.asm_user_name
    }
    connection_string {
      hosts = [var.asm_host_name]
      port = var.asm_port
      protocol = "TCP"
      service = var.asm_service_name
    }
  }
}

resource "oci_database_management_external_db_system_connector" "test_external_listener_connector" {
  #Required
  connector_type        = "MACS"
  external_db_system_id = oci_database_management_external_db_system.test_external_db_system.id
  agent_id              = var.connector_agent_id

  #Optional
  display_name = var.external_listener_connector_display_name
}

# List ExternalDbSystemConnector resources in ExternalDbSystem
data "oci_database_management_external_db_system_connectors" "test_external_db_system_connectors" {
  #Required
  external_db_system_id = oci_database_management_external_db_system.test_external_db_system.id

  #Optional
  compartment_id        = var.compartment_id
  display_name          = var.external_asm_connector_display_name
  depends_on = [oci_database_management_external_db_system_connector.test_external_asm_connector]
}

## Disable database management for ExternalDbSystem
#resource "oci_database_management_external_db_system_database_managements_management" "test_external_db_system_database_managements_management" {
#  #Required
#  external_db_system_id      = oci_database_management_external_db_system.test_external_db_system.id
#  enable_database_management = false
#}

# Enable Stack Monitoring for ExternalDbSystem
resource "oci_database_management_external_db_system_stack_monitorings_management" "test_external_db_system_stack_monitoring_management" {
  #Required
  external_db_system_id      = oci_database_management_external_db_system.test_external_db_system.id
  enable_stack_monitoring    = true
  is_enabled                 = true
}

## Disable Stack Monitoring for ExternalDbSystem
#resource "oci_database_management_external_db_system_stack_monitorings_management" "test_external_db_system_stack_monitoring_management" {
#  #Required
#  external_db_system_id      = oci_database_management_external_db_system.test_external_db_system.id
#  enable_stack_monitoring    = false
#}

####################### SQL Plan Management #########################
variable "managed_database_sql_plan_baseline_is_accepted" {
  default = false
}

variable "managed_database_sql_plan_baseline_is_adaptive" {
  default = false
}

variable "managed_database_sql_plan_baseline_is_enabled" {
  default = true
}

variable "managed_database_sql_plan_baseline_is_fixed" {
  default = false
}

variable "managed_database_sql_plan_baseline_is_reproduced" {
  default = false
}

variable "managed_database_sql_plan_baseline_origin" {
  default = "AUTO_CAPTURE"
}

variable "managed_database_sql_plan_baseline_plan_name" {
  default = "planName"
}

variable "managed_database_sql_plan_baseline_sql_handle" {
  default = "sqlHandle"
}

variable "managed_database_sql_plan_baseline_sql_text" {
  default = "sqlText"
}

variable "managed_database_sql_plan_baseline_job_name" {
  default = "TestJobName"
}

# Get SQL Plan Baseline configuration details for the managed database
data "oci_database_management_managed_database_sql_plan_baseline_configuration" "test_managed_database_sql_plan_baseline_configuration" {
  #Required
  managed_database_id = var.managed_database_id
}

# List SQL Plan Baselines
data "oci_database_management_managed_database_sql_plan_baselines" "test_managed_database_sql_plan_baselines" {
  #Required
  managed_database_id = var.managed_database_id

  #Optional
  origin        = var.managed_database_sql_plan_baseline_origin
  plan_name     = var.managed_database_sql_plan_baseline_plan_name
  is_enabled    = var.managed_database_sql_plan_baseline_is_enabled
  is_accepted   = var.managed_database_sql_plan_baseline_is_accepted
  is_adaptive   = var.managed_database_sql_plan_baseline_is_adaptive
  is_fixed      = var.managed_database_sql_plan_baseline_is_fixed
  is_reproduced = var.managed_database_sql_plan_baseline_is_reproduced
  sql_handle    = var.managed_database_sql_plan_baseline_sql_handle
  sql_text      = var.managed_database_sql_plan_baseline_sql_text
}

data "oci_database_management_managed_database_sql_plan_baseline_jobs" "test_managed_database_sql_plan_baseline_jobs" {
  #Required
  managed_database_id = var.managed_database_id

  #Optional
  name = var.managed_database_sql_plan_baseline_job_name
}

data "oci_database_management_managed_database_cursor_cache_statements" "test_managed_database_cursor_cache_statements" {
  #Required
  managed_database_id = var.managed_database_id

  #Optional
  sql_text = var.managed_database_sql_plan_baseline_sql_text
}

####################### Exadata Infrastructure Monitoring #########################

variable "external_exadata_infrastructure_database_managements_management_license_model" {
  default = "LICENSE_INCLUDED"
}

variable "enable_exadata" {
  default = false
}

variable "external_exadata_infrastructure_db_system_ids" {
  default = []
}

variable "external_exadata_infrastructure_display_name" {
  default = "exadata-Terraform-Testing"
}

variable "external_exadata_infrastructure_license_model" {
  default = "LICENSE_INCLUDED"
}

variable "external_exadata_infrastructure_storage_server_names" {
  default = []
}

variable "external_exadata_storage_connector_connection_uri" {
  default = "https://exaInfra01celadm01.us.oracle.com/MS/RESTService/"
}

variable "external_exadata_storage_connector_connector_name" {
  default = "connectorName"
}

variable "external_exadata_storage_connector_credential_info_password" {
  default = "BEstrO0ng_#11"
}

variable "external_exadata_storage_connector_credential_info_ssl_trust_store_location" {
  default = "sslTrustStoreLocation"
}

variable "external_exadata_storage_connector_credential_info_ssl_trust_store_password" {
  default = "sslTrustStorePassword"
}

variable "external_exadata_storage_connector_credential_info_ssl_trust_store_type" {
  default = "JKS"
}

variable "external_exadata_storage_connector_credential_info_username" {
  default = "username"
}

variable "external_exadata_storage_connector_display_name" {
  default = "exaInfra01celadm01-conn"
}

variable "external_exadata_storage_grid_id" {
  default = "id"
}

variable "external_exadata_storage_server_id" {
  default = "id"
}

resource "oci_database_management_external_exadata_infrastructure" "test_external_exadata_infrastructure" {
  #Required
  compartment_id = var.compartment_id
  db_system_ids  = var.external_exadata_infrastructure_db_system_ids
  display_name   = var.external_exadata_infrastructure_display_name

  #Optional
  discovery_key        = var.external_exadata_infrastructure_discovery_key
  license_model        = var.external_exadata_infrastructure_license_model
  storage_server_names = var.external_exadata_infrastructure_storage_server_names
}

data "oci_database_management_external_exadata_infrastructures" "test_external_exadata_infrastructures" {
  #Required
  compartment_id = var.compartment_id

  #Optional
  display_name = var.external_exadata_infrastructure_display_name
}

data "oci_database_management_external_exadata_infrastructure" "test_external_exadata_infrastructure" {
	#Required
	external_exadata_infrastructure_id = oci_database_management_external_exadata_infrastructure.test_external_exadata_infrastructure.id
}

resource "oci_database_management_external_exadata_storage_connector" "test_external_exadata_storage_connector" {
	#Required
	agent_id = oci_cloud_bridge_agent.test_agent.id
	connection_uri = var.external_exadata_storage_connector_connection_uri
	connector_name = var.external_exadata_storage_connector_connector_name
	credential_info {
		#Required
		password = var.external_exadata_storage_connector_credential_info_password
		username = var.external_exadata_storage_connector_credential_info_username

		#Optional
		ssl_trust_store_location = var.external_exadata_storage_connector_credential_info_ssl_trust_store_location
		ssl_trust_store_password = var.external_exadata_storage_connector_credential_info_ssl_trust_store_password
		ssl_trust_store_type = var.external_exadata_storage_connector_credential_info_ssl_trust_store_type
	}
	storage_server_id = oci_database_management_storage_server.test_storage_server.id
}

data "oci_database_management_external_exadata_storage_connectors" "test_external_exadata_storage_connectors" {
	#Required
	compartment_id = var.compartment_id
	external_exadata_infrastructure_id = oci_database_management_external_exadata_infrastructure.test_external_exadata_infrastructure.id

	#Optional
	display_name = var.external_exadata_storage_connector_display_name
}

data "oci_database_management_external_exadata_storage_connector" "test_external_exadata_storage_connector" {
	#Required
	external_exadata_storage_connector_id = oci_database_management_external_exadata_storage_connector.test_external_exadata_storage_connector.id
}

data "oci_database_management_external_exadata_storage_grid" "test_external_exadata_storage_grid" {
	#Required
	external_exadata_storage_grid_id = var.external_exadata_storage_grid_id
}

data "oci_database_management_external_exadata_storage_server_iorm_plan" "test_external_exadata_storage_server_iorm_plan" {
	#Required
	external_exadata_storage_server_id = var.external_exadata_storage_server_id
}

data "oci_database_management_external_exadata_storage_server_open_alert_history" "test_external_exadata_storage_server_open_alert_history" {
	#Required
	external_exadata_storage_server_id = var.external_exadata_storage_server_id
}

data "oci_database_management_external_exadata_storage_server_top_sql_cpu_activity" "test_external_exadata_storage_server_top_sql_cpu_activity" {
	#Required
	external_exadata_storage_server_id = var.external_exadata_storage_server_id
}

data "oci_database_management_external_exadata_storage_server" "test_external_exadata_storage_server" {
	#Required
	external_exadata_storage_server_id = var.external_exadata_storage_server_id
}

data "oci_database_management_external_exadata_storage_servers" "test_external_exadata_storage_servers" {
	#Required
	compartment_id = var.compartment_id
	external_exadata_infrastructure_id = oci_database_management_external_exadata_infrastructure.test_external_exadata_infrastructure.id

	#Optional
	display_name = var.external_exadata_storage_server_display_name
}

resource "oci_database_management_external_exadata_infrastructure_exadata_management" "test_external_exadata_infrastructure_exadata_management" {
  #Required
  external_exadata_infrastructure_id = oci_database_management_external_exadata_infrastructure.test_external_exadata_infrastructure.id
  enable_exadata = var.enable_exadata

  #Optional
  license_model = var.external_exadata_infrastructure_database_managements_management_license_model
}

# List managed MySQL database resources in a compartment
data "oci_database_management_managed_my_sql_databases" "test_managed_my_sql_databases" {
  #Required
  compartment_id = var.compartment_id
}

# Get managed MySQL database resource
data "oci_database_management_managed_my_sql_database" "test_managed_my_sql_database" {
  #Required
  managed_my_sql_database_id = data.oci_database_management_managed_my_sql_databases.test_managed_my_sql_databases.managed_my_sql_database_collection.0.items.0.id
}

# Get configuration data for a managed MySQL database resource
data "oci_database_management_managed_my_sql_database_configuration_data" "test_managed_my_sql_database_configuration_data" {
  #Required
  managed_my_sql_database_id = data.oci_database_management_managed_my_sql_databases.test_managed_my_sql_databases.managed_my_sql_database_collection.0.items.0.id
}

# Get SQL data for a managed MySQL database resource
data "oci_database_management_managed_my_sql_database_sql_data" "test_managed_my_sql_database_sql_data" {
  #Required
  managed_my_sql_database_id = data.oci_database_management_managed_my_sql_databases.test_managed_my_sql_databases.managed_my_sql_database_collection.0.items.0.id
  filter_column = "COUNT_STAR"
  start_time = replace(timeadd(timestamp(), "-2h"), "/Z/", ".000Z")
  end_time = replace(timestamp(), "/Z/", ".000Z")
}
