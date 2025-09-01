package jms_utils

import (
	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("jms_utils", jmsUtilsResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportJmsUtilsSubscriptionAcknowledgmentConfigurationHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_jms_utils_subscription_acknowledgment_configuration",
	DatasourceClass:      "oci_jms_utils_subscription_acknowledgment_configuration",
	ResourceAbbreviation: "subscription_acknowledgment_configuration",
}

var exportJmsUtilsAnalyzeApplicationsConfigurationHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_jms_utils_analyze_applications_configuration",
	DatasourceClass:      "oci_jms_utils_analyze_applications_configuration",
	ResourceAbbreviation: "analyze_applications_configuration",
}

var jmsUtilsResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportJmsUtilsSubscriptionAcknowledgmentConfigurationHints},
		{TerraformResourceHints: exportJmsUtilsAnalyzeApplicationsConfigurationHints},
	},
}
