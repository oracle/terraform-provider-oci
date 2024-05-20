// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"encoding/json"
	"fmt"

	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_demand_signal "github.com/oracle/oci-go-sdk/v65/demandsignal"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"

	//"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DemandSignalOccDemandSignalSingularDataSourceRepresentation = map[string]interface{}{
		"occ_demand_signal_id": acctest.Representation{RepType: acctest.Required, Create: `${var.occ_demand_signal_id}`},
	}

	DemandSignalOccDemandSignalDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
	}

	DemandSignalOccDemandSignalOccDemandSignalsRepresentation map[string]interface{}

	DemandSignalOccDemandSignalRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"is_active":      acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		//"occ_demand_signal_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_demand_signal_occ_demand_signal.test_occ_demand_signal.id}`},
		//"occ_demand_signals": acctest.RepresentationGroup{RepType: acctest.Required, Group: DemandSignalOccDemandSignalOccDemandSignalsRepresentationRepresentation},
		"occ_demand_signals": acctest.RepresentationGroup{RepType: acctest.Required, Group: map[string]interface{}{}},

		//"defined_tags":       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		//"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `string`, Update: `displayName2`},
		//"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		//"patch_operations": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DemandSignalOccDemandSignalPatchOperationsRepresentation},
	}
	DemandSignalOccDemandSignalUpdateRepresentation = map[string]interface{}{
		"compartment_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"is_active":            acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"occ_demand_signal_id": acctest.Representation{RepType: acctest.Required, Update: `${var.occ_demand_signal_id}`},
		"occ_demand_signals":   acctest.RepresentationGroup{RepType: acctest.Required, Group: DemandSignalOccDemandSignalOccDemandSignalsRepresentationRepresentation},

		//"defined_tags":       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `testDemandSignal`, Update: `displayName2`},
		//"patch_operations": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DemandSignalOccDemandSignalPatchOperationsRepresentation},
	}

	DemandSignalOccDemandSignalOccDemandSignalsRepresentationRepresentation = map[string]interface{}{
		"resource_type": acctest.Representation{RepType: acctest.Required, Create: `Compute - Std Intel`},
		"units":         acctest.Representation{RepType: acctest.Required, Create: `(Cores)`},
		"values":        acctest.RepresentationGroup{RepType: acctest.Required, Group: DemandSignalOccDemandSignalValuesRepresentationRepresentation},
		//"values":        acctest.Representation{RepType: acctest.Required, Create: "${var.occ_demand_signal_test_value}"},
	}
	DemandSignalOccDemandSignalValuesRepresentationRepresentation = map[string]interface{}{
		"time_expected": acctest.Representation{RepType: acctest.Required, Create: `2025-01-05T00:00:00.000Z`},
		"value":         acctest.Representation{RepType: acctest.Required, Create: `100`},
	}
) //end var

/*"time_expected": acctest.Representation{RepType: acctest.Required, Create: []string{`2025-01-05T00:00:00.000Z, 2025-02-05T00:00:00.000Z, 2025-03-05T00:00:00.000Z, 2025-04-05T00:00:00.000Z, 2025-05-05T00:00:00.000Z, 2025-06-05T00:00:00.000Z, 2025-07-05T00:00:00.000Z, 2025-08-05T00:00:00.000Z, 2025-09-05T00:00:00.000Z, 2025-10-05T00:00:00.000Z, 2025-11-05T00:00:00.000Z, 2025-12-05T00:00:00.000Z`}},
"value":         acctest.Representation{RepType: acctest.Required, Create: []string{`0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0`}},
"time_expected": acctest.Representation{RepType: acctest.Required, Create: []string{"2025-01-05T00:00:00.000Z", "2025-02-05T00:00:00.000Z", "2025-03-05T00:00:00.000Z", "2025-04-05T00:00:00.000Z", "2025-05-05T00:00:00.000Z", "2025-06-05T00:00:00.000Z", "2025-07-05T00:00:00.000Z", "2025-08-05T00:00:00.000Z", "2025-09-05T00:00:00.000Z", "2025-10-05T00:00:00.000Z", "2025-11-05T00:00:00.000Z", "2025-12-05T00:00:00.000Z"}},
"value":         acctest.Representation{RepType: acctest.Required, Create: []float64{100.0, 100.0, 100.0, 100.0, 100.0, 100.0, 100.0, 100.0, 100.0, 100.0, 100.0, 100.0}},
	}
*/

// issue-routing-tag: demand_signal/default
func TestDemandSignalOccDemandSignalResource_basic(t *testing.T) {
	fmt.Println(DemandSignalOccDemandSignalOccDemandSignalsRepresentation)

	jsonData, err := json.MarshalIndent(DemandSignalOccDemandSignalOccDemandSignalsRepresentation, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData))

	httpreplay.SetScenario("TestDemandSignalOccDemandSignalResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	occDemandSignalId := utils.GetEnvSettingWithBlankDefault("occ_demand_signal_id")
	occDemandSignalIdVariableStr := fmt.Sprintf("variable \"occ_demand_signal_id\" { default = \"%s\" }\n", occDemandSignalId)

	//resourceName := "oci_demand_signal_occ_demand_signal.test_occ_demand_signal"
	datasourceName := "data.oci_demand_signal_occ_demand_signals.test_occ_demand_signals"
	singularDatasourceName := "data.oci_demand_signal_occ_demand_signal.test_occ_demand_signal"

	//var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	//acctest.SaveConfigContent(config+compartmentIdVariableStr+ acctest.GenerateResourceFromRepresentationMap("oci_demand_signal_occ_demand_signal", "test_occ_demand_signal", acctest.Optional, acctest.Create, DemandSignalOccDemandSignalRepresentation), "demandsignal", "occDemandSignal", t)

	acctest.ResourceTest(t, testAccCheckDemandSignalOccDemandSignalDestroy, []resource.TestStep{

		//verify Create
		/*{
			Config: config + compartmentIdVariableStr + //DemandSignalOccDemandSignalResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_demand_signal_occ_demand_signal", "test_occ_demand_signal", acctest.Required, acctest.Create, DemandSignalOccDemandSignalRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "is_active", "false"),
				//resource.TestCheckResourceAttr(resourceName, "occ_demand_signals.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "occ_demand_signals.0.resource_type", "string"),
				//resource.TestCheckResourceAttr(resourceName, "occ_demand_signals.0.units", "string"),
				//resource.TestCheckResourceAttr(resourceName, "occ_demand_signals.0.values.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "occ_demand_signals.0.values.0.time_expected", "2024-05-01T00:00:00.000Z"),
				//resource.TestCheckResourceAttr(resourceName, "occ_demand_signals.0.values.0.value", "1.0"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},*/

		//// delete before next Create
		//{
		//	Config: config + compartmentIdVariableStr, // + DemandSignalOccDemandSignalResourceDependencies,
		//},
		//// verify Create with optionals
		//{
		//	Config: config + compartmentIdVariableStr + //DemandSignalOccDemandSignalResourceDependencies +
		//		acctest.GenerateResourceFromRepresentationMap("oci_demand_signal_occ_demand_signal", "test_occ_demand_signal", acctest.Optional, acctest.Create, DemandSignalOccDemandSignalRepresentation),
		//	Check: acctest.ComposeAggregateTestCheckFuncWrapper(
		//		resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
		//		resource.TestCheckResourceAttr(resourceName, "display_name", "string"),
		//		resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
		//		resource.TestCheckResourceAttrSet(resourceName, "id"),
		//		resource.TestCheckResourceAttr(resourceName, "is_active", "false"),
		//		resource.TestCheckResourceAttr(resourceName, "occ_demand_signals.#", "1"),
		//		resource.TestCheckResourceAttr(resourceName, "occ_demand_signals.0.resource_type", "string"),
		//		resource.TestCheckResourceAttr(resourceName, "occ_demand_signals.0.units", "string"),
		//		resource.TestCheckResourceAttr(resourceName, "occ_demand_signals.0.values.#", "1"),
		//		resource.TestCheckResourceAttr(resourceName, "occ_demand_signals.0.values.0.comments", "string"),
		//		resource.TestCheckResourceAttr(resourceName, "occ_demand_signals.0.values.0.time_expected", "2024-05-01T00:00:00.000Z"),
		//		resource.TestCheckResourceAttr(resourceName, "occ_demand_signals.0.values.0.value", "1.0"),
		//		resource.TestCheckResourceAttrSet(resourceName, "state"),
		//		resource.TestCheckResourceAttrSet(resourceName, "time_created"),
		//
		//		func(s *terraform.State) (err error) {
		//			resId, err = acctest.FromInstanceState(s, resourceName, "id")
		//			if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
		//				if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
		//					return errExport
		//				}
		//			}
		//			return err
		//		},
		//	),
		//},
		//
		//// verify updates to updatable parameters
		/*{
			Config: config + compartmentIdVariableStr + //DemandSignalOccDemandSignalResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_demand_signal_occ_demand_signal", "test_occ_demand_signal", acctest.Optional, acctest.Update, DemandSignalOccDemandSignalUpdateRepresentation) + occDemandSignalIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_active", "false"),
				//resource.TestCheckResourceAttr(resourceName, "occ_demand_signals.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "occ_demand_signals.0.resource_type", "string"),
				//resource.TestCheckResourceAttr(resourceName, "occ_demand_signals.0.units", "string"),
				//resource.TestCheckResourceAttr(resourceName, "occ_demand_signals.0.values.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "occ_demand_signals.0.values.0.comments", "string"),
				//resource.TestCheckResourceAttr(resourceName, "occ_demand_signals.0.values.0.time_expected", "2024-05-01T00:00:00.000Z"),
				//resource.TestCheckResourceAttr(resourceName, "occ_demand_signals.0.values.0.value", "1.0"),
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
		},*/

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_demand_signal_occ_demand_signals", "test_occ_demand_signals", acctest.Optional, acctest.Update, DemandSignalOccDemandSignalDataSourceRepresentation) +
				occDemandSignalIdVariableStr +
				compartmentIdVariableStr, // + DemandSignalOccDemandSignalResourceDependencies +

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "occ_demand_signal_collection.0.items.0.display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "occ_demand_signal_collection.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_demand_signal_occ_demand_signal", "test_occ_demand_signal", acctest.Required, acctest.Create, DemandSignalOccDemandSignalSingularDataSourceRepresentation) +
				occDemandSignalIdVariableStr +
				compartmentIdVariableStr, //+ DemandSignalOccDemandSignalResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "testing-cli-update"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_active", "false"),
				/*resource.TestCheckResourceAttr(singularDatasourceName, "occ_demand_signals.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "occ_demand_signals.0.resource_type", "string"),
				resource.TestCheckResourceAttr(singularDatasourceName, "occ_demand_signals.0.units", "string"),
				resource.TestCheckResourceAttr(singularDatasourceName, "occ_demand_signals.0.values.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "occ_demand_signals.0.values.0.comments", "string"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "occ_demand_signals.0.values.0.time_expected"),
				resource.TestCheckResourceAttr(singularDatasourceName, "occ_demand_signals.0.values.0.value", "1.0"),*/
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},

		//// verify resource import
		//{
		//	Config:            config + DemandSignalOccDemandSignalRequiredOnlyResource,
		//	ImportState:       true,
		//	ImportStateVerify: true,
		//	ImportStateVerifyIgnore: []string{
		//		"patch_operations",
		//	},
		//	ResourceName: resourceName,
		//},
	})
}

func testAccCheckDemandSignalOccDemandSignalDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).OccDemandSignalClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_demand_signal_occ_demand_signal" {
			noResourceFound = false
			request := oci_demand_signal.GetOccDemandSignalRequest{}

			tmp := rs.Primary.ID
			request.OccDemandSignalId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "demand_signal")

			response, err := client.GetOccDemandSignal(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_demand_signal.OccDemandSignalLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					continue
					//resource lifecycle state is not in expected deleted lifecycle states.
					//return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
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
	if !acctest.InSweeperExcludeList("DemandSignalOccDemandSignal") {
		resource.AddTestSweepers("DemandSignalOccDemandSignal", &resource.Sweeper{
			Name:         "DemandSignalOccDemandSignal",
			Dependencies: acctest.DependencyGraph["occDemandSignal"],
			F:            sweepDemandSignalOccDemandSignalResource,
		})
	}

	var err error
	DemandSignalOccDemandSignalOccDemandSignalsRepresentation, err = GenerateValues()
	if err != nil {
		panic(err)
	}
}

func sweepDemandSignalOccDemandSignalResource(compartment string) error {
	occDemandSignalClient := acctest.GetTestClients(&schema.ResourceData{}).OccDemandSignalClient()
	occDemandSignalIds, err := getDemandSignalOccDemandSignalIds(compartment)
	if err != nil {
		return err
	}
	for _, occDemandSignalId := range occDemandSignalIds {
		if ok := acctest.SweeperDefaultResourceId[occDemandSignalId]; !ok {
			deleteOccDemandSignalRequest := oci_demand_signal.DeleteOccDemandSignalRequest{}

			deleteOccDemandSignalRequest.OccDemandSignalId = &occDemandSignalId

			deleteOccDemandSignalRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "demand_signal")
			_, error := occDemandSignalClient.DeleteOccDemandSignal(context.Background(), deleteOccDemandSignalRequest)
			if error != nil {
				fmt.Printf("Error deleting OccDemandSignal %s %s, It is possible that the resource is already deleted. Please verify manually \n", occDemandSignalId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &occDemandSignalId, DemandSignalOccDemandSignalSweepWaitCondition, time.Duration(3*time.Minute),
				DemandSignalOccDemandSignalSweepResponseFetchOperation, "demand_signal", true)
		}
	}
	return nil
}

func getDemandSignalOccDemandSignalIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "OccDemandSignalId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	occDemandSignalClient := acctest.GetTestClients(&schema.ResourceData{}).OccDemandSignalClient()

	listOccDemandSignalsRequest := oci_demand_signal.ListOccDemandSignalsRequest{}
	listOccDemandSignalsRequest.CompartmentId = &compartmentId
	listOccDemandSignalsRequest.LifecycleState = oci_demand_signal.OccDemandSignalLifecycleStateActive
	listOccDemandSignalsResponse, err := occDemandSignalClient.ListOccDemandSignals(context.Background(), listOccDemandSignalsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting OccDemandSignal list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, occDemandSignal := range listOccDemandSignalsResponse.Items {
		id := *occDemandSignal.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "OccDemandSignalId", id)
	}
	return resourceIds, nil
}

func DemandSignalOccDemandSignalSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if occDemandSignalResponse, ok := response.Response.(oci_demand_signal.GetOccDemandSignalResponse); ok {
		return occDemandSignalResponse.LifecycleState != oci_demand_signal.OccDemandSignalLifecycleStateDeleted
	}
	return false
}

func DemandSignalOccDemandSignalSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.OccDemandSignalClient().GetOccDemandSignal(context.Background(), oci_demand_signal.GetOccDemandSignalRequest{
		OccDemandSignalId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}

// GenerateValues creates a map with the specified structure and values starting from the current date
func GenerateValues() (map[string]interface{}, error) {
	// Get the current date
	currentTime := time.Now()

	// Create the values slice
	values := make([]map[string]interface{}, 12)

	for i := 0; i < 12; i++ {
		// Get the first date of each month starting from the current month
		firstDateOfMonth := time.Date(currentTime.Year(), currentTime.Month()+time.Month(i), 1, 0, 0, 0, 0, time.UTC)
		date := firstDateOfMonth.Format(time.RFC3339)
		value := 0
		if i == 0 {
			value = 100
		}

		values[i] = map[string]interface{}{
			"timeExpected": date,
			"value":        value,
			"comments":     nil,
		}
	}

	return map[string]interface{}{
		"resource_type": "Compute - Std Intel",
		"units":         "(Cores)",
		"values":        values, // Directly use the values slice
	}, nil
}
