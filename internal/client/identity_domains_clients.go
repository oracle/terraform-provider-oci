// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_identity_domains "github.com/oracle/oci-go-sdk/v65/identitydomains"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_identity_domains.IdentityDomainsClient", &OracleClient{InitClientFn: initIdentitydomainsIdentityDomainsClient})
}

func initIdentitydomainsIdentityDomainsClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_identity_domains.NewIdentityDomainsClientWithConfigurationProvider(configProvider, "DUMMY_ENDPOINT")
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

func (m *OracleClients) IdentityDomainsClient() *oci_identity_domains.IdentityDomainsClient {
	return m.GetClient("oci_identity_domains.IdentityDomainsClient").(*oci_identity_domains.IdentityDomainsClient)
}
