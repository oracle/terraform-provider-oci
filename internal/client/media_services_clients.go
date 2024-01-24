// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_media_services "github.com/oracle/oci-go-sdk/v65/mediaservices"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_media_services.MediaServicesClient", &OracleClient{InitClientFn: initMediaservicesMediaServicesClient})
}

func initMediaservicesMediaServicesClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_media_services.NewMediaServicesClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) MediaServicesClient() *oci_media_services.MediaServicesClient {
	return m.GetClient("oci_media_services.MediaServicesClient").(*oci_media_services.MediaServicesClient)
}
