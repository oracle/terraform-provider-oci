// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_optimizer "github.com/oracle/oci-go-sdk/v28/optimizer"

	oci_common "github.com/oracle/oci-go-sdk/v28/common"
)

func init() {
	RegisterOracleClient("oci_optimizer.OptimizerClient", &OracleClient{initClientFn: initOptimizerOptimizerClient})
}

func initOptimizerOptimizerClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_optimizer.NewOptimizerClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) optimizerClient() *oci_optimizer.OptimizerClient {
	return m.GetClient("oci_optimizer.OptimizerClient").(*oci_optimizer.OptimizerClient)
}
