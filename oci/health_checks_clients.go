// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_health_checks "github.com/oracle/oci-go-sdk/v27/healthchecks"

	oci_common "github.com/oracle/oci-go-sdk/v27/common"
)

func init() {
	RegisterOracleClient("oci_health_checks.HealthChecksClient", &OracleClient{initClientFn: initHealthchecksHealthChecksClient})
}

func initHealthchecksHealthChecksClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_health_checks.NewHealthChecksClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) healthChecksClient() *oci_health_checks.HealthChecksClient {
	return m.GetClient("oci_health_checks.HealthChecksClient").(*oci_health_checks.HealthChecksClient)
}
