// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_management_agent "github.com/oracle/oci-go-sdk/v48/managementagent"

	oci_common "github.com/oracle/oci-go-sdk/v48/common"
)

func init() {
	RegisterOracleClient("oci_management_agent.ManagementAgentClient", &OracleClient{initClientFn: initManagementagentManagementAgentClient})
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

	if serviceClientOverrides.hostUrlOverride != "" {
		client.Host = serviceClientOverrides.hostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) managementAgentClient() *oci_management_agent.ManagementAgentClient {
	return m.GetClient("oci_management_agent.ManagementAgentClient").(*oci_management_agent.ManagementAgentClient)
}
