// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_management_agent "github.com/oracle/oci-go-sdk/v27/managementagent"

	oci_common "github.com/oracle/oci-go-sdk/v27/common"
)

func init() {
	RegisterOracleClient("oci_management_agent.ManagementAgentClient", &OracleClient{initClientFn: initManagementagentManagementAgentClient})
}

func initManagementagentManagementAgentClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_management_agent.NewManagementAgentClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) managementAgentClient() *oci_management_agent.ManagementAgentClient {
	return m.GetClient("oci_management_agent.ManagementAgentClient").(*oci_management_agent.ManagementAgentClient)
}
