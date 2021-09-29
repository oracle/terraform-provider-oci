// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_generic_artifacts_content "github.com/oracle/oci-go-sdk/v48/genericartifactscontent"

	oci_common "github.com/oracle/oci-go-sdk/v48/common"
)

func init() {
	RegisterOracleClient("oci_generic_artifacts_content.GenericArtifactsContentClient", &OracleClient{initClientFn: initGenericartifactscontentGenericArtifactsContentClient})
}

func initGenericartifactscontentGenericArtifactsContentClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_generic_artifacts_content.NewGenericArtifactsContentClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) genericArtifactsContentClient() *oci_generic_artifacts_content.GenericArtifactsContentClient {
	return m.GetClient("oci_generic_artifacts_content.GenericArtifactsContentClient").(*oci_generic_artifacts_content.GenericArtifactsContentClient)
}
