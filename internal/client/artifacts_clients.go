// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_artifacts "github.com/oracle/oci-go-sdk/v56/artifacts"

	oci_common "github.com/oracle/oci-go-sdk/v56/common"
)

func init() {
	RegisterOracleClient("oci_artifacts.ArtifactsClient", &OracleClient{InitClientFn: initArtifactsArtifactsClient})
}

func initArtifactsArtifactsClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_artifacts.NewArtifactsClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) ArtifactsClient() *oci_artifacts.ArtifactsClient {
	return m.GetClient("oci_artifacts.ArtifactsClient").(*oci_artifacts.ArtifactsClient)
}
