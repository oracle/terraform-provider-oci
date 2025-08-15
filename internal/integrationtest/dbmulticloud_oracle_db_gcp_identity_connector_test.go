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
	DbmulticloudOracleDbGcpIdentityConnectorRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_gcp_identity_connector", "test_oracle_db_gcp_identity_connector", acctest.Required, acctest.Create, DbmulticloudOracleDbGcpIdentityConnectorRepresentation)

	DbmulticloudOracleDbGcpIdentityConnectorResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_gcp_identity_connector", "test_oracle_db_gcp_identity_connector", acctest.Optional, acctest.Update, DbmulticloudOracleDbGcpIdentityConnectorRepresentation)

	DbmulticloudOracleDbGcpIdentityConnectorSingularDataSourceRepresentation = map[string]interface{}{
		"oracle_db_gcp_identity_connector_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dbmulticloud_oracle_db_gcp_identity_connector.test_oracle_db_gcp_identity_connector.id}`},
	}

	DbmulticloudOracleDbGcpIdentityConnectorDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `TersiGCPConnector`, Update: `TersiGCPConnectorUpdated`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DbmulticloudOracleDbGcpIdentityConnectorDataSourceFilterRepresentation}}
	DbmulticloudOracleDbGcpIdentityConnectorDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_dbmulticloud_oracle_db_gcp_identity_connector.test_oracle_db_gcp_identity_connector.id}`}},
	}

	DbmulticloudOracleDbGcpIdentityConnectorRepresentation = map[string]interface{}{
		"compartment_id":                    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                      acctest.Representation{RepType: acctest.Required, Create: `TersiGCPConnector`, Update: `TersiGCPConnectorUpdated`},
		"gcp_location":                      acctest.Representation{RepType: acctest.Required, Create: `global`},
		"gcp_resource_service_agent_id":     acctest.Representation{RepType: acctest.Required, Create: `dbqa-hxeqe1@db-mc-dataplane.iam.gserviceaccount.com`},
		"gcp_workload_identity_pool_id":     acctest.Representation{RepType: acctest.Required, Create: `dbmci`},
		"gcp_workload_identity_provider_id": acctest.Representation{RepType: acctest.Required, Create: `projects/823581690520/locations/global/workloadIdentityPools/dbmci/providers/odbgsandboxdevexa`},
		"issuer_url":                        acctest.Representation{RepType: acctest.Required, Create: `https://idcs-a28ecb28ca124051898b751e1438136d.identity.oraclecloud.com`},
		"project_id":                        acctest.Representation{RepType: acctest.Required, Create: `db-mc-dataplane`},
		"resource_id":                       acctest.Representation{RepType: acctest.Required, Create: `ocid1.cloudvmcluster.test..tersitest`},
	}
)

// issue-routing-tag: dbmulticloud/default
func TestDbmulticloudOracleDbGcpIdentityConnectorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDbmulticloudOracleDbGcpIdentityConnectorResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_dbmulticloud_oracle_db_gcp_identity_connector.test_oracle_db_gcp_identity_connector"
	datasourceName := "data.oci_dbmulticloud_oracle_db_gcp_identity_connectors.test_oracle_db_gcp_identity_connectors"
	singularDatasourceName := "data.oci_dbmulticloud_oracle_db_gcp_identity_connector.test_oracle_db_gcp_identity_connector"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_gcp_identity_connector", "test_oracle_db_gcp_identity_connector", acctest.Optional, acctest.Create, DbmulticloudOracleDbGcpIdentityConnectorRepresentation), "dbmulticloud", "oracleDbGcpIdentityConnector", t)

	acctest.ResourceTest(t, testAccCheckDbmulticloudOracleDbGcpIdentityConnectorDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_gcp_identity_connector", "test_oracle_db_gcp_identity_connector", acctest.Required, acctest.Create, DbmulticloudOracleDbGcpIdentityConnectorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TersiGCPConnector"),
				resource.TestCheckResourceAttr(resourceName, "gcp_location", "global"),
				resource.TestCheckResourceAttrSet(resourceName, "gcp_resource_service_agent_id"),
				resource.TestCheckResourceAttrSet(resourceName, "gcp_workload_identity_pool_id"),
				resource.TestCheckResourceAttrSet(resourceName, "gcp_workload_identity_provider_id"),
				resource.TestCheckResourceAttr(resourceName, "issuer_url", "https://idcs-a28ecb28ca124051898b751e1438136d.identity.oraclecloud.com"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_gcp_identity_connector", "test_oracle_db_gcp_identity_connector", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DbmulticloudOracleDbGcpIdentityConnectorRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TersiGCPConnector"),
				resource.TestCheckResourceAttr(resourceName, "gcp_location", "global"),
				resource.TestCheckResourceAttrSet(resourceName, "gcp_resource_service_agent_id"),
				resource.TestCheckResourceAttrSet(resourceName, "gcp_workload_identity_pool_id"),
				resource.TestCheckResourceAttrSet(resourceName, "gcp_workload_identity_provider_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "issuer_url", "https://idcs-a28ecb28ca124051898b751e1438136d.identity.oraclecloud.com"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_id"),

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
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_gcp_identity_connector", "test_oracle_db_gcp_identity_connector", acctest.Optional, acctest.Update, DbmulticloudOracleDbGcpIdentityConnectorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TersiGCPConnectorUpdated"),
				resource.TestCheckResourceAttr(resourceName, "gcp_location", "global"),
				resource.TestCheckResourceAttrSet(resourceName, "gcp_resource_service_agent_id"),
				resource.TestCheckResourceAttrSet(resourceName, "gcp_workload_identity_pool_id"),
				resource.TestCheckResourceAttrSet(resourceName, "gcp_workload_identity_provider_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "issuer_url", "https://idcs-a28ecb28ca124051898b751e1438136d.identity.oraclecloud.com"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_dbmulticloud_oracle_db_gcp_identity_connectors", "test_oracle_db_gcp_identity_connectors", acctest.Optional, acctest.Update, DbmulticloudOracleDbGcpIdentityConnectorDataSourceRepresentation) +
				compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_gcp_identity_connector", "test_oracle_db_gcp_identity_connector", acctest.Optional, acctest.Update, DbmulticloudOracleDbGcpIdentityConnectorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "TersiGCPConnectorUpdated"),

				resource.TestCheckResourceAttr(datasourceName, "oracle_db_gcp_identity_connector_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "oracle_db_gcp_identity_connector_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dbmulticloud_oracle_db_gcp_identity_connector", "test_oracle_db_gcp_identity_connector", acctest.Required, acctest.Create, DbmulticloudOracleDbGcpIdentityConnectorSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DbmulticloudOracleDbGcpIdentityConnectorResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "oracle_db_gcp_identity_connector_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "TersiGCPConnectorUpdated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "gcp_identity_connectivity_status"),
				resource.TestCheckResourceAttr(singularDatasourceName, "gcp_location", "global"),
				resource.TestCheckResourceAttr(singularDatasourceName, "gcp_nodes.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "issuer_url", "https://idcs-a28ecb28ca124051898b751e1438136d.identity.oraclecloud.com"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		//// verify resource import
		//{
		//	Config:                  config + DbmulticloudOracleDbGcpIdentityConnectorRequiredOnlyResource,
		//	ImportState:             true,
		//	ImportStateVerify:       true,
		//	ImportStateVerifyIgnore: []string{},
		//	ResourceName:            resourceName,
		//},
	})
}

func testAccCheckDbmulticloudOracleDbGcpIdentityConnectorDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DbMulticloudGCPProviderClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dbmulticloud_oracle_db_gcp_identity_connector" {
			noResourceFound = false
			request := oci_dbmulticloud.GetOracleDbGcpIdentityConnectorRequest{}

			tmp := rs.Primary.ID
			request.OracleDbGcpIdentityConnectorId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dbmulticloud")

			response, err := client.GetOracleDbGcpIdentityConnector(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_dbmulticloud.OracleDbGcpIdentityConnectorLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DbmulticloudOracleDbGcpIdentityConnector") {
		resource.AddTestSweepers("DbmulticloudOracleDbGcpIdentityConnector", &resource.Sweeper{
			Name:         "DbmulticloudOracleDbGcpIdentityConnector",
			Dependencies: acctest.DependencyGraph["oracleDbGcpIdentityConnector"],
			F:            sweepDbmulticloudOracleDbGcpIdentityConnectorResource,
		})
	}
}

func sweepDbmulticloudOracleDbGcpIdentityConnectorResource(compartment string) error {
	dbMulticloudGCPProviderClient := acctest.GetTestClients(&schema.ResourceData{}).DbMulticloudGCPProviderClient()
	oracleDbGcpIdentityConnectorIds, err := getDbmulticloudOracleDbGcpIdentityConnectorIds(compartment)
	if err != nil {
		return err
	}
	for _, oracleDbGcpIdentityConnectorId := range oracleDbGcpIdentityConnectorIds {
		if ok := acctest.SweeperDefaultResourceId[oracleDbGcpIdentityConnectorId]; !ok {
			deleteOracleDbGcpIdentityConnectorRequest := oci_dbmulticloud.DeleteOracleDbGcpIdentityConnectorRequest{}

			deleteOracleDbGcpIdentityConnectorRequest.OracleDbGcpIdentityConnectorId = &oracleDbGcpIdentityConnectorId

			deleteOracleDbGcpIdentityConnectorRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dbmulticloud")
			_, error := dbMulticloudGCPProviderClient.DeleteOracleDbGcpIdentityConnector(context.Background(), deleteOracleDbGcpIdentityConnectorRequest)
			if error != nil {
				fmt.Printf("Error deleting OracleDbGcpIdentityConnector %s %s, It is possible that the resource is already deleted. Please verify manually \n", oracleDbGcpIdentityConnectorId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &oracleDbGcpIdentityConnectorId, DbmulticloudOracleDbGcpIdentityConnectorSweepWaitCondition, time.Duration(3*time.Minute),
				DbmulticloudOracleDbGcpIdentityConnectorSweepResponseFetchOperation, "dbmulticloud", true)
		}
	}
	return nil
}

func getDbmulticloudOracleDbGcpIdentityConnectorIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "OracleDbGcpIdentityConnectorId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dbMulticloudGCPProviderClient := acctest.GetTestClients(&schema.ResourceData{}).DbMulticloudGCPProviderClient()

	listOracleDbGcpIdentityConnectorsRequest := oci_dbmulticloud.ListOracleDbGcpIdentityConnectorsRequest{}
	listOracleDbGcpIdentityConnectorsRequest.CompartmentId = &compartmentId
	listOracleDbGcpIdentityConnectorsRequest.LifecycleState = oci_dbmulticloud.OracleDbGcpIdentityConnectorLifecycleStateActive
	listOracleDbGcpIdentityConnectorsResponse, err := dbMulticloudGCPProviderClient.ListOracleDbGcpIdentityConnectors(context.Background(), listOracleDbGcpIdentityConnectorsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting OracleDbGcpIdentityConnector list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, oracleDbGcpIdentityConnector := range listOracleDbGcpIdentityConnectorsResponse.Items {
		id := *oracleDbGcpIdentityConnector.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "OracleDbGcpIdentityConnectorId", id)
	}
	return resourceIds, nil
}

func DbmulticloudOracleDbGcpIdentityConnectorSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if oracleDbGcpIdentityConnectorResponse, ok := response.Response.(oci_dbmulticloud.GetOracleDbGcpIdentityConnectorResponse); ok {
		return oracleDbGcpIdentityConnectorResponse.LifecycleState != oci_dbmulticloud.OracleDbGcpIdentityConnectorLifecycleStateDeleted
	}
	return false
}

func DbmulticloudOracleDbGcpIdentityConnectorSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DbMulticloudGCPProviderClient().GetOracleDbGcpIdentityConnector(context.Background(), oci_dbmulticloud.GetOracleDbGcpIdentityConnectorRequest{
		OracleDbGcpIdentityConnectorId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
