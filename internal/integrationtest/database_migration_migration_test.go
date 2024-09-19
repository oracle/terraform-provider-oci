// // // Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// // // Licensed under the Mozilla Public License v2.0
package integrationtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database_migration "github.com/oracle/oci-go-sdk/v65/databasemigration"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseMigrationMigrationRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_database_migration_migration", "test_migration", acctest.Required, acctest.Create, DatabaseMigrationMigrationRepresentationRDS)

	DatabaseMigrationMigrationResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_migration_migration", "test_migration", acctest.Optional, acctest.Update, DatabaseMigrationMigrationRepresentationRDS)

	DatabaseMigrationMigrationSingularDataSourceRepresentation = map[string]interface{}{
		"migration_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_migration_migration.test_migration.id}`},
	}

	DatabaseMigrationMigrationDataSourceRepresentation = map[string]interface{}{
		"migration_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_migration_migration.test_migration.id}`},
	}
	DatabaseMigrationMigrationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_migration_migration.test_migration.id}`}},
	}

	DatabaseMigrationMigrationRepresentation = map[string]interface{}{
		"compartment_id":                          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"database_combination":                    acctest.Representation{RepType: acctest.Required, Create: `ORACLE`},
		"source_database_connection_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.source_connection_oracle_id}`},
		"target_database_connection_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.target_connection_oracle_id}`},
		"type":                                    acctest.Representation{RepType: acctest.Required, Create: `ONLINE`, Update: `OFFLINE`},
		"advanced_parameters":                     acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseMigrationMigrationAdvancedParametersRepresentation},
		"advisor_settings":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseMigrationMigrationAdvisorSettingsRepresentation},
		"data_transfer_medium_details":            acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseMigrationMigrationDataTransferMediumDetailsRepresentation},
		"description":                             acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":                            acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"include_objects":                         acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseMigrationMigrationIncludeObjectsRepresentation},
		"initial_load_settings":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseMigrationMigrationInitialLoadSettingsOracleRepresentation},
		"source_container_database_connection_id": acctest.Representation{RepType: acctest.Required, Create: `${var.source_connection_container_oracle_id}`},
	}

	DatabaseMigrationMigrationRepresentationMySQL = map[string]interface{}{
		"compartment_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"database_combination":          acctest.Representation{RepType: acctest.Required, Create: `MYSQL`},
		"source_database_connection_id": acctest.Representation{RepType: acctest.Required, Create: `${var.source_connection_mysql_id}`},
		"target_database_connection_id": acctest.Representation{RepType: acctest.Required, Create: `${var.target_connection_mysql_id}`},
		"type":                          acctest.Representation{RepType: acctest.Required, Create: `ONLINE`, Update: `OFFLINE`},
		"advisor_settings":              acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseMigrationMigrationAdvisorSettingsRepresentation},
		"bulk_include_exclude_data":     acctest.Representation{RepType: acctest.Optional, Create: `bulkIncludeExcludeData`},
		"data_transfer_medium_details":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseMigrationMigrationDataTransferMediumDetailsRepresentation},
		"description":                   acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":                  acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"exclude_objects":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseMigrationMigrationExcludeObjectsRepresentation},
		"include_objects":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseMigrationMigrationIncludeObjectsRepresentation},
		"initial_load_settings":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseMigrationMigrationInitialLoadSettingsRepresentation},
	}
	DatabaseMigrationMigrationRepresentationRDS = map[string]interface{}{
		"compartment_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"database_combination":          acctest.Representation{RepType: acctest.Required, Create: `ORACLE`},
		"source_database_connection_id": acctest.Representation{RepType: acctest.Required, Create: `${var.source_connection_oracle_rds_id}`},
		"target_database_connection_id": acctest.Representation{RepType: acctest.Required, Create: `${var.target_connection_oracle_id}`},
		"type":                          acctest.Representation{RepType: acctest.Required, Create: `ONLINE`, Update: `OFFLINE`},
		"advanced_parameters":           acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseMigrationMigrationAdvancedParametersRepresentation},
		"advisor_settings":              acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseMigrationMigrationAdvisorSettingsRepresentation},
		"data_transfer_medium_details":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseMigrationMigrationDataTransferMediumDetailsAWS3Representation},
		"description":                   acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":                  acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"include_objects":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseMigrationMigrationIncludeObjectsRepresentation},
		"initial_load_settings":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseMigrationMigrationInitialLoadSettingsOracleRepresentation},
	}
	DatabaseMigrationMigrationAdvancedParametersRepresentation = map[string]interface{}{
		"data_type": acctest.Representation{RepType: acctest.Required, Create: `STRING`, Update: `INTEGER`},
		"name":      acctest.Representation{RepType: acctest.Required, Create: `DATAPUMPSETTINGS_METADATAONLY`, Update: `DATAPUMPSETTINGS_DUMPFILESIZE`},
		"value":     acctest.Representation{RepType: acctest.Required, Create: `True`, Update: `5000`},
	}
	DatabaseMigrationMigrationAdvisorSettingsRepresentation = map[string]interface{}{
		"is_ignore_errors": acctest.Representation{RepType: acctest.Optional, Update: `true`},
		"is_skip_advisor":  acctest.Representation{RepType: acctest.Optional, Update: `true`},
	}
	DatabaseMigrationMigrationDataTransferMediumDetailsRepresentation = map[string]interface{}{
		"type":                  acctest.Representation{RepType: acctest.Required, Create: `OBJECT_STORAGE`, Update: `OBJECT_STORAGE`},
		"object_storage_bucket": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseMigrationMigrationDataTransferMediumDetailsObjectStorageBucketRepresentation},
		"source":                acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseMigrationMigrationDataTransferMediumDetailsSourceRepresentation},
	}
	DatabaseMigrationMigrationDataTransferMediumDetailsAWS3Representation = map[string]interface{}{
		"type":                  acctest.Representation{RepType: acctest.Required, Create: `AWS_S3`, Update: `AWS_S3`},
		"object_storage_bucket": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseMigrationMigrationDataTransferMediumDetailsObjectStorageBucketRepresentation},
		"name":                  acctest.Representation{RepType: acctest.Required, Create: `rdsbucket`, Update: `rdsbucket2`},
		"region":                acctest.Representation{RepType: acctest.Required, Create: `us-east-1`, Update: `us-east-2`},
		"secret_access_key":     acctest.Representation{RepType: acctest.Required, Create: `12345/12345`, Update: `6789/6789`},
		"access_key_id":         acctest.Representation{RepType: acctest.Required, Create: `12345`, Update: `6789`},
	}

	DatabaseMigrationMigrationExcludeObjectsRepresentation = map[string]interface{}{
		"object": acctest.Representation{RepType: acctest.Required, Create: `.*`},
		"is_omit_excluded_table_from_replication": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"owner":  acctest.Representation{RepType: acctest.Optional, Create: `owner`},
		"schema": acctest.Representation{RepType: acctest.Optional, Create: `schema`},
		"type":   acctest.Representation{RepType: acctest.Optional, Create: `ALL`},
	}
	DatabaseMigrationMigrationGgsDetailsRepresentation = map[string]interface{}{
		"acceptable_lag": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"replicat":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseMigrationMigrationGgsDetailsReplicatRepresentation},
	}
	DatabaseMigrationMigrationHubDetailsRepresentation = map[string]interface{}{
		"key_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.kms_key_id}`},
		"rest_admin_credentials": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseMigrationMigrationHubDetailsRestAdminCredentialsRepresentation},
		"url":                    acctest.Representation{RepType: acctest.Required, Create: `url`, Update: `url2`},
		"vault_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.kms_vault_id}`},
		"acceptable_lag":         acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"extract":                acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseMigrationMigrationHubDetailsExtractRepresentation},
		"replicat":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseMigrationMigrationHubDetailsReplicatRepresentation},
	}
	DatabaseMigrationMigrationIncludeObjectsRepresentation = map[string]interface{}{
		"object": acctest.Representation{RepType: acctest.Required, Create: `.*`},
		"is_omit_excluded_table_from_replication": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"owner":  acctest.Representation{RepType: acctest.Optional, Create: `owner`},
		"schema": acctest.Representation{RepType: acctest.Optional, Create: `schema`},
		"type":   acctest.Representation{RepType: acctest.Optional, Create: `ALL`},
	}
	DatabaseMigrationMigrationInitialLoadSettingsRepresentation = map[string]interface{}{
		"job_mode":                   acctest.Representation{RepType: acctest.Required, Create: `FULL`, Update: `SCHEMA`},
		"handle_grant_errors":        acctest.Representation{RepType: acctest.Optional, Create: `ABORT`, Update: `DROP_ACCOUNT`},
		"is_consistent":              acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_ignore_existing_objects": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_tz_utc":                  acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	DatabaseMigrationMigrationInitialLoadSettingsOracleRepresentation = map[string]interface{}{
		"job_mode":                acctest.Representation{RepType: acctest.Required, Create: `SCHEMA`, Update: `SCHEMA`},
		"data_pump_parameters":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseMigrationMigrationInitialLoadSettingsDataPumpParametersRepresentation},
		"export_directory_object": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseMigrationMigrationInitialLoadSettingsExportDirectoryObjectAWS3Representation},
	}
	DatabaseMigrationMigrationDataTransferMediumDetailsObjectStorageBucketRepresentation = map[string]interface{}{
		"bucket":    acctest.Representation{RepType: acctest.Optional, Create: `${var.bucket_mysql_id}`, Update: `${var.bucket_mysql_id}`},
		"namespace": acctest.Representation{RepType: acctest.Optional, Create: `namespace`, Update: `namespace2`},
	}
	DatabaseMigrationMigrationDataTransferMediumDetailsSourceRepresentation = map[string]interface{}{
		"kind":            acctest.Representation{RepType: acctest.Required, Create: `OCI_CLI`, Update: `OCI_CLI`},
		"oci_home":        acctest.Representation{RepType: acctest.Optional, Create: `ociHome`, Update: `ociHome2`},
		"wallet_location": acctest.Representation{RepType: acctest.Optional, Create: `walletLocation`, Update: `walletLocation2`},
	}
	DatabaseMigrationMigrationDataTransferMediumDetailsTargetRepresentation = map[string]interface{}{
		"kind":            acctest.Representation{RepType: acctest.Required, Create: `CURL`, Update: `OCI_CLI`},
		"oci_home":        acctest.Representation{RepType: acctest.Optional, Create: `ociHome`, Update: `ociHome2`},
		"wallet_location": acctest.Representation{RepType: acctest.Optional, Create: `walletLocation`, Update: `walletLocation2`},
	}
	DatabaseMigrationMigrationGgsDetailsExtractRepresentation = map[string]interface{}{
		"long_trans_duration": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"performance_profile": acctest.Representation{RepType: acctest.Optional, Create: `LOW`, Update: `MEDIUM`},
	}
	DatabaseMigrationMigrationGgsDetailsReplicatRepresentation = map[string]interface{}{
		"performance_profile": acctest.Representation{RepType: acctest.Optional, Create: `LOW`, Update: `HIGH`},
	}
	DatabaseMigrationMigrationHubDetailsRestAdminCredentialsRepresentation = map[string]interface{}{
		"password": acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`, Update: `BEstrO0ng_#12`},
		"username": acctest.Representation{RepType: acctest.Required, Create: `username`, Update: `username2`},
	}
	DatabaseMigrationMigrationHubDetailsExtractRepresentation = map[string]interface{}{
		"long_trans_duration": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"performance_profile": acctest.Representation{RepType: acctest.Optional, Create: `LOW`, Update: `MEDIUM`},
	}
	DatabaseMigrationMigrationHubDetailsReplicatRepresentation = map[string]interface{}{
		"performance_profile": acctest.Representation{RepType: acctest.Optional, Create: `LOW`, Update: `HIGH`},
	}
	DatabaseMigrationMigrationInitialLoadSettingsDataPumpParametersRepresentation = map[string]interface{}{
		"estimate":                  acctest.Representation{RepType: acctest.Optional, Create: `BLOCKS`, Update: `STATISTICS`},
		"export_parallelism_degree": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"import_parallelism_degree": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"is_cluster":                acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"table_exists_action":       acctest.Representation{RepType: acctest.Optional, Create: `TRUNCATE`, Update: `REPLACE`},
	}
	DatabaseMigrationMigrationInitialLoadSettingsExportDirectoryObjectAWS3Representation = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
	}
	DatabaseMigrationMigrationInitialLoadSettingsExportDirectoryObjectRepresentation = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
		"path": acctest.Representation{RepType: acctest.Optional, Create: `/u01/app/oracle/dumpdir`, Update: `/u01/app/oracle/dumpdir2`},
	}
	DatabaseMigrationMigrationInitialLoadSettingsImportDirectoryObjectRepresentation = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Optional, Create: `name`, Update: `name2`},
		"path": acctest.Representation{RepType: acctest.Optional, Create: `path`, Update: `path2`},
	}
	DatabaseMigrationMigrationInitialLoadSettingsMetadataRemapsRepresentation = map[string]interface{}{
		"new_value": acctest.Representation{RepType: acctest.Optional, Create: `newValue`, Update: `newValue2`},
		"old_value": acctest.Representation{RepType: acctest.Optional, Create: `oldValue`, Update: `oldValue2`},
		"type":      acctest.Representation{RepType: acctest.Optional, Create: `SCHEMA`, Update: `TABLESPACE`},
	}
	DatabaseMigrationMigrationInitialLoadSettingsTablespaceDetailsRepresentation = map[string]interface{}{
		"target_type": acctest.Representation{RepType: acctest.Required, Create: `ADB_S_REMAP`, Update: `ADB_S_REMAP`},
	}
)

// issue-routing-tag: database_migration/default
func TestDatabaseMigrationMigrationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseMigrationMigrationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithBlankDefault("compartment_id_for_update")
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	kmsKeyId := utils.GetEnvSettingWithBlankDefault("kms_key_id")
	kmsKeyIdVariableStr := fmt.Sprintf("variable \"kms_key_id\" { default = \"%s\" }\n", kmsKeyId)

	kmsVaultId := utils.GetEnvSettingWithBlankDefault("kms_vault_id")
	kmsVaultIdVariableStr := fmt.Sprintf("variable \"kms_vault_id\" { default = \"%s\" }\n", kmsVaultId)

	sourceConnectionId := utils.GetEnvSettingWithBlankDefault("source_connection_mysql_id")
	sourceConnectionIdVariableStr := fmt.Sprintf("variable \"source_connection_mysql_id\" { default = \"%s\" }\n", sourceConnectionId)

	targetConnectionId := utils.GetEnvSettingWithBlankDefault("target_connection_mysql_id")
	targetConnectionIdVariableStr := fmt.Sprintf("variable \"target_connection_mysql_id\" { default = \"%s\" }\n", targetConnectionId)

	sourceConnectionOracleId := utils.GetEnvSettingWithBlankDefault("source_connection_oracle_id")
	sourceConnectionOracleIdVariableStr := fmt.Sprintf("variable \"source_connection_oracle_id\" { default = \"%s\" }\n", sourceConnectionOracleId)

	sourceConnectionRDSId := utils.GetEnvSettingWithBlankDefault("source_connection_oracle_rds_id")
	sourceConnectionRDSIdVariableStr := fmt.Sprintf("variable \"source_connection_oracle_rds_id\" { default = \"%s\" }\n", sourceConnectionRDSId)

	sourceConnectionContainerOracleId := utils.GetEnvSettingWithBlankDefault("source_connection_container_oracle_id")
	sourceConnectionContainerOracleIdVariableStr := fmt.Sprintf("variable \"source_connection_container_oracle_id\" { default = \"%s\" }\n", sourceConnectionContainerOracleId)

	targetConnectionOracleId := utils.GetEnvSettingWithBlankDefault("target_connection_oracle_id")
	targetConnectionOracleIdVariableStr := fmt.Sprintf("variable \"target_connection_oracle_id\" { default = \"%s\" }\n", targetConnectionOracleId)

	bucketId := utils.GetEnvSettingWithBlankDefault("bucket_mysql_id")
	bucketIdVariableStr := fmt.Sprintf("variable \"bucket_mysql_id\" { default = \"%s\" }\n", bucketId)

	resourceName := "oci_database_migration_migration.test_migration"
	datasourceName := "data.oci_database_migration_migrations.test_migrations"
	singularDatasourceName := "data.oci_database_migration_migration.test_migration"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+kmsKeyIdVariableStr+kmsVaultIdVariableStr+sourceConnectionIdVariableStr+targetConnectionIdVariableStr+bucketIdVariableStr+sourceConnectionOracleIdVariableStr+sourceConnectionContainerOracleIdVariableStr+targetConnectionOracleIdVariableStr+sourceConnectionRDSIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_database_migration_migration", "test_migration", acctest.Optional, acctest.Create, DatabaseMigrationMigrationRepresentationRDS), "databasemigration", "migration", t)

	acctest.ResourceTest(t, testAccCheckDatabaseMigrationMigrationDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + kmsKeyIdVariableStr + kmsVaultIdVariableStr + sourceConnectionIdVariableStr + targetConnectionIdVariableStr + bucketIdVariableStr + sourceConnectionOracleIdVariableStr + sourceConnectionContainerOracleIdVariableStr + targetConnectionOracleIdVariableStr + sourceConnectionRDSIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_migration", "test_migration", acctest.Required, acctest.Create, DatabaseMigrationMigrationRepresentationRDS),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "database_combination", "ORACLE"),
				resource.TestCheckResourceAttrSet(resourceName, "source_database_connection_id"),
				resource.TestCheckResourceAttrSet(resourceName, "target_database_connection_id"),
				resource.TestCheckResourceAttr(resourceName, "type", "ONLINE"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + kmsKeyIdVariableStr + kmsVaultIdVariableStr + sourceConnectionIdVariableStr + targetConnectionIdVariableStr + bucketIdVariableStr + sourceConnectionOracleIdVariableStr + sourceConnectionContainerOracleIdVariableStr + targetConnectionOracleIdVariableStr + sourceConnectionRDSIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_migration", "test_migration", acctest.Optional, acctest.Create, DatabaseMigrationMigrationRepresentationRDS),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "advanced_parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "advanced_parameters.0.data_type", "STRING"),
				resource.TestCheckResourceAttr(resourceName, "advanced_parameters.0.name", "DATAPUMPSETTINGS_METADATAONLY"),
				resource.TestCheckResourceAttr(resourceName, "advanced_parameters.0.value", "True"),
				resource.TestCheckResourceAttr(resourceName, "advisor_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "advisor_settings.0.is_ignore_errors", "false"),
				resource.TestCheckResourceAttr(resourceName, "advisor_settings.0.is_skip_advisor", "false"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "data_transfer_medium_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_transfer_medium_details.0.object_storage_bucket.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_transfer_medium_details.0.object_storage_bucket.0.namespace", "namespace"),
				resource.TestCheckResourceAttr(resourceName, "data_transfer_medium_details.0.type", "AWS_S3"),
				resource.TestCheckResourceAttr(resourceName, "data_transfer_medium_details.0.name", "rdsbucket"),
				resource.TestCheckResourceAttr(resourceName, "data_transfer_medium_details.0.region", "us-east-1"),
				resource.TestCheckResourceAttr(resourceName, "data_transfer_medium_details.0.secret_access_key", "12345/12345"),
				resource.TestCheckResourceAttr(resourceName, "data_transfer_medium_details.0.access_key_id", "12345"),
				resource.TestCheckResourceAttr(resourceName, "database_combination", "ORACLE"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "initial_load_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "initial_load_settings.0.is_consistent", "false"),
				resource.TestCheckResourceAttr(resourceName, "initial_load_settings.0.is_ignore_existing_objects", "false"),
				resource.TestCheckResourceAttr(resourceName, "initial_load_settings.0.is_tz_utc", "false"),
				resource.TestCheckResourceAttr(resourceName, "initial_load_settings.0.job_mode", "SCHEMA"),
				resource.TestCheckResourceAttr(resourceName, "include_objects.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "include_objects.0.is_omit_excluded_table_from_replication", "false"),
				resource.TestCheckResourceAttr(resourceName, "include_objects.0.object", ".*"),
				resource.TestCheckResourceAttr(resourceName, "include_objects.0.owner", "owner"),
				resource.TestCheckResourceAttr(resourceName, "include_objects.0.schema", "schema"),
				resource.TestCheckResourceAttr(resourceName, "include_objects.0.type", "ALL"),
				resource.TestCheckResourceAttrSet(resourceName, "source_database_connection_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_database_connection_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "ONLINE"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					time.Sleep(1 * time.Minute)
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + kmsKeyIdVariableStr + kmsVaultIdVariableStr + sourceConnectionIdVariableStr + targetConnectionIdVariableStr + bucketIdVariableStr + sourceConnectionOracleIdVariableStr + sourceConnectionContainerOracleIdVariableStr + targetConnectionOracleIdVariableStr + sourceConnectionRDSIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_migration", "test_migration", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatabaseMigrationMigrationRepresentationRDS, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "advanced_parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "advanced_parameters.0.data_type", "STRING"),
				resource.TestCheckResourceAttr(resourceName, "advanced_parameters.0.name", "DATAPUMPSETTINGS_METADATAONLY"),
				resource.TestCheckResourceAttr(resourceName, "advanced_parameters.0.value", "True"),
				resource.TestCheckResourceAttr(resourceName, "advisor_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "advisor_settings.0.is_ignore_errors", "false"),
				resource.TestCheckResourceAttr(resourceName, "advisor_settings.0.is_skip_advisor", "false"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "data_transfer_medium_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_transfer_medium_details.0.object_storage_bucket.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_transfer_medium_details.0.object_storage_bucket.0.namespace", "namespace"),
				resource.TestCheckResourceAttr(resourceName, "data_transfer_medium_details.0.type", "AWS_S3"),
				resource.TestCheckResourceAttr(resourceName, "data_transfer_medium_details.0.name", "rdsbucket"),
				resource.TestCheckResourceAttr(resourceName, "data_transfer_medium_details.0.region", "us-east-1"),
				resource.TestCheckResourceAttr(resourceName, "data_transfer_medium_details.0.secret_access_key", "12345/12345"),
				resource.TestCheckResourceAttr(resourceName, "data_transfer_medium_details.0.access_key_id", "12345"),
				resource.TestCheckResourceAttr(resourceName, "database_combination", "ORACLE"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "include_objects.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "include_objects.0.is_omit_excluded_table_from_replication", "false"),
				resource.TestCheckResourceAttr(resourceName, "include_objects.0.object", ".*"),
				resource.TestCheckResourceAttr(resourceName, "include_objects.0.owner", "owner"),
				resource.TestCheckResourceAttr(resourceName, "include_objects.0.schema", "schema"),
				resource.TestCheckResourceAttr(resourceName, "include_objects.0.type", "ALL"),
				resource.TestCheckResourceAttr(resourceName, "initial_load_settings.#", "1"),

				resource.TestCheckResourceAttr(resourceName, "initial_load_settings.0.is_consistent", "false"),
				resource.TestCheckResourceAttr(resourceName, "initial_load_settings.0.is_ignore_existing_objects", "false"),
				resource.TestCheckResourceAttr(resourceName, "initial_load_settings.0.is_tz_utc", "false"),
				resource.TestCheckResourceAttr(resourceName, "initial_load_settings.0.job_mode", "SCHEMA"),
				resource.TestCheckResourceAttrSet(resourceName, "source_database_connection_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_database_connection_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "ONLINE"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + kmsKeyIdVariableStr + kmsVaultIdVariableStr + sourceConnectionIdVariableStr + targetConnectionIdVariableStr + bucketIdVariableStr + sourceConnectionOracleIdVariableStr + sourceConnectionContainerOracleIdVariableStr + targetConnectionOracleIdVariableStr + sourceConnectionRDSIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_migration", "test_migration", acctest.Optional, acctest.Update, DatabaseMigrationMigrationRepresentationRDS),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "advanced_parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "advanced_parameters.0.data_type", "INTEGER"),
				resource.TestCheckResourceAttr(resourceName, "advanced_parameters.0.name", "DATAPUMPSETTINGS_DUMPFILESIZE"),
				resource.TestCheckResourceAttr(resourceName, "advanced_parameters.0.value", "5000"),
				resource.TestCheckResourceAttr(resourceName, "advisor_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "advisor_settings.0.is_ignore_errors", "true"),
				resource.TestCheckResourceAttr(resourceName, "advisor_settings.0.is_skip_advisor", "true"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "data_transfer_medium_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_transfer_medium_details.0.object_storage_bucket.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "data_transfer_medium_details.0.object_storage_bucket.0.namespace", "namespace2"),
				resource.TestCheckResourceAttr(resourceName, "data_transfer_medium_details.0.type", "AWS_S3"),
				resource.TestCheckResourceAttr(resourceName, "data_transfer_medium_details.0.name", "rdsbucket2"),
				resource.TestCheckResourceAttr(resourceName, "data_transfer_medium_details.0.region", "us-east-2"),
				resource.TestCheckResourceAttr(resourceName, "data_transfer_medium_details.0.secret_access_key", "6789/6789"),
				resource.TestCheckResourceAttr(resourceName, "data_transfer_medium_details.0.access_key_id", "6789"),
				resource.TestCheckResourceAttr(resourceName, "database_combination", "ORACLE"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),

				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "include_objects.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "include_objects.0.is_omit_excluded_table_from_replication", "false"),
				resource.TestCheckResourceAttr(resourceName, "include_objects.0.object", ".*"),
				resource.TestCheckResourceAttr(resourceName, "include_objects.0.owner", "owner"),
				resource.TestCheckResourceAttr(resourceName, "include_objects.0.schema", "schema"),
				resource.TestCheckResourceAttr(resourceName, "include_objects.0.type", "ALL"),
				resource.TestCheckResourceAttr(resourceName, "initial_load_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "initial_load_settings.0.is_consistent", "false"),
				resource.TestCheckResourceAttr(resourceName, "initial_load_settings.0.is_ignore_existing_objects", "false"),
				resource.TestCheckResourceAttr(resourceName, "initial_load_settings.0.is_tz_utc", "false"),
				resource.TestCheckResourceAttr(resourceName, "initial_load_settings.0.job_mode", "SCHEMA"),
				resource.TestCheckResourceAttrSet(resourceName, "source_database_connection_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_database_connection_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "OFFLINE"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_migration_migrations", "test_migrations", acctest.Optional, acctest.Update, DatabaseMigrationMigrationDataSourceRepresentation) +
				compartmentIdVariableStr + kmsKeyIdVariableStr + kmsVaultIdVariableStr + sourceConnectionIdVariableStr + targetConnectionIdVariableStr + bucketIdVariableStr + sourceConnectionOracleIdVariableStr + sourceConnectionContainerOracleIdVariableStr + targetConnectionOracleIdVariableStr + sourceConnectionRDSIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_migration", "test_migration", acctest.Optional, acctest.Update, DatabaseMigrationMigrationRepresentationRDS),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_migration_migration", "test_migration", acctest.Required, acctest.Create, DatabaseMigrationMigrationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + kmsKeyIdVariableStr + sourceConnectionIdVariableStr + targetConnectionIdVariableStr + bucketIdVariableStr + kmsVaultIdVariableStr + sourceConnectionOracleIdVariableStr + sourceConnectionContainerOracleIdVariableStr + targetConnectionOracleIdVariableStr + sourceConnectionRDSIdVariableStr + DatabaseMigrationMigrationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "migration_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "advanced_parameters.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "advanced_parameters.0.data_type", "INTEGER"),
				resource.TestCheckResourceAttr(singularDatasourceName, "advanced_parameters.0.name", "DATAPUMPSETTINGS_DUMPFILESIZE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "advanced_parameters.0.value", "5000"),
				resource.TestCheckResourceAttr(singularDatasourceName, "advisor_settings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "advisor_settings.0.is_ignore_errors", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "advisor_settings.0.is_skip_advisor", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_transfer_medium_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_transfer_medium_details.0.object_storage_bucket.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_transfer_medium_details.0.object_storage_bucket.0.namespace", "namespace2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_transfer_medium_details.0.type", "AWS_S3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_transfer_medium_details.0.name", "rdsbucket2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_transfer_medium_details.0.region", "us-east-2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_transfer_medium_details.0.secret_access_key", "6789/6789"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_transfer_medium_details.0.access_key_id", "6789"),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_combination", "ORACLE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "initial_load_settings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "initial_load_settings.0.is_consistent", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "initial_load_settings.0.is_ignore_existing_objects", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "initial_load_settings.0.is_tz_utc", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "initial_load_settings.0.job_mode", "SCHEMA"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "OFFLINE"),
			),
		},
		// verify resource import
		{
			Config:            config + DatabaseMigrationMigrationRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"bulk_include_exclude_data",
				"exclude_objects",
				"include_objects",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatabaseMigrationMigrationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DatabaseMigrationClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_migration_migration" {
			noResourceFound = false
			request := oci_database_migration.GetMigrationRequest{}

			tmp := rs.Primary.ID
			request.MigrationId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_migration")

			response, err := client.GetMigration(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database_migration.LifecycleStatesDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.GetLifecycleState())]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.GetLifecycleState())
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DatabaseMigrationMigration") {
		resource.AddTestSweepers("DatabaseMigrationMigration", &resource.Sweeper{
			Name:         "DatabaseMigrationMigration",
			Dependencies: acctest.DependencyGraph["migration"],
			F:            sweepDatabaseMigrationMigrationResource,
		})
	}
}

func sweepDatabaseMigrationMigrationResource(compartment string) error {
	databaseMigrationClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseMigrationClient()
	migrationIds, err := getDatabaseMigrationMigrationIds(compartment)
	if err != nil {
		return err
	}
	for _, migrationId := range migrationIds {
		if ok := acctest.SweeperDefaultResourceId[migrationId]; !ok {
			deleteMigrationRequest := oci_database_migration.DeleteMigrationRequest{}

			deleteMigrationRequest.MigrationId = &migrationId

			deleteMigrationRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_migration")
			_, error := databaseMigrationClient.DeleteMigration(context.Background(), deleteMigrationRequest)
			if error != nil {
				fmt.Printf("Error deleting Migration %s %s, It is possible that the resource is already deleted. Please verify manually \n", migrationId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &migrationId, DatabaseMigrationMigrationSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseMigrationMigrationSweepResponseFetchOperation, "database_migration", true)
		}
	}
	return nil
}

func getDatabaseMigrationMigrationIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "MigrationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseMigrationClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseMigrationClient()

	listMigrationsRequest := oci_database_migration.ListMigrationsRequest{}
	listMigrationsRequest.CompartmentId = &compartmentId
	listMigrationsRequest.LifecycleState = oci_database_migration.ListMigrationsLifecycleStateActive
	listMigrationsResponse, err := databaseMigrationClient.ListMigrations(context.Background(), listMigrationsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Migration list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, migration := range listMigrationsResponse.Items {
		id := *migration.GetId()
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "MigrationId", id)
	}
	return resourceIds, nil
}

func DatabaseMigrationMigrationSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if migrationResponse, ok := response.Response.(oci_database_migration.GetMigrationResponse); ok {
		return migrationResponse.GetLifecycleState() != oci_database_migration.MigrationLifecycleStatesDeleted
	}
	return false
}

func DatabaseMigrationMigrationSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseMigrationClient().GetMigration(context.Background(), oci_database_migration.GetMigrationRequest{
		MigrationId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
