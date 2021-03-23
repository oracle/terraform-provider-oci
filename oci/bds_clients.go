// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_bds "github.com/oracle/oci-go-sdk/v37/bds"

	oci_common "github.com/oracle/oci-go-sdk/v37/common"
)

func init() {
	RegisterOracleClient("oci_bds.BdsClient", &OracleClient{initClientFn: initBdsBdsClient})
}

func initBdsBdsClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_bds.NewBdsClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) bdsClient() *oci_bds.BdsClient {
	return m.GetClient("oci_bds.BdsClient").(*oci_bds.BdsClient)
}
