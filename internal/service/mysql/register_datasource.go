// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package mysql

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_mysql_channel", MysqlChannelDataSource())
	tfresource.RegisterDatasource("oci_mysql_channels", MysqlChannelsDataSource())
	tfresource.RegisterDatasource("oci_mysql_heat_wave_cluster", MysqlHeatWaveClusterDataSource())
	tfresource.RegisterDatasource("oci_mysql_mysql_backup", MysqlMysqlBackupDataSource())
	tfresource.RegisterDatasource("oci_mysql_mysql_backups", MysqlMysqlBackupsDataSource())
	tfresource.RegisterDatasource("oci_mysql_mysql_configuration", MysqlMysqlConfigurationDataSource())
	tfresource.RegisterDatasource("oci_mysql_mysql_configurations", MysqlMysqlConfigurationsDataSource())
	tfresource.RegisterDatasource("oci_mysql_mysql_db_system", MysqlMysqlDbSystemDataSource())
	tfresource.RegisterDatasource("oci_mysql_mysql_db_systems", MysqlMysqlDbSystemsDataSource())
	tfresource.RegisterDatasource("oci_mysql_mysql_versions", MysqlMysqlVersionsDataSource())
	tfresource.RegisterDatasource("oci_mysql_replica", MysqlReplicaDataSource())
	tfresource.RegisterDatasource("oci_mysql_replicas", MysqlReplicasDataSource())
	tfresource.RegisterDatasource("oci_mysql_shapes", MysqlShapesDataSource())
}
