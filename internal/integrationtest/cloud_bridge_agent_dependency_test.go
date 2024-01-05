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
	oci_cloud_bridge "github.com/oracle/oci-go-sdk/v65/cloudbridge"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ocbNamespaceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	ocbNamespace      = "axzjoajp6qu6"
	ocbBucketName     = "test_bucket"
	ocbObjectName     = "test_object"
	ocbDependencyName = "VDDK"

	CloudBridgeAgentDependencyRequiredOnlyResource = CloudBridgeAgentDependencyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_agent_dependency", "test_agent_dependency", acctest.Required, acctest.Create, CloudBridgeAgentDependencyRepresentation)

	CloudBridgeAgentDependencyResourceConfig = CloudBridgeAgentDependencyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_agent_dependency", "test_agent_dependency", acctest.Optional, acctest.Update, CloudBridgeAgentDependencyRepresentation)

	CloudBridgeCloudBridgeAgentDependencySingularDataSourceRepresentation = map[string]interface{}{
		"agent_dependency_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_cloud_bridge_agent_dependency.test_agent_dependency.id}`},
	}

	CloudBridgeCloudBridgeAgentDependencyDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: CloudBridgeAgentDependencyDataSourceFilterRepresentation}}
	CloudBridgeAgentDependencyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_cloud_bridge_agent_dependency.test_agent_dependency.id}`}},
	}

	CloudBridgeAgentDependencyRepresentation = map[string]interface{}{
		"bucket":          acctest.Representation{RepType: acctest.Required, Create: ocbBucketName, Update: ocbBucketName},
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"dependency_name": acctest.Representation{RepType: acctest.Required, Create: ocbDependencyName, Update: ocbDependencyName},
		"display_name":    acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName`},
		"namespace":       acctest.Representation{RepType: acctest.Required, Create: ocbNamespace, Update: ocbNamespace},
		"object":          acctest.Representation{RepType: acctest.Required, Create: ocbObjectName, Update: ocbObjectName},
		"lifecycle":       acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreSystemTagsChangesRep},
	}

	ignoreSystemTagsChangesRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`system_tags`}},
	}

	CloudBridgeAgentDependencyResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_agent", "test_agent", acctest.Required, acctest.Create, CloudBridgeAgentRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_environment", "test_environment", acctest.Required, acctest.Create, CloudBridgeEnvironmentRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "tenancy_namespace", acctest.Required, acctest.Create, ocbNamespaceRepresentation)
)

// issue-routing-tag: cloud_bridge/default
func TestCloudBridgeAgentDependencyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCloudBridgeAgentDependencyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_cloud_bridge_agent_dependency.test_agent_dependency"
	datasourceName := "data.oci_cloud_bridge_agent_dependencies.test_agent_dependencies"
	singularDatasourceName := "data.oci_cloud_bridge_agent_dependency.test_agent_dependency"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CloudBridgeAgentDependencyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_agent_dependency", "test_agent_dependency", acctest.Optional, acctest.Create, CloudBridgeAgentDependencyRepresentation), "cloudbridge", "agentDependency", t)

	acctest.ResourceTest(t, testAccCheckCloudBridgeAgentDependencyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CloudBridgeAgentDependencyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_agent_dependency", "test_agent_dependency", acctest.Required, acctest.Create, CloudBridgeAgentDependencyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "bucket", ocbBucketName),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "dependency_name", ocbDependencyName),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "namespace", ocbNamespace),
				resource.TestCheckResourceAttr(resourceName, "object", ocbObjectName),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CloudBridgeAgentDependencyResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CloudBridgeAgentDependencyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_agent_dependency", "test_agent_dependency", acctest.Optional, acctest.Create, CloudBridgeAgentDependencyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "bucket", ocbBucketName),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "dependency_name", ocbDependencyName),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "namespace", ocbNamespace),
				resource.TestCheckResourceAttr(resourceName, "object", ocbObjectName),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + CloudBridgeAgentDependencyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_agent_dependency", "test_agent_dependency", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CloudBridgeAgentDependencyRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "bucket", ocbBucketName),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "dependency_name", ocbDependencyName),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "namespace", ocbNamespace),
				resource.TestCheckResourceAttr(resourceName, "object", ocbObjectName),

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
			Config: config + compartmentIdVariableStr + CloudBridgeAgentDependencyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_agent_dependency", "test_agent_dependency", acctest.Optional, acctest.Update, CloudBridgeAgentDependencyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "bucket", ocbBucketName),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "dependency_name", ocbDependencyName),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "object", ocbObjectName),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_bridge_agent_dependencies", "test_agent_dependencies", acctest.Optional, acctest.Update, CloudBridgeCloudBridgeAgentDependencyDataSourceRepresentation) +
				compartmentIdVariableStr + CloudBridgeAgentDependencyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_agent_dependency", "test_agent_dependency", acctest.Optional, acctest.Update, CloudBridgeAgentDependencyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(datasourceName, "agent_dependency_collection.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_bridge_agent_dependency", "test_agent_dependency", acctest.Required, acctest.Create, CloudBridgeCloudBridgeAgentDependencySingularDataSourceRepresentation) +
				compartmentIdVariableStr + CloudBridgeAgentDependencyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "agent_dependency_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "bucket", ocbBucketName),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "checksum"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "dependency_name", ocbDependencyName),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "e_tag"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "namespace", ocbNamespace),
				resource.TestCheckResourceAttr(singularDatasourceName, "object", ocbObjectName),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config + CloudBridgeAgentDependencyRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCloudBridgeAgentDependencyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).OcbAgentSvcClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_cloud_bridge_agent_dependency" {
			noResourceFound = false
			request := oci_cloud_bridge.GetAgentDependencyRequest{}

			tmp := rs.Primary.ID
			request.AgentDependencyId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cloud_bridge")

			response, err := client.GetAgentDependency(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_cloud_bridge.AgentDependencyLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("CloudBridgeAgentDependency") {
		resource.AddTestSweepers("CloudBridgeAgentDependency", &resource.Sweeper{
			Name:         "CloudBridgeAgentDependency",
			Dependencies: acctest.DependencyGraph["agentDependency"],
			F:            sweepCloudBridgeAgentDependencyResource,
		})
	}
}

func sweepCloudBridgeAgentDependencyResource(compartment string) error {
	ocbAgentSvcClient := acctest.GetTestClients(&schema.ResourceData{}).OcbAgentSvcClient()
	agentDependencyIds, err := getCloudBridgeAgentDependencyIds(compartment)
	if err != nil {
		return err
	}
	for _, agentDependencyId := range agentDependencyIds {
		if ok := acctest.SweeperDefaultResourceId[agentDependencyId]; !ok {
			deleteAgentDependencyRequest := oci_cloud_bridge.DeleteAgentDependencyRequest{}

			deleteAgentDependencyRequest.AgentDependencyId = &agentDependencyId

			deleteAgentDependencyRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cloud_bridge")
			_, error := ocbAgentSvcClient.DeleteAgentDependency(context.Background(), deleteAgentDependencyRequest)
			if error != nil {
				fmt.Printf("Error deleting AgentDependency %s %s, It is possible that the resource is already deleted. Please verify manually \n", agentDependencyId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &agentDependencyId, CloudBridgeAgentDependencySweepWaitCondition, time.Duration(3*time.Minute),
				CloudBridgeAgentDependencySweepResponseFetchOperation, "cloud_bridge", true)
		}
	}
	return nil
}

func getCloudBridgeAgentDependencyIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "AgentDependencyId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	ocbAgentSvcClient := acctest.GetTestClients(&schema.ResourceData{}).OcbAgentSvcClient()

	listAgentDependenciesRequest := oci_cloud_bridge.ListAgentDependenciesRequest{}
	listAgentDependenciesRequest.CompartmentId = &compartmentId
	listAgentDependenciesRequest.LifecycleState = oci_cloud_bridge.AgentDependencyLifecycleStateActive
	listAgentDependenciesResponse, err := ocbAgentSvcClient.ListAgentDependencies(context.Background(), listAgentDependenciesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting AgentDependency list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, agentDependency := range listAgentDependenciesResponse.Items {
		id := *agentDependency.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "AgentDependencyId", id)
	}
	return resourceIds, nil
}

func CloudBridgeAgentDependencySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if agentDependencyResponse, ok := response.Response.(oci_cloud_bridge.GetAgentDependencyResponse); ok {
		return agentDependencyResponse.LifecycleState != oci_cloud_bridge.AgentDependencyLifecycleStateDeleted
	}
	return false
}

func CloudBridgeAgentDependencySweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.OcbAgentSvcClient().GetAgentDependency(context.Background(), oci_cloud_bridge.GetAgentDependencyRequest{
		AgentDependencyId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
