// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_monitoring "github.com/oracle/oci-go-sdk/monitoring"

	oci_common "github.com/oracle/oci-go-sdk/common"
)

func init() {
	RegisterOracleClient("oci_monitoring.MonitoringClient", &OracleClient{initClientFn: initMonitoringMonitoringClient})
}

func initMonitoringMonitoringClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_monitoring.NewMonitoringClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) monitoringClient() *oci_monitoring.MonitoringClient {
	return m.GetClient("oci_monitoring.MonitoringClient").(*oci_monitoring.MonitoringClient)
}
