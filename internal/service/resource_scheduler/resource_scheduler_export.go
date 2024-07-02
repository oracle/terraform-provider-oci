package resource_scheduler

import (
	oci_resource_scheduler "github.com/oracle/oci-go-sdk/v65/resourcescheduler"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("resource_scheduler", resourceSchedulerResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportResourceSchedulerScheduleHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_resource_scheduler_schedule",
	DatasourceClass:        "oci_resource_scheduler_schedules",
	DatasourceItemsAttr:    "schedule_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "schedule",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_resource_scheduler.ScheduleLifecycleStateActive),
	},
}

var resourceSchedulerResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportResourceSchedulerScheduleHints},
	},
}
