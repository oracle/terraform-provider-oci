// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	oci_apiaccesscontrol "github.com/oracle/oci-go-sdk/v65/apiaccesscontrol"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ApiaccesscontrolPrivilegedApiControlRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_apiaccesscontrol_privileged_api_control", "test_privileged_api_control", acctest.Required, acctest.Create, ApiaccesscontrolPrivilegedApiControlRepresentation)

	ApiaccesscontrolPrivilegedApiControlResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_apiaccesscontrol_privileged_api_control", "test_privileged_api_control", acctest.Optional, acctest.Update, ApiaccesscontrolPrivilegedApiControlRepresentation)

	ApiaccesscontrolPrivilegedApiControlSingularDataSourceRepresentation = map[string]interface{}{
		"privileged_api_control_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apiaccesscontrol_privileged_api_control.test_privileged_api_control.id}`},
	}

	ApiaccesscontrolPrivilegedApiControlDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `terraformprivilegedapicontrolUpdated`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_apiaccesscontrol_privileged_api_control.test_privileged_api_control.id}`},
		"resource_type":  acctest.Representation{RepType: acctest.Optional, Create: `EXADATAINFRASTRUCTURE`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: ApiaccesscontrolPrivilegedApiControlDataSourceFilterRepresentation}}
	ApiaccesscontrolPrivilegedApiControlDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_apiaccesscontrol_privileged_api_control.test_privileged_api_control.id}`}},
	}

	ApiaccesscontrolPrivilegedApiControlRepresentation = map[string]interface{}{
		"approver_group_id_list":    acctest.Representation{RepType: acctest.Required, Create: []string{`use_iam_policy`}},
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"notification_topic_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.topic_id}`},
		"privileged_operation_list": acctest.RepresentationGroup{RepType: acctest.Required, Group: ApiaccesscontrolPrivilegedApiControlPrivilegedOperationListRepresentation},
		"resource_type":             acctest.Representation{RepType: acctest.Required, Create: `EXADATAINFRASTRUCTURE`},
		"resources":                 acctest.Representation{RepType: acctest.Required, Create: []string{`${var.resource_id}`}},
		"description":               acctest.Representation{RepType: acctest.Required, Create: `TerraformPrivilegedApiControl test create`},
		"display_name":              acctest.Representation{RepType: acctest.Required, Create: `terraformprivilegedapicontrol`, Update: `terraformprivilegedapicontrolUpdated`},
		"number_of_approvers":       acctest.Representation{RepType: acctest.Required, Create: `1`},
	}
	ApiaccesscontrolPrivilegedApiControlPrivilegedOperationListRepresentation = map[string]interface{}{
		"api_name":        acctest.Representation{RepType: acctest.Required, Create: `UpdateVmCluster`},
		"attribute_names": acctest.Representation{RepType: acctest.Optional, Create: []string{}},
		"entity_type":     acctest.Representation{RepType: acctest.Required, Create: `DbaasExadataVmCluster`},
	}
)

// issue-routing-tag: apiaccesscontrol/default
func TestApiaccesscontrolPrivilegedApiControlResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApiaccesscontrolPrivilegedApiControlResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceId := utils.GetEnvSettingWithBlankDefault("resource_id")
	resourceIdVariableStr := fmt.Sprintf("variable \"resource_id\" { default = \"%s\" }\n", resourceId)

	onsTopicId := utils.GetEnvSettingWithBlankDefault("topic_id")
	onsTopicIdVariableStr := fmt.Sprintf("variable \"topic_id\" { default = \"%s\" }\n", onsTopicId)

	resourceName := "oci_apiaccesscontrol_privileged_api_control.test_privileged_api_control"
	datasourceName := "data.oci_apiaccesscontrol_privileged_api_controls.test_privileged_api_controls"
	singularDatasourceName := "data.oci_apiaccesscontrol_privileged_api_control.test_privileged_api_control"

	var resId, resId2 string
	//var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+resourceIdVariableStr+onsTopicIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_apiaccesscontrol_privileged_api_control", "test_privileged_api_control", acctest.Optional, acctest.Create, ApiaccesscontrolPrivilegedApiControlRepresentation), "apiaccesscontrol", "privilegedApiControl", t)

	acctest.ResourceTest(t, testAccCheckApiaccesscontrolPrivilegedApiControlDestroy, []resource.TestStep{
		//verify Create
		{
			Config: config + compartmentIdVariableStr + resourceIdVariableStr + onsTopicIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_apiaccesscontrol_privileged_api_control", "test_privileged_api_control", acctest.Required, acctest.Create, ApiaccesscontrolPrivilegedApiControlRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "approver_group_id_list.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "notification_topic_id"),
				resource.TestCheckResourceAttr(resourceName, "privileged_operation_list.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "privileged_operation_list.0.api_name"),
				resource.TestCheckResourceAttr(resourceName, "resource_type", "EXADATAINFRASTRUCTURE"),
				resource.TestCheckResourceAttr(resourceName, "resources.#", "1"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if err != nil {
						return err
					}
					err2 := dummyDeleteCallApiaccesscontrolPrivilegedApiControlResource(compartmentId, resId)
					if err2 != nil {
						return err2
					}
					err3 := dummyUpdateCallApiaccesscontrolPrivilegedApiControlResource(compartmentId, resId)
					if err3 != nil {
						return err3
					}
					return approveApiaccesscontrolPrivilegedApiRequestResources(compartmentId, resId)

				},
			),
		},

		// delete before next Create
		//{
		//	Config: config + compartmentIdVariableStr,
		//},
		// verify Create with optionals
		//{
		//	Config: config + compartmentIdVariableStr  +
		//		acctest.GenerateResourceFromRepresentationMap("oci_apiaccesscontrol_privileged_api_control", "test_privileged_api_control", acctest.Optional, acctest.Create, ApiaccesscontrolPrivilegedApiControlRepresentation),
		//	Check: acctest.ComposeAggregateTestCheckFuncWrapper(
		//		resource.TestCheckResourceAttr(resourceName, "approver_group_id_list.#", "1"),
		//		resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
		//		resource.TestCheckResourceAttr(resourceName, "description", "Control for pre approving the apis"),
		//		resource.TestCheckResourceAttr(resourceName, "display_name", "TestPrivilegedApiControl"),
		//		resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
		//		resource.TestCheckResourceAttrSet(resourceName, "id"),
		//		resource.TestCheckResourceAttrSet(resourceName, "notification_topic_id"),
		//		resource.TestCheckResourceAttr(resourceName, "number_of_approvers", "10"),
		//		resource.TestCheckResourceAttr(resourceName, "privileged_operation_list.#", "1"),
		//		resource.TestCheckResourceAttrSet(resourceName, "privileged_operation_list.0.api_name"),
		//		resource.TestCheckResourceAttr(resourceName, "privileged_operation_list.0.attribute_names.#", "1"),
		//		resource.TestCheckResourceAttr(resourceName, "privileged_operation_list.0.entity_type", "DbaasVmCluster"),
		//		resource.TestCheckResourceAttr(resourceName, "resource_type", "EXADATAINFRASTRUCTURE"),
		//		resource.TestCheckResourceAttr(resourceName, "resources.#", "1"),
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
		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + resourceIdVariableStr + onsTopicIdVariableStr + compartmentIdVariableStr + compartmentIdUVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_apiaccesscontrol_privileged_api_control", "test_privileged_api_control", acctest.Required, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(ApiaccesscontrolPrivilegedApiControlRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "approver_group_id_list.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "TerraformPrivilegedApiControl test create"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "terraformprivilegedapicontrol"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "notification_topic_id"),
				resource.TestCheckResourceAttr(resourceName, "number_of_approvers", "1"),
				resource.TestCheckResourceAttr(resourceName, "privileged_operation_list.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "privileged_operation_list.0.api_name"),
				resource.TestCheckResourceAttr(resourceName, "privileged_operation_list.0.entity_type", "DbaasExadataVmCluster"),
				resource.TestCheckResourceAttr(resourceName, "resource_type", "EXADATAINFRASTRUCTURE"),
				resource.TestCheckResourceAttr(resourceName, "resources.#", "1"),
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
			Config: config + resourceIdVariableStr + onsTopicIdVariableStr + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_apiaccesscontrol_privileged_api_control", "test_privileged_api_control", acctest.Required, acctest.Update, ApiaccesscontrolPrivilegedApiControlRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "approver_group_id_list.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "TerraformPrivilegedApiControl test create"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "terraformprivilegedapicontrolUpdated"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "notification_topic_id"),
				resource.TestCheckResourceAttr(resourceName, "number_of_approvers", "1"),
				resource.TestCheckResourceAttr(resourceName, "privileged_operation_list.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "privileged_operation_list.0.api_name"),
				resource.TestCheckResourceAttr(resourceName, "privileged_operation_list.0.entity_type", "DbaasExadataVmCluster"),
				resource.TestCheckResourceAttr(resourceName, "resource_type", "EXADATAINFRASTRUCTURE"),
				resource.TestCheckResourceAttr(resourceName, "resources.#", "1"),
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
			Config: config + onsTopicIdVariableStr + resourceIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apiaccesscontrol_privileged_api_controls", "test_privileged_api_controls", acctest.Optional, acctest.Update, ApiaccesscontrolPrivilegedApiControlDataSourceRepresentation) +
				compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_apiaccesscontrol_privileged_api_control", "test_privileged_api_control", acctest.Optional, acctest.Update, ApiaccesscontrolPrivilegedApiControlRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "terraformprivilegedapicontrolUpdated"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "resource_type", "EXADATAINFRASTRUCTURE"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "privileged_api_control_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "privileged_api_control_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config + onsTopicIdVariableStr + resourceIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apiaccesscontrol_privileged_api_control", "test_privileged_api_control", acctest.Required, acctest.Create, ApiaccesscontrolPrivilegedApiControlSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ApiaccesscontrolPrivilegedApiControlResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "privileged_api_control_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "approver_group_id_list.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "TerraformPrivilegedApiControl test create"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "terraformprivilegedapicontrolUpdated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "number_of_approvers", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "privileged_operation_list.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "privileged_operation_list.0.entity_type", "DbaasExadataVmCluster"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resource_type", "EXADATAINFRASTRUCTURE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resources.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state_details"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + ApiaccesscontrolPrivilegedApiControlRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckApiaccesscontrolPrivilegedApiControlDestroy(s *terraform.State) error {
	println("testAccCheckApiaccesscontrolPrivilegedApiControlDestroy")
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).PrivilegedApiControlClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_apiaccesscontrol_privileged_api_control" {
			noResourceFound = false
			request := oci_apiaccesscontrol.GetPrivilegedApiControlRequest{}

			tmp := rs.Primary.ID
			request.PrivilegedApiControlId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "apiaccesscontrol")

			response, err := client.GetPrivilegedApiControl(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_apiaccesscontrol.PrivilegedApiControlLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("ApiaccesscontrolPrivilegedApiControl") {
		resource.AddTestSweepers("ApiaccesscontrolPrivilegedApiControl", &resource.Sweeper{
			Name:         "ApiaccesscontrolPrivilegedApiControl",
			Dependencies: acctest.DependencyGraph["privilegedApiControl"],
			F:            sweepApiaccesscontrolPrivilegedApiControlResource,
		})
	}
}

func approveApiaccesscontrolPrivilegedApiRequestResources(compartment string, clusterId string) error {
	privilegedApiRequestClient := acctest.GetTestClients(&schema.ResourceData{}).PrivilegedApiRequestsClient()
	privilegedApiRequestIds, err := getApiaccesscontrolPrivilegedApiRequestIds(compartment, clusterId)
	if err != nil {
		return err
	}
	for _, privilegedApiRequestId := range privilegedApiRequestIds {

		approvePrivilegedApiRequest := oci_apiaccesscontrol.ApprovePrivilegedApiRequestRequest{}

		approvePrivilegedApiRequest.PrivilegedApiRequestId = &privilegedApiRequestId
		approvePrivilegedApiRequest.ApprovePrivilegedApiRequestDetails = oci_apiaccesscontrol.ApprovePrivilegedApiRequestDetails{
			ApproverComment: common.String("Terraform Approved"),
			// Set other fields if needed
		}
		approvePrivilegedApiRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "apiaccesscontrol")
		_, error := privilegedApiRequestClient.ApprovePrivilegedApiRequest(context.Background(), approvePrivilegedApiRequest)
		if error != nil {
			fmt.Printf("Error approving privilegedApiRequestId %s %s, It is possible that the resource is already approved. Please verify manually \n", privilegedApiRequestId, error)
			continue
		}
		acctest.WaitTillCondition(acctest.TestAccProvider, &privilegedApiRequestId, ApiaccesscontrolPrivilegedApiRequestApprovedWaitCondition, time.Duration(3*time.Minute),
			ApiaccesscontrolPrivilegedApiRequestResponseFetchOperation, "apiaccesscontrol", true)

	}
	return nil
}

func sweepApiaccesscontrolPrivilegedApiControlResource(compartment string) error {
	privilegedApiControlClient := acctest.GetTestClients(&schema.ResourceData{}).PrivilegedApiControlClient()
	privilegedApiControlIds, err := getApiaccesscontrolPrivilegedApiControlIds(compartment)
	if err != nil {
		return err
	}
	for _, privilegedApiControlId := range privilegedApiControlIds {
		if ok := acctest.SweeperDefaultResourceId[privilegedApiControlId]; !ok {
			deletePrivilegedApiControlRequest := oci_apiaccesscontrol.DeletePrivilegedApiControlRequest{}

			deletePrivilegedApiControlRequest.PrivilegedApiControlId = &privilegedApiControlId

			deletePrivilegedApiControlRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "apiaccesscontrol")
			_, error := privilegedApiControlClient.DeletePrivilegedApiControl(context.Background(), deletePrivilegedApiControlRequest)
			if error != nil {
				fmt.Printf("Error deleting PrivilegedApiControl %s %s, It is possible that the resource is already deleted. Please verify manually \n", privilegedApiControlId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &privilegedApiControlId, ApiaccesscontrolPrivilegedApiControlSweepWaitCondition, time.Duration(3*time.Minute),
				ApiaccesscontrolPrivilegedApiControlSweepResponseFetchOperation, "apiaccesscontrol", true)
		}
	}
	return nil
}

func dummyDeleteCallApiaccesscontrolPrivilegedApiControlResource(compartment string, controlId string) error {
	privilegedApiControlClient := acctest.GetTestClients(&schema.ResourceData{}).PrivilegedApiControlClient()
	privilegedApiControlIds := []string{controlId}
	for _, privilegedApiControlId := range privilegedApiControlIds {
		deletePrivilegedApiControlRequest := oci_apiaccesscontrol.DeletePrivilegedApiControlRequest{}

		deletePrivilegedApiControlRequest.PrivilegedApiControlId = &privilegedApiControlId

		deletePrivilegedApiControlRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "apiaccesscontrol")
		_, error := privilegedApiControlClient.DeletePrivilegedApiControl(context.Background(), deletePrivilegedApiControlRequest)
		if error != nil {
			fmt.Printf("Error deleting PrivilegedApiControl %s %s, It is possible that the resource is already deleted. Please verify manually \n", privilegedApiControlId, error)
			continue
		}

	}
	return nil
}

func dummyUpdateCallApiaccesscontrolPrivilegedApiControlResource(compartment string, controlId string) error {
	privilegedApiControlClient := acctest.GetTestClients(&schema.ResourceData{}).PrivilegedApiControlClient()
	privilegedApiControlIds := []string{controlId}
	for _, privilegedApiControlId := range privilegedApiControlIds {
		updatePrivilegedApiControlRequest := oci_apiaccesscontrol.UpdatePrivilegedApiControlRequest{}

		updatePrivilegedApiControlRequest.PrivilegedApiControlId = &privilegedApiControlId

		updatePrivilegedApiControlRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "apiaccesscontrol")
		_, error := privilegedApiControlClient.UpdatePrivilegedApiControl(context.Background(), updatePrivilegedApiControlRequest)
		if error != nil {
			fmt.Printf("Error updating PrivilegedApiControl %s %s, It is possible that the resource is already updated. Please verify manually \n", privilegedApiControlId, error)
			continue
		}

	}
	return nil
}

func getApiaccesscontrolPrivilegedApiRequestIds(compartment string, resourceId string) ([]string, error) {
	var resourceIds []string
	compartmentId := compartment
	privilegedApiRequestClient := acctest.GetTestClients(&schema.ResourceData{}).PrivilegedApiRequestsClient()

	listPrivilegedApiRequestsRequest := oci_apiaccesscontrol.ListPrivilegedApiRequestsRequest{}
	listPrivilegedApiRequestsRequest.CompartmentId = &compartmentId
	listPrivilegedApiRequestsRequest.State = oci_apiaccesscontrol.PrivilegedApiRequestStateApprovalWaiting
	listPrivilegedApiRequestsRequest.ResourceId = &resourceId
	listPrivilegedApiRequestsResponse, err := privilegedApiRequestClient.ListPrivilegedApiRequests(context.Background(), listPrivilegedApiRequestsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting PrivilegedApiRequests list for compartment id : %s state: %s, %s \n", compartmentId, `approval waiting`, err)
	}
	for _, privilegedApiRequest := range listPrivilegedApiRequestsResponse.Items {
		id := *privilegedApiRequest.Id
		resourceIds = append(resourceIds, id)
	}
	return resourceIds, nil
}

func getApiaccesscontrolPrivilegedApiControlIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "PrivilegedApiControlId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	privilegedApiControlClient := acctest.GetTestClients(&schema.ResourceData{}).PrivilegedApiControlClient()

	listPrivilegedApiControlsRequest := oci_apiaccesscontrol.ListPrivilegedApiControlsRequest{}
	listPrivilegedApiControlsRequest.CompartmentId = &compartmentId
	listPrivilegedApiControlsRequest.LifecycleState = oci_apiaccesscontrol.PrivilegedApiControlLifecycleStateNeedsAttention
	listPrivilegedApiControlsResponse, err := privilegedApiControlClient.ListPrivilegedApiControls(context.Background(), listPrivilegedApiControlsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting PrivilegedApiControl list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, privilegedApiControl := range listPrivilegedApiControlsResponse.Items {
		id := *privilegedApiControl.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "PrivilegedApiControlId", id)
	}
	return resourceIds, nil
}

func ApiaccesscontrolPrivilegedApiRequestApprovedWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if privilegedApiRequestResponse, ok := response.Response.(oci_apiaccesscontrol.GetPrivilegedApiRequestResponse); ok {
		return privilegedApiRequestResponse.State != oci_apiaccesscontrol.PrivilegedApiRequestStateApproved
	}
	return false
}

func ApiaccesscontrolPrivilegedApiRequestResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.PrivilegedApiRequestsClient().GetPrivilegedApiRequest(context.Background(), oci_apiaccesscontrol.GetPrivilegedApiRequestRequest{
		PrivilegedApiRequestId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}

func ApiaccesscontrolPrivilegedApiControlSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if privilegedApiControlResponse, ok := response.Response.(oci_apiaccesscontrol.GetPrivilegedApiControlResponse); ok {
		return privilegedApiControlResponse.LifecycleState != oci_apiaccesscontrol.PrivilegedApiControlLifecycleStateDeleted
	}
	return false
}

func ApiaccesscontrolPrivilegedApiControlSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.PrivilegedApiControlClient().GetPrivilegedApiControl(context.Background(), oci_apiaccesscontrol.GetPrivilegedApiControlRequest{
		PrivilegedApiControlId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
