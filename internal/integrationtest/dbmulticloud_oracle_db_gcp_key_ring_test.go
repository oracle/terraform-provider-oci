// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_dbmulticloud "github.com/oracle/oci-go-sdk/v65/dbmulticloud"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DbmulticloudOracleDbGcpKeyRingRequiredOnlyResource = DbmulticloudOracleDbGcpIdentityConnectorRequiredOnlyResource + DbmulticloudOracleDbGcpIdentityConnectorRequiredOnlyResource +
		acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_gcp_key_ring", "test_oracle_db_gcp_key_ring", acctest.Required, acctest.Create, DbmulticloudOracleDbGcpKeyRingRepresentation)

	DbmulticloudOracleDbGcpKeyRingResourceConfig = DbmulticloudOracleDbGcpKeyRingResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_gcp_key_ring", "test_oracle_db_gcp_key_ring", acctest.Optional, acctest.Update, DbmulticloudOracleDbGcpKeyRingRepresentation)

	DbmulticloudOracleDbGcpKeyRingSingularDataSourceRepresentation = map[string]interface{}{
		"oracle_db_gcp_key_ring_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dbmulticloud_oracle_db_gcp_key_ring.test_oracle_db_gcp_key_ring.id}`},
	}

	DbmulticloudOracleDbGcpKeyRingDataSourceRepresentation = map[string]interface{}{
		"compartment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":               acctest.Representation{RepType: acctest.Optional, Create: `TestDbGcpVault`, Update: `TestDbGcpVault2`},
		"oracle_db_gcp_connector_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_dbmulticloud_oracle_db_gcp_identity_connector.test_oracle_db_gcp_identity_connector.id}`},
		//"oracle_db_gcp_key_ring_id":  acctest.Representation{RepType: acctest.Optional, Create: `${oci_dbmulticloud_oracle_db_gcp_key_ring.test_oracle_db_gcp_key_ring.id}`},
		"filter": acctest.RepresentationGroup{RepType: acctest.Required, Group: DbmulticloudOracleDbGcpKeyRingDataSourceFilterRepresentation}}
	DbmulticloudOracleDbGcpKeyRingDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_dbmulticloud_oracle_db_gcp_key_ring.test_oracle_db_gcp_key_ring.id}`}},
	}

	DbmulticloudOracleDbGcpKeyRingRepresentation = map[string]interface{}{
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":           acctest.Representation{RepType: acctest.Required, Create: `TestDbGcpVault`, Update: `TestDbGcpVault2`},
		"oracle_db_connector_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dbmulticloud_oracle_db_gcp_identity_connector.test_oracle_db_gcp_identity_connector.id}`},
		"gcp_key_ring_id":        acctest.Representation{RepType: acctest.Required, Create: `test key ring`},
		"location":               acctest.Representation{RepType: acctest.Optional, Create: `location`},
		"type":                   acctest.Representation{RepType: acctest.Optional, Create: `type`},
	}

	DbmulticloudOracleDbGcpKeyRingResourceDependencies = ""
)

// issue-routing-tag: dbmulticloud/default
func TestDbmulticloudOracleDbGcpKeyRingResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDbmulticloudOracleDbGcpKeyRingResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_dbmulticloud_oracle_db_gcp_key_ring.test_oracle_db_gcp_key_ring"
	datasourceName := "data.oci_dbmulticloud_oracle_db_gcp_key_rings.test_oracle_db_gcp_key_rings"
	singularDatasourceName := "data.oci_dbmulticloud_oracle_db_gcp_key_ring.test_oracle_db_gcp_key_ring"
	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	//acctest.SaveConfigContent(config+compartmentIdVariableStr+DbmulticloudOracleDbGcpKeyRingResourceDependencies+
	//	acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_gcp_key_ring", "test_oracle_db_gcp_key_ring", acctest.Optional, acctest.Create, DbmulticloudOracleDbGcpKeyRingRepresentation), "dbmulticloud", "oracleDbGcpKeyRing", t)

	acctest.ResourceTest(t, testAccCheckDbmulticloudOracleDbGcpKeyRingDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DbmulticloudOracleDbGcpKeyRingResourceDependencies + DbmulticloudOracleDbGcpIdentityConnectorRequiredOnlyResource +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_gcp_key_ring", "test_oracle_db_gcp_key_ring", acctest.Optional, acctest.Create, DbmulticloudOracleDbGcpKeyRingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TestDbGcpVault"),
				resource.TestCheckResourceAttrSet(resourceName, "oracle_db_connector_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DbmulticloudOracleDbGcpKeyRingResourceDependencies + DbmulticloudOracleDbGcpIdentityConnectorRequiredOnlyResource +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_gcp_key_ring", "test_oracle_db_gcp_key_ring", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DbmulticloudOracleDbGcpKeyRingRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TestDbGcpVault"),
				resource.TestCheckResourceAttrSet(resourceName, "gcp_key_ring_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "location", "location"),
				resource.TestCheckResourceAttrSet(resourceName, "oracle_db_connector_id"),
				//resource.TestCheckResourceAttr(resourceName, "properties.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "type", "type"),

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
			Config: config + compartmentIdVariableStr + DbmulticloudOracleDbGcpKeyRingResourceDependencies + DbmulticloudOracleDbGcpIdentityConnectorRequiredOnlyResource +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_gcp_key_ring", "test_oracle_db_gcp_key_ring", acctest.Optional, acctest.Update, DbmulticloudOracleDbGcpKeyRingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TestDbGcpVault2"),
				resource.TestCheckResourceAttrSet(resourceName, "gcp_key_ring_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "location", "location"),
				resource.TestCheckResourceAttrSet(resourceName, "oracle_db_connector_id"),
				//resource.TestCheckResourceAttr(resourceName, "properties.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "type", "type"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_dbmulticloud_oracle_db_gcp_key_rings", "test_oracle_db_gcp_key_rings", acctest.Optional, acctest.Update, DbmulticloudOracleDbGcpKeyRingDataSourceRepresentation) +
				compartmentIdVariableStr + DbmulticloudOracleDbGcpKeyRingResourceDependencies +
				DbmulticloudOracleDbGcpIdentityConnectorRequiredOnlyResource +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_gcp_key_ring", "test_oracle_db_gcp_key_ring", acctest.Optional, acctest.Update, DbmulticloudOracleDbGcpKeyRingRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "TestDbGcpVault2"),
				resource.TestCheckResourceAttrSet(datasourceName, "oracle_db_gcp_connector_id"),

				resource.TestCheckResourceAttr(datasourceName, "oracle_db_gcp_key_ring_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "oracle_db_gcp_key_ring_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dbmulticloud_oracle_db_gcp_key_ring", "test_oracle_db_gcp_key_ring", acctest.Required, acctest.Create, DbmulticloudOracleDbGcpKeyRingSingularDataSourceRepresentation) + DbmulticloudOracleDbGcpIdentityConnectorRequiredOnlyResource +
				compartmentIdVariableStr + DbmulticloudOracleDbGcpKeyRingResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "oracle_db_gcp_key_ring_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "TestDbGcpVault2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "location", "location"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "properties.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "type"),
			),
		},
		//// verify resource import
		//{
		//	Config:                  config + DbmulticloudOracleDbGcpKeyRingRequiredOnlyResource,
		//	ImportState:             true,
		//	ImportStateVerify:       true,
		//	ImportStateVerifyIgnore: []string{},
		//	ResourceName:            resourceName,
		//},
	})
}

func testAccCheckDbmulticloudOracleDbGcpKeyRingDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DbMulticloudGCPProviderClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dbmulticloud_oracle_db_gcp_key_ring" {
			noResourceFound = false
			request := oci_dbmulticloud.GetOracleDbGcpKeyRingRequest{}

			tmp := rs.Primary.ID
			request.OracleDbGcpKeyRingId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dbmulticloud")

			response, err := client.GetOracleDbGcpKeyRing(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_dbmulticloud.OracleDbGcpKeyRingLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DbmulticloudOracleDbGcpKeyRing") {
		resource.AddTestSweepers("DbmulticloudOracleDbGcpKeyRing", &resource.Sweeper{
			Name:         "DbmulticloudOracleDbGcpKeyRing",
			Dependencies: acctest.DependencyGraph["oracleDbGcpKeyRing"],
			F:            sweepDbmulticloudOracleDbGcpKeyRingResource,
		})
	}
}

func sweepDbmulticloudOracleDbGcpKeyRingResource(compartment string) error {
	dbMulticloudGCPProviderClient := acctest.GetTestClients(&schema.ResourceData{}).DbMulticloudGCPProviderClient()
	oracleDbGcpKeyRingIds, err := getDbmulticloudOracleDbGcpKeyRingIds(compartment)
	if err != nil {
		return err
	}
	for _, oracleDbGcpKeyRingId := range oracleDbGcpKeyRingIds {
		if ok := acctest.SweeperDefaultResourceId[oracleDbGcpKeyRingId]; !ok {
			deleteOracleDbGcpKeyRingRequest := oci_dbmulticloud.DeleteOracleDbGcpKeyRingRequest{}

			deleteOracleDbGcpKeyRingRequest.OracleDbGcpKeyRingId = &oracleDbGcpKeyRingId

			deleteOracleDbGcpKeyRingRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dbmulticloud")
			_, error := dbMulticloudGCPProviderClient.DeleteOracleDbGcpKeyRing(context.Background(), deleteOracleDbGcpKeyRingRequest)
			if error != nil {
				fmt.Printf("Error deleting OracleDbGcpKeyRing %s %s, It is possible that the resource is already deleted. Please verify manually \n", oracleDbGcpKeyRingId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &oracleDbGcpKeyRingId, DbmulticloudOracleDbGcpKeyRingSweepWaitCondition, time.Duration(3*time.Minute),
				DbmulticloudOracleDbGcpKeyRingSweepResponseFetchOperation, "dbmulticloud", true)
		}
	}
	return nil
}

func getDbmulticloudOracleDbGcpKeyRingIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "OracleDbGcpKeyRingId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dbMulticloudGCPProviderClient := acctest.GetTestClients(&schema.ResourceData{}).DbMulticloudGCPProviderClient()

	listOracleDbGcpKeyRingsRequest := oci_dbmulticloud.ListOracleDbGcpKeyRingsRequest{}
	listOracleDbGcpKeyRingsRequest.CompartmentId = &compartmentId
	listOracleDbGcpKeyRingsRequest.LifecycleState = oci_dbmulticloud.OracleDbGcpKeyRingLifecycleStateActive
	listOracleDbGcpKeyRingsResponse, err := dbMulticloudGCPProviderClient.ListOracleDbGcpKeyRings(context.Background(), listOracleDbGcpKeyRingsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting OracleDbGcpKeyRing list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, oracleDbGcpKeyRing := range listOracleDbGcpKeyRingsResponse.Items {
		id := *oracleDbGcpKeyRing.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "OracleDbGcpKeyRingId", id)
	}
	return resourceIds, nil
}

func DbmulticloudOracleDbGcpKeyRingSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if oracleDbGcpKeyRingResponse, ok := response.Response.(oci_dbmulticloud.GetOracleDbGcpKeyRingResponse); ok {
		return oracleDbGcpKeyRingResponse.LifecycleState != oci_dbmulticloud.OracleDbGcpKeyRingLifecycleStateDeleted
	}
	return false
}

func DbmulticloudOracleDbGcpKeyRingSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DbMulticloudGCPProviderClient().GetOracleDbGcpKeyRing(context.Background(), oci_dbmulticloud.GetOracleDbGcpKeyRingRequest{
		OracleDbGcpKeyRingId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
