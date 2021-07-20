// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_artifacts "github.com/oracle/oci-go-sdk/v45/artifacts"

	oci_common "github.com/oracle/oci-go-sdk/v45/common"
)

func init() {
	RegisterOracleClient("oci_artifacts.ArtifactsClient", &OracleClient{initClientFn: initArtifactsArtifactsClient})
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

	if serviceClientOverrides.hostUrlOverride != "" {
		client.Host = serviceClientOverrides.hostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) artifactsClient() *oci_artifacts.ArtifactsClient {
	return m.GetClient("oci_artifacts.ArtifactsClient").(*oci_artifacts.ArtifactsClient)
}
