// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_psql "github.com/oracle/oci-go-sdk/v65/psql"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	PsqlBackupRequiredOnlyResource = PsqlBackupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_psql_backup", "test_backup", acctest.Required, acctest.Create, PsqlBackupRepresentation)

	PsqlBackupResourceConfig = PsqlBackupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_psql_backup", "test_backup", acctest.Optional, acctest.Update, PsqlBackupRepresentation)

	PsqlBackupSingularDataSourceRepresentation = map[string]interface{}{
		"backup_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_psql_backup.test_backup.id}`},
	}

	PsqlBackupDataSourceRepresentation = map[string]interface{}{
		// "backup_id":      acctest.Representation{RepType: acctest.Optional, Create: `${oci_psql_backup.test_backup.id}`},
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `${oci_psql_backup.test_backup.display_name}`},
		"id":             acctest.Representation{RepType: acctest.Required, Create: `${var.db_ocid}`},
		"state":          acctest.Representation{RepType: acctest.Required, Create: `${oci_psql_backup.test_backup.state}`},
		"time_ended":     acctest.Representation{RepType: acctest.Optional, Create: `2024-01-02T15:04:05Z`},
		"time_started":   acctest.Representation{RepType: acctest.Optional, Create: `2000-01-02T15:04:05Z`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: PsqlBackupDataSourceFilterRepresentation},
	}

	PsqlBackupIDOnlyDataSourceRepresentation = map[string]interface{}{
		"backup_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_psql_backup.test_backup.id}`},
		"filter":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: PsqlBackupDataSourceFilterRepresentation},
	}

	PsqlBackupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Optional, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_psql_backup.test_backup.id}`}},
	}

	PsqlBackupRepresentation = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"db_system_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.db_ocid}`},
		"display_name":     acctest.Representation{RepType: acctest.Required, Create: `terrafrom-backup-test`, Update: `terrafrom-backup-test-2`},
		"defined_tags":     acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":      acctest.Representation{RepType: acctest.Optional, Create: `description1`, Update: `description2`},
		"freeform_tags":    acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"retention_period": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"lifecycle":        acctest.RepresentationGroup{RepType: acctest.Required, Group: ignorePsqlBackupDefinedTagsChangesRepresentation},
	}

	ignorePsqlBackupDefinedTagsChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	PsqlBackupResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: psql/default
func TestPsqlBackupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestPsqlBackupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	subnetId := utils.GetEnvSettingWithBlankDefault("subnet_ocid")
	subnetIdVariableStr := fmt.Sprintf("variable \"subnet_id\" { default = \"%s\" }\n", subnetId)

	dbSystemId := utils.GetEnvSettingWithBlankDefault("db_ocid")
	dbSystemIdVariableStr := fmt.Sprintf("variable \"db_ocid\" { default = \"%s\" }\n", dbSystemId)

	resourceName := "oci_psql_backup.test_backup"
	datasourceName := "data.oci_psql_backups.test_backups"
	singularDatasourceName := "data.oci_psql_backup.test_backup"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+PsqlBackupResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_psql_backup", "test_backup", acctest.Optional, acctest.Create, PsqlBackupRepresentation), "psql", "backup", t)

	acctest.ResourceTest(t, testAccCheckPsqlBackupDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + dbSystemIdVariableStr + PsqlBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_psql_backup", "test_backup", acctest.Required, acctest.Create, PsqlBackupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "db_system_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "terrafrom-backup-test"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + subnetIdVariableStr + PsqlBackupResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + dbSystemIdVariableStr + PsqlBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_psql_backup", "test_backup", acctest.Optional, acctest.Create, PsqlBackupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "backup_size"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "db_system_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "db_system_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "terrafrom-backup-test"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "retention_period", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + dbSystemIdVariableStr + subnetIdVariableStr + PsqlBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_psql_backup", "test_backup", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(PsqlBackupRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "backup_size"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "db_system_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "db_system_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "terrafrom-backup-test"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "retention_period", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + dbSystemIdVariableStr + subnetIdVariableStr + PsqlBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_psql_backup", "test_backup", acctest.Optional, acctest.Update, PsqlBackupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "backup_size"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "db_system_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "db_system_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "terrafrom-backup-test-2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "retention_period", "11"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + dbSystemIdVariableStr + subnetIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_psql_backups", "test_backups", acctest.Optional, acctest.Update, PsqlBackupDataSourceRepresentation) +
				compartmentIdVariableStr + PsqlBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_psql_backup", "test_backup", acctest.Optional, acctest.Update, PsqlBackupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "terrafrom-backup-test-2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),

				resource.TestCheckResourceAttrSet(datasourceName, "backup_collection.0.items.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "backup_collection.0.items.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "backup_collection.0.items.0.display_name", "terrafrom-backup-test-2"),
				resource.TestCheckResourceAttrSet(datasourceName, "backup_collection.0.items.0.db_system_id"),
				resource.TestCheckResourceAttr(datasourceName, "backup_collection.0.items.0.state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_ended"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_started"),

				resource.TestCheckResourceAttr(datasourceName, "backup_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "backup_collection.0.items.#", "1"),
			),
		},
		// verify id only datasource
		{
			Config: config + dbSystemIdVariableStr + subnetIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_psql_backups", "test_backups", acctest.Optional, acctest.Update, PsqlBackupIDOnlyDataSourceRepresentation) +
				compartmentIdVariableStr + PsqlBackupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_psql_backup", "test_backup", acctest.Optional, acctest.Update, PsqlBackupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "backup_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "backup_collection.0.items.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "backup_collection.0.items.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "backup_collection.0.items.0.display_name", "terrafrom-backup-test-2"),
				resource.TestCheckResourceAttrSet(datasourceName, "backup_collection.0.items.0.db_system_id"),
				resource.TestCheckResourceAttr(datasourceName, "backup_collection.0.items.0.state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "backup_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "backup_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config + dbSystemIdVariableStr + subnetIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_psql_backup", "test_backup", acctest.Required, acctest.Create, PsqlBackupSingularDataSourceRepresentation) +
				compartmentIdVariableStr + PsqlBackupResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "backup_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "backup_size"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_system_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "terrafrom-backup-test-2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "retention_period", "11"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "source_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + dbSystemIdVariableStr + subnetIdVariableStr + PsqlBackupRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"description",
				"last_accepted_request_token",
				"last_completed_request_token",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckPsqlBackupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).PostgresqlClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_psql_backup" {
			noResourceFound = false
			request := oci_psql.GetBackupRequest{}

			tmp := rs.Primary.ID
			request.BackupId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "psql")

			response, err := client.GetBackup(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_psql.BackupLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("PsqlBackup") {
		resource.AddTestSweepers("PsqlBackup", &resource.Sweeper{
			Name:         "PsqlBackup",
			Dependencies: acctest.DependencyGraph["backup"],
			F:            sweepPsqlBackupResource,
		})
	}
}

func sweepPsqlBackupResource(compartment string) error {
	postgresqlClient := acctest.GetTestClients(&schema.ResourceData{}).PostgresqlClient()
	backupIds, err := getPsqlBackupIds(compartment)
	if err != nil {
		return err
	}
	for _, backupId := range backupIds {
		if ok := acctest.SweeperDefaultResourceId[backupId]; !ok {
			deleteBackupRequest := oci_psql.DeleteBackupRequest{}

			deleteBackupRequest.BackupId = &backupId

			deleteBackupRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "psql")
			_, error := postgresqlClient.DeleteBackup(context.Background(), deleteBackupRequest)
			if error != nil {
				fmt.Printf("Error deleting Backup %s %s, It is possible that the resource is already deleted. Please verify manually \n", backupId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &backupId, PsqlBackupSweepWaitCondition, time.Duration(3*time.Minute),
				PsqlBackupSweepResponseFetchOperation, "psql", true)
		}
	}
	return nil
}

func getPsqlBackupIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "BackupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	postgresqlClient := acctest.GetTestClients(&schema.ResourceData{}).PostgresqlClient()

	listBackupsRequest := oci_psql.ListBackupsRequest{}
	listBackupsRequest.CompartmentId = &compartmentId
	listBackupsRequest.LifecycleState = oci_psql.BackupLifecycleStateActive
	listBackupsResponse, err := postgresqlClient.ListBackups(context.Background(), listBackupsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Backup list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, backup := range listBackupsResponse.Items {
		id := *backup.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "BackupId", id)
	}
	return resourceIds, nil
}

func PsqlBackupSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if backupResponse, ok := response.Response.(oci_psql.GetBackupResponse); ok {
		return backupResponse.LifecycleState != oci_psql.BackupLifecycleStateDeleted
	}
	return false
}

func PsqlBackupSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.PostgresqlClient().GetBackup(context.Background(), oci_psql.GetBackupRequest{
		BackupId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
