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
	oci_cloud_bridge "github.com/oracle/oci-go-sdk/v65/cloudbridge"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

//fake

var (
	CloudBridgeEnvironmentRequiredOnlyResource = CloudBridgeEnvironmentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_environment", "test_environment", acctest.Required, acctest.Create, CloudBridgeEnvironmentRepresentation)

	CloudBridgeEnvironmentResourceConfig = CloudBridgeEnvironmentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_environment", "test_environment", acctest.Optional, acctest.Update, CloudBridgeEnvironmentRepresentation)

	CloudBridgeCloudBridgeEnvironmentSingularDataSourceRepresentation = map[string]interface{}{
		"environment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_cloud_bridge_environment.test_environment.id}`},
	}

	CloudBridgeCloudBridgeEnvironmentDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"environment_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_cloud_bridge_environment.test_environment.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: CloudBridgeEnvironmentDataSourceFilterRepresentation}}
	CloudBridgeEnvironmentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_cloud_bridge_environment.test_environment.id}`}},
	}

	CloudBridgeEnvironmentRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreSystemTagsChangesRep},
	}

	CloudBridgeEnvironmentResourceDependencies = ""
)

// issue-routing-tag: cloud_bridge/default
func TestCloudBridgeEnvironmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCloudBridgeEnvironmentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_cloud_bridge_environment.test_environment"
	datasourceName := "data.oci_cloud_bridge_environments.test_environments"
	singularDatasourceName := "data.oci_cloud_bridge_environment.test_environment"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CloudBridgeEnvironmentResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_environment", "test_environment", acctest.Optional, acctest.Create, CloudBridgeEnvironmentRepresentation), "cloudbridge", "environment", t)

	acctest.ResourceTest(t, testAccCheckCloudBridgeEnvironmentDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CloudBridgeEnvironmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_environment", "test_environment", acctest.Required, acctest.Create, CloudBridgeEnvironmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CloudBridgeEnvironmentResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CloudBridgeEnvironmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_environment", "test_environment", acctest.Optional, acctest.Create, CloudBridgeEnvironmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CloudBridgeEnvironmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_environment", "test_environment", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CloudBridgeEnvironmentRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
			Config: config + compartmentIdVariableStr + CloudBridgeEnvironmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_environment", "test_environment", acctest.Optional, acctest.Update, CloudBridgeEnvironmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_bridge_environments", "test_environments", acctest.Optional, acctest.Update, CloudBridgeCloudBridgeEnvironmentDataSourceRepresentation) +
				compartmentIdVariableStr + CloudBridgeEnvironmentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_environment", "test_environment", acctest.Optional, acctest.Update, CloudBridgeEnvironmentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "environment_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "environment_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "environment_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_bridge_environment", "test_environment", acctest.Required, acctest.Create, CloudBridgeCloudBridgeEnvironmentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CloudBridgeEnvironmentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "environment_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + CloudBridgeEnvironmentRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCloudBridgeEnvironmentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).OcbAgentSvcClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_cloud_bridge_environment" {
			noResourceFound = false
			request := oci_cloud_bridge.GetEnvironmentRequest{}

			tmp := rs.Primary.ID
			request.EnvironmentId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cloud_bridge")

			response, err := client.GetEnvironment(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_cloud_bridge.EnvironmentLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("CloudBridgeEnvironment") {
		resource.AddTestSweepers("CloudBridgeEnvironment", &resource.Sweeper{
			Name:         "CloudBridgeEnvironment",
			Dependencies: acctest.DependencyGraph["environment"],
			F:            sweepCloudBridgeEnvironmentResource,
		})
	}
}

func sweepCloudBridgeEnvironmentResource(compartment string) error {
	ocbAgentSvcClient := acctest.GetTestClients(&schema.ResourceData{}).OcbAgentSvcClient()
	environmentIds, err := getCloudBridgeEnvironmentIds(compartment)
	if err != nil {
		return err
	}
	for _, environmentId := range environmentIds {
		if ok := acctest.SweeperDefaultResourceId[environmentId]; !ok {
			deleteEnvironmentRequest := oci_cloud_bridge.DeleteEnvironmentRequest{}

			deleteEnvironmentRequest.EnvironmentId = &environmentId

			deleteEnvironmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cloud_bridge")
			_, error := ocbAgentSvcClient.DeleteEnvironment(context.Background(), deleteEnvironmentRequest)
			if error != nil {
				fmt.Printf("Error deleting Environment %s %s, It is possible that the resource is already deleted. Please verify manually \n", environmentId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &environmentId, CloudBridgeEnvironmentSweepWaitCondition, time.Duration(3*time.Minute),
				CloudBridgeEnvironmentSweepResponseFetchOperation, "cloud_bridge", true)
		}
	}
	return nil
}

func getCloudBridgeEnvironmentIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "EnvironmentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	ocbAgentSvcClient := acctest.GetTestClients(&schema.ResourceData{}).OcbAgentSvcClient()

	listEnvironmentsRequest := oci_cloud_bridge.ListEnvironmentsRequest{}
	listEnvironmentsRequest.CompartmentId = &compartmentId
	listEnvironmentsRequest.LifecycleState = oci_cloud_bridge.EnvironmentLifecycleStateActive
	listEnvironmentsResponse, err := ocbAgentSvcClient.ListEnvironments(context.Background(), listEnvironmentsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Environment list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, environment := range listEnvironmentsResponse.Items {
		id := *environment.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "EnvironmentId", id)
	}
	return resourceIds, nil
}

func CloudBridgeEnvironmentSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if environmentResponse, ok := response.Response.(oci_cloud_bridge.GetEnvironmentResponse); ok {
		return environmentResponse.LifecycleState != oci_cloud_bridge.EnvironmentLifecycleStateDeleted
	}
	return false
}

func CloudBridgeEnvironmentSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.OcbAgentSvcClient().GetEnvironment(context.Background(), oci_cloud_bridge.GetEnvironmentRequest{
		EnvironmentId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
