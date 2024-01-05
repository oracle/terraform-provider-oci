// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_tools

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_database_tools_database_tools_connection", DatabaseToolsDatabaseToolsConnectionDataSource())
	tfresource.RegisterDatasource("oci_database_tools_database_tools_connections", DatabaseToolsDatabaseToolsConnectionsDataSource())
	tfresource.RegisterDatasource("oci_database_tools_database_tools_endpoint_service", DatabaseToolsDatabaseToolsEndpointServiceDataSource())
	tfresource.RegisterDatasource("oci_database_tools_database_tools_endpoint_services", DatabaseToolsDatabaseToolsEndpointServicesDataSource())
	tfresource.RegisterDatasource("oci_database_tools_database_tools_private_endpoint", DatabaseToolsDatabaseToolsPrivateEndpointDataSource())
	tfresource.RegisterDatasource("oci_database_tools_database_tools_private_endpoints", DatabaseToolsDatabaseToolsPrivateEndpointsDataSource())
}
