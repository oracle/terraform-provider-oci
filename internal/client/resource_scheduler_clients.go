// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_resource_scheduler "github.com/oracle/oci-go-sdk/v65/resourcescheduler"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_resource_scheduler.ScheduleClient", &OracleClient{InitClientFn: initResourceschedulerScheduleClient})
}

func initResourceschedulerScheduleClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_resource_scheduler.NewScheduleClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) ScheduleClient() *oci_resource_scheduler.ScheduleClient {
	return m.GetClient("oci_resource_scheduler.ScheduleClient").(*oci_resource_scheduler.ScheduleClient)
}
