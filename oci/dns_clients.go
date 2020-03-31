// Copyright (c) 2020, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	oci_dns "github.com/oracle/oci-go-sdk/dns"

	oci_common "github.com/oracle/oci-go-sdk/common"
)

func init() {
	RegisterOracleClient("oci_dns.DnsClient", &OracleClient{initClientFn: initDnsDnsClient})
}

func initDnsDnsClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_dns.NewDnsClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) dnsClient() *oci_dns.DnsClient {
	return m.GetClient("oci_dns.DnsClient").(*oci_dns.DnsClient)
}
