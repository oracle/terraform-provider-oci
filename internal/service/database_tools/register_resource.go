// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_tools

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_database_tools_database_tools_connection", DatabaseToolsDatabaseToolsConnectionResource())
	tfresource.RegisterResource("oci_database_tools_database_tools_database_api_gateway_config", DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigResource())
	tfresource.RegisterResource("oci_database_tools_database_tools_identity", DatabaseToolsDatabaseToolsIdentityResource())
	tfresource.RegisterResource("oci_database_tools_database_tools_mcp_server", DatabaseToolsDatabaseToolsMcpServerResource())
	tfresource.RegisterResource("oci_database_tools_database_tools_mcp_toolset", DatabaseToolsDatabaseToolsMcpToolsetResource())
	tfresource.RegisterResource("oci_database_tools_database_tools_private_endpoint", DatabaseToolsDatabaseToolsPrivateEndpointResource())
	tfresource.RegisterResource("oci_database_tools_database_tools_sql_report", DatabaseToolsDatabaseToolsSqlReportResource())
}
