// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_events "github.com/oracle/oci-go-sdk/v42/events"

	oci_common "github.com/oracle/oci-go-sdk/v42/common"
)

func init() {
	RegisterOracleClient("oci_events.EventsClient", &OracleClient{initClientFn: initEventsEventsClient})
}

func initEventsEventsClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_events.NewEventsClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) eventsClient() *oci_events.EventsClient {
	return m.GetClient("oci_events.EventsClient").(*oci_events.EventsClient)
}
