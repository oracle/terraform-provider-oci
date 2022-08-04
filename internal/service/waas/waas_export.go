package waas

import (
	oci_waas "github.com/oracle/oci-go-sdk/v65/waas"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("waas", waasResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportWaasAddressListHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_waas_address_list",
	DatasourceClass:        "oci_waas_address_lists",
	DatasourceItemsAttr:    "address_lists",
	ResourceAbbreviation:   "address_list",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_waas.LifecycleStatesActive),
	},
}

var exportWaasCustomProtectionRuleHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_waas_custom_protection_rule",
	DatasourceClass:        "oci_waas_custom_protection_rules",
	DatasourceItemsAttr:    "custom_protection_rules",
	ResourceAbbreviation:   "custom_protection_rule",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_waas.LifecycleStatesActive),
	},
}

var exportWaasHttpRedirectHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_waas_http_redirect",
	DatasourceClass:      "oci_waas_http_redirects",
	DatasourceItemsAttr:  "http_redirects",
	ResourceAbbreviation: "http_redirect",
	DiscoverableLifecycleStates: []string{
		string(oci_waas.LifecycleStatesActive),
	},
}

var exportWaasWaasPolicyHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_waas_waas_policy",
	DatasourceClass:        "oci_waas_waas_policies",
	DatasourceItemsAttr:    "waas_policies",
	ResourceAbbreviation:   "waas_policy",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_waas.WaasPolicyLifecycleStateActive),
	},
}

var waasResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportWaasAddressListHints},
		{TerraformResourceHints: exportWaasCustomProtectionRuleHints},
		{TerraformResourceHints: exportWaasHttpRedirectHints},
		{TerraformResourceHints: exportWaasWaasPolicyHints},
	},
}
