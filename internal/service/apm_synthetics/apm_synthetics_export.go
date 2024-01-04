package apm_synthetics

import (
	"fmt"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportApmSyntheticsScriptHints.GetIdFn = getApmSyntheticsScriptId
	exportApmSyntheticsMonitorHints.GetIdFn = getApmSyntheticsMonitorId
	//exportApmSyntheticsDedicatedVantagePointHints.GetIdFn = getApmSyntheticsDedicatedVantagePointId
	exportApmSyntheticsOnPremiseVantagePointHints.GetIdFn = getApmSyntheticsOnPremiseVantagePointId
	exportApmSyntheticsOnPremiseVantagePointWorkerHints.GetIdFn = getApmSyntheticsOnPremiseVantagePointWorkerId
	tf_export.RegisterCompartmentGraphs("apm_synthetics", apmSyntheticsResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

func getApmSyntheticsScriptId(resource *tf_export.OCIResource) (string, error) {

	apmDomainId := resource.Parent.Id
	scriptId, ok := resource.SourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find scriptId for ApmSynthetics Script")
	}

	return GetScriptCompositeId(scriptId, apmDomainId), nil
}

func getApmSyntheticsMonitorId(resource *tf_export.OCIResource) (string, error) {

	apmDomainId := resource.Parent.Id
	monitorId, ok := resource.SourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find monitorId for ApmSynthetics Monitor")
	}

	return GetMonitorCompositeId(monitorId, apmDomainId), nil
}

func getApmSyntheticsOnPremiseVantagePointWorkerId(resource *tf_export.OCIResource) (string, error) {

	onPremiseVantagePointId := resource.Parent.Id
	workerId, ok := resource.SourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find id for ApmSynthetics OnPremiseVantagePointWorker")
	}

	apmDomainId := resource.Parent.Parent.Id

	return GetOnPremiseVantagePointWorkerCompositeId(onPremiseVantagePointId, workerId, apmDomainId), nil
}

func getApmSyntheticsOnPremiseVantagePointId(resource *tf_export.OCIResource) (string, error) {

	onPremiseVantagePointId, ok := resource.SourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find id for ApmSynthetics OnPremiseVantagePoint")
	}
	apmDomainId := resource.Parent.Id

	return GetOnPremiseVantagePointCompositeId(onPremiseVantagePointId, apmDomainId), nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportApmSyntheticsScriptHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_apm_synthetics_script",
	DatasourceClass:        "oci_apm_synthetics_scripts",
	DatasourceItemsAttr:    "script_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "script",
	RequireResourceRefresh: true,
}

var exportApmSyntheticsMonitorHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_apm_synthetics_monitor",
	DatasourceClass:        "oci_apm_synthetics_monitors",
	DatasourceItemsAttr:    "monitor_collection",
	IsDatasourceCollection: true,
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

var exportApmSyntheticsOnPremiseVantagePointWorkerHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_apm_synthetics_on_premise_vantage_point_worker",
	DatasourceClass:        "oci_apm_synthetics_on_premise_vantage_point_workers",
	DatasourceItemsAttr:    "worker_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "on_premise_vantage_point_worker",
	RequireResourceRefresh: true,
}

var exportApmSyntheticsOnPremiseVantagePointHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_apm_synthetics_on_premise_vantage_point",
	DatasourceClass:        "oci_apm_synthetics_on_premise_vantage_points",
	DatasourceItemsAttr:    "on_premise_vantage_point_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "on_premise_vantage_point",
	RequireResourceRefresh: true,
}

var apmSyntheticsResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {},
}
