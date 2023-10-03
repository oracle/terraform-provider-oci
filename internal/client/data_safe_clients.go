// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_data_safe.DataSafeClient", &OracleClient{InitClientFn: initDatasafeDataSafeClient})
}

func initDatasafeDataSafeClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_data_safe.NewDataSafeClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) DataSafeClient() *oci_data_safe.DataSafeClient {
	return m.GetClient("oci_data_safe.DataSafeClient").(*oci_data_safe.DataSafeClient)
}
