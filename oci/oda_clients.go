// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_oda "github.com/oracle/oci-go-sdk/v34/oda"

	oci_common "github.com/oracle/oci-go-sdk/v34/common"
)

func init() {
	RegisterOracleClient("oci_oda.OdaClient", &OracleClient{initClientFn: initOdaOdaClient})
}

func initOdaOdaClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_oda.NewOdaClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) odaClient() *oci_oda.OdaClient {
	return m.GetClient("oci_oda.OdaClient").(*oci_oda.OdaClient)
}
