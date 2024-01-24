// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_auto_scaling "github.com/oracle/oci-go-sdk/v65/autoscaling"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_auto_scaling.AutoScalingClient", &OracleClient{InitClientFn: initAutoscalingAutoScalingClient})
}

func initAutoscalingAutoScalingClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_auto_scaling.NewAutoScalingClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) AutoScalingClient() *oci_auto_scaling.AutoScalingClient {
	return m.GetClient("oci_auto_scaling.AutoScalingClient").(*oci_auto_scaling.AutoScalingClient)
}
