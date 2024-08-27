// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_delegate_access_control "github.com/oracle/oci-go-sdk/v65/delegateaccesscontrol"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_delegate_access_control.DelegateAccessControlClient", &OracleClient{InitClientFn: initDelegateaccesscontrolDelegateAccessControlClient})
	RegisterOracleClient("oci_delegate_access_control.WorkRequestClient", &OracleClient{InitClientFn: initDelegateaccesscontrolWorkRequestClient})
}

func initDelegateaccesscontrolDelegateAccessControlClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_delegate_access_control.NewDelegateAccessControlClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) DelegateAccessControlClient() *oci_delegate_access_control.DelegateAccessControlClient {
	return m.GetClient("oci_delegate_access_control.DelegateAccessControlClient").(*oci_delegate_access_control.DelegateAccessControlClient)
}

func initDelegateaccesscontrolWorkRequestClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_delegate_access_control.NewWorkRequestClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) DelegateAccessControlWorkRequestClient() *oci_delegate_access_control.WorkRequestClient {
	return m.GetClient("oci_delegate_access_control.WorkRequestClient").(*oci_delegate_access_control.WorkRequestClient)
}
