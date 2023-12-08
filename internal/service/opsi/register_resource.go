// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opsi

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_opsi_awr_hub", OpsiAwrHubResource())
	tfresource.RegisterResource("oci_opsi_awr_hub_source", OpsiAwrHubSourceResource())
	tfresource.RegisterResource("oci_opsi_awr_hub_source_awrhubsources_management", OpsiAwrHubSourceAwrhubsourcesManagementResource())
	tfresource.RegisterResource("oci_opsi_database_insight", OpsiDatabaseInsightResource())
	tfresource.RegisterResource("oci_opsi_enterprise_manager_bridge", OpsiEnterpriseManagerBridgeResource())
	tfresource.RegisterResource("oci_opsi_exadata_insight", OpsiExadataInsightResource())
	tfresource.RegisterResource("oci_opsi_host_insight", OpsiHostInsightResource())
	tfresource.RegisterResource("oci_opsi_news_report", OpsiNewsReportResource())
	tfresource.RegisterResource("oci_opsi_operations_insights_private_endpoint", OpsiOperationsInsightsPrivateEndpointResource())
	tfresource.RegisterResource("oci_opsi_operations_insights_warehouse", OpsiOperationsInsightsWarehouseResource())
	tfresource.RegisterResource("oci_opsi_operations_insights_warehouse_download_warehouse_wallet", OpsiOperationsInsightsWarehouseDownloadWarehouseWalletResource())
	tfresource.RegisterResource("oci_opsi_operations_insights_warehouse_rotate_warehouse_wallet", OpsiOperationsInsightsWarehouseRotateWarehouseWalletResource())
	tfresource.RegisterResource("oci_opsi_operations_insights_warehouse_user", OpsiOperationsInsightsWarehouseUserResource())
	tfresource.RegisterResource("oci_opsi_opsi_configuration", OpsiOpsiConfigurationResource())
}
