package dblm

import (
	oci_dblm "github.com/oracle/oci-go-sdk/v65/dblm"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("dblm", dblmResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportDblmVulnerabilityScanHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_dblm_vulnerability_scan",
	DatasourceClass:        "oci_dblm_vulnerability_scans",
	DatasourceItemsAttr:    "vulnerability_scan_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "vulnerability_scan",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_dblm.VulnerabilityScanLifecycleStateActive),
	},
}

var dblmResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDblmVulnerabilityScanHints},
	},
}
