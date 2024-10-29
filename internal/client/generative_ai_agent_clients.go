// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_generative_ai_agent "github.com/oracle/oci-go-sdk/v65/generativeaiagent"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_generative_ai_agent.GenerativeAiAgentClient", &OracleClient{InitClientFn: initGenerativeaiagentGenerativeAiAgentClient})
}

func initGenerativeaiagentGenerativeAiAgentClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_generative_ai_agent.NewGenerativeAiAgentClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) GenerativeAiAgentClient() *oci_generative_ai_agent.GenerativeAiAgentClient {
	return m.GetClient("oci_generative_ai_agent.GenerativeAiAgentClient").(*oci_generative_ai_agent.GenerativeAiAgentClient)
}
