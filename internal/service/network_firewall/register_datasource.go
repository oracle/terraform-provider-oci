// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package network_firewall

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_network_firewall_network_firewall", NetworkFirewallNetworkFirewallDataSource())
	tfresource.RegisterDatasource("oci_network_firewall_network_firewall_policies", NetworkFirewallNetworkFirewallPoliciesDataSource())
	tfresource.RegisterDatasource("oci_network_firewall_network_firewall_policy", NetworkFirewallNetworkFirewallPolicyDataSource())
	tfresource.RegisterDatasource("oci_network_firewall_network_firewalls", NetworkFirewallNetworkFirewallsDataSource())
}
