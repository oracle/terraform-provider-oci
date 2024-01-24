// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_disaster_recovery "github.com/oracle/oci-go-sdk/v65/disasterrecovery"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_disaster_recovery.DisasterRecoveryClient", &OracleClient{InitClientFn: initDisasterrecoveryDisasterRecoveryClient})
}

func initDisasterrecoveryDisasterRecoveryClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_disaster_recovery.NewDisasterRecoveryClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) DisasterRecoveryClient() *oci_disaster_recovery.DisasterRecoveryClient {
	return m.GetClient("oci_disaster_recovery.DisasterRecoveryClient").(*oci_disaster_recovery.DisasterRecoveryClient)
}
