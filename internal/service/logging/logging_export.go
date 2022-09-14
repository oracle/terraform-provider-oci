package logging

import (
	"fmt"

	oci_logging "github.com/oracle/oci-go-sdk/v65/logging"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportLoggingLogHints.GetIdFn = getLoggingLogId
	exportLoggingLogHints.GetIdFn = getLogId
	tf_export.RegisterCompartmentGraphs("logging", loggingResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework
func getLogId(resource *tf_export.OCIResource) (string, error) {
	logId, ok := resource.SourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find log_id for Log")
	}
	logGroupId := resource.Parent.Id
	return GetLogCompositeId(logGroupId, logId), nil
}

func getLoggingLogId(resource *tf_export.OCIResource) (string, error) {

	logGroupId := resource.Parent.Id
	logId, ok := resource.SourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find logId for Logging Log")
	}
	return GetLogCompositeId(logGroupId, logId), nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportLoggingLogGroupHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_logging_log_group",
	DatasourceClass:      "oci_logging_log_groups",
	DatasourceItemsAttr:  "log_groups",
	ResourceAbbreviation: "log_group",
	DiscoverableLifecycleStates: []string{
		string(oci_logging.LogGroupLifecycleStateActive),
	},
}

var exportLoggingLogHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_logging_log",
	DatasourceClass:      "oci_logging_logs",
	DatasourceItemsAttr:  "logs",
	ResourceAbbreviation: "log",
	DiscoverableLifecycleStates: []string{
		string(oci_logging.LogLifecycleStateActive),
	},
}

var exportLoggingUnifiedAgentConfigurationHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_logging_unified_agent_configuration",
	DatasourceClass:        "oci_logging_unified_agent_configurations",
	DatasourceItemsAttr:    "unified_agent_configuration_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "unified_agent_configuration",
	RequireResourceRefresh: true,
}

var loggingResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportLoggingLogGroupHints},
		{TerraformResourceHints: exportLoggingUnifiedAgentConfigurationHints},
	},
	"oci_logging_log_group": {
		{
			TerraformResourceHints: exportLoggingLogHints,
			DatasourceQueryParams: map[string]string{
				"log_group_id": "id",
			},
		},
	},
}
