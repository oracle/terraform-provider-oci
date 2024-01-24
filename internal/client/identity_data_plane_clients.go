// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_identity_data_plane "github.com/oracle/oci-go-sdk/v65/identitydataplane"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_identity_data_plane.DataplaneClient", &OracleClient{InitClientFn: initIdentitydataplaneDataplaneClient})
}

func initIdentitydataplaneDataplaneClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_identity_data_plane.NewDataplaneClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) DataplaneClient() *oci_identity_data_plane.DataplaneClient {
	return m.GetClient("oci_identity_data_plane.DataplaneClient").(*oci_identity_data_plane.DataplaneClient)
}
