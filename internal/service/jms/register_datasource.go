// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_jms_fleet", JmsFleetDataSource())
	tfresource.RegisterDatasource("oci_jms_fleet_blocklists", JmsFleetBlocklistsDataSource())
	tfresource.RegisterDatasource("oci_jms_fleet_installation_site", JmsFleetInstallationSiteDataSource())
	tfresource.RegisterDatasource("oci_jms_fleet_installation_sites", JmsFleetInstallationSitesDataSource())
	tfresource.RegisterDatasource("oci_jms_fleets", JmsFleetsDataSource())
	tfresource.RegisterDatasource("oci_jms_java_families", JmsJavaFamiliesDataSource())
	tfresource.RegisterDatasource("oci_jms_java_family", JmsJavaFamilyDataSource())
	tfresource.RegisterDatasource("oci_jms_java_release", JmsJavaReleaseDataSource())
	tfresource.RegisterDatasource("oci_jms_java_releases", JmsJavaReleasesDataSource())
	tfresource.RegisterDatasource("oci_jms_list_jre_usage", JmsListJreUsageDataSource())
	tfresource.RegisterDatasource("oci_jms_summarize_resource_inventory", JmsSummarizeResourceInventoryDataSource())
}
