// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_analytics "github.com/oracle/oci-go-sdk/v65/analytics"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_analytics.AnalyticsClient", &OracleClient{InitClientFn: initAnalyticsAnalyticsClient})
}

func initAnalyticsAnalyticsClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_analytics.NewAnalyticsClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) AnalyticsClient() *oci_analytics.AnalyticsClient {
	return m.GetClient("oci_analytics.AnalyticsClient").(*oci_analytics.AnalyticsClient)
}
