// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_dblm "github.com/oracle/oci-go-sdk/v65/dblm"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_dblm.DbLifeCycleManagementClient", &OracleClient{InitClientFn: initDblmDbLifeCycleManagementClient})
}

func initDblmDbLifeCycleManagementClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_dblm.NewDbLifeCycleManagementClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) DbLifeCycleManagementClient() *oci_dblm.DbLifeCycleManagementClient {
	return m.GetClient("oci_dblm.DbLifeCycleManagementClient").(*oci_dblm.DbLifeCycleManagementClient)
}
