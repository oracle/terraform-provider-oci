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
	oci_capacity_management "github.com/oracle/oci-go-sdk/v65/capacitymanagement"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CapacityManagementOccCapacityRequestRequiredOnlyResource = CapacityManagementOccCapacityRequestResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_occ_capacity_request", "test_occ_capacity_request", acctest.Required, acctest.Create, CapacityManagementOccCapacityRequestRepresentation)

	CapacityManagementOccCapacityRequestSingularDataSourceRepresentation = map[string]interface{}{
		"occ_capacity_request_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_capacity_management_occ_capacity_request.test_occ_capacity_request.id}`},
	}

	CapacityManagementOccCapacityRequestDataSourceRepresentation = map[string]interface{}{
		"compartment_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                acctest.Representation{RepType: acctest.Optional, Create: `UI test request`},
		"id":                          acctest.Representation{RepType: acctest.Optional, Create: `${oci_capacity_management_occ_capacity_request.test_occ_capacity_request.id}`},
		"namespace":                   acctest.Representation{RepType: acctest.Optional, Create: `COMPUTE`},
		"occ_availability_catalog_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.occ_availability_catalog_id}`},
		"filter":                      acctest.RepresentationGroup{RepType: acctest.Required, Group: CapacityManagementOccCapacityRequestDataSourceFilterRepresentation}}
	CapacityManagementOccCapacityRequestDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_capacity_management_occ_capacity_request.test_occ_capacity_request.id}`}},
	}

	CapacityManagementOccCapacityRequestRepresentation = map[string]interface{}{
		"availability_domain":             acctest.Representation{RepType: acctest.Required, Create: `US-ASHBURN-1-AD-2`},
		"compartment_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"date_expected_capacity_handover": acctest.Representation{RepType: acctest.Required, Create: `2025-04-05T00:00:00.000Z`},
		"details":                         acctest.RepresentationGroup{RepType: acctest.Required, Group: CapacityManagementOccCapacityRequestDetailsRepresentation},
		"display_name":                    acctest.Representation{RepType: acctest.Required, Create: `UI test request`, Update: `displayName2`},
		"namespace":                       acctest.Representation{RepType: acctest.Required, Create: `COMPUTE`},
		"occ_availability_catalog_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.occ_availability_catalog_id}`},
		"region":                          acctest.Representation{RepType: acctest.Required, Create: `US-ASHBURN-1`},
		"description":                     acctest.Representation{RepType: acctest.Optional, Create: `This is the test request created for UI`},
		"freeform_tags":                   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"request_state":                   acctest.Representation{RepType: acctest.Optional, Create: `CREATED`},
	}
	CapacityManagementOccCapacityRequestDetailsRepresentation = map[string]interface{}{
		"demand_quantity": acctest.Representation{RepType: acctest.Required, Create: `10`},
		"resource_name":   acctest.Representation{RepType: acctest.Required, Create: `BM.Standard3.64`},
		"resource_type":   acctest.Representation{RepType: acctest.Required, Create: `SERVER_HW`},
		"workload_type":   acctest.Representation{RepType: acctest.Required, Create: `US_PROD`},
	}

	CapacityManagementOccCapacityRequestResourceDependencies = ""
)

// issue-routing-tag: capacity_management/default
func TestCapacityManagementOccCapacityRequestResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCapacityManagementOccCapacityRequestResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	occAvailabilityCatalogId := utils.GetEnvSettingWithBlankDefault("occ_availability_catalog_ocid")
	occAvailabilityCatalogIdVariableStr := fmt.Sprintf("variable \"occ_availability_catalog_id\" { default = \"%s\" }\n", occAvailabilityCatalogId)

	resourceName := "oci_capacity_management_occ_capacity_request.test_occ_capacity_request"
	singularDatasourceName := "data.oci_capacity_management_occ_capacity_request.test_occ_capacity_request"

	//var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CapacityManagementOccCapacityRequestResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_occ_capacity_request", "test_occ_capacity_request", acctest.Optional, acctest.Create, CapacityManagementOccCapacityRequestRepresentation), "capacitymanagement", "occCapacityRequest", t)

	acctest.ResourceTest(t, testAccCheckCapacityManagementOccCapacityRequestDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + occAvailabilityCatalogIdVariableStr + CapacityManagementOccCapacityRequestResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_occ_capacity_request", "test_occ_capacity_request", acctest.Required, acctest.Create, CapacityManagementOccCapacityRequestRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "date_expected_capacity_handover"),
				resource.TestCheckResourceAttr(resourceName, "details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "details.0.demand_quantity", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "details.0.resource_name"),
				resource.TestCheckResourceAttr(resourceName, "details.0.resource_type", "SERVER_HW"),
				resource.TestCheckResourceAttr(resourceName, "details.0.workload_type", "US_PROD"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "UI test request"),
				resource.TestCheckResourceAttr(resourceName, "namespace", "COMPUTE"),
				resource.TestCheckResourceAttrSet(resourceName, "occ_availability_catalog_id"),
				resource.TestCheckResourceAttr(resourceName, "region", "US-ASHBURN-1"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + occAvailabilityCatalogIdVariableStr + CapacityManagementOccCapacityRequestResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + occAvailabilityCatalogIdVariableStr + CapacityManagementOccCapacityRequestResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_occ_capacity_request", "test_occ_capacity_request", acctest.Optional, acctest.Create, CapacityManagementOccCapacityRequestRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "date_expected_capacity_handover"),
				resource.TestCheckResourceAttr(resourceName, "description", "This is the test request created for UI"),
				resource.TestCheckResourceAttr(resourceName, "details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "details.0.demand_quantity", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "details.0.resource_name"),
				resource.TestCheckResourceAttr(resourceName, "details.0.resource_type", "SERVER_HW"),
				resource.TestCheckResourceAttr(resourceName, "details.0.workload_type", "US_PROD"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "UI test request"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "namespace", "COMPUTE"),
				resource.TestCheckResourceAttrSet(resourceName, "occ_availability_catalog_id"),
				resource.TestCheckResourceAttrSet(resourceName, "occ_customer_group_id"),
				resource.TestCheckResourceAttr(resourceName, "region", "US-ASHBURN-1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
			),
		},

		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_capacity_management_occ_capacity_request", "test_occ_capacity_request", acctest.Required, acctest.Create, CapacityManagementOccCapacityRequestSingularDataSourceRepresentation) +
				compartmentIdVariableStr + occAvailabilityCatalogIdVariableStr + CapacityManagementOccCapacityRequestRequiredOnlyResource,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "date_expected_capacity_handover"),
				resource.TestCheckResourceAttr(singularDatasourceName, "details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "details.0.demand_quantity", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "details.0.resource_type", "SERVER_HW"),
				resource.TestCheckResourceAttr(singularDatasourceName, "details.0.workload_type", "US_PROD"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "namespace", "COMPUTE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "occ_customer_group_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "region", "US-ASHBURN-1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + occAvailabilityCatalogIdVariableStr + CapacityManagementOccCapacityRequestRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"patch_operations",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckCapacityManagementOccCapacityRequestDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).CapacityManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_capacity_management_occ_capacity_request" {
			noResourceFound = false
			request := oci_capacity_management.GetOccCapacityRequestRequest{}

			tmp := rs.Primary.ID
			request.OccCapacityRequestId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "capacity_management")

			response, err := client.GetOccCapacityRequest(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_capacity_management.OccCapacityRequestLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("CapacityManagementOccCapacityRequest") {
		resource.AddTestSweepers("CapacityManagementOccCapacityRequest", &resource.Sweeper{
			Name:         "CapacityManagementOccCapacityRequest",
			Dependencies: acctest.DependencyGraph["occCapacityRequest"],
			F:            sweepCapacityManagementOccCapacityRequestResource,
		})
	}
}

func sweepCapacityManagementOccCapacityRequestResource(compartment string) error {
	capacityManagementClient := acctest.GetTestClients(&schema.ResourceData{}).CapacityManagementClient()
	occCapacityRequestIds, err := getCapacityManagementOccCapacityRequestIds(compartment)
	if err != nil {
		return err
	}
	for _, occCapacityRequestId := range occCapacityRequestIds {
		if ok := acctest.SweeperDefaultResourceId[occCapacityRequestId]; !ok {
			deleteOccCapacityRequestRequest := oci_capacity_management.DeleteOccCapacityRequestRequest{}

			deleteOccCapacityRequestRequest.OccCapacityRequestId = &occCapacityRequestId

			deleteOccCapacityRequestRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "capacity_management")
			_, error := capacityManagementClient.DeleteOccCapacityRequest(context.Background(), deleteOccCapacityRequestRequest)
			if error != nil {
				fmt.Printf("Error deleting OccCapacityRequest %s %s, It is possible that the resource is already deleted. Please verify manually \n", occCapacityRequestId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &occCapacityRequestId, CapacityManagementOccCapacityRequestSweepWaitCondition, time.Duration(3*time.Minute),
				CapacityManagementOccCapacityRequestSweepResponseFetchOperation, "capacity_management", true)
		}
	}
	return nil
}

func getCapacityManagementOccCapacityRequestIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "OccCapacityRequestId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	capacityManagementClient := acctest.GetTestClients(&schema.ResourceData{}).CapacityManagementClient()

	listOccCapacityRequestsRequest := oci_capacity_management.ListOccCapacityRequestsRequest{}
	listOccCapacityRequestsRequest.CompartmentId = &compartmentId
	listOccCapacityRequestsResponse, err := capacityManagementClient.ListOccCapacityRequests(context.Background(), listOccCapacityRequestsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting OccCapacityRequest list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, occCapacityRequest := range listOccCapacityRequestsResponse.Items {
		id := *occCapacityRequest.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "OccCapacityRequestId", id)
	}
	return resourceIds, nil
}

func CapacityManagementOccCapacityRequestSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if occCapacityRequestResponse, ok := response.Response.(oci_capacity_management.GetOccCapacityRequestResponse); ok {
		return occCapacityRequestResponse.LifecycleState != oci_capacity_management.OccCapacityRequestLifecycleStateDeleted
	}
	return false
}

func CapacityManagementOccCapacityRequestSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.CapacityManagementClient().GetOccCapacityRequest(context.Background(), oci_capacity_management.GetOccCapacityRequestRequest{
		OccCapacityRequestId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
