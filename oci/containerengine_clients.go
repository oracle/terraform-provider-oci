// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_containerengine "github.com/oracle/oci-go-sdk/v26/containerengine"

	oci_common "github.com/oracle/oci-go-sdk/v26/common"
)

func init() {
	RegisterOracleClient("oci_containerengine.ContainerEngineClient", &OracleClient{initClientFn: initContainerengineContainerEngineClient})
}

func initContainerengineContainerEngineClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_containerengine.NewContainerEngineClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) containerEngineClient() *oci_containerengine.ContainerEngineClient {
	return m.GetClient("oci_containerengine.ContainerEngineClient").(*oci_containerengine.ContainerEngineClient)
}
