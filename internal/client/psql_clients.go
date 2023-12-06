// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_psql "github.com/oracle/oci-go-sdk/v65/psql"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_psql.PostgresqlClient", &OracleClient{InitClientFn: initPsqlPostgresqlClient})
}

func initPsqlPostgresqlClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_psql.NewPostgresqlClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) PostgresqlClient() *oci_psql.PostgresqlClient {
	return m.GetClient("oci_psql.PostgresqlClient").(*oci_psql.PostgresqlClient)
}
