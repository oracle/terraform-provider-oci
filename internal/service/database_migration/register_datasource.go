// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_migration

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_database_migration_assessment", DatabaseMigrationAssessmentDataSource())
	tfresource.RegisterDatasource("oci_database_migration_assessment_assessor", DatabaseMigrationAssessmentAssessorDataSource())
	tfresource.RegisterDatasource("oci_database_migration_assessment_assessor_check", DatabaseMigrationAssessmentAssessorCheckDataSource())
	tfresource.RegisterDatasource("oci_database_migration_assessment_assessor_check_affected_objects", DatabaseMigrationAssessmentAssessorCheckAffectedObjectsDataSource())
	tfresource.RegisterDatasource("oci_database_migration_assessment_assessor_checks", DatabaseMigrationAssessmentAssessorChecksDataSource())
	tfresource.RegisterDatasource("oci_database_migration_assessment_assessors", DatabaseMigrationAssessmentAssessorsDataSource())
	tfresource.RegisterDatasource("oci_database_migration_assessment_object_types", DatabaseMigrationAssessmentObjectTypesDataSource())
	tfresource.RegisterDatasource("oci_database_migration_assessments", DatabaseMigrationAssessmentsDataSource())

	tfresource.RegisterDatasource("oci_database_migration_connection", DatabaseMigrationConnectionDataSource())
	tfresource.RegisterDatasource("oci_database_migration_connection_databaseconnectiontypes", DatabaseMigrationConnectionDatabaseconnectiontypesDataSource())
	tfresource.RegisterDatasource("oci_database_migration_connections", DatabaseMigrationConnectionsDataSource())
	tfresource.RegisterDatasource("oci_database_migration_job", DatabaseMigrationJobDataSource())

	//tfresource.RegisterDatasource("oci_database_migration_agent", DatabaseMigrationAgentDataSource())
	//tfresource.RegisterDatasource("oci_database_migration_agent_images", DatabaseMigrationAgentImagesDataSource())
	//tfresource.RegisterDatasource("oci_database_migration_agents", DatabaseMigrationAgentsDataSource())
	//tfresource.RegisterDatasource("oci_database_migration_connection", DatabaseMigrationConnectionDataSource())
	//tfresource.RegisterDatasource("oci_database_migration_connections", DatabaseMigrationConnectionsDataSource())
	//tfresource.RegisterDatasource("oci_database_migration_job", DatabaseMigrationJobDataSource())

	tfresource.RegisterDatasource("oci_database_migration_job_advisor_report", DatabaseMigrationJobAdvisorReportDataSource())
	tfresource.RegisterDatasource("oci_database_migration_job_advisor_report_check_objects", DatabaseMigrationJobAdvisorReportCheckObjectsDataSource())
	tfresource.RegisterDatasource("oci_database_migration_job_advisor_report_checks", DatabaseMigrationJobAdvisorReportChecksDataSource())
	tfresource.RegisterDatasource("oci_database_migration_job_output", DatabaseMigrationJobOutputDataSource())
	tfresource.RegisterDatasource("oci_database_migration_jobs", DatabaseMigrationJobsDataSource())
	tfresource.RegisterDatasource("oci_database_migration_migration", DatabaseMigrationMigrationDataSource())
	tfresource.RegisterDatasource("oci_database_migration_migration_object_types", DatabaseMigrationMigrationObjectTypesDataSource())

	tfresource.RegisterDatasource("oci_database_migration_script", DatabaseMigrationScriptDataSource())

	tfresource.RegisterDatasource("oci_database_migration_migrations", DatabaseMigrationMigrationDataSource())

}
