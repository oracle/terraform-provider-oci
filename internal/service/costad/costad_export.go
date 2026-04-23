package costad

import (
	oci_costad "github.com/oracle/oci-go-sdk/v65/costad"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("costad", costadResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportCostadCostAlertSubscriptionHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_costad_cost_alert_subscription",
	DatasourceClass:        "oci_costad_cost_alert_subscriptions",
	DatasourceItemsAttr:    "cost_alert_subscription_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "cost_alert_subscription",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_costad.CostAlertSubscriptionLifecycleStateActive),
	},
}

var exportCostadCostAnomalyMonitorHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_costad_cost_anomaly_monitor",
	DatasourceClass:        "oci_costad_cost_anomaly_monitors",
	DatasourceItemsAttr:    "cost_anomaly_monitor_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "cost_anomaly_monitor",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_costad.CostAnomalyMonitorLifecycleStateActive),
	},
}

var exportCostadCostAnomalyEventHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_costad_cost_anomaly_event",
	DatasourceClass:        "oci_costad_cost_anomaly_events",
	DatasourceItemsAttr:    "cost_anomaly_event_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "cost_anomaly_event",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_costad.CostAnomalyEventLifecycleStateActive),
	},
}

var costadResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportCostadCostAlertSubscriptionHints},
		{TerraformResourceHints: exportCostadCostAnomalyMonitorHints},
		{TerraformResourceHints: exportCostadCostAnomalyEventHints},
	},
}
