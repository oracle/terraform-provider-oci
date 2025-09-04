// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_multicloud "github.com/oracle/oci-go-sdk/v65/multicloud"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_multicloud.MetadataClient", &OracleClient{InitClientFn: initMulticloudMetadataClient})
	RegisterOracleClient("oci_multicloud.MultiCloudsMetadataClient", &OracleClient{InitClientFn: initMulticloudMultiCloudsMetadataClient})
	RegisterOracleClient("oci_multicloud.OmhubNetworkAnchorClient", &OracleClient{InitClientFn: initMulticloudOmhubNetworkAnchorClient})
	RegisterOracleClient("oci_multicloud.OmhubResourceAnchorClient", &OracleClient{InitClientFn: initMulticloudOmhubResourceAnchorClient})
}

func initMulticloudMetadataClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_multicloud.NewMetadataClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) MetadataClient() *oci_multicloud.MetadataClient {
	return m.GetClient("oci_multicloud.MetadataClient").(*oci_multicloud.MetadataClient)
}

func initMulticloudMultiCloudsMetadataClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_multicloud.NewMultiCloudsMetadataClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) MultiCloudsMetadataClient() *oci_multicloud.MultiCloudsMetadataClient {
	return m.GetClient("oci_multicloud.MultiCloudsMetadataClient").(*oci_multicloud.MultiCloudsMetadataClient)
}

func initMulticloudOmhubNetworkAnchorClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_multicloud.NewOmhubNetworkAnchorClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) OmhubNetworkAnchorClient() *oci_multicloud.OmhubNetworkAnchorClient {
	return m.GetClient("oci_multicloud.OmhubNetworkAnchorClient").(*oci_multicloud.OmhubNetworkAnchorClient)
}

func initMulticloudOmhubResourceAnchorClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_multicloud.NewOmhubResourceAnchorClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) OmhubResourceAnchorClient() *oci_multicloud.OmhubResourceAnchorClient {
	return m.GetClient("oci_multicloud.OmhubResourceAnchorClient").(*oci_multicloud.OmhubResourceAnchorClient)
}
