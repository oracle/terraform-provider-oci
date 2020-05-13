// Copyright (c) 2020, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	oci_identity "github.com/oracle/oci-go-sdk/identity"

	oci_common "github.com/oracle/oci-go-sdk/common"
)

func init() {
	RegisterOracleClient("oci_identity.IdentityClient", &OracleClient{initClientFn: initIdentityIdentityClient})
}

func initIdentityIdentityClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_identity.NewIdentityClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) identityClient() *oci_identity.IdentityClient {
	return m.GetClient("oci_identity.IdentityClient").(*oci_identity.IdentityClient)
}
