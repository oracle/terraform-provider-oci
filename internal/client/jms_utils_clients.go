// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_jms_utils "github.com/oracle/oci-go-sdk/v65/jmsutils"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_jms_utils.JmsUtilsClient", &OracleClient{InitClientFn: initJmsutilsJmsUtilsClient})
}

func initJmsutilsJmsUtilsClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_jms_utils.NewJmsUtilsClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) JmsUtilsClient() *oci_jms_utils.JmsUtilsClient {
	return m.GetClient("oci_jms_utils.JmsUtilsClient").(*oci_jms_utils.JmsUtilsClient)
}
