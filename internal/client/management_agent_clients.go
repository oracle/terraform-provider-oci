// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_management_agent "github.com/oracle/oci-go-sdk/v56/managementagent"

	oci_common "github.com/oracle/oci-go-sdk/v56/common"
)

func init() {
	RegisterOracleClient("oci_management_agent.ManagementAgentClient", &OracleClient{InitClientFn: initManagementagentManagementAgentClient})
}

func initManagementagentManagementAgentClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_management_agent.NewManagementAgentClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) ManagementAgentClient() *oci_management_agent.ManagementAgentClient {
	return m.GetClient("oci_management_agent.ManagementAgentClient").(*oci_management_agent.ManagementAgentClient)
}
