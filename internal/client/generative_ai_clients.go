// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_generative_ai "github.com/oracle/oci-go-sdk/v65/generativeai"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_generative_ai.GenerativeAiClient", &OracleClient{InitClientFn: initGenerativeaiGenerativeAiClient})
}

func initGenerativeaiGenerativeAiClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_generative_ai.NewGenerativeAiClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) GenerativeAiClient() *oci_generative_ai.GenerativeAiClient {
	return m.GetClient("oci_generative_ai.GenerativeAiClient").(*oci_generative_ai.GenerativeAiClient)
}
