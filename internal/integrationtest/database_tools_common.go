package integrationtest

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database_tools "github.com/oracle/oci-go-sdk/v65/databasetools"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

func databaseToolsInitDependencyGraphAndSweeper(name string) {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList(name) {
		resource.AddTestSweepers(name, &resource.Sweeper{
			Name:         name,
			Dependencies: acctest.DependencyGraph["databaseToolsConnection"],
			F:            sweepDatabaseToolsDatabaseToolsConnectionResource,
		})
	}
}

func databaseToolsOciProvider() map[string]*schema.Provider {
	return map[string]*schema.Provider{
		"oci": acctest.TestAccProvider,
	}
}

func databaseToolsStandardVariables() string {
	secretId := fmt.Sprintf("variable \"secret_id\" { default = \"%s\" }\n", utils.GetEnvSettingWithBlankDefault("secret_id"))
	relatedResourceId := fmt.Sprintf("variable \"related_resource_id\" { default = \"%s\" }\n", utils.GetEnvSettingWithBlankDefault("related_resource_id"))
	compartmentId := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", utils.GetEnvSettingWithBlankDefault("compartment_ocid"))
	return secretId + relatedResourceId + compartmentId
}

func testAccCheckDatabaseToolsDatabaseToolsConnectionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.GetTestClients(&schema.ResourceData{}).DatabaseToolsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_tools_database_tools_connection" {
			noResourceFound = false
			request := oci_database_tools.GetDatabaseToolsConnectionRequest{}

			tmp := rs.Primary.ID
			request.DatabaseToolsConnectionId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_tools")

			response, err := client.GetDatabaseToolsConnection(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database_tools.LifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.DatabaseToolsConnection.GetLifecycleState())]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.DatabaseToolsConnection.GetLifecycleState())
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

func sweepDatabaseToolsDatabaseToolsConnectionResource(compartment string) error {
	fmt.Printf("Sweeping database tools connections")
	databaseToolsClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseToolsClient()
	databaseToolsConnectionIds, err := getDatabaseToolsConnectionIds(compartment)
	if err != nil {
		return err
	}
	for _, databaseToolsConnectionId := range databaseToolsConnectionIds {
		if ok := acctest.SweeperDefaultResourceId[databaseToolsConnectionId]; !ok {
			deleteDatabaseToolsConnectionRequest := oci_database_tools.DeleteDatabaseToolsConnectionRequest{}

			deleteDatabaseToolsConnectionRequest.DatabaseToolsConnectionId = &databaseToolsConnectionId

			deleteDatabaseToolsConnectionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_tools")
			_, error := databaseToolsClient.DeleteDatabaseToolsConnection(context.Background(), deleteDatabaseToolsConnectionRequest)
			if error != nil {
				fmt.Printf("Error deleting DatabaseToolsConnection %s %s, It is possible that the resource is already deleted. Please verify manually \n", databaseToolsConnectionId, error)
				continue
			} else {
				fmt.Printf("Sweeper initiated deletion of DatabaseToolsConnection %s", databaseToolsConnectionId)
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &databaseToolsConnectionId, databaseToolsConnectionSweepWaitCondition, time.Duration(3*time.Minute),
				databaseToolsConnectionSweepResponseFetchOperation, "database_tools", true)
		}
	}
	return nil
}

func getDatabaseToolsConnectionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DatabaseToolsConnectionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseToolsClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseToolsClient()

	listDatabaseToolsConnectionsRequest := oci_database_tools.ListDatabaseToolsConnectionsRequest{}
	listDatabaseToolsConnectionsRequest.CompartmentId = &compartment
	listDatabaseToolsConnectionsRequest.LifecycleState = oci_database_tools.ListDatabaseToolsConnectionsLifecycleStateActive
	listDatabaseToolsConnectionsResponse, err := databaseToolsClient.ListDatabaseToolsConnections(context.Background(), listDatabaseToolsConnectionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DatabaseToolsConnection list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, databaseToolsConnection := range listDatabaseToolsConnectionsResponse.Items {
		id := *databaseToolsConnection.GetId()
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DatabaseToolsConnectionId", id)
	}
	return resourceIds, nil
}

func databaseToolsConnectionSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if databaseToolsConnectionResponse, ok := response.Response.(oci_database_tools.GetDatabaseToolsConnectionResponse); ok {
		return databaseToolsConnectionResponse.DatabaseToolsConnection.GetLifecycleState() != oci_database_tools.LifecycleStateDeleted
	}
	return false
}

func databaseToolsConnectionSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseToolsClient().GetDatabaseToolsConnection(context.Background(), oci_database_tools.GetDatabaseToolsConnectionRequest{
		DatabaseToolsConnectionId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
