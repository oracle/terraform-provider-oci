package functions

import (
	oci_functions "github.com/oracle/oci-go-sdk/v65/functions"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("functions", functionsResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportFunctionsApplicationHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_functions_application",
	DatasourceClass:        "oci_functions_applications",
	DatasourceItemsAttr:    "applications",
	ResourceAbbreviation:   "application",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_functions.ApplicationLifecycleStateActive),
	},
}

var exportFunctionsFunctionHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_functions_function",
	DatasourceClass:        "oci_functions_functions",
	DatasourceItemsAttr:    "functions",
	ResourceAbbreviation:   "function",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_functions.FunctionLifecycleStateActive),
	},
}

var functionsResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportFunctionsApplicationHints},
	},
	"oci_functions_application": {
		{
			TerraformResourceHints: exportFunctionsFunctionHints,
			DatasourceQueryParams: map[string]string{
				"application_id": "id",
			},
		},
	},
}
