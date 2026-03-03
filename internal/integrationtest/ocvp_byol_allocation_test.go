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
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_ocvp "github.com/oracle/oci-go-sdk/v65/ocvp"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OcvpByolAllocationRequiredOnlyResource = OcvpByolAllocationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ocvp_byol_allocation", "test_byol_allocation", acctest.Required, acctest.Create, OcvpByolAllocationRepresentation)

	OcvpByolAllocationResourceConfig = OcvpByolAllocationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ocvp_byol_allocation", "test_byol_allocation", acctest.Optional, acctest.Update, OcvpByolAllocationRepresentation)

	OcvpByolAllocationSingularDataSourceRepresentation = map[string]interface{}{
		"byol_allocation_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_ocvp_byol_allocation.test_byol_allocation.id}`},
	}

	OcvpByolAllocationDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"available_units_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `1`},
		"byol_allocation_id":                       acctest.Representation{RepType: acctest.Optional, Create: `${oci_ocvp_byol_allocation.test_byol_allocation.id}`},
		"byol_id":                                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_ocvp_byol.test_byol.id}`},
		"display_name":                             acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"software_type":                            acctest.Representation{RepType: acctest.Optional, Create: `VCF`},
		"state":                                    acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                                   acctest.RepresentationGroup{RepType: acctest.Required, Group: OcvpByolAllocationDataSourceFilterRepresentation}}

	OcvpByolAllocationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_ocvp_byol_allocation.test_byol_allocation.id}`}},
	}

	OcvpByolAllocationRepresentation = map[string]interface{}{
		"allocated_units": acctest.Representation{RepType: acctest.Required, Create: `100`, Update: `110`},
		"byol_id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_ocvp_byol.test_byol.id}`},
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":    acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"defined_tags":    acctest.Representation{RepType: acctest.Optional, Create: ocvpDefinedTag, Update: ocvpDefinedTagUpdate},
		"freeform_tags":   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":       acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesRepresentation},
	}

	OcvpByolAllocationVdefendRepresentation = acctest.RepresentationCopyWithNewProperties(
		acctest.RepresentationCopyWithRemovedProperties(OcvpByolAllocationRepresentation, []string{"byol_id"}),
		map[string]interface{}{
			"byol_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_ocvp_byol.test_byol_vdefend.id}`},
		})

	OcvpByolAllocationLbRepresentation = acctest.RepresentationCopyWithNewProperties(
		acctest.RepresentationCopyWithRemovedProperties(OcvpByolAllocationRepresentation, []string{"byol_id"}),
		map[string]interface{}{
			"byol_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_ocvp_byol.test_byol_lb.id}`},
		})

	OcvpByolAllocationResourceDependencies = ocvpDefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ocvp_byol", "test_byol", acctest.Required, acctest.Create, OcvpByolRepresentation)
)

// issue-routing-tag: ocvp/default
func TestOcvpByolAllocationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOcvpByolAllocationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_ocvp_byol_allocation.test_byol_allocation"
	datasourceName := "data.oci_ocvp_byol_allocations.test_byol_allocations"
	singularDatasourceName := "data.oci_ocvp_byol_allocation.test_byol_allocation"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OcvpByolAllocationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_ocvp_byol_allocation", "test_byol_allocation", acctest.Optional, acctest.Create, OcvpByolAllocationRepresentation), "ocvp", "byolAllocation", t)

	acctest.ResourceTest(t, testAccCheckOcvpByolAllocationDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + OcvpByolAllocationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_byol_allocation", "test_byol_allocation", acctest.Required, acctest.Create, OcvpByolAllocationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "allocated_units", "100"),
				resource.TestCheckResourceAttrSet(resourceName, "byol_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OcvpByolAllocationResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + OcvpByolAllocationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_byol_allocation", "test_byol_allocation", acctest.Optional, acctest.Create, OcvpByolAllocationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "allocated_units", "100"),
				resource.TestCheckResourceAttrSet(resourceName, "available_units"),
				resource.TestCheckResourceAttrSet(resourceName, "byol_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "entitlement_key"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "software_type"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_term_end"),
				resource.TestCheckResourceAttrSet(resourceName, "time_term_start"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + OcvpByolAllocationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_byol_allocation", "test_byol_allocation", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(OcvpByolAllocationRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "allocated_units", "100"),
				resource.TestCheckResourceAttrSet(resourceName, "available_units"),
				resource.TestCheckResourceAttrSet(resourceName, "byol_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "entitlement_key"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "software_type"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_term_end"),
				resource.TestCheckResourceAttrSet(resourceName, "time_term_start"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			Config: config + compartmentIdVariableStr + OcvpByolAllocationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_byol_allocation", "test_byol_allocation", acctest.Optional, acctest.Update, OcvpByolAllocationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "allocated_units", "110"),
				resource.TestCheckResourceAttrSet(resourceName, "available_units"),
				resource.TestCheckResourceAttrSet(resourceName, "byol_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "entitlement_key"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "software_type"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_term_end"),
				resource.TestCheckResourceAttrSet(resourceName, "time_term_start"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_byol_allocations", "test_byol_allocations", acctest.Optional, acctest.Update, OcvpByolAllocationDataSourceRepresentation) +
				compartmentIdVariableStr + OcvpByolAllocationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ocvp_byol_allocation", "test_byol_allocation", acctest.Optional, acctest.Update, OcvpByolAllocationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "available_units_greater_than_or_equal_to", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "byol_allocation_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "byol_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "software_type", "VCF"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "byol_allocation_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "byol_allocation_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_byol_allocation", "test_byol_allocation", acctest.Required, acctest.Create, OcvpByolAllocationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OcvpByolAllocationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "byol_allocation_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "allocated_units", "110"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "available_units"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "entitlement_key"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "software_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_term_end"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_term_start"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + OcvpByolAllocationRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckOcvpByolAllocationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).ByolAllocationClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_ocvp_byol_allocation" {
			noResourceFound = false
			request := oci_ocvp.GetByolAllocationRequest{}

			tmp := rs.Primary.ID
			request.ByolAllocationId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ocvp")

			response, err := client.GetByolAllocation(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_ocvp.ByolAllocationLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("OcvpByolAllocation") {
		resource.AddTestSweepers("OcvpByolAllocation", &resource.Sweeper{
			Name:         "OcvpByolAllocation",
			Dependencies: acctest.DependencyGraph["byolAllocation"],
			F:            sweepOcvpByolAllocationResource,
		})
	}
}

func sweepOcvpByolAllocationResource(compartment string) error {
	byolAllocationClient := acctest.GetTestClients(&schema.ResourceData{}).ByolAllocationClient()
	byolAllocationIds, err := getOcvpByolAllocationIds(compartment)
	if err != nil {
		return err
	}
	for _, byolAllocationId := range byolAllocationIds {
		if ok := acctest.SweeperDefaultResourceId[byolAllocationId]; !ok {
			deleteByolAllocationRequest := oci_ocvp.DeleteByolAllocationRequest{}

			deleteByolAllocationRequest.ByolAllocationId = &byolAllocationId

			deleteByolAllocationRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ocvp")
			_, error := byolAllocationClient.DeleteByolAllocation(context.Background(), deleteByolAllocationRequest)
			if error != nil {
				fmt.Printf("Error deleting ByolAllocation %s %s, It is possible that the resource is already deleted. Please verify manually \n", byolAllocationId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &byolAllocationId, OcvpByolAllocationSweepWaitCondition, time.Duration(3*time.Minute),
				OcvpByolAllocationSweepResponseFetchOperation, "ocvp", true)
		}
	}
	return nil
}

func getOcvpByolAllocationIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ByolAllocationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	byolAllocationClient := acctest.GetTestClients(&schema.ResourceData{}).ByolAllocationClient()

	listByolAllocationsRequest := oci_ocvp.ListByolAllocationsRequest{}
	listByolAllocationsRequest.CompartmentId = &compartmentId
	listByolAllocationsRequest.LifecycleState = oci_ocvp.ByolAllocationLifecycleStateActive
	listByolAllocationsResponse, err := byolAllocationClient.ListByolAllocations(context.Background(), listByolAllocationsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ByolAllocation list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, byolAllocation := range listByolAllocationsResponse.Items {
		id := *byolAllocation.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ByolAllocationId", id)
	}
	return resourceIds, nil
}

func OcvpByolAllocationSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if byolAllocationResponse, ok := response.Response.(oci_ocvp.GetByolAllocationResponse); ok {
		return byolAllocationResponse.LifecycleState != oci_ocvp.ByolAllocationLifecycleStateDeleted
	}
	return false
}

func OcvpByolAllocationSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.ByolAllocationClient().GetByolAllocation(context.Background(), oci_ocvp.GetByolAllocationRequest{
		ByolAllocationId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
