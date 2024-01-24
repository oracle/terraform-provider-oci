// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package mysql

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_mysql_channel", MysqlChannelResource())
	tfresource.RegisterResource("oci_mysql_heat_wave_cluster", MysqlHeatWaveClusterResource())
	tfresource.RegisterResource("oci_mysql_mysql_backup", MysqlMysqlBackupResource())
	tfresource.RegisterResource("oci_mysql_mysql_configuration", MysqlMysqlConfigurationResource())
	tfresource.RegisterResource("oci_mysql_mysql_db_system", MysqlMysqlDbSystemResource())
	tfresource.RegisterResource("oci_mysql_replica", MysqlReplicaResource())
}
