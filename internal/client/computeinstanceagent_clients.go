// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_computeinstanceagent "github.com/oracle/oci-go-sdk/v58/computeinstanceagent"

	oci_common "github.com/oracle/oci-go-sdk/v58/common"
)

func init() {
	RegisterOracleClient("oci_computeinstanceagent.PluginClient", &OracleClient{InitClientFn: initComputeinstanceagentPluginClient})
	RegisterOracleClient("oci_computeinstanceagent.PluginconfigClient", &OracleClient{InitClientFn: initComputeinstanceagentPluginconfigClient})
}

func initComputeinstanceagentPluginClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_computeinstanceagent.NewPluginClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) PluginClient() *oci_computeinstanceagent.PluginClient {
	return m.GetClient("oci_computeinstanceagent.PluginClient").(*oci_computeinstanceagent.PluginClient)
}

func initComputeinstanceagentPluginconfigClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_computeinstanceagent.NewPluginconfigClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) PluginconfigClient() *oci_computeinstanceagent.PluginconfigClient {
	return m.GetClient("oci_computeinstanceagent.PluginconfigClient").(*oci_computeinstanceagent.PluginconfigClient)
}
