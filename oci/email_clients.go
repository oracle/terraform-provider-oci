// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_email "github.com/oracle/oci-go-sdk/v25/email"

	oci_common "github.com/oracle/oci-go-sdk/v25/common"
)

func init() {
	RegisterOracleClient("oci_email.EmailClient", &OracleClient{initClientFn: initEmailEmailClient})
}

func initEmailEmailClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_email.NewEmailClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) emailClient() *oci_email.EmailClient {
	return m.GetClient("oci_email.EmailClient").(*oci_email.EmailClient)
}
