// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_dbmulticloud "github.com/oracle/oci-go-sdk/v65/dbmulticloud"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_dbmulticloud.DbMulticloudGCPProviderClient", &OracleClient{InitClientFn: initDbmulticloudDbMulticloudGCPProviderClient})
	RegisterOracleClient("oci_dbmulticloud.WorkRequestClient", &OracleClient{InitClientFn: initDbmulticloudWorkRequestClient})
	RegisterOracleClient("oci_dbmulticloud.MultiCloudResourceDiscoveryClient", &OracleClient{InitClientFn: initDbmulticloudMultiCloudResourceDiscoveryClient})
	RegisterOracleClient("oci_dbmulticloud.OracleDBAzureBlobContainerClient", &OracleClient{InitClientFn: initDbmulticloudOracleDBAzureBlobContainerClient})
	RegisterOracleClient("oci_dbmulticloud.OracleDBAzureBlobMountClient", &OracleClient{InitClientFn: initDbmulticloudOracleDBAzureBlobMountClient})
	RegisterOracleClient("oci_dbmulticloud.OracleDBAzureConnectorClient", &OracleClient{InitClientFn: initDbmulticloudOracleDBAzureConnectorClient})
	RegisterOracleClient("oci_dbmulticloud.OracleDbAzureKeyClient", &OracleClient{InitClientFn: initDbmulticloudOracleDbAzureKeyClient})
	RegisterOracleClient("oci_dbmulticloud.OracleDbAzureVaultClient", &OracleClient{InitClientFn: initDbmulticloudOracleDbAzureVaultClient})
	RegisterOracleClient("oci_dbmulticloud.OracleDbAzureVaultAssociationClient", &OracleClient{InitClientFn: initDbmulticloudOracleDbAzureVaultAssociationClient})
}

func initDbmulticloudDbMulticloudGCPProviderClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_dbmulticloud.NewDbMulticloudGCPProviderClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) DbMulticloudGCPProviderClient() *oci_dbmulticloud.DbMulticloudGCPProviderClient {
	return m.GetClient("oci_dbmulticloud.DbMulticloudGCPProviderClient").(*oci_dbmulticloud.DbMulticloudGCPProviderClient)
}

func initDbmulticloudWorkRequestClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_dbmulticloud.NewWorkRequestClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) DbmulticloudWorkRequestClient() *oci_dbmulticloud.WorkRequestClient {
	return m.GetClient("oci_dbmulticloud.WorkRequestClient").(*oci_dbmulticloud.WorkRequestClient)
}

func initDbmulticloudMultiCloudResourceDiscoveryClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_dbmulticloud.NewMultiCloudResourceDiscoveryClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) MultiCloudResourceDiscoveryClient() *oci_dbmulticloud.MultiCloudResourceDiscoveryClient {
	return m.GetClient("oci_dbmulticloud.MultiCloudResourceDiscoveryClient").(*oci_dbmulticloud.MultiCloudResourceDiscoveryClient)
}

func initDbmulticloudOracleDBAzureBlobContainerClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_dbmulticloud.NewOracleDBAzureBlobContainerClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) OracleDBAzureBlobContainerClient() *oci_dbmulticloud.OracleDBAzureBlobContainerClient {
	return m.GetClient("oci_dbmulticloud.OracleDBAzureBlobContainerClient").(*oci_dbmulticloud.OracleDBAzureBlobContainerClient)
}

func initDbmulticloudOracleDBAzureBlobMountClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_dbmulticloud.NewOracleDBAzureBlobMountClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) OracleDBAzureBlobMountClient() *oci_dbmulticloud.OracleDBAzureBlobMountClient {
	return m.GetClient("oci_dbmulticloud.OracleDBAzureBlobMountClient").(*oci_dbmulticloud.OracleDBAzureBlobMountClient)
}

func initDbmulticloudOracleDBAzureConnectorClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_dbmulticloud.NewOracleDBAzureConnectorClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) OracleDBAzureConnectorClient() *oci_dbmulticloud.OracleDBAzureConnectorClient {
	return m.GetClient("oci_dbmulticloud.OracleDBAzureConnectorClient").(*oci_dbmulticloud.OracleDBAzureConnectorClient)
}

func initDbmulticloudOracleDbAzureKeyClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_dbmulticloud.NewOracleDbAzureKeyClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) OracleDbAzureKeyClient() *oci_dbmulticloud.OracleDbAzureKeyClient {
	return m.GetClient("oci_dbmulticloud.OracleDbAzureKeyClient").(*oci_dbmulticloud.OracleDbAzureKeyClient)
}

func initDbmulticloudOracleDbAzureVaultClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_dbmulticloud.NewOracleDbAzureVaultClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) OracleDbAzureVaultClient() *oci_dbmulticloud.OracleDbAzureVaultClient {
	return m.GetClient("oci_dbmulticloud.OracleDbAzureVaultClient").(*oci_dbmulticloud.OracleDbAzureVaultClient)
}

func initDbmulticloudOracleDbAzureVaultAssociationClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_dbmulticloud.NewOracleDbAzureVaultAssociationClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) OracleDbAzureVaultAssociationClient() *oci_dbmulticloud.OracleDbAzureVaultAssociationClient {
	return m.GetClient("oci_dbmulticloud.OracleDbAzureVaultAssociationClient").(*oci_dbmulticloud.OracleDbAzureVaultAssociationClient)
}
