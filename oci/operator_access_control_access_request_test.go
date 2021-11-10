// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	accessRequestSingularDataSourceRepresentation = map[string]interface{}{
		"access_request_id": Representation{RepType: Required, Create: `${data.oci_operator_access_control_access_requests.test_access_requests.access_request_collection.0.items.0.id}`},
	}

	accessRequestDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
	}

	AccessRequestResourceConfig = ""
)

// issue-routing-tag: operator_access_control/default
func TestOperatorAccessControlAccessRequestResource_basic(t *testing.T) {
	t.Skip("Access Requests are created outside customer api. Access requests may not be available all the time")
	httpreplay.SetScenario("TestOperatorAccessControlAccessRequestResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_operator_access_control_access_requests.test_access_requests"
	singularDatasourceName := "data.oci_operator_access_control_access_request.test_access_request"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource step
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_operator_access_control_access_requests", "test_access_requests", Required, Create, accessRequestDataSourceRepresentation) +
				compartmentIdVariableStr + AccessRequestResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "access_request_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "access_request_collection.0.items.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_operator_access_control_access_requests", "test_access_requests", Required, Create, accessRequestDataSourceRepresentation) +
				GenerateDataSourceFromRepresentationMap("oci_operator_access_control_access_request", "test_access_request", Required, Create, accessRequestSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AccessRequestResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "access_reason_summary"),
				resource.TestCheckResourceAttr(singularDatasourceName, "action_requests_list.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "approver_comment"),
				resource.TestCheckResourceAttr(singularDatasourceName, "audit_type.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "duration"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "extend_duration"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_auto_approved"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "opctl_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "opctl_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "reason"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "request_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "resource_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "severity"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "system_message"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_of_creation"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_of_modification"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "workflow_id.#", "1"),
			),
		},
	})
}
