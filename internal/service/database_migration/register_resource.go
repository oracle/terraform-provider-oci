// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_migration

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {

	tfresource.RegisterResource("oci_database_migration_connection", DatabaseMigrationConnectionResource())

	//tfresource.RegisterResource("oci_database_migration_agent", DatabaseMigrationAgentResource())
	//tfresource.RegisterResource("oci_database_migration_connection", DatabaseMigrationConnectionResource())

	tfresource.RegisterResource("oci_database_migration_job", DatabaseMigrationJobResource())
	tfresource.RegisterResource("oci_database_migration_migration", DatabaseMigrationMigrationResource())
}
