// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_service_mesh "github.com/oracle/oci-go-sdk/v65/servicemesh"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_service_mesh.ServiceMeshClient", &OracleClient{InitClientFn: initServicemeshServiceMeshClient})
}

func initServicemeshServiceMeshClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_service_mesh.NewServiceMeshClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) ServiceMeshClient() *oci_service_mesh.ServiceMeshClient {
	return m.GetClient("oci_service_mesh.ServiceMeshClient").(*oci_service_mesh.ServiceMeshClient)
}
