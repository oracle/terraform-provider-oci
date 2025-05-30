// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	oci_apiaccesscontrol "github.com/oracle/oci-go-sdk/v65/apiaccesscontrol"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ApiaccesscontrolPrivilegedApiRequestRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_apiaccesscontrol_privileged_api_request", "test_privileged_api_request", acctest.Required, acctest.Create, ApiaccesscontrolPrivilegedApiRequestRepresentation)

	ApiaccesscontrolPrivilegedApiRequestResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_apiaccesscontrol_privileged_api_request", "test_privileged_api_request", acctest.Optional, acctest.Update, ApiaccesscontrolPrivilegedApiRequestRepresentation)

	ApiaccesscontrolPrivilegedApiRequestSingularDataSourceRepresentation = map[string]interface{}{
		"privileged_api_request_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apiaccesscontrol_privileged_api_request.test_privileged_api_request.id}`},
	}

	ApiaccesscontrolPrivilegedApiRequestDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_apiaccesscontrol_privileged_api_request.test_privileged_api_request.id}`},
		"resource_id":    acctest.Representation{RepType: acctest.Optional, Create: `${var.cluster_id}`},
		"resource_type":  acctest.Representation{RepType: acctest.Optional, Create: `resourceType`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: ApiaccesscontrolPrivilegedApiRequestDataSourceFilterRepresentation}}
	ApiaccesscontrolPrivilegedApiRequestDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_apiaccesscontrol_privileged_api_request.test_privileged_api_request.id}`}},
	}

	ApiaccesscontrolPrivilegedApiRequestRepresentation = map[string]interface{}{
		"privileged_operation_list":        acctest.RepresentationGroup{RepType: acctest.Required, Group: ApiaccesscontrolPrivilegedApiRequestPrivilegedOperationListRepresentation},
		"reason_summary":                   acctest.Representation{RepType: acctest.Required, Create: `TerraformTestPrivilegedApiControl`},
		"resource_id":                      acctest.Representation{RepType: acctest.Required, Create: `${var.cluster_id}`},
		"compartment_id":                   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"duration_in_hrs":                  acctest.Representation{RepType: acctest.Required, Create: `1`},
		"notification_topic_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.topic_id}`},
		"reason_detail":                    acctest.Representation{RepType: acctest.Required, Create: `reasonDetail`},
		"severity":                         acctest.Representation{RepType: acctest.Required, Create: `SEV_3`},
		"sub_resource_name_list":           acctest.Representation{RepType: acctest.Optional, Create: []string{`subResourceNameList`}},
		"ticket_numbers":                   acctest.Representation{RepType: acctest.Required, Create: []string{`ticketNumbers`}},
		"time_requested_for_future_access": acctest.Representation{RepType: acctest.Optional, Create: ``},
	}
	ApiaccesscontrolPrivilegedApiRequestPrivilegedOperationListRepresentation = map[string]interface{}{
		"api_name":        acctest.Representation{RepType: acctest.Required, Create: `UpdateVmCluster`},
		"attribute_names": acctest.Representation{RepType: acctest.Optional, Create: []string{`cpuCoreCount`}},
	}

	//ApiaccesscontrolPrivilegedApiRequestResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_apigateway_api", "test_api", acctest.Required, acctest.Create, ApigatewayApiRepresentation) +
	//	acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_guard_resources", "test_resources", acctest.Required, acctest.Create, CloudGuardResourceDataSourceRepresentation) +
	//	DefinedTagsDependencies +
	//	acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, OnsNotificationTopicRepresentation)
)

// issue-routing-tag: apiaccesscontrol/default
func TestApiaccesscontrolPrivilegedApiRequestResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApiaccesscontrolPrivilegedApiRequestResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()
	resourceId := utils.GetEnvSettingWithBlankDefault("resource_id")
	resourceIdVariableStr := fmt.Sprintf("variable \"resource_id\" { default = \"%s\" }\n", resourceId)
	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	onsTopicId := utils.GetEnvSettingWithBlankDefault("topic_id")
	onsTopicIdVariableStr := fmt.Sprintf("variable \"topic_id\" { default = \"%s\" }\n", onsTopicId)
	clusterId := utils.GetEnvSettingWithBlankDefault("cluster_id")
	clusterIdVariableStr := fmt.Sprintf("variable \"cluster_id\" { default = \"%s\" }\n", clusterId)
	controlResourceName := "oci_apiaccesscontrol_privileged_api_control.test_privileged_api_control"
	resourceName := "oci_apiaccesscontrol_privileged_api_request.test_privileged_api_request"
	datasourceName := "data.oci_apiaccesscontrol_privileged_api_requests.test_privileged_api_requests"
	singularDatasourceName := "data.oci_apiaccesscontrol_privileged_api_request.test_privileged_api_request"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+clusterIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_apiaccesscontrol_privileged_api_request", "test_privileged_api_request", acctest.Optional, acctest.Create, ApiaccesscontrolPrivilegedApiRequestRepresentation), "apiaccesscontrol", "privilegedApiRequest", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		//Prerequisite control creation
		{
			Config: config + compartmentIdVariableStr + resourceIdVariableStr + onsTopicIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_apiaccesscontrol_privileged_api_control", "test_privileged_api_control", acctest.Required, acctest.Create, ApiaccesscontrolPrivilegedApiControlRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(controlResourceName, "approver_group_id_list.#", "1"),
				resource.TestCheckResourceAttr(controlResourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(controlResourceName, "notification_topic_id"),
				resource.TestCheckResourceAttr(controlResourceName, "privileged_operation_list.#", "1"),
				resource.TestCheckResourceAttrSet(controlResourceName, "privileged_operation_list.0.api_name"),
				resource.TestCheckResourceAttr(controlResourceName, "resource_type", "EXADATAINFRASTRUCTURE"),
				resource.TestCheckResourceAttr(controlResourceName, "resources.#", "1"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, controlResourceName, "id")
					if err != nil {
						return err
					}
					err2 := dummyDeleteCallApiaccesscontrolPrivilegedApiControlResource(compartmentId, resId)
					if err2 != nil {
						return err2
					}
					return approveApiaccesscontrolPrivilegedApiRequestResources(compartmentId, resId)

				},
			),
		},

		// verify Create
		{
			Config: config + compartmentIdVariableStr + resourceIdVariableStr + onsTopicIdVariableStr + clusterIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_apiaccesscontrol_privileged_api_request", "test_privileged_api_request", acctest.Required, acctest.Create, ApiaccesscontrolPrivilegedApiRequestRepresentation) + acctest.GenerateResourceFromRepresentationMap("oci_apiaccesscontrol_privileged_api_control", "test_privileged_api_control", acctest.Required, acctest.Create, ApiaccesscontrolPrivilegedApiControlRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "privileged_operation_list.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "privileged_operation_list.0.api_name"),
				resource.TestCheckResourceAttr(resourceName, "reason_summary", "TerraformTestPrivilegedApiControl"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_id"),
			),
		},

		//// delete before next Create
		//{
		//	Config: config + compartmentIdVariableStr + resourceIdVariableStr + onsTopicIdVariableStr + clusterIdVariableStr,
		//},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + resourceIdVariableStr + onsTopicIdVariableStr + clusterIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_apiaccesscontrol_privileged_api_request", "test_privileged_api_request", acctest.Optional, acctest.Create, ApiaccesscontrolPrivilegedApiRequestRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_apiaccesscontrol_privileged_api_control", "test_privileged_api_control", acctest.Required, acctest.Create, ApiaccesscontrolPrivilegedApiControlRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "duration_in_hrs", "1"),
				//resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "notification_topic_id"),
				resource.TestCheckResourceAttr(resourceName, "privileged_operation_list.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "privileged_operation_list.0.api_name"),
				resource.TestCheckResourceAttr(resourceName, "privileged_operation_list.0.attribute_names.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "reason_detail", "reasonDetail"),
				resource.TestCheckResourceAttr(resourceName, "reason_summary", "TerraformTestPrivilegedApiControl"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_id"),
				resource.TestCheckResourceAttr(resourceName, "severity", "SEV_3"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "sub_resource_name_list.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "ticket_numbers.#", "1"),

				func(s *terraform.State) (err error) {
					return rejectApiaccesscontrolPrivilegedApiRequestResources(compartmentId, clusterId)
					//resId, err = acctest.FromInstanceState(s, resourceName, "id")
					//if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
					//	if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
					//		return errExport
					//	}
					//}
					//return err
				},
			),
		},

		// verify datasource
		{
			Config: config + onsTopicIdVariableStr + resourceIdVariableStr + clusterIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apiaccesscontrol_privileged_api_requests", "test_privileged_api_requests", acctest.Optional, acctest.Update, ApiaccesscontrolPrivilegedApiRequestDataSourceRepresentation) +
				compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_apiaccesscontrol_privileged_api_request", "test_privileged_api_request", acctest.Optional, acctest.Update, ApiaccesscontrolPrivilegedApiRequestRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_id"),
				resource.TestCheckResourceAttr(datasourceName, "resource_type", "resourceType"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "privileged_api_request_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "privileged_api_request_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config + onsTopicIdVariableStr + resourceIdVariableStr + clusterIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apiaccesscontrol_privileged_api_request", "test_privileged_api_request", acctest.Required, acctest.Create, ApiaccesscontrolPrivilegedApiRequestSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ApiaccesscontrolPrivilegedApiRequestResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "privileged_api_request_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "approver_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "duration_in_hrs", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "entity_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "number_of_approvers_required"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "privileged_api_control_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "privileged_api_control_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "privileged_operation_list.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "privileged_operation_list.0.attribute_names.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "reason_detail", "reasonDetail"),
				resource.TestCheckResourceAttr(singularDatasourceName, "reason_summary", "TerraformTestPrivilegedApiControl"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "request_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "requested_by.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "severity", "SEV_3"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "sub_resource_name_list.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ticket_numbers.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + ApiaccesscontrolPrivilegedApiRequestRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func rejectApiaccesscontrolPrivilegedApiRequestResources(compartment string, resourceId string) error {
	privilegedApiRequestClient := acctest.GetTestClients(&schema.ResourceData{}).PrivilegedApiRequestsClient()
	privilegedApiRequestIds, err := getApiaccesscontrolPrivilegedApiRequestIds(compartment, resourceId)
	if err != nil {
		return err
	}
	for _, privilegedApiRequestId := range privilegedApiRequestIds {

		rejectPrivilegedApiRequest := oci_apiaccesscontrol.RejectPrivilegedApiRequestRequest{}

		rejectPrivilegedApiRequest.PrivilegedApiRequestId = &privilegedApiRequestId
		rejectPrivilegedApiRequest.RejectPrivilegedApiRequestDetails = oci_apiaccesscontrol.RejectPrivilegedApiRequestDetails{
			ApproverComment: common.String("Terraform Rejected"),
			// Set other fields if needed
		}
		rejectPrivilegedApiRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "apiaccesscontrol")
		_, error := privilegedApiRequestClient.RejectPrivilegedApiRequest(context.Background(), rejectPrivilegedApiRequest)
		if error != nil {
			fmt.Printf("Error rejecting privilegedApiRequestId %s %s, It is possible that the resource is already approved. Please verify manually \n", privilegedApiRequestId, error)
			continue
		}
		//acctest.WaitTillCondition(acctest.TestAccProvider, &privilegedApiRequestId, ApiaccesscontrolPrivilegedApiRequestApprovedWaitCondition, time.Duration(3*time.Minute),
		//	ApiaccesscontrolPrivilegedApiRequestResponseFetchOperation, "apiaccesscontrol", true)

	}
	return nil
}
