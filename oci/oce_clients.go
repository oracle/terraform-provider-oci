// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_oce "github.com/oracle/oci-go-sdk/v32/oce"

	oci_common "github.com/oracle/oci-go-sdk/v32/common"
)

func init() {
	RegisterOracleClient("oci_oce.OceInstanceClient", &OracleClient{initClientFn: initOceOceInstanceClient})
}

func initOceOceInstanceClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_oce.NewOceInstanceClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) oceInstanceClient() *oci_oce.OceInstanceClient {
	return m.GetClient("oci_oce.OceInstanceClient").(*oci_oce.OceInstanceClient)
}
