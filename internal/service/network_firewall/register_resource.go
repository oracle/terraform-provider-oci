// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package network_firewall

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_network_firewall_network_firewall", NetworkFirewallNetworkFirewallResource())
	tfresource.RegisterResource("oci_network_firewall_network_firewall_policy", NetworkFirewallNetworkFirewallPolicyResource())
}
