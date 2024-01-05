// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_kms "github.com/oracle/oci-go-sdk/v65/keymanagement"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_kms.EkmClient", &OracleClient{InitClientFn: initKeymanagementEkmClient})
	RegisterOracleClient("oci_kms.KmsCryptoClient", &OracleClient{InitClientFn: initKeymanagementKmsCryptoClient})
	RegisterOracleClient("oci_kms.KmsManagementClient", &OracleClient{InitClientFn: initKeymanagementKmsManagementClient})
	RegisterOracleClient("oci_kms.KmsVaultClient", &OracleClient{InitClientFn: initKeymanagementKmsVaultClient})
}

func initKeymanagementEkmClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_kms.NewEkmClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) EkmClient() *oci_kms.EkmClient {
	return m.GetClient("oci_kms.EkmClient").(*oci_kms.EkmClient)
}

func initKeymanagementKmsCryptoClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_kms.NewKmsCryptoClientWithConfigurationProvider(configProvider, "DUMMY_ENDPOINT")
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

func (m *OracleClients) KmsCryptoClient() *oci_kms.KmsCryptoClient {
	return m.GetClient("oci_kms.KmsCryptoClient").(*oci_kms.KmsCryptoClient)
}

func initKeymanagementKmsManagementClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_kms.NewKmsManagementClientWithConfigurationProvider(configProvider, "DUMMY_ENDPOINT")
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

func (m *OracleClients) KmsManagementClient() *oci_kms.KmsManagementClient {
	return m.GetClient("oci_kms.KmsManagementClient").(*oci_kms.KmsManagementClient)
}

func initKeymanagementKmsVaultClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_kms.NewKmsVaultClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) KmsVaultClient() *oci_kms.KmsVaultClient {
	return m.GetClient("oci_kms.KmsVaultClient").(*oci_kms.KmsVaultClient)
}
