// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v44/common"
	oci_database_migration "github.com/oracle/oci-go-sdk/v44/databasemigration"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	MigrationRequiredOnlyResource = MigrationResourceDependenciesMig +
		generateResourceFromRepresentationMap("oci_database_migration_migration", "test_migration", Required, Create, migrationRepresentationMig)

	MigrationResourceConfig = MigrationResourceDependenciesMig +
		generateResourceFromRepresentationMap("oci_database_migration_migration", "test_migration", Optional, Update, migrationRepresentationMig)

	migrationSingularDataSourceRepresentation = map[string]interface{}{
		"migration_id": Representation{repType: Required, create: `${oci_database_migration_migration.test_migration.id}`},
	}

	migrationDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"state":          Representation{repType: Optional, create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, migrationDataSourceFilterRepresentation}}
	migrationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_database_migration_migration.test_migration.id}`}},
	}

	migrationRepresentationMig = map[string]interface{}{
		"compartment_id":                Representation{repType: Required, create: `${var.compartment_id}`},
		"source_database_connection_id": Representation{repType: Required, create: `${oci_database_migration_connection.test_connection_source.id}`},
		"target_database_connection_id": Representation{repType: Required, create: `${oci_database_migration_connection.test_connection.id}`},
		"type":                          Representation{repType: Required, create: `ONLINE`},
		"datapump_settings":             RepresentationGroup{Required, migrationDatapumpSettingsRepresentation},
		"display_name":                  Representation{repType: Optional, create: `TF_ONLINE_MIG`, update: `TF_ONLINE_MIG`},
		"golden_gate_details":           RepresentationGroup{Required, migrationGoldenGateDetailsRepresentation},
		"vault_details":                 RepresentationGroup{Required, migrationVaultDetailsRepresentation},
	}
	migrationDataTransferMediumDetailsRepresentation = map[string]interface{}{
		"database_link_details":  RepresentationGroup{Optional, migrationDataTransferMediumDetailsDatabaseLinkDetailsRepresentation},
		"object_storage_details": RepresentationGroup{Optional, migrationDataTransferMediumDetailsObjectStorageDetailsRepresentation},
	}
	migrationDatapumpSettingsRepresentation = map[string]interface{}{
		"export_directory_object": RepresentationGroup{Required, migrationDatapumpSettingsExportDirectoryObjectRepresentation},
		"metadata_remaps":         RepresentationGroup{Required, migrationDatapumpSettingsMetadataRemapsRepresentation},
	}
	migrationExcludeObjectsRepresentation = map[string]interface{}{
		"object": Representation{repType: Required, create: `object`, update: `object2`},
		"owner":  Representation{repType: Required, create: `owner`, update: `owner2`},
	}
	migrationGoldenGateDetailsRepresentation = map[string]interface{}{
		"hub":      RepresentationGroup{Required, migrationGoldenGateDetailsHubRepresentation},
		"settings": RepresentationGroup{Optional, migrationGoldenGateDetailsSettingsRepresentation},
	}
	migrationVaultDetailsRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"key_id":         Representation{repType: Required, create: `${var.kms_key_id}`},
		"vault_id":       Representation{repType: Required, create: `${var.kms_vault_id}`},
	}
	migrationDataTransferMediumDetailsDatabaseLinkDetailsRepresentation = map[string]interface{}{
		"name": Representation{repType: Required, create: `name`, update: `name2`},
	}
	migrationDataTransferMediumDetailsObjectStorageDetailsRepresentation = map[string]interface{}{
		"bucket":    Representation{repType: Required, create: `bucket`, update: `bucket2`},
		"namespace": Representation{repType: Required, create: `namespace`, update: `namespace2`},
	}
	migrationDatapumpSettingsDataPumpParametersRepresentation = map[string]interface{}{
		"estimate":                  Representation{repType: Optional, create: `BLOCKS`, update: `STATISTICS`},
		"exclude_parameters":        Representation{repType: Optional, create: []string{`excludeParameters`}, update: []string{`excludeParameters2`}},
		"export_parallelism_degree": Representation{repType: Optional, create: `10`, update: `11`},
		"import_parallelism_degree": Representation{repType: Optional, create: `10`, update: `11`},
		"is_cluster":                Representation{repType: Optional, create: `false`, update: `true`},
		"table_exists_action":       Representation{repType: Optional, create: `TRUNCATE`, update: `REPLACE`},
	}
	migrationDatapumpSettingsExportDirectoryObjectRepresentation = map[string]interface{}{
		"name": Representation{repType: Required, create: `test_export_dir`, update: `test_export_dir`},
		"path": Representation{repType: Required, create: `/u01/app/oracle/product/19.0.0.0/dbhome_1/rdbms/log`, update: `/u01/app/oracle/product/19.0.0.0/dbhome_1/rdbms/log`},
	}
	migrationDatapumpSettingsImportDirectoryObjectRepresentation = map[string]interface{}{
		"name": Representation{repType: Required, create: `name`, update: `name2`},
		"path": Representation{repType: Required, create: `path`, update: `path2`},
	}
	migrationDatapumpSettingsMetadataRemapsRepresentation = map[string]interface{}{
		"new_value": Representation{repType: Required, create: `DATA`, update: `DATA`},
		"old_value": Representation{repType: Required, create: `USERS`, update: `USERS`},
		"type":      Representation{repType: Required, create: `TABLESPACE`, update: `TABLESPACE`},
	}
	migrationGoldenGateDetailsHubRepresentation = map[string]interface{}{
		"rest_admin_credentials":               RepresentationGroup{Required, migrationGoldenGateDetailsHubRestAdminCredentialsRepresentation},
		"source_db_admin_credentials":          RepresentationGroup{Required, migrationGoldenGateDetailsHubSourceDbAdminCredentialsRepresentation},
		"source_microservices_deployment_name": Representation{repType: Required, create: `Target`},
		"target_db_admin_credentials":          RepresentationGroup{Required, migrationGoldenGateDetailsHubTargetDbAdminCredentialsRepresentation},
		"target_microservices_deployment_name": Representation{repType: Required, create: `Target`},
		"url":                                  Representation{repType: Required, create: `https://130.35.83.125`, update: `https://130.35.83.125`},
		//"source_container_db_admin_credentials": RepresentationGroup{Required, migrationGoldenGateDetailsHubSourceContainerDbAdminCredentialsRepresentation},
	}
	migrationGoldenGateDetailsSettingsRepresentation = map[string]interface{}{
		"acceptable_lag": Representation{repType: Optional, create: `10`, update: `11`},
		"extract":        RepresentationGroup{Optional, migrationGoldenGateDetailsSettingsExtractRepresentation},
		"replicat":       RepresentationGroup{Optional, migrationGoldenGateDetailsSettingsReplicatRepresentation},
	}
	migrationGoldenGateDetailsHubRestAdminCredentialsRepresentation = map[string]interface{}{
		"password": Representation{repType: Required, create: `n5j2LRy0X%A2VRxY`, update: `n5j2LRy0X%A2VRxY`},
		"username": Representation{repType: Required, create: `oggadmin`, update: `oggadmin`},
	}
	migrationGoldenGateDetailsHubSourceDbAdminCredentialsRepresentation = map[string]interface{}{
		"password": Representation{repType: Required, create: `GG$$admin128`, update: `GG$$admin128`},
		"username": Representation{repType: Required, create: `ggadmin`, update: `ggadmin`},
	}
	migrationGoldenGateDetailsHubTargetDbAdminCredentialsRepresentation = map[string]interface{}{
		"password": Representation{repType: Required, create: `ORcl##4567890`, update: `ORcl##4567890`},
		"username": Representation{repType: Required, create: `ggadmin`, update: `ggadmin`},
	}
	migrationGoldenGateDetailsHubSourceContainerDbAdminCredentialsRepresentation = map[string]interface{}{
		"password": Representation{repType: Required, create: `GG$$admin128`, update: `GG$$admin128`},
		"username": Representation{repType: Required, create: `c##ggadmin`, update: `c##ggadmin`},
	}
	migrationGoldenGateDetailsSettingsExtractRepresentation = map[string]interface{}{
		"long_trans_duration": Representation{repType: Optional, create: `10`, update: `11`},
		"performance_profile": Representation{repType: Optional, create: `LOW`, update: `MEDIUM`},
	}
	migrationGoldenGateDetailsSettingsReplicatRepresentation = map[string]interface{}{
		"map_parallelism":       Representation{repType: Optional, create: `10`, update: `11`},
		"max_apply_parallelism": Representation{repType: Optional, create: `10`, update: `11`},
		"min_apply_parallelism": Representation{repType: Optional, create: `10`, update: `11`},
	}

	MigrationResourceDependenciesMig = ConnectionResourceDependenciesTargetCommon +
		generateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection", Required, Create, connectionRepresentationTarget) +
		ConnectionResourceDependenciesSource +
		generateResourceFromRepresentationMap("oci_database_migration_connection", "test_connection_source", Required, Create, connectionRepresentationSource)
)

func TestDatabaseMigrationMigrationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseMigrationMigrationResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithBlankDefault("compartment_id_for_update")
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_migration_migration.test_migration"
	datasourceName := "data.oci_database_migration_migrations.test_migrations"
	singularDatasourceName := "data.oci_database_migration_migration.test_migration"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+MigrationResourceDependenciesMig+
		generateResourceFromRepresentationMap("oci_database_migration_migration", "test_migration", Optional, Create, migrationRepresentationMig), "databasemigration", "migration", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseMigrationMigrationDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + MigrationResourceDependenciesMig +
					generateResourceFromRepresentationMap("oci_database_migration_migration", "test_migration", Required, Create, migrationRepresentationMig),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(resourceName, "source_database_connection_id"),
					resource.TestCheckResourceAttrSet(resourceName, "target_database_connection_id"),
					resource.TestCheckResourceAttr(resourceName, "type", "ONLINE"),
					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr, //+ MigrationResourceDependenciesMig,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + MigrationResourceDependenciesMig +
					generateResourceFromRepresentationMap("oci_database_migration_migration", "test_migration", Optional, Create, migrationRepresentationMig),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "datapump_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "datapump_settings.0.export_directory_object.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "datapump_settings.0.export_directory_object.0.name", "test_export_dir"),
					resource.TestCheckResourceAttr(resourceName, "datapump_settings.0.export_directory_object.0.path", "/u01/app/oracle/product/19.0.0.0/dbhome_1/rdbms/log"),
					resource.TestCheckResourceAttr(resourceName, "datapump_settings.0.metadata_remaps.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "datapump_settings.0.metadata_remaps", map[string]string{
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
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + MigrationResourceDependenciesMig +
					generateResourceFromRepresentationMap("oci_database_migration_migration", "test_migration", Optional, Create,
						representationCopyWithNewProperties(migrationRepresentationMig, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),

					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "datapump_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "datapump_settings.0.export_directory_object.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "datapump_settings.0.export_directory_object.0.name", "test_export_dir"),
					resource.TestCheckResourceAttr(resourceName, "datapump_settings.0.export_directory_object.0.path", "/u01/app/oracle/product/19.0.0.0/dbhome_1/rdbms/log"),
					resource.TestCheckResourceAttr(resourceName, "datapump_settings.0.metadata_remaps.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "datapump_settings.0.metadata_remaps.0.new_value", "DATA"),
					resource.TestCheckResourceAttr(resourceName, "datapump_settings.0.metadata_remaps.0.old_value", "USERS"),
					resource.TestCheckResourceAttr(resourceName, "datapump_settings.0.metadata_remaps.0.type", "TABLESPACE"),
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
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateResourceFromRepresentationMap("oci_database_migration_migration", "test_migration", Optional, Update, migrationRepresentationMig),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "datapump_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "datapump_settings.0.export_directory_object.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "datapump_settings.0.export_directory_object.0.name", "test_export_dir"),
					resource.TestCheckResourceAttr(resourceName, "datapump_settings.0.export_directory_object.0.path", "/u01/app/oracle/product/19.0.0.0/dbhome_1/rdbms/log"),
					resource.TestCheckResourceAttr(resourceName, "datapump_settings.0.metadata_remaps.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "datapump_settings.0.metadata_remaps.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "datapump_settings.0.metadata_remaps", map[string]string{
						"new_value": "DATA",
						"old_value": "USERS",
						"type":      "TABLESPACE",
					},
						[]string{}),
					resource.TestCheckResourceAttr(resourceName, "datapump_settings.0.metadata_remaps.0.type", "TABLESPACE"),
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
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_database_migration_migrations", "test_migrations", Optional, Update, migrationDataSourceRepresentation) +
					compartmentIdVariableStr + MigrationResourceDependenciesMig +
					generateResourceFromRepresentationMap("oci_database_migration_migration", "test_migration", Optional, Update, migrationRepresentationMig),
				Check: ComposeAggregateTestCheckFuncWrapper(
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
					generateDataSourceFromRepresentationMap("oci_database_migration_migration", "test_migration", Required, Create, migrationSingularDataSourceRepresentation) +
					compartmentIdVariableStr + MigrationResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "migration_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "datapump_settings.0.export_directory_object.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "datapump_settings.0.export_directory_object.0.name", "test_export_dir"),
					resource.TestCheckResourceAttr(singularDatasourceName, "datapump_settings.0.export_directory_object.0.path", "/u01/app/oracle/product/19.0.0.0/dbhome_1/rdbms/log"),
					resource.TestCheckResourceAttr(singularDatasourceName, "datapump_settings.0.metadata_remaps.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "datapump_settings.0.metadata_remaps.0.new_value", "DATA"),
					resource.TestCheckResourceAttr(singularDatasourceName, "datapump_settings.0.metadata_remaps.0.old_value", "USERS"),
					resource.TestCheckResourceAttr(singularDatasourceName, "datapump_settings.0.metadata_remaps.0.type", "TABLESPACE"),
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
		},
	})
}

func testAccCheckDatabaseMigrationMigrationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).databaseMigrationClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_migration_migration" {
			noResourceFound = false
			request := oci_database_migration.GetMigrationRequest{}

			tmp := rs.Primary.ID
			request.MigrationId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database_migration")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("DatabaseMigrationMigration") {
		resource.AddTestSweepers("DatabaseMigrationMigration", &resource.Sweeper{
			Name:         "DatabaseMigrationMigration",
			Dependencies: DependencyGraph["migration"],
			F:            sweepDatabaseMigrationMigrationResource,
		})
	}
}

func sweepDatabaseMigrationMigrationResource(compartment string) error {
	databaseMigrationClient := GetTestClients(&schema.ResourceData{}).databaseMigrationClient()
	migrationIds, err := getMigrationIds(compartment)
	if err != nil {
		return err
	}
	for _, migrationId := range migrationIds {
		if ok := SweeperDefaultResourceId[migrationId]; !ok {
			deleteMigrationRequest := oci_database_migration.DeleteMigrationRequest{}

			deleteMigrationRequest.MigrationId = &migrationId

			deleteMigrationRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database_migration")
			_, error := databaseMigrationClient.DeleteMigration(context.Background(), deleteMigrationRequest)
			if error != nil {
				fmt.Printf("Error deleting Migration %s %s, It is possible that the resource is already deleted. Please verify manually \n", migrationId, error)
				continue
			}
			waitTillCondition(testAccProvider, &migrationId, migrationSweepWaitCondition, time.Duration(3*time.Minute),
				migrationSweepResponseFetchOperation, "database_migration", true)
		}
	}
	return nil
}

func getMigrationIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "MigrationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseMigrationClient := GetTestClients(&schema.ResourceData{}).databaseMigrationClient()

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
		addResourceIdToSweeperResourceIdMap(compartmentId, "MigrationId", id)
	}
	return resourceIds, nil
}

func migrationSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if migrationResponse, ok := response.Response.(oci_database_migration.GetMigrationResponse); ok {
		return migrationResponse.LifecycleState != oci_database_migration.LifecycleStatesDeleted
	}
	return false
}

func migrationSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.databaseMigrationClient().GetMigration(context.Background(), oci_database_migration.GetMigrationRequest{
		MigrationId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
