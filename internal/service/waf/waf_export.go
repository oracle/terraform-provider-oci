package waf

import (
	oci_waf "github.com/oracle/oci-go-sdk/v65/waf"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("waf", wafResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportWafWebAppFirewallPolicyHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_waf_web_app_firewall_policy",
	DatasourceClass:        "oci_waf_web_app_firewall_policies",
	DatasourceItemsAttr:    "web_app_firewall_policy_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "web_app_firewall_policy",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_waf.WebAppFirewallPolicyLifecycleStateActive),
	},
}

var exportWafWebAppFirewallHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_waf_web_app_firewall",
	DatasourceClass:        "oci_waf_web_app_firewalls",
	DatasourceItemsAttr:    "web_app_firewall_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "web_app_firewall",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_waf.WebAppFirewallLifecycleStateActive),
	},
}

var exportWafNetworkAddressListHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_waf_network_address_list",
	DatasourceClass:        "oci_waf_network_address_lists",
	DatasourceItemsAttr:    "network_address_list_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "network_address_list",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_waf.NetworkAddressListLifecycleStateActive),
	},
}

var wafResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportWafWebAppFirewallPolicyHints},
		{TerraformResourceHints: exportWafWebAppFirewallHints},
		{TerraformResourceHints: exportWafNetworkAddressListHints},
	},
}
