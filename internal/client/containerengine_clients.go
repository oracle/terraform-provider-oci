// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_containerengine "github.com/oracle/oci-go-sdk/v56/containerengine"

	oci_common "github.com/oracle/oci-go-sdk/v56/common"
)

func init() {
	RegisterOracleClient("oci_containerengine.ContainerEngineClient", &OracleClient{InitClientFn: initContainerengineContainerEngineClient})
}

func initContainerengineContainerEngineClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_containerengine.NewContainerEngineClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) ContainerEngineClient() *oci_containerengine.ContainerEngineClient {
	return m.GetClient("oci_containerengine.ContainerEngineClient").(*oci_containerengine.ContainerEngineClient)
}
