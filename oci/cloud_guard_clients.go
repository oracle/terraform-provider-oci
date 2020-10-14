// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_cloud_guard "github.com/oracle/oci-go-sdk/v27/cloudguard"

	oci_common "github.com/oracle/oci-go-sdk/v27/common"
)

func init() {
	RegisterOracleClient("oci_cloud_guard.CloudGuardClient", &OracleClient{initClientFn: initCloudguardCloudGuardClient})
}

func initCloudguardCloudGuardClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_cloud_guard.NewCloudGuardClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) cloudGuardClient() *oci_cloud_guard.CloudGuardClient {
	return m.GetClient("oci_cloud_guard.CloudGuardClient").(*oci_cloud_guard.CloudGuardClient)
}
