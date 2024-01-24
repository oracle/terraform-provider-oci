// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_compute_cloud_at_customer "github.com/oracle/oci-go-sdk/v65/computecloudatcustomer"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_compute_cloud_at_customer.ComputeCloudAtCustomerClient", &OracleClient{InitClientFn: initComputecloudatcustomerComputeCloudAtCustomerClient})
}

func initComputecloudatcustomerComputeCloudAtCustomerClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_compute_cloud_at_customer.NewComputeCloudAtCustomerClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) ComputeCloudAtCustomerClient() *oci_compute_cloud_at_customer.ComputeCloudAtCustomerClient {
	return m.GetClient("oci_compute_cloud_at_customer.ComputeCloudAtCustomerClient").(*oci_compute_cloud_at_customer.ComputeCloudAtCustomerClient)
}
