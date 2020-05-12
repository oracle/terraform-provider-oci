// Copyright (c) 2020, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	oci_auto_scaling "github.com/oracle/oci-go-sdk/autoscaling"

	oci_common "github.com/oracle/oci-go-sdk/common"
)

func init() {
	RegisterOracleClient("oci_auto_scaling.AutoScalingClient", &OracleClient{initClientFn: initAutoscalingAutoScalingClient})
}

func initAutoscalingAutoScalingClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_auto_scaling.NewAutoScalingClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) autoScalingClient() *oci_auto_scaling.AutoScalingClient {
	return m.GetClient("oci_auto_scaling.AutoScalingClient").(*oci_auto_scaling.AutoScalingClient)
}
