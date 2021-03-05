// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v36/common"
	oci_management_agent "github.com/oracle/oci-go-sdk/v36/managementagent"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ManagementAgentResourceConfig = ManagementAgentResourceDependencies +
		generateResourceFromRepresentationMap("oci_management_agent_management_agent", "test_management_agent", Optional, Update, managementAgentRepresentation)

	managementAgentSingularDataSourceRepresentation = map[string]interface{}{
		"management_agent_id": Representation{repType: Required, create: `${oci_management_agent_management_agent.test_management_agent.id}`},
	}

	managementAgentDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"platform_type":  Representation{repType: Optional, create: `LINUX`},
		"state":          Representation{repType: Optional, create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, managementAgentDataSourceFilterRepresentation}}
	managementAgentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_management_agent_management_agent.test_management_agent.id}`}},
	}

	managementAgentRepresentation = map[string]interface{}{
		"managed_agent_id":         Representation{repType: Required, create: `${var.managed_agent_id}`},
		"display_name":             Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"is_agent_auto_upgradable": Representation{repType: Optional, create: `false`},
		"deploy_plugins_id":        Representation{repType: Optional, create: []string{`${data.oci_management_agent_management_agent_plugins.test_management_agent_plugins.management_agent_plugins.0.id}`}},
	}

	ManagementAgentResourceDependencies = ""
)

func TestManagementAgentManagementAgentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestManagementAgentManagementAgentResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	managementAgentId := getEnvSettingWithBlankDefault("managed_agent_id")
	if managementAgentId == "" {
		t.Skip("Manual install agent and set managed_agent_id to run this test")
	}
	managementAgentIdVariableStr := fmt.Sprintf("variable \"managed_agent_id\" { default = \"%s\" }\n", managementAgentId)

	resourceName := "oci_management_agent_management_agent.test_management_agent"
	datasourceName := "data.oci_management_agent_management_agents.test_management_agents"
	singularDatasourceName := "data.oci_management_agent_management_agent.test_management_agent"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+ManagementAgentResourceDependencies+
		generateResourceFromRepresentationMap("oci_management_agent_management_agent", "test_management_agent", Optional, Create, managementAgentRepresentation), "managementagent", "managementAgent", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckManagementAgentManagementAgentDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + managementAgentIdVariableStr + ManagementAgentResourceDependencies +
					generateDataSourceFromRepresentationMap("oci_management_agent_management_agent_plugins", "test_management_agent_plugins", Required, Create, managementAgentPluginDataSourceRepresentation) +
					generateResourceFromRepresentationMap("oci_management_agent_management_agent", "test_management_agent", Required, Create, managementAgentRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + managementAgentIdVariableStr + ManagementAgentResourceDependencies +
					generateDataSourceFromRepresentationMap("oci_management_agent_management_agent_plugins", "test_management_agent_plugins", Required, Create, managementAgentPluginDataSourceRepresentation) +
					generateResourceFromRepresentationMap("oci_management_agent_management_agent", "test_management_agent", Optional, Update, managementAgentRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "version"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_management_agent_management_agents", "test_management_agents", Optional, Update, managementAgentDataSourceRepresentation) +
					generateDataSourceFromRepresentationMap("oci_management_agent_management_agent_plugins", "test_management_agent_plugins", Required, Create, managementAgentPluginDataSourceRepresentation) +
					compartmentIdVariableStr + managementAgentIdVariableStr + ManagementAgentResourceDependencies +
					generateResourceFromRepresentationMap("oci_management_agent_management_agent", "test_management_agent", Optional, Update, managementAgentRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttr(datasourceName, "management_agents.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.availability_status"),
					resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.compartment_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.display_name"),
					resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.host"),
					resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.install_key_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.is_agent_auto_upgradable"),
					resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.platform_name"),
					resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.platform_type"),
					resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.platform_version"),
					resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.time_created"),
					resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.time_last_heartbeat"),
					resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.version"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_management_agent_management_agent", "test_management_agent", Required, Create, managementAgentSingularDataSourceRepresentation) +
					generateDataSourceFromRepresentationMap("oci_management_agent_management_agent_plugins", "test_management_agent_plugins", Required, Create, managementAgentPluginDataSourceRepresentation) +
					compartmentIdVariableStr + managementAgentIdVariableStr + ManagementAgentResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "platform_type", "LINUX"),
					resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "management_agent_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_status"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "host"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "install_key_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "install_path"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "is_agent_auto_upgradable"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "platform_name"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "platform_version"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_last_heartbeat"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "version"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + managementAgentIdVariableStr + ManagementAgentResourceConfig +
					generateDataSourceFromRepresentationMap("oci_management_agent_management_agent_plugins", "test_management_agent_plugins", Required, Create, managementAgentPluginDataSourceRepresentation),
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       false,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckManagementAgentManagementAgentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).managementAgentClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_management_agent_management_agent" {
			noResourceFound = false
			request := oci_management_agent.GetManagementAgentRequest{}

			tmp := rs.Primary.ID
			request.ManagementAgentId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "management_agent")

			response, err := client.GetManagementAgent(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_management_agent.LifecycleStatesTerminated): true, string(oci_management_agent.LifecycleStatesDeleted): true,
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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("ManagementAgentManagementAgent") {
		resource.AddTestSweepers("ManagementAgentManagementAgent", &resource.Sweeper{
			Name:         "ManagementAgentManagementAgent",
			Dependencies: DependencyGraph["managementAgent"],
			F:            sweepManagementAgentManagementAgentResource,
		})
	}
}

func sweepManagementAgentManagementAgentResource(compartment string) error {
	managementAgentClient := GetTestClients(&schema.ResourceData{}).managementAgentClient()
	managementAgentIds, err := getManagementAgentIds(compartment)
	if err != nil {
		return err
	}
	for _, managementAgentId := range managementAgentIds {
		if ok := SweeperDefaultResourceId[managementAgentId]; !ok {
			deleteManagementAgentRequest := oci_management_agent.DeleteManagementAgentRequest{}

			deleteManagementAgentRequest.ManagementAgentId = &managementAgentId

			deleteManagementAgentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "management_agent")
			_, error := managementAgentClient.DeleteManagementAgent(context.Background(), deleteManagementAgentRequest)
			if error != nil {
				fmt.Printf("Error deleting ManagementAgent %s %s, It is possible that the resource is already deleted. Please verify manually \n", managementAgentId, error)
				continue
			}
			waitTillCondition(testAccProvider, &managementAgentId, managementAgentSweepWaitCondition, time.Duration(3*time.Minute),
				managementAgentSweepResponseFetchOperation, "management_agent", true)
		}
	}
	return nil
}

func getManagementAgentIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "ManagementAgentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	managementAgentClient := GetTestClients(&schema.ResourceData{}).managementAgentClient()

	listManagementAgentsRequest := oci_management_agent.ListManagementAgentsRequest{}
	listManagementAgentsRequest.CompartmentId = &compartmentId
	listManagementAgentsRequest.LifecycleState = oci_management_agent.ListManagementAgentsLifecycleStateActive
	listManagementAgentsResponse, err := managementAgentClient.ListManagementAgents(context.Background(), listManagementAgentsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ManagementAgent list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, managementAgent := range listManagementAgentsResponse.Items {
		id := *managementAgent.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "ManagementAgentId", id)
	}
	return resourceIds, nil
}

func managementAgentSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if managementAgentResponse, ok := response.Response.(oci_management_agent.GetManagementAgentResponse); ok {
		return managementAgentResponse.LifecycleState != oci_management_agent.LifecycleStatesTerminated
	}
	return false
}

func managementAgentSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.managementAgentClient().GetManagementAgent(context.Background(), oci_management_agent.GetManagementAgentRequest{
		ManagementAgentId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
