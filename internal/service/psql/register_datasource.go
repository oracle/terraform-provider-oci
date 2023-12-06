// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package psql

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_psql_backup", PsqlBackupDataSource())
	tfresource.RegisterDatasource("oci_psql_backups", PsqlBackupsDataSource())
	tfresource.RegisterDatasource("oci_psql_configuration", PsqlConfigurationDataSource())
	tfresource.RegisterDatasource("oci_psql_configurations", PsqlConfigurationsDataSource())
	tfresource.RegisterDatasource("oci_psql_db_system", PsqlDbSystemDataSource())
	tfresource.RegisterDatasource("oci_psql_db_system_connection_detail", PsqlDbSystemConnectionDetailDataSource())
	tfresource.RegisterDatasource("oci_psql_db_system_primary_db_instance", PsqlDbSystemPrimaryDbInstanceDataSource())
	tfresource.RegisterDatasource("oci_psql_db_systems", PsqlDbSystemsDataSource())
	tfresource.RegisterDatasource("oci_psql_default_configuration", PsqlDefaultConfigurationDataSource())
	tfresource.RegisterDatasource("oci_psql_default_configurations", PsqlDefaultConfigurationsDataSource())
	tfresource.RegisterDatasource("oci_psql_shapes", PsqlShapesDataSource())
}
