// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_nosql "github.com/oracle/oci-go-sdk/v39/nosql"

	oci_common "github.com/oracle/oci-go-sdk/v39/common"
)

func init() {
	RegisterOracleClient("oci_nosql.NosqlClient", &OracleClient{initClientFn: initNosqlNosqlClient})
}

func initNosqlNosqlClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_nosql.NewNosqlClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) nosqlClient() *oci_nosql.NosqlClient {
	return m.GetClient("oci_nosql.NosqlClient").(*oci_nosql.NosqlClient)
}
