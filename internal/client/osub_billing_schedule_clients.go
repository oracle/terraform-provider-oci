// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_osub_billing_schedule "github.com/oracle/oci-go-sdk/v65/osubbillingschedule"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_osub_billing_schedule.BillingScheduleClient", &OracleClient{InitClientFn: initOsubbillingscheduleBillingScheduleClient})
}

func initOsubbillingscheduleBillingScheduleClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_osub_billing_schedule.NewBillingScheduleClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) BillingScheduleClient() *oci_osub_billing_schedule.BillingScheduleClient {
	return m.GetClient("oci_osub_billing_schedule.BillingScheduleClient").(*oci_osub_billing_schedule.BillingScheduleClient)
}
