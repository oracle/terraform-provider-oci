// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_apm_config "github.com/oracle/oci-go-sdk/v51/apmconfig"

	oci_common "github.com/oracle/oci-go-sdk/v51/common"
)

func init() {
	RegisterOracleClient("oci_apm_config.ConfigClient", &OracleClient{InitClientFn: initApmconfigConfigClient})
}

func initApmconfigConfigClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_apm_config.NewConfigClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) configClient() *oci_apm_config.ConfigClient {
	return m.GetClient("oci_apm_config.ConfigClient").(*oci_apm_config.ConfigClient)
}
