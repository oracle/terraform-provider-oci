// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_mysql "github.com/oracle/oci-go-sdk/v26/mysql"

	oci_common "github.com/oracle/oci-go-sdk/v26/common"
)

func init() {
	RegisterOracleClient("oci_mysql.DbBackupsClient", &OracleClient{initClientFn: initMysqlDbBackupsClient})
	RegisterOracleClient("oci_mysql.DbSystemClient", &OracleClient{initClientFn: initMysqlDbSystemClient})
	RegisterOracleClient("oci_mysql.MysqlaasClient", &OracleClient{initClientFn: initMysqlMysqlaasClient})
}

func initMysqlDbBackupsClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_mysql.NewDbBackupsClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) dbBackupsClient() *oci_mysql.DbBackupsClient {
	return m.GetClient("oci_mysql.DbBackupsClient").(*oci_mysql.DbBackupsClient)
}

func initMysqlDbSystemClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_mysql.NewDbSystemClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) dbSystemClient() *oci_mysql.DbSystemClient {
	return m.GetClient("oci_mysql.DbSystemClient").(*oci_mysql.DbSystemClient)
}

func initMysqlMysqlaasClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_mysql.NewMysqlaasClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) mysqlaasClient() *oci_mysql.MysqlaasClient {
	return m.GetClient("oci_mysql.MysqlaasClient").(*oci_mysql.MysqlaasClient)
}
