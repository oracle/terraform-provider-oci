// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_jms_announcements", JmsAnnouncementsDataSource())
	tfresource.RegisterDatasource("oci_jms_fleet", JmsFleetDataSource())
	tfresource.RegisterDatasource("oci_jms_fleet_advanced_feature_configuration", JmsFleetAdvancedFeatureConfigurationDataSource())
	tfresource.RegisterDatasource("oci_jms_fleet_blocklists", JmsFleetBlocklistsDataSource())
	tfresource.RegisterDatasource("oci_jms_fleet_crypto_analysis_result", JmsFleetCryptoAnalysisResultDataSource())
	tfresource.RegisterDatasource("oci_jms_fleet_crypto_analysis_results", JmsFleetCryptoAnalysisResultsDataSource())
	tfresource.RegisterDatasource("oci_jms_fleet_diagnoses", JmsFleetDiagnosesDataSource())
	tfresource.RegisterDatasource("oci_jms_fleet_drs_file", JmsFleetDrsFileDataSource())
	tfresource.RegisterDatasource("oci_jms_fleet_drs_files", JmsFleetDrsFilesDataSource())
	tfresource.RegisterDatasource("oci_jms_fleet_export_setting", JmsFleetExportSettingDataSource())
	tfresource.RegisterDatasource("oci_jms_fleet_export_status", JmsFleetExportStatusDataSource())
	tfresource.RegisterDatasource("oci_jms_fleet_installation_site", JmsFleetInstallationSiteDataSource())
	tfresource.RegisterDatasource("oci_jms_fleet_installation_sites", JmsFleetInstallationSitesDataSource())
	tfresource.RegisterDatasource("oci_jms_fleet_java_migration_analysis_result", JmsFleetJavaMigrationAnalysisResultDataSource())
	tfresource.RegisterDatasource("oci_jms_fleet_java_migration_analysis_results", JmsFleetJavaMigrationAnalysisResultsDataSource())
	tfresource.RegisterDatasource("oci_jms_fleet_performance_tuning_analysis_result", JmsFleetPerformanceTuningAnalysisResultDataSource())
	tfresource.RegisterDatasource("oci_jms_fleet_performance_tuning_analysis_results", JmsFleetPerformanceTuningAnalysisResultsDataSource())
	tfresource.RegisterDatasource("oci_jms_fleets", JmsFleetsDataSource())
	tfresource.RegisterDatasource("oci_jms_java_families", JmsJavaFamiliesDataSource())
	tfresource.RegisterDatasource("oci_jms_java_family", JmsJavaFamilyDataSource())
	tfresource.RegisterDatasource("oci_jms_java_release", JmsJavaReleaseDataSource())
	tfresource.RegisterDatasource("oci_jms_java_releases", JmsJavaReleasesDataSource())
	tfresource.RegisterDatasource("oci_jms_list_jre_usage", JmsListJreUsageDataSource())
	tfresource.RegisterDatasource("oci_jms_summarize_resource_inventory", JmsSummarizeResourceInventoryDataSource())
}
