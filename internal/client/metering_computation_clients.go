// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_metering_computation "github.com/oracle/oci-go-sdk/v58/usageapi"

	oci_common "github.com/oracle/oci-go-sdk/v58/common"
)

func init() {
	RegisterOracleClient("oci_metering_computation.UsageapiClient", &OracleClient{InitClientFn: initUsageapiUsageapiClient})
}

func initUsageapiUsageapiClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_metering_computation.NewUsageapiClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) UsageapiClient() *oci_metering_computation.UsageapiClient {
	return m.GetClient("oci_metering_computation.UsageapiClient").(*oci_metering_computation.UsageapiClient)
}
