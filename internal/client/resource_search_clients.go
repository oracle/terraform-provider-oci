// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_resource_search "github.com/oracle/oci-go-sdk/v65/resourcesearch"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_resource_search.ResourceSearchClient", &OracleClient{InitClientFn: initResourcesearchResourceSearchClient})
}

func initResourcesearchResourceSearchClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_resource_search.NewResourceSearchClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) ResourceSearchClient() *oci_resource_search.ResourceSearchClient {
	return m.GetClient("oci_resource_search.ResourceSearchClient").(*oci_resource_search.ResourceSearchClient)
}
