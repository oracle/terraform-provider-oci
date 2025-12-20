package database_migration

import (
	"fmt"

	oci_database_migration "github.com/oracle/oci-go-sdk/v65/databasemigration"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportDatabaseMigrationJobAdvisorReportCheckHints.GetIdFn = getDatabaseMigrationJobAdvisorReportCheckId
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
