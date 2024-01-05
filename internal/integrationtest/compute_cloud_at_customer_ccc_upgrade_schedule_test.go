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
	oci_compute_cloud_at_customer "github.com/oracle/oci-go-sdk/v65/computecloudatcustomer"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ComputeCloudAtCustomerCccUpgradeScheduleRequiredOnlyResource = ComputeCloudAtCustomerCccUpgradeScheduleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_compute_cloud_at_customer_ccc_upgrade_schedule", "test_ccc_upgrade_schedule", acctest.Required, acctest.Create, ComputeCloudAtCustomerCccUpgradeScheduleRepresentation)

	ComputeCloudAtCustomerCccUpgradeScheduleResourceConfig = ComputeCloudAtCustomerCccUpgradeScheduleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_compute_cloud_at_customer_ccc_upgrade_schedule", "test_ccc_upgrade_schedule", acctest.Optional, acctest.Update, ComputeCloudAtCustomerCccUpgradeScheduleRepresentation)

	ComputeCloudAtCustomerCccUpgradeScheduleSingularDataSourceRepresentation = map[string]interface{}{
		"ccc_upgrade_schedule_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_compute_cloud_at_customer_ccc_upgrade_schedule.test_ccc_upgrade_schedule.id}`},
	}

	ComputeCloudAtCustomerCccUpgradeScheduleDataSourceRepresentation = map[string]interface{}{
		"access_level":              acctest.Representation{RepType: acctest.Optional, Create: `RESTRICTED`},
		"ccc_upgrade_schedule_id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_compute_cloud_at_customer_ccc_upgrade_schedule.test_ccc_upgrade_schedule.id}`},
		"compartment_id":            acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `example_cccUpgradeSchedule`, Update: `displayName2`},
		"display_name_contains":     acctest.Representation{RepType: acctest.Optional, Create: `displayNameContains`},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: ComputeCloudAtCustomerCccUpgradeScheduleDataSourceFilterRepresentation}}
	ComputeCloudAtCustomerCccUpgradeScheduleDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_compute_cloud_at_customer_ccc_upgrade_schedule.test_ccc_upgrade_schedule.id}`}},
	}

	ComputeCloudAtCustomerCccUpgradeScheduleRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `example_cccUpgradeSchedule`, Update: `displayName2`},
		"events":         acctest.RepresentationGroup{RepType: acctest.Required, Group: ComputeCloudAtCustomerCccUpgradeScheduleEventsRepresentation},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `Month-start upgrade window`, Update: `description2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
	}

	ComputeCloudAtCustomerCccUpgradeScheduleEventsRepresentation = map[string]interface{}{
		"description":                acctest.Representation{RepType: acctest.Required, Create: `description`, Update: `description2`},
		"schedule_event_duration":    acctest.Representation{RepType: acctest.Required, Create: `PT49H`, Update: `PT49H`},
		"time_start":                 acctest.Representation{RepType: acctest.Required, Create: `2023-09-09T16:10:25Z`, Update: `2023-09-09T16:10:25Z`},
		"schedule_event_recurrences": acctest.Representation{RepType: acctest.Required, Create: `FREQ=MONTHLY;INTERVAL=3;`, Update: `FREQ=MONTHLY;INTERVAL=3;`},
	}

	ComputeCloudAtCustomerCccUpgradeScheduleResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_compute_cloud_at_customer_ccc_upgrade_schedule", "test_ccc_upgrade_schedule_def", acctest.Required, acctest.Create, ComputeCloudAtCustomerCccUpgradeScheduleRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: compute_cloud_at_customer/default
func TestComputeCloudAtCustomerCccUpgradeScheduleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestComputeCloudAtCustomerCccUpgradeScheduleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_compute_cloud_at_customer_ccc_upgrade_schedule.test_ccc_upgrade_schedule"
	datasourceName := "data.oci_compute_cloud_at_customer_ccc_upgrade_schedules.test_ccc_upgrade_schedules"
	singularDatasourceName := "data.oci_compute_cloud_at_customer_ccc_upgrade_schedule.test_ccc_upgrade_schedule"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ComputeCloudAtCustomerCccUpgradeScheduleResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_compute_cloud_at_customer_ccc_upgrade_schedule", "test_ccc_upgrade_schedule", acctest.Optional, acctest.Create, ComputeCloudAtCustomerCccUpgradeScheduleRepresentation), "computecloudatcustomer", "cccUpgradeSchedule", t)

	acctest.ResourceTest(t, testAccCheckComputeCloudAtCustomerCccUpgradeScheduleDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ComputeCloudAtCustomerCccUpgradeScheduleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_compute_cloud_at_customer_ccc_upgrade_schedule", "test_ccc_upgrade_schedule", acctest.Required, acctest.Create, ComputeCloudAtCustomerCccUpgradeScheduleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_cccUpgradeSchedule"),
				resource.TestCheckResourceAttr(resourceName, "events.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "events.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "events.0.schedule_event_duration", "PT49H"),
				resource.TestCheckResourceAttr(resourceName, "events.0.time_start", "2023-09-09T16:10:25Z"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ComputeCloudAtCustomerCccUpgradeScheduleResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ComputeCloudAtCustomerCccUpgradeScheduleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_compute_cloud_at_customer_ccc_upgrade_schedule", "test_ccc_upgrade_schedule", acctest.Optional, acctest.Create, ComputeCloudAtCustomerCccUpgradeScheduleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "Month-start upgrade window"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_cccUpgradeSchedule"),
				resource.TestCheckResourceAttr(resourceName, "events.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "events.0.description", "description"),
				resource.TestCheckResourceAttrSet(resourceName, "events.0.name"),
				resource.TestCheckResourceAttr(resourceName, "events.0.schedule_event_duration", "PT49H"),
				resource.TestCheckResourceAttr(resourceName, "events.0.schedule_event_recurrences", "FREQ=MONTHLY;INTERVAL=3;"),
				resource.TestCheckResourceAttr(resourceName, "events.0.time_start", "2023-09-09T16:10:25Z"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ComputeCloudAtCustomerCccUpgradeScheduleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_compute_cloud_at_customer_ccc_upgrade_schedule", "test_ccc_upgrade_schedule", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(ComputeCloudAtCustomerCccUpgradeScheduleRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "Month-start upgrade window"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "example_cccUpgradeSchedule"),
				resource.TestCheckResourceAttr(resourceName, "events.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "events.0.description", "description"),
				resource.TestCheckResourceAttrSet(resourceName, "events.0.name"),
				resource.TestCheckResourceAttr(resourceName, "events.0.schedule_event_duration", "PT49H"),
				resource.TestCheckResourceAttr(resourceName, "events.0.schedule_event_recurrences", "FREQ=MONTHLY;INTERVAL=3;"),
				resource.TestCheckResourceAttr(resourceName, "events.0.time_start", "2023-09-09T16:10:25Z"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
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
			Config: config + compartmentIdVariableStr + ComputeCloudAtCustomerCccUpgradeScheduleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_compute_cloud_at_customer_ccc_upgrade_schedule", "test_ccc_upgrade_schedule", acctest.Optional, acctest.Update, ComputeCloudAtCustomerCccUpgradeScheduleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "events.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "events.0.description", "description2"),
				resource.TestCheckResourceAttrSet(resourceName, "events.0.name"),
				resource.TestCheckResourceAttr(resourceName, "events.0.schedule_event_duration", "PT49H"),
				resource.TestCheckResourceAttr(resourceName, "events.0.schedule_event_recurrences", "FREQ=MONTHLY;INTERVAL=3;"),
				resource.TestCheckResourceAttr(resourceName, "events.0.time_start", "2023-09-09T16:10:25Z"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_compute_cloud_at_customer_ccc_upgrade_schedules", "test_ccc_upgrade_schedules", acctest.Optional, acctest.Update, ComputeCloudAtCustomerCccUpgradeScheduleDataSourceRepresentation) +
				compartmentIdVariableStr + ComputeCloudAtCustomerCccUpgradeScheduleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_compute_cloud_at_customer_ccc_upgrade_schedule", "test_ccc_upgrade_schedule", acctest.Optional, acctest.Update, ComputeCloudAtCustomerCccUpgradeScheduleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "access_level", "RESTRICTED"),
				resource.TestCheckResourceAttrSet(datasourceName, "ccc_upgrade_schedule_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "display_name_contains", "displayNameContains"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "ccc_upgrade_schedule_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "ccc_upgrade_schedule_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_compute_cloud_at_customer_ccc_upgrade_schedule", "test_ccc_upgrade_schedule", acctest.Required, acctest.Create, ComputeCloudAtCustomerCccUpgradeScheduleSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ComputeCloudAtCustomerCccUpgradeScheduleResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ccc_upgrade_schedule_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "events.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "events.0.description", "description2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "events.0.name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "events.0.schedule_event_duration", "PT49H"),
				resource.TestCheckResourceAttr(singularDatasourceName, "events.0.schedule_event_recurrences", "FREQ=MONTHLY;INTERVAL=3;"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "events.0.time_start"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				// resource.TestCheckResourceAttr(singularDatasourceName, "infrastructure_ids.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + ComputeCloudAtCustomerCccUpgradeScheduleRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckComputeCloudAtCustomerCccUpgradeScheduleDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ComputeCloudAtCustomerClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_compute_cloud_at_customer_ccc_upgrade_schedule" {
			noResourceFound = false
			request := oci_compute_cloud_at_customer.GetCccUpgradeScheduleRequest{}

			tmp := rs.Primary.ID
			request.CccUpgradeScheduleId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "compute_cloud_at_customer")

			response, err := client.GetCccUpgradeSchedule(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_compute_cloud_at_customer.CccUpgradeScheduleLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("ComputeCloudAtCustomerCccUpgradeSchedule") {
		resource.AddTestSweepers("ComputeCloudAtCustomerCccUpgradeSchedule", &resource.Sweeper{
			Name:         "ComputeCloudAtCustomerCccUpgradeSchedule",
			Dependencies: acctest.DependencyGraph["cccUpgradeSchedule"],
			F:            sweepComputeCloudAtCustomerCccUpgradeScheduleResource,
		})
	}
}

func sweepComputeCloudAtCustomerCccUpgradeScheduleResource(compartment string) error {
	computeCloudAtCustomerClient := acctest.GetTestClients(&schema.ResourceData{}).ComputeCloudAtCustomerClient()
	cccUpgradeScheduleIds, err := getComputeCloudAtCustomerCccUpgradeScheduleIds(compartment)
	if err != nil {
		return err
	}
	for _, cccUpgradeScheduleId := range cccUpgradeScheduleIds {
		if ok := acctest.SweeperDefaultResourceId[cccUpgradeScheduleId]; !ok {
			deleteCccUpgradeScheduleRequest := oci_compute_cloud_at_customer.DeleteCccUpgradeScheduleRequest{}

			deleteCccUpgradeScheduleRequest.CccUpgradeScheduleId = &cccUpgradeScheduleId

			deleteCccUpgradeScheduleRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "compute_cloud_at_customer")
			_, error := computeCloudAtCustomerClient.DeleteCccUpgradeSchedule(context.Background(), deleteCccUpgradeScheduleRequest)
			if error != nil {
				fmt.Printf("Error deleting CccUpgradeSchedule %s %s, It is possible that the resource is already deleted. Please verify manually \n", cccUpgradeScheduleId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &cccUpgradeScheduleId, ComputeCloudAtCustomerCccUpgradeScheduleSweepWaitCondition, time.Duration(3*time.Minute),
				ComputeCloudAtCustomerCccUpgradeScheduleSweepResponseFetchOperation, "compute_cloud_at_customer", true)
		}
	}
	return nil
}

func getComputeCloudAtCustomerCccUpgradeScheduleIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "CccUpgradeScheduleId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	computeCloudAtCustomerClient := acctest.GetTestClients(&schema.ResourceData{}).ComputeCloudAtCustomerClient()

	listCccUpgradeSchedulesRequest := oci_compute_cloud_at_customer.ListCccUpgradeSchedulesRequest{}
	listCccUpgradeSchedulesRequest.CompartmentId = &compartmentId
	listCccUpgradeSchedulesRequest.LifecycleState = oci_compute_cloud_at_customer.CccUpgradeScheduleLifecycleStateActive
	listCccUpgradeSchedulesResponse, err := computeCloudAtCustomerClient.ListCccUpgradeSchedules(context.Background(), listCccUpgradeSchedulesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting CccUpgradeSchedule list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, cccUpgradeSchedule := range listCccUpgradeSchedulesResponse.Items {
		id := *cccUpgradeSchedule.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "CccUpgradeScheduleId", id)
	}
	return resourceIds, nil
}

func ComputeCloudAtCustomerCccUpgradeScheduleSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if cccUpgradeScheduleResponse, ok := response.Response.(oci_compute_cloud_at_customer.GetCccUpgradeScheduleResponse); ok {
		return cccUpgradeScheduleResponse.LifecycleState != oci_compute_cloud_at_customer.CccUpgradeScheduleLifecycleStateDeleted
	}
	return false
}

func ComputeCloudAtCustomerCccUpgradeScheduleSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ComputeCloudAtCustomerClient().GetCccUpgradeSchedule(context.Background(), oci_compute_cloud_at_customer.GetCccUpgradeScheduleRequest{
		CccUpgradeScheduleId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
