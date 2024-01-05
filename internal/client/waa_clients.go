// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_waa "github.com/oracle/oci-go-sdk/v65/waa"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_waa.WaaClient", &OracleClient{InitClientFn: initWaaWaaClient})
	RegisterOracleClient("oci_waa.WorkRequestClient", &OracleClient{InitClientFn: initWaaWorkRequestClient})
}

func initWaaWaaClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_waa.NewWaaClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) WaaClient() *oci_waa.WaaClient {
	return m.GetClient("oci_waa.WaaClient").(*oci_waa.WaaClient)
}

func initWaaWorkRequestClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_waa.NewWorkRequestClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) WaaWorkRequestClient() *oci_waa.WorkRequestClient {
	return m.GetClient("oci_waa.WorkRequestClient").(*oci_waa.WorkRequestClient)
}
