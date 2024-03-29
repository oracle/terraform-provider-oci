package metering_computation

import (
	"fmt"

	oci_metering_computation "github.com/oracle/oci-go-sdk/v65/usageapi"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportMeteringComputationUsageStatementEmailRecipientsGroupHints.GetIdFn = getMeteringComputationUsageStatementEmailRecipientsGroupId
	tf_export.RegisterCompartmentGraphs("metering_computation", meteringComputationResourceGraph)
	tf_export.RegisterTenancyGraphs("metering_computation", meteringComputationResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

func getMeteringComputationUsageStatementEmailRecipientsGroupId(resource *tf_export.OCIResource) (string, error) {

	emailRecipientsGroupId, ok := resource.SourceAttributes["email_recipients_group_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find emailRecipientsGroupId for MeteringComputation UsageStatementEmailRecipientsGroup")
	}
	subscriptionId := resource.Parent.Id
	compartmentId := resource.Parent.CompartmentId
	return GetUsageStatementEmailRecipientsGroupCompositeId(emailRecipientsGroupId, subscriptionId, compartmentId), nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportMeteringComputationQueryHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_metering_computation_query",
	DatasourceClass:        "oci_metering_computation_queries",
	IsDatasourceCollection: true,
	DatasourceItemsAttr:    "query_collection",
	ResourceAbbreviation:   "query",
	RequireResourceRefresh: true,
}

var exportMeteringComputationCustomTableHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_metering_computation_custom_table",
	DatasourceClass:        "oci_metering_computation_custom_tables",
	DatasourceItemsAttr:    "custom_table_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "custom_table",
	RequireResourceRefresh: true,
}

var exportMeteringComputationScheduleHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_metering_computation_schedule",
	DatasourceClass:        "oci_metering_computation_schedules",
	DatasourceItemsAttr:    "schedule_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "schedule",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_metering_computation.ScheduleLifecycleStateActive),
	},
}

var exportMeteringComputationUsageCarbonEmissionsQueryHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_metering_computation_usage_carbon_emissions_query",
	DatasourceClass:        "oci_metering_computation_usage_carbon_emissions_queries",
	DatasourceItemsAttr:    "usage_carbon_emissions_query_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "usage_carbon_emissions_query",
	RequireResourceRefresh: true,
}

var exportMeteringComputationUsageCarbonEmissionHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_metering_computation_usage_carbon_emission",
	ResourceAbbreviation: "usage_carbon_emission",
}

var exportMeteringComputationUsageStatementEmailRecipientsGroupHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_metering_computation_usage_statement_email_recipients_group",
	DatasourceClass:        "oci_metering_computation_usage_statement_email_recipients_groups",
	DatasourceItemsAttr:    "email_recipients_group_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "usage_statement_email_recipients_group",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_metering_computation.EmailRecipientsGroupLifecycleStateActive),
	},
}

var meteringComputationResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_tenancy": {
		{TerraformResourceHints: exportMeteringComputationQueryHints},
		{TerraformResourceHints: exportMeteringComputationScheduleHints},
		{TerraformResourceHints: exportMeteringComputationUsageCarbonEmissionsQueryHints},
	},
	"oci_metering_computation_query": {
		{
			TerraformResourceHints: exportMeteringComputationCustomTableHints,
			DatasourceQueryParams: map[string]string{
				"saved_report_id": "id",
			},
		},
	},
}
