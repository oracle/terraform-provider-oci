// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_limits "github.com/oracle/oci-go-sdk/v58/limits"

	oci_common "github.com/oracle/oci-go-sdk/v58/common"
)

func init() {
	RegisterOracleClient("oci_limits.LimitsClient", &OracleClient{InitClientFn: initLimitsLimitsClient})
	RegisterOracleClient("oci_limits.QuotasClient", &OracleClient{InitClientFn: initLimitsQuotasClient})
}

func initLimitsLimitsClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_limits.NewLimitsClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) LimitsClient() *oci_limits.LimitsClient {
	return m.GetClient("oci_limits.LimitsClient").(*oci_limits.LimitsClient)
}

func initLimitsQuotasClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_limits.NewQuotasClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) QuotasClient() *oci_limits.QuotasClient {
	return m.GetClient("oci_limits.QuotasClient").(*oci_limits.QuotasClient)
}
