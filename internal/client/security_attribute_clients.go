// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_security_attribute "github.com/oracle/oci-go-sdk/v65/securityattribute"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_security_attribute.SecurityAttributeClient", &OracleClient{InitClientFn: initSecurityattributeSecurityAttributeClient})
}

func initSecurityattributeSecurityAttributeClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_security_attribute.NewSecurityAttributeClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) SecurityAttributeClient() *oci_security_attribute.SecurityAttributeClient {
	return m.GetClient("oci_security_attribute.SecurityAttributeClient").(*oci_security_attribute.SecurityAttributeClient)
}
