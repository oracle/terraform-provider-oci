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
	ResourceClass:          "oci_stack_monitoring_monitored_resource",
	DatasourceClass:        "oci_stack_monitoring_monitored_resources",
	DatasourceItemsAttr:    "monitored_resource_collection",
	IsDatasourceCollection: true,
	RequireResourceRefresh: true,
	ResourceAbbreviation:   "monitored_resource",
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

var exportStackMonitoringConfigHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_stack_monitoring_config",
	DatasourceClass:        "oci_stack_monitoring_configs",
	DatasourceItemsAttr:    "config_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "config",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_stack_monitoring.ConfigLifecycleStateActive),
	},
}
var exportStackMonitoringMonitoredResourceTaskHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_stack_monitoring_monitored_resource_task",
	DatasourceClass:        "oci_stack_monitoring_monitored_resource_tasks",
	DatasourceItemsAttr:    "monitored_resource_tasks_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "monitored_resource_task",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_stack_monitoring.MonitoredResourceTaskLifecycleStateSucceeded),
		string(oci_stack_monitoring.MonitoredResourceTaskLifecycleStateNeedsAttention),
	},
}

var exportStackMonitoringMonitoredResourceTypeHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_stack_monitoring_monitored_resource_type",
	DatasourceClass:        "oci_stack_monitoring_monitored_resource_types",
	DatasourceItemsAttr:    "monitored_resource_types_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "monitored_resource_type",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_stack_monitoring.ResourceTypeLifecycleStateActive),
	},
}

var exportStackMonitoringMetricExtensionHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_stack_monitoring_metric_extension",
	DatasourceClass:        "oci_stack_monitoring_metric_extensions",
	DatasourceItemsAttr:    "metric_extension_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "metric_extension",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_stack_monitoring.MetricExtensionLifeCycleStatesActive),
	},
}

var exportStackMonitoringBaselineableMetricHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_stack_monitoring_baselineable_metric",
	DatasourceClass:        "oci_stack_monitoring_baselineable_metrics",
	DatasourceItemsAttr:    "baselineable_metric_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "baselineable_metric",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_stack_monitoring.LifecycleStateActive),
	},
}

var exportStackMonitoringProcessSetHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_stack_monitoring_process_set",
	DatasourceClass:        "oci_stack_monitoring_process_sets",
	DatasourceItemsAttr:    "process_set_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "process_set",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_stack_monitoring.LifecycleStateActive),
	},
}

var exportStackMonitoringMaintenanceWindowHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_stack_monitoring_maintenance_window",
	DatasourceClass:        "oci_stack_monitoring_maintenance_windows",
	DatasourceItemsAttr:    "maintenance_window_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "maintenance_window",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_stack_monitoring.MaintenanceWindowLifecycleStateActive),
		string(oci_stack_monitoring.MaintenanceWindowLifecycleStateNeedsAttention),
	},
}

var exportStackMonitoringMaintenanceWindowsRetryFailedOperationHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_stack_monitoring_maintenance_windows_retry_failed_operation",
	ResourceAbbreviation: "maintenance_windows_retry_failed_operation",
}

var exportStackMonitoringMaintenanceWindowsStopHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_stack_monitoring_maintenance_windows_stop",
	ResourceAbbreviation: "maintenance_windows_stop",
}

var stackMonitoringResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportStackMonitoringDiscoveryJobHints},
		{TerraformResourceHints: exportStackMonitoringConfigHints},
		{TerraformResourceHints: exportStackMonitoringMetricExtensionHints},
		{TerraformResourceHints: exportStackMonitoringBaselineableMetricHints},
		{TerraformResourceHints: exportStackMonitoringProcessSetHints},
		{TerraformResourceHints: exportStackMonitoringMaintenanceWindowHints},
	},
	"oci_stack_monitoring_monitored_resource": {
		{
			TerraformResourceHints: exportStackMonitoringMonitoredResourceHints,
			DatasourceQueryParams: map[string]string{
				"monitored_resource_id": "id",
			},
		},
	},
	"oci_stack_monitoring_monitored_resource_type": {
		{
			TerraformResourceHints: exportStackMonitoringMonitoredResourceTypeHints,
			DatasourceQueryParams: map[string]string{
				"monitored_resource_type_id": "id",
			},
		},
	},
	"oci_stack_monitoring_monitored_resource_task": {
		{
			TerraformResourceHints: exportStackMonitoringMonitoredResourceTaskHints,
			DatasourceQueryParams: map[string]string{
				"monitored_resource_task_id": "id",
			},
		},
	},
}
