// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_monitoring "github.com/oracle/oci-go-sdk/v65/monitoring"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_monitoring.MonitoringClient", &OracleClient{InitClientFn: initMonitoringMonitoringClient})
}

func initMonitoringMonitoringClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_monitoring.NewMonitoringClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) MonitoringClient() *oci_monitoring.MonitoringClient {
	return m.GetClient("oci_monitoring.MonitoringClient").(*oci_monitoring.MonitoringClient)
}
