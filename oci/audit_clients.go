// Copyright (c) 2020, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	oci_audit "github.com/oracle/oci-go-sdk/audit"

	oci_common "github.com/oracle/oci-go-sdk/common"
)

func init() {
	RegisterOracleClient("oci_audit.AuditClient", &OracleClient{initClientFn: initAuditAuditClient})
}

func initAuditAuditClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_audit.NewAuditClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) auditClient() *oci_audit.AuditClient {
	return m.GetClient("oci_audit.AuditClient").(*oci_audit.AuditClient)
}
