// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OsManagementHubEventRequiredOnlyResource = OsManagementHubEventResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_event", "test_event", acctest.Required, acctest.Create, OsManagementHubEventRepresentation)

	//OsManagementHubEventResourceConfig = OsManagementHubEventResourceDependencies +
	//	acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_event", "test_event", acctest.Optional, acctest.Update, OsManagementHubEventRepresentation)

	OsManagementHubEventSingularDataSourceRepresentation = map[string]interface{}{
		"event_id": acctest.Representation{RepType: acctest.Required, Create: `${var.event_id2}`},
	}

	OsManagementHubEventDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		//"event_fingerprint":              acctest.Representation{RepType: acctest.Optional, Create: `eventFingerprint`},
		"event_summary":                         acctest.Representation{RepType: acctest.Optional, Create: `Manually created event 2 for testing caused by <Yijiu>`},
		"event_summary_contains":                acctest.Representation{RepType: acctest.Optional, Create: `testing`},
		"id":                                    acctest.Representation{RepType: acctest.Optional, Create: `${var.event_id2}`},
		"is_managed_by_autonomous_linux":        acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"resource_id":                           acctest.Representation{RepType: acctest.Optional, Create: utils.GetEnvSettingWithBlankDefault("osmh_managed_instance_ocid")},
		"state":                                 acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"time_created_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `2018-01-01T00:00:00.000Z`},
		"time_created_less_than":                acctest.Representation{RepType: acctest.Optional, Create: `2088-01-01T00:00:00.000Z`},
		"type":                                  acctest.Representation{RepType: acctest.Optional, Create: []string{`EXPLOIT_ATTEMPT`}},
		"filter":                                acctest.RepresentationGroup{RepType: acctest.Required, Group: OsManagementHubEventDataSourceFilterRepresentation}}
	OsManagementHubEventDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${var.event_id2}`}},
	}

	OsManagementHubEventRepresentation = map[string]interface{}{
		"event_id": acctest.Representation{RepType: acctest.Required, Create: `${var.event_id}`},
		//"defined_tags":  acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	OsManagementHubEventRepresentation2 = map[string]interface{}{
		"event_id": acctest.Representation{RepType: acctest.Required, Create: `${var.event_id2}`},
		//"defined_tags":  acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	OsManagementHubEventResourceDependencies = ``
	//DefinedTagsDependencies +
	//acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_event", "test_event", acctest.Required, acctest.Create, OsManagementHubEventRepresentation) +
	//acctest.GenerateDataSourceFromRepresentationMap("oci_usage_proxy_resources", "test_resources", acctest.Required, acctest.Create, UsageProxyResourceDataSourceRepresentation)
)

// issue-routing-tag: os_management_hub/default
func TestOsManagementHubEventResource_basic(t *testing.T) {
	/*	httpreplay.SetScenario("TestOsManagementHubEventResource_basic")
		defer httpreplay.SaveScenario()

		config := acctest.ProviderTestConfig()

		compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
		compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

		compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
		compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

		eventId := utils.GetEnvSettingWithBlankDefault("event_id")
		eventIdVariableStr := fmt.Sprintf("variable \"event_id\" { default = \"%s\" }\n", eventId)

		eventId2 := utils.GetEnvSettingWithDefault("event_id2", eventId)
		eventIdVariableStr2 := fmt.Sprintf("variable \"event_id2\" { default = \"%s\" }\n", eventId2)

		resourceName := "oci_os_management_hub_event.test_event"
		datasourceName := "data.oci_os_management_hub_events.test_events"
		singularDatasourceName := "data.oci_os_management_hub_event.test_event"

		var resId, resId2 string
		// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
		acctest.SaveConfigContent(config+compartmentIdVariableStr+compartmentIdUVariableStr+OsManagementHubEventResourceDependencies+eventIdVariableStr+
			acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_event", "test_event", acctest.Optional, acctest.Create,
				acctest.RepresentationCopyWithNewProperties(OsManagementHubEventRepresentation, map[string]interface{}{
					"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
				})), "osmanagementhub", "event", t)

		acctest.ResourceTest(t, nil, []resource.TestStep{
			// verify update with tag
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + OsManagementHubEventResourceDependencies + eventIdVariableStr +
					acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_event", "test_event", acctest.Optional, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(OsManagementHubEventRepresentation, map[string]interface{}{
							"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
						})),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "data.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "event_id"),
					resource.TestCheckResourceAttrSet(resourceName, "event_summary"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "type"),

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

			// verify updates to updatable parameters and compartmentId
			{
				Config: config + compartmentIdUVariableStr + OsManagementHubEventResourceDependencies + eventIdVariableStr +
					acctest.GenerateResourceFromRepresentationMap("oci_os_management_hub_event", "test_event", acctest.Optional, acctest.Update,
						acctest.RepresentationCopyWithNewProperties(OsManagementHubEventRepresentation, map[string]interface{}{
							"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
						})),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "data.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "event_id"),
					resource.TestCheckResourceAttrSet(resourceName, "event_summary"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "type"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify resource import
			{
				Config:                  config + OsManagementHubEventRequiredOnlyResource + eventIdVariableStr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
			// verify datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_events", "test_events", acctest.Optional, acctest.Update, OsManagementHubEventDataSourceRepresentation) +
					compartmentIdVariableStr + OsManagementHubEventResourceDependencies + eventIdVariableStr2,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					//resource.TestCheckResourceAttr(datasourceName, "event_fingerprint", "eventFingerprint"),
					resource.TestCheckResourceAttr(datasourceName, "event_summary", "Manually created event 2 for testing caused by <Yijiu>"),
					resource.TestCheckResourceAttr(datasourceName, "event_summary_contains", "testing"),
					//resource.TestCheckResourceAttr(datasourceName, "id", "id"),
					resource.TestCheckResourceAttr(datasourceName, "is_managed_by_autonomous_linux", "true"),
					resource.TestCheckResourceAttrSet(datasourceName, "resource_id"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
					resource.TestCheckResourceAttrSet(datasourceName, "time_created_greater_than_or_equal_to"),
					resource.TestCheckResourceAttrSet(datasourceName, "time_created_less_than"),
					resource.TestCheckResourceAttr(datasourceName, "type.#", "1"),

					resource.TestCheckResourceAttr(datasourceName, "event_collection.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "event_collection.0.items.#", "1"),
				),
			},
			//// delete before next step
			//{
			//	Config: config + compartmentIdVariableStr + OsManagementHubEventResourceDependencies + eventIdVariableStr2,
			//},
			// verify singular datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_os_management_hub_event", "test_event", acctest.Required, acctest.Create, OsManagementHubEventSingularDataSourceRepresentation) +
					compartmentIdVariableStr + eventIdVariableStr2,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "event_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "data.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "event_details"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "event_summary"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "is_managed_by_autonomous_linux"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttr(singularDatasourceName, "system_details.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_occurred"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "type"),
				),
			},
		})*/
}

// TODO
func testAccCheckOsManagementHubEventDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).OsmhEventClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_os_management_hub_event" {
			noResourceFound = false
			request := oci_os_management_hub.GetEventRequest{}

			tmp := rs.Primary.ID
			request.EventId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "os_management_hub")

			response, err := client.GetEvent(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_os_management_hub.EventLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("OsManagementHubEvent") {
		resource.AddTestSweepers("OsManagementHubEvent", &resource.Sweeper{
			Name:         "OsManagementHubEvent",
			Dependencies: acctest.DependencyGraph["event"],
			F:            sweepOsManagementHubEventResource,
		})
	}
}

func sweepOsManagementHubEventResource(compartment string) error {
	eventClient := acctest.GetTestClients(&schema.ResourceData{}).OsmhEventClient()
	eventIds, err := getOsManagementHubEventIds(compartment)
	if err != nil {
		return err
	}
	for _, eventId := range eventIds {
		if ok := acctest.SweeperDefaultResourceId[eventId]; !ok {
			deleteEventRequest := oci_os_management_hub.DeleteEventRequest{}

			deleteEventRequest.EventId = &eventId

			deleteEventRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "os_management_hub")
			_, error := eventClient.DeleteEvent(context.Background(), deleteEventRequest)
			if error != nil {
				fmt.Printf("Error deleting Event %s %s, It is possible that the resource is already deleted. Please verify manually \n", eventId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &eventId, OsManagementHubEventSweepWaitCondition, time.Duration(3*time.Minute),
				OsManagementHubEventSweepResponseFetchOperation, "os_management_hub", true)
		}
	}
	return nil
}

func getOsManagementHubEventIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "EventId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	eventClient := acctest.GetTestClients(&schema.ResourceData{}).OsmhEventClient()

	listEventsRequest := oci_os_management_hub.ListEventsRequest{}
	listEventsRequest.CompartmentId = &compartmentId
	listEventsRequest.LifecycleState = oci_os_management_hub.EventLifecycleStateActive
	listEventsResponse, err := eventClient.ListEvents(context.Background(), listEventsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Event list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, event := range listEventsResponse.Items {
		id := *event.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "EventId", id)
	}
	return resourceIds, nil
}

func OsManagementHubEventSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if eventResponse, ok := response.Response.(oci_os_management_hub.GetEventResponse); ok {
		return eventResponse.GetLifecycleState() != oci_os_management_hub.EventLifecycleStateDeleted
	}
	return false
}

func OsManagementHubEventSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.OsmhEventClient().GetEvent(context.Background(), oci_os_management_hub.GetEventRequest{
		EventId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
