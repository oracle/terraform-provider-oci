// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package osmanagement

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_osmanagement_managed_instance", OsmanagementManagedInstanceDataSource())
	tfresource.RegisterDatasource("oci_osmanagement_managed_instance_event_report", OsmanagementManagedInstanceEventReportDataSource())
	tfresource.RegisterDatasource("oci_osmanagement_managed_instance_group", OsmanagementManagedInstanceGroupDataSource())
	tfresource.RegisterDatasource("oci_osmanagement_managed_instance_groups", OsmanagementManagedInstanceGroupsDataSource())
	tfresource.RegisterDatasource("oci_osmanagement_managed_instance_module_streams", OsmanagementManagedInstanceModuleStreamsDataSource())
	tfresource.RegisterDatasource("oci_osmanagement_managed_instance_stream_profiles", OsmanagementManagedInstanceStreamProfilesDataSource())
	tfresource.RegisterDatasource("oci_osmanagement_managed_instances", OsmanagementManagedInstancesDataSource())
	tfresource.RegisterDatasource("oci_osmanagement_software_source", OsmanagementSoftwareSourceDataSource())
	tfresource.RegisterDatasource("oci_osmanagement_software_source_module_stream", OsmanagementSoftwareSourceModuleStreamDataSource())
	tfresource.RegisterDatasource("oci_osmanagement_software_source_module_stream_profile", OsmanagementSoftwareSourceModuleStreamProfileDataSource())
	tfresource.RegisterDatasource("oci_osmanagement_software_source_stream_profiles", OsmanagementSoftwareSourceStreamProfilesDataSource())
	tfresource.RegisterDatasource("oci_osmanagement_software_sources", OsmanagementSoftwareSourcesDataSource())
}
