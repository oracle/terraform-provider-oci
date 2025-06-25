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
	oci_capacity_management "github.com/oracle/oci-go-sdk/v65/capacitymanagement"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CapacityManagementInternalOccmDemandSignalDeliveryRequiredOnlyResource = CapacityManagementInternalOccmDemandSignalDeliveryResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_internal_occm_demand_signal_delivery", "test_internal_occm_demand_signal_delivery", acctest.Required, acctest.Create, CapacityManagementInternalOccmDemandSignalDeliveryRepresentation)

	CapacityManagementInternalOccmDemandSignalDeliveryResourceConfig = CapacityManagementInternalOccmDemandSignalDeliveryResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_internal_occm_demand_signal_delivery", "test_internal_occm_demand_signal_delivery", acctest.Optional, acctest.Update, CapacityManagementInternalOccmDemandSignalDeliveryRepresentation)

	CapacityManagementInternalOccmDemandSignalDeliverySingularDataSourceRepresentation = map[string]interface{}{
		"occm_demand_signal_delivery_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_capacity_management_internal_occm_demand_signal_delivery.test_internal_occm_demand_signal_delivery.id}`},
	}

	CapacityManagementInternalOccmDemandSignalDeliveryDataSourceRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"occ_customer_group_id": acctest.Representation{RepType: acctest.Required, Create: `${var.customergroup_id}`},
		//"id":                         acctest.Representation{RepType: acctest.Optional, Create: `${var.demandsignal_id}`},
		"occm_demand_signal_item_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.demandsignalitem_id}`},
		"filter":                     acctest.RepresentationGroup{RepType: acctest.Required, Group: CapacityManagementInternalOccmDemandSignalDeliveryDataSourceFilterRepresentation}}
	CapacityManagementInternalOccmDemandSignalDeliveryDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${var.demandsignal_id}`}},
	}

	CapacityManagementInternalOccmDemandSignalDeliveryRepresentation = map[string]interface{}{
		"accepted_quantity":     acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"demand_signal_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.demandsignal_id}`},
		"demand_signal_item_id": acctest.Representation{RepType: acctest.Required, Create: `${var.demandsignalitem_id}`},
		"occ_customer_group_id": acctest.Representation{RepType: acctest.Required, Create: `${var.customergroup_id}`},
		//"defined_tags":          acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		//"freeform_tags":         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"justification": acctest.Representation{RepType: acctest.Optional, Create: `justification`, Update: `justification2`},
		"notes":         acctest.Representation{RepType: acctest.Optional, Create: `notes`, Update: `notes2`},
		//"lifecycle_details": acctest.Representation{RepType: acctest.Optional, Create: `IN_REVIEW`},
	}

	CapacityManagementInternalOccmDemandSignalDeliveryResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: capacity_management/default
func TestCapacityManagementInternalOccmDemandSignalDeliveryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCapacityManagementInternalOccmDemandSignalDeliveryResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("prod_compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	customerGroupId := utils.GetEnvSettingWithBlankDefault("customergroup_id")
	customerGroupIdVariableStr := fmt.Sprintf("variable \"customergroup_id\" { default = \"%s\" }\n", customerGroupId)

	demandsignalId := utils.GetEnvSettingWithBlankDefault("inprogress_demand_signal_id")
	demandsignalIdVariableStr := fmt.Sprintf("variable \"demandsignal_id\" { default = \"%s\" }\n", demandsignalId)

	demandsignalitemId := utils.GetEnvSettingWithBlankDefault("demand_signal_item_id")
	demandsignalitemIdVariableStr := fmt.Sprintf("variable \"demandsignalitem_id\" { default = \"%s\" }\n", demandsignalitemId)

	resourceName := "oci_capacity_management_internal_occm_demand_signal_delivery.test_internal_occm_demand_signal_delivery"
	datasourceName := "data.oci_capacity_management_internal_occm_demand_signal_deliveries.test_internal_occm_demand_signal_deliveries"
	//singularDatasourceName := "data.oci_capacity_management_internal_occm_demand_signal_delivery.test_internal_occm_demand_signal_delivery"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CapacityManagementInternalOccmDemandSignalDeliveryResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_internal_occm_demand_signal_delivery", "test_internal_occm_demand_signal_delivery", acctest.Optional, acctest.Create, CapacityManagementInternalOccmDemandSignalDeliveryRepresentation), "capacitymanagement", "internalOccmDemandSignalDelivery", t)

	acctest.ResourceTest(t, testAccCheckCapacityManagementInternalOccmDemandSignalDeliveryDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + customerGroupIdVariableStr + demandsignalIdVariableStr + demandsignalitemIdVariableStr + CapacityManagementInternalOccmDemandSignalDeliveryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_internal_occm_demand_signal_delivery", "test_internal_occm_demand_signal_delivery", acctest.Required, acctest.Create, CapacityManagementInternalOccmDemandSignalDeliveryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "accepted_quantity", "10"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "demand_signal_id"),
				resource.TestCheckResourceAttrSet(resourceName, "demand_signal_item_id"),
				resource.TestCheckResourceAttrSet(resourceName, "occ_customer_group_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + customerGroupIdVariableStr + demandsignalIdVariableStr + demandsignalitemIdVariableStr + CapacityManagementInternalOccmDemandSignalDeliveryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_internal_occm_demand_signal_delivery", "test_internal_occm_demand_signal_delivery", acctest.Optional, acctest.Create, CapacityManagementInternalOccmDemandSignalDeliveryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "accepted_quantity", "10"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "demand_signal_id"),
				resource.TestCheckResourceAttrSet(resourceName, "demand_signal_item_id"),
				//resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "justification", "justification"),
				resource.TestCheckResourceAttr(resourceName, "notes", "notes"),
				resource.TestCheckResourceAttrSet(resourceName, "occ_customer_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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
			Config: config + compartmentIdVariableStr + customerGroupIdVariableStr + demandsignalIdVariableStr + demandsignalitemIdVariableStr + CapacityManagementInternalOccmDemandSignalDeliveryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_internal_occm_demand_signal_delivery", "test_internal_occm_demand_signal_delivery", acctest.Optional, acctest.Update, CapacityManagementInternalOccmDemandSignalDeliveryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "accepted_quantity", "11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "demand_signal_id"),
				resource.TestCheckResourceAttrSet(resourceName, "demand_signal_item_id"),
				//resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "justification", "justification2"),
				resource.TestCheckResourceAttr(resourceName, "notes", "notes2"),
				resource.TestCheckResourceAttrSet(resourceName, "occ_customer_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				//func(s *terraform.State) (err error) {
				//	resId2, err = acctest.FromInstanceState(s, resourceName, "id")
				//	if resId != resId2 {
				//		return fmt.Errorf("Resource recreated when it was supposed to be updated.")
				//	}
				//	return err
				//},
			),
		},
		// verify datasource
		{
			Config: config + demandsignalitemIdVariableStr + demandsignalIdVariableStr + customerGroupIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_capacity_management_internal_occm_demand_signal_deliveries", "test_internal_occm_demand_signal_deliveries", acctest.Optional, acctest.Update, CapacityManagementInternalOccmDemandSignalDeliveryDataSourceRepresentation) +
				compartmentIdVariableStr + CapacityManagementInternalOccmDemandSignalDeliveryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_internal_occm_demand_signal_delivery", "test_internal_occm_demand_signal_delivery", acctest.Optional, acctest.Update, CapacityManagementInternalOccmDemandSignalDeliveryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				// resource.TestCheckResourceAttr(datasourceName, "id", "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "occ_customer_group_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "occm_demand_signal_item_id"),

				resource.TestCheckResourceAttr(datasourceName, "internal_occm_demand_signal_delivery_collection.#", "1"),
				//resource.TestCheckResourceAttr(datasourceName, "internal_occm_demand_signal_delivery_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		//{
		//	Config: config +
		//		acctest.GenerateDataSourceFromRepresentationMap("oci_capacity_management_internal_occm_demand_signal_delivery", "test_internal_occm_demand_signal_delivery", acctest.Required, acctest.Create, CapacityManagementInternalOccmDemandSignalDeliverySingularDataSourceRepresentation) +
		//		compartmentIdVariableStr + demandsignalIdVariableStr + customerGroupIdVariableStr + demandsignalitemIdVariableStr + CapacityManagementInternalOccmDemandSignalDeliveryResourceConfig,
		//	Check: acctest.ComposeAggregateTestCheckFuncWrapper(
		//		resource.TestCheckResourceAttrSet(singularDatasourceName, "occm_demand_signal_delivery_id"),
		//
		//		resource.TestCheckResourceAttr(singularDatasourceName, "accepted_quantity", "11"),
		//		resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
		//		//resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
		//		//resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
		//		resource.TestCheckResourceAttr(singularDatasourceName, "justification", "justification2"),
		//		resource.TestCheckResourceAttr(singularDatasourceName, "notes", "notes2"),
		//		resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
		//		resource.TestCheckResourceAttrSet(singularDatasourceName, "time_delivered"),
		//	),
		//},
		// verify resource import
		{
			Config:                  config + CapacityManagementInternalOccmDemandSignalDeliveryRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCapacityManagementInternalOccmDemandSignalDeliveryDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).InternalDemandSignalClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_capacity_management_internal_occm_demand_signal_delivery" {
			noResourceFound = false
			request := oci_capacity_management.GetInternalOccmDemandSignalDeliveryRequest{}

			if rs.Primary.ID == "" {
				return fmt.Errorf("resource ID is empty for resource %s", rs.Primary.ID)
			}
			request.OccmDemandSignalDeliveryId = &rs.Primary.ID

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "capacity_management")

			response, err := client.GetInternalOccmDemandSignalDelivery(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_capacity_management.InternalOccmDemandSignalDeliveryLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("CapacityManagementInternalOccmDemandSignalDelivery") {
		resource.AddTestSweepers("CapacityManagementInternalOccmDemandSignalDelivery", &resource.Sweeper{
			Name:         "CapacityManagementInternalOccmDemandSignalDelivery",
			Dependencies: acctest.DependencyGraph["internalOccmDemandSignalDelivery"],
			F:            sweepCapacityManagementInternalOccmDemandSignalDeliveryResource,
		})
	}
}

func sweepCapacityManagementInternalOccmDemandSignalDeliveryResource(compartment string) error {
	internalDemandSignalClient := acctest.GetTestClients(&schema.ResourceData{}).InternalDemandSignalClient()
	internalOccmDemandSignalDeliveryIds, err := getCapacityManagementInternalOccmDemandSignalDeliveryIds(compartment)
	if err != nil {
		return err
	}
	for _, internalOccmDemandSignalDeliveryId := range internalOccmDemandSignalDeliveryIds {
		if ok := acctest.SweeperDefaultResourceId[internalOccmDemandSignalDeliveryId]; !ok {
			deleteInternalOccmDemandSignalDeliveryRequest := oci_capacity_management.DeleteInternalOccmDemandSignalDeliveryRequest{}

			deleteInternalOccmDemandSignalDeliveryRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "capacity_management")
			_, error := internalDemandSignalClient.DeleteInternalOccmDemandSignalDelivery(context.Background(), deleteInternalOccmDemandSignalDeliveryRequest)
			if error != nil {
				fmt.Printf("Error deleting InternalOccmDemandSignalDelivery %s %s, It is possible that the resource is already deleted. Please verify manually \n", internalOccmDemandSignalDeliveryId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &internalOccmDemandSignalDeliveryId, CapacityManagementInternalOccmDemandSignalDeliverySweepWaitCondition, time.Duration(3*time.Minute),
				CapacityManagementInternalOccmDemandSignalDeliverySweepResponseFetchOperation, "capacity_management", true)
		}
	}
	return nil
}

func getCapacityManagementInternalOccmDemandSignalDeliveryIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "InternalOccmDemandSignalDeliveryId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	internalDemandSignalClient := acctest.GetTestClients(&schema.ResourceData{}).InternalDemandSignalClient()

	listInternalOccmDemandSignalDeliveriesRequest := oci_capacity_management.ListInternalOccmDemandSignalDeliveriesRequest{}
	listInternalOccmDemandSignalDeliveriesRequest.CompartmentId = &compartmentId

	occCustomerGroupIds, error := getCapacityManagementOccCustomerGroupIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting occCustomerGroupId required for InternalOccmDemandSignalDelivery resource requests \n")
	}
	for _, occCustomerGroupId := range occCustomerGroupIds {
		listInternalOccmDemandSignalDeliveriesRequest.OccCustomerGroupId = &occCustomerGroupId

		listInternalOccmDemandSignalDeliveriesResponse, err := internalDemandSignalClient.ListInternalOccmDemandSignalDeliveries(context.Background(), listInternalOccmDemandSignalDeliveriesRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting InternalOccmDemandSignalDelivery list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, internalOccmDemandSignalDelivery := range listInternalOccmDemandSignalDeliveriesResponse.Items {
			id := *internalOccmDemandSignalDelivery.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "InternalOccmDemandSignalDeliveryId", id)
		}

	}
	return resourceIds, nil
}

func CapacityManagementInternalOccmDemandSignalDeliverySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if internalOccmDemandSignalDeliveryResponse, ok := response.Response.(oci_capacity_management.GetInternalOccmDemandSignalDeliveryResponse); ok {
		return internalOccmDemandSignalDeliveryResponse.LifecycleState != oci_capacity_management.InternalOccmDemandSignalDeliveryLifecycleStateDeleted
	}
	return false
}

func CapacityManagementInternalOccmDemandSignalDeliverySweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	request := oci_capacity_management.GetInternalOccmDemandSignalDeliveryRequest{
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	}
	if resourceId != nil {
		request.OccmDemandSignalDeliveryId = resourceId
	} else {
		return fmt.Errorf("resourceId is nil in CapacityManagementInternalOccmDemandSignalDeliverySweepResponseFetchOperation")
	}

	_, err := client.InternalDemandSignalClient().GetInternalOccmDemandSignalDelivery(context.Background(), request)
	return err
}
