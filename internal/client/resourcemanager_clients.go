// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_resourcemanager "github.com/oracle/oci-go-sdk/v56/resourcemanager"

	oci_common "github.com/oracle/oci-go-sdk/v56/common"
)

func init() {
	RegisterOracleClient("oci_resourcemanager.ResourceManagerClient", &OracleClient{InitClientFn: initResourcemanagerResourceManagerClient})
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

	if serviceClientOverrides.HostUrlOverride != "" {
		client.Host = serviceClientOverrides.HostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) ResourceManagerClient() *oci_resourcemanager.ResourceManagerClient {
	return m.GetClient("oci_resourcemanager.ResourceManagerClient").(*oci_resourcemanager.ResourceManagerClient)
}
