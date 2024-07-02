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
	oci_management_agent "github.com/oracle/oci-go-sdk/v65/managementagent"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	managementAgentSingularDataSourceRepresentation = map[string]interface{}{
		"management_agent_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_management_agent_management_agent.test_management_agent.id}`},
	}

	managementAgentDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":              acctest.Representation{RepType: acctest.Optional, Create: `ACCESSIBLE`},
		"availability_status":       acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `terraformTest`, Update: `terraformTest2`},
		"host_id":                   acctest.Representation{RepType: acctest.Optional, Create: ``},
		"install_type":              acctest.Representation{RepType: acctest.Optional, Create: `AGENT`},
		"is_customer_deployed":      acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"platform_type":             acctest.Representation{RepType: acctest.Optional, Create: []string{`LINUX`}},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: managementAgentDataSourceFilterRepresentation}}

	managementAgentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_management_agent_management_agent.test_management_agent.id}`}},
	}

	ManagementAgentResourceDependencies = ""
)

// issue-routing-tag: management_agent/default
// This test can only be run against production where RQS scope is not in staging mode. this will fail in dev environments
func TestManagementAgentManagementAgentResource_dataInSubcompartment(t *testing.T) {
	httpreplay.SetScenario("TestManagementAgentManagementAgentResource_dataInSubcompartment")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	rootCompartmentId := utils.GetEnvSettingWithBlankDefault("root_compartment_ocid")
	if rootCompartmentId == "" {
		rootCompartmentId = utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	}
	rootCompartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", rootCompartmentId)
	managementAgentDataSourceRepresentationInSubtree := map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: rootCompartmentId},
		"access_level":              acctest.Representation{RepType: acctest.Optional, Create: `ACCESSIBLE`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `terraformTest`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `true`}}
	datasourceName := "data.oci_management_agent_management_agents.test_management_agents"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource with compartment_id_in_subtree
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_management_agent_management_agents", "test_management_agents", acctest.Optional, acctest.Update, managementAgentDataSourceRepresentationInSubtree) +
				rootCompartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", rootCompartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.availability_status"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.host"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.install_key_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.install_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.is_customer_deployed"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.is_agent_auto_upgradable"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.platform_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.platform_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.platform_version"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.time_last_heartbeat"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.version"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.time_updated"),
			),
		},
	})
}
func TestManagementAgentManagementAgentResource_gatewayId(t *testing.T) {
	// This test searches for a specific agent (terraformGWTest) which is known to have a gatewayId set
	// and validates the management_agent_properties are set for that agent
	httpreplay.SetScenario("TestManagementAgentManagementAgentResource_gatewayId")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	rootCompartmentId := utils.GetEnvSettingWithBlankDefault("root_compartment_ocid")
	if rootCompartmentId == "" {
		rootCompartmentId = utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	}
	rootCompartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", rootCompartmentId)
	managementAgentDataSourceRepresentationWithGW := map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: rootCompartmentId},
		"access_level":              acctest.Representation{RepType: acctest.Optional, Create: `ACCESSIBLE`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `terraformGWTest`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `true`}}
	datasourceName := "data.oci_management_agent_management_agents.test_management_agents"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource with compartment_id_in_subtree
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_management_agent_management_agents", "test_management_agents", acctest.Optional, acctest.Update, managementAgentDataSourceRepresentationWithGW) +
				rootCompartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", rootCompartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.availability_status"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.host"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.install_key_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.install_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.is_customer_deployed"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.is_agent_auto_upgradable"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.platform_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.platform_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.platform_version"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.time_last_heartbeat"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.version"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.time_updated"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.management_agent_properties.#"),
			),
		},
	})
}

// issue-routing-tag: management_agent/default
func TestManagementAgentManagementAgentResource_basic(t *testing.T) {

	//  Basic test performs the following
	//  1. List all agents in compartment (env:TF_VAR_compartment_ocid) with status=ACTIVE and displayName=terraformTest
	//  2.  Verifies the first response contains the values defined managementAgentDataSourceRepresentation above
	//
	httpreplay.SetScenario("TestManagementAgentManagementAgentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	//  1. List all agents in compartment (env:TF_VAR_compartment_ocid) with status=ACTIVE and displayName=terraformTest
	managementAgentIds, err := getManagementAgentIds(compartmentId)
	if err != nil {
		t.Errorf("Failed to get agents in compartment %s", err)
	}
	if len(managementAgentIds) == 0 {
		t.Errorf("Failed to find any active agents in compartment %s", compartmentId)
	}
	managementAgentId := managementAgentIds[0]
	managementAgentIdVariableStr := fmt.Sprintf("variable \"managed_agent_id\" { default = \"%s\" }\n", managementAgentId)

	resourceName := "oci_management_agent_management_agent.test_management_agent"
	datasourceName := "data.oci_management_agent_management_agents.test_management_agents"
	singularDatasourceName := "data.oci_management_agent_management_agent.test_management_agent"
	managementAgentRepresentation := map[string]interface{}{
		"managed_agent_id":  acctest.Representation{RepType: acctest.Required, Create: managementAgentId},
		"display_name":      acctest.Representation{RepType: acctest.Optional, Create: `terraformTest`, Update: `terraformTest2`},
		"deploy_plugins_id": acctest.Representation{RepType: acctest.Optional, Create: []string{`${data.oci_management_agent_management_agent_plugins.test_management_agent_plugins.management_agent_plugins.0.id}`}},
		"freeform_tags":     acctest.Representation{RepType: acctest.Required, Create: map[string]string{"TestingTag": "TestingValue"}, Update: map[string]string{"TestingTag2": "TestingValue2"}},
	}
	ManagementAgentRequiredOnlyResource := acctest.GenerateResourceFromRepresentationMap("oci_management_agent_management_agent", "test_management_agent", acctest.Required, acctest.Create, managementAgentRepresentation)

	ManagementAgentResourceConfig := ManagementAgentResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_management_agent_management_agent", "test_management_agent", acctest.Optional, acctest.Update, managementAgentRepresentation)

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ManagementAgentResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_management_agent_management_agent", "test_management_agent", acctest.Optional, acctest.Create, managementAgentRepresentation)+managementAgentIdVariableStr, "managementagent", "managementAgent", t)

	acctest.ResourceTest(t, testAccCheckManagementAgentManagementAgentDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + managementAgentIdVariableStr + ManagementAgentResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_management_agent_management_agent_plugins",
					"test_management_agent_plugins", acctest.Required, acctest.Create, managementAgentPluginDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_management_agent_management_agent",
					"test_management_agent", acctest.Required, acctest.Create, managementAgentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),

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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + managementAgentIdVariableStr + ManagementAgentResourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_management_agent_management_agent_plugins", "test_management_agent_plugins", acctest.Required, acctest.Create, managementAgentPluginDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_management_agent_management_agent", "test_management_agent", acctest.Optional, acctest.Update, managementAgentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "version"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "terraformTest2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_management_agent_management_agents", "test_management_agents", acctest.Optional, acctest.Update, managementAgentDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_management_agent_management_agent_plugins", "test_management_agent_plugins", acctest.Required, acctest.Create, managementAgentPluginDataSourceRepresentation) +
				compartmentIdVariableStr + managementAgentIdVariableStr + ManagementAgentResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_management_agent_management_agent", "test_management_agent", acctest.Optional, acctest.Update, managementAgentRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "availability_status", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "terraformTest2"),
				resource.TestCheckResourceAttr(datasourceName, "install_type", "AGENT"),
				resource.TestCheckResourceAttr(datasourceName, "is_customer_deployed", "true"),
				resource.TestCheckResourceAttr(datasourceName, "platform_type.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "management_agents.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.availability_status"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.host"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.install_key_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.install_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.is_customer_deployed"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.is_agent_auto_upgradable"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.platform_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.platform_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.platform_version"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.time_last_heartbeat"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.version"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.time_updated"),
				resource.TestCheckResourceAttr(datasourceName, "management_agents.0.plugin_list.0.plugin_status", "RUNNING"),
				resource.TestCheckResourceAttrSet(datasourceName, "management_agents.0.plugin_list.0.plugin_display_name"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_management_agent_management_agent", "test_management_agent", acctest.Required, acctest.Create, managementAgentSingularDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_management_agent_management_agent_plugins", "test_management_agent_plugins", acctest.Required, acctest.Create, managementAgentPluginDataSourceRepresentation) +
				compartmentIdVariableStr + managementAgentIdVariableStr + ManagementAgentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "terraformTest2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "platform_type", "LINUX"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "management_agent_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "data_source_list.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "host"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "install_key_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "install_path"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "install_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_agent_auto_upgradable"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_customer_deployed"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "platform_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "platform_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_last_heartbeat"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "version"),
			),
		},
		// verify resource import
		{
			Config:                  config + ManagementAgentRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       false,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckManagementAgentManagementAgentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ManagementAgentClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_management_agent_management_agent" {
			noResourceFound = false
			request := oci_management_agent.GetManagementAgentRequest{}

			tmp := rs.Primary.ID
			request.ManagementAgentId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "management_agent")

			_, err := client.GetManagementAgent(context.Background(), request)

			//Destroy for resource does not delete the resource from MACS backend, it only removes it from terraform state file.
			//resources should be removed manually by Uninstalling the agent from the on premise system or OCA compute node

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
	if !acctest.InSweeperExcludeList("ManagementAgentManagementAgent") {
		resource.AddTestSweepers("ManagementAgentManagementAgent", &resource.Sweeper{
			Name:         "ManagementAgentManagementAgent",
			Dependencies: acctest.DependencyGraph["managementAgent"],
			F:            sweepManagementAgentManagementAgentResource,
		})
	}
}

func sweepManagementAgentManagementAgentResource(compartment string) error {
	managementAgentClient := acctest.GetTestClients(&schema.ResourceData{}).ManagementAgentClient()
	managementAgentIds, err := getManagementAgentIds(compartment)
	if err != nil {
		return err
	}
	for _, managementAgentId := range managementAgentIds {
		if ok := acctest.SweeperDefaultResourceId[managementAgentId]; !ok {
			deleteManagementAgentRequest := oci_management_agent.DeleteManagementAgentRequest{}

			deleteManagementAgentRequest.ManagementAgentId = &managementAgentId

			deleteManagementAgentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "management_agent")
			_, error := managementAgentClient.DeleteManagementAgent(context.Background(), deleteManagementAgentRequest)
			if error != nil {
				fmt.Printf("Error deleting ManagementAgent %s %s, It is possible that the resource is already deleted. Please verify manually \n", managementAgentId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &managementAgentId, managementAgentSweepWaitCondition, time.Duration(3*time.Minute),
				managementAgentSweepResponseFetchOperation, "management_agent", true)
		}
	}
	return nil
}

func getManagementAgentIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ManagementAgentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	terraformTest := "terraformTest"
	managementAgentClient := acctest.GetTestClients(&schema.ResourceData{}).ManagementAgentClient()

	listManagementAgentsRequest := oci_management_agent.ListManagementAgentsRequest{}
	listManagementAgentsRequest.CompartmentId = &compartmentId
	listManagementAgentsRequest.LifecycleState = oci_management_agent.ListManagementAgentsLifecycleStateActive
	listManagementAgentsRequest.DisplayName = &terraformTest
	listManagementAgentsResponse, err := managementAgentClient.ListManagementAgents(context.Background(), listManagementAgentsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ManagementAgent list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, managementAgent := range listManagementAgentsResponse.Items {
		id := *managementAgent.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ManagementAgentId", id)
	}
	return resourceIds, nil
}

func managementAgentSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if managementAgentResponse, ok := response.Response.(oci_management_agent.GetManagementAgentResponse); ok {
		return managementAgentResponse.LifecycleState != oci_management_agent.LifecycleStatesTerminated && managementAgentResponse.LifecycleState != oci_management_agent.LifecycleStatesDeleted
	}
	return false
}

func managementAgentSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ManagementAgentClient().GetManagementAgent(context.Background(), oci_management_agent.GetManagementAgentRequest{
		ManagementAgentId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
