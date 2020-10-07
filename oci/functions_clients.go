// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_functions "github.com/oracle/oci-go-sdk/v26/functions"

	oci_common "github.com/oracle/oci-go-sdk/v26/common"
)

func init() {
	RegisterOracleClient("oci_functions.FunctionsInvokeClient", &OracleClient{initClientFn: initFunctionsFunctionsInvokeClient})
	RegisterOracleClient("oci_functions.FunctionsManagementClient", &OracleClient{initClientFn: initFunctionsFunctionsManagementClient})
}

func initFunctionsFunctionsInvokeClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_functions.NewFunctionsInvokeClientWithConfigurationProvider(configProvider, "DUMMY_ENDPOINT")
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) functionsInvokeClient() *oci_functions.FunctionsInvokeClient {
	return m.GetClient("oci_functions.FunctionsInvokeClient").(*oci_functions.FunctionsInvokeClient)
}

func initFunctionsFunctionsManagementClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_functions.NewFunctionsManagementClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) functionsManagementClient() *oci_functions.FunctionsManagementClient {
	return m.GetClient("oci_functions.FunctionsManagementClient").(*oci_functions.FunctionsManagementClient)
}
