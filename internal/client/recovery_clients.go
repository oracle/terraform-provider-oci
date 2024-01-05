// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_recovery "github.com/oracle/oci-go-sdk/v65/recovery"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_recovery.DatabaseRecoveryClient", &OracleClient{InitClientFn: initRecoveryDatabaseRecoveryClient})
}

func initRecoveryDatabaseRecoveryClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_recovery.NewDatabaseRecoveryClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) DatabaseRecoveryClient() *oci_recovery.DatabaseRecoveryClient {
	return m.GetClient("oci_recovery.DatabaseRecoveryClient").(*oci_recovery.DatabaseRecoveryClient)
}
