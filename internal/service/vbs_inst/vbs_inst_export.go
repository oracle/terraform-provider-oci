package vbs_inst

import (
	oci_vbs_inst "github.com/oracle/oci-go-sdk/v65/vbsinst"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterTenancyGraphs("vbs_inst", vbsInstResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportVbsInstVbsInstanceHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_vbs_inst_vbs_instance",
	DatasourceClass:        "oci_vbs_inst_vbs_instances",
	DatasourceItemsAttr:    "vbs_instance_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "vbs_instance",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_vbs_inst.LifecycleStateActive),
	},
}

var vbsInstResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_tenancy": {
		{TerraformResourceHints: exportVbsInstVbsInstanceHints},
	},
}
