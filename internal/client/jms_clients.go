// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_jms "github.com/oracle/oci-go-sdk/v65/jms"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_jms.JavaManagementServiceClient", &OracleClient{InitClientFn: initJmsJavaManagementServiceClient})
}

func initJmsJavaManagementServiceClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_jms.NewJavaManagementServiceClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) JavaManagementServiceClient() *oci_jms.JavaManagementServiceClient {
	return m.GetClient("oci_jms.JavaManagementServiceClient").(*oci_jms.JavaManagementServiceClient)
}
