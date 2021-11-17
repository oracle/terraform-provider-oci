// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_service_manager_proxy "github.com/oracle/oci-go-sdk/v52/servicemanagerproxy"

	oci_common "github.com/oracle/oci-go-sdk/v52/common"
)

func init() {
	RegisterOracleClient("oci_service_manager_proxy.ServiceManagerProxyClient", &OracleClient{InitClientFn: initServicemanagerproxyServiceManagerProxyClient})
}

func initServicemanagerproxyServiceManagerProxyClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_service_manager_proxy.NewServiceManagerProxyClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) serviceManagerProxyClient() *oci_service_manager_proxy.ServiceManagerProxyClient {
	return m.GetClient("oci_service_manager_proxy.ServiceManagerProxyClient").(*oci_service_manager_proxy.ServiceManagerProxyClient)
}
