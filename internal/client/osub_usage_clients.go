// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_osub_usage "github.com/oracle/oci-go-sdk/v65/osubusage"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_osub_usage.ComputedUsageClient", &OracleClient{InitClientFn: initOsubusageComputedUsageClient})
}

func initOsubusageComputedUsageClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_osub_usage.NewComputedUsageClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) ComputedUsageClient() *oci_osub_usage.ComputedUsageClient {
	return m.GetClient("oci_osub_usage.ComputedUsageClient").(*oci_osub_usage.ComputedUsageClient)
}
