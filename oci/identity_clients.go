// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_identity "github.com/oracle/oci-go-sdk/v46/identity"

	oci_common "github.com/oracle/oci-go-sdk/v46/common"
)

func init() {
	RegisterOracleClient("oci_identity.IdentityClient", &OracleClient{initClientFn: initIdentityIdentityClient})
}

func initIdentityIdentityClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_identity.NewIdentityClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) identityClient() *oci_identity.IdentityClient {
	return m.GetClient("oci_identity.IdentityClient").(*oci_identity.IdentityClient)
}
