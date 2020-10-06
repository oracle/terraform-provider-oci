// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_logging "github.com/oracle/oci-go-sdk/v26/logging"

	oci_common "github.com/oracle/oci-go-sdk/v26/common"
)

func init() {
	RegisterOracleClient("oci_logging.LoggingManagementClient", &OracleClient{initClientFn: initLoggingLoggingManagementClient})
}

func initLoggingLoggingManagementClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_logging.NewLoggingManagementClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) loggingManagementClient() *oci_logging.LoggingManagementClient {
	return m.GetClient("oci_logging.LoggingManagementClient").(*oci_logging.LoggingManagementClient)
}
