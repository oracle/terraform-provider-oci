// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_network_firewall "github.com/oracle/oci-go-sdk/v65/networkfirewall"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_network_firewall.NetworkFirewallClient", &OracleClient{InitClientFn: initNetworkfirewallNetworkFirewallClient})
}

func initNetworkfirewallNetworkFirewallClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_network_firewall.NewNetworkFirewallClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) NetworkFirewallClient() *oci_network_firewall.NetworkFirewallClient {
	return m.GetClient("oci_network_firewall.NetworkFirewallClient").(*oci_network_firewall.NetworkFirewallClient)
}
