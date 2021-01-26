// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_log_analytics "github.com/oracle/oci-go-sdk/v34/loganalytics"

	oci_common "github.com/oracle/oci-go-sdk/v34/common"
)

func init() {
	RegisterOracleClient("oci_log_analytics.LogAnalyticsClient", &OracleClient{initClientFn: initLoganalyticsLogAnalyticsClient})
}

func initLoganalyticsLogAnalyticsClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_log_analytics.NewLogAnalyticsClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) logAnalyticsClient() *oci_log_analytics.LogAnalyticsClient {
	return m.GetClient("oci_log_analytics.LogAnalyticsClient").(*oci_log_analytics.LogAnalyticsClient)
}
