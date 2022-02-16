// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_health_checks "github.com/oracle/oci-go-sdk/v58/healthchecks"

	oci_common "github.com/oracle/oci-go-sdk/v58/common"
)

func init() {
	RegisterOracleClient("oci_health_checks.HealthChecksClient", &OracleClient{InitClientFn: initHealthchecksHealthChecksClient})
}

func initHealthchecksHealthChecksClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_health_checks.NewHealthChecksClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) HealthChecksClient() *oci_health_checks.HealthChecksClient {
	return m.GetClient("oci_health_checks.HealthChecksClient").(*oci_health_checks.HealthChecksClient)
}
