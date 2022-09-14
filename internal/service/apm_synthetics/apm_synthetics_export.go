package apm_synthetics

import (
	"fmt"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportApmSyntheticsScriptHints.GetIdFn = getApmSyntheticsScriptId
	exportApmSyntheticsMonitorHints.GetIdFn = getApmSyntheticsMonitorId
	tf_export.RegisterCompartmentGraphs("apm_synthetics", apmSyntheticsResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

func getApmSyntheticsScriptId(resource *tf_export.OCIResource) (string, error) {

	scriptId, ok := resource.SourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find scriptId for ApmSynthetics Script")
	}
	apmDomainId, ok := resource.SourceAttributes["apm_domain_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find apmDomainId for ApmSynthetics Script")
	}
	return GetScriptCompositeId(scriptId, apmDomainId), nil
}

func getApmSyntheticsMonitorId(resource *tf_export.OCIResource) (string, error) {

	monitorId, ok := resource.SourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find monitorId for ApmSynthetics Monitor")
	}
	apmDomainId, ok := resource.SourceAttributes["apm_domain_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find apmDomainId for ApmSynthetics Monitor")
	}
	return GetMonitorCompositeId(monitorId, apmDomainId), nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportApmSyntheticsScriptHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_apm_synthetics_script",
	DatasourceClass:        "oci_apm_synthetics_scripts",
	DatasourceItemsAttr:    "script_collection",
	ResourceAbbreviation:   "script",
	RequireResourceRefresh: true,
}

var exportApmSyntheticsMonitorHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_apm_synthetics_monitor",
	DatasourceClass:        "oci_apm_synthetics_monitors",
	DatasourceItemsAttr:    "monitor_collection",
	ResourceAbbreviation:   "monitor",
	RequireResourceRefresh: true,
}

var exportApmSyntheticsDedicatedVantagePointHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_apm_synthetics_dedicated_vantage_point",
	DatasourceClass:        "oci_apm_synthetics_dedicated_vantage_points",
	DatasourceItemsAttr:    "dedicated_vantage_point_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "dedicated_vantage_point",
	RequireResourceRefresh: true,
}

var apmSyntheticsResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {},
}
