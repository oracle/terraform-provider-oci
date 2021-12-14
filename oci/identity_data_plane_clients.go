// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_identity_data_plane "github.com/oracle/oci-go-sdk/v54/identitydataplane"

	oci_common "github.com/oracle/oci-go-sdk/v54/common"
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

	if serviceClientOverrides.hostUrlOverride != "" {
		client.Host = serviceClientOverrides.hostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) dataplaneClient() *oci_identity_data_plane.DataplaneClient {
	return m.GetClient("oci_identity_data_plane.DataplaneClient").(*oci_identity_data_plane.DataplaneClient)
}
