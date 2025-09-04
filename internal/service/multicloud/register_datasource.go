// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package multicloud

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_multicloud_external_location_mapping_metadata", MulticloudExternalLocationMappingMetadataDataSource())
	tfresource.RegisterDatasource("oci_multicloud_external_location_summaries_metadata", MulticloudExternalLocationSummariesMetadataDataSource())
	tfresource.RegisterDatasource("oci_multicloud_external_locations_metadata", MulticloudExternalLocationsMetadataDataSource())
	tfresource.RegisterDatasource("oci_multicloud_network_anchor", MulticloudNetworkAnchorDataSource())
	tfresource.RegisterDatasource("oci_multicloud_network_anchors", MulticloudNetworkAnchorsDataSource())
	tfresource.RegisterDatasource("oci_multicloud_om_hub_multi_cloud_metadata", MulticloudOmHubMultiCloudMetadataDataSource())
	tfresource.RegisterDatasource("oci_multicloud_om_hub_multi_clouds_metadata", MulticloudOmHubMultiCloudsMetadataDataSource())
	tfresource.RegisterDatasource("oci_multicloud_resource_anchor", MulticloudResourceAnchorDataSource())
	tfresource.RegisterDatasource("oci_multicloud_resource_anchors", MulticloudResourceAnchorsDataSource())
}
