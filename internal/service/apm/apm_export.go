package apm

import (
	oci_apm "github.com/oracle/oci-go-sdk/v65/apmcontrolplane"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("apm", apmResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportApmApmDomainHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_apm_apm_domain",
	DatasourceClass:      "oci_apm_apm_domains",
	DatasourceItemsAttr:  "apm_domains",
	ResourceAbbreviation: "apm_domain",
	DiscoverableLifecycleStates: []string{
		string(oci_apm.LifecycleStatesActive),
	},
}

var apmResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportApmApmDomainHints},
	},
}
