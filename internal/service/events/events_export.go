package events

import (
	oci_events "github.com/oracle/oci-go-sdk/v65/events"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("events", eventsResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportEventsRuleHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_events_rule",
	DatasourceClass:        "oci_events_rules",
	DatasourceItemsAttr:    "rules",
	ResourceAbbreviation:   "rule",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_events.RuleLifecycleStateActive),
	},
}

var eventsResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportEventsRuleHints},
	},
}
