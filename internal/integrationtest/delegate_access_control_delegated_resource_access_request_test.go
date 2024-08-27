// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DelegateAccessControlDelegatedResourceAccessRequestSingularDataSourceRepresentation = map[string]interface{}{
		"delegated_resource_access_request_id": acctest.Representation{RepType: acctest.Required, Create: `${var.getAccReqId}`},
	}

	DelegateAccessControlDelegatedResourceAccessRequestDataSourceRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"delegation_control_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.delCtrlId}`},
		"request_status":        acctest.Representation{RepType: acctest.Optional, Create: `CREATED`},
		//"resource_id":           acctest.Representation{RepType: acctest.Optional, Create: `${oci_cloud_guard_resource.test_resource.id}`},
		"state":      acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"time_end":   acctest.Representation{RepType: acctest.Optional, Create: `timeEnd`},
		"time_start": acctest.Representation{RepType: acctest.Optional, Create: `timeStart`},
	}

	/*DelegateAccessControlDelegatedResourceAccessRequestResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_guard_resources", "test_resources", acctest.Required, acctest.Create, CloudGuardResourceDataSourceRepresentation) +
	acctest.GenerateResourceFromRepresentationMap("oci_delegate_access_control_delegation_control", "test_delegation_control", acctest.Required, acctest.Create, DelegateAccessControlDelegationControlRepresentation) +
	acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, OnsNotificationTopicRepresentation)*/
)

// issue-routing-tag: delegate_access_control/default
func TestDelegateAccessControlDelegatedResourceAccessRequestResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDelegateAccessControlDelegatedResourceAccessRequestResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_delegate_access_control_delegated_resource_access_requests.test_delegated_resource_access_requests"
	singularDatasourceName := "data.oci_delegate_access_control_delegated_resource_access_request.test_delegated_resource_access_request"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_delegate_access_control_delegated_resource_access_requests", "test_delegated_resource_access_requests", acctest.Required, acctest.Create, DelegateAccessControlDelegatedResourceAccessRequestDataSourceRepresentation) +
				compartmentIdVariableStr, /* + DelegateAccessControlDelegatedResourceAccessRequestResourceConfig*/
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				// Prakash, items collection,below attributes will not exist
				/*resource.TestCheckResourceAttrSet(datasourceName, "delegation_control_id"),
				resource.TestCheckResourceAttr(datasourceName, "request_status", "CREATED"),
				resource.TestCheckResourceAttrSet(datasourceName, "resource_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_end"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_start"),
				*/
				resource.TestCheckResourceAttrSet(datasourceName, "delegated_resource_access_request_summary_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_delegate_access_control_delegated_resource_access_request", "test_delegated_resource_access_request", acctest.Required, acctest.Create, DelegateAccessControlDelegatedResourceAccessRequestSingularDataSourceRepresentation) +
				compartmentIdVariableStr, /* + DelegateAccessControlDelegatedResourceAccessRequestResourceConfig*/
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "delegated_resource_access_request_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "approval_info.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "audit_types.#", "1"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "closure_comment"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "database_name_list.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "delegation_subscription_ids.#", "2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "duration_in_hours"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "extend_duration_in_hours"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_auto_approved"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_pending_more_info"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "num_extension_approvals"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "num_initial_approvals"),
				resource.TestCheckResourceAttr(singularDatasourceName, "provided_service_types.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "reason_for_request"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "request_status"),
				resource.TestCheckResourceAttr(singularDatasourceName, "requested_action_names.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "requester_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "severity"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ticket_numbers.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_access_requested"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}
