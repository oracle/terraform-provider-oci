package network_firewall

import (
	oci_network_firewall "github.com/oracle/oci-go-sdk/v65/networkfirewall"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("network_firewall", networkFirewallResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportNetworkFirewallNetworkFirewallPolicyHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_network_firewall_network_firewall_policy",
	DatasourceClass:        "oci_network_firewall_network_firewall_policies",
	DatasourceItemsAttr:    "network_firewall_policy_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "network_firewall_policy",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_network_firewall.LifecycleStateActive),
	},
}

var exportNetworkFirewallNetworkFirewallHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_network_firewall_network_firewall",
	DatasourceClass:        "oci_network_firewall_network_firewalls",
	DatasourceItemsAttr:    "network_firewall_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "network_firewall",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_network_firewall.LifecycleStateActive),
	},
}

var networkFirewallResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportNetworkFirewallNetworkFirewallPolicyHints},
		{TerraformResourceHints: exportNetworkFirewallNetworkFirewallHints},
	},
}
