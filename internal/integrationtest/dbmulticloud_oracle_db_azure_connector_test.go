// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"

	// "strconv"
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

	// "github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	dbclusterResourceId = `ocid1.cloudvmcluster.test..` + utils.RandomString(50, utils.CharsetWithoutDigits)

	DbmulticloudOracleDbAzureConnectorRequiredOnlyResource = DbmulticloudOracleDbAzureConnectorResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_connector", "test_oracle_db_azure_connector", acctest.Required, acctest.Create, DbmulticloudOracleDbAzureConnectorRepresentation)

	DbmulticloudOracleDbAzureConnectorResourceConfig = DbmulticloudOracleDbAzureConnectorResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_connector", "test_oracle_db_azure_connector", acctest.Optional, acctest.Update, DbmulticloudOracleDbAzureConnectorRepresentation)

	DbmulticloudOracleDbAzureConnectorSingularDataSourceRepresentation = map[string]interface{}{
		"oracle_db_azure_connector_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dbmulticloud_oracle_db_azure_connector.test_oracle_db_azure_connector.id}`},
	}

	DbmulticloudOracleDbAzureConnectorDataSourceRepresentation = map[string]interface{}{
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: `AzureConnectorTest-Tersi`, Update: `AzureConnectorTest-Tersi`},
		"db_cluster_resource_id": acctest.Representation{RepType: acctest.Required, Create: dbclusterResourceId},
		"state":                  acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		// "access_token":           acctest.Representation{RepType: acctest.Required, Create: `AzureAccessToken`, Update: `AzureAccessToken`},
		"filter": acctest.RepresentationGroup{RepType: acctest.Required, Group: DbmulticloudOracleDbAzureConnectorDataSourceFilterRepresentation}}
	DbmulticloudOracleDbAzureConnectorDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_dbmulticloud_oracle_db_azure_connector.test_oracle_db_azure_connector.id}`}},
	}

	DbmulticloudOracleDbAzureConnectorRepresentation = map[string]interface{}{
		"azure_identity_mechanism": acctest.Representation{RepType: acctest.Required, Create: `ARC_AGENT`, Update: `ARC_AGENT`},
		"azure_resource_group":     acctest.Representation{RepType: acctest.Required, Create: `Prasanna.RG`, Update: `Prasanna.RG`},
		"azure_subscription_id":    acctest.Representation{RepType: acctest.Required, Create: `7080446f-ee76-4aa2-b9dd-c2625f63cab0`},
		"azure_tenant_id":          acctest.Representation{RepType: acctest.Required, Create: `5b743bc7-c1e2-4d46-b4b5-a32eddac0286`},
		"compartment_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"db_cluster_resource_id":   acctest.Representation{RepType: acctest.Required, Create: dbclusterResourceId},
		"display_name":             acctest.Representation{RepType: acctest.Required, Create: `AzureConnectorTest-Tersi`, Update: `AzureConnectorTest-Tersi`},
		"access_token":             acctest.Representation{RepType: acctest.Required, Create: `AzureAccessToken`, Update: `AzureAccessToken`},
	}

	DbmulticloudOracleDbAzureConnectorResourceDependencies = ""
)

// issue-routing-tag: dbmulticloud/default
func TestDbmulticloudOracleDbAzureConnectorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDbmulticloudOracleDbAzureConnectorResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_dbmulticloud_oracle_db_azure_connector.test_oracle_db_azure_connector"
	datasourceName := "data.oci_dbmulticloud_oracle_db_azure_connectors.test_oracle_db_azure_connectors"
	singularDatasourceName := "data.oci_dbmulticloud_oracle_db_azure_connector.test_oracle_db_azure_connector"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DbmulticloudOracleDbAzureConnectorResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_connector", "test_oracle_db_azure_connector", acctest.Optional, acctest.Create, DbmulticloudOracleDbAzureConnectorRepresentation), "dbmulticloud", "oracleDbAzureConnector", t)

	acctest.ResourceTest(t, testAccCheckDbmulticloudOracleDbAzureConnectorDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DbmulticloudOracleDbAzureConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_connector", "test_oracle_db_azure_connector", acctest.Required, acctest.Create, DbmulticloudOracleDbAzureConnectorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// resource.TestCheckResourceAttr(resourceName, "access_token", "Azure-accessToken"),
				resource.TestCheckResourceAttrSet(resourceName, "access_token"),
				resource.TestCheckResourceAttr(resourceName, "azure_identity_mechanism", "ARC_AGENT"),
				resource.TestCheckResourceAttr(resourceName, "azure_resource_group", "Prasanna.RG"),
				resource.TestCheckResourceAttrSet(resourceName, "azure_subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "azure_tenant_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "db_cluster_resource_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "AzureConnectorTest-Tersi"),

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
			Config: config + compartmentIdVariableStr + DbmulticloudOracleDbAzureConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_connector", "test_oracle_db_azure_connector", acctest.Optional, acctest.Create, DbmulticloudOracleDbAzureConnectorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// resource.TestCheckResourceAttr(resourceName, "access_token", "Azure-accessToken"),
				resource.TestCheckResourceAttrSet(resourceName, "access_token"),
				resource.TestCheckResourceAttr(resourceName, "azure_identity_mechanism", "ARC_AGENT"),
				resource.TestCheckResourceAttr(resourceName, "azure_resource_group", "Prasanna.RG"),
				resource.TestCheckResourceAttrSet(resourceName, "azure_subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "azure_tenant_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "db_cluster_resource_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "AzureConnectorTest-Tersi"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					// if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
					// 	if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
					// 		return errExport
					// 	}
					// }
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DbmulticloudOracleDbAzureConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_connector", "test_oracle_db_azure_connector", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DbmulticloudOracleDbAzureConnectorRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// resource.TestCheckResourceAttr(resourceName, "access_token", "Azure-accessToken"),
				resource.TestCheckResourceAttrSet(resourceName, "access_token"),
				resource.TestCheckResourceAttr(resourceName, "azure_identity_mechanism", "ARC_AGENT"),
				resource.TestCheckResourceAttr(resourceName, "azure_resource_group", "Prasanna.RG"),
				resource.TestCheckResourceAttrSet(resourceName, "azure_subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "azure_tenant_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "db_cluster_resource_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "AzureConnectorTest-Tersi"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

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
			Config: config + compartmentIdVariableStr + DbmulticloudOracleDbAzureConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_connector", "test_oracle_db_azure_connector", acctest.Optional, acctest.Update, DbmulticloudOracleDbAzureConnectorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// resource.TestCheckResourceAttr(resourceName, "access_token", "Azure-accessToken"),
				resource.TestCheckResourceAttrSet(resourceName, "access_token"),
				resource.TestCheckResourceAttr(resourceName, "azure_identity_mechanism", "ARC_AGENT"),
				resource.TestCheckResourceAttr(resourceName, "azure_resource_group", "Prasanna.RG"),
				resource.TestCheckResourceAttrSet(resourceName, "azure_subscription_id"),
				resource.TestCheckResourceAttrSet(resourceName, "azure_tenant_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "db_cluster_resource_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "AzureConnectorTest-Tersi"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_connectors", "test_oracle_db_azure_connectors", acctest.Optional, acctest.Update, DbmulticloudOracleDbAzureConnectorDataSourceRepresentation) +
				compartmentIdVariableStr + DbmulticloudOracleDbAzureConnectorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_connector", "test_oracle_db_azure_connector", acctest.Optional, acctest.Update, DbmulticloudOracleDbAzureConnectorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "db_cluster_resource_id"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dbmulticloud_oracle_db_azure_connector", "test_oracle_db_azure_connector", acctest.Required, acctest.Create, DbmulticloudOracleDbAzureConnectorSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DbmulticloudOracleDbAzureConnectorResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				//resource.TestCheckResourceAttr(singularDatasourceName, "access_token", "accessToken2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "arc_agent_nodes.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "azure_identity_connectivity_status"),
				resource.TestCheckResourceAttr(singularDatasourceName, "azure_identity_mechanism", "ARC_AGENT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "azure_resource_group", "Prasanna.RG"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "AzureConnectorTest-Tersi"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "db_cluster_resource_id"),
			),
		},
		// verify resource import
		{
			Config:            config + DbmulticloudOracleDbAzureConnectorRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"new_version",
				"access_token",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDbmulticloudOracleDbAzureConnectorDestroy(s *terraform.State) error {

	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).OracleDBAzureConnectorClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dbmulticloud_oracle_db_azure_connector" {
			noResourceFound = false
			request := oci_dbmulticloud.GetOracleDbAzureConnectorRequest{}

			tmp := rs.Primary.ID
			request.OracleDbAzureConnectorId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dbmulticloud")

			response, err := client.GetOracleDbAzureConnector(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_dbmulticloud.OracleDbAzureConnectorLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DbmulticloudOracleDbAzureConnector") {
		resource.AddTestSweepers("DbmulticloudOracleDbAzureConnector", &resource.Sweeper{
			Name:         "DbmulticloudOracleDbAzureConnector",
			Dependencies: acctest.DependencyGraph["oracleDbAzureConnector"],
			F:            sweepDbmulticloudOracleDbAzureConnectorResource,
		})
	}
}

func sweepDbmulticloudOracleDbAzureConnectorResource(compartment string) error {
	oracleDBAzureConnectorClient := acctest.GetTestClients(&schema.ResourceData{}).OracleDBAzureConnectorClient()
	oracleDbAzureConnectorIds, err := getDbmulticloudOracleDbAzureConnectorIds(compartment)
	if err != nil {
		return err
	}
	for _, oracleDbAzureConnectorId := range oracleDbAzureConnectorIds {
		if ok := acctest.SweeperDefaultResourceId[oracleDbAzureConnectorId]; !ok {
			deleteOracleDbAzureConnectorRequest := oci_dbmulticloud.DeleteOracleDbAzureConnectorRequest{}

			deleteOracleDbAzureConnectorRequest.OracleDbAzureConnectorId = &oracleDbAzureConnectorId

			deleteOracleDbAzureConnectorRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dbmulticloud")
			_, error := oracleDBAzureConnectorClient.DeleteOracleDbAzureConnector(context.Background(), deleteOracleDbAzureConnectorRequest)
			if error != nil {
				fmt.Printf("Error deleting OracleDbAzureConnector %s %s, It is possible that the resource is already deleted. Please verify manually \n", oracleDbAzureConnectorId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &oracleDbAzureConnectorId, DbmulticloudOracleDbAzureConnectorSweepWaitCondition, time.Duration(3*time.Minute),
				DbmulticloudOracleDbAzureConnectorSweepResponseFetchOperation, "dbmulticloud", true)
		}
	}
	return nil
}

func getDbmulticloudOracleDbAzureConnectorIds(compartment string) ([]string, error) {

	ids := acctest.GetResourceIdsToSweep(compartment, "OracleDbAzureConnectorId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	oracleDBAzureConnectorClient := acctest.GetTestClients(&schema.ResourceData{}).OracleDBAzureConnectorClient()

	listOracleDbAzureConnectorsRequest := oci_dbmulticloud.ListOracleDbAzureConnectorsRequest{}
	listOracleDbAzureConnectorsRequest.CompartmentId = &compartmentId
	listOracleDbAzureConnectorsRequest.LifecycleState = oci_dbmulticloud.OracleDbAzureConnectorLifecycleStateActive
	listOracleDbAzureConnectorsResponse, err := oracleDBAzureConnectorClient.ListOracleDbAzureConnectors(context.Background(), listOracleDbAzureConnectorsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting OracleDbAzureConnector list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, oracleDbAzureConnector := range listOracleDbAzureConnectorsResponse.Items {
		id := *oracleDbAzureConnector.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "OracleDbAzureConnectorId", id)
	}
	return resourceIds, nil
}

func DbmulticloudOracleDbAzureConnectorSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if oracleDbAzureConnectorResponse, ok := response.Response.(oci_dbmulticloud.GetOracleDbAzureConnectorResponse); ok {
		return oracleDbAzureConnectorResponse.LifecycleState != oci_dbmulticloud.OracleDbAzureConnectorLifecycleStateDeleted
	}
	return false
}

func DbmulticloudOracleDbAzureConnectorSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.OracleDBAzureConnectorClient().GetOracleDbAzureConnector(context.Background(), oci_dbmulticloud.GetOracleDbAzureConnectorRequest{
		OracleDbAzureConnectorId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
