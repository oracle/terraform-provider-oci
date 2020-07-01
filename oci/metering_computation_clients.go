// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_metering_computation "github.com/oracle/oci-go-sdk/usageapi"

	oci_common "github.com/oracle/oci-go-sdk/common"
)

func init() {
	RegisterOracleClient("oci_metering_computation.UsageapiClient", &OracleClient{initClientFn: initUsageapiUsageapiClient})
}

func initUsageapiUsageapiClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_metering_computation.NewUsageapiClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) usageapiClient() *oci_metering_computation.UsageapiClient {
	return m.GetClient("oci_metering_computation.UsageapiClient").(*oci_metering_computation.UsageapiClient)
}
