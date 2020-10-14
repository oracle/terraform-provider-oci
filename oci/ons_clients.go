// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_ons "github.com/oracle/oci-go-sdk/v27/ons"

	oci_common "github.com/oracle/oci-go-sdk/v27/common"
)

func init() {
	RegisterOracleClient("oci_ons.NotificationControlPlaneClient", &OracleClient{initClientFn: initOnsNotificationControlPlaneClient})
	RegisterOracleClient("oci_ons.NotificationDataPlaneClient", &OracleClient{initClientFn: initOnsNotificationDataPlaneClient})
}

func initOnsNotificationControlPlaneClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_ons.NewNotificationControlPlaneClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) notificationControlPlaneClient() *oci_ons.NotificationControlPlaneClient {
	return m.GetClient("oci_ons.NotificationControlPlaneClient").(*oci_ons.NotificationControlPlaneClient)
}

func initOnsNotificationDataPlaneClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_ons.NewNotificationDataPlaneClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) notificationDataPlaneClient() *oci_ons.NotificationDataPlaneClient {
	return m.GetClient("oci_ons.NotificationDataPlaneClient").(*oci_ons.NotificationDataPlaneClient)
}
