// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_generic_artifacts_content "github.com/oracle/oci-go-sdk/v56/genericartifactscontent"

	oci_common "github.com/oracle/oci-go-sdk/v56/common"
)

func init() {
	RegisterOracleClient("oci_generic_artifacts_content.GenericArtifactsContentClient", &OracleClient{InitClientFn: initGenericartifactscontentGenericArtifactsContentClient})
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

	if serviceClientOverrides.HostUrlOverride != "" {
		client.Host = serviceClientOverrides.HostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) GenericArtifactsContentClient() *oci_generic_artifacts_content.GenericArtifactsContentClient {
	return m.GetClient("oci_generic_artifacts_content.GenericArtifactsContentClient").(*oci_generic_artifacts_content.GenericArtifactsContentClient)
}
