// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_waas "github.com/oracle/oci-go-sdk/v27/waas"

	oci_common "github.com/oracle/oci-go-sdk/v27/common"
)

func init() {
	RegisterOracleClient("oci_waas.RedirectClient", &OracleClient{initClientFn: initWaasRedirectClient})
	RegisterOracleClient("oci_waas.WaasClient", &OracleClient{initClientFn: initWaasWaasClient})
}

func initWaasRedirectClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_waas.NewRedirectClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) redirectClient() *oci_waas.RedirectClient {
	return m.GetClient("oci_waas.RedirectClient").(*oci_waas.RedirectClient)
}

func initWaasWaasClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_waas.NewWaasClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) waasClient() *oci_waas.WaasClient {
	return m.GetClient("oci_waas.WaasClient").(*oci_waas.WaasClient)
}
