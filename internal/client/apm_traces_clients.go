// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_apm_traces "github.com/oracle/oci-go-sdk/v65/apmtraces"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_apm_traces.QueryClient", &OracleClient{InitClientFn: initApmtracesQueryClient})
	RegisterOracleClient("oci_apm_traces.TraceClient", &OracleClient{InitClientFn: initApmtracesTraceClient})
}

func initApmtracesQueryClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_apm_traces.NewQueryClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) QueryClient() *oci_apm_traces.QueryClient {
	return m.GetClient("oci_apm_traces.QueryClient").(*oci_apm_traces.QueryClient)
}

func initApmtracesTraceClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_apm_traces.NewTraceClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) TraceClient() *oci_apm_traces.TraceClient {
	return m.GetClient("oci_apm_traces.TraceClient").(*oci_apm_traces.TraceClient)
}
