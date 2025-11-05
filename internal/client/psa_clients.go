// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_psa "github.com/oracle/oci-go-sdk/v65/psa"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_psa.PrivateServiceAccessClient", &OracleClient{InitClientFn: initPsaPrivateServiceAccessClient})
}

func initPsaPrivateServiceAccessClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_psa.NewPrivateServiceAccessClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) PrivateServiceAccessClient() *oci_psa.PrivateServiceAccessClient {
	return m.GetClient("oci_psa.PrivateServiceAccessClient").(*oci_psa.PrivateServiceAccessClient)
}
