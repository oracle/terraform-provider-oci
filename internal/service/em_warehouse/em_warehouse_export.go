package em_warehouse

import (
	oci_em_warehouse "github.com/oracle/oci-go-sdk/v65/emwarehouse"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("em_warehouse", emWarehouseResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportEmWarehouseEmWarehouseHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_em_warehouse_em_warehouse",
	DatasourceClass:        "oci_em_warehouse_em_warehouses",
	DatasourceItemsAttr:    "em_warehouse_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "em_warehouse",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_em_warehouse.EmWarehouseLifecycleStateActive),
	},
}

var emWarehouseResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportEmWarehouseEmWarehouseHints},
	},
}
