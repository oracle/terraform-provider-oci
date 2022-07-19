// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_jms "github.com/oracle/oci-go-sdk/v65/jms"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	JmsFleetRequiredOnlyResource = JmsFleetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_jms_fleet", "test_fleet", acctest.Required, acctest.Create, JmsFleetRepresentation)

	JmsFleetResourceConfig = JmsFleetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_jms_fleet", "test_fleet", acctest.Optional, acctest.Update, JmsFleetRepresentation)

	JmsJmsFleetSingularDataSourceRepresentation = map[string]interface{}{
		"fleet_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_jms_fleet.test_fleet.id}`},
	}

	JmsJmsFleetDataSourceRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `Created Fleet`, Update: `displayName2`},
		"display_name_contains": acctest.Representation{RepType: acctest.Optional, Create: `displayName2`},
		"id":                    acctest.Representation{RepType: acctest.Optional, Create: `${oci_jms_fleet.test_fleet.id}`},
		"state":                 acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                acctest.RepresentationGroup{RepType: acctest.Required, Group: JmsFleetDataSourceFilterRepresentation},
	}

	JmsFleetDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_jms_fleet.test_fleet.id}`}},
	}

	JmsFleetRepresentation = map[string]interface{}{
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                 acctest.Representation{RepType: acctest.Required, Create: `Created Fleet`, Update: `displayName2`},
		"inventory_log":                acctest.RepresentationGroup{RepType: acctest.Required, Group: JmsFleetInventoryLogRepresentation},
		"operation_log":                acctest.RepresentationGroup{RepType: acctest.Optional, Group: JmsFleetOperationLogRepresentation},
		"defined_tags":                 acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                  acctest.Representation{RepType: acctest.Optional, Create: `Created Fleet`, Update: `description2`},
		"freeform_tags":                acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"is_advanced_features_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"lifecycle":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: JmsFleetIgnoreChangesRepresentation},
	}

	JmsFleetInventoryLogRepresentation = map[string]interface{}{
		"log_group_id": acctest.Representation{RepType: acctest.Required, Create: `${var.inventory_log_group_id_for_create}`, Update: `${var.inventory_log_group_id_for_update}`},
		"log_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.inventory_log_id_for_create}`, Update: `${var.inventory_log_id_for_update}`},
	}

	JmsFleetOperationLogRepresentation = map[string]interface{}{
		"log_group_id": acctest.Representation{RepType: acctest.Required, Create: `${var.operation_log_group_id_for_create}`, Update: `${var.operation_log_group_id_for_update}`},
		"log_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.operation_log_id_for_create}`, Update: `${var.operation_log_id_for_update}`},
	}

	JmsFleetIgnoreChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	JmsFleetResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: jms/default
func TestJmsFleetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsFleetResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	inventoryLogGroupId := utils.GetEnvSettingWithBlankDefault("inventory_log_group_ocid_for_create")
	inventoryLogGroupIdVariableStr := fmt.Sprintf("variable \"inventory_log_group_id_for_create\" { default = \"%s\" }\n", inventoryLogGroupId)

	inventoryLogGroupIdU := utils.GetEnvSettingWithBlankDefault("inventory_log_group_ocid_for_update")
	inventoryLogGroupIdUVariableStr := fmt.Sprintf("variable \"inventory_log_group_id_for_update\" { default = \"%s\" }\n", inventoryLogGroupIdU)

	operationLogGroupId := utils.GetEnvSettingWithBlankDefault("operation_log_group_ocid_for_create")
	operationLogGroupIdVariableStr := fmt.Sprintf("variable \"operation_log_group_id_for_create\" { default = \"%s\" }\n", operationLogGroupId)

	operationLogGroupIdU := utils.GetEnvSettingWithBlankDefault("operation_log_group_ocid_for_update")
	operationLogGroupIdUVariableStr := fmt.Sprintf("variable \"operation_log_group_id_for_update\" { default = \"%s\" }\n", operationLogGroupIdU)

	inventoryLogId := utils.GetEnvSettingWithBlankDefault("inventory_log_ocid_for_create")
	inventoryLogIdVariableStr := fmt.Sprintf("variable \"inventory_log_id_for_create\" { default = \"%s\" }\n", inventoryLogId)

	inventoryLogIdU := utils.GetEnvSettingWithBlankDefault("inventory_log_ocid_for_update")
	inventoryLogIdUVariableStr := fmt.Sprintf("variable \"inventory_log_id_for_update\" { default = \"%s\" }\n", inventoryLogIdU)

	operationLogId := utils.GetEnvSettingWithBlankDefault("operation_log_ocid_for_create")
	operationLogIdVariableStr := fmt.Sprintf("variable \"operation_log_id_for_create\" { default = \"%s\" }\n", operationLogId)

	operationLogIdU := utils.GetEnvSettingWithBlankDefault("operation_log_ocid_for_update")
	operationLogIdUVariableStr := fmt.Sprintf("variable \"operation_log_id_for_update\" { default = \"%s\" }\n", operationLogIdU)

	resourceName := "oci_jms_fleet.test_fleet"
	datasourceName := "data.oci_jms_fleets.test_fleets"
	singularDatasourceName := "data.oci_jms_fleet.test_fleet"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties.
	// This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+
		compartmentIdVariableStr+
		compartmentIdUVariableStr+
		inventoryLogGroupIdVariableStr+
		inventoryLogIdVariableStr+
		operationLogGroupIdVariableStr+
		operationLogIdVariableStr+
		inventoryLogGroupIdUVariableStr+
		inventoryLogIdUVariableStr+
		operationLogGroupIdUVariableStr+
		operationLogIdUVariableStr+
		JmsFleetResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_jms_fleet", "test_fleet", acctest.Optional, acctest.Create, JmsFleetRepresentation), "jms", "fleet", t)

	acctest.ResourceTest(t, testAccCheckJmsFleetDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config +
				compartmentIdVariableStr +
				inventoryLogGroupIdVariableStr +
				inventoryLogIdVariableStr +
				JmsFleetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_jms_fleet", "test_fleet", acctest.Required, acctest.Create, JmsFleetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Created Fleet"),
				resource.TestCheckResourceAttr(resourceName, "inventory_log.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "inventory_log.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "inventory_log.0.log_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete before next create
		{
			Config: config + compartmentIdVariableStr + JmsFleetResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config +
				compartmentIdVariableStr +
				inventoryLogGroupIdVariableStr +
				inventoryLogIdVariableStr +
				operationLogGroupIdVariableStr +
				operationLogIdVariableStr +
				JmsFleetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_jms_fleet", "test_fleet", acctest.Optional, acctest.Create, JmsFleetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "approximate_application_count"),
				resource.TestCheckResourceAttrSet(resourceName, "approximate_installation_count"),
				resource.TestCheckResourceAttrSet(resourceName, "approximate_jre_count"),
				resource.TestCheckResourceAttrSet(resourceName, "approximate_managed_instance_count"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "Created Fleet"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Created Fleet"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "inventory_log.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "inventory_log.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "inventory_log.0.log_id"),
				resource.TestCheckResourceAttr(resourceName, "is_advanced_features_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "operation_log.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "operation_log.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "operation_log.0.log_id"),
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
			Config: config +
				compartmentIdVariableStr +
				compartmentIdUVariableStr +
				inventoryLogGroupIdVariableStr +
				inventoryLogIdVariableStr +
				operationLogGroupIdVariableStr +
				operationLogIdVariableStr +
				JmsFleetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_jms_fleet", "test_fleet", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(JmsFleetRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "approximate_application_count"),
				resource.TestCheckResourceAttrSet(resourceName, "approximate_installation_count"),
				resource.TestCheckResourceAttrSet(resourceName, "approximate_jre_count"),
				resource.TestCheckResourceAttrSet(resourceName, "approximate_managed_instance_count"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "Created Fleet"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Created Fleet"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "inventory_log.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "inventory_log.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "inventory_log.0.log_id"),
				resource.TestCheckResourceAttr(resourceName, "is_advanced_features_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "operation_log.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "operation_log.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "operation_log.0.log_id"),
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
			Config: config +
				compartmentIdVariableStr +
				inventoryLogGroupIdUVariableStr +
				inventoryLogIdUVariableStr +
				operationLogGroupIdUVariableStr +
				operationLogIdUVariableStr +
				JmsFleetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_jms_fleet", "test_fleet", acctest.Optional, acctest.Update, JmsFleetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "approximate_application_count"),
				resource.TestCheckResourceAttrSet(resourceName, "approximate_installation_count"),
				resource.TestCheckResourceAttrSet(resourceName, "approximate_jre_count"),
				resource.TestCheckResourceAttrSet(resourceName, "approximate_managed_instance_count"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "inventory_log.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "inventory_log.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "inventory_log.0.log_id"),
				resource.TestCheckResourceAttr(resourceName, "is_advanced_features_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "operation_log.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "operation_log.0.log_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "operation_log.0.log_id"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_jms_fleets", "test_fleets", acctest.Optional, acctest.Update, JmsJmsFleetDataSourceRepresentation) +
				compartmentIdVariableStr +
				inventoryLogGroupIdUVariableStr +
				inventoryLogIdUVariableStr +
				operationLogGroupIdUVariableStr +
				operationLogIdUVariableStr +
				JmsFleetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_jms_fleet", "test_fleet", acctest.Optional, acctest.Update, JmsFleetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "display_name_contains", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "fleet_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "fleet_collection.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_jms_fleet", "test_fleet", acctest.Required, acctest.Create, JmsJmsFleetSingularDataSourceRepresentation) +
				compartmentIdVariableStr +
				inventoryLogGroupIdUVariableStr +
				inventoryLogIdUVariableStr +
				operationLogGroupIdUVariableStr +
				operationLogIdUVariableStr +
				JmsFleetResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fleet_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "approximate_application_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "approximate_installation_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "approximate_jre_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "approximate_managed_instance_count"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "inventory_log.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_advanced_features_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "operation_log.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config + JmsFleetRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckJmsFleetDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).JavaManagementServiceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_jms_fleet" {
			noResourceFound = false
			request := oci_jms.GetFleetRequest{}

			tmp := rs.Primary.ID
			request.FleetId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "jms")

			response, err := client.GetFleet(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_jms.LifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("JmsFleet") {
		resource.AddTestSweepers("JmsFleet", &resource.Sweeper{
			Name:         "JmsFleet",
			Dependencies: acctest.DependencyGraph["fleet"],
			F:            sweepJmsFleetResource,
		})
	}
}

func sweepJmsFleetResource(compartment string) error {
	javaManagementServiceClient := acctest.GetTestClients(&schema.ResourceData{}).JavaManagementServiceClient()
	fleetIds, err := getJmsFleetIds(compartment)
	if err != nil {
		return err
	}
	for _, fleetId := range fleetIds {
		if ok := acctest.SweeperDefaultResourceId[fleetId]; !ok {
			deleteFleetRequest := oci_jms.DeleteFleetRequest{}

			deleteFleetRequest.FleetId = &fleetId

			deleteFleetRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "jms")
			_, error := javaManagementServiceClient.DeleteFleet(context.Background(), deleteFleetRequest)
			if error != nil {
				fmt.Printf("Error deleting Fleet %s %s, It is possible that the resource is already deleted. Please verify manually \n", fleetId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &fleetId, JmsFleetSweepWaitCondition, time.Duration(3*time.Minute),
				JmsFleetSweepResponseFetchOperation, "jms", true)
		}
	}
	return nil
}

func getJmsFleetIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "FleetId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	javaManagementServiceClient := acctest.GetTestClients(&schema.ResourceData{}).JavaManagementServiceClient()

	listFleetsRequest := oci_jms.ListFleetsRequest{}
	listFleetsRequest.CompartmentId = &compartmentId
	listFleetsRequest.LifecycleState = oci_jms.ListFleetsLifecycleStateActive
	listFleetsResponse, err := javaManagementServiceClient.ListFleets(context.Background(), listFleetsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Fleet list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, fleet := range listFleetsResponse.Items {
		id := *fleet.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "FleetId", id)
	}
	return resourceIds, nil
}

func JmsFleetSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if fleetResponse, ok := response.Response.(oci_jms.GetFleetResponse); ok {
		return fleetResponse.LifecycleState != oci_jms.LifecycleStateDeleted
	}
	return false
}

func JmsFleetSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.JavaManagementServiceClient().GetFleet(context.Background(), oci_jms.GetFleetRequest{
		FleetId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
