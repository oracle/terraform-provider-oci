package jms_utils

import (
	"fmt"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportJmsUtilsSubscriptionAcknowledgmentConfigurationHints.GetIdFn = getJmsUtilsSubscriptionAcknowledgmentConfigurationId
	exportJmsUtilsAnalyzeApplicationsConfigurationHints.GetIdFn = getJmsUtilsAnalyzeApplicationsConfigurationId
	tf_export.RegisterCompartmentGraphs("jms_utils", jmsUtilsResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework
func getJmsUtilsSubscriptionAcknowledgmentConfigurationId(resource *tf_export.OCIResource) (string, error) {
	compartmentId, ok := resource.SourceAttributes["compartment_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find compartmentId for Subscription Acknowledgment Configuration")
	}
	return compartmentId, nil
}

func getJmsUtilsAnalyzeApplicationsConfigurationId(resource *tf_export.OCIResource) (string, error) {
	compartmentId, ok := resource.SourceAttributes["compartment_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find compartmentId for Analyze Applications Configuration")
	}
	return compartmentId, nil
}

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
		{TerraformResourceHints: exportJmsUtilsSubscriptionAcknowledgmentConfigurationHints,
			DatasourceQueryParams: map[string]string{
				"compartment_id": "id",
			}},
		{TerraformResourceHints: exportJmsUtilsAnalyzeApplicationsConfigurationHints,
			DatasourceQueryParams: map[string]string{
				"compartment_id": "id",
			}},
	},
}
