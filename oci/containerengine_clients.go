// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_containerengine "github.com/oracle/oci-go-sdk/v47/containerengine"

	oci_common "github.com/oracle/oci-go-sdk/v47/common"
)

func init() {
	RegisterOracleClient("oci_containerengine.ContainerEngineClient", &OracleClient{initClientFn: initContainerengineContainerEngineClient})
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

	if serviceClientOverrides.hostUrlOverride != "" {
		client.Host = serviceClientOverrides.hostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) containerEngineClient() *oci_containerengine.ContainerEngineClient {
	return m.GetClient("oci_containerengine.ContainerEngineClient").(*oci_containerengine.ContainerEngineClient)
}
