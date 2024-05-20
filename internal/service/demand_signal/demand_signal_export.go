package demand_signal

import (
	oci_demand_signal "github.com/oracle/oci-go-sdk/v65/demandsignal"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("demand_signal", demandSignalResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportDemandSignalOccDemandSignalHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_demand_signal_occ_demand_signal",
	DatasourceClass:        "oci_demand_signal_occ_demand_signals",
	DatasourceItemsAttr:    "occ_demand_signal_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "occ_demand_signal",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_demand_signal.OccDemandSignalLifecycleStateActive),
	},
}

var demandSignalResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDemandSignalOccDemandSignalHints},
	},
}
