// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_managed_kafka "github.com/oracle/oci-go-sdk/v65/managedkafka"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_managed_kafka.KafkaClusterClient", &OracleClient{InitClientFn: initManagedkafkaKafkaClusterClient})
}

func initManagedkafkaKafkaClusterClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_managed_kafka.NewKafkaClusterClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) KafkaClusterClient() *oci_managed_kafka.KafkaClusterClient {
	return m.GetClient("oci_managed_kafka.KafkaClusterClient").(*oci_managed_kafka.KafkaClusterClient)
}
