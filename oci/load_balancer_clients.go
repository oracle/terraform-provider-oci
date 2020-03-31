// Copyright (c) 2020, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	oci_load_balancer "github.com/oracle/oci-go-sdk/loadbalancer"

	oci_common "github.com/oracle/oci-go-sdk/common"
)

func init() {
	RegisterOracleClient("oci_load_balancer.LoadBalancerClient", &OracleClient{initClientFn: initLoadbalancerLoadBalancerClient})
}

func initLoadbalancerLoadBalancerClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_load_balancer.NewLoadBalancerClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) loadBalancerClient() *oci_load_balancer.LoadBalancerClient {
	return m.GetClient("oci_load_balancer.LoadBalancerClient").(*oci_load_balancer.LoadBalancerClient)
}
