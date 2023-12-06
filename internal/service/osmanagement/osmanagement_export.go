package osmanagement

import (
	oci_osmanagement "github.com/oracle/oci-go-sdk/v65/osmanagement"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("osmanagement", osmanagementResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportOsmanagementManagedInstanceHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_osmanagement_managed_instance",
	DatasourceClass:        "oci_osmanagement_managed_instances",
	DatasourceItemsAttr:    "managed_instances",
	ResourceAbbreviation:   "managed_instance",
	RequireResourceRefresh: true,
}

var exportOsmanagementManagedInstanceGroupHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_osmanagement_managed_instance_group",
	DatasourceClass:      "oci_osmanagement_managed_instance_groups",
	DatasourceItemsAttr:  "managed_instance_groups",
	ResourceAbbreviation: "managed_instance_group",
	DiscoverableLifecycleStates: []string{
		string(oci_osmanagement.ListManagedInstanceGroupsLifecycleStateActive),
	},
}

var exportOsmanagementSoftwareSourceHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_osmanagement_software_source",
	DatasourceClass:        "oci_osmanagement_software_sources",
	DatasourceItemsAttr:    "software_sources",
	ResourceAbbreviation:   "software_source",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_osmanagement.ListSoftwareSourcesLifecycleStateActive),
	},
}

var osmanagementResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportOsmanagementManagedInstanceHints},
		{TerraformResourceHints: exportOsmanagementManagedInstanceGroupHints},
		{TerraformResourceHints: exportOsmanagementSoftwareSourceHints},
	},
}
