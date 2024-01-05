// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_network_load_balancer "github.com/oracle/oci-go-sdk/v65/networkloadbalancer"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_network_load_balancer.NetworkLoadBalancerClient", &OracleClient{InitClientFn: initNetworkloadbalancerNetworkLoadBalancerClient})
}

func initNetworkloadbalancerNetworkLoadBalancerClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_network_load_balancer.NewNetworkLoadBalancerClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) NetworkLoadBalancerClient() *oci_network_load_balancer.NetworkLoadBalancerClient {
	return m.GetClient("oci_network_load_balancer.NetworkLoadBalancerClient").(*oci_network_load_balancer.NetworkLoadBalancerClient)
}
