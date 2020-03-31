// Copyright (c) 2020, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	oci_nosql "github.com/oracle/oci-go-sdk/nosql"

	oci_common "github.com/oracle/oci-go-sdk/common"
)

func init() {
	RegisterOracleClient("oci_nosql.NosqlClient", &OracleClient{initClientFn: initNosqlNosqlClient})
}

func initNosqlNosqlClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_nosql.NewNosqlClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) nosqlClient() *oci_nosql.NosqlClient {
	return m.GetClient("oci_nosql.NosqlClient").(*oci_nosql.NosqlClient)
}
