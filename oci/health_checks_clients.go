// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_health_checks "github.com/oracle/oci-go-sdk/v31/healthchecks"

	oci_common "github.com/oracle/oci-go-sdk/v31/common"
)

func init() {
	RegisterOracleClient("oci_health_checks.HealthChecksClient", &OracleClient{initClientFn: initHealthchecksHealthChecksClient})
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

	if serviceClientOverrides.hostUrlOverride != "" {
		client.Host = serviceClientOverrides.hostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) healthChecksClient() *oci_health_checks.HealthChecksClient {
	return m.GetClient("oci_health_checks.HealthChecksClient").(*oci_health_checks.HealthChecksClient)
}
