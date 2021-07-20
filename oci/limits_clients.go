// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_limits "github.com/oracle/oci-go-sdk/v45/limits"

	oci_common "github.com/oracle/oci-go-sdk/v45/common"
)

func init() {
	RegisterOracleClient("oci_limits.LimitsClient", &OracleClient{initClientFn: initLimitsLimitsClient})
	RegisterOracleClient("oci_limits.QuotasClient", &OracleClient{initClientFn: initLimitsQuotasClient})
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

	if serviceClientOverrides.hostUrlOverride != "" {
		client.Host = serviceClientOverrides.hostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) limitsClient() *oci_limits.LimitsClient {
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

	if serviceClientOverrides.hostUrlOverride != "" {
		client.Host = serviceClientOverrides.hostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) quotasClient() *oci_limits.QuotasClient {
	return m.GetClient("oci_limits.QuotasClient").(*oci_limits.QuotasClient)
}
