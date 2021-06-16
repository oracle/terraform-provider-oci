// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_dns "github.com/oracle/oci-go-sdk/v42/dns"

	oci_common "github.com/oracle/oci-go-sdk/v42/common"
)

func init() {
	RegisterOracleClient("oci_dns.DnsClient", &OracleClient{initClientFn: initDnsDnsClient})
}

func initDnsDnsClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_dns.NewDnsClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) dnsClient() *oci_dns.DnsClient {
	return m.GetClient("oci_dns.DnsClient").(*oci_dns.DnsClient)
}
