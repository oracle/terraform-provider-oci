// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_desktops "github.com/oracle/oci-go-sdk/v65/desktops"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_desktops.DesktopServiceClient", &OracleClient{InitClientFn: initDesktopsDesktopServiceClient})
}

func initDesktopsDesktopServiceClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_desktops.NewDesktopServiceClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) DesktopServiceClient() *oci_desktops.DesktopServiceClient {
	return m.GetClient("oci_desktops.DesktopServiceClient").(*oci_desktops.DesktopServiceClient)
}
