package stack_monitoring

import (
	oci_stack_monitoring "github.com/oracle/oci-go-sdk/v65/stackmonitoring"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("stack_monitoring", stackMonitoringResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportStackMonitoringMonitoredResourceHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_stack_monitoring_monitored_resource",
	DatasourceClass:      "oci_stack_monitoring_monitored_resource",
	ResourceAbbreviation: "monitored_resource",
	DiscoverableLifecycleStates: []string{
		string(oci_stack_monitoring.ResourceLifecycleStateActive),
	},
}

var exportStackMonitoringDiscoveryJobHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_stack_monitoring_discovery_job",
	DatasourceClass:        "oci_stack_monitoring_discovery_jobs",
	DatasourceItemsAttr:    "discovery_job_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "discovery_job",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_stack_monitoring.LifecycleStateActive),
		string(oci_stack_monitoring.LifecycleStateFailed),
	},
}

var exportStackMonitoringMonitoredResourcesListMemberHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_stack_monitoring_monitored_resources_list_member",
	ResourceAbbreviation: "monitored_resources_list_member",
}

var exportStackMonitoringMonitoredResourcesSearchAssociationHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_stack_monitoring_monitored_resources_search_association",
	ResourceAbbreviation: "monitored_resources_search_association",
}

var exportStackMonitoringMonitoredResourcesSearchHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_stack_monitoring_monitored_resources_search",
	ResourceAbbreviation: "monitored_resources_search",
}

var exportStackMonitoringMonitoredResourcesAssociateMonitoredResourceHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_stack_monitoring_monitored_resources_associate_monitored_resource",
	ResourceAbbreviation: "monitored_resources_associate_monitored_resource",
}

var stackMonitoringResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportStackMonitoringDiscoveryJobHints},
	},
	"oci_stack_monitoring_monitored_resource": {
		{
			TerraformResourceHints: exportStackMonitoringMonitoredResourceHints,
			DatasourceQueryParams: map[string]string{
				"monitored_resource_id": "id",
			},
		},
	},
}
