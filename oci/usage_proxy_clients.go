// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_usage_proxy "github.com/oracle/oci-go-sdk/v50/usage"

	oci_common "github.com/oracle/oci-go-sdk/v50/common"
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

	if serviceClientOverrides.hostUrlOverride != "" {
		client.Host = serviceClientOverrides.hostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) rewardsClient() *oci_usage_proxy.RewardsClient {
	return m.GetClient("oci_usage_proxy.RewardsClient").(*oci_usage_proxy.RewardsClient)
}
