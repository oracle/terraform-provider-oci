// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opsi

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_opsi_awr_hub", OpsiAwrHubDataSource())
	tfresource.RegisterDatasource("oci_opsi_awr_hub_awr_snapshot", OpsiAwrHubAwrSnapshotDataSource())
	tfresource.RegisterDatasource("oci_opsi_awr_hub_awr_snapshots", OpsiAwrHubAwrSnapshotsDataSource())
	tfresource.RegisterDatasource("oci_opsi_awr_hub_awr_sources_summary", OpsiAwrHubAwrSourcesSummaryDataSource())
	tfresource.RegisterDatasource("oci_opsi_awr_hub_source", OpsiAwrHubSourceDataSource())
	tfresource.RegisterDatasource("oci_opsi_awr_hub_sources", OpsiAwrHubSourcesDataSource())
	tfresource.RegisterDatasource("oci_opsi_awr_hubs", OpsiAwrHubsDataSource())
	tfresource.RegisterDatasource("oci_opsi_database_insight", OpsiDatabaseInsightDataSource())
	tfresource.RegisterDatasource("oci_opsi_database_insights", OpsiDatabaseInsightsDataSource())
	tfresource.RegisterDatasource("oci_opsi_enterprise_manager_bridge", OpsiEnterpriseManagerBridgeDataSource())
	tfresource.RegisterDatasource("oci_opsi_enterprise_manager_bridges", OpsiEnterpriseManagerBridgesDataSource())
	tfresource.RegisterDatasource("oci_opsi_exadata_insight", OpsiExadataInsightDataSource())
	tfresource.RegisterDatasource("oci_opsi_exadata_insights", OpsiExadataInsightsDataSource())
	tfresource.RegisterDatasource("oci_opsi_host_insight", OpsiHostInsightDataSource())
	tfresource.RegisterDatasource("oci_opsi_host_insights", OpsiHostInsightsDataSource())
	tfresource.RegisterDatasource("oci_opsi_importable_compute_entities", OpsiImportableComputeEntitiesDataSource())
	tfresource.RegisterDatasource("oci_opsi_news_report", OpsiNewsReportDataSource())
	tfresource.RegisterDatasource("oci_opsi_news_reports", OpsiNewsReportsDataSource())
	tfresource.RegisterDatasource("oci_opsi_importable_compute_entity", OpsiImportableComputeEntityDataSource())
	tfresource.RegisterDatasource("oci_opsi_operations_insights_private_endpoint", OpsiOperationsInsightsPrivateEndpointDataSource())
	tfresource.RegisterDatasource("oci_opsi_operations_insights_private_endpoints", OpsiOperationsInsightsPrivateEndpointsDataSource())
	tfresource.RegisterDatasource("oci_opsi_operations_insights_warehouse", OpsiOperationsInsightsWarehouseDataSource())
	tfresource.RegisterDatasource("oci_opsi_operations_insights_warehouse_resource_usage_summary", OpsiOperationsInsightsWarehouseResourceUsageSummaryDataSource())
	tfresource.RegisterDatasource("oci_opsi_operations_insights_warehouse_user", OpsiOperationsInsightsWarehouseUserDataSource())
	tfresource.RegisterDatasource("oci_opsi_operations_insights_warehouse_users", OpsiOperationsInsightsWarehouseUsersDataSource())
	tfresource.RegisterDatasource("oci_opsi_operations_insights_warehouses", OpsiOperationsInsightsWarehousesDataSource())
	tfresource.RegisterDatasource("oci_opsi_opsi_configuration", OpsiOpsiConfigurationDataSource())
	tfresource.RegisterDatasource("oci_opsi_opsi_configuration_configuration_item", OpsiOpsiConfigurationConfigurationItemDataSource())
	tfresource.RegisterDatasource("oci_opsi_opsi_configurations", OpsiOpsiConfigurationsDataSource())
	tfresource.RegisterDatasource("oci_opsi_importable_agent_entities", OpsiImportableAgentEntitiesDataSource())
	tfresource.RegisterDatasource("oci_opsi_importable_agent_entity", OpsiImportableAgentEntityDataSource())
}
