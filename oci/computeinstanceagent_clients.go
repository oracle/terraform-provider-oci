// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_computeinstanceagent "github.com/oracle/oci-go-sdk/v46/computeinstanceagent"

	oci_common "github.com/oracle/oci-go-sdk/v46/common"
)

func init() {
	RegisterOracleClient("oci_computeinstanceagent.PluginClient", &OracleClient{initClientFn: initComputeinstanceagentPluginClient})
	RegisterOracleClient("oci_computeinstanceagent.PluginconfigClient", &OracleClient{initClientFn: initComputeinstanceagentPluginconfigClient})
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

	if serviceClientOverrides.hostUrlOverride != "" {
		client.Host = serviceClientOverrides.hostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) pluginClient() *oci_computeinstanceagent.PluginClient {
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

	if serviceClientOverrides.hostUrlOverride != "" {
		client.Host = serviceClientOverrides.hostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) pluginconfigClient() *oci_computeinstanceagent.PluginconfigClient {
	return m.GetClient("oci_computeinstanceagent.PluginconfigClient").(*oci_computeinstanceagent.PluginconfigClient)
}
