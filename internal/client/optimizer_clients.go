// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_optimizer "github.com/oracle/oci-go-sdk/v58/optimizer"

	oci_common "github.com/oracle/oci-go-sdk/v58/common"
)

func init() {
	RegisterOracleClient("oci_optimizer.OptimizerClient", &OracleClient{InitClientFn: initOptimizerOptimizerClient})
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

	if serviceClientOverrides.HostUrlOverride != "" {
		client.Host = serviceClientOverrides.HostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) OptimizerClient() *oci_optimizer.OptimizerClient {
	return m.GetClient("oci_optimizer.OptimizerClient").(*oci_optimizer.OptimizerClient)
}
