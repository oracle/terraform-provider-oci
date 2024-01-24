// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
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
	oci_stack_monitoring "github.com/oracle/oci-go-sdk/v65/stackmonitoring"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	StackMonitoringProcessSetRequiredOnlyResource = StackMonitoringProcessSetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_process_set", "test_process_set", acctest.Required, acctest.Create, StackMonitoringProcessSetRepresentation)

	StackMonitoringProcessSetResourceConfig = StackMonitoringProcessSetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_process_set", "test_process_set", acctest.Optional, acctest.Update, StackMonitoringProcessSetRepresentation)

	StackMonitoringProcessSetSingularDataSourceRepresentation = map[string]interface{}{
		"process_set_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_stack_monitoring_process_set.test_process_set.id}`},
	}

	StackMonitoringProcessSetDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `nginx workers`, Update: `displayName2`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: StackMonitoringProcessSetDataSourceFilterRepresentation}}
	StackMonitoringProcessSetDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_stack_monitoring_process_set.test_process_set.id}`}},
	}

	StackMonitoringProcessSetRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `nginx workers`, Update: `displayName2`},
		"specification":  acctest.RepresentationGroup{RepType: acctest.Required, Group: StackMonitoringProcessSetSpecificationRepresentation},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
	}
	StackMonitoringProcessSetSpecificationRepresentation = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Required, Group: StackMonitoringProcessSetSpecificationItemsRepresentation},
	}
	StackMonitoringProcessSetSpecificationItemsRepresentation = map[string]interface{}{
		"label":                      acctest.Representation{RepType: acctest.Optional, Create: `nginx-workers`, Update: `label2`},
		"process_command":            acctest.Representation{RepType: acctest.Required, Create: `nginx`, Update: `processCommand2`},
		"process_line_regex_pattern": acctest.Representation{RepType: acctest.Optional, Create: `nginx: worker.`, Update: `processLineRegexPattern2`},
		"process_user":               acctest.Representation{RepType: acctest.Optional, Create: `opc`, Update: `processUser2`},
	}

	StackMonitoringProcessSetResourceDependencies = ""
)

// issue-routing-tag: stack_monitoring/default
func TestStackMonitoringProcessSetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestStackMonitoringProcessSetResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_stack_monitoring_process_set.test_process_set"
	datasourceName := "data.oci_stack_monitoring_process_sets.test_process_sets"
	singularDatasourceName := "data.oci_stack_monitoring_process_set.test_process_set"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+StackMonitoringProcessSetResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_process_set", "test_process_set", acctest.Optional, acctest.Create, StackMonitoringProcessSetRepresentation), "stackmonitoring", "processSet", t)

	acctest.ResourceTest(t, testAccCheckStackMonitoringProcessSetDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + StackMonitoringProcessSetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_process_set", "test_process_set", acctest.Required, acctest.Create, StackMonitoringProcessSetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "nginx workers"),
				resource.TestCheckResourceAttr(resourceName, "specification.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.items.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + StackMonitoringProcessSetResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + StackMonitoringProcessSetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_process_set", "test_process_set", acctest.Optional, acctest.Create, StackMonitoringProcessSetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "nginx workers"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "revision"),
				resource.TestCheckResourceAttr(resourceName, "specification.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.items.0.label", "nginx-workers"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.items.0.process_command", "nginx"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.items.0.process_line_regex_pattern", "nginx: worker."),
				resource.TestCheckResourceAttr(resourceName, "specification.0.items.0.process_user", "opc"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + StackMonitoringProcessSetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_process_set", "test_process_set", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(StackMonitoringProcessSetRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "nginx workers"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "revision"),
				resource.TestCheckResourceAttr(resourceName, "specification.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.items.0.label", "nginx-workers"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.items.0.process_command", "nginx"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.items.0.process_line_regex_pattern", "nginx: worker."),
				resource.TestCheckResourceAttr(resourceName, "specification.0.items.0.process_user", "opc"),
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
			Config: config + compartmentIdVariableStr + StackMonitoringProcessSetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_process_set", "test_process_set", acctest.Optional, acctest.Update, StackMonitoringProcessSetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "revision"),
				resource.TestCheckResourceAttr(resourceName, "specification.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.items.0.label", "label2"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.items.0.process_command", "processCommand2"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.items.0.process_line_regex_pattern", "processLineRegexPattern2"),
				resource.TestCheckResourceAttr(resourceName, "specification.0.items.0.process_user", "processUser2"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_stack_monitoring_process_sets", "test_process_sets", acctest.Optional, acctest.Update, StackMonitoringProcessSetDataSourceRepresentation) +
				compartmentIdVariableStr + StackMonitoringProcessSetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_stack_monitoring_process_set", "test_process_set", acctest.Optional, acctest.Update, StackMonitoringProcessSetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

				resource.TestCheckResourceAttr(datasourceName, "process_set_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "process_set_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_stack_monitoring_process_set", "test_process_set", acctest.Required, acctest.Create, StackMonitoringProcessSetSingularDataSourceRepresentation) +
				compartmentIdVariableStr + StackMonitoringProcessSetResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "process_set_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "revision"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.items.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.items.0.label", "label2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.items.0.process_command", "processCommand2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.items.0.process_line_regex_pattern", "processLineRegexPattern2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "specification.0.items.0.process_user", "processUser2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + StackMonitoringProcessSetRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckStackMonitoringProcessSetDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).StackMonitoringClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_stack_monitoring_process_set" {
			noResourceFound = false
			request := oci_stack_monitoring.GetProcessSetRequest{}

			tmp := rs.Primary.ID
			request.ProcessSetId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "stack_monitoring")

			response, err := client.GetProcessSet(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_stack_monitoring.LifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("StackMonitoringProcessSet") {
		resource.AddTestSweepers("StackMonitoringProcessSet", &resource.Sweeper{
			Name:         "StackMonitoringProcessSet",
			Dependencies: acctest.DependencyGraph["processSet"],
			F:            sweepStackMonitoringProcessSetResource,
		})
	}
}

func sweepStackMonitoringProcessSetResource(compartment string) error {
	stackMonitoringClient := acctest.GetTestClients(&schema.ResourceData{}).StackMonitoringClient()
	processSetIds, err := getStackMonitoringProcessSetIds(compartment)
	if err != nil {
		return err
	}
	for _, processSetId := range processSetIds {
		if ok := acctest.SweeperDefaultResourceId[processSetId]; !ok {
			deleteProcessSetRequest := oci_stack_monitoring.DeleteProcessSetRequest{}

			deleteProcessSetRequest.ProcessSetId = &processSetId

			deleteProcessSetRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "stack_monitoring")
			_, error := stackMonitoringClient.DeleteProcessSet(context.Background(), deleteProcessSetRequest)
			if error != nil {
				fmt.Printf("Error deleting ProcessSet %s %s, It is possible that the resource is already deleted. Please verify manually \n", processSetId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &processSetId, StackMonitoringProcessSetSweepWaitCondition, time.Duration(3*time.Minute),
				StackMonitoringProcessSetSweepResponseFetchOperation, "stack_monitoring", true)
		}
	}
	return nil
}

func getStackMonitoringProcessSetIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ProcessSetId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	stackMonitoringClient := acctest.GetTestClients(&schema.ResourceData{}).StackMonitoringClient()

	listProcessSetsRequest := oci_stack_monitoring.ListProcessSetsRequest{}
	listProcessSetsRequest.CompartmentId = &compartmentId
	listProcessSetsResponse, err := stackMonitoringClient.ListProcessSets(context.Background(), listProcessSetsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ProcessSet list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, processSet := range listProcessSetsResponse.Items {
		id := *processSet.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ProcessSetId", id)
	}
	return resourceIds, nil
}

func StackMonitoringProcessSetSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if processSetResponse, ok := response.Response.(oci_stack_monitoring.GetProcessSetResponse); ok {
		return processSetResponse.LifecycleState != oci_stack_monitoring.LifecycleStateDeleted
	}
	return false
}

func StackMonitoringProcessSetSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.StackMonitoringClient().GetProcessSet(context.Background(), oci_stack_monitoring.GetProcessSetRequest{
		ProcessSetId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
