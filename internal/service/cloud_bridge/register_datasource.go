// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_bridge

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_cloud_bridge_agent", CloudBridgeAgentDataSource())
	tfresource.RegisterDatasource("oci_cloud_bridge_agent_dependencies", CloudBridgeAgentDependenciesDataSource())
	tfresource.RegisterDatasource("oci_cloud_bridge_agent_dependency", CloudBridgeAgentDependencyDataSource())
	tfresource.RegisterDatasource("oci_cloud_bridge_agent_plugin", CloudBridgeAgentPluginDataSource())
	tfresource.RegisterDatasource("oci_cloud_bridge_agents", CloudBridgeAgentsDataSource())
	tfresource.RegisterDatasource("oci_cloud_bridge_appliance_image", CloudBridgeApplianceImageDataSource())
	tfresource.RegisterDatasource("oci_cloud_bridge_appliance_images", CloudBridgeApplianceImagesDataSource())
	tfresource.RegisterDatasource("oci_cloud_bridge_asset", CloudBridgeAssetDataSource())
	tfresource.RegisterDatasource("oci_cloud_bridge_asset_source", CloudBridgeAssetSourceDataSource())
	tfresource.RegisterDatasource("oci_cloud_bridge_asset_sources", CloudBridgeAssetSourcesDataSource())
	tfresource.RegisterDatasource("oci_cloud_bridge_assets", CloudBridgeAssetsDataSource())
	tfresource.RegisterDatasource("oci_cloud_bridge_discovery_schedule", CloudBridgeDiscoveryScheduleDataSource())
	tfresource.RegisterDatasource("oci_cloud_bridge_discovery_schedules", CloudBridgeDiscoverySchedulesDataSource())
	tfresource.RegisterDatasource("oci_cloud_bridge_environment", CloudBridgeEnvironmentDataSource())
	tfresource.RegisterDatasource("oci_cloud_bridge_environments", CloudBridgeEnvironmentsDataSource())
	tfresource.RegisterDatasource("oci_cloud_bridge_inventories", CloudBridgeInventoriesDataSource())
	tfresource.RegisterDatasource("oci_cloud_bridge_inventory", CloudBridgeInventoryDataSource())
}
