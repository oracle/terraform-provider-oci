// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_api_platform "github.com/oracle/oci-go-sdk/v65/apiplatform"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_api_platform.ApiPlatformClient", &OracleClient{InitClientFn: initApiplatformApiPlatformClient})
}

func initApiplatformApiPlatformClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_api_platform.NewApiPlatformClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) ApiPlatformClient() *oci_api_platform.ApiPlatformClient {
	return m.GetClient("oci_api_platform.ApiPlatformClient").(*oci_api_platform.ApiPlatformClient)
}
