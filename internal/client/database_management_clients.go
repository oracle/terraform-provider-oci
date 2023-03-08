// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_database_management.DbManagementClient", &OracleClient{InitClientFn: initDatabasemanagementDbManagementClient})
	RegisterOracleClient("oci_database_management.ManagedMySqlDatabasesClient", &OracleClient{InitClientFn: initDatabasemanagementManagedMySqlDatabasesClient})
	RegisterOracleClient("oci_database_management.DiagnosabilityClient", &OracleClient{InitClientFn: initDatabasemanagementDiagnosabilityClient})
	RegisterOracleClient("oci_database_management.SqlTuningClient", &OracleClient{InitClientFn: initDatabasemanagementSqlTuningClient})
}

func initDatabasemanagementDbManagementClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_database_management.NewDbManagementClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) DbManagementClient() *oci_database_management.DbManagementClient {
	return m.GetClient("oci_database_management.DbManagementClient").(*oci_database_management.DbManagementClient)
}

func initDatabasemanagementDiagnosabilityClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_database_management.NewDiagnosabilityClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) DiagnosabilityClient() *oci_database_management.DiagnosabilityClient {
	return m.GetClient("oci_database_management.DiagnosabilityClient").(*oci_database_management.DiagnosabilityClient)
}

func initDatabasemanagementManagedMySqlDatabasesClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_database_management.NewManagedMySqlDatabasesClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) ManagedMySqlDatabasesClient() *oci_database_management.ManagedMySqlDatabasesClient {
	return m.GetClient("oci_database_management.ManagedMySqlDatabasesClient").(*oci_database_management.ManagedMySqlDatabasesClient)
}

func initDatabasemanagementSqlTuningClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_database_management.NewSqlTuningClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) SqlTuningClient() *oci_database_management.SqlTuningClient {
	return m.GetClient("oci_database_management.SqlTuningClient").(*oci_database_management.SqlTuningClient)
}
