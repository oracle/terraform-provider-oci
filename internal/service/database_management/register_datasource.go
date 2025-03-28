// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_database_management_db_management_private_endpoint", DatabaseManagementDbManagementPrivateEndpointDataSource())
	tfresource.RegisterDatasource("oci_database_management_db_management_private_endpoint_associated_database", DatabaseManagementDbManagementPrivateEndpointAssociatedDatabaseDataSource())
	tfresource.RegisterDatasource("oci_database_management_db_management_private_endpoint_associated_databases", DatabaseManagementDbManagementPrivateEndpointAssociatedDatabasesDataSource())
	tfresource.RegisterDatasource("oci_database_management_db_management_private_endpoints", DatabaseManagementDbManagementPrivateEndpointsDataSource())
	tfresource.RegisterDatasource("oci_database_management_external_asm", DatabaseManagementExternalAsmDataSource())
	tfresource.RegisterDatasource("oci_database_management_external_asm_configuration", DatabaseManagementExternalAsmConfigurationDataSource())
	tfresource.RegisterDatasource("oci_database_management_external_asm_disk_groups", DatabaseManagementExternalAsmDiskGroupsDataSource())
	tfresource.RegisterDatasource("oci_database_management_external_asm_instance", DatabaseManagementExternalAsmInstanceDataSource())
	tfresource.RegisterDatasource("oci_database_management_external_asm_instances", DatabaseManagementExternalAsmInstancesDataSource())
	tfresource.RegisterDatasource("oci_database_management_external_asm_users", DatabaseManagementExternalAsmUsersDataSource())
	tfresource.RegisterDatasource("oci_database_management_external_asms", DatabaseManagementExternalAsmsDataSource())
	tfresource.RegisterDatasource("oci_database_management_external_cluster", DatabaseManagementExternalClusterDataSource())
	tfresource.RegisterDatasource("oci_database_management_external_cluster_instance", DatabaseManagementExternalClusterInstanceDataSource())
	tfresource.RegisterDatasource("oci_database_management_external_cluster_instances", DatabaseManagementExternalClusterInstancesDataSource())
	tfresource.RegisterDatasource("oci_database_management_external_clusters", DatabaseManagementExternalClustersDataSource())
	tfresource.RegisterDatasource("oci_database_management_external_databases", DatabaseManagementExternalDatabasesDataSource())
	tfresource.RegisterDatasource("oci_database_management_external_db_home", DatabaseManagementExternalDbHomeDataSource())
	tfresource.RegisterDatasource("oci_database_management_external_db_homes", DatabaseManagementExternalDbHomesDataSource())
	tfresource.RegisterDatasource("oci_database_management_external_db_node", DatabaseManagementExternalDbNodeDataSource())
	tfresource.RegisterDatasource("oci_database_management_external_db_nodes", DatabaseManagementExternalDbNodesDataSource())
	tfresource.RegisterDatasource("oci_database_management_external_db_system", DatabaseManagementExternalDbSystemDataSource())
	tfresource.RegisterDatasource("oci_database_management_external_db_system_connector", DatabaseManagementExternalDbSystemConnectorDataSource())
	tfresource.RegisterDatasource("oci_database_management_external_db_system_connectors", DatabaseManagementExternalDbSystemConnectorsDataSource())
	tfresource.RegisterDatasource("oci_database_management_external_db_system_discoveries", DatabaseManagementExternalDbSystemDiscoveriesDataSource())
	tfresource.RegisterDatasource("oci_database_management_external_db_system_discovery", DatabaseManagementExternalDbSystemDiscoveryDataSource())
	tfresource.RegisterDatasource("oci_database_management_external_db_systems", DatabaseManagementExternalDbSystemsDataSource())
	tfresource.RegisterDatasource("oci_database_management_external_exadata_infrastructure", DatabaseManagementExternalExadataInfrastructureDataSource())
	tfresource.RegisterDatasource("oci_database_management_external_exadata_infrastructures", DatabaseManagementExternalExadataInfrastructuresDataSource())
	tfresource.RegisterDatasource("oci_database_management_external_exadata_storage_connector", DatabaseManagementExternalExadataStorageConnectorDataSource())
	tfresource.RegisterDatasource("oci_database_management_external_exadata_storage_connectors", DatabaseManagementExternalExadataStorageConnectorsDataSource())
	tfresource.RegisterDatasource("oci_database_management_external_exadata_storage_grid", DatabaseManagementExternalExadataStorageGridDataSource())
	tfresource.RegisterDatasource("oci_database_management_external_exadata_storage_server", DatabaseManagementExternalExadataStorageServerDataSource())
	tfresource.RegisterDatasource("oci_database_management_external_exadata_storage_server_iorm_plan", DatabaseManagementExternalExadataStorageServerIormPlanDataSource())
	tfresource.RegisterDatasource("oci_database_management_external_exadata_storage_server_open_alert_history", DatabaseManagementExternalExadataStorageServerOpenAlertHistoryDataSource())
	tfresource.RegisterDatasource("oci_database_management_external_exadata_storage_server_top_sql_cpu_activity", DatabaseManagementExternalExadataStorageServerTopSqlCpuActivityDataSource())
	tfresource.RegisterDatasource("oci_database_management_external_exadata_storage_servers", DatabaseManagementExternalExadataStorageServersDataSource())
	tfresource.RegisterDatasource("oci_database_management_external_listener", DatabaseManagementExternalListenerDataSource())
	tfresource.RegisterDatasource("oci_database_management_external_listener_services", DatabaseManagementExternalListenerServicesDataSource())
	tfresource.RegisterDatasource("oci_database_management_external_listeners", DatabaseManagementExternalListenersDataSource())
	tfresource.RegisterDatasource("oci_database_management_external_my_sql_database", DatabaseManagementExternalMySqlDatabaseDataSource())
	tfresource.RegisterDatasource("oci_database_management_external_my_sql_database_connector", DatabaseManagementExternalMySqlDatabaseConnectorDataSource())
	tfresource.RegisterDatasource("oci_database_management_external_my_sql_database_connectors", DatabaseManagementExternalMySqlDatabaseConnectorsDataSource())
	tfresource.RegisterDatasource("oci_database_management_external_my_sql_databases", DatabaseManagementExternalMySqlDatabasesDataSource())
	tfresource.RegisterDatasource("oci_database_management_job_executions_status", DatabaseManagementJobExecutionsStatusDataSource())
	tfresource.RegisterDatasource("oci_database_management_job_executions_statuses", DatabaseManagementJobExecutionsStatusesDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database", DatabaseManagementManagedDatabaseDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_addm_task", DatabaseManagementManagedDatabaseAddmTaskDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_addm_tasks", DatabaseManagementManagedDatabaseAddmTasksDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_alert_log_count", DatabaseManagementManagedDatabaseAlertLogCountDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_alert_log_counts", DatabaseManagementManagedDatabaseAlertLogCountsDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_attention_log_count", DatabaseManagementManagedDatabaseAttentionLogCountDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_attention_log_counts", DatabaseManagementManagedDatabaseAttentionLogCountsDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_cursor_cache_statements", DatabaseManagementManagedDatabaseCursorCacheStatementsDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_group", DatabaseManagementManagedDatabaseGroupDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_groups", DatabaseManagementManagedDatabaseGroupsDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_sql_plan_baseline", DatabaseManagementManagedDatabaseSqlPlanBaselineDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_sql_plan_baseline_configuration", DatabaseManagementManagedDatabaseSqlPlanBaselineConfigurationDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_sql_plan_baseline_jobs", DatabaseManagementManagedDatabaseSqlPlanBaselineJobsDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_sql_plan_baselines", DatabaseManagementManagedDatabaseSqlPlanBaselinesDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_sql_tuning_advisor_task", DatabaseManagementManagedDatabaseSqlTuningAdvisorTaskDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_sql_tuning_advisor_tasks", DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_sql_tuning_advisor_tasks_execution_plan_stats_comparision", DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksExecutionPlanStatsComparisionDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_sql_tuning_advisor_tasks_finding", DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksFindingDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_sql_tuning_advisor_tasks_findings", DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksFindingsDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_sql_tuning_advisor_tasks_recommendation", DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksRecommendationDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_sql_tuning_advisor_tasks_recommendations", DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksRecommendationsDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_sql_tuning_advisor_tasks_sql_execution_plan", DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSqlExecutionPlanDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_sql_tuning_advisor_tasks_summary_report", DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSummaryReportDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_sql_tuning_set", DatabaseManagementManagedDatabaseSqlTuningSetDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_sql_tuning_sets", DatabaseManagementManagedDatabaseSqlTuningSetsDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_user", DatabaseManagementManagedDatabaseUserDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_user_consumer_group_privilege", DatabaseManagementManagedDatabaseUserConsumerGroupPrivilegeDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_user_consumer_group_privileges", DatabaseManagementManagedDatabaseUserConsumerGroupPrivilegesDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_user_data_access_container", DatabaseManagementManagedDatabaseUserDataAccessContainerDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_user_data_access_containers", DatabaseManagementManagedDatabaseUserDataAccessContainersDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_user_object_privilege", DatabaseManagementManagedDatabaseUserObjectPrivilegeDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_user_object_privileges", DatabaseManagementManagedDatabaseUserObjectPrivilegesDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_user_proxied_for_user", DatabaseManagementManagedDatabaseUserProxiedForUserDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_user_proxied_for_users", DatabaseManagementManagedDatabaseUserProxiedForUsersDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_user_role", DatabaseManagementManagedDatabaseUserRoleDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_user_roles", DatabaseManagementManagedDatabaseUserRolesDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_users", DatabaseManagementManagedDatabaseUsersDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_databases", DatabaseManagementManagedDatabasesDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_databases_asm_properties", DatabaseManagementManagedDatabasesAsmPropertiesDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_databases_asm_property", DatabaseManagementManagedDatabasesAsmPropertyDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_databases_database_parameter", DatabaseManagementManagedDatabasesDatabaseParameterDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_databases_database_parameters", DatabaseManagementManagedDatabasesDatabaseParametersDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_databases_user_proxy_user", DatabaseManagementManagedDatabasesUserProxyUserDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_databases_user_proxy_users", DatabaseManagementManagedDatabasesUserProxyUsersDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_databases_user_system_privilege", DatabaseManagementManagedDatabasesUserSystemPrivilegeDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_databases_user_system_privileges", DatabaseManagementManagedDatabasesUserSystemPrivilegesDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_my_sql_database", DatabaseManagementManagedMySqlDatabaseDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_my_sql_database_configuration_data", DatabaseManagementManagedMySqlDatabaseConfigurationDataDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_my_sql_database_sql_data", DatabaseManagementManagedMySqlDatabaseSqlDataDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_my_sql_databases", DatabaseManagementManagedMySqlDatabasesDataSource())
	tfresource.RegisterDatasource("oci_database_management_named_credential", DatabaseManagementNamedCredentialDataSource())
	tfresource.RegisterDatasource("oci_database_management_named_credentials", DatabaseManagementNamedCredentialsDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_optimizer_statistics_advisor_execution", DatabaseManagementManagedDatabaseOptimizerStatisticsAdvisorExecutionDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_optimizer_statistics_advisor_execution_script", DatabaseManagementManagedDatabaseOptimizerStatisticsAdvisorExecutionScriptDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_optimizer_statistics_advisor_executions", DatabaseManagementManagedDatabaseOptimizerStatisticsAdvisorExecutionsDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_optimizer_statistics_collection_aggregations", DatabaseManagementManagedDatabaseOptimizerStatisticsCollectionAggregationsDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_optimizer_statistics_collection_operation", DatabaseManagementManagedDatabaseOptimizerStatisticsCollectionOperationDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_optimizer_statistics_collection_operations", DatabaseManagementManagedDatabaseOptimizerStatisticsCollectionOperationsDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_table_statistics", DatabaseManagementManagedDatabaseTableStatisticsDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_preferred_credential", DatabaseManagementManagedDatabasePreferredCredentialDataSource())
	tfresource.RegisterDatasource("oci_database_management_managed_database_preferred_credentials", DatabaseManagementManagedDatabasePreferredCredentialsDataSource())
}
