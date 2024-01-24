// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_bridge

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_cloud_bridge_agent", CloudBridgeAgentResource())
	tfresource.RegisterResource("oci_cloud_bridge_agent_dependency", CloudBridgeAgentDependencyResource())
	tfresource.RegisterResource("oci_cloud_bridge_agent_plugin", CloudBridgeAgentPluginResource())
	tfresource.RegisterResource("oci_cloud_bridge_asset", CloudBridgeAssetResource())
	tfresource.RegisterResource("oci_cloud_bridge_asset_source", CloudBridgeAssetSourceResource())
	tfresource.RegisterResource("oci_cloud_bridge_discovery_schedule", CloudBridgeDiscoveryScheduleResource())
	tfresource.RegisterResource("oci_cloud_bridge_environment", CloudBridgeEnvironmentResource())
	tfresource.RegisterResource("oci_cloud_bridge_inventory", CloudBridgeInventoryResource())
}
