// Copyright (c) 2020, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	oci_resourcemanager "github.com/oracle/oci-go-sdk/resourcemanager"

	oci_common "github.com/oracle/oci-go-sdk/common"
)

func init() {
	RegisterOracleClient("oci_resourcemanager.ResourceManagerClient", &OracleClient{initClientFn: initResourcemanagerResourceManagerClient})
}

func initResourcemanagerResourceManagerClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_resourcemanager.NewResourceManagerClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) resourceManagerClient() *oci_resourcemanager.ResourceManagerClient {
	return m.GetClient("oci_resourcemanager.ResourceManagerClient").(*oci_resourcemanager.ResourceManagerClient)
}
