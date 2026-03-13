// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package distributed_database

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_distributed_database_distributed_autonomous_database", DistributedDatabaseDistributedAutonomousDatabaseDataSource())
	tfresource.RegisterDatasource("oci_distributed_database_distributed_autonomous_databases", DistributedDatabaseDistributedAutonomousDatabasesDataSource())
	tfresource.RegisterDatasource("oci_distributed_database_distributed_database", DistributedDatabaseDistributedDatabaseDataSource())
	tfresource.RegisterDatasource("oci_distributed_database_distributed_database_private_endpoint", DistributedDatabaseDistributedDatabasePrivateEndpointDataSource())
	tfresource.RegisterDatasource("oci_distributed_database_distributed_database_private_endpoints", DistributedDatabaseDistributedDatabasePrivateEndpointsDataSource())
	tfresource.RegisterDatasource("oci_distributed_database_distributed_databases", DistributedDatabaseDistributedDatabasesDataSource())
}
