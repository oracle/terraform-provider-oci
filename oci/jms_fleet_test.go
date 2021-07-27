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
	"github.com/oracle/oci-go-sdk/v45/common"
	oci_jms "github.com/oracle/oci-go-sdk/v45/jms"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	FleetRequiredOnlyResource = FleetResourceDependencies +
		generateResourceFromRepresentationMap("oci_jms_fleet", "test_fleet", Required, Create, fleetRepresentation)

	FleetResourceConfig = FleetResourceDependencies +
		generateResourceFromRepresentationMap("oci_jms_fleet", "test_fleet", Optional, Update, fleetRepresentation)

	fleetSingularDataSourceRepresentation = map[string]interface{}{
		"fleet_id": Representation{repType: Required, create: `${oci_jms_fleet.test_fleet.id}`},
	}

	fleetDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Optional, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `Created Fleet`, update: `displayName2`},
		"id":             Representation{repType: Optional, create: `${oci_jms_fleet.test_fleet.id}`},
		"state":          Representation{repType: Optional, create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, fleetDataSourceFilterRepresentation}}
	fleetDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_jms_fleet.test_fleet.id}`}},
	}

	fleetRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Required, create: `Created Fleet`, update: `displayName2`},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    Representation{repType: Optional, create: `Created Fleet`, update: `description2`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"bar-key": "value"}, update: map[string]string{"Department": "Accounting"}},
	}

	FleetResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: jms/default
func TestJmsFleetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsFleetResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_jms_fleet.test_fleet"
	datasourceName := "data.oci_jms_fleets.test_fleets"
	singularDatasourceName := "data.oci_jms_fleet.test_fleet"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+FleetResourceDependencies+
		generateResourceFromRepresentationMap("oci_jms_fleet", "test_fleet", Optional, Create, fleetRepresentation), "jms", "fleet", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckJmsFleetDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + FleetResourceDependencies +
					generateResourceFromRepresentationMap("oci_jms_fleet", "test_fleet", Required, Create, fleetRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "Created Fleet"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + FleetResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + FleetResourceDependencies +
					generateResourceFromRepresentationMap("oci_jms_fleet", "test_fleet", Optional, Create, fleetRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "approximate_application_count"),
					resource.TestCheckResourceAttrSet(resourceName, "approximate_installation_count"),
					resource.TestCheckResourceAttrSet(resourceName, "approximate_jre_count"),
					resource.TestCheckResourceAttrSet(resourceName, "approximate_managed_instance_count"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "Created Fleet"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "Created Fleet"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + FleetResourceDependencies +
					generateResourceFromRepresentationMap("oci_jms_fleet", "test_fleet", Optional, Create,
						representationCopyWithNewProperties(fleetRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "approximate_application_count"),
					resource.TestCheckResourceAttrSet(resourceName, "approximate_installation_count"),
					resource.TestCheckResourceAttrSet(resourceName, "approximate_jre_count"),
					resource.TestCheckResourceAttrSet(resourceName, "approximate_managed_instance_count"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "Created Fleet"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "Created Fleet"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + FleetResourceDependencies +
					generateResourceFromRepresentationMap("oci_jms_fleet", "test_fleet", Optional, Update, fleetRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "approximate_application_count"),
					resource.TestCheckResourceAttrSet(resourceName, "approximate_installation_count"),
					resource.TestCheckResourceAttrSet(resourceName, "approximate_jre_count"),
					resource.TestCheckResourceAttrSet(resourceName, "approximate_managed_instance_count"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
					generateDataSourceFromRepresentationMap("oci_jms_fleets", "test_fleets", Optional, Update, fleetDataSourceRepresentation) +
					compartmentIdVariableStr + FleetResourceDependencies +
					generateResourceFromRepresentationMap("oci_jms_fleet", "test_fleet", Optional, Update, fleetRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "id"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttr(datasourceName, "fleet_collection.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "fleet_collection.0.items.#", "1"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_jms_fleet", "test_fleet", Required, Create, fleetSingularDataSourceRepresentation) +
					compartmentIdVariableStr + FleetResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "fleet_id"),

					resource.TestCheckResourceAttrSet(singularDatasourceName, "approximate_application_count"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "approximate_installation_count"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "approximate_jre_count"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "approximate_managed_instance_count"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + FleetResourceConfig,
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckJmsFleetDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).javaManagementServiceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_jms_fleet" {
			noResourceFound = false
			request := oci_jms.GetFleetRequest{}

			tmp := rs.Primary.ID
			request.FleetId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "jms")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("JmsFleet") {
		resource.AddTestSweepers("JmsFleet", &resource.Sweeper{
			Name:         "JmsFleet",
			Dependencies: DependencyGraph["fleet"],
			F:            sweepJmsFleetResource,
		})
	}
}

func sweepJmsFleetResource(compartment string) error {
	javaManagementServiceClient := GetTestClients(&schema.ResourceData{}).javaManagementServiceClient()
	fleetIds, err := getFleetIds(compartment)
	if err != nil {
		return err
	}
	for _, fleetId := range fleetIds {
		if ok := SweeperDefaultResourceId[fleetId]; !ok {
			deleteFleetRequest := oci_jms.DeleteFleetRequest{}

			deleteFleetRequest.FleetId = &fleetId

			deleteFleetRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "jms")
			_, error := javaManagementServiceClient.DeleteFleet(context.Background(), deleteFleetRequest)
			if error != nil {
				fmt.Printf("Error deleting Fleet %s %s, It is possible that the resource is already deleted. Please verify manually \n", fleetId, error)
				continue
			}
			waitTillCondition(testAccProvider, &fleetId, fleetSweepWaitCondition, time.Duration(3*time.Minute),
				fleetSweepResponseFetchOperation, "jms", true)
		}
	}
	return nil
}

func getFleetIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "FleetId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	javaManagementServiceClient := GetTestClients(&schema.ResourceData{}).javaManagementServiceClient()

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
		addResourceIdToSweeperResourceIdMap(compartmentId, "FleetId", id)
	}
	return resourceIds, nil
}

func fleetSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if fleetResponse, ok := response.Response.(oci_jms.GetFleetResponse); ok {
		return fleetResponse.LifecycleState != oci_jms.LifecycleStateDeleted
	}
	return false
}

func fleetSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.javaManagementServiceClient().GetFleet(context.Background(), oci_jms.GetFleetRequest{
		FleetId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
