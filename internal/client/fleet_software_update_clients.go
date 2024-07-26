// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_fleet_software_update "github.com/oracle/oci-go-sdk/v65/fleetsoftwareupdate"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_fleet_software_update.FleetSoftwareUpdateClient", &OracleClient{InitClientFn: initFleetsoftwareupdateFleetSoftwareUpdateClient})
}

func initFleetsoftwareupdateFleetSoftwareUpdateClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_fleet_software_update.NewFleetSoftwareUpdateClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) FleetSoftwareUpdateClient() *oci_fleet_software_update.FleetSoftwareUpdateClient {
	return m.GetClient("oci_fleet_software_update.FleetSoftwareUpdateClient").(*oci_fleet_software_update.FleetSoftwareUpdateClient)
}
