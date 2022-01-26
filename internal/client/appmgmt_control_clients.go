// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_appmgmt_control "github.com/oracle/oci-go-sdk/v56/appmgmtcontrol"

	oci_common "github.com/oracle/oci-go-sdk/v56/common"
)

func init() {
	RegisterOracleClient("oci_appmgmt_control.AppmgmtControlClient", &OracleClient{InitClientFn: initAppmgmtcontrolAppmgmtControlClient})
}

func initAppmgmtcontrolAppmgmtControlClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_appmgmt_control.NewAppmgmtControlClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) AppmgmtControlClient() *oci_appmgmt_control.AppmgmtControlClient {
	return m.GetClient("oci_appmgmt_control.AppmgmtControlClient").(*oci_appmgmt_control.AppmgmtControlClient)
}
