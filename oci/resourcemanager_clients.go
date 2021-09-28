// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_resourcemanager "github.com/oracle/oci-go-sdk/v48/resourcemanager"

	oci_common "github.com/oracle/oci-go-sdk/v48/common"
)

func init() {
	RegisterOracleClient("oci_resourcemanager.ResourceManagerClient", &OracleClient{initClientFn: initResourcemanagerResourceManagerClient})
}

func initResourcemanagerResourceManagerClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_resourcemanager.NewResourceManagerClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) resourceManagerClient() *oci_resourcemanager.ResourceManagerClient {
	return m.GetClient("oci_resourcemanager.ResourceManagerClient").(*oci_resourcemanager.ResourceManagerClient)
}
