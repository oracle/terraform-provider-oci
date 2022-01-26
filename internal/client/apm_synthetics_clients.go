// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_apm_synthetics "github.com/oracle/oci-go-sdk/v56/apmsynthetics"

	oci_common "github.com/oracle/oci-go-sdk/v56/common"
)

func init() {
	RegisterOracleClient("oci_apm_synthetics.ApmSyntheticClient", &OracleClient{InitClientFn: initApmsyntheticsApmSyntheticClient})
}

func initApmsyntheticsApmSyntheticClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_apm_synthetics.NewApmSyntheticClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) ApmSyntheticClient() *oci_apm_synthetics.ApmSyntheticClient {
	return m.GetClient("oci_apm_synthetics.ApmSyntheticClient").(*oci_apm_synthetics.ApmSyntheticClient)
}
