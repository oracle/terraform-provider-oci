// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_jms_java_downloads "github.com/oracle/oci-go-sdk/v65/jmsjavadownloads"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_jms_java_downloads.JavaDownloadClient", &OracleClient{InitClientFn: initJmsjavadownloadsJavaDownloadClient})
}

func initJmsjavadownloadsJavaDownloadClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_jms_java_downloads.NewJavaDownloadClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) JavaDownloadClient() *oci_jms_java_downloads.JavaDownloadClient {
	return m.GetClient("oci_jms_java_downloads.JavaDownloadClient").(*oci_jms_java_downloads.JavaDownloadClient)
}
