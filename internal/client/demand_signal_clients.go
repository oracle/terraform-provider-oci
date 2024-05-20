// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_demand_signal "github.com/oracle/oci-go-sdk/v65/demandsignal"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_demand_signal.OccDemandSignalClient", &OracleClient{InitClientFn: initDemandsignalOccDemandSignalClient})
}

func initDemandsignalOccDemandSignalClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_demand_signal.NewOccDemandSignalClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) OccDemandSignalClient() *oci_demand_signal.OccDemandSignalClient {
	return m.GetClient("oci_demand_signal.OccDemandSignalClient").(*oci_demand_signal.OccDemandSignalClient)
}
