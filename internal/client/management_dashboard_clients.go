// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_management_dashboard "github.com/oracle/oci-go-sdk/v58/managementdashboard"

	oci_common "github.com/oracle/oci-go-sdk/v58/common"
)

func init() {
	RegisterOracleClient("oci_management_dashboard.DashxApisClient", &OracleClient{InitClientFn: initManagementdashboardDashxApisClient})
}

func initManagementdashboardDashxApisClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_management_dashboard.NewDashxApisClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) DashxApisClient() *oci_management_dashboard.DashxApisClient {
	return m.GetClient("oci_management_dashboard.DashxApisClient").(*oci_management_dashboard.DashxApisClient)
}
