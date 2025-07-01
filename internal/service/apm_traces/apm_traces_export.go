package apm_traces

import (
	"fmt"

	oci_apm_traces "github.com/oracle/oci-go-sdk/v65/apmtraces"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportApmTracesScheduledQueryHints.GetIdFn = getApmTracesScheduledQueryId
	tf_export.RegisterCompartmentGraphs("apm_traces", apmTracesResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

func getApmTracesScheduledQueryId(resource *tf_export.OCIResource) (string, error) {

	scheduledQueryId, ok := resource.SourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find scheduledQueryId for ApmTraces ScheduledQuery")
	}

	apmDomainId, ok := resource.SourceAttributes["apm_domain_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find apmDomainId for ApmConfig Config")
	}
	return GetScheduledQueryCompositeId(apmDomainId, scheduledQueryId), nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportApmTracesScheduledQueryHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_apm_traces_scheduled_query",
	DatasourceClass:        "oci_apm_traces_scheduled_queries",
	DatasourceItemsAttr:    "scheduled_query_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "scheduled_query",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_apm_traces.LifecycleStatesActive),
	},
}

var apmTracesResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {},
}
