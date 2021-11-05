// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_identity "github.com/oracle/oci-go-sdk/v49/identity"

	oci_common "github.com/oracle/oci-go-sdk/v49/common"

	tf_client "github.com/terraform-providers/terraform-provider-oci/oci/client"
)

func init() {
	tf_client.RegisterOracleClient("oci_identity.IdentityClient", &tf_client.OracleClient{InitClientFn: initIdentityIdentityClient})
}

type OracleIdentityClients struct {
	*tf_client.OracleClients
}

func initIdentityIdentityClient(configProvider oci_common.ConfigurationProvider, configureClient tf_client.ConfigureClient, serviceClientOverrides tf_client.ServiceClientOverrides) (interface{}, error) {
	client, err := oci_identity.NewIdentityClientWithConfigurationProvider(configProvider)
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

func (m *OracleIdentityClients) identityClient() *oci_identity.IdentityClient {
	return m.GetClient("oci_identity.IdentityClient").(*oci_identity.IdentityClient)
}
