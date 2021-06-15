// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_waas "github.com/oracle/oci-go-sdk/v42/waas"

	oci_common "github.com/oracle/oci-go-sdk/v42/common"
)

func init() {
	RegisterOracleClient("oci_waas.RedirectClient", &OracleClient{initClientFn: initWaasRedirectClient})
	RegisterOracleClient("oci_waas.WaasClient", &OracleClient{initClientFn: initWaasWaasClient})
}

func initWaasRedirectClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_waas.NewRedirectClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.hostUrlOverride != "" {
		client.Host = serviceClientOverrides.hostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) redirectClient() *oci_waas.RedirectClient {
	return m.GetClient("oci_waas.RedirectClient").(*oci_waas.RedirectClient)
}

func initWaasWaasClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_waas.NewWaasClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.hostUrlOverride != "" {
		client.Host = serviceClientOverrides.hostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) waasClient() *oci_waas.WaasClient {
	return m.GetClient("oci_waas.WaasClient").(*oci_waas.WaasClient)
}
