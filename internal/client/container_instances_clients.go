// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_container_instances "github.com/oracle/oci-go-sdk/v65/containerinstances"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_container_instances.ContainerInstanceClient", &OracleClient{InitClientFn: initContainerinstancesContainerInstanceClient})
}

func initContainerinstancesContainerInstanceClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_container_instances.NewContainerInstanceClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) ContainerInstanceClient() *oci_container_instances.ContainerInstanceClient {
	return m.GetClient("oci_container_instances.ContainerInstanceClient").(*oci_container_instances.ContainerInstanceClient)
}
