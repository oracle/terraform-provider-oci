// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_wlms "github.com/oracle/oci-go-sdk/v65/wlms"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_wlms.WeblogicManagementServiceClient", &OracleClient{InitClientFn: initWlmsWeblogicManagementServiceClient})
}

func initWlmsWeblogicManagementServiceClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_wlms.NewWeblogicManagementServiceClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) WeblogicManagementServiceClient() *oci_wlms.WeblogicManagementServiceClient {
	return m.GetClient("oci_wlms.WeblogicManagementServiceClient").(*oci_wlms.WeblogicManagementServiceClient)
}
