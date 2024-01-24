// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseOneoffPatchRequiredOnlyResource = DatabaseOneoffPatchResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_oneoff_patch", "test_oneoff_patch", acctest.Required, acctest.Create, DatabaseOneoffPatchRepresentation)

	DatabaseOneoffPatchResourceConfig = DatabaseOneoffPatchResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_oneoff_patch", "test_oneoff_patch", acctest.Optional, acctest.Update, DatabaseOneoffPatchRepresentation)

	DatabaseOneoffPatchSingularDataSourceRepresentation = map[string]interface{}{
		"oneoff_patch_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_oneoff_patch.test_oneoff_patch.id}`},
	}

	DatabaseOneoffPatchDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `19.18_RU`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseOneoffPatchDataSourceFilterRepresentation}}
	DatabaseOneoffPatchDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_oneoff_patch.test_oneoff_patch.id}`}},
	}

	DatabaseOneoffPatchRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"db_version":      acctest.Representation{RepType: acctest.Required, Create: `19.0.0.0`},
		"display_name":    acctest.Representation{RepType: acctest.Required, Create: `19.18_RU`},
		"release_update":  acctest.Representation{RepType: acctest.Required, Create: `19.18.0.0`},
		"defined_tags":    acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"one_off_patches": acctest.Representation{RepType: acctest.Optional, Create: []string{`31908573`}},
		//"download_oneoff_patch_trigger": acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `1`},
	}

	DatabaseOneoffPatchResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: database/default
func TestDatabaseOneoffPatchResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseOneoffPatchResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_oneoff_patch.test_oneoff_patch"
	datasourceName := "data.oci_database_oneoff_patches.test_oneoff_patches"
	singularDatasourceName := "data.oci_database_oneoff_patch.test_oneoff_patch"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseOneoffPatchResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_oneoff_patch", "test_oneoff_patch", acctest.Optional, acctest.Create, DatabaseOneoffPatchRepresentation), "database", "oneoffPatch", t)

	acctest.ResourceTest(t, testAccCheckDatabaseOneoffPatchDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseOneoffPatchResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_oneoff_patch", "test_oneoff_patch", acctest.Required, acctest.Create, DatabaseOneoffPatchRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "db_version", "19.0.0.0"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "19.18_RU"),
				resource.TestCheckResourceAttr(resourceName, "release_update", "19.18.0.0"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseOneoffPatchResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseOneoffPatchResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_oneoff_patch", "test_oneoff_patch", acctest.Optional, acctest.Create, DatabaseOneoffPatchRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "db_version", "19.0.0.0"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "19.18_RU"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "one_off_patches.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "release_update", "19.18.0.0"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatabaseOneoffPatchResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_oneoff_patch", "test_oneoff_patch", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatabaseOneoffPatchRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "db_version", "19.0.0.0"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "19.18_RU"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "one_off_patches.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "release_update", "19.18.0.0"),
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
			Config: config + compartmentIdVariableStr + DatabaseOneoffPatchResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_oneoff_patch", "test_oneoff_patch", acctest.Optional, acctest.Update, DatabaseOneoffPatchRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "db_version", "19.0.0.0"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "19.18_RU"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "one_off_patches.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "release_update", "19.18.0.0"),
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
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_oneoff_patches", "test_oneoff_patches", acctest.Optional, acctest.Update, DatabaseOneoffPatchDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseOneoffPatchResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_oneoff_patch", "test_oneoff_patch", acctest.Optional, acctest.Update, DatabaseOneoffPatchRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "19.18_RU"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "oneoff_patches.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "oneoff_patches.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "oneoff_patches.0.db_version", "19.0.0.0"),
				resource.TestCheckResourceAttr(datasourceName, "oneoff_patches.0.display_name", "19.18_RU"),
				resource.TestCheckResourceAttr(datasourceName, "oneoff_patches.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "oneoff_patches.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "oneoff_patches.0.one_off_patches.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "oneoff_patches.0.release_update", "19.18.0.0"),
				resource.TestCheckResourceAttrSet(datasourceName, "oneoff_patches.0.sha256sum"),
				resource.TestCheckResourceAttrSet(datasourceName, "oneoff_patches.0.size_in_kbs"),
				resource.TestCheckResourceAttrSet(datasourceName, "oneoff_patches.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "oneoff_patches.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "oneoff_patches.0.time_of_expiration"),
				resource.TestCheckResourceAttrSet(datasourceName, "oneoff_patches.0.time_updated"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_oneoff_patch", "test_oneoff_patch", acctest.Required, acctest.Create, DatabaseOneoffPatchSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseOneoffPatchResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "oneoff_patch_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_version", "19.0.0.0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "19.18_RU"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "one_off_patches.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "release_update", "19.18.0.0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sha256sum"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "size_in_kbs"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_of_expiration"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + DatabaseOneoffPatchRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDatabaseOneoffPatchDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DatabaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_oneoff_patch" {
			noResourceFound = false
			request := oci_database.GetOneoffPatchRequest{}

			tmp := rs.Primary.ID
			request.OneoffPatchId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")

			response, err := client.GetOneoffPatch(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.OneoffPatchLifecycleStateDeleted): true, string(oci_database.OneoffPatchLifecycleStateTerminated): true,
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
	if !acctest.InSweeperExcludeList("DatabaseOneoffPatch") {
		resource.AddTestSweepers("DatabaseOneoffPatch", &resource.Sweeper{
			Name:         "DatabaseOneoffPatch",
			Dependencies: acctest.DependencyGraph["oneoffPatch"],
			F:            sweepDatabaseOneoffPatchResource,
		})
	}
}

func sweepDatabaseOneoffPatchResource(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	oneoffPatchIds, err := getDatabaseOneoffPatchIds(compartment)
	if err != nil {
		return err
	}
	for _, oneoffPatchId := range oneoffPatchIds {
		if ok := acctest.SweeperDefaultResourceId[oneoffPatchId]; !ok {
			deleteOneoffPatchRequest := oci_database.DeleteOneoffPatchRequest{}

			deleteOneoffPatchRequest.OneoffPatchId = &oneoffPatchId

			deleteOneoffPatchRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteOneoffPatch(context.Background(), deleteOneoffPatchRequest)
			if error != nil {
				fmt.Printf("Error deleting OneoffPatch %s %s, It is possible that the resource is already deleted. Please verify manually \n", oneoffPatchId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &oneoffPatchId, DatabaseOneoffPatchSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseOneoffPatchSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getDatabaseOneoffPatchIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "OneoffPatchId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()

	listOneoffPatchesRequest := oci_database.ListOneoffPatchesRequest{}
	listOneoffPatchesRequest.CompartmentId = &compartmentId
	listOneoffPatchesRequest.LifecycleState = oci_database.OneoffPatchSummaryLifecycleStateAvailable
	listOneoffPatchesResponse, err := databaseClient.ListOneoffPatches(context.Background(), listOneoffPatchesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting OneoffPatch list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, oneoffPatch := range listOneoffPatchesResponse.Items {
		id := *oneoffPatch.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "OneoffPatchId", id)
	}
	return resourceIds, nil
}

func DatabaseOneoffPatchSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if oneoffPatchResponse, ok := response.Response.(oci_database.GetOneoffPatchResponse); ok {
		return oneoffPatchResponse.LifecycleState != oci_database.OneoffPatchLifecycleStateDeleted
	}
	return false
}

func DatabaseOneoffPatchSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetOneoffPatch(context.Background(), oci_database.GetOneoffPatchRequest{
		OneoffPatchId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
