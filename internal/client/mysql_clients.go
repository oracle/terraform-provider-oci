// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_mysql "github.com/oracle/oci-go-sdk/v65/mysql"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_mysql.ChannelsClient", &OracleClient{InitClientFn: initMysqlChannelsClient})
	RegisterOracleClient("oci_mysql.DbBackupsClient", &OracleClient{InitClientFn: initMysqlDbBackupsClient})
	RegisterOracleClient("oci_mysql.DbSystemClient", &OracleClient{InitClientFn: initMysqlDbSystemClient})
	RegisterOracleClient("oci_mysql.MysqlaasClient", &OracleClient{InitClientFn: initMysqlMysqlaasClient})
	RegisterOracleClient("oci_mysql.ReplicasClient", &OracleClient{InitClientFn: initMysqlReplicasClient})
}

func initMysqlChannelsClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_mysql.NewChannelsClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.HostUrlOverride != "" {
		client.Host = serviceClientOverrides.HostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) ChannelsClient() *oci_mysql.ChannelsClient {
	return m.GetClient("oci_mysql.ChannelsClient").(*oci_mysql.ChannelsClient)
}

func initMysqlDbBackupsClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_mysql.NewDbBackupsClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.HostUrlOverride != "" {
		client.Host = serviceClientOverrides.HostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) DbBackupsClient() *oci_mysql.DbBackupsClient {
	return m.GetClient("oci_mysql.DbBackupsClient").(*oci_mysql.DbBackupsClient)
}

func initMysqlDbSystemClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_mysql.NewDbSystemClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.HostUrlOverride != "" {
		client.Host = serviceClientOverrides.HostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) DbSystemClient() *oci_mysql.DbSystemClient {
	return m.GetClient("oci_mysql.DbSystemClient").(*oci_mysql.DbSystemClient)
}

func initMysqlMysqlaasClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_mysql.NewMysqlaasClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.HostUrlOverride != "" {
		client.Host = serviceClientOverrides.HostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) MysqlaasClient() *oci_mysql.MysqlaasClient {
	return m.GetClient("oci_mysql.MysqlaasClient").(*oci_mysql.MysqlaasClient)
}

func initMysqlReplicasClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_mysql.NewReplicasClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.HostUrlOverride != "" {
		client.Host = serviceClientOverrides.HostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) ReplicasClient() *oci_mysql.ReplicasClient {
	return m.GetClient("oci_mysql.ReplicasClient").(*oci_mysql.ReplicasClient)
}
