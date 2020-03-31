// Copyright (c) 2020, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	oci_oce "github.com/oracle/oci-go-sdk/oce"

	oci_common "github.com/oracle/oci-go-sdk/common"
)

func init() {
	RegisterOracleClient("oci_oce.OceInstanceClient", &OracleClient{initClientFn: initOceOceInstanceClient})
}

func initOceOceInstanceClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_oce.NewOceInstanceClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) oceInstanceClient() *oci_oce.OceInstanceClient {
	return m.GetClient("oci_oce.OceInstanceClient").(*oci_oce.OceInstanceClient)
}
