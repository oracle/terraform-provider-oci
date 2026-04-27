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
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database_tools "github.com/oracle/oci-go-sdk/v65/databasetools"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigRequiredOnlyResource = DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_database_api_gateway_config", "test_database_tools_database_api_gateway_config", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigRepresentation)

	DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigResourceConfig = DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_database_api_gateway_config", "test_database_tools_database_api_gateway_config", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigRepresentation)

	DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigSingularDataSourceRepresentation = map[string]interface{}{
		"database_tools_database_api_gateway_config_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_tools_database_tools_database_api_gateway_config.test_database_tools_database_api_gateway_config.id}`},
	}

	DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `MyDbApiConfig1`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"type":           acctest.Representation{RepType: acctest.Optional, Create: []string{`DEFAULT`}},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigDataSourceFilterRepresentation}}
	DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_tools_database_tools_database_api_gateway_config.test_database_tools_database_api_gateway_config.id}`}},
	}

	DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":    acctest.Representation{RepType: acctest.Required, Create: `MyDbApiConfig1`, Update: `displayName2`},
		"metadata_source": acctest.Representation{RepType: acctest.Required, Create: `DATABASE`},
		"type":            acctest.Representation{RepType: acctest.Required, Create: `DEFAULT`},
		"freeform_tags":   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
	}

	DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigResourceDependencies = ""
)

// issue-routing-tag: database_tools/default
func TestDatabaseToolsDatabaseToolsDatabaseApiGatewayConfigResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseToolsDatabaseToolsDatabaseApiGatewayConfigResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	relatedResourceId := utils.GetEnvSettingWithBlankDefault("TF_VAR_related_resource_id")
	relatedResourceIdVariableStr := fmt.Sprintf("variable \"related_resource_id\" { default = \"%s\" }\n", relatedResourceId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_tools_database_tools_database_api_gateway_config.test_database_tools_database_api_gateway_config"
	datasourceName := "data.oci_database_tools_database_tools_database_api_gateway_configs.test_database_tools_database_api_gateway_configs"
	singularDatasourceName := "data.oci_database_tools_database_tools_database_api_gateway_config.test_database_tools_database_api_gateway_config"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+relatedResourceIdVariableStr+DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_database_api_gateway_config", "test_database_tools_database_api_gateway_config", acctest.Optional, acctest.Create, DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigRepresentation), "databasetools", "databaseToolsDatabaseApiGatewayConfig", t)

	acctest.ResourceTest(t, testAccCheckDatabaseToolsDatabaseToolsDatabaseApiGatewayConfigDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + relatedResourceIdVariableStr + DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_database_api_gateway_config", "test_database_tools_database_api_gateway_config", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MyDbApiConfig1"),
				resource.TestCheckResourceAttr(resourceName, "metadata_source", "DATABASE"),
				resource.TestCheckResourceAttr(resourceName, "type", "DEFAULT"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + relatedResourceIdVariableStr + DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + relatedResourceIdVariableStr + DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_database_api_gateway_config", "test_database_tools_database_api_gateway_config", acctest.Optional, acctest.Create, DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MyDbApiConfig1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "metadata_source", "DATABASE"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "type", "DEFAULT"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + relatedResourceIdVariableStr + DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_database_api_gateway_config", "test_database_tools_database_api_gateway_config", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MyDbApiConfig1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "metadata_source", "DATABASE"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "type", "DEFAULT"),

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
			Config: config + compartmentIdVariableStr + relatedResourceIdVariableStr + DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_database_api_gateway_config", "test_database_tools_database_api_gateway_config", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "metadata_source", "DATABASE"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "type", "DEFAULT"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_database_api_gateway_configs", "test_database_tools_database_api_gateway_configs", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigDataSourceRepresentation) +
				compartmentIdVariableStr + relatedResourceIdVariableStr + DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_database_api_gateway_config", "test_database_tools_database_api_gateway_config", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "type.#", "1"),

				resource.TestCheckResourceAttr(datasourceName, "database_tools_database_api_gateway_config_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "database_tools_database_api_gateway_config_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_database_api_gateway_config", "test_database_tools_database_api_gateway_config", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigSingularDataSourceRepresentation) +
				compartmentIdVariableStr + relatedResourceIdVariableStr + DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_tools_database_api_gateway_config_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metadata_source", "DATABASE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "DEFAULT"),
			),
		},
		// verify resource import
		{
			Config:                  config + compartmentIdVariableStr + relatedResourceIdVariableStr + DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDatabaseToolsDatabaseToolsDatabaseApiGatewayConfigDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DatabaseToolsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_tools_database_tools_database_api_gateway_config" {
			noResourceFound = false
			request := oci_database_tools.GetDatabaseToolsDatabaseApiGatewayConfigRequest{}

			tmp := rs.Primary.ID
			request.DatabaseToolsDatabaseApiGatewayConfigId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_tools")

			response, err := client.GetDatabaseToolsDatabaseApiGatewayConfig(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database_tools.DatabaseToolsDatabaseApiGatewayConfigLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.GetLifecycleState())]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.GetLifecycleState())
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
	if !acctest.InSweeperExcludeList("DatabaseToolsDatabaseToolsDatabaseApiGatewayConfig") {
		resource.AddTestSweepers("DatabaseToolsDatabaseToolsDatabaseApiGatewayConfig", &resource.Sweeper{
			Name:         "DatabaseToolsDatabaseToolsDatabaseApiGatewayConfig",
			Dependencies: acctest.DependencyGraph["databaseToolsDatabaseApiGatewayConfig"],
			F:            sweepDatabaseToolsDatabaseToolsDatabaseApiGatewayConfigResource,
		})
	}
}

func sweepDatabaseToolsDatabaseToolsDatabaseApiGatewayConfigResource(compartment string) error {
	databaseToolsClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseToolsClient()
	databaseToolsDatabaseApiGatewayConfigIds, err := getDatabaseToolsDatabaseToolsDatabaseApiGatewayConfigIds(compartment)
	if err != nil {
		return err
	}
	for _, databaseToolsDatabaseApiGatewayConfigId := range databaseToolsDatabaseApiGatewayConfigIds {
		if ok := acctest.SweeperDefaultResourceId[databaseToolsDatabaseApiGatewayConfigId]; !ok {
			deleteDatabaseToolsDatabaseApiGatewayConfigRequest := oci_database_tools.DeleteDatabaseToolsDatabaseApiGatewayConfigRequest{}

			deleteDatabaseToolsDatabaseApiGatewayConfigRequest.DatabaseToolsDatabaseApiGatewayConfigId = &databaseToolsDatabaseApiGatewayConfigId

			deleteDatabaseToolsDatabaseApiGatewayConfigRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_tools")
			_, error := databaseToolsClient.DeleteDatabaseToolsDatabaseApiGatewayConfig(context.Background(), deleteDatabaseToolsDatabaseApiGatewayConfigRequest)
			if error != nil {
				fmt.Printf("Error deleting DatabaseToolsDatabaseApiGatewayConfig %s %s, It is possible that the resource is already deleted. Please verify manually \n", databaseToolsDatabaseApiGatewayConfigId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &databaseToolsDatabaseApiGatewayConfigId, DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigSweepResponseFetchOperation, "database_tools", true)
		}
	}
	return nil
}

func getDatabaseToolsDatabaseToolsDatabaseApiGatewayConfigIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DatabaseToolsDatabaseApiGatewayConfigId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseToolsClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseToolsClient()

	listDatabaseToolsDatabaseApiGatewayConfigsRequest := oci_database_tools.ListDatabaseToolsDatabaseApiGatewayConfigsRequest{}
	listDatabaseToolsDatabaseApiGatewayConfigsRequest.CompartmentId = &compartmentId
	listDatabaseToolsDatabaseApiGatewayConfigsRequest.LifecycleState = oci_database_tools.ListDatabaseToolsDatabaseApiGatewayConfigsLifecycleStateActive
	listDatabaseToolsDatabaseApiGatewayConfigsResponse, err := databaseToolsClient.ListDatabaseToolsDatabaseApiGatewayConfigs(context.Background(), listDatabaseToolsDatabaseApiGatewayConfigsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DatabaseToolsDatabaseApiGatewayConfig list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, databaseToolsDatabaseApiGatewayConfig := range listDatabaseToolsDatabaseApiGatewayConfigsResponse.Items {
		id := *databaseToolsDatabaseApiGatewayConfig.GetId()
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DatabaseToolsDatabaseApiGatewayConfigId", id)
	}
	return resourceIds, nil
}

func DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if databaseToolsDatabaseApiGatewayConfigResponse, ok := response.Response.(oci_database_tools.GetDatabaseToolsDatabaseApiGatewayConfigResponse); ok {
		return databaseToolsDatabaseApiGatewayConfigResponse.GetLifecycleState() != oci_database_tools.DatabaseToolsDatabaseApiGatewayConfigLifecycleStateDeleted
	}
	return false
}

func DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseToolsClient().GetDatabaseToolsDatabaseApiGatewayConfig(context.Background(), oci_database_tools.GetDatabaseToolsDatabaseApiGatewayConfigRequest{
		DatabaseToolsDatabaseApiGatewayConfigId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
