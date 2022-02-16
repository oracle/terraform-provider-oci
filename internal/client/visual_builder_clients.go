// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_visual_builder "github.com/oracle/oci-go-sdk/v58/visualbuilder"

	oci_common "github.com/oracle/oci-go-sdk/v58/common"
)

func init() {
	RegisterOracleClient("oci_visual_builder.VbInstanceClient", &OracleClient{InitClientFn: initVisualbuilderVbInstanceClient})
}

func initVisualbuilderVbInstanceClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_visual_builder.NewVbInstanceClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) VbInstanceClient() *oci_visual_builder.VbInstanceClient {
	return m.GetClient("oci_visual_builder.VbInstanceClient").(*oci_visual_builder.VbInstanceClient)
}
