// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v58/common"
	oci_database_migration "github.com/oracle/oci-go-sdk/v58/databasemigration"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	MigrationRequiredOnlyResource = MigrationResourceDependenciesMig +
		acctest.GenerateResourceFromRepresentationMap("oci_database_migration_migration", "test_migration", acctest.Required, acctest.Create, migrationRepresentationMig)

	MigrationResourceConfig = MigrationResourceDependenciesMig +
		acctest.GenerateResourceFromRepresentationMap("oci_database_migration_migration", "test_migration", acctest.Optional, acctest.Update, migrationRepresentationMig)

	migrationSingularDataSourceRepresentation = map[string]interface{}{
		"migration_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_migration_migration.test_migration.id}`},
	}

	migrationDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: migrationDataSourceFilterRepresentation}}
	migrationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_migration_migration.test_migration.id}`}},
	}

	migrationRepresentationMig = map[string]interface{}{
		"compartment_id":                acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"source_database_connection_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_migration_connection.test_connection_source.id}`},
		"target_database_connection_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_migration_connection.test_connection.id}`},
		"type":                          acctest.Representation{RepType: acctest.Required, Create: `ONLINE`},
		"exclude_objects":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: migrationExcludeObjectsRepresentation},
		"data_transfer_medium_details":  acctest.RepresentationGroup{RepType: acctest.Required, Group: migrationDataTransferMediumDetailsRepresentation},
		"datapump_settings":             acctest.RepresentationGroup{RepType: acctest.Required, Group: migrationDatapumpSettingsRepresentation},
		"display_name":                  acctest.Representation{RepType: acctest.Optional, Create: `TF_ONLINE_MIG`, Update: `TF_ONLINE_MIG`},
		"golden_gate_details":           acctest.RepresentationGroup{RepType: acctest.Required, Group: migrationGoldenGateDetailsRepresentation},
		"vault_details":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: migrationVaultDetailsRepresentation},
	}
	migrationDataTransferMediumDetailsRepresentation = map[string]interface{}{
		"object_storage_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: migrationDataTransferMediumDetailsObjectStorageDetailsRepresentation},
	}
	migrationDataTransferMediumDetailsRepresentationBeforeCPAT = map[string]interface{}{
		"database_link_details":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: migrationDataTransferMediumDetailsDatabaseLinkDetailsRepresentation},
		"object_storage_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: migrationDataTransferMediumDetailsObjectStorageDetailsRepresentation},
	}
	migrationDatapumpSettingsRepresentation = map[string]interface{}{
		"export_directory_object": acctest.RepresentationGroup{RepType: acctest.Required, Group: migrationDatapumpSettingsExportDirectoryObjectRepresentation},
		"metadata_remaps":         acctest.RepresentationGroup{RepType: acctest.Required, Group: migrationDatapumpSettingsMetadataRemapsRepresentation},
	}
	migrationExcludeObjectsRepresentation = map[string]interface{}{
		"object": acctest.Representation{RepType: acctest.Required, Create: `.*`, Update: `object2`},
		"owner":  acctest.Representation{RepType: acctest.Required, Create: `owner`, Update: `owner2`},
		"type":   acctest.Representation{RepType: acctest.Optional, Create: `ALL`, Update: `TABLE`},
	}

	migrationDumpTransferRepresentation = map[string]interface{}{
		"source": acctest.RepresentationGroup{RepType: acctest.Optional, Group: migrationSourceHostDumpTransferDetailsRepresentation},
		"target": acctest.RepresentationGroup{RepType: acctest.Optional, Group: migrationTargetHostDumpTransferDetailsRepresentation},
	}

	migrationSourceHostDumpTransferDetailsRepresentation = map[string]interface{}{
		"kind":     acctest.Representation{RepType: acctest.Required, Create: `OCI_CLI`},
		"oci_home": acctest.Representation{RepType: acctest.Optional, Create: `/path/to/ociCli`},
	}

	migrationTargetHostDumpTransferDetailsRepresentation = map[string]interface{}{
		"kind": acctest.Representation{RepType: acctest.Required, Create: `CURL`},
	}
	migrationGoldenGateDetailsRepresentation = map[string]interface{}{
		"hub":      acctest.RepresentationGroup{RepType: acctest.Required, Group: migrationGoldenGateDetailsHubRepresentation},
		"settings": acctest.RepresentationGroup{RepType: acctest.Optional, Group: migrationGoldenGateDetailsSettingsRepresentation},
	}
	migrationVaultDetailsRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"key_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.kms_key_id}`},
		"vault_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.kms_vault_id}`},
	}
	migrationDataTransferMediumDetailsDatabaseLinkDetailsRepresentation = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
	}
	migrationDataTransferMediumDetailsObjectStorageDetailsRepresentation = map[string]interface{}{
		"bucket":    acctest.Representation{RepType: acctest.Required, Create: `bucket`, Update: `bucket2`},
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `namespace`, Update: `namespace2`},
	}
	migrationDatapumpSettingsDataPumpParametersRepresentation = map[string]interface{}{
		"estimate":                  acctest.Representation{RepType: acctest.Optional, Create: `BLOCKS`, Update: `STATISTICS`},
		"exclude_parameters":        acctest.Representation{RepType: acctest.Optional, Create: []string{`excludeParameters`}, Update: []string{`excludeParameters2`}},
		"export_parallelism_degree": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"import_parallelism_degree": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"is_cluster":                acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"table_exists_action":       acctest.Representation{RepType: acctest.Optional, Create: `TRUNCATE`, Update: `REPLACE`},
	}
	migrationDatapumpSettingsExportDirectoryObjectRepresentation = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `test_export_dir`, Update: `test_export_dir`},
		"path": acctest.Representation{RepType: acctest.Required, Create: `/u01/app/oracle/product/19.0.0.0/dbhome_1/rdbms/log`, Update: `/u01/app/oracle/product/19.0.0.0/dbhome_1/rdbms/log`},
	}
	migrationDatapumpSettingsImportDirectoryObjectRepresentation = map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"path": acctest.Representation{RepType: acctest.Required, Create: `path`, Update: `path2`},
	}
	migrationDatapumpSettingsMetadataRemapsRepresentation = map[string]interface{}{
		"new_value": acctest.Representation{RepType: acctest.Required, Create: `DATA`, Update: `DATA`},
		"old_value": acctest.Representation{RepType: acctest.Required, Create: `USERS`, Update: `USERS`},
		"type":      acctest.Representation{RepType: acctest.Required, Create: `TABLESPACE`, Update: `TABLESPACE`},
	}
	migrationGoldenGateDetailsHubRepresentation = map[string]interface{}{
		"rest_admin_credentials":               acctest.RepresentationGroup{RepType: acctest.Required, Group: migrationGoldenGateDetailsHubRestAdminCredentialsRepresentation},
		"source_db_admin_credentials":          acctest.RepresentationGroup{RepType: acctest.Required, Group: migrationGoldenGateDetailsHubSourceDbAdminCredentialsRepresentation},
		"source_microservices_deployment_name": acctest.Representation{RepType: acctest.Required, Create: `Target`},
		"target_db_admin_credentials":          acctest.RepresentationGroup{RepType: acctest.Required, Group: migrationGoldenGateDetailsHubTargetDbAdminCredentialsRepresentation},
		"target_microservices_deployment_name": acctest.Representation{RepType: acctest.Required, Create: `Target`},
		"url":                                  acctest.Representation{RepType: acctest.Required, Create: `https://130.35.83.125`, Update: `https://130.35.83.125`},
		//"source_container_db_admin_credentials": acctest.RepresentationGroup{RepType:acctest.Required,Group: migrationGoldenGateDetailsHubSourceContainerDbAdminCredentialsRepresentation},
	}
	migrationGoldenGateDetailsSettingsRepresentation = map[string]interface{}{
		"acceptable_lag": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"extract":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: migrationGoldenGateDetailsSettingsExtractRepresentation},
		"replicat":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: migrationGoldenGateDetailsSettingsReplicatRepresentation},
	}
	migrationGoldenGateDetailsHubRestAdminCredentialsRepresentation = map[string]interface{}{
		"password": acctest.Representation{RepType: acctest.Required, Create: `n5j2LRy0X%A2VRxY`, Update: `n5j2LRy0X%A2VRxY`},
		"username": acctest.Representation{RepType: acctest.Required, Create: `oggadmin`, Update: `oggadmin`},
	}
	migrationGoldenGateDetailsHubSourceDbAdminCredentialsRepresentation = map[string]interface{}{
		"password": acctest.Representation{RepType: acctest.Required, Create: `GG$$admin128`, Update: `GG$$admin128`},
		"username": acctest.Representation{RepType: acctest.Required, Create: `ggadmin`, Update: `ggadmin`},
	}
	migrationGoldenGateDetailsHubTargetDbAdminCredentialsRepresentation = map[string]interface{}{
		"password": acctest.Representation{RepType: acctest.Required, Create: `ORcl##4567890`, Update: `ORcl##4567890`},
		"username": acctest.Representation{RepType: acctest.Required, Create: `ggadmin`, Update: `ggadmin`},
	}
	migrationGoldenGateDetailsHubSourceContainerDbAdminCredentialsRepresentation = map[string]interface{}{
		"password": acctest.Representation{RepType: acctest.Required, Create: `GG$$admin128`, Update: `GG$$admin128`},
		"username": acctest.Representation{RepType: acctest.Required, Create: `c##ggadmin`, Update: `c##ggadmin`},
	}
	migrationGoldenGateDetailsSettingsExtractRepresentation = map[string]interface{}{
		"long_trans_duration": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"performance_profile": acctest.Representation{RepType: acctest.Optional, Create: `LOW`, Update: `MEDIUM`},
	}
	migrationGoldenGateDetailsSettingsReplicatRepresentation = map[string]interface{}{
		"map_parallelism":       acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"max_apply_parallelism": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"min_apply_parallelism": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	}

	MigrationResourceDependenciesMig = ConnectionResourceDependenciesTargetCommon +
		acctest.GenerateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection", acctest.Required, acctest.Create, connectionRepresentationTarget) +
		ConnectionResourceDependenciesSource +
		acctest.GenerateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection_source", acctest.Required, acctest.Create, connectionRepresentationSource)
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

	resourceName := "oci_database_migration_migration.test_migration"
	datasourceName := "data.oci_database_migration_migrations.test_migrations"
	singularDatasourceName := "data.oci_database_migration_migration.test_migration"
	var resId, resId2 string
	acctest.SaveConfigContent(config+compartmentIdVariableStr+MigrationResourceDependenciesMig+
		acctest.GenerateResourceFromRepresentationMap("oci_database_migration_migration", "test_migration", acctest.Optional, acctest.Create, migrationRepresentationMig), "databasemigration", "migration", t)

	acctest.ResourceTest(t, testAccCheckDatabaseMigrationMigrationDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + MigrationResourceDependenciesMig +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_migration", "test_migration", acctest.Required, acctest.Create, migrationRepresentationMig),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
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
			Config: config + compartmentIdVariableStr, //+ MigrationResourceDependenciesMig,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + MigrationResourceDependenciesMig +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_migration", "test_migration", acctest.Optional, acctest.Create, migrationRepresentationMig),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "datapump_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "datapump_settings.0.export_directory_object.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "datapump_settings.0.export_directory_object.0.name", "test_export_dir"),
				resource.TestCheckResourceAttr(resourceName, "datapump_settings.0.export_directory_object.0.path", "/u01/app/oracle/product/19.0.0.0/dbhome_1/rdbms/log"),
				resource.TestCheckResourceAttr(resourceName, "datapump_settings.0.metadata_remaps.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "datapump_settings.0.metadata_remaps", map[string]string{
					"new_value": "DATA",
					"old_value": "USERS",
					"type":      "TABLESPACE",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TF_ONLINE_MIG"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.hub.0.rest_admin_credentials.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.hub.0.rest_admin_credentials.0.password", "n5j2LRy0X%A2VRxY"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.hub.0.rest_admin_credentials.0.username", "oggadmin"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.hub.0.source_db_admin_credentials.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.hub.0.source_db_admin_credentials.0.password", "GG$$admin128"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.hub.0.source_db_admin_credentials.0.username", "ggadmin"),
				resource.TestCheckResourceAttrSet(resourceName, "golden_gate_details.0.hub.0.source_microservices_deployment_name"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.hub.0.target_db_admin_credentials.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.hub.0.target_db_admin_credentials.0.password", "ORcl##4567890"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.hub.0.target_db_admin_credentials.0.username", "ggadmin"),
				resource.TestCheckResourceAttrSet(resourceName, "golden_gate_details.0.hub.0.target_microservices_deployment_name"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.hub.0.url", "https://130.35.83.125"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.settings.0.acceptable_lag", "10"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.settings.0.extract.0.long_trans_duration", "10"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.settings.0.extract.0.performance_profile", "LOW"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.settings.0.replicat.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.settings.0.replicat.0.map_parallelism", "10"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.settings.0.replicat.0.max_apply_parallelism", "10"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.settings.0.replicat.0.min_apply_parallelism", "10"),

				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "source_database_connection_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_database_connection_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "ONLINE"),
				resource.TestCheckResourceAttr(resourceName, "vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "vault_details.0.key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_details.0.vault_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + MigrationResourceDependenciesMig +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_migration", "test_migration", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(migrationRepresentationMig, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),

				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "datapump_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "datapump_settings.0.export_directory_object.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "datapump_settings.0.export_directory_object.0.name", "test_export_dir"),
				resource.TestCheckResourceAttr(resourceName, "datapump_settings.0.export_directory_object.0.path", "/u01/app/oracle/product/19.0.0.0/dbhome_1/rdbms/log"),
				resource.TestCheckResourceAttr(resourceName, "datapump_settings.0.metadata_remaps.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "datapump_settings.0.metadata_remaps", map[string]string{
					"new_value": "DATA",
					"old_value": "USERS",
					"type":      "TABLESPACE",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TF_ONLINE_MIG"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.hub.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.hub.0.rest_admin_credentials.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.hub.0.rest_admin_credentials.0.password", "n5j2LRy0X%A2VRxY"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.hub.0.rest_admin_credentials.0.username", "oggadmin"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.hub.0.source_db_admin_credentials.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.hub.0.source_db_admin_credentials.0.password", "GG$$admin128"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.hub.0.source_db_admin_credentials.0.username", "ggadmin"),
				resource.TestCheckResourceAttrSet(resourceName, "golden_gate_details.0.hub.0.source_microservices_deployment_name"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.hub.0.target_db_admin_credentials.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.hub.0.target_db_admin_credentials.0.password", "ORcl##4567890"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.hub.0.target_db_admin_credentials.0.username", "ggadmin"),
				resource.TestCheckResourceAttrSet(resourceName, "golden_gate_details.0.hub.0.target_microservices_deployment_name"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.hub.0.url", "https://130.35.83.125"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.settings.0.acceptable_lag", "10"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.settings.0.extract.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.settings.0.extract.0.long_trans_duration", "10"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.settings.0.extract.0.performance_profile", "LOW"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.settings.0.replicat.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.settings.0.replicat.0.map_parallelism", "10"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.settings.0.replicat.0.max_apply_parallelism", "10"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.settings.0.replicat.0.min_apply_parallelism", "10"),

				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "source_database_connection_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_database_connection_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "ONLINE"),
				resource.TestCheckResourceAttr(resourceName, "vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "vault_details.0.key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_details.0.vault_id"),

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
			Config: config + compartmentIdVariableStr + MigrationResourceDependenciesMig +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_migration", "test_migration", acctest.Optional, acctest.Update, migrationRepresentationMig),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "datapump_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "datapump_settings.0.export_directory_object.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "datapump_settings.0.export_directory_object.0.name", "test_export_dir"),
				resource.TestCheckResourceAttr(resourceName, "datapump_settings.0.export_directory_object.0.path", "/u01/app/oracle/product/19.0.0.0/dbhome_1/rdbms/log"),
				resource.TestCheckResourceAttr(resourceName, "datapump_settings.0.metadata_remaps.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "datapump_settings.0.metadata_remaps.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "datapump_settings.0.metadata_remaps", map[string]string{
					"new_value": "DATA",
					"old_value": "USERS",
					"type":      "TABLESPACE",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TF_ONLINE_MIG"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.hub.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.hub.0.rest_admin_credentials.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.hub.0.rest_admin_credentials.0.password", "n5j2LRy0X%A2VRxY"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.hub.0.rest_admin_credentials.0.username", "oggadmin"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.hub.0.source_db_admin_credentials.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.hub.0.source_db_admin_credentials.0.password", "GG$$admin128"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.hub.0.source_db_admin_credentials.0.username", "ggadmin"),
				resource.TestCheckResourceAttrSet(resourceName, "golden_gate_details.0.hub.0.source_microservices_deployment_name"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.hub.0.target_db_admin_credentials.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.hub.0.target_db_admin_credentials.0.password", "ORcl##4567890"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.hub.0.target_db_admin_credentials.0.username", "ggadmin"),
				resource.TestCheckResourceAttrSet(resourceName, "golden_gate_details.0.hub.0.target_microservices_deployment_name"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.hub.0.url", "https://130.35.83.125"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.settings.0.acceptable_lag", "11"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.settings.0.extract.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.settings.0.extract.0.long_trans_duration", "11"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.settings.0.extract.0.performance_profile", "MEDIUM"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.settings.0.replicat.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.settings.0.replicat.0.map_parallelism", "11"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.settings.0.replicat.0.max_apply_parallelism", "11"),
				resource.TestCheckResourceAttr(resourceName, "golden_gate_details.0.settings.0.replicat.0.min_apply_parallelism", "11"),

				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "source_database_connection_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "target_database_connection_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "type", "ONLINE"),
				resource.TestCheckResourceAttr(resourceName, "vault_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "vault_details.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "vault_details.0.key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_details.0.vault_id"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_migration_migrations", "test_migrations", acctest.Optional, acctest.Update, migrationDataSourceRepresentation) +
				compartmentIdVariableStr + MigrationResourceDependenciesMig +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_migration", "test_migration", acctest.Optional, acctest.Update, migrationRepresentationMig),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "migration_collection.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_migration_migration", "test_migration", acctest.Required, acctest.Create, migrationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + MigrationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "migration_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "datapump_settings.0.export_directory_object.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "datapump_settings.0.export_directory_object.0.name", "test_export_dir"),
				resource.TestCheckResourceAttr(singularDatasourceName, "datapump_settings.0.export_directory_object.0.path", "/u01/app/oracle/product/19.0.0.0/dbhome_1/rdbms/log"),
				resource.TestCheckResourceAttr(singularDatasourceName, "datapump_settings.0.metadata_remaps.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "datapump_settings.0.metadata_remaps", map[string]string{
					"new_value": "DATA",
					"old_value": "USERS",
					"type":      "TABLESPACE",
				},
					[]string{}),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "TF_ONLINE_MIG"),
				resource.TestCheckResourceAttr(singularDatasourceName, "golden_gate_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "golden_gate_details.0.hub.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "golden_gate_details.0.hub.0.rest_admin_credentials.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "golden_gate_details.0.hub.0.rest_admin_credentials.0.username", "oggadmin"),
				resource.TestCheckResourceAttr(singularDatasourceName, "golden_gate_details.0.hub.0.source_db_admin_credentials.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "golden_gate_details.0.hub.0.source_db_admin_credentials.0.username", "ggadmin"),
				resource.TestCheckResourceAttr(singularDatasourceName, "golden_gate_details.0.hub.0.target_db_admin_credentials.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "golden_gate_details.0.hub.0.target_db_admin_credentials.0.username", "ggadmin"),
				resource.TestCheckResourceAttr(singularDatasourceName, "golden_gate_details.0.hub.0.url", "https://130.35.83.125"),
				resource.TestCheckResourceAttr(singularDatasourceName, "golden_gate_details.0.settings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "golden_gate_details.0.settings.0.acceptable_lag", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "golden_gate_details.0.settings.0.extract.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "golden_gate_details.0.settings.0.extract.0.long_trans_duration", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "golden_gate_details.0.settings.0.extract.0.performance_profile", "MEDIUM"),
				resource.TestCheckResourceAttr(singularDatasourceName, "golden_gate_details.0.settings.0.replicat.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "golden_gate_details.0.settings.0.replicat.0.map_parallelism", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "golden_gate_details.0.settings.0.replicat.0.max_apply_parallelism", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "golden_gate_details.0.settings.0.replicat.0.min_apply_parallelism", "11"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "ONLINE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vault_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vault_details.0.compartment_id", compartmentId),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + MigrationResourceConfig,
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			//ImportStateVerifyIgnore: []string{},
			ImportStateVerifyIgnore: []string{
				"golden_gate_details.0.hub.0.rest_admin_credentials.0.password",
				"golden_gate_details.0.hub.0.source_container_db_admin_credentials.0.password",
				"golden_gate_details.0.hub.0.source_db_admin_credentials.0.password",
				"golden_gate_details.0.hub.0.target_db_admin_credentials.0.password",
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
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
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
	migrationIds, err := getMigrationIds(compartment)
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
			acctest.WaitTillCondition(acctest.TestAccProvider, &migrationId, migrationSweepWaitCondition, time.Duration(3*time.Minute),
				migrationSweepResponseFetchOperation, "database_migration", true)
		}
	}
	return nil
}

func getMigrationIds(compartment string) ([]string, error) {
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
		id := *migration.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "MigrationId", id)
	}
	return resourceIds, nil
}

func migrationSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	/*if migrationResponse, ok := response.Response.(oci_database_migration.GetMigrationResponse); ok {
		return migrationResponse.LifecycleState != oci_database_migration.LifecycleStatesDeleted
	}*/
	return false
}

func migrationSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseMigrationClient().GetMigration(context.Background(), oci_database_migration.GetMigrationRequest{
		MigrationId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
