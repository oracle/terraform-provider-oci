// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_analytics "github.com/oracle/oci-go-sdk/v26/analytics"

	oci_common "github.com/oracle/oci-go-sdk/v26/common"
)

func init() {
	RegisterOracleClient("oci_analytics.AnalyticsClient", &OracleClient{initClientFn: initAnalyticsAnalyticsClient})
}

func initAnalyticsAnalyticsClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_analytics.NewAnalyticsClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) analyticsClient() *oci_analytics.AnalyticsClient {
	return m.GetClient("oci_analytics.AnalyticsClient").(*oci_analytics.AnalyticsClient)
}
