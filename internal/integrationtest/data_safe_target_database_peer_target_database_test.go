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
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataSafeTargetDatabasePeerTargetDatabaseRequiredOnlyResource = DataSafeTargetDatabasePeerTargetDatabaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_target_database_peer_target_database", "test_target_database_peer_target_database", acctest.Required, acctest.Create, DataSafeTargetDatabasePeerTargetDatabaseRepresentation)

	DataSafeTargetDatabasePeerTargetDatabaseResourceConfig = DataSafeTargetDatabasePeerTargetDatabaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_target_database_peer_target_database", "test_target_database_peer_target_database", acctest.Optional, acctest.Update, DataSafeTargetDatabasePeerTargetDatabaseRepresentation)

	DataSafeTargetDatabasePeerTargetDatabaseSingularDataSourceRepresentation = map[string]interface{}{
		"peer_target_database_id": acctest.Representation{RepType: acctest.Required, Create: `{}`},
		"target_database_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.target_id}`},
	}

	DataSafeTargetDatabasePeerTargetDatabaseDataSourceRepresentation = map[string]interface{}{
		"target_database_id": acctest.Representation{RepType: acctest.Required, Create: `${var.target_id}`},
		"filter":             acctest.RepresentationGroup{RepType: acctest.Required, Group: DataSafeTargetDatabasePeerTargetDatabaseDataSourceFilterRepresentation},
	}
	DataSafeTargetDatabasePeerTargetDatabaseDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_data_safe_target_database_peer_target_database.test_target_database_peer_target_database.id}`}},
	}

	DataSafeTargetDatabasePeerTargetDatabaseRepresentation = map[string]interface{}{
		"database_details":   acctest.RepresentationGroup{RepType: acctest.Required, Group: DataSafeTargetDatabasePeerTargetDatabaseDatabaseDetailsRepresentation},
		"target_database_id": acctest.Representation{RepType: acctest.Required, Create: `${var.target_id}`},
		"description":        acctest.Representation{RepType: acctest.Required, Create: `description`, Update: `description2`},
		"display_name":       acctest.Representation{RepType: acctest.Required, Create: `standby`, Update: `displayName2`},
	}
	DataSafeTargetDatabasePeerTargetDatabaseDatabaseDetailsRepresentation = map[string]interface{}{
		"database_type":       acctest.Representation{RepType: acctest.Required, Create: `DATABASE_CLOUD_SERVICE`, Update: `DATABASE_CLOUD_SERVICE`},
		"infrastructure_type": acctest.Representation{RepType: acctest.Required, Create: `ORACLE_CLOUD`, Update: `ORACLE_CLOUD`},
		"db_system_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.db_system_id}`},
		"listener_port":       acctest.Representation{RepType: acctest.Required, Create: `1521`, Update: `1521`},
		"service_name":        acctest.Representation{RepType: acctest.Required, Create: `DB1116_pdb1.sub06132343240.testtargetvcnad.oraclevcn.com`},
		//"lifecycle":              acctest.RepresentationGroup{RepType: acctest.Required, Group: ignorePeerTargetDatabaseRep},
	}

	DataSafeTargetDatabasePeerTargetDatabaseResourceDependencies = utils.OciImageIdsVariable + DefinedTagsDependencies

	ignorePeerTargetDatabaseRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `freeform_tags`}},
	}
)

// issue-routing-tag: data_safe/default
func TestDataSafeTargetDatabasePeerTargetDatabaseResource_basic(t *testing.T) {
	t.Skip("Needs real DBCS/ExaCS, not a fake resource. Skipping due to resource and maintainability constraints")
	httpreplay.SetScenario("TestDataSafeTargetDatabasePeerTargetDatabaseResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_adg_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	dbSystemId := utils.GetEnvSettingWithBlankDefault("data_safe_adg_db_system_ocid")
	dbSystemIdVariableStr := fmt.Sprintf("variable \"db_system_id\" { default = \"%s\" }\n", dbSystemId)

	resourceName := "oci_data_safe_target_database_peer_target_database.test_target_database_peer_target_database"
	datasourceName := "data.oci_data_safe_target_database_peer_target_databases.test_target_database_peer_target_databases"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DataSafeTargetDatabasePeerTargetDatabaseResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_data_safe_target_database_peer_target_database", "test_target_database_peer_target_database", acctest.Optional, acctest.Create, DataSafeTargetDatabasePeerTargetDatabaseRepresentation), "datasafe", "targetDatabasePeerTargetDatabase", t)

	acctest.ResourceTest(t, testAccCheckDataSafeTargetDatabasePeerTargetDatabaseDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DataSafeTargetDatabasePeerTargetDatabaseResourceDependencies + targetIdVariableStr + dbSystemIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_target_database_peer_target_database", "test_target_database_peer_target_database", acctest.Required, acctest.Create, DataSafeTargetDatabasePeerTargetDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "database_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "database_details.0.db_system_id"),
				resource.TestCheckResourceAttr(resourceName, "database_details.0.database_type", "DATABASE_CLOUD_SERVICE"),
				resource.TestCheckResourceAttr(resourceName, "database_details.0.infrastructure_type", "ORACLE_CLOUD"),
				resource.TestCheckResourceAttr(resourceName, "database_details.0.listener_port", "1521"),
				resource.TestCheckResourceAttrSet(resourceName, "database_details.0.service_name"),
				resource.TestCheckResourceAttrSet(resourceName, "target_database_id"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DataSafeTargetDatabasePeerTargetDatabaseResourceDependencies + targetIdVariableStr + dbSystemIdVariableStr,
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DataSafeTargetDatabasePeerTargetDatabaseResourceDependencies + targetIdVariableStr + dbSystemIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_target_database_peer_target_database", "test_target_database_peer_target_database", acctest.Optional, acctest.Update, DataSafeTargetDatabasePeerTargetDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "database_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_target_database_peer_target_databases", "test_target_database_peer_target_databases", acctest.Optional, acctest.Update, DataSafeTargetDatabasePeerTargetDatabaseDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeTargetDatabasePeerTargetDatabaseResourceDependencies + targetIdVariableStr + dbSystemIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_data_safe_target_database_peer_target_database", "test_target_database_peer_target_database", acctest.Optional, acctest.Update, DataSafeTargetDatabasePeerTargetDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "target_database_id"),
				resource.TestCheckResourceAttr(datasourceName, "peer_target_database_collection.#", "1"),
			),
		},
		// verify resource import
		{
			Config:                  config + DataSafeTargetDatabasePeerTargetDatabaseRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDataSafeTargetDatabasePeerTargetDatabaseDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataSafeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_data_safe_target_database_peer_target_database" {
			noResourceFound = false
			request := oci_data_safe.GetPeerTargetDatabaseRequest{}

			if value, ok := rs.Primary.Attributes["key"]; ok {
				tmp, _ := strconv.Atoi(value)
				request.PeerTargetDatabaseId = &tmp
			}

			if value, ok := rs.Primary.Attributes["target_database_id"]; ok {
				request.TargetDatabaseId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")

			response, err := client.GetPeerTargetDatabase(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_data_safe.TargetDatabaseLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DataSafeTargetDatabasePeerTargetDatabase") {
		resource.AddTestSweepers("DataSafeTargetDatabasePeerTargetDatabase", &resource.Sweeper{
			Name:         "DataSafeTargetDatabasePeerTargetDatabase",
			Dependencies: acctest.DependencyGraph["targetDatabasePeerTargetDatabase"],
			F:            sweepDataSafeTargetDatabasePeerTargetDatabaseResource,
		})
	}
}

func sweepDataSafeTargetDatabasePeerTargetDatabaseResource(compartment string) error {
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()
	targetDatabasePeerTargetDatabaseIds, err := getDataSafeTargetDatabasePeerTargetDatabaseIds(compartment)
	if err != nil {
		return err
	}
	for _, targetDatabasePeerTargetDatabaseId := range targetDatabasePeerTargetDatabaseIds {
		if ok := acctest.SweeperDefaultResourceId[targetDatabasePeerTargetDatabaseId]; !ok {
			deletePeerTargetDatabaseRequest := oci_data_safe.DeletePeerTargetDatabaseRequest{}

			deletePeerTargetDatabaseRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeletePeerTargetDatabase(context.Background(), deletePeerTargetDatabaseRequest)
			if error != nil {
				fmt.Printf("Error deleting TargetDatabasePeerTargetDatabase %s %s, It is possible that the resource is already deleted. Please verify manually \n", targetDatabasePeerTargetDatabaseId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &targetDatabasePeerTargetDatabaseId, DataSafeTargetDatabasePeerTargetDatabaseSweepWaitCondition, time.Duration(3*time.Minute),
				DataSafeTargetDatabasePeerTargetDatabaseSweepResponseFetchOperation, "data_safe", true)
		}
	}
	return nil
}

func getDataSafeTargetDatabasePeerTargetDatabaseIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "TargetDatabasePeerTargetDatabaseId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()

	listPeerTargetDatabasesRequest := oci_data_safe.ListPeerTargetDatabasesRequest{}

	targetDatabaseIds, error := getDataSafeTargetDatabaseIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting targetDatabaseId required for TargetDatabasePeerTargetDatabase resource requests \n")
	}
	for _, targetDatabaseId := range targetDatabaseIds {
		listPeerTargetDatabasesRequest.TargetDatabaseId = &targetDatabaseId

		listPeerTargetDatabasesResponse, err := dataSafeClient.ListPeerTargetDatabases(context.Background(), listPeerTargetDatabasesRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting TargetDatabasePeerTargetDatabase list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, targetDatabasePeerTargetDatabase := range listPeerTargetDatabasesResponse.Items {
			id := *targetDatabasePeerTargetDatabase.Key
			resourceIds = append(resourceIds, strconv.Itoa(id))
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "TargetDatabasePeerTargetDatabaseId", strconv.Itoa(id))
		}

	}
	return resourceIds, nil
}

func DataSafeTargetDatabasePeerTargetDatabaseSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if targetDatabasePeerTargetDatabaseResponse, ok := response.Response.(oci_data_safe.GetPeerTargetDatabaseResponse); ok {
		return targetDatabasePeerTargetDatabaseResponse.LifecycleState != oci_data_safe.TargetDatabaseLifecycleStateDeleted
	}
	return false
}

func DataSafeTargetDatabasePeerTargetDatabaseSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataSafeClient().GetPeerTargetDatabase(context.Background(), oci_data_safe.GetPeerTargetDatabaseRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}
