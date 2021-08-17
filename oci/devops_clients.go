// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_devops "github.com/oracle/oci-go-sdk/v46/devops"

	oci_common "github.com/oracle/oci-go-sdk/v46/common"
)

func init() {
	RegisterOracleClient("oci_devops.DevopsClient", &OracleClient{initClientFn: initDevopsDevopsClient})
}

func initDevopsDevopsClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_devops.NewDevopsClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) devopsClient() *oci_devops.DevopsClient {
	return m.GetClient("oci_devops.DevopsClient").(*oci_devops.DevopsClient)
}
