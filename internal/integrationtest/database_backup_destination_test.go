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
	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v56/common"
	oci_database "github.com/oracle/oci-go-sdk/v56/database"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	BackupDestinationRequiredOnlyResource = BackupDestinationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_backup_destination", "test_backup_destination", acctest.Required, acctest.Create, backupDestinationRepresentation)

	BackupDestinationResourceConfig = BackupDestinationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_backup_destination", "test_backup_destination", acctest.Optional, acctest.Update, backupDestinationRepresentation)

	backupDestinationSingularDataSourceRepresentation = map[string]interface{}{
		"backup_destination_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_backup_destination.test_backup_destination.id}`},
	}

	backupDestinationDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"type":           acctest.Representation{RepType: acctest.Optional, Create: `RECOVERY_APPLIANCE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: backupDestinationDataSourceFilterRepresentation}}
	backupDestinationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_backup_destination.test_backup_destination.id}`}},
	}

	backupDestinationRepresentation = map[string]interface{}{
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":      acctest.Representation{RepType: acctest.Required, Create: `Recovery Appliance1`},
		"type":              acctest.Representation{RepType: acctest.Required, Create: `RECOVERY_APPLIANCE`},
		"connection_string": acctest.Representation{RepType: acctest.Optional, Create: `connectionString`, Update: `connectionString2`},
		"defined_tags":      acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":     acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"vpc_users":         acctest.Representation{RepType: acctest.Optional, Create: []string{`bkupUser1`}, Update: []string{`bkupUser1`, `bkupUser2`}},
	}

	BackupDestinationResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: database/ExaCC
func TestDatabaseBackupDestinationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseBackupDestinationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_backup_destination.test_backup_destination"
	datasourceName := "data.oci_database_backup_destinations.test_backup_destinations"
	singularDatasourceName := "data.oci_database_backup_destination.test_backup_destination"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+BackupDestinationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_backup_destination", "test_backup_destination", acctest.Optional, acctest.Create, backupDestinationRepresentation), "database", "backupDestination", t)

	acctest.ResourceTest(t, testAccCheckDatabaseBackupDestinationDestroy, []resource.TestStep{
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + BackupDestinationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_backup_destination", "test_backup_destination", acctest.Optional, acctest.Create, backupDestinationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "connection_string", "connectionString"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Recovery Appliance1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "type", "RECOVERY_APPLIANCE"),
				resource.TestCheckResourceAttr(resourceName, "vpc_users.#", "1"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + BackupDestinationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_backup_destination", "test_backup_destination", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(backupDestinationRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "connection_string", "connectionString"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Recovery Appliance1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "type", "RECOVERY_APPLIANCE"),
				resource.TestCheckResourceAttr(resourceName, "vpc_users.#", "1"),

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
			Config: config + compartmentIdVariableStr + BackupDestinationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_backup_destination", "test_backup_destination", acctest.Optional, acctest.Update, backupDestinationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "connection_string", "connectionString2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Recovery Appliance1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "type", "RECOVERY_APPLIANCE"),
				resource.TestCheckResourceAttr(resourceName, "vpc_users.#", "2"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_backup_destinations", "test_backup_destinations", acctest.Optional, acctest.Update, backupDestinationDataSourceRepresentation) +
				compartmentIdVariableStr + BackupDestinationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_backup_destination", "test_backup_destination", acctest.Optional, acctest.Update, backupDestinationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(datasourceName, "backup_destinations.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "backup_destinations.0.associated_databases.#", "0"),
				resource.TestCheckResourceAttr(datasourceName, "backup_destinations.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "backup_destinations.0.connection_string", "connectionString2"),
				resource.TestCheckResourceAttr(datasourceName, "backup_destinations.0.display_name", "Recovery Appliance1"),
				resource.TestCheckResourceAttr(datasourceName, "backup_destinations.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "backup_destinations.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "backup_destinations.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "backup_destinations.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName, "backup_destinations.0.type", "RECOVERY_APPLIANCE"),
				resource.TestCheckResourceAttr(datasourceName, "backup_destinations.0.vpc_users.#", "2"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_backup_destination", "test_backup_destination", acctest.Required, acctest.Create, backupDestinationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + BackupDestinationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "backup_destination_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "associated_databases.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_string", "connectionString2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "Recovery Appliance1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "RECOVERY_APPLIANCE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vpc_users.#", "2"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + BackupDestinationResourceConfig,
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"mount_type_details",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatabaseBackupDestinationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*client.OracleClients).DatabaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_backup_destination" {
			noResourceFound = false
			request := oci_database.GetBackupDestinationRequest{}

			tmp := rs.Primary.ID
			request.BackupDestinationId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")

			response, err := client.GetBackupDestination(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.BackupDestinationLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DatabaseBackupDestination") {
		resource.AddTestSweepers("DatabaseBackupDestination", &resource.Sweeper{
			Name:         "DatabaseBackupDestination",
			Dependencies: acctest.DependencyGraph["backupDestination"],
			F:            sweepDatabaseBackupDestinationResource,
		})
	}
}

func sweepDatabaseBackupDestinationResource(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	backupDestinationIds, err := getBackupDestinationIds(compartment)
	if err != nil {
		return err
	}
	for _, backupDestinationId := range backupDestinationIds {
		if ok := acctest.SweeperDefaultResourceId[backupDestinationId]; !ok {
			deleteBackupDestinationRequest := oci_database.DeleteBackupDestinationRequest{}

			deleteBackupDestinationRequest.BackupDestinationId = &backupDestinationId

			deleteBackupDestinationRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteBackupDestination(context.Background(), deleteBackupDestinationRequest)
			if error != nil {
				fmt.Printf("Error deleting BackupDestination %s %s, It is possible that the resource is already deleted. Please verify manually \n", backupDestinationId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &backupDestinationId, backupDestinationSweepWaitCondition, time.Duration(3*time.Minute),
				backupDestinationSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getBackupDestinationIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "BackupDestinationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()

	listBackupDestinationRequest := oci_database.ListBackupDestinationRequest{}
	listBackupDestinationRequest.CompartmentId = &compartmentId
	listBackupDestinationResponse, err := databaseClient.ListBackupDestination(context.Background(), listBackupDestinationRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting BackupDestination list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, backupDestination := range listBackupDestinationResponse.Items {
		id := *backupDestination.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "BackupDestinationId", id)
	}
	return resourceIds, nil
}

func backupDestinationSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if backupDestinationResponse, ok := response.Response.(oci_database.GetBackupDestinationResponse); ok {
		return backupDestinationResponse.LifecycleState != oci_database.BackupDestinationLifecycleStateDeleted
	}
	return false
}

func backupDestinationSweepResponseFetchOperation(client *client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetBackupDestination(context.Background(), oci_database.GetBackupDestinationRequest{
		BackupDestinationId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
