// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_load_balancer "github.com/oracle/oci-go-sdk/v65/loadbalancer"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_load_balancer.LoadBalancerClient", &OracleClient{InitClientFn: initLoadbalancerLoadBalancerClient})
}

func initLoadbalancerLoadBalancerClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_load_balancer.NewLoadBalancerClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) LoadBalancerClient() *oci_load_balancer.LoadBalancerClient {
	return m.GetClient("oci_load_balancer.LoadBalancerClient").(*oci_load_balancer.LoadBalancerClient)
}
