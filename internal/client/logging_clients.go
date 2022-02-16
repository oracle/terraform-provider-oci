// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_logging "github.com/oracle/oci-go-sdk/v58/logging"

	oci_common "github.com/oracle/oci-go-sdk/v58/common"
)

func init() {
	RegisterOracleClient("oci_logging.LoggingManagementClient", &OracleClient{InitClientFn: initLoggingLoggingManagementClient})
}

func initLoggingLoggingManagementClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_logging.NewLoggingManagementClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) LoggingManagementClient() *oci_logging.LoggingManagementClient {
	return m.GetClient("oci_logging.LoggingManagementClient").(*oci_logging.LoggingManagementClient)
}
