package health_checks

import (
	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("health_checks", healthChecksResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportHealthChecksHttpMonitorHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_health_checks_http_monitor",
	DatasourceClass:        "oci_health_checks_http_monitors",
	DatasourceItemsAttr:    "http_monitors",
	ResourceAbbreviation:   "http_monitor",
	RequireResourceRefresh: true,
}

var exportHealthChecksPingMonitorHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_health_checks_ping_monitor",
	DatasourceClass:        "oci_health_checks_ping_monitors",
	DatasourceItemsAttr:    "ping_monitors",
	ResourceAbbreviation:   "ping_monitor",
	RequireResourceRefresh: true,
}

var healthChecksResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportHealthChecksHttpMonitorHints},
		{TerraformResourceHints: exportHealthChecksPingMonitorHints},
	},
}
