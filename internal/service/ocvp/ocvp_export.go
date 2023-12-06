package ocvp

import (
	oci_ocvp "github.com/oracle/oci-go-sdk/v65/ocvp"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("ocvp", ocvpResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportOcvpSddcHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_ocvp_sddc",
	DatasourceClass:        "oci_ocvp_sddcs",
	DatasourceItemsAttr:    "sddc_collection",
	ResourceAbbreviation:   "sddc",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_ocvp.LifecycleStatesActive),
		string(oci_ocvp.LifecycleStatesFailed),
	},
}

var exportOcvpEsxiHostHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_ocvp_esxi_host",
	DatasourceClass:        "oci_ocvp_esxi_hosts",
	DatasourceItemsAttr:    "esxi_host_collection",
	ResourceAbbreviation:   "esxi_host",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_ocvp.LifecycleStatesActive),
		string(oci_ocvp.LifecycleStatesFailed),
	},
}

var exportOcvpClusterHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_ocvp_cluster",
	DatasourceClass:        "oci_ocvp_clusters",
	DatasourceItemsAttr:    "cluster_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "cluster",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_ocvp.LifecycleStatesActive),
		string(oci_ocvp.LifecycleStatesFailed),
	},
}

var ocvpResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportOcvpSddcHints},
		{TerraformResourceHints: exportOcvpClusterHints},
	},
	"oci_ocvp_sddc": {
		{
			TerraformResourceHints: exportOcvpEsxiHostHints,
			DatasourceQueryParams: map[string]string{
				"sddc_id": "id",
			},
		},
	},
}
