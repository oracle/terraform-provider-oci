// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_database_management_db_management_private_endpoint", DatabaseManagementDbManagementPrivateEndpointResource())
	tfresource.RegisterResource("oci_database_management_managed_database_group", DatabaseManagementManagedDatabaseGroupResource())
	tfresource.RegisterResource("oci_database_management_managed_databases_change_database_parameter", DatabaseManagementManagedDatabasesChangeDatabaseParameterResource())
	tfresource.RegisterResource("oci_database_management_managed_databases_reset_database_parameter", DatabaseManagementManagedDatabasesResetDatabaseParameterResource())
}
