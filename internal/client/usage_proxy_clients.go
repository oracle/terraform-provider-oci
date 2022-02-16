// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_usage_proxy "github.com/oracle/oci-go-sdk/v58/usage"

	oci_common "github.com/oracle/oci-go-sdk/v58/common"
)

func init() {
	RegisterOracleClient("oci_usage_proxy.RewardsClient", &OracleClient{InitClientFn: initUsageRewardsClient})
}

func initUsageRewardsClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_usage_proxy.NewRewardsClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) RewardsClient() *oci_usage_proxy.RewardsClient {
	return m.GetClient("oci_usage_proxy.RewardsClient").(*oci_usage_proxy.RewardsClient)
}
