package database_migration

import (
	"context"
	"fmt"
	"strings"

	oci_database_migration "github.com/oracle/oci-go-sdk/v65/databasemigration"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func init() {
	exportDatabaseMigrationJobAdvisorReportCheckHints.GetIdFn = getDatabaseMigrationJobAdvisorReportCheckId
	exportDatabaseMigrationMigrationHints.FindResourcesOverrideFn = findDatabaseMigrationMigrations
	exportDatabaseMigrationAssessmentHints.FindResourcesOverrideFn = findDatabaseMigrationAssessments
	tf_export.RegisterCompartmentGraphs("database_migration", databaseMigrationResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

func getDatabaseMigrationJobAdvisorReportCheckId(resource *tf_export.OCIResource) (string, error) {

	advisorReportCheckId, ok := resource.SourceAttributes["advisor_report_check_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find advisorReportCheckId for DatabaseMigration JobAdvisorReportCheck")
	}
	jobId := resource.Parent.Id
	return GetJobAdvisorReportCheckCompositeId(advisorReportCheckId, jobId), nil
}

func findDatabaseMigrationMigrations(ctx *tf_export.ResourceDiscoveryContext, tfMeta *tf_export.TerraformResourceAssociation, parent *tf_export.OCIResource, resourceGraph *tf_export.TerraformResourceGraph) ([]*tf_export.OCIResource, error) {
	request := oci_database_migration.ListMigrationsRequest{
		CompartmentId: &parent.CompartmentId,
	}
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_migration")

	response, err := ctx.Clients.DatabaseMigrationClient().ListMigrations(context.Background(), request)
	if err != nil {
		return nil, err
	}

	request.Page = response.OpcNextPage
	for request.Page != nil {
		listResponse, err := ctx.Clients.DatabaseMigrationClient().ListMigrations(context.Background(), request)
		if err != nil {
			return nil, err
		}
		response.Items = append(response.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	migrationResource := tf_export.ResourcesMap[tfMeta.ResourceClass]
	resources := make([]*tf_export.OCIResource, 0, len(response.Items))
	for _, migration := range response.Items {
		if migration == nil || migration.GetId() == nil {
			continue
		}
		if len(ctx.ExpectedResourceIds) > 0 {
			if _, expected := ctx.ExpectedResourceIds[*migration.GetId()]; !expected {
				continue
			}
		}

		d := migrationResource.TestResourceData()
		d.SetId(*migration.GetId())

		if diags := migrationResource.ReadContext(context.Background(), d, ctx.Clients); diags.HasError() {
			readErr := fmt.Errorf("%s", strings.Join(tf_export.ParseDiagToError(diags), " | "))
			ctx.AddErrorToList(&tf_export.ResourceDiscoveryError{
				ResourceType:   tfMeta.ResourceClass,
				ParentResource: parent.TerraformName,
				Error:          readErr,
				ResourceGraph:  resourceGraph,
			})
			continue
		}

		resource := &tf_export.OCIResource{
			CompartmentId:    parent.CompartmentId,
			SourceAttributes: tf_export.ConvertResourceDataToMap(migrationResource.Schema, d),
			RawResource:      migration,
			TerraformResource: tf_export.TerraformResource{
				Id:             d.Id(),
				TerraformClass: tfMeta.ResourceClass,
			},
			GetHclStringFn: tf_export.GetHclStringFromGenericMap,
			Parent:         parent,
		}

		resource.TerraformName, err = tf_export.GenerateTerraformNameFromResource(resource.SourceAttributes, migrationResource.Schema)
		if err != nil {
			resource.TerraformName = tf_export.CheckDuplicateResourceName(fmt.Sprintf("%s_%s", parent.TerraformName, d.Id()))
		}

		resources = append(resources, resource)
	}

	return resources, nil
}

func findDatabaseMigrationAssessments(ctx *tf_export.ResourceDiscoveryContext, tfMeta *tf_export.TerraformResourceAssociation, parent *tf_export.OCIResource, resourceGraph *tf_export.TerraformResourceGraph) ([]*tf_export.OCIResource, error) {
	request := oci_database_migration.ListAssessmentsRequest{
		CompartmentId: &parent.CompartmentId,
	}
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_migration")

	response, err := ctx.Clients.DatabaseMigrationClient().ListAssessments(context.Background(), request)
	if err != nil {
		return nil, err
	}

	request.Page = response.OpcNextPage
	for request.Page != nil {
		listResponse, err := ctx.Clients.DatabaseMigrationClient().ListAssessments(context.Background(), request)
		if err != nil {
			return nil, err
		}
		response.Items = append(response.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	assessmentResource := tf_export.ResourcesMap[tfMeta.ResourceClass]
	resources := make([]*tf_export.OCIResource, 0, len(response.Items))
	for _, assessment := range response.Items {
		if assessment == nil || assessment.GetId() == nil {
			continue
		}
		if len(ctx.ExpectedResourceIds) > 0 {
			if _, expected := ctx.ExpectedResourceIds[*assessment.GetId()]; !expected {
				continue
			}
		} else {
			switch assessment.GetLifecycleState() {
			case oci_database_migration.AssessmentLifecycleStatesActive,
				oci_database_migration.AssessmentLifecycleStatesSucceeded,
				oci_database_migration.AssessmentLifecycleStatesNeedsAttention:
			default:
				continue
			}
		}

		d := assessmentResource.TestResourceData()
		d.SetId(*assessment.GetId())

		if diags := assessmentResource.ReadContext(context.Background(), d, ctx.Clients); diags.HasError() {
			readErr := fmt.Errorf("%s", strings.Join(tf_export.ParseDiagToError(diags), " | "))
			ctx.AddErrorToList(&tf_export.ResourceDiscoveryError{
				ResourceType:   tfMeta.ResourceClass,
				ParentResource: parent.TerraformName,
				Error:          readErr,
				ResourceGraph:  resourceGraph,
			})
			continue
		}

		resource := &tf_export.OCIResource{
			CompartmentId:    parent.CompartmentId,
			SourceAttributes: tf_export.ConvertResourceDataToMap(assessmentResource.Schema, d),
			RawResource:      assessment,
			TerraformResource: tf_export.TerraformResource{
				Id:             d.Id(),
				TerraformClass: tfMeta.ResourceClass,
			},
			GetHclStringFn: tf_export.GetHclStringFromGenericMap,
			Parent:         parent,
		}

		resource.TerraformName, err = tf_export.GenerateTerraformNameFromResource(resource.SourceAttributes, assessmentResource.Schema)
		if err != nil {
			resource.TerraformName = tf_export.CheckDuplicateResourceName(fmt.Sprintf("%s_%s", parent.TerraformName, d.Id()))
		}

		resources = append(resources, resource)
	}

	return resources, nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportDatabaseMigrationConnectionHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_database_migration_connection",
	DatasourceClass:        "oci_database_migration_connections",
	DatasourceItemsAttr:    "connection_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "connection",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_database_migration.LifecycleStatesActive),
	},
}

var exportDatabaseMigrationMigrationHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_database_migration_migration",
	DatasourceClass:        "oci_database_migration_migrations",
	DatasourceItemsAttr:    "migration_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "migration",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_database_migration.LifecycleStatesActive),
	},
}

var exportDatabaseMigrationAssessmentHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_database_migration_assessment",
	DatasourceClass:        "oci_database_migration_assessments",
	DatasourceItemsAttr:    "assessment_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "assessment",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_database_migration.AssessmentLifecycleStatesActive),
		string(oci_database_migration.AssessmentLifecycleStatesSucceeded),
		string(oci_database_migration.AssessmentLifecycleStatesNeedsAttention),
	},
}

var exportDatabaseMigrationAssessmentAssessorActionHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_database_migration_assessment_assessor_action",
	ResourceAbbreviation: "assessment_assessor_action",
}

var exportDatabaseMigrationJobAdvisorReportCheckHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_database_migration_job_advisor_report_check",
	DatasourceClass:        "oci_database_migration_job_advisor_report_checks",
	DatasourceItemsAttr:    "advisor_report_check_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "job_advisor_report_check",
}

var databaseMigrationResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDatabaseMigrationMigrationHints},
		{TerraformResourceHints: exportDatabaseMigrationAssessmentHints},
		{TerraformResourceHints: exportDatabaseMigrationConnectionHints},
	},
}
