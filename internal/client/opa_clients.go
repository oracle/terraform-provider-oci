// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_opa "github.com/oracle/oci-go-sdk/v65/opa"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_opa.OpaInstanceClient", &OracleClient{InitClientFn: initOpaOpaInstanceClient})
}

func initOpaOpaInstanceClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_opa.NewOpaInstanceClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) OpaInstanceClient() *oci_opa.OpaInstanceClient {
	return m.GetClient("oci_opa.OpaInstanceClient").(*oci_opa.OpaInstanceClient)
}
