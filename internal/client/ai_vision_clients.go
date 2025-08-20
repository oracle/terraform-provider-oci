// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_ai_vision "github.com/oracle/oci-go-sdk/v65/aivision"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_ai_vision.AiServiceVisionClient", &OracleClient{InitClientFn: initAivisionAiServiceVisionClient})
}

func initAivisionAiServiceVisionClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_ai_vision.NewAIServiceVisionClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) AiServiceVisionClient() *oci_ai_vision.AIServiceVisionClient {
	return m.GetClient("oci_ai_vision.AiServiceVisionClient").(*oci_ai_vision.AIServiceVisionClient)
}
