// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_fusion_apps "github.com/oracle/oci-go-sdk/v65/fusionapps"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_fusion_apps.FusionApplicationsClient", &OracleClient{InitClientFn: initFusionappsFusionApplicationsClient})
}

func initFusionappsFusionApplicationsClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_fusion_apps.NewFusionApplicationsClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) FusionApplicationsClient() *oci_fusion_apps.FusionApplicationsClient {
	return m.GetClient("oci_fusion_apps.FusionApplicationsClient").(*oci_fusion_apps.FusionApplicationsClient)
}
